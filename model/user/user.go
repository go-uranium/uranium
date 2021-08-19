package user

import (
	"encoding/json"
	"time"

	"github.com/go-uranium/uranium/utils/hash"
	"github.com/go-uranium/uranium/utils/sqlmap"
	"github.com/go-uranium/uranium/utils/sqlnull"
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
	// regex: ^(?=.{1,20}$)(?!-)[a-zA-Z0-9-]{0,19}[a-zA-Z0-9]$
	// default: /
	// length: [1,20]
	// not null, unique
	Username string `json:"username"`

	// note: /
	// value: Lowercase = Lowercase(Username)
	// regex: ^(?=.{1,20}$)(?!-)[a-z0-9-]{0,19}[a-z0-9]$
	// default: Lowercase(Username)
	// length: [1,20]
	// not null, unique
	Lowercase string `json:"lowercase"`

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

type basic struct {
	UID      int32  `json:"uid"`
	Username string `json:"username"`
	Admin    int16  `json:"admin"`
}

// Basic is the basic information of a user (simplified version of User),
// ! Basic is read-only !
type Basic struct {
	basic      basic
	basicReady bool
	js         []byte
	jsReady    bool
}

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

	// Will not be worked on currently.
	//// note: When TwoFactor == true, user must pass 2FA challenge when login.
	//// value: TwoFactor is a boolean.
	//// regex: /
	//// default: false
	//// length: /
	//// not null
	//TwoFactor bool `json:"two_factor"`

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

type Profile struct {
	// reference: user.User.UID
	UID int32 `json:"uid"`

	// note: Name is the name displayed on user profile page.
	// value: Name is a string, which should be UTF-8 chars.
	// regex: /
	// default: /
	// length: [1,30]
	// not null
	Name string `json:"name"`

	// note: Bio is the bio displayed on user profile page.
	// value: Bio is a string, which should be UTF-8 chars.
	// regex: /
	// default: null
	// length: [1,255]
	// /
	Bio sqlnull.String `json:"bio"`

	// note: Location(geolocation) is the location displayed on user profile page.
	// value: Location is a string, which should be UTF-8 chars.
	// regex: /
	// default: null
	// length: [1,15]
	// /
	Location sqlnull.String `json:"location"`

	// note: Birthday is the birthday displayed on user profile page.
	// value: Birthday is a string, which should be UTF-8 chars.
	// regex: /
	// default: null
	// length: [1,15]
	// /
	Birthday sqlnull.Time `json:"birthday"`

	// note: Email is the email address displayed on profile page.
	// value: Email is a string, which can be null.
	// regex: ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// default: null
	// length: [4,320]
	// /
	Email sqlnull.String `json:"email"`

	// note: example: {"github":"@iochen","website":"https://iochen.com/"}
	// value: Social is a map[string]string.
	// regex: /
	// default: "{}"
	// length: [4,320]
	// not null
	Social sqlmap.StringString `json:"social"`
}

func NewBasicFromJSON(data []byte) *Basic {
	return &Basic{
		basic:   basic{},
		js:      data,
		jsReady: true,
	}
}

func (u *User) Basic() *Basic {
	return &Basic{
		basic: basic{
			UID:      u.UID,
			Username: u.Username,
			Admin:    u.Admin,
		},
	}
}

func (b *Basic) MarshalJSON() ([]byte, error) {
	if b.jsReady {
		return b.js, nil
	}
	data, err := json.Marshal(b.basic)
	if err != nil {
		return nil, err
	}
	b.js = data
	b.jsReady = true
	return data, nil
}

// Valid checks whether user auth info(password) is valid
func (auth *Auth) Valid(password []byte) bool {
	return hash.SHA256Validate(auth.Password, password)
}

func (auth *Auth) Masking() {
	auth.Password = nil
}
