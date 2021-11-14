/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package orm

import (
	"fmt"
	"liveearth/infrastructure/config"

	"gorm.io/driver/mysql"

	"github.com/sereiner/library/concurrent/cmap"
	"gorm.io/gorm"
)

type IComponentOrm interface {
	GetRegularGorm(names ...string) (d *gorm.DB)
	GetOrm(names ...string) (d *gorm.DB, err error)
	GetOrmBy(name string) (c *gorm.DB, err error)
	SaveOrmObject(name string, f func(conf config.DbConfig) (*gorm.DB, error)) (bool, *gorm.DB, error)
	Close() error
}

type StandardOrm struct {
	name  string
	dbMap cmap.ConcurrentMap
}

func NewStandardOrm(name ...string) IComponentOrm {

	if len(name) > 0 {
		return &StandardOrm{name: name[0], dbMap: cmap.New(2)}
	}
	return &StandardOrm{name: "default", dbMap: cmap.New(2)}

}

func (s *StandardOrm) GetRegularGorm(names ...string) (d *gorm.DB) {
	d, err := s.GetOrm(names...)
	if err != nil {
		panic(err)
	}

	return d
}

func (s *StandardOrm) GetOrm(names ...string) (d *gorm.DB, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetOrmBy(name)
}

func (s *StandardOrm) GetOrmBy(name string) (c *gorm.DB, err error) {

	_, c, err = s.SaveOrmObject(name, func(conf config.DbConfig) (*gorm.DB, error) {

		return gorm.Open(mysql.New(mysql.Config{
			DriverName: conf.Driver,
			DSN:        conf.Dsn,
		}), &gorm.Config{})

	})

	return c, err
}

func (s *StandardOrm) SaveOrmObject(name string, f func(conf config.DbConfig) (*gorm.DB, error)) (bool, *gorm.DB, error) {

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

	return ok, ch.(*gorm.DB), err
}

func (s *StandardOrm) Close() error {
	return nil
}
