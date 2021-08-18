package date

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var ErrExpectedJson = errors.New("unexpected type date.Date in json format.")

type Date struct {
	// safe between 0 to 32767
	// but should be 0 to 9999
	Year uint16 `json:"year"`
	// safe between 0 to 255
	// but should be 0 to 12
	Month uint8 `json:"month"`
	// safe between 0 to 255
	// but should be 0 to 31
	Day uint8 `json:"day"`
}

// Now returns a current Date struct.
func Now() Date {
	now := time.Now()
	return Date{
		Year:  uint16(now.Year()),
		Month: uint8(now.Month()),
		Day:   uint8(now.Day()),
	}
}

// New decodes the int32 type to a date.
func New(c int32) Date {
	return Date{
		Year:  uint16((uint32(c) & 0xFFFF0000) >> 16),
		Month: uint8((c & 0xFF00) >> 8),
		Day:   uint8(c & 0xFF),
	}
}

// Encode encodes the date to an int32 type
func (d *Date) Encode() int32 {
	return (int32(d.Year) << 16) & (int32(d.Month) << 8) & int32(d.Day)
}

func (d *Date) JSONStringJoin() ([]byte, error) {
	return d.marshalJSONStringJoin()
}

func (d *Date) UnmarshalJSON(b []byte) error {
	parts := bytes.Split(b, []byte(":"))
	if len(parts) != 4 {
		return ErrExpectedJson
	}
	yearEnd := bytes.IndexByte(parts[1], ',')
	year, err := strconv.Atoi(string(bytes.TrimSpace(parts[1][:yearEnd])))
	if err != nil {
		return err
	}
	d.Year = uint16(year)

	monthEnd := bytes.IndexByte(parts[2], ',')
	month, err := strconv.Atoi(string(bytes.TrimSpace(parts[2][:monthEnd])))
	if err != nil {
		return err
	}
	d.Month = uint8(month)

	dayEnd := bytes.IndexByte(parts[3], '}')
	day, err := strconv.Atoi(string(bytes.TrimSpace(parts[3][:dayEnd])))
	if err != nil {
		return err
	}
	d.Day = uint8(day)
	return nil
}

func (d *Date) marshalJSONStringJoin() ([]byte, error) {
	return []byte(`{"year":` + strconv.Itoa(int(d.Year)) + `,"month":` +
		strconv.Itoa(int(d.Month)) + `,"day":` + strconv.Itoa(int(d.Day)) + `}`), nil
}

func (d *Date) marshalJSONFmt() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"year":%d,"month":%d,"day":%d}`, d.Year, d.Month, d.Day)), nil
}
