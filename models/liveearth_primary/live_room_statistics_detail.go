package liveearth_primary

// LiveRoomStatisticsDetail ...
type LiveRoomStatisticsDetail struct {
	Id uint64 `json:"id"`
	// LiveRoomID 直播间id
	LiveRoomId int64 `json:"live_room_id"`
	// Exposure 曝光量
	Exposure int32 `json:"exposure"`
	// ClickConversionRate 点击转化率
	ClickConversionRate int32 `json:"click_conversion_rate"`
	LiveRoomUv          int32 `json:"live_room_uv"`
	LiveRoomPv          int32 `json:"live_room_pv"`
	// CommentNum 评论数
	CommentNum int32 `json:"comment_num"`
	// LikeNum 点赞数
	LikeNum int32 `json:"like_num"`
	// CollectNum 收藏数
	CollectNum int32 `json:"collect_num"`
	// AppShareNum 分享次数
	AppShareNum int32 `json:"app_share_num"`
	// ShareUv 分享链接浏览uv
	ShareUv int32 `json:"share_uv"`
	// SharePv 分享链接浏览pv
	SharePv int32 `json:"share_pv"`
	// ShareForwardNum 转发次数
	ShareForwardNum int32 `json:"share_forward_num"`
	// NavigateClickNum 导航点击次数
	NavigateClickNum int32 `json:"navigate_click_num"`
	// FullScreenNum 全屏观看次数
	FullScreenNum int32 `json:"full_screen_num"`
	// ShareBarrageNum 分享页面弹幕数
	ShareBarrageNum int32 `json:"share_barrage_num"`
	// ShareLikeNum 分享页面点赞数
	ShareLikeNum int32 `json:"share_like_num"`
	// ShareDownloadNum 分享下载触发次数
	ShareDownloadNum int32 `json:"share_download_num"`
	// ShareLoginNum 分享页登录成功次数
	ShareLoginNum int32 `json:"share_login_num"`
	// CreateAt 日期
	CreateAt string `json:"create_at"`
}

func (*LiveRoomStatisticsDetail) TableName() string {
	return "liveearth_primary.live_room_statistics_detail"
}
