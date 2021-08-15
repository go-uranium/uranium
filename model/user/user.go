package user

import (
	"time"

	"github.com/go-uranium/uranium/utils/clean"
	"github.com/go-uranium/uranium/utils/hash"
)

type User struct {
	UID int64 `json:"uid"`
	// "name" here stands for "display name"
	Name     string `json:"name"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	// "electrons" is something like "karma" in reddit
	Electrons int64 `json:"electrons"`
	// when set to 0, means the root mod
	Mod       uint8     `json:"mod"`
	CreatedAt time.Time `json:"created_at"`
}

type Basic struct {
	UID      int64  `json:"uid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Mod      uint8  `json:"mod"`
}

type Auth struct {
	UID           int64
	Email         string
	Password      []byte
	SecurityEmail string
	Locked        time.Time
	Disabled      bool
}

type Profile struct {
	Email  string            `json:"email"`
	Bio    string            `json:"bio"`
	Social map[string]string `json:"social"`
}

// Valid checks whether user auth info(password) is valid
func (auth *Auth) Valid(password []byte) bool {
	return hash.SHA256Compare(auth.Password, password)
}

// Tidy tidies user info and generates default avatar
func (u *User) Tidy() {
	u.Username = clean.String(u.Username)
}
