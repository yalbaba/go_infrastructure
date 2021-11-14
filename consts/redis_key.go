package consts

// 计数器以counter开通
// 内容缓存以各个服务名称简称开头

const (
	// redis 中的直播间数据 key
	LiveRoomKey = "content:live_room:{@live_room_id}"

	// 直播间每天的数据统计key(Hash)
	LiveRoomViewStatisticsKey = "statistics:live_room:{@live_room_id}:{@date}"

	// 直播间每天的UV统计key(Hyperloglog)
	LiveRoomViewUVKey = "statistics:live_room_uv:{@live_room_id}:{@date}"

	// 直播间分享模板数据, 这是一个 kv
	LiveRoomShareTemplateKey = "content:live_room_share_template:{@live_room_id}"

	// 点赞直播间的用户id集合(真实数据)
	LikeRoomUsersKey = "counter:like-live_room-user_id:{@live_room_id}"

	// 点赞直播间的用户数(假数据)
	LikeRoomUsersNumKey = "counter:like-robot_live_room:{@live_room_id}"

	// 点赞主播发言的用户集合
	LikeAnchorCommentUsersKey = "counter:like-anchor_comment-user_id:{@anchor_comment_id}"

	// 足迹详情
	FootprintDetailKey = "content:footprint:{@footprint_id}"

	// 点赞足迹的用户集合
	LikeFootprintUsersKey = "counter:like-footprint-user_id:{@footprint_id}"

	// 天气数据
	AreaWeatherKey = "weather:{@area_id}"

	// 百科喜欢的人集合
	EncyclopediaLikeKey = "counter:like-encyclopedia-user_id:{@encyclopedia_id}"

	// 百科收藏的人集合
	EncyclopediaCollectKey = "counter:collect-encyclopedia-user_id:{@encyclopedia_id}"

	// 百科详情
	EncyclopediaDetailKey = "content:encyclopedia:{@encyclopedia_id}"

	// 直播间收藏的人的集合
	LiveRoomCollectKey = "counter:collect-live_room-user_id:{@live_room_id}"

	// 足迹收藏的人的集合
	FootprintCollectKey = "counter:collect-footprint-user_id:{@footprint_id}"

	LiveRoomShareUVKey = "counter:live_room_share_uv:{@live_room_id}:{@date}"

	LiveRoomShareDateKey = "counter:live_room_share_data:{@live_room_id}:{@date}"

	// 热门搜索词点击数记录, 这是一个hash, key是id, value是次数
	HotSearchWordClickNumKey = "counter:host_search_word:click"

	// 直播间集合id列表(一个list)
	LiveRoomCollectionIdsKey = "content:live_room_collections"

	// 直播间集合详情
	LiveRoomCollectionKey = "content:live_room_collection:{@collection_id}"

	// 每个直播间集合的首页直播间id
	CollectionFirstPageLiveKey = "content:collection_live_room:{@collection_id}"
)

