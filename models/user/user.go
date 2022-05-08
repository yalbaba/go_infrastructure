package user

import (
	"github.com/zlyuancn/zstr"

	"github.com/yalbaba/go_infrastructure/models"

	"github.com/iris-contrib/middleware/jwt"
	jsoniter "github.com/json-iterator/go"
)

type UserInfo struct {
	UserId   string `json:"user_id"`   // 用户id
	IsLogin  bool   `json:"is_login"`  // 用户是否登录
	UserName string `json:"user_name"` // 用户名
	*DeviceData
}

// 将id转为int
func (u *UserInfo) IntId() int {
	return zstr.ToIntDefault(u.UserId)
}

type DeviceData struct {
	DeviceId    string `json:"device_id"`    // 设备id
	Version     string `json:"version"`      // app版本
	VersionNum  int    `json:"version_num"`  // 内部版本号
	Channel     string `json:"channel"`      // 手机下载渠道
	DeviceType  string `json:"device_type"`  // 手机型号：如 min 9
	DeviceBrand string `json:"device_brand"` // 手机厂商 ：如 huawei,ios，xiaomi
	Platform    int    `json:"platform"`     // 平台: 20=安卓; 30=ios; 40=web
}

func GetToken(userInfo *UserInfo) string {

	var m jwt.MapClaims
	s, _ := jsoniter.Marshal(userInfo)
	_ = jsoniter.Unmarshal(s, &m)

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, m)
	tokenString, _ := token.SignedString(models.JwtSecret)

	return tokenString
}

func GetTokenFromParms(userId string, isLogin bool) string {
	parms := &UserInfo{
		UserId:  userId,
		IsLogin: isLogin,
	}
	return GetToken(parms)
}
