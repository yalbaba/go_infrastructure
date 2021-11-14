package liveearth_primary

import (
	"time"
)

type FootprintMediaDetail struct {
	Id                     int64     `json:"id" xorm:"pk autoincr comment('足迹详情id') BIGINT(20)"`
	FootprintId            int       `json:"footprint_id" xorm:"not null comment('足迹id') index INT(10)"`
	UserId                 string    `json:"user_id" xorm:"not null comment('用户id') index VARCHAR(32)"`
	MediaCreateTime        time.Time `json:"media_create_time" xorm:"not null default CURRENT_TIMESTAMP comment('媒体的创建时间') TIMESTAMP"`
	MediaDesc              string    `json:"media_desc" xorm:"comment('media描述') VARCHAR(64)"`
	SortId                 int       `json:"sort_id" xorm:"not null comment('生成故事内容的列表顺序') TINYINT(4)"`
	AreaId                 int       `json:"area_id" xorm:"not null comment('区域排序id') TINYINT(4)"`
	MediaType              int       `json:"media_type"`
	MediaUrl               string    `json:"media_url" xorm:"not null comment('图片') VARCHAR(1024)"`
	VideoCoverUrl          string    `json:"video_cover_url" xorm:"not null comment('视频封面') VARCHAR(1024)"`
	VideoDuration          string    `json:"video_duration" xorm:"not null comment('视频播放时长') VARCHAR(32)"`
	VideoType              int       `json:"video_type"`
	VideoHeight            int       `json:"video_height"`
	VideoWidth             int       `json:"video_width"`
	ImageHeight            int       `json:"image_height" xorm:"not null comment('图片高度') INT(11)"`
	ImageWidth             int       `json:"image_width" xorm:"not null comment('图片宽度') INT(11)"`
	CLocationsXzqhNation   string    `json:"c_locations_xzqh_nation" xorm:"default '' comment('国家') VARCHAR(64)"`
	CLocationsXzqhProvince string    `json:"c_locations_xzqh_province" xorm:"default '' comment('省份') VARCHAR(64)"`
	CLocationsXzqhCity     string    `json:"c_locations_xzqh_city" xorm:"default '' comment('城市') VARCHAR(64)"`
	CLocationsName         string    `json:"c_locations_name" xorm:"default '' comment('地名') VARCHAR(64)"`
	CLocationsHeight       int       `json:"c_locations_height" xorm:"comment('海拔') INT(11)"`
	CLocationsLocationLon  float64   `json:"c_locations_location_lon" xorm:"not null comment('经度') DOUBLE"`
	CLocationsLocationLat  float64   `json:"c_locations_location_lat" xorm:"not null comment('维度') DOUBLE"`
	CreateTime             time.Time `json:"create_time" xorm:"not null comment('生成时间') DATETIME"`
	UpdateTime             time.Time `json:"update_time" xorm:"not null comment('更新时间') DATETIME"`
	IsChina                int       `json:"is_china"`      // 是否是国内(10：国内，20：国外)
	PositionType           int       `json:"position_type"` // 地址类型（1：区域，2：点位）
	Town                   string    `json:"town"`
	Address                string    `json:"address"`
	County                 string    `json:"county"`
}

func (*FootprintMediaDetail) TableName() string {
	return "liveearth_primary.footprint_media_detail"
}
