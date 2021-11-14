package liveearth_primary

import "time"

type StoryBaseInfo struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	StoryTitle string `json:"story_title" gorm:"column:story_title"` // 故事标题
	CoverUrl   string `json:"cover_url" gorm:"column:cover_url"`     // 封面图片
	MapUrl     string `json:"map_url" gorm:"column:map_url"`         // 遥感图片
	//StoryPoint   unknown   `json:"story_point" gorm:"column:story_point"`     // 定位点
	LiveRoomId   int64     `json:"live_room_id" gorm:"column:live_room_id"`   // 直播间id
	TopicId      int       `json:"topic_id" gorm:"column:topic_id"`           // 话题id
	MusicId      int       `json:"music_id" gorm:"column:music_id"`           // 背景音乐id
	State        int8      `json:"state" gorm:"column:state"`                 // 审核状态
	UserId       string    `json:"user_id" gorm:"column:user_id"`             // 作者用户id
	Province     string    `json:"province" gorm:"column:province"`           // 省份
	City         string    `json:"city" gorm:"column:city"`                   // 城市
	PublishTime  time.Time `json:"publish_time" gorm:"column:publish_time"`   // 发布时间
	Keywords     string    `json:"keywords" gorm:"column:keywords"`           // 关键词
	Operator     string    `json:"operator" gorm:"column:operator"`           // 操作人
	VerifyReason string    `json:"verify_reason" gorm:"column:verify_reason"` // 审核原因
	StoryLevel   int8      `json:"story_level" gorm:"column:story_level"`     // 内容等级
	BeLikeNum    int       `json:"be_like_num" gorm:"column:be_like_num"`     // 真实点赞数
	ViewNum      int64     `json:"view_num" gorm:"column:view_num"`           // 真实观看量
	BaseViewNum  int64     `json:"base_view_num" gorm:"column:base_view_num"` // 基础观看量
	AuditTime    time.Time `json:"audit_time"`                                // 审核时间
	BaseLikeNum  int64     `json:"base_like_num"`                             // 基础点赞数量

}

func (m *StoryBaseInfo) TableName() string {
	return "liveearth_primary.story_base_info"
}
