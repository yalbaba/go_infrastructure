/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/18
   Description :
-------------------------------------------------
*/

package consts

// web平台
const (
	WebPlatformAdmin      = "admin"       // 管理后台
	WebPlatformSourceOpen = "source_open" // 源开放平台
)

// 设备类型
const (
	DevicePlatformAll     = 10 // 所有
	DevicePlatformAndroid = 20 // 安卓
	DevicePlatformIos     = 30 // ios
	DevicePlatformWeb     = 40 // web
	DevicePlatformApplet  = 50 // 小程序
	DevicePlatformH5      = 60 // H5
)

// 删除状态
const (
	DeletedNo  = 10 // 未删除
	DeletedYes = 20 // 已删除
)

// 开启状态
const (
	StateOff = 10 // 关闭
	StateOn  = 20 // 开启
)

// 是否状态
const (
	No  = 10 // 否
	Yes = 20 // 是
)

const (
	SatelliteTVFirstTagId = 264 // 卫星tv的一级标签
)

//消息推送相关枚举 admin_message_push.device_platform
const (
	IosPush     = 10 //只推ios
	AndroidPush = 20 //只推android
	AllPush     = 30 //推全部
	AppPush     = 40 //只推app内
)

//user_message 文章类型（供系统消息点击跳转的）
const (
	UserMessageLive    = 90
	UserMessageVideo   = 140
	UserMessageTopic   = 150
	UserMessageSubject = 160
)
