package liveearth_primary

import (
	"time"
)

type LiveCyclopedia struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	DataId           string    `json:"data_id" xorm:"VARCHAR(32)"`
	Title            string    `json:"title" xorm:"VARCHAR(255)"`
	Summary          string    `json:"summary" xorm:"MEDIUMTEXT"`
	Cover            string    `json:"cover" xorm:"MEDIUMTEXT"`
	CoverIsVideo     int       `json:"cover_is_video" xorm:"TINYINT(4)"`
	Content          string    `json:"content" xorm:"TEXT"`
	AuthorName       string    `json:"author_name" xorm:"VARCHAR(255)"`
	Editorial        string    `json:"editorial" xorm:"VARCHAR(255)"`
	Address          string    `json:"address" xorm:"VARCHAR(255)"`
	Area             string    `json:"area" xorm:"VARCHAR(255)"`
	Coordinate       string    `json:"coordinate" xorm:"VARCHAR(255)"`
	FlyToHeight      int       `json:"fly_to_height" xorm:"INT(12)"`
	ZoneId           string    `json:"zone_id" xorm:"VARCHAR(255)"`
	Url              string    `json:"url" xorm:"VARCHAR(255)"`
	RealSource       string    `json:"real_source" xorm:"VARCHAR(255)"`
	VideoUrl         string    `json:"video_url" xorm:"VARCHAR(255)"`
	MoreTag          string    `json:"more_tag" xorm:"VARCHAR(255)"`
	TopFlag          int       `json:"top_flag" xorm:"TINYINT(4)"`
	TopicFlag        int       `json:"topic_flag" xorm:"TINYINT(4)"`
	RelevantType     int       `json:"relevant_type" xorm:"INT(6)"`
	ShowType         int       `json:"show_type" xorm:"INT(6)"`
	FlashImgFlag     int       `json:"flash_img_flag" xorm:"TINYINT(4)"`
	HotMode          int       `json:"hot_mode" xorm:"TINYINT(4)"`
	NewFlag          int       `json:"new_flag" xorm:"TINYINT(4)"`
	FirstPublishTime time.Time `json:"first_publish_time" xorm:"TIMESTAMP"`
	RelatedPublicMm  time.Time `json:"related_public_mm" xorm:"TIMESTAMP"`
	ArticleType      int       `json:"article_type" xorm:"INT(4)"`
	DevelopDatas     string    `json:"develop_datas" xorm:"VARCHAR(255)"`
	CustomDatas      string    `json:"custom_datas" xorm:"TEXT"`
	Other            string    `json:"other" xorm:"TEXT"`
	Source           string    `json:"source" xorm:"VARCHAR(255)"`
	PublishTimeText  time.Time `json:"publish_time_text" xorm:"DATETIME"`
}

func (*LiveCyclopedia) TableName() string {
	return "liveearth_primary.live_cyclopedia"
}
