package liveearth_primary

import "time"

type StoryTag struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	StoryId    int       `json:"story_id" gorm:"column:story_id"`       // 故事id
	TagId      int       `json:"tag_id" gorm:"column:tag_id"`           // 标签id
	TagLevel   int       `json:"tag_level" gorm:"column:tag_level"`     // 标签级别
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
}

func (m *StoryTag) TableName() string {
	return "liveearth_primary.story_tag"
}
