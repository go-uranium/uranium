package session

import "time"

type Session struct {
	Token     string
	UID       int
	UA        string
	IP        string
	CreatedAt time.Time
	ExpireAt  time.Time
}

type Basic struct {
	Token  string
	UID    int
	Expire time.Time
}

func (s *Session) Valid() bool {
	return time.Now().Before(s.ExpireAt)
}

func (bs *Basic) Valid() bool {
	return time.Now().Before(bs.Expire)
}
