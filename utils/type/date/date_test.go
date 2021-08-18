package date

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDate_Now(t *testing.T) {
	type T struct {
		A int
		B string
		C *Date
	}
	data := &T{
		A: 2021,
		B: "hello",
		C: Now(),
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))

	bytes, err = json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))

	data1 := &T{}
	err = json.Unmarshal(bytes, data1)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(*data1)
}

func TestDate_Encode(t *testing.T) {
	d := Now()
	e := d.Encode()
	d1 := New(e)
	fmt.Println(d, e, d1)
}

func TestDate_Compare(t *testing.T) {
	d1 := &Date{
		Year:  2021,
		Month: 07,
		Day:   30,
	}

	d2 := &Date{
		Year:  2001,
		Month: 12,
		Day:   1,
	}

	d3 := &Date{
		Year:  2001,
		Month: 12,
		Day:   31,
	}

	if d1.Compare(d1) != 0 {
		t.Error("expected: 0")
	}
	if d3.Compare(d1) != -1 {
		t.Error("expected: 1")
	}
	if d3.Compare(d2) != 1 {
		t.Error("expected: -1")
	}
}

func BenchmarkDate_UnmarshalJSON(b *testing.B) {
	js, err := json.MarshalIndent(Now(), "", "\t")
	if err != nil {
		b.Error(err)
		return
	}
	d := &Date{}
	for n := 0; n < b.N; n++ {
		_ = d.UnmarshalJSON(js)
	}
}

func BenchmarkDate_marshalJSONFmt(b *testing.B) {
	d := Now()
	for n := 0; n < b.N; n++ {
		_, _ = d.marshalJSONFmt()
	}
}

func BenchmarkDate_marshalJSONStringJoin(b *testing.B) {
	d := Now()
	for n := 0; n < b.N; n++ {
		_, _ = d.marshalJSONStringJoin()
	}
}

func BenchmarkDate_jsonUnmarshal(b *testing.B) {
	js, err := json.MarshalIndent(Now(), "", "\t")
	if err != nil {
		b.Error(err)
		return
	}
	type D struct {
		Year  uint16 `json:"year"`
		Month uint8  `json:"month"`
		Day   uint8  `json:"day"`
	}
	d := &D{}
	for n := 0; n < b.N; n++ {
		_ = json.Unmarshal(js, d)
	}
}
