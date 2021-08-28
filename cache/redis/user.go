package rcache

import (
	"database/sql"
	"strconv"

	"github.com/go-redis/redis/v8"

	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/clean"
)

func (r *RCache) UserBasicByUID(uid int32) (*user.Basic, bool, error) {
	result, err := r.rdb.Get(ctx, "userb:"+strconv.Itoa(int(uid))).Result()
	if err != nil {
		if err == redis.Nil {
			return r.refreshUserBasicByUID(uid)
		}
		return &user.Basic{}, false, err
	}
	return user.NewBasicFromJSON([]byte(result)), true, nil
}

func (r *RCache) refreshUserBasicByUID(uid int32) (*user.Basic, bool, error) {
	basic, err := r.storage.UserBasicByUID(uid)
	if err != nil {
		return &user.Basic{}, false, err
	}
	js, err := basic.MarshalJSON()
	if err != nil {
		return &user.Basic{}, false, err
	}
	_, err = r.rdb.Set(ctx, "userb:"+strconv.Itoa(int(uid)), string(js), r.ttl.UserBasic).Result()
	if err != nil {
		return &user.Basic{}, false, err
	}
	return basic, false, nil
}

func (r *RCache) RefreshUserBasicByUID(uid int32) (*user.Basic, error) {
	basic, _, err := r.refreshUserBasicByUID(uid)
	if err == sql.ErrNoRows {
		_, err := r.rdb.Del(ctx, "userb:"+strconv.Itoa(int(uid))).Result()
		if err != nil {
			return &user.Basic{}, err
		}
	}
	return basic, err
}

func (r *RCache) UserUIDByUsername(username string) (int32, bool, error) {
	result, err := r.rdb.Get(ctx, "uid:"+clean.Lowercase(username)).Result()
	if err != nil {
		if err == redis.Nil {
			return r.refreshUserUIDByUsername(username)
		}
		return 0, false, err
	}
	uid, err := strconv.Atoi(result)
	if err != nil {
		return 0, false, err
	}
	return int32(uid), true, err
}

func (r *RCache) refreshUserUIDByUsername(username string) (int32, bool, error) {
	uid, err := r.storage.UserUIDByUsername(username)
	if err != nil {
		return 0, false, err
	}
	uidStr := strconv.Itoa(int(uid))
	_, err = r.rdb.Set(ctx, "uid:"+clean.Lowercase(username), uidStr, r.ttl.UserUID).Result()
	if err != nil {
		return 0, false, err
	}
	return uid, false, nil
}

func (r *RCache) RefreshUserUIDByUsername(username string) (int32, error) {
	uid, _, err := r.refreshUserUIDByUsername(username)
	if err == sql.ErrNoRows {
		_, err := r.rdb.Del(ctx, "uid:"+clean.Lowercase(username)).Result()
		if err != nil {
			return 0, err
		}
	}
	return uid, err
}
