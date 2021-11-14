package liveearth_primary

import "time"

type StoryVideoInfo struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	StoryId  int    `json:"story_id" gorm:"column:story_id"`   // 故事主键id
	VideoUrl string `json:"video_url" gorm:"column:video_url"` // 视频链接
	CoverUrl string `json:"cover_url" gorm:"column:cover_url"` // 封面图片/视频
	//VideoPoint unknown   `json:"video_point" gorm:"column:video_point"` // 定位点
	Province   string    `json:"province" gorm:"column:province"`       // 省份
	City       string    `json:"city" gorm:"column:city"`               // 城市
	ShowIndex  int8      `json:"show_index" gorm:"column:show_index"`   // 页面展示顺序
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
}

func (m *StoryVideoInfo) TableName() string {
	return "liveearth_primary.story_video_info"
}
