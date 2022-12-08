package go_infrastructure

import (
	"fmt"
	app2 "github.com/yalbaba/go_infrastructure/app"
	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/pkg/iris"
	"github.com/yalbaba/go_infrastructure/servers/http"
	"testing"
)

func TestNewGApp(t *testing.T) {
	app := app2.NewGApp(
		app2.WithPlatName("test"),
		app2.WithAppName("t"),
		app2.WithAPI())

	app.RegisterAPIRouter(func(container component.Container, party iris.Party) {
		test := party.Party("/test")
		test.Get("/123", http.Wrap(testHandler))
	})

	fmt.Println(app.Run())
}

func testHandler(ctx iris.Context) interface{} {
	fmt.Println("11111111")

	return nil
}
