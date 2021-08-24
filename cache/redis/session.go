package rcache

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/go-uranium/uranium/model/session"
)

var (
	ErrCannotDecode = errors.New("cannot decode the session in redis")
)

func (r *RCache) ValidSessionByToken(token string) (*session.Cache, bool, error) {
	if len(token) == 0 {
		return &session.Cache{}, false, sql.ErrNoRows
	}
	result, err := r.rdb.Get(ctx, "session:"+token).Result()
	if err != nil {
		if err == redis.Nil {
			return r.refreshValidSessionByToken(token)
		}
		return &session.Cache{}, false, err
	}
	return decodeSession(result)
}

func (r *RCache) refreshValidSessionByToken(token string) (*session.Cache, bool, error) {
	sess, err := r.storage.SessionBasicByToken(token)
	if err != nil {
		return &session.Cache{}, false, err
	}
	d := minDuration(r.ttl.Session, time.Until(sess.Expire))
	if d <= time.Second {
		return &session.Cache{}, false, nil
	}
	sc := &session.Cache{
		UID:   sess.UID,
		Mode:  sess.Mode,
		Valid: true,
	}
	_, err = r.rdb.Set(ctx, "session:"+token, encodeSession(sc), d).Result()
	return sc, false, err
}

func (r *RCache) RefreshValidSessionByToken(token string) (*session.Cache, error) {
	sc, _, err := r.refreshValidSessionByToken(token)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := r.rdb.Del(ctx, "session:"+token).Result()
			if err != nil {
				return &session.Cache{}, err
			}
		}
		return &session.Cache{}, err
	}
	return sc, err
}

func encodeSession(sess *session.Cache) string {
	return string(encodeMode(sess.Mode)) + strconv.Itoa(int(sess.UID))
}

func decodeSession(str string) (*session.Cache, bool, error) {
	if len(str) < 2 {
		return &session.Cache{}, false, ErrCannotDecode
	}
	uid, err := strconv.Atoi(str[1:])
	if err != nil {
		return &session.Cache{}, false, err
	}
	return &session.Cache{
		UID:   int32(uid),
		Mode:  decodeMode(str[0]),
		Valid: true,
	}, true, nil
}

func decodeMode(b byte) int16 {
	return int16(b - 48)
}

func encodeMode(mode int16) byte {
	return byte(mode + 48)
}

func minDuration(d1, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d1
	}
	return d2
}
