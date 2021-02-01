package cache

import (
	"github.com/go-ushio/ushio/core/user"
)

func (cache *Cache) UserByUID(uid int) (*user.User, error) {
	cache.refresh.RLock()
	defer cache.refresh.RUnlock()
	v, ok := cache.userByUID[uid]
	if ok {
		return v, nil
	}

	u, err := cache.data.UserByUID(uid)
	if err != nil {
		return &user.User{}, err
	}
	cache.refresh.RUnlock()
	cache.refresh.Lock()
	cache.userByUID[u.UID] = u
	cache.refresh.Unlock()
	// cause defer at first
	cache.refresh.RLock()
	return u, nil
}

func (cache *Cache) UserByUsername(username string) (*user.User, error) {
	cache.refresh.RLock()
	defer cache.refresh.RUnlock()
	v, ok := cache.userByUsername[username]
	if ok {
		return v, nil
	}

	u, err := cache.data.UserByUsername(username)
	if err != nil {
		return &user.User{}, err
	}

	cache.refresh.RUnlock()
	cache.refresh.Lock()
	cache.userByUsername[u.Username] = u
	cache.refresh.Unlock()
	// cause defer at first
	cache.refresh.RLock()
	return u, nil
}

func (cache *Cache) UserDrop() error {
	cache.refresh.Lock()
	defer cache.refresh.Unlock()
	cache.userByUID = map[int]*user.User{}
	cache.userByUsername = map[string]*user.User{}
	return nil
}
