package cache

import (
	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/user"
)

type Cacher interface {
	UserBasicByUID(uid int32) (*user.Basic, error)
	UserUIDByLowercase(lowercase string) (int32, error)
	RefreshUserBasicByUID(uid int32) (*user.Basic, error)
	RefreshUserUIDByLowercase(lowercase string) (int32, error)

	CategoryByTName(tname string) (*category.Category, error)
	RefreshCategory() error
}
