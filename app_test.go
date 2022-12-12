package go_infrastructure

import (
	"fmt"
	"github.com/mikegleasonjr/workers"
	"github.com/yalbaba/go_infrastructure/app"
	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/pkg/iris"
	"github.com/yalbaba/go_infrastructure/rpc_clients"
	"github.com/yalbaba/go_infrastructure/servers/http"
	"testing"
)

// 测试http服务
func TestHttpServer(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithAPI())

	myapp.RegisterAPIRouter(func(container component.Container, party iris.Party) {
		test := party.Party("/test")
		test.Get("/123", http.Wrap(testHandler))
	})

	fmt.Println(myapp.Run())
}

func testHandler(ctx iris.Context) interface{} {
	fmt.Println("11111111")

	return nil
}

// 测试rpc服务
func TestRpcServer(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithGRPC())

	myapp.RegisterRpcService(rpc_clients.NewTestRpcServer)
	fmt.Println(myapp.Run())
}

// 测试消息队列服务
func TestMqcServer(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithMQC())

	install(myapp)

	fmt.Println(myapp.Run())
}

func install(a app.IApp) {
	t := NewTestMqcService(a.GetContainer())
	a.RegisterMqcWorker("test-topic", t.TestMqcHandler)
}

type TestMqcService struct {
	c component.Container
}

func NewTestMqcService(c component.Container) *TestMqcService {
	return &TestMqcService{c: c}
}

func (s *TestMqcService) TestMqcHandler(job *workers.Job) {
	defer job.Delete()
	s.c.Debug("job begin")
}

func TestMqcClient(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithMQC())

	myapp.GetContainer().GetRegularMQ().Send("test-topic", []byte(""), 1, 0, 0)
}

// 测试长连接websocket协议服务
func TestWsServer(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithWs())
	myapp.RegisterWsRouter("/test/ws", wsHandler)

	fmt.Println(myapp.Run())
}

func wsHandler(ctx iris.Context, message []byte) interface{} {

	fmt.Println("message      ", string(message))

	return nil
}

// 测试定时任务
func TestCronJob(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithCron())
	// @every 1s 1m 1h
	//myapp.RegisterCronJob("test-cron-job", "@every 10s", false, func(ctx iris.Context) (err error) {
	//	fmt.Println(111111111)
	//	return
	//})

	// cron表达式
	myapp.RegisterCronJob("test-cron-job", "0 1 * * *", false, func(ctx iris.Context) (err error) {
		fmt.Println(2222222)
		return
	})

	fmt.Println(myapp.Run())
}
