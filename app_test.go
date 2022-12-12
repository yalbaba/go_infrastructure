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

func TestRpcServer(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithGRPC())

	myapp.RegisterRpcService(rpc_clients.NewTestRpcServer)
	fmt.Println(myapp.Run())
}

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

func TestWsServer(t *testing.T) {
	myapp := app.NewGApp(
		app.WithPlatName("test"),
		app.WithAppName("t"),
		app.WithWs())
	myapp.RegisterWs("/test/ws", wsHandler)

	fmt.Println(myapp.Run())
}

func wsHandler(ctx iris.Context, message []byte) interface{} {

	fmt.Println("message      ", string(message))

	return nil
}
