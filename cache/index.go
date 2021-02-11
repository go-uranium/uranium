package cache

import (
	"github.com/go-ushio/ushio/core/post"
)

func (cache *Cache) IndexPostInfo() []*post.Info {
	cache.indexRefresh.RLock()
	defer cache.indexRefresh.RUnlock()
	return cache.indexPostInfo
}

func (cache *Cache) IndexPostInfoRefresh() error {
	cache.indexRefresh.Lock()
	defer cache.indexRefresh.Unlock()
	infos, err := cache.data.PostInfoIndex(cache.indexSize)
	if err != nil {
		return err
	}
	cache.indexPostInfo = infos
	return nil
}
