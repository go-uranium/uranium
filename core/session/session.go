package session

import "time"

type Session struct {
	Token  string
	UID    int
	UA     string
	IP     string
	Time   time.Time
	Expire time.Time
}

type Basic struct {
	Token  string
	UID    int
	Expire time.Time
}

func (s *Session) IsValid() bool {
	return time.Now().Before(s.Expire)
}

func (bs *Basic) IsValid() bool {
	return time.Now().Before(bs.Expire)
}
