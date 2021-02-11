package cache

import (
	"github.com/go-ushio/ushio/core/category"
)

func (cache *Cache) Categories() ([]*category.Category, error) {
	cache.cateRefresh.RLock()
	defer cache.cateRefresh.RUnlock()
	return cache.categories, nil
}

func (cache *Cache) Category(key interface{}) *category.Category {
	cache.cateRefresh.RLock()
	defer cache.cateRefresh.RUnlock()
	switch key.(type) {
	case int:
		return cache.categoryByTID[key.(int)]
	case string:
		return cache.categoryByTName[key.(string)]
	default:
		return nil
	}
}

func (cache *Cache) CategoryRefresh() error {
	cache.cateRefresh.Lock()
	defer cache.cateRefresh.Unlock()
	categories, err := cache.data.GetCategories()
	if err != nil {
		return err
	}
	cache.categories = categories
	cache.categoryByTID = map[int]*category.Category{}
	cache.categoryByTName = map[string]*category.Category{}
	for i := range categories {
		cache.categoryByTID[categories[i].TID] = categories[i]
		cache.categoryByTName[categories[i].TName] = categories[i]
	}
	return nil
}
