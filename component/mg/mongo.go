package mg

import (
	"context"
	"fmt"
	"time"

	"github.com/yalbaba/go_infrastructure/config"

	"github.com/sereiner/library/concurrent/cmap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IComponentMongo interface {
	GetRegularMongo(names ...string) (d *mongo.Client)
	GetMongo(names ...string) (d *mongo.Client, err error)
	GetMongoBy(name string) (c *mongo.Client, err error)
	SaveMongoObject(name string, f func(conf config.MongoConfig) (*mongo.Client, error)) (bool, *mongo.Client, error)
	Close() error
}

type StandardMg struct {
	name  string
	mgMap cmap.ConcurrentMap
}

func NewStandardMg(name ...string) IComponentMongo {
	if len(name) > 0 {
		mg := &StandardMg{name: name[0], mgMap: cmap.New(2)}
		return mg
	}
	return &StandardMg{name: "default", mgMap: cmap.New(2)}
}

func (s *StandardMg) GetRegularMongo(names ...string) (d *mongo.Client) {
	d, err := s.GetMongo(names...)
	if err != nil {
		panic(err)
	}

	return d
}

func (s *StandardMg) GetMongo(names ...string) (d *mongo.Client, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetMongoBy(name)
}

func (s *StandardMg) GetMongoBy(name string) (c *mongo.Client, err error) {

	_, c, err = s.SaveMongoObject(name, func(conf config.MongoConfig) (*mongo.Client, error) {
		var sockTimeout time.Duration = conf.SocketTimeout * time.Second
		var conTimeout time.Duration = conf.ConnectTimeout * time.Second
		opts := &options.ClientOptions{
			Hosts:          conf.Address,
			MaxPoolSize:    &conf.PoolSize,
			SocketTimeout:  &sockTimeout,
			ConnectTimeout: &conTimeout,
		}
		if conf.UserName != "" && conf.Password != "" {
			var cred = options.Credential{
				Username:   conf.UserName,
				Password:   conf.Password,
				AuthSource: conf.Database,
				//AuthMechanism: "SCRAM-SHA-1",
			}
			opts.SetAuth(cred)
		}
		c, err := mongo.Connect(context.Background(), opts)
		if err != nil {
			return nil, fmt.Errorf("mongo连接失败 err:%v", err)
		}
		//if err := c.Ping(context.Background(), nil); err != nil {
		//	return nil, fmt.Errorf("mongo Ping err:%v", err)
		//}
		return c, nil
	})

	return c, err
}

func (s *StandardMg) SaveMongoObject(name string, f func(conf config.MongoConfig) (*mongo.Client, error)) (bool, *mongo.Client, error) {

	key := fmt.Sprintf("%s/%s", "mongo", name)

	ok, ch, err := s.mgMap.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {

		mgConf, ok := config.C.Mongo[name]
		if !ok {
			panic(fmt.Sprintf("mongo配置不存在 name:%s", name))
		}

		return f(mgConf)
	})
	if err != nil {
		err = fmt.Errorf("创建mongo失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch.(*mongo.Client), err
}

func (s *StandardMg) Close() error {
	s.mgMap.RemoveIterCb(func(k string, v interface{}) bool {
		v.(*mongo.Client).Disconnect(context.Background())
		return true
	})
	return nil
}
