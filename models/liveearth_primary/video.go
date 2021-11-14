package liveearth_primary

import "time"

// Video ...
type Video struct {
	Id int64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	// VideoName 视频名称
	VideoName string `json:"video_name"`
	// CoverUrl 封面图片/视频
	CoverUrl string `json:"cover_url"`
	// CoverMediaType 封面媒体类型: 10 = 图片; 20 = 视频
	CoverMediaType int `json:"cover_media_type"`
	// VideoDesc 视频描述
	VideoDesc string `json:"video_desc"`
	// VideoSource 视频来源
	VideoSource string `json:"video_source"`
	// VideoDuration 视频时长
	VideoDuration string `json:"video_duration"`
	// VideoUrl 视频链接
	VideoUrl string `json:"video_url"`
	// VideoOriginalUrl 视频原始链接(爬取用)
	VideoOriginalUrl string `json:"video_original_url"`
	// AuthorPhone 作者手机
	AuthorPhone string `json:"author_phone"`
	// Author 作者
	Author string `json:"author"`
	// PublishState 发布状态: 10 = 已下架; 20 = 未发布 ;30=发布
	PublishState int `json:"publish_state"`
	// Operator 操作人
	Operator string `json:"operator"`
	// VerifyReason 审核原因
	VerifyReason string `json:"verify_reason"`
	// VerifyState 审核状态: 10 = 等待抓取;11 = 正在抓取;19 = 抓取失败; 20 = 等待审核; 30= 审核不通过; 40 = 草稿; 90 = 审核通过
	VerifyState int `json:"verify_state"`
	// Mark 备注
	Mark string `json:"mark"`
	// BatchId 抓取批次号
	BatchId string `json:"batch_id"`
	// FetchExpireTime 抓取过期时间
	FetchExpireTime time.Time `json:"fetch_expire_time"`
	// FetchTimes 抓取次数(会尝试三次)
	FetchTimes int `json:"fetch_times"`
	// Deleted 删除标记: 10 = 未删除 ; 20 = 删除
	Deleted int8 `json:"deleted"`
	// Nation 国家
	Nation string `json:"nation"`
	// Province 省
	Province string `json:"province"`
	// City 市
	City string `json:"city"`
	// OriginalTime 原创时间
	OriginalTime time.Time `json:"original_time"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// FetchTime 抓取时间
	FetchTime time.Time `json:"fetch_time"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`
	// TopLevel 置顶等级
	TopLevel     int `json:"top_level"`
	TopLevelTemp int `json:"top_level_temp"`
	// TopStartTime 置顶开始时间
	TopStartTime time.Time `json:"top_start_time"`
	// TopEndTime 置顶结束时间
	TopEndTime time.Time `json:"top_end_time"`
	// RecommendLevel 推荐等级
	RecommendLevel     int `json:"recommend_level"`
	RecommendLevelTemp int `json:"recommend_level_temp"`
	// RecommendStartTime 推荐开始时间
	RecommendStartTime time.Time `json:"recommend_start_time"`
	// RecommendEndTime 推荐结束时间
	RecommendEndTime time.Time `json:"recommend_end_time"`
	// MatchUser 可见人
	MatchUser string `json:"match_user"`
	// MatchDistrict 可见区域
	MatchDistrict    string    `json:"match_district"`
	VisibleStartTime time.Time `json:"visible_start_time"`
	// VisibleEndTime 可见截止时间
	VisibleEndTime time.Time `json:"visible_end_time"`
	// VisibleType 可见时效类型: 10 = 24小时;20=48小时;30=72小时;40=7天;50=1个月;60=3个月;70=6个月;90=不限
	VisibleType int `json:"visible_type"`
	// UploadType 视频上传类型: 10 = 聚合; 20 = 本地上传
	UploadType             int       `json:"upload_type"`
	VideoOriginalName      string    `json:"video_original_name"`       //文件下载源名称，为了防止下载视频成功但上传失败,此时不用再重新下载
	VideoOriginalCoverName string    `json:"video_original_cover_name"` //文件爬取的图片名称
	UserId                 string    `json:"user_id"`                   //作者id 用户id
	Keywords               string    `json:"keywords"`                  //关键词
	BeLikeNum              int64     `json:"be_like_num"`               // 获赞数
	ViewNum                int64     `json:"view_num"`                  //观看量
	PublishTime            time.Time `json:"publish_time"`
	BaseViewNum            int64     `json:"base_view_num"` //基础观看人数
	BaseThumbUp            int64     `json:"base_thumb_up"` //基础点赞人数
	VideoWidth             int       `json:"video_width"`
	VideoHeight            int       `json:"video_height"`
	IsPushed               int       `json:"is_pushed"`     //该是否推送过消息（10：未推送过，20：已推送了）
	VideoType              int       `json:"video_type"`    //视频类型
	IsChina                int       `json:"is_china"`      //是否是国内(10：国内，20：国外)
	PositionType           int       `json:"position_type"` //地址类型（1：区域，2：点位）
	Continent              string    `json:"continent"`     //洲
	County                 string    `json:"county"`
	Town                   string    `json:"town"`
	Address                string    `json:"address"`
}

func (*Video) TableName() string {
	return "liveearth_primary.video"
}
