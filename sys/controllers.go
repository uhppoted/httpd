package system

import (
	//	"encoding/json"
	//	"fmt"
	//	"io/ioutil"
	//	"log"
	//	"net/http"
	//	"os"
	//	"path/filepath"
	//	"strings"
	"time"
	//	core "github.com/uhppoted/uhppote-core/types"
	//	"github.com/uhppoted/uhppoted-api/acl"
	"github.com/uhppoted/uhppoted-httpd/types"
)

type Controller struct {
	ID       string           `json:"ID"`
	Created  time.Time        `json:"created"`
	Name     *types.Name      `json:"name"`
	DeviceID uint32           `json:"device-id"`
	IP       address          `json:"address"`
	Doors    map[uint8]string `json:"-"`
	TimeZone string           `json:"timezone"`
}

// type datetime struct {
// 	DateTime *types.DateTime
// 	TimeZone *time.Location
// 	Status   status
// }
//
// type records uint32
//
// func (r *records) String() string {
// 	if r != nil {
// 		return fmt.Sprintf("%v", uint32(*r))
// 	}
//
// 	return ""
// }
//
// type system struct {
// 	file  string
// 	Doors map[string]types.Door `json:"doors"`
// 	Local []*Local              `json:"local"`
// }
//
// func (s *system) refresh() {
// 	if s != nil {
// 		for _, l := range s.Local {
// 			go l.refresh()
// 		}
// 	}
// }
//
// var sys = system{
// 	Doors: map[string]types.Door{},
// 	Local: []*Local{},
// }
//
// func init() {
// 	go func() {
// 		c := time.Tick(30 * time.Second)
// 		for _ = range c {
// 			sys.refresh()
// 		}
// 	}()
// }
//
// func Init(conf string) error {
// 	bytes, err := ioutil.ReadFile(conf)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = json.Unmarshal(bytes, &sys)
// 	if err != nil {
// 		return err
// 	}
//
// 	sys.file = conf
//
// 	for _, l := range sys.Local {
// 		l.Init()
// 	}
//
// 	return nil
// }
//
// func System() interface{} {
// 	controllers := []Controller{}
//
// 	for _, l := range sys.Local {
// 		controllers = append(controllers, l.Controllers()...)
// 	}
//
// 	for _, c := range controllers {
// 		for _, d := range sys.Doors {
// 			if d.ControllerID == c.DeviceID {
// 				c.Doors[d.Door] = d.Name
// 			}
// 		}
// 	}
//
// 	return struct {
// 		Controllers []Controller
// 	}{
// 		Controllers: controllers,
// 	}
// }
//
// func Update(permissions []types.Permissions) {
// 	for _, l := range sys.Local {
// 		l.Update(permissions)
// 	}
// }
//
// //func Post(m map[string]interface{}, auth db.IAuth) (interface{}, error) {
// func Post(m map[string]interface{}) (interface{}, error) {
// 	//	d.Lock()
// 	//
// 	//	defer d.Unlock()
//
// 	// add/update ?
//
// 	controllers, err := unpack(m)
// 	if err != nil {
// 		return nil, &types.HttpdError{
// 			Status: http.StatusBadRequest,
// 			Err:    fmt.Errorf("Invalid request"),
// 			Detail: fmt.Errorf("Error unpacking 'post' request (%w)", err),
// 		}
// 	}
//
// 	list := struct {
// 		Updated []interface{} `json:"updated"`
// 		Deleted []interface{} `json:"deleted"`
// 	}{}
//
// loop:
// 	for _, c := range controllers {
// 		if c.ID == "" {
// 			return nil, &types.HttpdError{
// 				Status: http.StatusBadRequest,
// 				Err:    fmt.Errorf("Invalid controller ID"),
// 				Detail: fmt.Errorf("Invalid 'post' request (%w)", fmt.Errorf("Invalid controller ID '%v'", c.ID)),
// 			}
// 		}
//
// 		for _, l := range sys.Local {
// 			for k, d := range l.Devices {
// 				if ID(k) == c.ID {
// 					//				if c.Name != nil && *c.Name == "" && c.Card != nil && *c.Card == 0 {
// 					//					if r, err := d.delete(shadow, c, auth); err != nil {
// 					//						return nil, err
// 					//					} else if r != nil {
// 					//						list.Deleted = append(list.Deleted, r)
// 					//						continue loop
// 					//					}
// 					//				}
//
// 					//					if r, err := d.update(shadow, c, auth); err != nil {
// 					//						return nil, err
// 					//					} else if r != nil {
// 					//						list.Updated = append(list.Updated, r)
// 					//					}
//
// 					if c.Name != nil {
// 						d.Name = c.Name.String()
// 					}
//
// 					if cc := l.Controller(k); cc != nil {
// 						list.Updated = append(list.Updated, cc)
// 					}
//
// 					continue loop
// 				}
// 			}
// 		}
//
// 		//		if r, err := d.add(shadow, c, auth); err != nil {
// 		//			return nil, err
// 		//		} else if r != nil {
// 		//			list.Updated = append(list.Updated, r)
// 		//		}
// 	}
//
// 	if err := save(&sys, sys.file); err != nil {
// 		return nil, err
// 	}
//
// 	//	d.data = *shadow
//
// 	return list, nil
// }
//
// func save(s *system, file string) error {
// 	//	if err := validate(s); err != nil {
// 	//		return err
// 	//	}
// 	//
// 	//	if err := clean(s); err != nil {
// 	//		return err
// 	//	}
//
// 	if file == "" {
// 		return nil
// 	}
//
// 	b, err := json.MarshalIndent(s, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	tmp, err := ioutil.TempFile(os.TempDir(), "uhppoted-system.json")
// 	if err != nil {
// 		return err
// 	}
//
// 	defer os.Remove(tmp.Name())
//
// 	if _, err := tmp.Write(b); err != nil {
// 		return err
// 	}
//
// 	if err := tmp.Close(); err != nil {
// 		return err
// 	}
//
// 	if err := os.MkdirAll(filepath.Dir(file), 0770); err != nil {
// 		return err
// 	}
//
// 	return os.Rename(tmp.Name(), file)
// }
//
// func consolidate(list []types.Permissions) (*acl.ACL, error) {
// 	// initialise empty ACL
// 	acl := make(acl.ACL)
//
// 	for _, d := range sys.Doors {
// 		if _, ok := acl[d.ControllerID]; !ok {
// 			acl[d.ControllerID] = make(map[uint32]core.Card)
// 		}
// 	}
//
// 	// create ACL with all cards on all controllers
// 	for _, p := range list {
// 		for _, l := range acl {
// 			if _, ok := l[p.CardNumber]; !ok {
// 				from := core.Date(p.From)
// 				to := core.Date(p.To)
//
// 				l[p.CardNumber] = core.Card{
// 					CardNumber: p.CardNumber,
// 					From:       &from,
// 					To:         &to,
// 					Doors:      map[uint8]bool{1: false, 2: false, 3: false, 4: false},
// 				}
// 			}
// 		}
// 	}
//
// 	// update ACL cards from permissions
// 	for _, p := range list {
// 		for _, d := range p.Doors {
// 			if door, ok := sys.Doors[d]; !ok {
// 				log.Printf("WARN %v", fmt.Errorf("Invalid door %v for card %v", d, p.CardNumber))
// 			} else if l, ok := acl[door.ControllerID]; !ok {
// 				log.Printf("WARN %v", fmt.Errorf("Door %v - invalid configuration (no controller defined for  %v)", d, door.ControllerID))
// 			} else if card, ok := l[p.CardNumber]; !ok {
// 				log.Printf("WARN %v", fmt.Errorf("Card %v not initialised for controller %v", p.CardNumber, door.ControllerID))
// 			} else {
// 				card.Doors[door.Door] = true
// 			}
// 		}
// 	}
//
// 	return &acl, nil
// }
//
// func unpack(m map[string]interface{}) ([]Controller, error) {
// 	o := struct {
// 		Controllers []struct {
// 			ID   string
// 			Name *string
// 		}
// 	}{}
//
// 	blob, err := json.Marshal(m)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := json.Unmarshal(blob, &o); err != nil {
// 		return nil, err
// 	}
//
// 	controllers := []Controller{}
//
// 	for _, r := range o.Controllers {
// 		record := Controller{
// 			ID: strings.TrimSpace(r.ID),
// 		}
//
// 		if r.Name != nil {
// 			name := types.Name(*r.Name)
// 			record.Name = &name
// 		}
//
// 		controllers = append(controllers, record)
// 	}
//
// 	return controllers, nil
// }
//
// func clean(s string) string {
// 	return strings.ReplaceAll(strings.ToLower(s), " ", "")
// }
//
// func warn(err error) {
// 	log.Printf("ERROR %v", err)
// }
