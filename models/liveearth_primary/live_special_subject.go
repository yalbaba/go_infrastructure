package liveearth_primary

import (
	"time"
)

type LiveSpecialSubject struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Title         string `json:"title" xorm:"not null comment('专题名称') VARCHAR(16)"`
	Introduction  string `json:"introduction"`
	BasePeopleNum int    `json:"base_people_num"`
	CoverUrl      string `json:"cover_url" xorm:"not null comment('专题封面') VARCHAR(1024)"`
	// LiveRoomNum 直播间数量
	LiveRoomNum    int       `json:"live_room_num"`
	RecommendLevel int       `json:"recommend_level" xorm:"not null default 0 comment('推荐等级,按此字段非0排序') TINYINT(4)"`
	CreateTime     time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	State          int       `json:"state" xorm:"state"` //状态
}

func (*LiveSpecialSubject) TableName() string {
	return "liveearth_primary.live_special_subject"
}