const (

	// 百科评论计数
	EncyclopediaCounterKey = "counter:comment_encyclopedia_count:{@encyclopedia_id}"

	// 百科查看计数, 这是一个hash, key是id, value是次数
	EncyclopediaViewCounterKey = "counter:encyclopedia:view"

	// 故事评论计数
	StoryCounterKey = "counter:comment_story_count:{@story_id}"

	// 主播分享评论计数
	AnchorTalkCounterKey = "counter:comment_anchor_count:{@anchor_talk_id}"

	// 评论回复计数
	CommentReplyKey = "counter:comment_reply_count:{@comment_id}"

	// 直播间历史消息
	LiveRoomHistoryMsgKey = "im:live_room-history:{@live_room_id}"

	// 直播间历史来了
	LiveRoomHistoryComeKey = "im:live_room-history-come:{@live_room_id}"

	// 直播间发言数量统计
	LiveRoomSpeakTotalKey = "counter:im-live_room-speak_total:{@live_room_id}"

	// 直播间发言人数去重
	LiveRoomSpeakUserKey = "counter:im-live_room-speak-user_id:{@live_room_id}"

	// 进入直播间的历史用户集合(观看人数真实数据)
	LiveRoomHistoryUserKey = "counter:im-live_room-user_id:{@live_room_id}"

	// 直播间机器人观看数据(假数据)
	LiveRoomLookRobotNumKey = "counter:im-live_room-robot:{@live_room_id}"

	// 单个足迹的浏览量
	FootprintViewCountKey = "counter:footprint_view:{@footprint_id}"

	// 话题下的足迹浏览数量
	TopicFootprintViewCountKey = "counter:topic_footprint_view:{@topic_id}"

	// 栏目访问的人的集合
	ColumnViewUsersKey = "counter:view-column-user_id:{@column_id}"

	// 栏目访问人次
	ColumnViewsKey = "counter:view-column:{@column_id}"

	// 直播专题访问人次计数器, 这是一个set, value为人次
	LiveSpecialSubjectViewsKey = "counter:view-special_subject:{@id}"

	// 节目单访问人次, 这是一个set, value为人次
	LiveEPGViewsKey = "counter:view-epg:{@id}"

	MessagePushQueue = "message_push:queue:{@message_id}"

	// 视频标签分布
	VideoTagsDistribution = "distribution:video:tag-{@tag_id}"

	// 直播标签分布
	LiveRoomTagsDistribution = "distribution:live-room:tag-{@tag_id}"

	// 新闻标签分布
	NewsTagsDistribution = "distribution:news:tag-{@tag_id}"

	// VR视频标签分布
	VRVideoTagsDistribution = "distribution:vr_video:tag-{@tag_id}"

	// Map标签分布
	MapTagsDistribution = "distribution:map:tag-{@tag_id}"

	// 用户标签分布
	UserTagsDistribution = "distribution:user:user-{@user_id}"

	// 用户视频浏览游标
	UserScanVideoCursorKey = "cursor:user:user-{@user_id}:video"

	// 用户vr视频浏览游标
	UserScanVRVideoCursorKey = "cursor:user:user-{@user_id}:vr_video"

	// 用户直播间浏览游标
	UserScanLiveRoomCursorKey = "cursor:user:user-{@user_id}:live-room"

	// 用户新闻浏览游标
	UserScanNewsCursorKey = "cursor:user:user-{@user_id}:news"

	// 用户Map浏览游标
	UserScanMapCursorKey = "cursor:user:user-{@user_id}:map"
)

const (
	// 点赞评论集合
	CommentLikeKey = "comment:like_user_set:{@comment_id}"

	// 历史点赞评论集合
	CommentRawLikeKey = "comment:like_user_raw_set:{@comment_id}"

	// 主播在直播间分享的内容的数据
	AnchorCommentKey = "anchor_comment:{@anchor_comment_id}"
)

// 用户直播源
const (
	// 用户转播源加水印, 用于节约成本, 相同的源加同一个水印时只产生一个流名称, 这是一个set, value为 水印id_流名称
	UserLiveSourceRebroadcastWatermarkKey = "dup:user_source_rebroadcast_watermark"
)

// app配置
const (
	// hash key
	AppConfigKey = "app_config"
	// 每日一图文案
	ACKeyEDICopywriting = "edi_copywriting"
	// h5urls
	ACKeyH5Urls = "h5urls"
	// 启动页
	ACKeyStartPage = "start_page"
	// 兴趣页
	ACKeyInterestPage = "interest_page"
	// 文本资料
	ACKeyTextData = "text_data"
	// 安卓gis配置
	ACKeyGisConfigAndroid = "gis_config_android"
	// ios gis配置
	ACKeyGisConfigIos = "gis_config_ios"
)

// app版本(hash
const (
	AppNewVersionAndroid = "app_new_ver:android" // 安卓
	AppNewVersionIos     = "app_new_ver:ios"     // ios
)

// job id
const (
	// 足迹置顶时间范围, 这是一个hash, key是足迹id, val是job_id
	JobIdFootprintTopTimeKey = "job_id:footprint:top_time"
)

// 活动行为记录
const (
	// --------点击--------
	// 累计总数, 这是一个hash表, 只有一个hash表, key是活动id, value是总数
	ActivityBehaviorTotalPVKey = "behavior:activity:click_total:pv"
	// 累计人次, 这是一个HyperLogLog, 每个活动一个HyperLogLog, element是用户id
	ActivityBehaviorTotalUVKey = "behavior:activity:click_total:uv:{@activity_id}"
	// 当日总数, 这是一个hash表, 每天一个hash表, key是活动id, value是总数
	ActivityBehaviorDatePVKey = "behavior:activity:click_count:{@date}:pv"
	// 当日人次, 这是一个HyperLogLog, 每个活动每天一个HyperLogLog, element是用户id
	ActivityBehaviorDateUVKey = "behavior:activity:click_count:{@date}:uv:{@activity_id}"
	// --------曝光--------
	// 累计总数, 这是一个hash表, 只有一个hash表, key是活动id, value是总数
	ActivityBehaviorExposurePVKey = "behavior:activity:exposure:pv"
	// 累计人次, 这是一个HyperLogLog, 每个活动一个HyperLogLog, element是用户id
	ActivityBehaviorExposureUVKey = "behavior:activity:exposure:uv:{@activity_id}"
)

