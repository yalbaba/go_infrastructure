package liveearth_primary

import "time"

type StoryContentInfo struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	StoryId    int       `json:"story_id" gorm:"column:story_id"`       // 故事主键id
	StoryText  string    `json:"story_text" gorm:"column:story_text"`   // 故事正文
	ShowIndex  int8      `json:"show_index" gorm:"column:show_index"`   // 页面展示顺序
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
}

func (m *StoryContentInfo) TableName() string {
	return "liveearth_primary.story_content_info"
}
