package liveearth_primary

import "time"

// Columns ...
type Columns struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	ColumnCh        string    `json:"column_ch" xorm:"column_ch"`
	ColumnEn        string    `json:"column_en" xorm:"column_en"`
	FirstTagId      int64     `json:"first_tag_id" xorm:"first_tag_id"`
	Sort            int64     `json:"sort" xorm:"sort"`
	State           int32     `json:"state" xorm:"state"`
	MatchUserId     string    `json:"match_user_id"`
	MatchDistrictId string    `json:"match_district_id"`
	CreateTime      time.Time `json:"create_time" xorm:"create_time"`
	UpdateTime      time.Time `json:"update_time" xorm:"update_time"`
	Operator        string    `json:"operator"`
}

func (*Columns) TableName() string {
	return "liveearth_primary.columns"
}
