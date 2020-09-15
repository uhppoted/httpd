package memdb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/uhppoted/uhppoted-httpd/db"
	"github.com/uhppoted/uhppoted-httpd/sys"
)

type fdb struct {
	sync.RWMutex
	file string
	data data
}

type data struct {
	Tables tables `json:"tables"`
}

type tables struct {
	Doors       []*db.Door       `json:"doors"`
	Groups      []*db.Group      `json:"groups"`
	CardHolders []*db.CardHolder `json:"cardholders"`
}

func (d *data) copy() *data {
	shadow := data{
		Tables: tables{
			Doors:       make([]*db.Door, len(d.Tables.Doors)),
			Groups:      make([]*db.Group, len(d.Tables.Groups)),
			CardHolders: make([]*db.CardHolder, len(d.Tables.CardHolders)),
		},
	}

	for i, v := range d.Tables.Doors {
		shadow.Tables.Doors[i] = v.Copy()
	}

	for i, v := range d.Tables.Groups {
		shadow.Tables.Groups[i] = v.Copy()
	}

	for i, v := range d.Tables.CardHolders {
		shadow.Tables.CardHolders[i] = v.Copy()
	}

	return &shadow
}

func NewDB(file string) (*fdb, error) {
	f := fdb{
		file: file,
		data: data{
			Tables: tables{
				Groups:      []*db.Group{},
				CardHolders: []*db.CardHolder{},
			},
		},
	}

	if err := load(&f.data, f.file); err != nil {
		return nil, err
	}

	return &f, nil
}

func (d *fdb) Groups() []*db.Group {
	d.RLock()

	defer d.RUnlock()

	return d.data.Tables.Groups
}

func (d *fdb) CardHolders() ([]*db.CardHolder, error) {
	d.RLock()

	defer d.RUnlock()

	return d.data.Tables.CardHolders, nil
}

func (d *fdb) ACL() ([]system.Permissions, error) {
	d.RLock()

	defer d.RUnlock()

	list := []system.Permissions{}

	for _, c := range d.data.Tables.CardHolders {
		doors := []string{}

		for _, p := range c.Groups {
			if p.Value {
				for _, group := range d.data.Tables.Groups {
					if p.GID == group.ID {
						for _, doorID := range group.Doors {
							for _, door := range d.data.Tables.Doors {
								if doorID == door.ID {
									doors = append(doors, door.DoorID)
								}
							}
						}
					}
				}
			}
		}

		list = append(list, system.Permissions{
			CardNumber: c.CardNumber,
			From:       c.From,
			To:         c.To,
			Doors:      doors,
		})
	}

	return list, nil
}

func (d *fdb) Update(u map[string]interface{}) (interface{}, error) {
	d.Lock()

	defer d.Unlock()

	list := struct {
		Updated map[string]interface{} `json:"updated"`
	}{
		Updated: map[string]interface{}{},
	}

	shadow := d.data.copy()

	for k, v := range u {
		gid := k

		if value, ok := v.(bool); ok {
			for _, c := range shadow.Tables.CardHolders {
				for _, g := range c.Groups {
					if g.ID == gid {
						g.Value = value
						list.Updated[gid] = g.Value
					}
				}
			}
		}
	}

	if err := save(shadow, d.file); err != nil {
		return nil, err
	}

	d.data = *shadow

	return list, nil
}

func save(data interface{}, file string) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	tmp, err := ioutil.TempFile(os.TempDir(), "uhppoted-*.db")
	if err != nil {
		return err
	}

	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(b); err != nil {
		return err
	}

	if err := tmp.Close(); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(file), 0770); err != nil {
		return err
	}

	return os.Rename(tmp.Name(), file)
}

func load(data interface{}, file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return json.Unmarshal(b, data)
}
