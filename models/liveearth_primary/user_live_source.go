package liveearth_primary

import "time"

// UserLiveSource ...
type UserLiveSource struct {
	Id int `json:"id" xorm:"not null pk autoincr comment('pk') INT(10)"`
	// UserID 购买资源用户ID
	UserId int `json:"user_id"`
	// LiveSourceID 直播源id
	LiveSourceId int `json:"live_source_id"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// Deleted 10=否; 20=是
	Deleted     int       `json:"deleted"`
	DeletedTime time.Time `json:"deleted_time"`
}

func (*UserLiveSource) TableName() string {
	return "liveearth_primary.user_live_source"
}
