package user

import (
	"github.com/go-uranium/uranium/utils/sqlmap"
	"github.com/go-uranium/uranium/utils/sqlnull"
)

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
	// length: [1,30]
	// /
	Location sqlnull.String `json:"location"`

	// note: Birthday is the birthday displayed on user profile page.
	// value: Birthday is a Date, which is stored as int32.
	// regex: /
	// default: null
	// length: /
	// /
	Birthday sqlnull.Date `json:"birthday"`

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
