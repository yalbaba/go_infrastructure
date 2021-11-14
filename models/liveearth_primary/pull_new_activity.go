package liveearth_primary

import "time"

// PullNewActivity ...
type PullNewActivity struct {
	Id int64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	// ActivityName 活动名称
	ActivityName string `json:"activity_name"`
	// CoverURL 活动图片
	CoverUrl string `json:"cover_url"`
	// WinnerNum 中奖人数
	WinnerNum int `json:"winner_num"`
	// HelperNum 需要助力人数
	HelperNum int `json:"helper_num"`
	// HelperType 助力对象: 10 = 全部用户; 20 = 新注册用户
	HelperType int `json:"helper_type"`
	// ShareTitle 分享的标题
	ShareTitle string `json:"share_title"`
	// ShareImageURL 分享的图片
	ShareImageUrl string `json:"share_image_url"`
	// AppCommunityQrcode app用户社群二维码
	AppCommunityQrcode string `json:"app_community_qrcode"`
	// MiniappCommunityQrcode 小程序用户,社群二维码
	MiniappCommunityQrcode string `json:"miniapp_community_qrcode"`
	// LotteryType 抽奖方式: 10 = 系统随机; 20 = 指定用户
	LotteryType int `json:"lottery_type"`
	// LotteryAssignUserID 指定中奖的用户ID的数组
	LotteryAssignUserId string `json:"lottery_assign_user_id"`
	// State 状态: 10 = 未开始; 20 = 进行中; 30 = 已结束
	State int `json:"state"`
	// DrawTime 开奖时间
	DrawTime time.Time `json:"draw_time"`
	// StartTime 活动开始时间
	StartTime time.Time `json:"start_time"`
	// EndTime 活动结束时间
	EndTime time.Time `json:"end_time"`
	// Operator 操作人
	Operator string `json:"operator"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`
}

func (*PullNewActivity) TableName() string {
	return "liveearth_primary.pull_new_activity"
}
