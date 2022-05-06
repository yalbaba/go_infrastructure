package component

import (
	"fmt"
	"reflect"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis"
	"github.com/iris-contrib/middleware/jwt"
	jsoniter "github.com/json-iterator/go"
	nnsq "github.com/nsqio/go-nsq"
	"github.com/olivere/elastic/v7"
	"github.com/sereiner/library/db"
	logger "github.com/sereiner/library/log"
	"github.com/sereiner/library/types"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"go_infrastructure/component/cache"
	idb "go_infrastructure/component/db"
	"go_infrastructure/component/es"
	"go_infrastructure/component/mg"
	"go_infrastructure/component/mq"
	"go_infrastructure/component/nsq"
	"go_infrastructure/component/orm"
	"go_infrastructure/component/registry"
	"go_infrastructure/component/rpccli"
	iuser "go_infrastructure/models/user"
	"go_infrastructure/pkg/errno"
	"go_infrastructure/pkg/iris"
	"go_infrastructure/protos/push_stream"
	"go_infrastructure/utils"
)

// Container 容器接口, 需要的组件在这里添加
type Container interface {
	GetRegularDB(names ...string) (d *db.DB)
	GetRegularGorm(names ...string) (d *gorm.DB)
	GetRegularCache(names ...string) (d *redis.Client)
	GetRegularES(names ...string) (d *elastic.Client)
	GetRegularMQ(names ...string) mq.Mq
	GetRegularMongo(names ...string) (d *mongo.Client)
	GetNsq(names ...string) *nnsq.Producer
	Bind(ctx iris.Context, obj interface{}) error
	GetUserInfo(ctx iris.Context) (*iuser.UserInfo, error)
	GetRealIP(ctx iris.Context) string
	RefreshWeight(target string, server string) error

	GetPushStreamServiceClient() push_stream.PushStreamServiceClient
	logger.ILogger
}

// IComponent 组件接口
type IComponent interface {
	GetRegularDB(names ...string) (d *db.DB)
	GetRegularGorm(names ...string) (d *gorm.DB)
	GetRegularCache(names ...string) (d *redis.Client)
	GetRegularES(names ...string) (d *elastic.Client)
	GetRegularMQ(names ...string) mq.Mq
	GetRegularMongo(names ...string) (d *mongo.Client)
	GetNsq(names ...string) *nnsq.Producer
	GetRegistry() registry.IRegistry
	RefreshWeight(target string, server string) error

	Bind(ctx iris.Context, obj interface{}) error
	GetUserInfo(ctx iris.Context) (*iuser.UserInfo, error)
	GetRealIP(ctx iris.Context) string

	GetPushStreamServiceClient() push_stream.PushStreamServiceClient

	logger.ILogger
	Close()
}

type component struct {
	idb.IComponentDB
	orm.IComponentOrm
	cache.IComponentCache
	es.IComponentES
	mq.IComponentMQ
	mg.IComponentMongo
	nsq.IComponentNsq
	rpccli.IComponentRpcClient
	registry.IRegistry
	*logger.Logger
}

func NewComponent(logger *logger.Logger) *component {
	c := &component{
		IComponentDB:        idb.NewStandardDB(),
		IComponentOrm:       orm.NewStandardOrm(),
		IComponentCache:     cache.NewStandardCache(),
		IComponentES:        es.NewStandardES(),
		IComponentMQ:        mq.NewStandardMQ(),
		IComponentMongo:     mg.NewStandardMg(),
		IComponentNsq:       nsq.NewNsqProducer(),
		IComponentRpcClient: rpccli.NewStandardRpcClient(),
		Logger:              logger,
	}

	return c
}

// https://github.com/asaskevich/govalidator
// 字段必填   valid:"required"
// 邮箱   	valid:"email"
// 范围  	valid:"range(min|max)"
// byte长度  valid:"length(min|max)"
// rune长度  valid:"runelength(min|max)"
// string长度 valid:"stringlength(min|max)"
// in  valid:"in(string1|string2|...|stringN)"
func (c *component) Bind(ctx iris.Context, obj interface{}) error {

	bt, err := ctx.GetBody()
	if err != nil {
		return err
	}

	if len(bt) != 0 {
		if err := jsoniter.Unmarshal(bt, obj); err != nil {
			return err
		}
	} else {
		if err := ctx.ReadQuery(obj); err != nil {
			return err
		}
	}

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	if _, err := govalidator.ValidateStruct(obj); err != nil {
		err = errno.New(errno.ErrParam, fmt.Errorf("输入参数有误 %v", err))
		return err
	}
	c.Infof("请求参数: %+v", obj)

	return nil
}

func (c *component) GetRealIP(ctx iris.Context) string {

	ips := ctx.Request().Header.Values("X-Real-Ip")
	if len(ips) == 0 {
		ips = ctx.Request().Header.Values("X-Forwarded-For")
	}

	if len(ips) > 0 {
		return ips[0]
	}

	return "127.0.0.1"
}

func (c *component) GetUserInfo(ctx iris.Context) (*iuser.UserInfo, error) {

	deviceData := &iuser.DeviceData{
		DeviceId:    ctx.GetHeader("device_id"),
		Version:     ctx.GetHeader("version"),
		VersionNum:  types.GetInt(ctx.GetHeader("version_num")),
		Channel:     ctx.GetHeader("channel"),
		DeviceType:  ctx.GetHeader("device_type"),
		DeviceBrand: ctx.GetHeader("device_brand"),
		Platform:    types.GetInt(ctx.GetHeader("platform")),
	}

	token, ok := ctx.Values().Get("jwt").(*jwt.Token)
	if !ok {
		return &iuser.UserInfo{DeviceData: deviceData}, nil
	}

	bt, err := jsoniter.Marshal(token.Claims.(jwt.MapClaims))
	if err != nil {
		return nil, err
	}

	u := &iuser.UserInfo{}
	if err := jsoniter.Unmarshal(bt, u); err != nil {
		return nil, err
	}
	u.DeviceData = deviceData

	return u, nil
}

func (c *component) GetRegistry() registry.IRegistry {
	return c.IRegistry
}

func (c *component) RefreshWeight(target string, server string) error {
	return c.GetRegistry().RefreshWeight(target, server)
}

func (c *component) Close() {
	g := errgroup.Group{}
	g.Go(func() error {
		return c.IComponentDB.Close()
	})
	g.Go(func() error {
		return c.IComponentCache.Close()
	})
	g.Go(func() error {
		return c.IComponentRpcClient.Close()
	})
	g.Go(func() error {
		c.IComponentNsq.Close()
		return nil
	})
	g.Go(func() error {
		if c.IRegistry == nil {
			return nil
		}
		return c.IRegistry.Close()
	})
	_ = g.Wait()

}

func init() {
	govalidator.TagMap["time"] = func(str string) bool {
		_, err := time.ParseInLocation(utils.DefaultLayout, str, time.Local)
		return err == nil
	}
	govalidator.TagMap["date"] = func(str string) bool {
		_, err := time.ParseInLocation(utils.DateLayout, str, time.Local)
		return err == nil
	}
}
