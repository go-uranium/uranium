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

func (s *Session) IsValid() bool {
	return time.Now().After(s.Expire)
}
