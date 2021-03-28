package cache

import (
	"github.com/go-ushio/ushio/core/category"
	"github.com/go-ushio/ushio/core/post"
	"github.com/go-ushio/ushio/core/session"
	"github.com/go-ushio/ushio/core/user"
)

type Cacher interface {
	Init() error

	User(interface{}) (*user.User, error)
	UserDrop(interface{}) error
	UserDropAll() error

	IndexPostInfo() []*post.Info
	IndexPostInfoRefresh() error

	IndexSize() int64

	Session(token string) (*session.Basic, error)
	SessionDropAll() error

	Categories() []*category.Category
	Category(interface{}) *category.Category
	CategoryRefresh() error

	DropAll() error
}
