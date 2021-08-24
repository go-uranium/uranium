package rcache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/go-uranium/uranium/storage"
)

type TTLConfig struct {
	UserBasic time.Duration
	UserUID   time.Duration
	Category  time.Duration
	Session   time.Duration
}

type RCache struct {
	rdb                *redis.Client
	ttl                *TTLConfig
	storage            storage.Provider
	cacheCategoryInMem bool
}

var ctx = context.Background()

func New(opt *redis.Options, ttl *TTLConfig, storage storage.Provider) (*RCache, error) {
	rdb := redis.NewClient(opt)
	if err := rdb.Ping(ctx).Err(); err != nil {
		return &RCache{}, err
	}
	ccim := CacheCategoryInMem
	return &RCache{
		rdb:                rdb,
		ttl:                ttl,
		storage:            storage,
		cacheCategoryInMem: ccim,
	}, nil
}
