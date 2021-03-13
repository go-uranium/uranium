package user

import (
	"encoding/hex"
	"time"

	"github.com/go-ushio/ushio/common/scan"
	"github.com/go-ushio/ushio/utils/clean"
	"github.com/go-ushio/ushio/utils/hash"
)

type User struct {
	UID       int64     `json:"uid"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	Artifact  int64     `json:"artifact"`
}

type Simple struct {
	UID      int64  `json:"uid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Auth struct {
	UID           int64
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
		&u.Bio, &u.CreatedAt, &u.Artifact)
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
