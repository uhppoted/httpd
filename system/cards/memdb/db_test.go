package memdb

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/uhppoted/uhppoted-httpd/audit"
	"github.com/uhppoted/uhppoted-httpd/auth"
	"github.com/uhppoted/uhppoted-httpd/types"
)

var hagrid = cardholder("C01", "Hagrid", 6514231)
var dobby = cardholder("C02", "Dobby", 1234567, "G05")

type stub struct {
	write func(e audit.LogEntry)
}

func (x *stub) UID() string {
	return "stub"
}

func (x *stub) CanAddController(controller auth.Operant) error {
	return fmt.Errorf("not authorised")
}

func (x *stub) CanUpdateController(before, after auth.Operant) error {
	return fmt.Errorf("not authorised")
}

func (x *stub) CanDeleteController(controller auth.Operant) error {
	return fmt.Errorf("not authorised")
}

func (x *stub) CanAddCardHolder(cardHolder auth.Operant) error {
	return fmt.Errorf("not authorised")
}

func (x *stub) CanUpdateCardHolder(before, after auth.Operant) error {
	return fmt.Errorf("not authorised")
}

func (x *stub) CanDeleteCardHolder(cardHolder auth.Operant) error {
	return fmt.Errorf("not authorised")
}

func (x *stub) Write(e audit.LogEntry) {
	x.write(e)
}

func date(s string) *types.Date {
	date, _ := time.ParseInLocation("2006-01-02", s, time.Local)
	d := types.Date(date)

	return &d
}

func dbx(cardholders ...types.CardHolder) *fdb {
	p := fdb{
		data: data{
			Tables: tables{
				Groups: types.Groups{
					"G05": group("G05"),
				},
				CardHolders: types.CardHolders{},
			},
		},
	}

	for i, _ := range cardholders {
		c := cardholders[i].Clone()
		p.data.Tables.CardHolders[c.ID] = c
	}

	return &p
}

func group(id string) types.Group {
	return types.Group{
		ID:   id,
		Name: "",
	}
}

func cardholder(id, name string, card uint32, groups ...string) types.CardHolder {
	n := types.Name(name)
	c := types.Card(card)

	cardholder := types.CardHolder{
		ID:     id,
		Name:   &n,
		Card:   &c,
		From:   date("2021-01-02"),
		To:     date("2021-12-30"),
		Groups: map[string]bool{},
	}

	for _, g := range groups {
		cardholder.Groups[g] = true
	}

	return cardholder
}

func compare(got, expected interface{}, t *testing.T) {
	p, _ := json.Marshal(got)
	q, _ := json.Marshal(expected)

	if string(p) != string(q) {
		t.Errorf("'got' does not match 'expected'\nexpected:%s\ngot:     %s", string(q), string(p))
	}
}

func compareDB(db, expected *fdb, t *testing.T) {
	compare(db.data.Tables, expected.data.Tables, t)
}
