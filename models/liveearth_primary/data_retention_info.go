package liveearth_primary

import "time"

type DataRetentionInfo struct {
	Id               int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	TotalInstallUser int       `json:"total_install_user" gorm:"column:total_install_user"` // 当天安装数
	RetentionRate    int       `json:"retention_rate" gorm:"column:retention_rate"`         // 留存率
	AppType          int8      `json:"app_type" gorm:"column:app_type"`                     // app类型：20：安卓，30：ios
	Date             time.Time `json:"date" gorm:"column:date"`                             // 日期
}

func (m *DataRetentionInfo) TableName() string {
	return "liveearth_primary.data_retention_info"
}
