package liveearth_primary

type LiveSourceScore struct {
	Id           int64  `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveSourceId int64  `json:"live_source_id" `
	ScoreTime    string `json:"score_time"`
	ScoreName    string `json:"score_name" `
	Score        int    `json:"score"`
}

func (*LiveSourceScore) TableName() string {
	return "liveearth_primary.live_source_score"
}
