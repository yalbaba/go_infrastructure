package liveearth_primary

import (
	"time"
)

type Footprint struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('足迹id') INT(10)"`
	UserId             string    `json:"user_id" xorm:"not null comment('用户id') index VARCHAR(32)"`
	FootprintDesc      string    `json:"footprint_desc" xorm:"comment('足迹的总文案') VARCHAR(1024)"`
	CoverUrl           string    `json:"cover_url" xorm:"not null comment('足迹首图') VARCHAR(1024)"`
	CoverWidth         int       `json:"cover_width"`
	CoverHeight        int       `json:"cover_height"`
	MainPictureTime    time.Time `json:"main_picture_time" xorm:"not null default CURRENT_TIMESTAMP comment('足迹首图的时间') TIMESTAMP"`
	IsPublish          int       `json:"is_publish" xorm:"not null comment('是否发布；10：不发布，20：发布') TINYINT(4)"`
	State              int       `json:"state" xorm:"not null default 0 comment('状态：10：新增等待审核，20：修改后等待审核，30：通过审核，40:用户不发布，且该故事没有经过审核，50：用户不发布，但故事已经审核通过，60：审核不通过，70:三方审核不通过，90：删除状态') TINYINT(4)"`
	Score              int       `json:"score" xorm:"not null default 0 comment('前端排序使用的评分') INT(11)"`
	NationCount        int       `json:"nation_count" xorm:"not null default 0 comment('国家数') INT(11)"`
	ProvinceCount      int       `json:"province_count" xorm:"not null default 0 comment('省的数量') INT(11)"`
	CityCount          int       `json:"city_count" xorm:"not null default 0 comment('城市数') INT(11)"`
	PictureCount       int       `json:"picture_count" xorm:"not null default 0 comment('图片数') INT(11)"`
	VideoCount         int       `json:"video_count" xorm:"not null default 0 comment('视频数') INT(11)"`
	CLocations         string    `json:"c_locations" xorm:"comment('位置信息') VARCHAR(1024)"`
	CFlytoXzqhNation   string    `json:"c_flyto_xzqh_nation" xorm:"default '' comment('此条足迹的国家') VARCHAR(32)"`
	CFlytoXzqhProvince string    `json:"c_flyto_xzqh_province" xorm:"default '' comment('此条足迹的省') VARCHAR(32)"`
	CFlytoXzqhCity     string    `json:"c_flyto_xzqh_city" xorm:"default '' comment('此条足迹城市') VARCHAR(32)"`
	CFlytoName         string    `json:"c_flyto_name" xorm:"default '' comment('足迹的位置数据') VARCHAR(64)"`
	Altitude           int       `json:"altitude" xorm:"default 0 comment('海拔') INT(11)"`
	Lon                float64   `json:"lon" xorm:"comment('经度') DOUBLE"`
	Lat                float64   `json:"lat" xorm:"comment('纬度') DOUBLE"`
	CFlytoRectangle    string    `json:"c_flyto_rectangle" xorm:"comment('可视矩形') VARCHAR(255)"`
	StartPosition      string    `json:"start_position" xorm:"not null default '' comment('开始位置') VARCHAR(16)"`
	EndPosition        string    `json:"end_position" xorm:"not null default '' comment('结束位置') VARCHAR(16)"`
	ExposureNum        int       `json:"exposure_num" xorm:"not null default 0 comment('曝光量') INT(11)"`
	ExposureUser       int       `json:"exposure_user" xorm:"not null default 0 comment('曝光用户数') INT(11)"`
	LikeNum            int       `json:"like_num" xorm:"not null default 0 INT(11)"`
	ReplyNum           int       `json:"reply_num" xorm:"not null default 0 INT(11)"`
	CreateTime         time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime         time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('用户更新时间') TIMESTAMP"`
	RecommendScore     int       `json:"recommend_score" xorm:"not null default 0 TINYINT(4)"`
	RecommendScoreTemp int       `json:"recommend_score_temp" xorm:"not null default 0 TINYINT(4)"`
	RecommendTimeType  int       `json:"recommend_time_type"` // 推荐时间类型: 10=自定义; 20=1小时; 30=3小时; 40=6小时; 50=12小时; 60=1天; 70=3天; 80=7天
	RecommendStartTime time.Time `json:"recommend_start_time" xorm:"TIMESTAMP"`
	RecommendEndTime   time.Time `json:"recommend_end_time" xorm:"TIMESTAMP"`
	Top                int       `json:"top" xorm:"not null default 0 TINYINT(4)"`
	TopTimeType        int       `json:"top_time_type"` // 置顶时间类型: 10=自定义; 20=1小时; 30=3小时; 40=6小时; 50=12小时; 60=1天; 70=3天; 80=7天
	TopTemp            int       `json:"top_temp" xorm:"not null default 0 TINYINT(4)"`
	TopStartTime       time.Time `json:"top_start_time" xorm:"TIMESTAMP"`
	TopEndTime         time.Time `json:"top_end_time" xorm:"TIMESTAMP"`
	FirstAuditUserId   int       `json:"first_audit_user_id" xorm:"comment('首次审核人员id') INT(11)"`
	FirstAuditTime     time.Time `json:"first_audit_time" xorm:"comment('首次审核通过时间') TIMESTAMP"`
	AuditUserId        int       `json:"audit_user_id" xorm:"comment('审核人员id') INT(11)"`
	AuditUserName      string    `json:"audit_user_name" xorm:"comment('审核人名字') VARCHAR(32)"`
	AuditTime          time.Time `json:"audit_time" xorm:"comment('审核通过时间') TIMESTAMP"`
	IsAdminCreate      int       `json:"is_admin_create" xorm:"not null default 0 comment('是否为管理员创建; 10=否; 20=是') TINYINT(4)"`
	CreateUserName     string    `json:"create_user_name" xorm:"comment('创建人名字') VARCHAR(32)"`
	MatchUserIds       string    `json:"match_user_ids"`
	MatchDistrictIds   string    `json:"match_district_ids"`
	Keywords           string    `json:"keywords"`
	IsPgc              int       `json:"is_pgc"`
	HasVr              int       `json:"has_vr"`
	BaseViewNum        int       `json:"base_view_num"`
}

func (*Footprint) TableName() string {
	return "liveearth_primary.footprint"
}
