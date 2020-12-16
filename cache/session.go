package cache

import (
	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/session"
)

var sessionByToken = map[string]*session.SimpleSession{}

func SessionByToken(token string) (*session.SimpleSession, error) {
	v, ok := sessionByToken[token]
	if ok {
		return v, nil
	}

	s, err := data.SessionByToken(token)
	if err != nil {
		return &session.SimpleSession{}, err
	}
	ss := &session.SimpleSession{
		UID:    s.UID,
		Token:  s.Token,
		Expire: s.Expire,
	}
	sessionByToken[s.Token] = ss
	return ss, nil
}
