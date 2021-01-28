package cache

import (
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/user"
)

type Cacher interface {
	UserByUID(uid int) (*user.User, error)
	UserByUsername(username string) (*user.User, error)
	UserDrop() error

	IndexPostInfo(size int) ([]*post.Info, error)
	IndexPostInfoDrop() error

	SessionByToken(token string) (*session.Basic, error)
	SessionDrop() error

	DropAll() error
}
