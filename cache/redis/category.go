package rcache

import (
	"database/sql"
	"errors"
	"strings"
	"sync"

	"github.com/go-uranium/uranium/model/category"
)

var (
	CacheCategoryInMem = true
)

var (
	categoriesInMem     map[string]*category.Category
	categoriesInMemLock = &sync.RWMutex{}
)

func (r *RCache) CategoryByTName(tname string) (*category.Category, bool, error) {
	if r.cacheCategoryInMem {
		categoriesInMemLock.RLock()
		defer categoriesInMemLock.RUnlock()
		cate, ok := categoriesInMem[tname]
		if !ok {
			return nil, false, sql.ErrNoRows
		}
		return cate, true, nil
	}
	c, err := r.rdb.Get(ctx, "category:"+tname).Result()
	if err != nil {
		return &category.Category{}, false, err
	}
	parts := strings.Split(c, ",")
	if len(parts) != 2 {
		return nil, false,
			errors.New("unexpected length when splitting marshaled category in redis")
	}
	return &category.Category{
		TName: tname,
		Name:  parts[0],
		Color: parts[1],
	}, true, nil
}

func (r *RCache) RefreshCategory() error {
	var err error
	if r.cacheCategoryInMem {
		err = r.refreshCategoryInMem()
	} else {
		err = r.refreshCategoryInRedis()
	}
	return err
}

func (r *RCache) refreshCategoryInMem() error {
	categories, err := r.storage.Categories()
	if err != nil {
		return err
	}
	newCategoriesInMem := make(map[string]*category.Category)
	for i := range categories {
		newCategoriesInMem[categories[i].TName] = categories[i]
	}
	categoriesInMemLock.Lock()
	categoriesInMem = newCategoriesInMem
	categoriesInMemLock.Unlock()
	return nil
}

func (r *RCache) refreshCategoryInRedis() error {
	categories, err := r.storage.Categories()
	if err != nil {
		return err
	}
	for i := range categories {
		_, err := r.rdb.Set(ctx, "category:"+categories[i].TName,
			strings.Join([]string{categories[i].Name, categories[i].Color}, ","),
			r.ttl.Category).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
