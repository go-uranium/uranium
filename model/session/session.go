package session

import (
	"time"
)

type Session struct {
	Token string
	UID   int64
	UA    string
	IP    string
	// if is sudo mode
	Sudo      bool
	CreatedAt time.Time
	ExpireAt  time.Time
}

type Basic struct {
	Token    string
	UID      int64
	Sudo     bool
	ExpireAt time.Time
}

func (sess *Session) Valid() bool {
	return time.Now().Before(sess.ExpireAt)
}

func (bsc *Basic) Valid() bool {
	return time.Now().Before(bsc.ExpireAt)
}
