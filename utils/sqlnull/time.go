package sqlnull

import (
	"time"

	"github.com/lib/pq"
)

type Time struct {
	pq.NullTime
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 3 && string(b) == "null" {
		t.Valid = false
		return nil
	}
	t.Time = time.Time{}
	err := t.Time.UnmarshalJSON(b)
	if err != nil {
		t.Valid = false
		return err
	}
	t.Valid = true
	return nil
}
