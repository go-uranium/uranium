package cache

import (
	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/session"
	"github.com/go-uranium/uranium/model/user"
)

type Cacher interface {
	UserBasicByUID(uid int32) (*user.Basic, error)
	UserUIDByUsername(username string) (int32, error)
	RefreshUserBasicByUID(uid int32) (*user.Basic, error)
	RefreshUserUIDByUsername(username string) (int32, error)

	ValidSessionByToken(token string) (*session.Cache, error)
	RefreshValidSessionByToken(token string) (*session.Cache, error)

	CategoryByTName(tname string) (*category.Category, error)
	RefreshCategory() error
}
