package user

import (
	"time"
)

const (
	ADMIN_NOT_ADMIN int16 = 0
	ADMIN_MODERATOR int16 = 1
	ADMIN_WEBMASTER int16 = 2
)

type User struct {
	// note: type int32 is sufficient, and UID must be unique
	// value: UID is an auto-incremented integer, must be greater than zero.
	// regex: /
	// default: /
	// length: /
	// not null, auto increment, unique
	UID int32 `json:"uid"`

	// note: Lowercase(Username) must be unique.
	//       When user is deleted, username would not be released automatically.
	// value: Username is a string, which only contains alphanumeric characters or single hyphens,
	//        and cannot begin or end with a hyphen.
	// regex: ^[a-zA-Z0-9][a-zA-Z0-9-]{0,18}[a-zA-Z0-9]$
	// default: /
	// length: [2,20]
	// not null, unique
	Username string `json:"username"`

	// note: Electrons is something like "karma" in Reddit or "coin" in V2EX
	// value: Electrons is an integer, which can be less than zero.
	// regex: /
	// default: 30
	// length: /
	// not null
	Electrons int32 `json:"electrons"`

	// note: /
	// value: Admin is an integer, which indicates whether the user is an admin and which role he/she is.
	// regex: /
	// default: 0
	// length: /
	// not null
	Admin int16 `json:"admin"`

	// note: /
	// value: Created is a timestamp, which records the date when the user registered.
	// regex: /
	// default: time.Now()
	// length: /
	// not null
	Created time.Time `json:"created"`

	// note: When user deletes his/her account, all data except UID and Username would be removed,
	//		 field Deleted would be set to true, and user page would be redirected to "ghost".
	// value: Deleted is an boolean, which indicates whether the user is deleted.
	// regex: /
	// default: false
	// length: /
	// not null
	Deleted bool `json:"deleted"`
}

func (u *User) Basic() *Basic {
	return &Basic{
		basic: &BasicCore{
			UID:      u.UID,
			Username: u.Username,
			Admin:    u.Admin,
		},
		basicReady: true,
	}
}
