package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sereiner/library/concurrent/cmap"
	"liveearth/infrastructure/config"
	"time"
)

type IComponentES interface {
	GetRegularES(names ...string) (d *elastic.Client)
	GetES(names ...string) (d *elastic.Client, err error)
	GetESBy(name string) (c *elastic.Client, err error)
	SaveESObject(name string, f func(conf config.EsConfig) (*elastic.Client, error)) (bool, *elastic.Client, error)
	Close() error
}

type StandardES struct {
	name  string
	esMap cmap.ConcurrentMap
}

func NewStandardES(name ...string) IComponentES {
	if len(name) > 0 {
		es := &StandardES{name: name[0], esMap: cmap.New(2)}
		return es
	}
	return &StandardES{name: "default", esMap: cmap.New(2)}

}

func (s *StandardES) GetRegularES(names ...string) (d *elastic.Client) {
	d, err := s.GetES(names...)
	if err != nil {
		panic(err)
	}

	return d
}

func (s *StandardES) GetES(names ...string) (d *elastic.Client, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetESBy(name)
}

func (s *StandardES) GetESBy(name string) (c *elastic.Client, err error) {

	_, c, err = s.SaveESObject(name, func(conf config.EsConfig) (*elastic.Client, error) {

		opts := []elastic.ClientOptionFunc{
			elastic.SetSniff(conf.Sniff),
			elastic.SetURL(conf.Address...),
			elastic.SetHealthcheck(conf.HealthCheck),
			elastic.SetGzip(conf.GZip),
		}
		if conf.UserName != "" || conf.Password != "" {
			opts = append(opts, elastic.SetBasicAuth(conf.UserName, conf.Password))
		}
		if conf.Retry > 0 {
			ticks := make([]int, conf.Retry)
			for i := 0; i < conf.Retry; i++ {
				ticks[i] = conf.RetryInterval
			}
			elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewSimpleBackoff(ticks...)))
		}

		ctx := context.Background()
		if conf.DialTimeout > 0 {
			ctx, _ = context.WithTimeout(ctx, time.Duration(conf.DialTimeout*1e6))
		}

		c, err := elastic.DialContext(ctx, opts...)
		if err != nil {
			return nil, fmt.Errorf("es连接失败 err:%v", err)
		}

		return c, nil
	})

	return c, err
}

func (s *StandardES) SaveESObject(name string, f func(conf config.EsConfig) (*elastic.Client, error)) (bool, *elastic.Client, error) {

	key := fmt.Sprintf("%s/%s", "es", name)

	ok, ch, err := s.esMap.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {

		esConf, ok := config.C.Es[name]
		if !ok {
			panic(fmt.Sprintf("es配置不存在 name:%s", name))
		}

		return f(esConf)
	})
	if err != nil {
		err = fmt.Errorf("创建es失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch.(*elastic.Client), err
}

func (s *StandardES) Close() error {
	s.esMap.RemoveIterCb(func(k string, v interface{}) bool {
		v.(*elastic.Client).Stop()
		return true
	})
	return nil
}
