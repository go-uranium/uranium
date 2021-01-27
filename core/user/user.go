package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-ushio/ushio/utils/flags"
	"github.com/go-ushio/ushio/utils/hash"
)

type User struct {
	UID      int    `json:"uid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"-"`
	// hashed password
	Password    []byte       `json:"-"`
	HashedEmail string       `json:"hashed_email"`
	CreatedAt   time.Time    `json:"created_at"`
	IsAdmin     bool         `json:"is_admin"`
	Banned      bool         `json:"-"`
	Locked      bool         `json:"-"`
	Flags       *flags.Flags `json:"-"`
}

type SimpleUser struct {
	UID         int    `json:"uid"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	HashedEmail string `json:"hashed_email"`
}

func (u *User) Json() ([]byte, error) {
	j, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (u *User) Valid(password string) bool {
	return hash.Compare(u.Password, password)
}

func (u *User) Tidy() {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)
}

func (u *User) HashEmail() {
	h := md5.Sum([]byte(u.Email))
	u.HashedEmail = hex.EncodeToString(h[:])
}
