package cache

import (
	"github.com/go-ushio/ushio/core/user"
)

func (cache *Cache) UserByUID(uid int) (*user.SimpleUser, error) {
	v, ok := cache.userByUID[uid]
	if ok {
		return v, nil
	}

	u, err := cache.data.UserByUID(uid)
	if err != nil {
		return &user.SimpleUser{}, err
	}
	u.HashEmail()

	su := &user.SimpleUser{
		UID:         u.UID,
		Name:        u.Name,
		Username:    u.Username,
		HashedEmail: u.HashedEmail,
	}
	cache.userByUID[u.UID] = su
	return su, nil
}

func (cache *Cache) UserByUsername(username string) (*user.SimpleUser, error) {
	v, ok := cache.userByUsername[username]
	if ok {
		return v, nil
	}

	u, err := cache.data.UserByUsername(username)
	if err != nil {
		return &user.SimpleUser{}, err
	}

	su := &user.SimpleUser{
		UID:      u.UID,
		Name:     u.Name,
		Username: u.Username,
	}
	cache.userByUsername[u.Username] = su
	return su, nil
}

func (cache *Cache) UserDrop() error {

	return nil
}
