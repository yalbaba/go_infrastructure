package liveearth_primary

import (
	"time"
)

type Activity struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ConfigType       int       `json:"config_type" xorm:"not null comment('活动类型: 10=开屏广告; 20=活动弹窗; 30=固定活动; 40=每日一图; 50=文章内广告; 60=消费电子地图') TINYINT(3)"`
	CreateTime       time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime       time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Operator         int       `json:"operator" xorm:"not null default 0 comment('操作人id') INT(10)"`
	OperatorName     string    `json:"operator_name" xorm:"not null default '' comment('操作人名字') VARCHAR(32)"`
	Deleted          int       `json:"deleted" xorm:"not null default 10 comment('删除标记 10 = 未删除; 20 = 删除') TINYINT(3)"`
	DeleteTime       time.Time `json:"delete_time" xorm:"comment('删除时间') TIMESTAMP"`
	StartTime        time.Time `json:"start_time" xorm:"comment('开始时间') TIMESTAMP"`
	EndTime          time.Time `json:"end_time" xorm:"comment('结束时间') TIMESTAMP"`
	Title            string    `json:"title" xorm:"not null comment('标题') VARCHAR(64)"`
	DeviceType       int8      `json:"device_type" xorm:"not null default 10 comment('设备类型: 10=所有; 20=安卓; 30=ios') index TINYINT(3)"`
	SkipType         int       `json:"skip_type" xorm:"not null default 10 comment('跳转类型: 10=无 20=h5链接; 30=文章; 40=跳转到经纬度; 50=消费电子地图') TINYINT(3)"`
	State            int       `json:"state" xorm:"not null default 10 comment('开启状态: 10=关闭; 20=开启') TINYINT(3)"`
	Payload          string    `json:"payload" xorm:"not null default '{}' comment('负载数据, 它是一个json') JSON"`
	MatchUserIds     string    `json:"match_user_ids"`
	MatchDistrictIds string    `json:"match_district_ids"`
	BannerPosition   int       `json:"banner_position"`
}

func (*Activity) TableName() string {
	return "liveearth_primary.activity"
}
