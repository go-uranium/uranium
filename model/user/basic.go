package user

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type BasicCore struct {
	UID      int32  `json:"uid"`
	Username string `json:"username"`
	Admin    int16  `json:"admin"`
}

// Basic is the basic information of a user (simplified version of User),
// ! Basic is read-only !
type Basic struct {
	UID        int32 `json:"_"`
	basic      *BasicCore
	basicReady bool
	js         []byte
	jsReady    bool
}

type BasicList struct {
	bl      []*Basic
	current int
	length  int
}

func NewBasicFromCore(c *BasicCore) *Basic {
	return &Basic{
		basic:      c,
		basicReady: true,
	}
}

func NewEmptyBasicList(list []int32) *BasicList {
	bl := &BasicList{
		bl:      make([]*Basic, len(list)),
		current: -1,
		length:  len(list),
	}
	for k, v := range list {
		bl.bl[k] = &Basic{
			UID: v,
		}
	}
	return bl
}

func (bl *BasicList) Next() bool {
	bl.current++
	if bl.current >= bl.length {
		return false
	}
	return true
}

func (bl *BasicList) CurrentUID() int32 {
	return bl.bl[bl.current].UID
}

func (bl *BasicList) FillCurrentFromJSON(data []byte) {
	bl.bl[bl.current].FillFromJSON(data)
}

func NewBasicFromJSON(data []byte) *Basic {
	return &Basic{
		js:      data,
		jsReady: true,
	}
}

func (b *Basic) FillFromJSON(data []byte) {
	b.js = data
	b.jsReady = true
}

func (b *Basic) MarshalJSON() ([]byte, error) {
	if b.jsReady {
		return b.js, nil
	}
	if !b.basicReady {
		return nil, errors.New("no available data source found for user.Basic to marshal")
	}
	data, err := json.Marshal(b.basic)
	if err != nil {
		return nil, err
	}
	b.js = data
	b.jsReady = true
	return data, nil
}

func (b *Basic) Core() *BasicCore {
	if !b.basicReady {
		if !b.jsReady {
			return &BasicCore{}
		}
		if err := b.parseJs(); err != nil {
			return &BasicCore{}
		}
	}
	return b.basic
}

func (b *Basic) parseJs() error {
	b.basic = &BasicCore{}
	err := json.Unmarshal(b.js, b.basic)
	if err != nil {
		return err
	}
	return nil
}

func (b *Basic) Equal(b1 *Basic) bool {
	if !b.basicReady {
		if !b.jsReady {
			return false
		}
		if err := b.parseJs(); err != nil {
			return false
		}
	}

	if !b1.basicReady {
		if !b1.jsReady {
			return false
		}
		if err := b1.parseJs(); err != nil {
			return false
		}
	}

	return b.basic.UID == b1.basic.UID &&
		b.basic.Admin == b1.basic.Admin &&
		b.basic.Username == b1.basic.Username
}
