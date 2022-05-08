package middleware

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/pkg/iris"
)

func Recover(c component.Container) iris.Handler {
	return func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(ctx))
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				ctx.Warn(logMessage)
				ctx.Values().Set("error", fmt.Errorf("panic: %v", err))

				ctx.JSON(map[string]interface{}{
					"errcode": 1,
					"errmsg":  "Internal server error",
				})
				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}

func getRequestLogs(ctx iris.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}
