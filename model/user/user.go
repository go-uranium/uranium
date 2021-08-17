package user

import (
	"time"

	"github.com/go-uranium/uranium/utils/clean"
	"github.com/go-uranium/uranium/utils/hash"
	"github.com/go-uranium/uranium/utils/sqlmap"
	"github.com/go-uranium/uranium/utils/sqlnull"
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
	// value: UsernameLowercase = Lowercase(Username)
	// regex: ^(?=.{1,20}$)(?!-)[a-z0-9-]{0,19}[a-z0-9]$
	// default: Lowercase(Username)
	// length: [1,20]
	// not null, unique
	UsernameLowercase string `json:"username_lowercase"`

	// note: Electrons is something like "karma" in Reddit or "coin" in V2EX
	// value: Electrons is an integer, which can be less than zero.
	// regex: /
	// default: 30
	// length: /
	// not null
	Electrons int32 `json:"electrons"`

	// note: /
	// value: IsMod is an boolean, which indicates whether the user is a mod.
	// regex: /
	// default: false
	// length: /
	// not null
	IsMod bool `json:"is_mod"`

	// note: ModPermission works only if IsMod == true
	// value: ModPermission is an unsigned integer, which indicates the permission of the mod.
	// regex: /
	// default: 0
	// length: /
	// not null
	ModPermission uint8 `json:"mod_permission"`

	// note: /
	// value: CreatedAt is a timestamp, which records the date when the user registered.
	// regex: /
	// default: time.Now()
	// length: /
	// not null
	CreatedAt time.Time `json:"created_at"`

	// note: When user deletes his/her account, all data except UID and Username would be removed,
	//		 field Deleted would be set to true, and user page would be redirected to "ghost".
	// value: Deleted is an boolean, which indicates whether the user is deleted.
	// regex: /
	// default: false
	// length: /
	// not null
	Deleted bool `json:"deleted"`
}

// Basic is the basic information of a user (simplified version of User)
type Basic struct {
	UID               int32  `json:"uid"`
	Username          string `json:"username"`
	UsernameLowercase string `json:"username_lowercase"`
	IsMod             bool   `json:"is_mod"`
	// In case of some situations.
	Deleted bool `json:"deleted"`
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
	// value: SecurityEmail is a string type, which can bu null.
	// regex: ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$
	// default: /
	// length: [4,320]
	// -
	SecurityEmail sqlnull.String `json:"security_email"`

	//TODO: finish the fields below

	// if user is locked
	// not null
	Locked bool `json:"locked"`

	// lock user until
	// must >= current time
	// otherwise, Locked would be set to false
	LockedTill sqlnull.Time `json:"locked_till"`

	// if user is disabled
	// not null
	Disabled bool `json:"disabled"`
}

type Profile struct {
	// reference: user.User.UID
	UID int32 `json:"uid"`

	// "name" here stands for "display name"
	// not null
	// length: [1,30]
	Name string `json:"name"`

	// bio
	// can be null
	// length: (0,255]
	Bio sqlnull.String `json:"bio"`

	// geolocation
	// can be null
	// length: [1,15]
	Location sqlnull.String `json:"location"`

	// birthday
	// can be null
	Birthday sqlnull.Time `json:"birthday"`

	// email to be displayed
	// can be null
	Email string `json:"email"`

	// social is a map[string]string type
	// example: {"github":"@iochen","website":"https://iochen.com/"}
	// not null
	// default: "{}"
	Social sqlmap.StringString `json:"social"`
}

// Valid checks whether user auth info(password) is valid
func (auth *Auth) Valid(password []byte) bool {
	return hash.SHA256Validate(auth.Password, password)
}

// Tidy tidies user info and generates default avatar
func (u *User) Tidy() {
	u.Username = clean.Username(u.Username)
}

func (auth *Auth) Masking() {
	auth.Password = nil
}
