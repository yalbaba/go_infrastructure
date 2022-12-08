package go_infrastructure

import (
	"fmt"
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
