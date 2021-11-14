/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/7
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type LivePlayback struct {
	Id            int       `json:"id"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	OperatorId    int       `json:"operator_id"`
	OperatorName  string    `json:"operator_name"`
	Deleted       int       `json:"deleted"` // Deleted 删除状态; 10=未删除; 20=删除
	DeleteTime    time.Time `json:"delete_time"`
	LiveSourceId  int       `json:"live_source_id"`
	SortValue     int       `json:"sort_value"` // 排序值
	Title         string    `json:"title"`      // Title 标题; 同一直播源的回放标题不能重复
	CoverUrl      string    `json:"cover_url"`  // CoverURL 封面图url
	LogoUrl       string    `json:"logo_url"`
	MediaUrl      string    `json:"media_url"`       // MediaURL 媒体url
	MediaPlayTime int       `json:"media_play_time"` // 媒体播放时间(秒)
	State         int       `json:"state"`           // State 开启状态; 10=关闭; 20=开启
	StartTime     int       `json:"start_time"`
	EndTime       int       `json:"end_time"`
	UnlimitedTime int       `json:"unlimited_time"` // UnlimitedTime 是否不限时间; 10=否; 20=是
}

func (*LivePlayback) TableName() string {
	return "liveearth_primary.live_playback"
}
