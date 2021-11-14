package activity

import "time"

type TokyoOlympicStatistic struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Pv             int64     `json:"pv" gorm:"column:pv"`                             // 首页pv
	Uv             int64     `json:"uv" gorm:"column:uv"`                             // 首页uv
	AssistancePv   int64     `json:"assistance_pv" gorm:"column:assistance_pv"`       // app内助力按钮pv
	AssistanceUv   int64     `json:"assistance_uv" gorm:"column:assistance_uv"`       // app内助力按钮uv
	H5AssistancePv int64     `json:"h5_assistance_pv" gorm:"column:h5_assistance_pv"` // app外助力按钮pv
	H5AssistanceUv int64     `json:"h5_assistance_uv" gorm:"column:h5_assistance_uv"` // app外助力按钮uv
	H5Pv           int64     `json:"h5_pv" gorm:"column:h5_pv"`                       // h5底部点击pv
	H5Uv           int64     `json:"h5_uv" gorm:"column:h5_uv"`                       // h5底部点击uv
	SharePv        int64     `json:"share_pv" gorm:"column:share_pv"`                 // 分享pv
	ShareUv        int64     `json:"share_uv" gorm:"column:share_uv"`                 // 分享uv
	Date           time.Time `json:"date" gorm:"column:date"`                         // 日期
}

func (m *TokyoOlympicStatistic) TableName() string {
	return "activity.tokyo_olympic_statistic"
}
