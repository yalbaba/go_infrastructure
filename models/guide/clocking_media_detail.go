package guide

import "time"

type ClockingMediaDetail struct {
	Id            int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"`
	ClockingId    int       `json:"clocking_id" gorm:"column:clocking_id"` // 打卡表关联id
	MediaType     int8      `json:"media_type" gorm:"column:media_type"`   // 类型; 10=图片; 20=视频
	MediaUrl      string    `json:"media_url" gorm:"column:media_url"`
	MediaHeight   int       `json:"media_height" gorm:"column:media_height"`
	MediaWidth    int       `json:"media_width" gorm:"column:media_width"`
	VideoDuration string    `json:"video_duration" gorm:"column:video_duration"` // 视频播放时长
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}

func (m *ClockingMediaDetail) TableName() string {
	return "guide.clocking_media_detail"
}
