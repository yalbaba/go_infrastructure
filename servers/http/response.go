package http

import (
	"go_infrastructure/pkg/errno"
	"go_infrastructure/pkg/iris"
)

type Handler func(ctx iris.Context) interface{}

type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data,omitempty"`
}

func Wrap(handler Handler) iris.Handler {
	return func(ctx iris.Context) {
		result := handler(ctx)
		WriteCtx(ctx, result)
		ctx.StopExecution()
	}
}

// 包装中间件
func WrapMiddleware(middleware Handler) iris.Handler {
	return func(ctx iris.Context) {
		result := middleware(ctx)
		if result == nil {
			ctx.Next()
			return
		}

		WriteCtx(ctx, result)
		ctx.StopExecution()
	}
}

func WriteCtx(ctx iris.Context, result interface{}) {
	var code int
	message := "OK"

	if v, ok := result.(error); ok {
		code, message = errno.DecodeErr(v)
		result = nil
		ctx.Values().Set("error", v)
	}

	// always return http.StatusOK
	ctx.JSON(Response{
		ErrCode: code,
		ErrMsg:  message,
		Data:    result,
	})
}
