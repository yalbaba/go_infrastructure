package liveearth_primary

import (
	"time"
)

type Encyclopedia struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	DataId             string    `json:"data_id" xorm:"VARCHAR(32)"`
	Title              string    `json:"title" xorm:"not null comment('标题') VARCHAR(32)"`
	Summary            string    `json:"summary" xorm:"not null comment('摘要') MEDIUMTEXT"`
	CoverUrl           string    `json:"cover_url" xorm:"not null comment('封面图') VARCHAR(1024)"`
	CoverMediaType     int       `json:"cover_media_type"`
	Content            string    `json:"content" xorm:"comment('内容') TEXT"`
	AuthorName         string    `json:"author_name" xorm:"comment('作者名称') VARCHAR(32)"`
	AuthorId           string    `json:"author_id"`
	ShowAuthor         int       `json:"show_author"`
	PositionType       int8      `json:"position_type"` // 点位类型; 10=区域; 20=点位
	Continent          string    `json:"continent"`     // 州
	Nation             string    `json:"nation" xorm:"VARCHAR(32)"`
	Province           string    `json:"province" xorm:"VARCHAR(32)"`
	City               string    `json:"city" xorm:"VARCHAR(32)"`
	County             string    `json:"county" xorm:"VARCHAR(32)"`
	Town               string    `json:"town" xorm:"VARCHAR(32)"`
	Address            string    `json:"address" xorm:"comment('地址') VARCHAR(255)"`
	Area               string    `json:"area" xorm:"comment('地区') VARCHAR(255)"`
	FlyToHeight        int       `json:"fly_to_height" xorm:"INT(11)"`
	ZoneId             string    `json:"zone_id" xorm:"VARCHAR(32)"`
	SpiderUrl          string    `json:"spider_url" xorm:"comment('爬取url') VARCHAR(1024)"`
	SpiderSourceName   int       `json:"spider_source_name" xorm:"VARCHAR(64)"`
	RealSource         string    `json:"real_source" xorm:"comment('对外展示的来源') VARCHAR(32)"`
	ArticleType        int       `json:"article_type" xorm:"TINYINT(4)"`
	DevelopDatas       string    `json:"develop_datas" xorm:"JSON"`
	CustomDatas        string    `json:"custom_datas" xorm:"JSON"`
	Lon                float64   `json:"lon" xorm:"comment('经度') DOUBLE"`
	Lat                float64   `json:"lat" xorm:"comment('纬度') DOUBLE"`
	FirstPublishTime   time.Time `json:"first_publish_time" xorm:"TIMESTAMP"`
	PublishTime        time.Time `json:"publish_time" xorm:"TIMESTAMP"`
	Keywords           string    `json:"keywords"`
	VisibleStartTime   time.Time `json:"visible_start_time"`
	VisibleEndTime     time.Time `json:"visible_end_time"`
	MatchUser          string    `json:"match_user"`
	MatchDistrict      string    `json:"match_district"`
	Mark               string    `json:"mark"`
	Editor             string    `json:"editor"`
	AuthorPhone        int       `json:"author_phone"`
	RecommendLevel     int       `json:"recommend_level"`
	RecommendLevelTemp int       `json:"recommend_level_temp"`
	RecommendStartTime time.Time `json:"recommend_start_time"`
	RecommendEndTime   time.Time `json:"recommend_end_time"`
	TopLevel           int       `json:"top_level"`
	TopLevelTemp       int       `json:"top_level_temp"`
	TopStartTime       time.Time `json:"top_start_time"`
	TopEndTime         time.Time `json:"top_end_time"`
	SortValue          int       `json:"sort_value"`
	CreateTime         time.Time `json:"create_time"`
	UpdateTime         time.Time `json:"update_time"`
	State              int       `json:"state"`
	ContentType        int       `json:"content_type"`
	AuditReason        string    `json:"audit_reason"`
	BaseViewNum        int       `json:"base_view_num"`
	BaseLikeNum        int       `json:"base_like_num"`
	ListLayout         int       `json:"list_layout"`
	DetailLayout       int       `json:"detail_layout"`
}

func (*Encyclopedia) TableName() string {
	return "liveearth_primary.encyclopedia"
}
