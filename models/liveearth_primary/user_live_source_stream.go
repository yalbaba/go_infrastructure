/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/11/20
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type UserLiveSourceStream struct {
	Id           int       `json:"id"`
	LiveSourceId int       `json:"live_source_id"`
	WatermarkId  int       `json:"watermark_id"` // WatermarkID 水印id, 这里表示数据库的id而不是腾讯云平台上的水印id
	StreamName   string    `json:"stream_name"`  // StreamName 新的流名称
	CreateTime   time.Time `json:"create_time"`
}

func (*UserLiveSourceStream) TableName() string {
	return "liveearth_primary.user_live_source_stream"
}
