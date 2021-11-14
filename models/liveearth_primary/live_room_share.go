/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

type LiveRoomShare struct {
	Id                int    `json:"id"`
	LiveRoomId        int    `json:"live_room_id"`
	AppDownloadUrl    string `json:"app_download_url"`
	ShareTemplateType int    `json:"share_template_type"` // 分享模板; 10=内容; 20=商品
}

func (*LiveRoomShare) TableName() string {
	return "liveearth_primary.live_room_share"
}
