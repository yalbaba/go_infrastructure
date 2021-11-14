package liveearth_primary

import "time"

// ImBarrage ...
type ImBarrage struct {
	Id uint64 `json:"id" xorm:"not null pk autoincr INT(20)"`
	// LiveRoomId 直播间id
	LiveRoomId int64 `json:"live_room_id"`
	// LiveRoomName 直播间名称
	LiveRoomName string `json:"live_room_name"`
	// LiveRoomUserName 直播用户名
	LiveRoomUserName string `json:"live_room_user_name"`
	CommentUserId    string `json:"comment_user_id"`
	// CommentUserName 评论用户名
	CommentUserName string `json:"comment_user_name"`
	// Content 内容(回调原始数据)
	Content string `json:"content"`
	// CommentStr 弹幕内容(后台展示)
	CommentStr string `json:"comment_str"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
}

func (*ImBarrage) TableName() string {
	return "liveearth_primary.im_barrage"
}
