package liveearth_primary

import "time"

type EuropeanCupLog struct {
	Id         int       `json:"id" xorm:"id"`
	User       string    `json:"user" gorm:"column:user"`               // 点赞的用户
	MatchId    int       `json:"match_id" gorm:"column:match_id"`       // 比赛id
	TeamName   string    `json:"team_name" gorm:"column:team_name"`     // 点赞的球队
	CreateDate time.Time `json:"create_date" gorm:"column:create_date"` // 点赞日期
}

func (m *EuropeanCupLog) TableName() string {
	return "liveearth_primary.european_cup_log"
}
