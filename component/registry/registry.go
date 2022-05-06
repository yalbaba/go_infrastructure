package registry

import (
	"context"
	"fmt"
	"go_infrastructure/component/rpccli/balancer/smooth_roundrobin"
	"go_infrastructure/config"
	"go_infrastructure/consts"
	"go_infrastructure/utils"

	jsoniter "github.com/json-iterator/go"
	"github.com/ozonru/etcd/clientv3"
	logger "github.com/sereiner/library/log"
	inet "github.com/sereiner/library/net"
	"google.golang.org/grpc/grpclog"
)

type IRegistry interface {
	Register() error
	Close() error
	RefreshWeight(target string, serverName string) error
}

const (
	SmoothRoundRobin = "smooth_roundrobin"
)

type EtcdRegistry struct {
	Cli *clientv3.Client
	Lci clientv3.Lease
	log *logger.Logger
}

func NewRegistry(l *logger.Logger) IRegistry {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: config.C.Registry.EndPoints,
		//DialTimeout: time.Duration(config.C.Registry.DialTimeout),
		Username: config.C.Registry.UserName,
		Password: config.C.Registry.Password,
	})
	if err != nil {
		panic(fmt.Errorf("连接注册中心失败,err: %v", err))
	}

	//初始化负载均衡器
	switch config.C.Registry.Balancer {
	case SmoothRoundRobin:
		smooth_roundrobin.NewSmoothRoundRobin(cli, l)
	default:
		panic(fmt.Errorf("balancer not found,type:%s", config.C.Registry.Balancer))
	}

	return &EtcdRegistry{
		Cli: cli,
		Lci: clientv3.NewLease(cli),
	}
}

func (r *EtcdRegistry) Register() error {
	var addr string
	if config.C.Debug {
		addr = "127.0.0.1" + config.C.RPC.Addr
	} else {
		addr = inet.GetLocalIPAddress() + config.C.RPC.Addr
	}

	for _, v := range config.C.RegisterServerList {
		key := consts.PushStream + "_" + v.ServiceName

		serverInfo := struct {
			Using int    `json:"using"`
			Rest  int    `json:"rest"`
			Addr  string `json:"addr"`
		}{}
		err := jsoniter.UnmarshalFromString(v.ServerInfo, &serverInfo)
		if err != nil {
			return fmt.Errorf("服务信息格式错误,err:%v,source:%s", err, v.ServerInfo)
		}

		serverInfo.Addr = addr

		bytes, _ := jsoniter.Marshal(serverInfo)

		lRsp, err := r.Lci.Grant(context.Background(), v.TTl)
		if err != nil {
			return fmt.Errorf("获取租约异常,err:%v", err)
		}

		opts := []clientv3.OpOption{clientv3.WithLease(lRsp.ID)}
		_, err = r.Cli.KV.Put(context.Background(), key, string(bytes), opts...)
		if err != nil {
			return fmt.Errorf("注册服务异常,err:%v", err)
		}

		lsRspChan, err := r.Lci.KeepAlive(context.Background(), lRsp.ID)
		if err != nil {
			return err
		}
		go func() {
			for {
				_, ok := <-lsRspChan
				if !ok {
					grpclog.Fatalf("%v 服务正在关闭", key)
					break
				}
			}
		}()
	}

	return nil
}

func (r *EtcdRegistry) RefreshWeight(target string, serverName string) error {

	switch target {
	case consts.PushStream:
		err := r.refreshPushStreamWeight(serverName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *EtcdRegistry) Close() error {
	err := r.Cli.Close()
	if err != nil {
		return fmt.Errorf("关闭注册中心客户端失败,err:%v", err)
	}

	err = r.Lci.Close()
	if err != nil {
		return fmt.Errorf("关闭注册中心租约客户端失败,err:%v", err)
	}

	return nil
}

/*
以下为各种服务的权重刷新实现
*/

func (r *EtcdRegistry) refreshPushStreamWeight(serverName string) error {
	key := utils.GetPushStreamKey(serverName)
	resp, err := r.Cli.Get(context.Background(), key)
	if err != nil {
		return fmt.Errorf("RefreshWeight 获取target服务失败,err:%v,key:%s", err, key)
	}

	serverInfo := struct {
		Using int    `json:"using"`
		Rest  int    `json:"rest"`
		Addr  string `json:"addr"`
	}{}
	_ = jsoniter.UnmarshalFromString(string(resp.Kvs[0].Value), &serverInfo)
	serverInfo.Rest -= 1
	serverInfo.Using += 1
	val, _ := jsoniter.MarshalToString(serverInfo)
	_, err = r.Cli.Put(context.Background(), key, val)
	if err != nil {
		return fmt.Errorf("refreshPushStreamWeight 更新注册中心服务信息失败,err:%v,key:%s", err, key)
	}

	return nil
}
