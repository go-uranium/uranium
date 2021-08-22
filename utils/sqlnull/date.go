package sqlnull

import (
	"database/sql"
	"database/sql/driver"

	"github.com/go-uranium/uranium/utils/type/date"
)

type Date struct {
	Date  *date.Date
	Valid bool
}

func (d *Date) Scan(value interface{}) error {
	nullint := sql.NullInt32{}
	err := nullint.Scan(value)
	if err != nil {
		return err
	}

	if !nullint.Valid {
		d.Valid = false
		return nil
	}
	d.Date = date.New(nullint.Int32)
	d.Valid = true
	return nil
}

// Value implements the driver Valuer interface.
func (d *Date) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}
	return int64(d.Date.Encode()), nil
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
