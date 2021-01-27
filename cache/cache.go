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
	userByUID      map[int]*user.SimpleUser
	userByUsername map[string]*user.SimpleUser
	sessionByToken map[string]*session.SimpleSession
}

func New(data *data.Data) *Cache {
	return &Cache{
		data:           data,
		indexPostInfo:  []*post.Info{},
		userByUID:      map[int]*user.SimpleUser{},
		userByUsername: map[string]*user.SimpleUser{},
		sessionByToken: map[string]*session.SimpleSession{},
	}
}