// 话题的相关key:预览图, zset, 存的key是足迹id,存的分是浏览量
const (
	TopicImageKey = "content:topic:image:{@topic_id}"
	// 以下是新话题使用的
	TopicViewUserAvatarsKey = "counter:topic:view_user_avatars:{@topic_id}" // 存浏览话题的用户头像
	TopicViewUsersKey       = "counter:topic:view_user_id:{@topic_id}"      // 存浏览话题的用户id
	TopicViewCountKey       = "counter:topic_view:{@topic_id}"              // 存浏览话题的参与数
	TopicKey                = "content:topic:{@topic_id}"                   // 话题基本信息key
	TopicArticleCountKey    = "counter:topic_Article:{@topic_id}"           // 话题动态计数
)

// 搜索
const (
	SearchTextCounterKey         = "counter:search:text"           // 搜索文本计数， 它是一个 sorted sets
	SearchDistrictChinaDataKey   = "content:district:china_data"   // 搜索国外区域数据
	SearchDistrictForeignDataKey = "content:district:foreign_data" // 搜索国内区域数据
)

// 小程序-发现id集合
const (
	AppletDiscoverIdsKey = "applet:discover_ids:{@column}"
)

// WeChatOA
const (
	WeChatOAAccessToken = "wechat_oa:access_token"
	WeChatOAJsApiTicket = "wechat_oa:jsapi_ticket"
)

// 看云
const (
	KanYunAccessToken  = "kanyun:access_token"
	KanYunShortUrlsMap = "kanyun:short_urls"
	KanYunLongUrlsMap  = "kanyun:long_urls"
)

const (
	// 系统消息推送的模板
	MessagePushTemplate = "message_push:template:{@template_id}"
	MessagePushID       = "message_push:message:{@platform}:{@push_id}"
)

// 用户缓存信息
const (
	UserCacheKey   = "user:{@user_id}"
	DeviceCacheKey = "device:{@device_id}"

	// 用户关注查询时间, 这是一个hash, key为用户id, value为查询时间标准文本
	UserAttentionQueueTimeKey = "user_attention_queue_time"
	// 用户关注的直播id列表, 这是一个zset, 每个用户一个set, 内容为直播id, 注意: 第一个值为占位符
	UserAttentionLiveList = "user_attention_live:{@user_id}"
	// 用户直播数计数器
	UserLiveNum = "content:live_num:{@user_id}"
	// 用户总视频数计数器
	UserAllVideoNum = "content:video_all_num:{@user_id}"
	// 用户审核通过视频数计数器
	UserVideoNum = "content:video_num:{@user_id}"
	// 用户总足迹数计数器
	UserAllFootprintNum = "content:footprint_all_num:{@user_id}"
	// 用户审核通过足迹数计数器
	UserFootprintNum = "content:footprint_num:{@user_id}"
	// 用户文章数计数器
	UserEncyclopediaNum = "content:encyclopedia_num:{@user_id}"
)

const (
	// 短视频详情数据 key
	VideoKey          = "content:video:{@video_id}"
	VideoShareUVKey   = "counter:video_share_uv:{@video_id}:{@date}"
	VideoShareDateKey = "counter:video_share_data:{@video_id}:{@date}"
	// 点赞视频的用户数(假数据)
	VideoUsersNumKey = "counter:like-robot_video:{@video_id}"
	// 视频每天的数据统计key(Hash)
	VideoViewStatisticsKey = "statistics:video:{@video_id}:{@date}"
	// 点赞视频的用户id集合(真实数据)
	VideoUsersKey = "counter:like-video-user_id:{@video_id}"
	// 视频收藏的人的集合
	VideoCollectKey = "counter:collect-video-user_id:{@video_id}"
	// 视频每天的UV统计key(Hyperloglog)
	VideoViewUVKey = "statistics:video_uv:{@video_id}:{@date}"
	// 视频累计的UV统计key
	VideoViewTotalUVKey = "statistics:video_uv:total:{@video_id}"
	// 单个视频的浏览量
	VideoViewCountKey = "counter:video_view:{@video_id}"
	// 视频评论计数
	VideoCounterKey = "counter:comment_video_count:{@video_id}"
	// 每天的视频观看排名统计
	VideoRankKey = "counter:rank_video:{@date}"
)

