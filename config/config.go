package config

import "time"

var (
	C Config
)

type Config struct {
	Debug   bool   `toml:"debug"`
	LogPath string `toml:"log_path"`
	Schema  string `toml:"schema"` //预留字段：程序启动时选择的运行环境（pre：预发布环境）

	API APIServer `toml:"api"`

	WS WsServer `toml:"ws"`

	RPC GrpcService `toml:"rpc"` //当前rpc服务

	Service map[string]Service `toml:"service"` //其他rpc服务

	DB map[string]DbConfig `toml:"db"`

	Redis map[string]RedisConfig `toml:"redis"`

	Es map[string]EsConfig `toml:"es"`

	MQ map[string]MqConfig `toml:"mq"`

	Mongo map[string]MongoConfig `toml:"mongo"`

	Nsq        map[string]NsqConfig `toml:"nsq"`         // nsq生产者配置
	NsqConsume NsqConsumeConfig     `toml:"nsq_consume"` // nsq消费服务配置

	Registry           Registry         `toml:"registry"`             //etcd注册中心（key为集群名）
	RegisterServerList []RegisterServer `toml:"register_server_list"` //要注册的RPC服务信息集合（目前只有推流服务）
}

type Registry struct {
	EndPoints   []string `toml:"end_points"`   //etcd的集群地址
	UserName    string   `toml:"user_name"`    //etcd用户名（可选）
	Password    string   `toml:"password"`     //etcd的密码（可选）
	DialTimeout int      `toml:"dial_timeout"` //连接注册中心的超时时间
	Balancer    string   `toml:"balancer"`     //负载均衡名称
}

type RegisterServer struct {
	ServiceName string `toml:"service_name"` //服务的名称
	ServerInfo  string `toml:"server_info"`  //服务的信息（json表示）
	TTl         int64  `toml:"ttl"`          //服务的存活时间
}

type Service struct {
	Addr string `toml:"addr"`
}

type DbConfig struct {
	Provider     string `toml:"provider"`
	Driver       string `toml:"driver"`
	Dsn          string `toml:"dsn"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
	ShowSql      bool   `toml:"show_sql"`
}

type RedisConfig struct {
	Network               string        `toml:"network"`
	Addr                  string        `toml:"addr"`
	Password              string        `toml:"password"`
	DialConnectionTimeout time.Duration `toml:"dial_connection_timeout"`
	DialReadTimeout       time.Duration `toml:"dial_read_timeout"`
	DialWriteTimeout      time.Duration `toml:"dial_write_timeout"`
	IdleTimeout           time.Duration `toml:"idle_timeout"`
	DB                    int           `toml:"db"`
	PoolSize              int           `toml:"pool_size"`
}

type EsConfig struct {
	Address       []string `toml:"address"`        // 地址
	UserName      string   `toml:"user_name"`      // 用户名
	Password      string   `toml:"password"`       // 密码
	DialTimeout   int64    `toml:"dial_timeout"`   // 连接超时(毫秒
	Sniff         bool     `toml:"sniff"`          // 嗅探器
	HealthCheck   bool     `toml:"health_check"`   // 心跳检查
	Retry         int      `toml:"retry"`          // 重试次数
	RetryInterval int      `toml:"retry_interval"` // 重试间隔(毫秒)
	GZip          bool     `toml:"g_zip"`          // 启用gzip压缩
}

type MongoConfig struct {
	Address        []string      `toml:"address"`         // 地址
	Database       string        `toml:"database"`        // db
	UserName       string        `toml:"user_name"`       // 用户名
	Password       string        `toml:"password"`        // 密码
	ConnectTimeout time.Duration `toml:"connect_timeout"` // 连接超时(秒
	PoolSize       uint64        `toml:"pool_size"`
	// DoTimeout      uint64   `toml:"do_timeout"`
	SocketTimeout time.Duration `toml:"socket_timeout"` // 秒
	Ping          bool          `toml:"ping"`
}

type MqConfig struct {
	Address string `toml:"address"`
}

type NsqConfig struct {
	Address           string `toml:"address"`            // 地址: localhost:4150
	AuthSecret        string `toml:"auth_secret"`        // 验证秘钥
	HeartbeatInterval int64  `toml:"heartbeat_interval"` // 心跳间隔(毫秒), 不能超过ReadTimeout
	ReadTimeout       int64  `toml:"read_timeout"`       // 超时(毫秒
	WriteTimeout      int64  `toml:"write_timeout"`      // 超时(毫秒
	DialTimeout       int64  `toml:"dial_timeout"`       // 超时(毫秒
}

type NsqConsumeConfig struct {
	Address           string `toml:"address"`             // nsqd地址, localhost1:4150,localhost2:4150
	NsqLookupdAddress string `toml:"nsq_lookupd_address"` // nsq发现服务地址, 优先级高于address, localhost1:4161,localhost2:4161
	AuthSecret        string `toml:"auth_secret"`         // 验证秘钥
	HeartbeatInterval int64  `toml:"heartbeat_interval"`  // 心跳间隔(毫秒), 不能超过ReadTimeout
	ReadTimeout       int64  `toml:"read_timeout"`        // 超时(毫秒)
	WriteTimeout      int64  `toml:"write_timeout"`       // 超时(毫秒)
	DialTimeout       int64  `toml:"dial_timeout"`        // 超时(毫秒)
	MaxInFlight       int    `toml:"max_in_flight"`       // Maximum number of messages to allow in flight (concurrency knob)
	// 线程数, 默认为0表示使用逻辑cpu数量
	//
	// 同时处理信息的goroutine数
	ThreadCount     int    `toml:"thread_count"`
	RequeueDelay    int64  `toml:"requeue_delay"`     // 默认延时时间, 延时时间为-1时和消费失败自动发送延时消息时生效, 实际延时时间=延时时间x尝试次数(毫秒)
	MaxRequeueDelay int64  `toml:"max_requeue_delay"` // 默认最大延时时间, 延时时间为-1时和消费失败自动发送延时消息时生效
	ConsumeAttempts uint16 `toml:"consume_attempts"`  // 消费尝试次数, 默认3, 最大65535
}

type GrpcService struct {
	Addr string `toml:"addr"`
}

type APIServer struct {
	Addr string `toml:"addr"`
}

type WsServer struct {
	Addr string `toml:"addr"`
}
