package sqlnull

import (
	"github.com/go-uranium/uranium/utils/type/date"
)

type Date struct {
	Date  *date.Date
	Valid bool
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}
	return d.Date.MarshalJSON()
}

func (d *Date) UnmarshalJSON(b []byte) error {
	if len(b) == 3 && string(b) == "null" {
		d.Valid = false
		return nil
	}
	d.Date = &date.Date{}
	err := d.Date.UnmarshalJSON(b)
	if err != nil {
		d.Valid = false
		return err
	}
	d.Valid = true
	return nil
}
