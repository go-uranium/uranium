package rcache_test

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"

	rcache "github.com/go-uranium/uranium/cache/redis"
	"github.com/go-uranium/uranium/model/category"
)

var cache *rcache.RCache

type testingStorage struct{}

func (*testingStorage) Categories() ([]*category.Category, error) {
	return nil, nil
}

func (*testingStorage) Init() error  { return nil }
func (*testingStorage) Close() error { return nil }

func Init() error {
	s := &testingStorage{}
	var err error
	cache, err = rcache.New(&redis.Options{}, &rcache.TTLConfig{
		UserBasic: 2 * time.Minute,
		UserUID:   2 * time.Minute,
		Category:  2 * time.Minute,
		Session:   2 * time.Minute,
	}, s)
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	err := cache.FlushAll()
	if err != nil {
		return err
	}
	if err := cache.Close(); err != nil {
		return err
	}
	return nil
}

func TestNew(t *testing.T) {
	if err := Init(); err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if err := Close(); err != nil {
			t.Error(err)
			return
		}
	}()
	result, err := cache.Ping()
	if err != nil {
		t.Error(err)
		return
	}
	if result != "PONG" {
		t.Errorf(`expected "PONG", got "%s"`, result)
	}
}
