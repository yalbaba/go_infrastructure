package liveearth_primary

type DataConfigInfo struct {
	Id                      int  `json:"id" xorm:"not null pk autoincr INT(10)"`
	TotalUserBase           int  `json:"total_user_base" gorm:"column:total_user_base"`                         // 累计用户初始值
	TotalLiveOnlineUserBase int  `json:"total_live_online_user_base" gorm:"column:total_live_online_user_base"` // 累计在线互动用户初始值
	TotalLiveWatchUserBase  int  `json:"total_live_watch_user_base" gorm:"column:total_live_watch_user_base"`   // 累计观看直播用户初始值
	NewUserFactor           int  `json:"new_user_factor" gorm:"column:new_user_factor"`                         // 新增用户量级
	ActiveUserFactor        int  `json:"active_user_factor" gorm:"column:active_user_factor"`                   // 活跃用户量级
	TotalActiveUserFactor   int  `json:"total_active_user_factor" gorm:"column:total_active_user_factor"`       // 近七日总活跃用户
	TotalUserFactor         int  `json:"total_user_factor" gorm:"column:total_user_factor"`                     // 累计用户数（不用）
	RetentionRateFactor     int  `json:"retention_rate_factor" gorm:"column:retention_rate_factor"`             // 新增用户次日留存率
	Duration                int  `json:"duration" gorm:"column:duration"`                                       // 使用时长量级（单位：秒）
	Frequency               int  `json:"frequency" gorm:"column:frequency"`                                     // 使用频率
	Pv                      int  `json:"pv" gorm:"column:pv"`                                                   // 访问页面
	AppType                 int8 `json:"app_type" gorm:"column:app_type"`                                       // 20：安卓，30：ios
	LiveOnlineUserFactor    int  `json:"live_online_user_factor" gorm:"column:live_online_user_factor"`         // 在线互动用户量级
	LiveWatchUserFactor     int  `json:"live_watch_user_factor" gorm:"column:live_watch_user_factor"`           // 观看直播用户量级
}

func (m *DataConfigInfo) TableName() string {
	return "liveearth_primary.data_config_info"
}
