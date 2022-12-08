package rpc_clients

import (
	"context"
	"fmt"
	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/protos/rpc_servers"
)

var _ rpc_servers.TestRpcServerServiceServer = &TestRpcService{}

type TestRpcService struct {
	container component.Container
}

func NewTestRpcServer(container component.Container) *TestRpcService {
	return &TestRpcService{container: container}
}

func (t TestRpcService) TestInvoke(ctx context.Context, req *rpc_servers.TestInvokeReq) (*rpc_servers.TestInvokeResp, error) {

	fmt.Println("req      ", req)

	return &rpc_servers.TestInvokeResp{ServiceName: req.StreamName}, nil
}
