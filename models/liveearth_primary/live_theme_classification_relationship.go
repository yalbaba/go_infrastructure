package liveearth_primary

import "time"

type LiveThemeClassificationRelationship struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(10)"`                      // 数据id
	LiveThemeId          int       `json:"live_theme_id" gorm:"column:live_theme_id"`                   // 直播主题id
	LiveClassificationId int       `json:"live_classification_id" gorm:"column:live_classification_id"` // 直播分类id
	Sort                 int8      `json:"sort"`
	CreateTime           time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime           time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *LiveThemeClassificationRelationship) TableName() string {
	return "liveearth_primary.live_theme_classification_relationship"
}
