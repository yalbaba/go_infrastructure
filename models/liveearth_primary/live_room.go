package liveearth_primary

import (
	"time"
)

type LiveRoom struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	RoomName           string    `json:"room_name" xorm:"not null comment('直播间名称') index VARCHAR(32)"`
	RoomDesc           string    `json:"room_desc" xorm:"comment('描述') VARCHAR(255)"`
	RecommendScore     int       `json:"recommend_score" xorm:"not null default 0 comment('推荐分值') INT(11)"`
	CoverUrl           string    `json:"cover_url" xorm:"not null comment('直播间封面图片') VARCHAR(255)"`
	LiveSourceId       int       `json:"live_source_id" xorm:"not null comment('直播源id') index INT(10)"`
	LiveSourceName     string    `json:"live_source_name" xorm:"not null comment('直播源名称') VARCHAR(64)"`
	UserId             string    `json:"user_id" xorm:"not null comment('直播用户id') VARCHAR(32)"`
	UserName           string    `json:"user_name" xorm:"comment('直播用户名') VARCHAR(64)"`
	AnchorId           string    `json:"anchor_id" xorm:"comment('主持人id') VARCHAR(32)"`
	AnchorName         string    `json:"anchor_name" xorm:"comment('主持人名称') VARCHAR(32)"`
	Nation             string    `json:"nation" xorm:"not null default '' comment('国家') VARCHAR(32)"`
	Province           string    `json:"province" xorm:"not null default '' comment('省份') VARCHAR(32)"`
	City               string    `json:"city" xorm:"not null default '' comment('城市') VARCHAR(32)"`
	Address            string    `json:"address" xorm:"not null default '' comment('地点名称') VARCHAR(32)"`
	Lon                float64   `json:"lon" xorm:"not null default 0 comment('经度') DOUBLE"`
	Lat                float64   `json:"lat" xorm:"not null default 0 comment('纬度') DOUBLE"`
	Altitude           int       `json:"altitude" xorm:"not null default 0 comment('海拔') INT(11)"`
	Label              string    `json:"label" xorm:"default '[]' comment('标签') VARCHAR(64)"`
	LikeNum            int64     `json:"like_num" xorm:"not null default 0 comment('点赞数量') BIGINT(20)"`
	LikeRobotNum       int       `json:"like_robot_num" xorm:"not null default 0 comment('机器人点赞数量') INT(11)"`
	OnlineNum          int       `json:"online_num" xorm:"not null default 0 comment('不存数据,结构中占位用') INT(11)"`
	LookNum            int       `json:"look_num" xorm:"not null default 0 comment('累计观看人数(不去重)') INT(11)"`
	LookRobotNum       int       `json:"look_robot_num" xorm:"not null default 0 comment('机器人观看数量(累加)') INT(11)"`
	SpeakNum           int       `json:"speak_num" xorm:"not null default 0 comment('发言人数(去重)') INT(11)"`
	SpeakNumTotal      int       `json:"speak_num_total" xorm:"not null default 0 comment('累计发言人数(不去重)') INT(11)"`
	CommentNum         int       `json:"comment_num" xorm:"not null default 0 comment('精彩瞬间评论的人数(去重)') INT(11)"`
	CommentNumTotal    int       `json:"comment_num_total" xorm:"not null default 0 comment('精彩瞬间评论的人次(不去重)') INT(11)"`
	State              int       `json:"state" xorm:"not null default 10 comment('状态 10 = 关闭  20 = 开启') TINYINT(4)"`
	Operator           string    `json:"operator" xorm:"not null default '' comment('操作人') VARCHAR(32)"`
	CreateTime         time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime         time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Deleted            int       `json:"deleted" xorm:"not null default 10 comment('删除标记 10 = 未删除; 20 = 删除') TINYINT(4)"`
	CreateState        int       `json:"create_state" xorm:"not null default 10 comment('创建直播间状态: 10 =  正在创建中; 20 = 创建成功') TINYINT(4)"`
	DeleteTime         time.Time `json:"delete_time" xorm:"TIMESTAMP"`
	AreaId             int       `json:"area_id" xorm:"comment('地区id，用于映射到天气表的主键') INT(11)"`
	OpenStartTime      time.Time `json:"open_start_time"`
	OpenEndTime        time.Time `json:"open_end_time"`
	MatchUser          string    `json:"match_user"`
	MatchDistrict      string    `json:"match_district"`
	IsSatelliteTv      int       `json:"is_satellite_tv"`
	WonderfulVideoId   int64     `json:"wonderful_video_id"`
	WonderfulVideoName string    `json:"wonderful_video_name"`
	RoomType           int8      `json:"room_type"`
	NoticeType         int8      `json:"notice_type"`
	LiveStartTime      time.Time `json:"live_start_time"`
	NoticeVideoId      int       `json:"notice_video_id"`
	NoticeImageUrl     string    `json:"notice_image_url"`
}

func (*LiveRoom) TableName() string {
	return "liveearth_primary.live_room"
}
