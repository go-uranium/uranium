package user

import (
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/go-ushio/ushio/common/put"
	"github.com/go-ushio/ushio/common/scan"
	"github.com/go-ushio/ushio/utils/clean"
	"github.com/go-ushio/ushio/utils/flags"
	"github.com/go-ushio/ushio/utils/hash"
)

type User struct {
	UID       int         `json:"uid"`
	Name      string      `json:"name"`
	Username  string      `json:"username"`
	Email     string      `json:"-"`
	Avatar    string      `json:"avatar"`
	CreatedAt time.Time   `json:"created_at"`
	IsAdmin   bool        `json:"is_admin"`
	Banned    bool        `json:"banned"`
	Flags     flags.Flags `json:"-"`
}

type Auth struct {
	UID           int
	Password      []byte
	Locked        bool
	SecurityEmail string
}

// Valid checks whether user auth info(password) is valid
func (auth *Auth) Valid(password []byte) bool {
	return hash.SHA256Compare(auth.Password, password)
}

// Tidy tidies user info and generates default avatar
func (u *User) Tidy() {
	u.Username = clean.String(u.Username)
	u.Email = clean.String(u.Email)
	if len(u.Avatar) == 0 && len(u.Email) != 0 {
		u.Avatar = hex.EncodeToString(hash.MD5([]byte(u.Email)))
	}
}

// ScanUser scans full User info
// for actions like: scan user info
func ScanUser(scanner scan.Scanner) (*User, error) {
	u := &User{}
	err := scanner.Scan(&u.UID, &u.Name, &u.Username, &u.Email, &u.Avatar,
		&u.CreatedAt, &u.IsAdmin, &u.Banned, &u.Flags)
	if err != nil {
		return &User{}, err
	}
	u.Tidy()
	return u, nil
}

// ScanAuth scans full Auth info
// for actions like: scan user info
func ScanAuth(scanner scan.Scanner) (*Auth, error) {
	auth := &Auth{}
	err := scanner.Scan(&auth.UID, &auth.Password, &auth.Locked, &auth.SecurityEmail)
	if err != nil {
		return &Auth{}, err
	}
	return auth, nil
}

// Put puts User info without user.UID
// for actions like: insert new user info
func (u *User) Put(putter put.Putter) (sql.Result, error) {
	return putter.Put(u.Name, u.Username, u.Email, u.Avatar,
		u.CreatedAt, u.IsAdmin, u.Banned, u.Flags)
}

// PutWithUID puts User info with user.UID at last
// for actions like: update user info
func (u *User) PutWithUID(putter put.Putter) (sql.Result, error) {
	return putter.Put(u.Name, u.Username, u.Email, u.Avatar,
		u.CreatedAt, u.IsAdmin, u.Banned, u.Flags, u.UID)
}

// PutWithUIDFirst puts Auth info with auth.UID at first
// for actions like: insert user auth
func (auth *Auth) PutWithUIDFirst(putter put.Putter) (sql.Result, error) {
	return putter.Put(auth.UID, auth.Password, auth.Locked, auth.SecurityEmail)
}

// PutWithUIDLast puts Auth info with auth.UID at last
// for actions like: update user auth
func (auth *Auth) PutWithUIDLast(putter put.Putter) (sql.Result, error) {
	return putter.Put(auth.Password, auth.Locked, auth.SecurityEmail, auth.UID)
}
