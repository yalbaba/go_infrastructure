package liveearth_primary

import (
	"time"
)

type LiveSource struct {
	Id                 int64     `json:"id" xorm:"not null pk autoincr INT(10)"` //
	LiveRoomId         int64     `json:"live_room_id"`                           // LiveRoomId 直播间id,不为0表示和直播间关联
	SourceName         string    `json:"source_name"`                            // SourceName 直播源名称
	SourceStream       string    `json:"source_stream"`                          // SourceStream 流名称: 可以有多个,不同的服务平台流名称可能不同
	RecommendScore     int       `json:"recommend_score"`                        // RecommendScore 推荐分值
	LiveSourceUrls     string    `json:"live_source_urls"`                       // LiveSourceUrls 直播源(可以有多个): [''url1'',''url2'']
	SourceType         int       `json:"source_type"`                            // SourceType 直播源类型 10 = 直播; 20 = 录播
	LogoUrl            string    `json:"logo_url"`                               // LogoURL logo 地址
	MusicUrl           string    `json:"music_url"`                              // MusicURL 配乐地址
	CoverUrl           string    `json:"cover_url"`                              // CoverURL 封面图
	PushPull           int       `json:"push_pull"`                              // PushPull 10推流 20拉流
	Ownership          int       `json:"ownership"`                              // Ownership 所有权 10 = 自己的 20 = 第三方
	TimeDiff           int       `json:"time_diff"`                              // TimeDiff 与北京时间的时差, 小时
	Resolution         int       `json:"resolution"`                             // Resolution 清晰度 10 = 480; 20 = 480@60; 30 = 720; 40 = 720@60; 50 = 1080; 60 = 1080@60
	ThirdSource        string    `json:"third_source"`                           // ThirdSource 第三方来源名: 如斗鱼,人民网
	PositionType       int8      `json:"position_type"`                          // 点位类型; 10=区域; 20=点位
	Continent          string    `json:"continent"`                              // 州
	Nation             string    `json:"nation"`                                 // Nation 国家
	Province           string    `json:"province"`                               // Province 省
	City               string    `json:"city"`                                   // City 市
	Address            string    `json:"address"`                                // Address 地点名称
	Lon                float64   `json:"lon"`                                    // Lon 经度
	Lat                float64   `json:"lat"`                                    // Lat 纬度
	County             string    `json:"county" xorm:"VARCHAR(32)"`              //
	Town               string    `json:"town" xorm:"VARCHAR(32)"`                //
	Altitude           int       `json:"altitude"`                               // Altitude 海拔
	State              int       `json:"state"`                                  // State 10 删除 20 启用
	Operator           string    `json:"operator"`                               // Operator 操作人
	Quality            int       `json:"quality"`                                // Quality 品质
	SourceClass        int       `json:"source_class"`                           // SourceClass 内容分类 10= 动物
	CreateTime         time.Time `json:"create_time"`                            //
	UpdateTime         time.Time `json:"update_time"`                            //
	TopLevel           int       `json:"top_level"`                              // TopLevel 置顶等级
	TopLevelTemp       int       `json:"top_level_temp"`                         //
	TopStartTime       time.Time `json:"top_start_time"`                         // TopStartTime 置顶开始时间
	TopEndTime         time.Time `json:"top_end_time"`                           // TopEndTime 置顶结束时间
	AreaId             int       `json:"area_id"`                                //
	AuthorName         string    `json:"author_name"`                            // AuthorName 作者名
	AuthorAvatar       string    `json:"author_avatar"`                          // 作者头像
	AuthorId           string    `json:"author_id"`                              // AuthorId 作者Id
	PlayNumber         int       `json:"play_number"`                            // PlayNumber 播放量
	StartLiveTime      time.Time `json:"start_live_time"`                        // StartLiveTime 开始直播时间
	CrawlTime          time.Time `json:"crawl_time"`                             // CrawlTime 抓取时间
	PublishTime        time.Time `json:"publish_time"`                           //
	PushStreamState    int       `json:"push_stream_state"`                      // PushStreamState 推流状态 10 推流中 20已中断
	AuditTime          time.Time `json:"audit_time"`                             // AuditTime 审核时间
	AuditUserName      string    `json:"audit_user_name"`                        // AuditUserName 审核人名称
	AuditState         int       `json:"audit_state"`                            // AuditState 审核状态 10 未审核 20 审核通过 30审核不通过
	Origin             int       `json:"origin"`                                 // 直播源来源 10聚合 20手动
	AuditRemark        string    `json:"audit_remark"`                           // 审核不通过原因
	SourceDesc         string    `json:"source_desc"`                            // 源描述
	BannerLevel        int       `json:"banner_level"`                           //
	BannerCoverUrl     string    `json:"banner_cover_url"`                       //
	BannerStartTime    time.Time `json:"banner_start_time"`                      //
	BannerEndTime      time.Time `json:"banner_end_time"`                        //
	IsFullData         int       `json:"is_full_data"`                           //
	SpiderImportRaw    string    `json:"spider_import_raw"`                      //
	MusicId            int       `json:"music_id"`                               // 直播源音乐id
	MusicType          int8      `json:"music_type"`                             //
	CanChoose          int       `json:"can_choose"`                             // 是否可选音乐（10，能，20，不能）
	Keywords           string    `json:"keywords"`                               //
	IsSatelliteTv      int       `json:"is_satellite_tv"`                        //
	IsShareScreen      int       `json:"is_share_screen"`                        //
	RecommendLevel     int       `json:"recommend_level"`                        //
	RecommendStartTime time.Time `json:"recommend_start_time"`                   //
	RecommendEndTime   time.Time `json:"recommend_end_time"`                     //
	LogoPosition       int8      `json:"logo_position"`                          // logo位置： 11 左上  12 中上  13 右上   21 左下  22 中下
	IsTranscribe       bool      `json:"is_transcribe"`                          // 是否录制
}

func (*LiveSource) TableName() string {
	return "liveearth_primary.live_source"
}
