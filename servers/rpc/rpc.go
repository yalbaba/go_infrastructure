package rpc

import (
	"context"
	"liveearth/infrastructure/protos/push_stream"
	"net"
	"reflect"
	"time"

	"liveearth/infrastructure/component"
	"liveearth/infrastructure/config"
	"liveearth/infrastructure/consts"
	"liveearth/infrastructure/servers"

	inet "github.com/sereiner/library/net"

	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
)

type RpcServer struct {
	server *grpc.Server
	c      component.Container
}

func (r *RpcServer) GetServerType() consts.ServerType {
	return consts.RpcServer
}

func NewRpcServer(c component.Container) *RpcServer {

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			UnaryServerLogInterceptor(c),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: 10,
		}),
	)

	return &RpcServer{
		server: server,
		c:      c,
	}
}

func (r *RpcServer) Start() error {

	r.c.Debug("开始启动 grpc 服务器...")
	reflection.Register(r.server)

	var addr string
	if config.C.Debug {
		addr = config.C.RPC.Addr
	} else {
		addr = inet.GetLocalIPAddress() + config.C.RPC.Addr
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	errChan := make(chan error, 1)
	go func(errChan chan error) {
		if err := r.server.Serve(listener); err != nil {
			errChan <- err
		}
	}(errChan)

	select {
	case <-time.After(time.Millisecond * 500):
	case err := <-errChan:
		r.c.Error(err.Error())
	}

	r.c.Debugf("grpc 服务器启动成功 addr->[ %s ]", config.C.RPC.Addr)
	return nil
}

func (r *RpcServer) RegisterService(sc ...interface{}) {

	for _, v := range sc {

		fv := reflect.ValueOf(v)
		tp := reflect.TypeOf(v)
		if fv.Kind() != reflect.Func {
			r.c.Error("服务构造器必须是函数")
			return
		}

		if tp.NumIn() != 1 || tp.NumOut() != 1 {
			r.c.Error("服务构造器参数错误")
			return
		}

		if tp.In(0).Name() != "Container" {
			r.c.Error("服务构造器入参类型必须是 component.Container")
			return
		}

		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(r.c)

		rs := fv.Call(params)
		if len(rs) != 1 {
			r.c.Error("服务构造器返回值错误")
			return
		}

		switch t := rs[0].Interface().(type) {
		case push_stream.PushStreamServiceServer:
			push_stream.RegisterPushStreamServiceServer(r.server, t)
		default:
			r.c.Error("未知的服务类型")
			return
		}
	}
}

func (r *RpcServer) Close() error {
	r.server.GracefulStop()
	return nil
}

func UnaryServerLogInterceptor(c component.Container) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		startTime := time.Now()
		c.Info("grpc.request", "method:", info.FullMethod, "args:", req)
		resp, err := handler(ctx, req)

		if err != nil {
			c.Error("grpc.response", "method:", info.FullMethod, err, "耗时:", time.Since(startTime).String())
		} else {
			c.Info("grpc.response", "method:", info.FullMethod, "resp:", resp, err, "耗时:", time.Since(startTime).String())
		}

		return resp, err
	}
}

type rpcServerAdapter struct {
}

func (h *rpcServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewRpcServer(c)
}

func init() {
	servers.Register(consts.RpcServer, &rpcServerAdapter{})
}
