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

type SimpleSession struct {
	UID    int
	Token  string
	Expire time.Time
}

func (s *Session) IsValid() bool {
	return time.Now().Before(s.Expire)
}

func (ss *SimpleSession) IsValid() bool {
	return time.Now().Before(ss.Expire)
}
