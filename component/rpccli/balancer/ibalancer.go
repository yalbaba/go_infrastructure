package balancer

import "google.golang.org/grpc"

type IBalancer interface {
	GetConn(target string) (cc *grpc.ClientConn, err error) //负载获取rpc服务
	Start(target string) error                              //开启负载均衡器
}

var B IBalancer
