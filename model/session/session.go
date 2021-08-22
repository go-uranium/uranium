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
	SESSION_MODE_UNKNOWN_CHAR   = '0'
	SESSION_MODE_USER_CHAR      = '1'
	SESSION_MODE_MODERATOR_CHAR = '2'
	SESSION_MODE_ADMIN_CHAR     = '3'
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

type Cache struct {
	Token string
	UID   int32
	Mode  int16
}

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.Expire)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.Expire)
}
