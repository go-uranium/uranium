package rcache

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"

	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/storage"
)

type RCache struct {
	rdb     *redis.Client
	storage storage.Provider
}

var ctx = context.Background()

func New(opt *redis.Options, storage storage.Provider) (*RCache, error) {
	rdb := redis.NewClient(opt)
	if err := rdb.Ping(ctx).Err(); err != nil {
		return &RCache{}, err
	}
	return &RCache{
		rdb:     rdb,
		storage: storage,
	}, nil
}

func (r *RCache) UserBasicByUID(uid int32) (*user.Basic, error) {
	result, err := r.rdb.Get(ctx, "userb:"+strconv.Itoa(int(uid))).Result()
	if err != nil {
		if err == redis.Nil {
			return r.RefreshUserBasicByUID(uid)
		}
	}
	return user.NewBasicFromJSON([]byte(result)), nil
}

func (r *RCache) RefreshUserBasicByUID(uid int32) (*user.Basic, error) {
	basic, err := r.storage.UserBasicByUID(uid)
	if err != nil {
		return &user.Basic{}, err
	}
	js, err := basic.MarshalJSON()
	if err != nil {
		return &user.Basic{}, err
	}
	_, err = r.rdb.Set(ctx, "userb:"+strconv.Itoa(int(uid)), string(js), 0).Result()
	if err != nil {
		return &user.Basic{}, err
	}
	return basic, nil
}