const (
	// 点播key
	LiveRoomPlayBack = "content:live_room_play_back:{@play_back_id}"
	// 回放视频评论计数
	PlayBackCounterKey = "counter:comment_play_back_count:{@play_back_id}"
)

const (
	//景区打卡评论计数
	ClockCounterKey = "counter:comment_clock_count:{@clocking_id}"
)

// 用户增长活动
const (
	// 用户增长活动显示入口, 值为 1=开; 0=关
	UserIncreaseActivityShowEntryKey = "user_incr_activity:show_entry"
	// 用户增长活动临时用户名跑马灯, 类型为set, 值为用户名
	UserIncreaseActivityTempUserNamesKey = "user_incr_activity:temp_user_names"
	// 用户增长活动规则, 类型为string
	UserIncreaseActivityRegulationKey = "user_incr_activity:regulation"
	// 新用户为别人助力过的记录, 这是一个set, value为用户id
	UserIncreaseActivityNewHelperUserIdKey = "user_incr_activity:new_helper_user"
	// 老用户为别人助力过的记录, 每个活动一个set, value为用户id
	UserIncreaseActivityOldHelperUserIdKey = "user_incr_activity:old_helper_user:{@id}"
)

// 拉新活动通知相关key
const (
	ActivityStartMiniAppJobIdKey = "activity:notify:mini_app:{@activity_id}" // 小程序活动开始通知任务key
	ActivityStartAppJobIdKey     = "activity:notify:app:{@activity_id}"      // app活动开始通知任务key
	ActivityDrawJobIdKey         = "activity:notify:app_draw:{@activity_id}" // app活动中奖通知任务key
	ActivityLuckDrawJobIdKey     = "activity:luck_draw:{@activity_id}"       // 抽奖任务key
)

// 标签id 获取 标签名称
const (
	// 根据标签id 获取标签名称
	TagId2TagName = "tag_manager:tag_id2tag_name"
	// 根据关键词Id 获取关键词名称
	KeywordsId2KeywordsName = "keywords:keywords_id2keywords_name"
	// 根据栏目Id   获取栏目名称
	ColumnId2ColumnName = "column:column_id2column_name"
)

// 景区
const (
	// 景区内容, 这是一个kv, 每个景区一条数据
	LandscapeInfoKey = "content:landscape:{@landscape_id}"
	// 景区解说包, 这是一个kv, 每个解说包一条数据
	GuideCommentPkgInfoKey = "content:guide_comment_pkg:{@pkg_id}"
	// 景区收藏的人的集合(Set), 每个景区一条数据, 里面存的是用户id
	LandscapeCollectKey = "counter:collect-landscape:{@landscape_id}"
	// 景区查看数(hash), 所有景区一条数据, key是景区id, value是查看数
	LandscapeViewKey = "counter:view-landscape"
	// 全国热门景区的首页数据(set 存第一页全国的景区id)
	HotFirstPageLandscapeIds = "content:hot_landscape_ids"
)

// 打卡
const (
	// 打卡数据内容(kv), 每个打卡一条数据
	ClockingInfoKey = "content:clocking:{@clocking_id}"
	// 打卡点赞的人的集合(set), 每个打卡一条数据, 里面存的是用户id
	ClockingLikeKey = "counter:like-clocking:{@clocking_id}"
	// 打卡查看数(hash), 所有打卡一条数据, key是打卡id, value是查看数
	ClockingViewKey = "counter:view-clocking"
)

//锦鲤红包活动
const (
	//每日活动各个阶段的奖品领取情况,数据结构：{"prize_count":奖品总数，是每个阶段多个奖品的总数,"taken_count":已领取总数，也是每个阶段奖品的总数}
	DailyActivityStepInfo = `activity:{@date}:{@activity_id}:{@step_id}`
	//每日红包领取金额总数
	DailyPrizeTakenAmountCount = `activity:daily:prize-taken`
	//每日的锦鲤报名人数假数据
	DailyEnterLuckyNum = `activity:daily:lucky-num`
	//用户领取红包的id缓存(set)
	UserTakenPrizeIdsKey = `activity:taken:{@activity_id}:{@user_id}`
)

