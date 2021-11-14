package liveearth_primary

import "time"

type DataTopChannelInfo struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Huawei      int       `json:"huawei" gorm:"column:huawei"`             // 华为（存比例数据乘以100）
	AppTreasure int       `json:"app_treasure" gorm:"column:app_treasure"` // 应用宝
	Xiaomi      int       `json:"xiaomi" gorm:"column:xiaomi"`             // 小米
	Vivo        int       `json:"vivo" gorm:"column:vivo"`                 // vivo
	Oppo        int       `json:"oppo" gorm:"column:oppo"`                 // oppo
	Headline    int       `json:"headline" gorm:"column:headline"`         // 头条
	Baidu       int       `json:"baidu" gorm:"column:baidu"`               // 百度
	Offical     int       `json:"offical" gorm:"column:offical"`           // 官方
	Other       int       `json:"other" gorm:"column:other"`               // 其他
	Appstore    int       `json:"appstore" gorm:"column:appstore"`         // 苹果商店
	AppType     int8      `json:"app_type" gorm:"column:app_type"`         // 设备类型（20：安卓，30：ios）
	DataType    int8      `json:"data_type" gorm:"column:data_type"`       // 数据类型（10：新增用户，20：活跃用户，30：累计用户）
	Date        time.Time `json:"date" gorm:"column:date"`                 // 日期
}

func (m *DataTopChannelInfo) TableName() string {
	return "liveearth_primary.data_top_channel_info"
}
