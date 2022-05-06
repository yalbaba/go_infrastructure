package cache

import (
	"fmt"
	"go_infrastructure/config"
	"time"

	"github.com/go-redis/redis"
	"github.com/sereiner/library/concurrent/cmap"
)

type IComponentCache interface {
	GetRegularCache(names ...string) (d *redis.Client)
	GetCache(names ...string) (d *redis.Client, err error)
	GetCacheBy(name string) (c *redis.Client, err error)
	SaveCacheObject(name string, f func(conf config.RedisConfig) (*redis.Client, error)) (bool, *redis.Client, error)
	Close() error
}

type StandardCache struct {
	name     string
	cacheMap cmap.ConcurrentMap
}

func NewStandardCache(name ...string) IComponentCache {
	if len(name) > 0 {
		return &StandardCache{name: name[0], cacheMap: cmap.New(2)}
	}
	return &StandardCache{name: "default", cacheMap: cmap.New(2)}
}

func (s *StandardCache) GetRegularCache(names ...string) (d *redis.Client) {
	d, err := s.GetCache(names...)
	if err != nil {
		panic(err)
	}

	return d
}

func (s *StandardCache) GetCache(names ...string) (d *redis.Client, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetCacheBy(name)
}

func (s *StandardCache) GetCacheBy(name string) (c *redis.Client, err error) {

	_, c, err = s.SaveCacheObject(name, func(conf config.RedisConfig) (*redis.Client, error) {

		return redis.NewClient(&redis.Options{
			Network:      conf.Network,
			Addr:         conf.Addr,
			Password:     conf.Password,
			DB:           conf.DB,
			DialTimeout:  time.Second * conf.DialConnectionTimeout,
			ReadTimeout:  time.Second * conf.DialReadTimeout,
			WriteTimeout: time.Second * conf.DialWriteTimeout,
			IdleTimeout:  time.Second * conf.IdleTimeout,
			PoolSize:     conf.PoolSize,
		}), nil
	})

	return c, err
}

func (s *StandardCache) SaveCacheObject(name string, f func(conf config.RedisConfig) (*redis.Client, error)) (bool, *redis.Client, error) {

	key := fmt.Sprintf("%s/%s", "cache", name)

	ok, ch, err := s.cacheMap.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {
		redisConfig, ok := config.C.Redis[name]
		if !ok {
			panic(fmt.Sprintf("redis配置不存在 name:%s", name))
		}

		return f(redisConfig)
	})
	if err != nil {
		err = fmt.Errorf("创建cahce失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch.(*redis.Client), err
}

func (s *StandardCache) Close() error {
	s.cacheMap.RemoveIterCb(func(k string, v interface{}) bool {
		v.(*redis.Client).Close()
		return true
	})
	return nil
}
