package cache

import (
	"github.com/go-ushio/ushio/core/session"
)

func (cache *Cache) SessionByToken(token string) (*session.Basic, error) {
	cache.refresh.RLock()
	defer cache.refresh.RUnlock()
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
	cache.refresh.RUnlock()
	cache.refresh.Lock()
	cache.sessionByToken[s.Token] = ss
	cache.refresh.Unlock()
	// cause defer at first
	cache.refresh.RLock()
	return ss, nil
}

func (cache *Cache) SessionDrop() error {
	cache.refresh.Lock()
	defer cache.refresh.Unlock()
	cache.sessionByToken = map[string]*session.Basic{}
	return nil
}
