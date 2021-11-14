/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

type LiveRoomShareContent struct {
	Id          int `json:"id"`
	LiveRoomId  int `json:"live_room_id"`
	ContentType int `json:"content_type"` // 内容类型; 10=直播; 20=视频
	ContentId   int `json:"content_id"`
	SortValue   int `json:"sort_value"`
}

func (*LiveRoomShareContent) TableName() string {
	return "liveearth_primary.live_room_share_content"
}
