package cache

import (
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/user"
	"github.com/go-ushio/ushio/data"
)

type Cache struct {
	data *data.Data

	indexPostInfo  []*post.Info
	postsNotEnough bool
	userByUID      map[int]*user.User
	userByUsername map[string]*user.User
	sessionByToken map[string]*session.Basic
}

func New(data *data.Data) *Cache {
	return &Cache{
		data: data,

		indexPostInfo:  []*post.Info{},
		postsNotEnough: false,
		userByUID:      map[int]*user.User{},
		userByUsername: map[string]*user.User{},
		sessionByToken: map[string]*session.Basic{},
	}
}

func (cache *Cache) DropAll() error {
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
