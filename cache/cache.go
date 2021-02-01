package cache

import (
	"sync"

	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/user"
	"github.com/go-ushio/ushio/data"
)

type Cache struct {
	data *data.Data

	refresh        *sync.RWMutex
	indexPostInfo  []*post.Info
	postsNotEnough bool
	userByUID      map[int]*user.User
	userByUsername map[string]*user.User
	sessionByToken map[string]*session.Basic
}

func New(data *data.Data) *Cache {
	return &Cache{
		data: data,

		refresh:        &sync.RWMutex{},
		indexPostInfo:  []*post.Info{},
		postsNotEnough: false,
		userByUID:      map[int]*user.User{},
		userByUsername: map[string]*user.User{},
		sessionByToken: map[string]*session.Basic{},
	}
}

func (cache *Cache) DropAll() error {
	cache.refresh.Lock()
	defer cache.refresh.Unlock()
	err := cache.UserDrop()
	if err != nil {
		return err
	}
	err = cache.IndexPostInfoDrop()
	if err != nil {
		return err
	}
	err = cache.SessionDrop()
	if err != nil {
		return err
	}
	return nil
}