//五一抽奖活动
const (
	// 活动列表key(list 存活动信息序列化数据的列表)
	WuyiActivityList = `wuyi_activity:activity_list`
	// 每个活动的中奖人信息（list）
	WuyiActivityLuckyManList = `wuyi_activity:lucky_man_list:{@activity_id}`
	// 抽奖结束最终的数据图片列表（list）
	WuyiActivityData = `wuyi_activity:data`
	// 活动分享次数统计
	WuyiActivityShare = `wuyi_activity:share:{@date}`
	// 活动app端pv统计
	WuyiActivityAppPv = `wuyi_activity:app:pv:{@date}`
	// 活动APP端uv统计
	WuyiActivityAppUv = `wuyi_activity:app:uv:{@date}`
	// 活动h5端pv统计
	WuyiActivityWebPv = `wuyi_activity:web:pv:{@date}`
	// 活动h5端uv统计
	WuyiActivityWebUv = `wuyi_activity:web:uv:{@date}`
	// 当日的ip访问池
	WuyiActivityIpSet = `wuyi_activity:ip_set:{@date}`
	// 当日的用户访问池
	WuyiActivityUserSet = `wuyi_activity:user_set:{@date}`
)

//欧洲杯数据统计相关key
const (
	//欧洲杯页面pv统计
	EuropeanCupPvKey = `european_cup:statistics:pv`
	//欧洲杯页面uv统计(set)
	EuropeanCupUvKey = `european_cup:statistics:uv`
	//欧洲杯分享次数统计
	EuropeanCupShareCount = `european_cup:statistics:share`
	//欧洲杯分享人数统计(set)
	EuropeanCupSharePeopleNum = `european_cup:statistics:share:people`
)

const (
	OlympicPvKey           = `tokyo_olympic:homepage:pv:{@date}`      //奥运会pv统计
	OlympicUvKey           = `tokyo_olympic:homepage:uv:{@date}`      //奥运会活动uv统计（set）
	OlympicAssistancePvKey = `tokyo_olympic:assistance:pv:{@date}`    //点击按钮pv统计
	OlympicAssistanceUvKey = `tokyo_olympic:assistance:uv:{@date}`    //点击按钮uv统计
	OlympicH5PvKey         = `tokyo_olympic:h5:pv:{@date}`            //h5底部点击pv
	OlympicH5UvKey         = `tokyo_olympic:h5:uv:{@date}`            //h5底部点击uv
	H5AssistancePvKey      = `tokyo_olympic:h5_assistance:pv:{@date}` //h5助力按钮pv
	H5AssistanceUvKey      = `tokyo_olympic:h5_assistance:uv:{@date}` //h5助力按钮uv
	SharePvKey             = `tokyo_olympic:share:pv:{@date}`         //app内h5分享pv
	ShareUvKey             = `tokyo_olympic:share:uv:{@date}`         //app内h5分享Uv
)

//故事相关key
const (
	StoryLikeUsersKey = "counter:like-story-user_id:{@story_id}" //故事点赞用户集合（真实数据） redis类型：set
	StoryViewCountKey = "counter:story_view:{@story_id}"         //故事观看真实数据 redis类型：计数器
	StoryKey          = "content:story:{@story_id}"              //故事详情key
)

//摄像头操控相关
const (
	CameraLine         = "camera:line:{@room_id}"           //摄像头操控排队（存deviceId）(redis类型：list)
	CurCameraUser      = "camera:cur_user:{@room_id}"       //当前操作用户（deviceId）（redis类型：string）
	CameraAccessToken  = "camera:access_token:{@plat_form}" //维护摄像头操控平台的token(redis类型：string)
	CameraRoomResetJob = "camera:reset:{@room_id}"          //维护直播间的摄像头的重置任务的任务id(redis类型：string)
)

//推流服务相关key
const (
	StreamList  = "stream:list"                 //存待推流的流名称（redis类型：list）
	StreamWhite = "stream:white:{@stream_name}" //推流白名单表示正在推流的流名称（redis类型：string，value：服务名称）
	StreamBlack = "stream:black"                //推流黑名单表示推流失败后的流名称（redis类型：zset,member：进程号+流名称，score：失败次数）
)
