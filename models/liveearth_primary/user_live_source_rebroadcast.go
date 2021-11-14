package liveearth_primary

import "time"

// UserLiveSourceRebroadcast ...
type UserLiveSourceRebroadcast struct {
	Id int `json:"id" xorm:"not null pk autoincr comment('pk') INT(10)"`
	// UserID 用户id
	UserId int `json:"user_id"`
	// LiveSourceID 直播源id
	LiveSourceId int `json:"live_source_id"`
	// RebroadcastName 转播渠道名称
	RebroadcastName string `json:"rebroadcast_name"`
	// 水印Id
	WatermarkId int `json:"watermark_id"`
	// 流名称
	StreamName string `json:"stream_name"`
	// RebroadcastEndTime 转播结束时间
	RebroadcastEndTime uint64 `json:"rebroadcast_end_time"`
	// RebroadcastIsPermanent 有效期是否永久;10=否;20=是
	RebroadcastIsPermanent int       `json:"rebroadcast_is_permanent"`
	CreateTime             time.Time `json:"create_time"`
	// Deleted 10=否;20=是
	Deleted    int       `json:"deleted"`
	DeleteTime time.Time `json:"delete_time"`
}

func (*UserLiveSourceRebroadcast) TableName() string {
	return "liveearth_primary.user_live_source_rebroadcast"
}
