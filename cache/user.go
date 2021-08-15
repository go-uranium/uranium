package cache

import (
	"errors"
	"sync"

	"github.com/go-ushio/ushio/model/user"
)

func (cache *Cache) User(key interface{}) (*user.User, error) {
	value, ok := cache.user.Load(key)
	u, isUser := value.(*user.User)
	if !isUser {
		ok = false
	}
	if !ok {
		u := &user.User{}
		var err error
		switch key.(type) {
		case string:
			u, err = cache.data.UserByUsername(key.(string))
			if err != nil {
				return &user.User{}, err
			}
		case int64:
			u, err = cache.data.UserByUID(key.(int64))
			if err != nil {
				return &user.User{}, err
			}
		default:
			return nil, errors.New("invalid cache key for user")
		}
		cache.user.Store(key, u)
		return u, nil
	}
	return u, nil
}

func (cache *Cache) UserDrop(key interface{}) error {
	cache.user.Delete(key)
	return nil
}

func (cache *Cache) UserDropAll() error {
	cache.user = &sync.Map{}
	return nil
}
