package rpccli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sereiner/library/concurrent/cmap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	"github.com/yalbaba/go_infrastructure/component/rpccli/balancer"
	"github.com/yalbaba/go_infrastructure/config"
	"github.com/yalbaba/go_infrastructure/consts"
	"github.com/yalbaba/go_infrastructure/protos/test"
)

type IComponentRpcClient interface {
	GetPushStreamServiceClient() push_stream.PushStreamServiceClient

	GetClientByBalancer(names ...string) (r interface{}, err error)
	GetClient(names ...string) (r interface{}, err error)
	GetClientBy(name string) (r interface{}, err error)
	SaveClientObject(name string, f func(conf config.Service) (interface{}, error)) (bool, interface{}, error)
	Close() error
}

type StandardRpcClient struct {
	name   string
	rpcCli cmap.ConcurrentMap
	ccMap  map[string]*grpc.ClientConn
}

func NewStandardRpcClient(name ...string) IComponentRpcClient {
	if len(name) > 0 {
		return &StandardRpcClient{
			name:   name[0],
			rpcCli: cmap.New(2),
			ccMap:  map[string]*grpc.ClientConn{},
		}
	}
	return &StandardRpcClient{
		name:   "default",
		rpcCli: cmap.New(2),
		ccMap:  map[string]*grpc.ClientConn{},
	}
}

func (s *StandardRpcClient) GetPushStreamServiceClient() push_stream.PushStreamServiceClient {
	var (
		r   interface{}
		err error
	)
	//判断是否开启负载均衡
	if config.C.Registry.Balancer == "" {
		//没开启走原来的流程
		r, err = s.GetClient(consts.PushStream)
		if err != nil {
			panic(err)
		}
	} else {
		r, err = s.GetClientByBalancer(consts.PushStream)
		if err != nil {
			panic(err)
		}
	}

	v, ok := r.(push_stream.PushStreamServiceClient)
	if !ok {
		panic("PushStreamServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetClientByBalancer(names ...string) (r interface{}, err error) {

	name := s.name
	if len(names) > 0 {
		name = names[0]
	}

	cc, err := balancer.B.GetConn(name)
	if err != nil {
		return nil, fmt.Errorf("负载均衡获取连接错误,err:%v,name:%s", err, name)
	}

	s.ccMap[name] = cc

	switch name {
	case consts.PushStream:
		r = push_stream.NewPushStreamServiceClient(cc)
	}

	return r, nil
}

func (s *StandardRpcClient) GetClient(names ...string) (r interface{}, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetClientBy(name)
}

func (s *StandardRpcClient) GetClientBy(name string) (r interface{}, err error) {

	_, c, err := s.SaveClientObject(name, func(conf config.Service) (interface{}, error) {

		re := balancer.NewResolver("live", strings.Split(conf.Addr, ","))
		resolver.Register(re)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		cc, err := grpc.DialContext(
			ctx,
			re.Scheme()+"://authority/",
			grpc.WithBalancerName(roundrobin.Name),
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			return nil, err
		}
		s.ccMap[name] = cc
		var r interface{}
		switch name {
		case consts.PushStream:
			r = push_stream.NewPushStreamServiceClient(cc)
		}

		return r, nil
	})

	return c, err
}

func (s *StandardRpcClient) SaveClientObject(name string, f func(conf config.Service) (interface{}, error)) (bool, interface{}, error) {

	key := fmt.Sprintf("%s/%s", "rpc", name)

	ok, ch, err := s.rpcCli.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {

		clientConf, ok := config.C.Service[name]
		if !ok {
			return nil, fmt.Errorf("grpc客户端配置不存在 name:%s", name)
		}

		return f(clientConf)
	})
	if err != nil {
		err = fmt.Errorf("创建rpc客户端失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch, err
}

func (s *StandardRpcClient) Close() error {
	for _, v := range s.ccMap {
		return v.Close()
	}

	return nil
}
