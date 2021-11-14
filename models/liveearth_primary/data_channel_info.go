package liveearth_primary

import "time"

type DataChannelInfo struct {
	Id             int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	Date           time.Time `json:"date" gorm:"column:date"`                         // 日期
	Duration       int       `json:"duration" gorm:"column:duration"`                 // 使用时长（单位：秒）
	ActiveUser     int       `json:"active_user" gorm:"column:active_user"`           // 活跃用户数
	NewUser        int       `json:"new_user" gorm:"column:new_user"`                 // 新增用户数
	TotalUser      int       `json:"total_user" gorm:"column:total_user"`             // 总用户数
	Channel        string    `json:"channel" gorm:"column:channel"`                   // 渠道名称
	ChannelId      string    `json:"channel_id" gorm:"column:channel_id"`             // 渠道id
	Launch         int       `json:"launch" gorm:"column:launch"`                     // 启动次数
	TotalUserRate  float32   `json:"total_user_rate" gorm:"column:total_user_rate"`   // 渠道用户占总用户比例
	AppType        int8      `json:"app_type" gorm:"column:app_type"`                 // app类型：20：安卓，30：ios   10：全部
	LiveOnlineUser int       `json:"live_online_user" gorm:"column:live_online_user"` // 在线互动用户数
	LiveWatchUser  int       `json:"live_watch_user" gorm:"column:live_watch_user"`   // 观看直播用户数
}

func (m *DataChannelInfo) TableName() string {
	return "liveearth_primary.data_channel_info"
}
