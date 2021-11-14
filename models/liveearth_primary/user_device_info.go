package liveearth_primary

import (
	"time"
)

type UserDeviceInfo struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('pk') INT(10)"`
	UserId             string    `json:"user_id" xorm:"default '' comment('用户id') VARCHAR(32)"`
	DeviceId           string    `json:"device_id" xorm:"not null comment('设备id') VARCHAR(32)"`
	DeviceToken        string    `json:"device_token" xorm:"not null comment('设备token，用于推送') VARCHAR(256)"`
	DeviceType         string    `json:"device_type" xorm:"not null comment('设备型号') VARCHAR(16)"`
	Pushed             int       `json:"pushed" xorm:"not null comment('用户是否开启推送  10 = 否 ; 20 = 是;') TINYINT(4)"`
	UserPushed         int       `json:"user_pushed" xorm:"not null default 20 comment('是否接收用户互动消息推送 10不接收 20接收') TINYINT(4)"`
	SystemPushed       int       `json:"system_pushed" xorm:"not null default 20 comment('是否接收系统的消息推送 10不接收 20接收') TINYINT(4)"`
	PopularPushed      int       `json:"popular_pushed" xorm:"not null default 20 comment('是否接收系统的消息推送 10不接收 20接收') TINYINT(4)"`
	DevicePlatform     int       `json:"device_platform" xorm:"comment('平台: 安卓= 20; iOS = 30; web = 40') VARCHAR(16)"`
	DeviceBrand        string    `json:"device_brand" xorm:"comment('手机生产厂商') VARCHAR(8)"`
	AppVersion         string    `json:"app_version" xorm:"not null comment('当前设备上的app版本') VARCHAR(8)"`
	LoginState         int       `json:"login_state" xorm:"not null comment('登录状态  10 = 未登录; 20 = 登录') TINYINT(4)"`
	Lon                float64   `json:"lon" xorm:"not null default 0 comment('经度') DOUBLE"`
	Lat                float64   `json:"lat" xorm:"not null default 0 comment('纬度') DOUBLE"`
	Nation             string    `json:"nation" xorm:"not null default '' comment('国家') VARCHAR(32)"`
	Province           string    `json:"province" xorm:"not null default '' comment('省') VARCHAR(32)"`
	City               string    `json:"city" xorm:"not null default '' comment('城市') VARCHAR(32)"`
	Country            string    `json:"country" xorm:"not null default '' comment('区县') VARCHAR(32)"`
	Adcode             int       `json:"adcode" xorm:"not null default '' comment('区域码') VARCHAR(32)"`
	DistrictId         int       `json:"district_id"`
	DownloadSource     string    `json:"download_source" xorm:"not null default '' comment('下载来源') VARCHAR(16)"`
	CreateTime         time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime         time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	DeviceTokenReverse string    `json:"device_token_reverse"`
}

func (*UserDeviceInfo) TableName() string {
	return "liveearth_primary.user_device_info"
}
