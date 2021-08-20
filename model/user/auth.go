package user

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/utils/hash"
	"github.com/go-uranium/uranium/utils/sqlnull"
)

type Auth struct {
	// reference: user.User.UID
	UID int32 `json:"uid"`

	// note: Auth.Email is the email for login action or user verification
	// value: Email is a string, must be lowercase.
	// regex: ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// default: /
	// length: [4,320]
	// not null, unique
	Email string `json:"email"`

	// note: /
	// value: Password is a []byte type, which is the SHA256 hash of user password.
	// regex: /
	// default: /
	// length: 32
	// not null
	Password []byte `json:"_"`

	// note: SecurityEmail is an alternative address for user verification,
	//       and it receives a copy of security alert.
	// value: SecurityEmail is a string, which can be null.
	// regex: ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// default: /
	// length: [4,320]
	// /
	SecurityEmail sqlnull.String `json:"security_email"`

	// !!! Will not be worked on currently.
	// note: When TwoFactor == true, user must pass 2FA challenge when login.
	// value: TwoFactor is a boolean.
	// regex: /
	// default: false
	// length: /
	// not null
	TwoFactor bool `json:"two_factor"`

	// note: When Locked == true, user cannot login, or perform any actions.
	// value: Locked is a boolean, which means that the user has been locked.
	// regex: /
	// default: false
	// length: /
	// not null
	Locked bool `json:"locked"`

	// note: LockedTill only works if Locked == true
	// value: LockedTill is a timestamp, after which Locked should be set to false.
	// regex: /
	// default: null
	// length: /
	// /
	LockedTill sqlnull.Time `json:"locked_till"`

	// note: When Disabled == true, user cannot login, or perform any actions.
	// value: Disabled is a boolean, which means that the user has been disabled.
	// regex: /
	// default: false
	// length: /
	// not null
	Disabled bool `json:"disabled"`
}

func (auth *Auth) Valid(password []byte) error {
	if auth.LockedOrDisabled() {
		return fiber.NewError(http.StatusUnauthorized, "User has been locked or disabled.")
	}
	if !auth.PasswordValid(password) {
		return fiber.NewError(http.StatusUnauthorized, "User password is wrong.")
	}
	return nil
}

// Valid checks whether user password info is valid
func (auth *Auth) PasswordValid(password []byte) bool {
	return hash.SHA256Validate(auth.Password, password)
}

func (auth *Auth) LockedOrDisabled() bool {
	if auth.Disabled {
		return true
	}
	if auth.Locked && auth.LockedTill.Valid && auth.LockedTill.Time.After(time.Now()) {
		return true
	}
	return false
}

func (auth *Auth) Masking() {
	auth.Password = nil
}
