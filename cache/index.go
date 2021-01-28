package cache

import (
	"github.com/go-ushio/ushio/core/post"
)

func (cache *Cache) IndexPostInfo(size int) ([]*post.Info, error) {
	if len(cache.indexPostInfo) < size {
		if cache.postsNotEnough {
			return cache.indexPostInfo, nil
		}
		infos, err := cache.data.PostInfoByPage(0, size)
		if err != nil {
			return nil, err
		}
		cache.indexPostInfo = infos
		if len(infos) < size {
			cache.postsNotEnough = true
		} else {
			cache.postsNotEnough = false
		}
		return infos, nil
	} else {
		return cache.indexPostInfo[0:size], nil
	}
}

func (cache *Cache) IndexPostInfoDrop() error {
	cache.indexPostInfo = []*post.Info{}
	cache.postsNotEnough = false
	return nil
}
