package liveearth_primary

import "time"

type IconConfig struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// IconUrl icon图片
	IconUrl string `json:"icon_url"`
	// ActiveIconUrl 活动icon
	ActiveIconUrl string `json:"active_icon_url"`
	// IconName icon名称
	IconName string `json:"icon_name"`
	// ActiveIconName 活动icon名称
	ActiveIconName string `json:"active_icon_name"`
	// Sort 排序(数字越小,越靠前)
	Sort int `json:"sort"`
	// IconType icon跳转类型: 1 直播, 2足迹,4发现,8 视频,16 卫星TV, 32 发现
	IconType int `json:"icon_type"`
	// State 状态: 10 = 不显示 ; 20 = 显示
	State int `json:"state"`
	// ActiveStartTime 活动开始时间
	ActiveStartTime time.Time `json:"active_start_time"`
	// ActiveEndTime 活动结束时间
	ActiveEndTime time.Time `json:"active_end_time"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`
	Operator   string    `json:"operator"`
}

func (*IconConfig) TableName() string {
	return "liveearth_primary.icon_config"
}
