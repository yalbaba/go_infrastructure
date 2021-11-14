package liveearth_primary

type LiveSharePosterBackground struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	MediaUrl  string `json:"media_url"`
	SortValue int    `json:"sort_value"`
	Deleted   int    `json:"deleted"`
}

func (*LiveSharePosterBackground) TableName() string {
	return "liveearth_primary.live_share_poster_background"
}
