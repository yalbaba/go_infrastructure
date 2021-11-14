package liveearth_primary

import (
	"time"
)

type AppConfig struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	DeviceType   int       `json:"device_type" xorm:"not null default 10 comment('设备类型: 10=所有; 20=安卓; 30=ios') TINYINT(3)"`
	Ckey         string    `json:"ckey" xorm:"not null comment('字段名') unique VARCHAR(64)"`
	Cvalue       string    `json:"cvalue" xorm:"not null default '' comment('字段值, 最大长度12k') VARCHAR(12288)"`
	Payload      string    `json:"payload" xorm:"not null default '{}' comment('负载数据, 它是一个json') JSON"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Operator     int       `json:"operator" xorm:"not null default 0 comment('操作人id') INT(10)"`
	OperatorName string    `json:"operator_name" xorm:"not null default '' comment('操作人名字') VARCHAR(32)"`
	LayoutType   int8      `json:"layout_type"`
	State        int       `json:"state" xorm:"not null comment('状态: 10=关闭; 20=开启') TINYINT(3)"`
}

func (*AppConfig) TableName() string {
	return "liveearth_primary.app_config"
}
