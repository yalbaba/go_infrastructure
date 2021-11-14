package liveearth_primary

import (
	"time"
)

type LiveRoomOpenTime struct {
	Id                  int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId          int       `json:"live_room_id" xorm:"not null comment('直播间id') index INT(10)"`
	OpenType            int       `json:"open_type" xorm:"not null comment('开启时段类型 10= 每天; 20=每周; 30=不限; 40 = 全天候开启;') TINYINT(4)"`
	WeekNum             int       `json:"week_num" xorm:"not null default 0 comment('星期 1,2,3,4,5,6,7') TINYINT(4)"`
	StartTime           string    `json:"start_time" xorm:"comment('open_type = 30 不使用') TIME"`
	EndTime             string    `json:"end_time" xorm:"comment('open_type = 30 不使用') TIME"`
	StartAt             time.Time `json:"start_at" xorm:"comment('open_type = 30 使用') TIMESTAMP"`
	EndAt               time.Time `json:"end_at" xorm:"comment('open_type = 30 使用') TIMESTAMP"`
	LiveAlldayStartTime string    `json:"live_allday_start_time"`
	LiveAlldayEndTime   string    `json:"live_allday_end_time"`
}

func (*LiveRoomOpenTime) TableName() string {
	return "liveearth_primary.live_room_open_time"
}
