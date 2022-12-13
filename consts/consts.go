package consts

type LikeOp uint8

type ServerType uint8

const (
	HttpServer ServerType = iota + 1
	RpcServer
	BeanstalkServer
	CronServer
	NsqConsumeServer
	WsServer
)

func (s ServerType) String() string {
	switch s {
	case HttpServer:
		return "http"
	case RpcServer:
		return "grpc"
	case BeanstalkServer:
		return "beanstalk"
	case CronServer:
		return "cron"
	case NsqConsumeServer:
		return "nsq"
	case WsServer:
		return "ws"
	}

	return ""
}
