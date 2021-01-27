package cache

import (
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/user"
)

type Cacher interface {
	UserByUID(uid int) (*user.SimpleUser, error)
	UserByUsername(username string) (*user.SimpleUser, error)
	UserDrop() error

	IndexPostInfo(size int) ([]*post.Info, error)
	IndexPostInfoDrop() error

	SessionByToken(token string) (*session.SimpleSession, error)
	SessionDrop() error
}
