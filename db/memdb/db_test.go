package memdb

import (
//	"reflect"
//	"testing"
//	"time"
//
//	"github.com/uhppoted/uhppoted-httpd/db"
//	"github.com/uhppoted/uhppoted-httpd/types"
)

//func date(s string) *types.Date {
//	date, _ := time.ParseInLocation("2006-01-02", s, time.Local)
//	d := types.Date(date)
//
//	return &d
//}
//
//func TestCardNumberUpdate(t *testing.T) {
//	name := types.Name("Hagrid")
//	card := types.Card(6514231)
//
//	dbt := fdb{
//		data: data{
//			Tables: tables{
//				Groups: types.Groups{},
//				CardHolders: []*db.CardHolder{
//					&db.CardHolder{
//						ID:     "C01",
//						Name:   &name,
//						Card:   &card,
//						From:   date("2021-01-02"),
//						To:     date("2021-12-31"),
//						Groups: []*db.Permission{},
//					},
//				},
//			},
//		},
//	}
//
//	u := map[string]interface{}{
//		"CARD01": "1234567",
//	}
//
//	expected := struct {
//		Updated map[string]interface{} `json:"updated"`
//	}{
//		Updated: map[string]interface{}{
//			"CARD01": uint32(1234567),
//		},
//	}
//
//	hagrid := types.Name("Hagrid")
//	c1234567 := types.Card(1234567)
//
//	updated := fdb{
//		data: data{
//			Tables: tables{
//				Groups: types.Groups{},
//				CardHolders: []*db.CardHolder{
//					&db.CardHolder{
//						ID:     "C01",
//						Name:   &hagrid,
//						Card:   &c1234567,
//						From:   date("2021-01-02"),
//						To:     date("2021-12-31"),
//						Groups: []*db.Permission{},
//					},
//				},
//			},
//		},
//	}
//
//	result, err := dbt.update("", u)
//	if err != nil {
//		t.Fatalf("Unexpected error updating DB: %v", err)
//	}
//
//	if !reflect.DeepEqual(result, expected) {
//		t.Errorf("Incorrect return value: expected:%#v, got:%#v", expected, result)
//	}
//
//	if !reflect.DeepEqual(dbt.data, updated.data) {
//		t.Errorf("DB updated incorrectly: expected:%v, got:%v", updated.data, dbt.data)
//	}
//}
//
//func TestDuplicateCardNumberUpdate(t *testing.T) {
//	dbt := fdb{
//		data: data{
//			Tables: tables{
//				Groups: []*db.Group{},
//				CardHolders: []*db.CardHolder{
//					&db.CardHolder{
//						ID:     "C01",
//						Name:   &db.Name{ID: "C01.name", Name: "Hagrid"},
//						Card:   &db.Card{ID: "CARD01", Number: 6514231},
//						From:   &types.Date{ID: "C01.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C01.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//
//					&db.CardHolder{
//						ID:     "C02",
//						Name:   &db.Name{ID: "C02.name", Name: "Dobby"},
//						Card:   &db.Card{ID: "CARD02", Number: 1234567},
//						From:   &types.Date{ID: "C02.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C02.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{}},
//				},
//			},
//		},
//	}
//
//	u := map[string]interface{}{
//		"CARD01": "1234567",
//	}
//
//	updated := fdb{
//		data: data{
//			Tables: tables{
//				Groups: []*db.Group{},
//				CardHolders: []*db.CardHolder{
//					&db.CardHolder{
//						ID:     "C01",
//						Name:   &db.Name{ID: "C01.name", Name: "Hagrid"},
//						Card:   &db.Card{ID: "CARD01", Number: 6514231},
//						From:   &types.Date{ID: "C01.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C01.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//
//					&db.CardHolder{
//						ID:     "C02",
//						Name:   &db.Name{ID: "C02.name", Name: "Dobby"},
//						Card:   &db.Card{ID: "CARD02", Number: 1234567},
//						From:   &types.Date{ID: "C02.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C02.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//				},
//			},
//		},
//	}
//
//	result, err := dbt.update("", u)
//	if err == nil {
//		t.Errorf("Expected error updating DB, got %v", err)
//	}
//
//	if result != nil {
//		t.Errorf("Incorrect return value: expected:%#v, got:%#v", nil, result)
//	}
//
//	if !reflect.DeepEqual(dbt.data, updated.data) {
//		t.Errorf("DB updated incorrectly: expected:%v, got:%v", updated.data, dbt.data)
//	}
//}
//
//func TestCardNumberSwap(t *testing.T) {
//	dbt := fdb{
//		data: data{
//			Tables: tables{
//				Groups: []*db.Group{},
//				CardHolders: []*db.CardHolder{
//					&db.CardHolder{
//						ID:     "C01",
//						Name:   &db.Name{ID: "C01.name", Name: "Hagrid"},
//						Card:   &db.Card{ID: "CARD01", Number: 6514231},
//						From:   &types.Date{ID: "C01.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C01.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//
//					&db.CardHolder{
//						ID:     "C02",
//						Card:   &db.Card{ID: "CARD02", Number: 6514231},
//						Name:   &db.Name{ID: "C02.name", Name: "Dobby"},
//						From:   &types.Date{ID: "C02.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C02.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//				},
//			},
//		},
//	}
//
//	u := map[string]interface{}{
//		"CARD01": "1234567",
//		"CARD02": "6514231",
//	}
//
//	expected := struct {
//		Updated map[string]interface{} `json:"updated"`
//	}{
//		Updated: map[string]interface{}{
//			"CARD01": uint32(1234567),
//			"CARD02": uint32(6514231),
//		},
//	}
//
//	updated := fdb{
//		data: data{
//			Tables: tables{
//				Groups: []*db.Group{},
//				CardHolders: []*db.CardHolder{
//					&db.CardHolder{
//						ID:     "C01",
//						Name:   &db.Name{ID: "C01.name", Name: "Hagrid"},
//						Card:   &db.Card{ID: "CARD01", Number: 1234567},
//						From:   &types.Date{ID: "C01.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C01.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//
//					&db.CardHolder{
//						ID:     "C02",
//						Name:   &db.Name{ID: "C02.name", Name: "Dobby"},
//						Card:   &db.Card{ID: "CARD02", Number: 6514231},
//						From:   &types.Date{ID: "C02.from", Date: date("2021-01-02")},
//						To:     &types.Date{ID: "C02.to", Date: date("2021-12-31")},
//						Groups: []*db.Permission{},
//					},
//				},
//			},
//		},
//	}
//
//	result, err := dbt.update("", u)
//	if err != nil {
//		t.Fatalf("Unexpected error updating DB: %v", err)
//	}
//
//	if !reflect.DeepEqual(result, expected) {
//		t.Errorf("Incorrect return value: expected:%#v, got:%#v", expected, result)
//	}
//
//	if !reflect.DeepEqual(dbt.data, updated.data) {
//		t.Errorf("DB updated incorrectly: expected:%v, got:%v", updated.data, dbt.data)
//	}
//}
