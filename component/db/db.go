package db

import (
	"fmt"

	"github.com/yalbaba/go_infrastructure/config"

	"github.com/sereiner/library/concurrent/cmap"
	"github.com/sereiner/library/db"
)

type IComponentDB interface {
	GetRegularDB(names ...string) (d *db.DB)
	GetDB(names ...string) (d *db.DB, err error)
	GetDBBy(name string) (c *db.DB, err error)
	SaveDBObject(name string, f func(conf config.DbConfig) (*db.DB, error)) (bool, *db.DB, error)
	Close() error
}

type StandardDB struct {
	name  string
	dbMap cmap.ConcurrentMap
}

func NewStandardDB(name ...string) IComponentDB {
	if len(name) > 0 {
		return &StandardDB{name: name[0], dbMap: cmap.New(2)}
	}
	return &StandardDB{name: "default", dbMap: cmap.New(2)}

}

func (s *StandardDB) GetRegularDB(names ...string) (d *db.DB) {
	d, err := s.GetDB(names...)
	if err != nil {
		panic(err)
	}

	return d
}

func (s *StandardDB) GetDB(names ...string) (d *db.DB, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetDBBy(name)
}

func (s *StandardDB) GetDBBy(name string) (c *db.DB, err error) {

	_, c, err = s.SaveDBObject(name, func(conf config.DbConfig) (*db.DB, error) {

		return db.NewDB(conf.Driver,
			conf.Dsn,
			conf.MaxOpenConns,
			conf.MaxIdleConns,
			600)
	})

	return c, err
}

func (s *StandardDB) SaveDBObject(name string, f func(conf config.DbConfig) (*db.DB, error)) (bool, *db.DB, error) {

	key := fmt.Sprintf("%s/%s", "db", name)

	ok, ch, err := s.dbMap.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {

		dbConf, ok := config.C.DB[name]
		if !ok {
			panic(fmt.Sprintf("数据库配置不存在 name:%s", name))
		}

		return f(dbConf)
	})
	if err != nil {
		err = fmt.Errorf("创建db失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch.(*db.DB), err
}

func (s *StandardDB) Close() error {
	s.dbMap.RemoveIterCb(func(k string, v interface{}) bool {
		v.(*db.DB).Close()
		return true
	})
	return nil
}
