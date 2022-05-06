package consts

type LikeOp uint8

type ServerType uint8

const (
	HttpServer ServerType = iota + 1
	RpcServer
	MqcServer
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
	case MqcServer:
		return "mqc"
	case CronServer:
		return "cron"
	case NsqConsumeServer:
		return "nsq_consume"
	case WsServer:
		return "ws"
	}

	return ""
}
