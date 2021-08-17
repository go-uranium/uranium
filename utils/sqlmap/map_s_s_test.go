package sqlmap_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-uranium/uranium/utils/sqlmap"
)

func TestMapStringString_MarshalJSON(t *testing.T) {
	type T struct {
		A string
		B int
		C *sqlmap.StringString
	}
	data := &T{
		A: "aaaa",
		B: 2021,
		C: sqlmap.NewMapSS(nil),
	}
	_, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}

	data.C = sqlmap.NewMapSS(map[string]string{
		"github":  "@iochen",
		"website": "https://iochen.com/",
		"twitter": "@realRichardChen",
	})

	_, err = json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}

	data.C = &sqlmap.StringString{}
	_, err = json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestName(t *testing.T) {
	bytes, err := json.Marshal(`a\a\ad\aaa`)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))
}
