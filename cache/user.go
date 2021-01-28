package cache

import (
	"github.com/go-ushio/ushio/core/user"
)

func (cache *Cache) UserByUID(uid int) (*user.User, error) {
	v, ok := cache.userByUID[uid]
	if ok {
		return v, nil
	}

	u, err := cache.data.UserByUID(uid)
	if err != nil {
		return &user.User{}, err
	}

	cache.userByUID[u.UID] = u
	return u, nil
}

func (cache *Cache) UserByUsername(username string) (*user.User, error) {
	v, ok := cache.userByUsername[username]
	if ok {
		return v, nil
	}

	u, err := cache.data.UserByUsername(username)
	if err != nil {
		return &user.User{}, err
	}

	cache.userByUsername[u.Username] = u
	return u, nil
}

func (cache *Cache) UserDrop() error {
	cache.userByUID = map[int]*user.User{}
	cache.userByUsername = map[string]*user.User{}
	return nil
}
