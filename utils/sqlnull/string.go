package sqlnull

import (
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
}

func (s *String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *String) UnmarshalJSON(b []byte) error {
	if len(b) == 3 && string(b) == "null" {
		s.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &s.String)
	if err != nil {
		s.Valid = false
		return err
	}
	s.Valid = true
	return nil
}
