package liveearth_primary

import "time"

type StoryImageInfo struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	StoryId  int    `json:"story_id" gorm:"column:story_id"`   // 故事主键id
	ImageUrl string `json:"image_url" gorm:"column:image_url"` // 图片url
	Width    int    `json:"width" gorm:"column:width"`         // 宽
	Height   int    `json:"height" gorm:"column:height"`       // 高
	//ImagePoint unknown   `json:"image_point" gorm:"column:image_point"` // 定位点
	Province   string    `json:"province" gorm:"column:province"`       // 省份
	City       string    `json:"city" gorm:"column:city"`               // 城市
	ShowIndex  int8      `json:"show_index" gorm:"column:show_index"`   // 页面展示顺序
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
}

func (m *StoryImageInfo) TableName() string {
	return "liveearth_primary.story_image_info"
}
