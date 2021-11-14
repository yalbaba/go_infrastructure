package middleware

import (
	"fmt"
	iuser "liveearth/infrastructure/models/user"
	"liveearth/infrastructure/pkg/iris"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	jsoniter "github.com/json-iterator/go"
)

func LoggerMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		start := time.Now()
		u := &iuser.UserInfo{}
		token, ok := ctx.Values().Get("jwt").(*jwt.Token)
		if ok {
			bt, _ := jsoniter.Marshal(token.Claims.(jwt.MapClaims))
			_ = jsoniter.Unmarshal(bt, u)
		}

		ctx.Infof("api.request method:%s path:%s Authorization:%s user_id:%s", ctx.Method(), ctx.Path(), ctx.GetHeader("Authorization"), u.UserId)

		ctx.Next()

		code := ctx.GetStatusCode()
		err, ok := ctx.Values().Get("error").(error)
		if err == nil {
			err = fmt.Errorf("nil")
		}
		if ok {
			ctx.Errorf("api.response method:%s path:%s code:%d error:%v 耗时:%s", ctx.Method(), ctx.Path(), code, err, time.Since(start).String())
		} else {
			ctx.Infof("api.response method:%s path:%s code:%d 耗时:%s", ctx.Method(), ctx.Path(), code, time.Since(start).String())
		}
	}
}
