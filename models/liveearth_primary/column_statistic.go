package liveearth_primary

type ColumnStatistics struct {
	Id         int32  `json:"id"`
	ColumnId   int32  `json:"column_id"`
	ViewNumber int32  `json:"view_number"`
	UserNumber int32  `json:"user_number"`
	CreateTime string `json:"create_time"`
	PvDay      int    `json:"pv_day"`
	UvDay      int    `json:"uv_day"`
}

func (*ColumnStatistics) TableName() string {
	return "liveearth_primary.column_statistics"
}
