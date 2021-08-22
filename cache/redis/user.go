package rcache

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"

	"github.com/go-uranium/uranium/model/user"
)

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
	_, err = r.rdb.Set(ctx, "userb:"+strconv.Itoa(int(uid)), string(js), r.ttl.UserBasic).Result()
	if err != nil {
		return &user.Basic{}, err
	}
	return basic, nil
}

func (r *RCache) UserUIDByUsername(username string) (int32, error) {
	result, err := r.rdb.Get(ctx, "uid:"+strings.ToLower(username)).Result()
	if err != nil {
		if err == redis.Nil {
			uid, err := r.RefreshUserUIDByUsername(username)
			if err != nil {
				return 0, err
			}
			return uid, err
		}
	}
	uid, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return int32(uid), err
}

func (r *RCache) RefreshUserUIDByUsername(username string) (int32, error) {
	uid, err := r.storage.UserUIDByUsername(username)
	if err != nil {
		return 0, err
	}
	uidStr := strconv.Itoa(int(uid))
	_, err = r.rdb.Set(ctx, "uid:"+strings.ToLower(username), uidStr, r.ttl.UserUID).Result()
	if err != nil {
		return 0, err
	}
	return uid, nil
}
