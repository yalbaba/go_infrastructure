package middleware

import (
	"fmt"
	"time"

	iuser "github.com/yalbaba/go_infrastructure/models/user"
	"github.com/yalbaba/go_infrastructure/pkg/iris"

	"github.com/iris-contrib/middleware/jwt"
	jsoniter "github.com/json-iterator/go"
)

func LoggerMiddleware(serverType string) iris.Handler {
	return func(ctx iris.Context) {
		start := time.Now()
		u := &iuser.UserInfo{}
		token, ok := ctx.Values().Get("jwt").(*jwt.Token)
		if ok {
			bt, _ := jsoniter.Marshal(token.Claims.(jwt.MapClaims))
			_ = jsoniter.Unmarshal(bt, u)
		}
		uuid := getUUID(ctx)
		setUUID(ctx, uuid)

		ctx.Infof(serverType+".request method:%s path:%s Authorization:%s user_id:%s", ctx.Method(), ctx.Path(), ctx.GetHeader("Authorization"), u.UserId)

		ctx.Next()

		code := ctx.GetStatusCode()
		err, ok := ctx.Values().Get("error").(error)
		if err == nil {
			err = fmt.Errorf("nil")
		}
		if ok {
			ctx.Errorf(serverType+".response method:%s path:%s code:%d error:%v 耗时:%s", ctx.Method(), ctx.Path(), code, err, time.Since(start).String())
		} else {
			ctx.Infof(serverType+".response method:%s path:%s code:%d 耗时:%s", ctx.Method(), ctx.Path(), code, time.Since(start).String())
		}
	}
}

func getUUID(c iris.Context) string {

	if v, ok := c.Values().GetEntry("__parrot_sid_"); ok {
		return v.String()
	}

	ck, err := c.Request().Cookie("parrot_sid")
	if err != nil || ck == nil || ck.Value == "" {
		return c.GetSessionID()
	}

	return ck.Value
}
func setUUID(c iris.Context, id string) {
	c.Values().Set("__parrot_sid_", id)
}
