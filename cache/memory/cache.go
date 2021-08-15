package memory

import (
	"sync"

	"github.com/go-uranium/uranium/model/category"
	"github.com/go-uranium/uranium/model/post"
	"github.com/go-uranium/uranium/storage"
)

type Cache struct {
	data storage.Provider

	indexSize     int64
	indexRefresh  *sync.RWMutex
	indexPostInfo [][]*post.Info
	indexMaxSize  int64

	user    *sync.Map
	session *sync.Map

	cateRefresh     *sync.RWMutex
	categories      []*category.Category
	categoryByTID   map[int64]*category.Category
	categoryByTName map[string]*category.Category
}

func New(data storage.Provider, indexSize int64) *Cache {
	return &Cache{
		data: data,

		indexSize:     indexSize,
		indexRefresh:  &sync.RWMutex{},
		indexPostInfo: [][]*post.Info{},
		indexMaxSize:  100,

		user:    &sync.Map{},
		session: &sync.Map{},

		cateRefresh:     &sync.RWMutex{},
		categories:      []*category.Category{},
		categoryByTID:   map[int64]*category.Category{},
		categoryByTName: map[string]*category.Category{},
	}
}

func (cache *Cache) Init() error {
	err := cache.CategoryRefresh()
	if err != nil {
		return err
	}
	err = cache.IndexPostInfoRefresh()
	if err != nil {
		return err
	}
	return nil
}

func (cache *Cache) DropAll() error {
	err := cache.UserDropAll()
	if err != nil {
		return err
	}
	err = cache.IndexPostInfoRefresh()
	if err != nil {
		return err
	}
	err = cache.SessionDropAll()
	if err != nil {
		return err
	}
	err = cache.CategoryRefresh()
	if err != nil {
		return err
	}
	return nil
}
