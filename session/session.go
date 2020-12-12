package session

import "time"

type Session struct {
	UID    int
	Token  string
	UA     string
	IP     string
	Time   time.Time
	Expire time.Time
}

func (s *Session) Valid() bool {
	return time.Now().Before(s.Expire)
}
