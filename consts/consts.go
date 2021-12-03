package consts

type ServerType uint8

const (
	HttpServer ServerType = iota + 1
	RpcServer
	MqcServer
	CronServer
	NsqConsumeServer
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
	}

	return ""
}
