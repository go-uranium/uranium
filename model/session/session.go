package session

import (
	"time"
)

const (
	SESSION_MODE_UNKNOWN int16 = iota
	SESSION_MODE_USER
	SESSION_MODE_MODERATOR
	SESSION_MODE_ADMIN
)

const (
	SESSION_MODE_UNKNOWN_STR   = "0"
	SESSION_MODE_USER_STR      = "1"
	SESSION_MODE_MODERATOR_STR = "2"
	SESSION_MODE_ADMIN_STR     = "3"
)

type Session struct {
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

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.Expire)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.Expire)
}
