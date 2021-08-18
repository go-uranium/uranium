package date

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDate(t *testing.T) {
	type T struct {
		A int
		B string
		C Date
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
