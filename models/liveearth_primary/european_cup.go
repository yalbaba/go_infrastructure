package liveearth_primary

import "time"

type EuropeanCup struct {
	Id        int    `json:"id" xorm:"id"`                          // id
	MatchId   int    `json:"match_id" gorm:"column:match_id"`       // 比赛id
	TeamAName string `json:"team_a_name" gorm:"column:team_a_name"` // 队伍a名称

	TeamBName      string `json:"team_b_name" gorm:"column:team_b_name"`             // 队伍b名称
	TeamALogoBlack string `json:"team_a_logo_black" gorm:"column:team_a_logo_black"` // 队伍A黑封面
	TeamBLogoBlack string `json:"team_b_logo_black" gorm:"column:team_b_logo_black"` // 队伍b黑封面

	TeamAScore     string    `json:"team_a_score" gorm:"column:team_a_score"`           // 队伍a得分
	TeamBScore     string    `json:"team_b_score" gorm:"column:team_b_score"`           // 队伍b得分
	StartPlay      time.Time `json:"start_play" gorm:"column:start_play"`               // 比赛开始时间
	TeamALike      int       `json:"team_a_like" gorm:"column:team_a_like"`             // 队伍a点赞数
	TeamBLike      int       `json:"team_b_like" gorm:"column:team_b_like"`             // 队伍b点赞数
	TeamALogoWhite string    `json:"team_a_logo_white" gorm:"column:team_a_logo_white"` // 队伍a彩色logo
	TeamBLogoWhite string    `json:"team_b_logo_white" gorm:"column:team_b_logo_white"` // 队伍b彩色logo
	MatchType      string    `json:"match_type" gorm:"column:match_type"`               // 比赛类型

	GroundName         string `json:"ground_name" gorm:"column:ground_name"`                 // 球场名称
	GroundIntroduction string `json:"ground_introduction" gorm:"column:ground_introduction"` // 球场介绍
	GroundLogo         string `json:"ground_logo" gorm:"column:ground_logo"`                 // 球场图片
	GroundLogoOutside  string `json:"ground_logo_outside" gorm:"column:ground_logo_outside"` // 球场背景小图
}

func (m *EuropeanCup) TableName() string {
	return "liveearth_primary.european_cup"
}
