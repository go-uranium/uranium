package cache

import (
	"github.com/go-ushio/ushio/core/session"
)

func (cache *Cache) SessionByToken(token string) (*session.Basic, error) {
	v, ok := cache.sessionByToken[token]
	if ok {
		return v, nil
	}

	s, err := cache.data.SessionByToken(token)
	if err != nil {
		return &session.Basic{}, err
	}
	ss := &session.Basic{
		Token:  s.Token,
		UID:    s.UID,
		Expire: s.Expire,
	}
	cache.sessionByToken[s.Token] = ss
	return ss, nil
}

func (cache *Cache) SessionDrop() error {
	cache.sessionByToken = map[string]*session.Basic{}
	return nil
}
