package session

import (
	"time"
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

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.Expire)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.Expire)
}
