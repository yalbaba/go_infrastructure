package liveearth_primary

import "time"

type DataDayInfo struct {
	Id                  int     `json:"id" xorm:"not null pk autoincr INT(10)"`
	NewUser             int     `json:"new_user" gorm:"column:new_user"`                             // 新增用户7日平均
	NewUserGrew         float32 `json:"new_user_grew" gorm:"column:new_user_grew"`                   // 新增用户7日平均同比
	ActiveUser          int     `json:"active_user" gorm:"column:active_user"`                       // 活跃用户7日平均
	ActiveUserGrew      float32 `json:"active_user_grew" gorm:"column:active_user_grew"`             // 活跃用户7日平均同比
	TotalActiveUser     int     `json:"total_active_user" gorm:"column:total_active_user"`           // 近7日总活跃
	TotalActiveUserGrew float32 `json:"total_active_user_grew" gorm:"column:total_active_user_grew"` // 近7日总活跃同比
	TotalUser           int     `json:"total_user" gorm:"column:total_user"`                         // 累计用户
	RetentionRate       float32 `json:"retention_rate" gorm:"column:retention_rate"`                 // 新用户次日留存率（7日平均）
	RetentionRateGrew   float32 `json:"retention_rate_grew" gorm:"column:retention_rate_grew"`       // 新用户次日留存率（7日平均）同比
	Frequency           float32 `json:"frequency" gorm:"column:frequency"`                           // 使用频率（7日平均）（单位：秒）
	FrequencyGrew       float32 `json:"frequency_grew" gorm:"column:frequency_grew"`                 // 使用频率（7日平均）同比
	Duration            int     `json:"duration" gorm:"column:duration"`                             // 使用时长（7日平均）
	DurationGrew        float32 `json:"duration_grew" gorm:"column:duration_grew"`                   // 使用时长（7日平均）同比
	Pv                  float32 `json:"pv" gorm:"column:pv"`                                         // 访问页面数（7日平均）
	PvGrew              float32 `json:"pv_grew" gorm:"column:pv_grew"`                               // 访问页面数（7日平均）同比
	LiveUser            int     `json:"live_user" gorm:"column:live_user"`                           // 在线互动用户（7日平均）
	LiveUserGrew        float32 `json:"live_user_grew" gorm:"column:live_user_grew"`                 // 在线互动用户（7日平均）同比
	TotalLiveUser       int     `json:"total_live_user" gorm:"column:total_live_user"`               // 累计在线互动用户
	WatchUser           int     `json:"watch_user" gorm:"column:watch_user"`                         // 观看直播用户（7日平均）
	WatchUserGrew       float32 `json:"watch_user_grew" gorm:"column:watch_user_grew"`               // 观看直播用户（7日平均）同比
	TotalWatchUser      int     `json:"total_watch_user" gorm:"column:total_watch_user"`             // 累计观看直播人数
	AppType             int8    `json:"app_type" gorm:"column:app_type"`                             // app类型：20安卓，30ios  10全部

	Date time.Time `json:"date" gorm:"column:date"` // 日期
}

func (m *DataDayInfo) TableName() string {
	return "liveearth_primary.data_day_info"
}
