package session

import (
	"time"
)

type Session struct {
	Token   string
	UID     int64
	UA      string
	IP      string
	Mode    bool
	Created time.Time
	Expire  time.Time
}

type Basic struct {
	Token  string
	UID    int64
	Sudo   bool
	Expire time.Time
}

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.Expire)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.Expire)
}
