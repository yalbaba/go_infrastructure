package liveearth_primary

import "time"

type LiveTheme struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"` // 主题id
	ThemeName  string    `json:"theme_name" gorm:"column:theme_name"`    // 主题名称
	IconUrl    string    `json:"icon_url" gorm:"column:icon_url"`        // icon地址
	State      int8      `json:"state" gorm:"column:state"`              // 上架状态  10 下架  20 上架
	Sort       int       `json:"sort" gorm:"column:sort"`                // 主题排序
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`  // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`  // 更新时间
	Deleted    int8      `json:"deleted" gorm:"column:deleted"`          // 10 否 20 是
	Operator   string    `json:"operator" gorm:"column:operator"`        // 操作人
}

func (m *LiveTheme) TableName() string {
	return "liveearth_primary.live_theme"
}
