package cache

import (
	"github.com/go-ushio/ushio/model/post"
)

func (cache *Cache) IndexPostInfo(page int64) ([]*post.Info,error) {
	cache.indexRefresh.RLock()
	defer cache.indexRefresh.RUnlock()
	if page < 1 {
		return nil,nil
	}

	if page > cache.indexMaxSize {
		return cache.data.PostsInfoByActivity(false,
			cache.indexSize, (page-1)*cache.indexSize)
	}

	if cache.indexPostInfo[page-1] == nil {
		err := cache.indexPostInfoRefresh(page)
		if err != nil {
			return nil,err
		}
	}

	return cache.indexPostInfo[page-1],nil
}

func (cache *Cache) IndexPostInfoRefresh() error {
	cache.indexRefresh.Lock()
	defer cache.indexRefresh.Unlock()
	//infos, err := cache.data.PostsInfoByActivity(false, cache.indexSize, 0)
	//if err != nil {
	//	return err
	//}
	//
	cache.indexPostInfo = make([][]*post.Info, cache.indexMaxSize)

	return cache.indexPostInfoRefresh(1)
}

func (cache *Cache) indexPostInfoRefresh(page int64) error {
	if page > cache.indexMaxSize {
		return nil
	}
	infos, err := cache.data.PostsInfoByActivity(false, cache.indexSize,
		(page-1)*cache.indexSize)
	if err != nil {
		return err
	}
	cache.indexPostInfo[page-1] = infos
	return nil
}
