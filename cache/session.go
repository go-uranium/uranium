package cache

import (
	"sync"

	"github.com/go-uranium/uranium/model/session"
)

func (cache *Cache) Session(token string) (*session.Basic, error) {
	value, ok := cache.session.Load(token)
	bsc, isBS := value.(*session.Basic)
	if !isBS {
		ok = false
	}
	if !ok {
		bsc, err := cache.data.SessionBasicByToken(token)
		if err != nil {
			return &session.Basic{}, err
		}
		cache.session.Store(token, bsc)
		return bsc, nil
	}
	return bsc, nil
}

func (cache *Cache) SessionDropAll() error {
	cache.session = &sync.Map{}
	return nil
}
