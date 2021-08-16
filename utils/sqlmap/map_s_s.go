package sqlmap

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// stores as a json text in database
// always not null
type MapStringString struct {
	mp      map[string]string
	mpReady bool
	js      []byte
	jsReady bool
}

func NewMapSS(m map[string]string) *MapStringString {
	if m == nil {
		m = map[string]string{}
	}
	return &MapStringString{
		mp:      m,
		mpReady: true,
	}
}

// Scan implements the Scanner interface.
func (mss *MapStringString) Scan(value interface{}) error {
	ns := &sql.NullString{}
	err := ns.Scan(value)
	if err != nil {
		return err
	}
	mss.js = []byte(ns.String)
	if !ns.Valid {
		mss.js = []byte("{}")
	}
	mss.jsReady = true
	return nil
}

// Value implements the driver Valuer interface.
func (mss *MapStringString) Value() (driver.Value, error) {
	if !mss.jsReady {
		if !mss.mpReady {
			mss.mp = map[string]string{}
			mss.mpReady = true
		}
		err := mss.jsFromMap()
		if err != nil {
			return nil, err
		}
	}
	ns := &sql.NullString{}
	ns.String = string(mss.js)
	ns.Valid = true
	return ns.Value()
}

func (mss *MapStringString) MarshalJSON() ([]byte, error) {
	// json is ready
	// map is/not ready
	if mss.jsReady {
		return mss.js, nil
	}

	// json is not ready
	// map is ready
	if mss.mpReady {
		err := mss.jsFromMap()
		if err != nil {
			return nil, err
		}
		return mss.js, nil
	}

	// json is not ready
	// map is not ready
	return []byte("{}"), nil
}

func (mss *MapStringString) UnmarshalJSON(b []byte) error {
	mss.js = b
	return nil
}

func (mss *MapStringString) Map() (map[string]string, error) {
	if mss.mpReady {
		return mss.mp, nil
	}
	if mss.jsReady {
		err := mss.mapFromJs()
		if err != nil {
			return nil, err
		}
		return mss.mp, nil
	}
	return map[string]string{}, nil
}

func (mss *MapStringString) jsFromMap() error {
	if !mss.mpReady {
		return errors.New("map status not ready")
	}
	js, err := json.Marshal(mss.mp)
	if err != nil {
		return err
	}
	mss.js = js
	mss.jsReady = true
	return nil
}

func (mss *MapStringString) mapFromJs() error {
	if !mss.jsReady {
		mss.mp = map[string]string{}
		mss.mpReady = true
		return nil
	}

	mss.mp = map[string]string{}
	err := json.Unmarshal(mss.js, &mss.mp)
	if err != nil {
		return err
	}
	mss.mpReady = true
	return nil
}
