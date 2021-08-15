package cache

import (
	"github.com/go-ushio/ushio/model/post"
)

func (cache *Cache) IndexPostInfo(page int64) []*post.Info {
	cache.indexRefresh.RLock()
	defer cache.indexRefresh.RUnlock()
	if page < 1 {
		return nil
	}

	if page > cache.indexMaxSize {
		infos, _ := cache.data.PostsInfoByActivity(false,
			cache.indexSize, (page-1)*cache.indexSize)
		return infos
	}

	if cache.indexPostInfo[page-1] == nil {
		_ = cache.indexPostInfoRefresh(page)
	}

	return cache.indexPostInfo[page-1]
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
