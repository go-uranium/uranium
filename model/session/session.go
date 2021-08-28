package session

import (
	"time"

	"github.com/go-uranium/uranium/model/user"
)

const (
	UNKNOWN int16 = iota
	USER
	SUDO
	MODERATOR
	ADMIN
)

const (
	UNKNOWN_CHAR   = '0' //30
	USER_CHAR      = '1' //31
	USER_SUDO_CHAR = '2' //32
	MODERATOR_CHAR = '3' //33
	ADMIN_CHAR     = '4' //34
)

type Session struct {
	// length: 32
	Token   string
	UID     int32
	Mode    int16
	UA      string
	IP      string
	Created time.Time
	Expire  time.Time
}

type Basic struct {
	Token  string
	UID    int32
	Mode   int16
	Expire time.Time
}

type Cache struct {
	UID  int32
	Mode int16
	// if not expired
	Valid bool
}

func CleanTypeWithDefault(t int16) int16 {
	// 1, 2, 3, 4
	if t >= 1 && t <= 4 {
		return t
	}
	// default
	return 1
}

func ValidSessionType(userType, sessType int16) bool {
	switch userType {
	case user.USER:
		return sessType == USER || sessType == SUDO
	case user.MODERATOR:
		return sessType == USER || sessType == SUDO || sessType == MODERATOR
	case user.ADMIN:
		return sessType == USER || sessType == SUDO || sessType == MODERATOR || sessType == ADMIN
	default:
		return false
	}
}

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.Expire)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.Expire)
}
