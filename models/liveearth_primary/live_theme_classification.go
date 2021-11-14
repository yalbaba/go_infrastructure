package liveearth_primary

import "time"

type LiveThemeClassification struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(10)"`                // 分类id
	ClassificationName string    `json:"classification_name" gorm:"column:classification_name"` // 分类名称
	State              int8      `json:"state"`
	Deleted            int8      `json:"deleted" gorm:"column:deleted"`         // 是否删除 10 否 20 是
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
	Operator           string    `json:"operator" gorm:"column:operator"`       // 操作人
}

func (m *LiveThemeClassification) TableName() string {
	return "liveearth_primary.live_theme_classification"
}
