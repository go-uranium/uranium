package cache

import (
	"github.com/go-ushio/ushio/model/category"
	"github.com/go-ushio/ushio/model/post"
	"github.com/go-ushio/ushio/model/session"
	"github.com/go-ushio/ushio/model/user"
)

type Cacher interface {
	Init() error

	User(interface{}) (*user.User, error)
	UserDrop(interface{}) error
	UserDropAll() error

	IndexPostInfo(page int64) ([]*post.Info,error)
	IndexPostInfoRefresh() error

	IndexSize() int64

	Session(token string) (*session.Basic, error)
	SessionDropAll() error

	Categories() []*category.Category
	Category(interface{}) *category.Category
	CategoryRefresh() error

	DropAll() error
}
