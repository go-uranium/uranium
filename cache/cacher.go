package cache

import (
	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/post"
	"github.com/go-uranium/uranium/model/session"
	"github.com/go-uranium/uranium/model/user"
)

type Cacher interface {
	Init() error

	User(interface{}) (*user.User, error)
	UserDrop(interface{}) error
	UserDropAll() error

	IndexPostInfo(page int64) ([]*post.Info, error)
	IndexPostInfoRefresh() error

	Session(token string) (*session.Basic, error)
	SessionDropAll() error

	Categories() []*category.Category
	Category(interface{}) *category.Category
	CategoryRefresh() error

	DropAll() error
}
