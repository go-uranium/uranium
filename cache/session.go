package cache

import (
	"github.com/go-ushio/ushio/core/session"
)

func (cache *Cache) SessionByToken(token string) (*session.SimpleSession, error) {
	v, ok := cache.sessionByToken[token]
	if ok {
		return v, nil
	}

	s, err := cache.data.SessionByToken(token)
	if err != nil {
		return &session.SimpleSession{}, err
	}
	ss := &session.SimpleSession{
		UID:    s.UID,
		Token:  s.Token,
		Expire: s.Expire,
	}
	cache.sessionByToken[s.Token] = ss
	return ss, nil
}

func (cache *Cache) SessionDrop() error {

	return nil
}
