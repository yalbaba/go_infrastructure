package middleware

import (
	"go_infrastructure/models"
	"go_infrastructure/pkg/errno"
	"go_infrastructure/pkg/iris"
	"go_infrastructure/servers/http/middleware/jwt"
	"strings"
)

// MustCheckAuth 必须检查Auth
func MustCheckAuth() iris.Handler {
	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return models.JwtSecret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return func(ctx iris.Context) {
		routePath := strings.TrimPrefix(ctx.Path(), "/rest/api/v5")
		ctx.Warn("routePath:::", routePath)
		//如果是登录相关直接放行
		if routePath == "/admin/admin_user/login" {
			ctx.Warn("登录不需要鉴权")
			ctx.Next()
			return
		}

		if err := j.CheckJWT(ctx); err != nil {

			ctx.Warn(errno.AuthorizationRequired.Message)
			ctx.Values().Set("error", errno.AuthorizationRequired)
			ctx.JSON(map[string]interface{}{
				"errcode": errno.AuthorizationRequired.Code,
				"errmsg":  errno.AuthorizationRequired.Message,
			})
			ctx.StopExecution()
			return
		}
		ctx.Next()
	}

}

func CheckAuth() iris.Handler {
	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return models.JwtSecret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return func(ctx iris.Context) {
		if err := j.CheckJWT(ctx); err == nil {
			j.Serve(ctx)
		}
		ctx.Next()
	}
}
