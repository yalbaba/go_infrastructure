package consts

//消息队列名称

const (
	// 直播源变更同步
	LiveSourceStreamSync = "live_source_stream_sync"

	// 机器人足迹点赞队列
	RobotLikeFootprint = "robot_like_footprint"

	// 机器人足迹关注队列
	RobotFollowFootprint = "robot_follow_footprint"

	// 机器人视频关注队列
	RobotFollowVideo = "robot_follow_video"

	// 机器人直播关注队列
	RobotFollowLive = "robot_follow_live"

	// 机器人发现关注队列
	RobotFollowEncyclopedia = "robot_follow_encyclopedia"

	// 机器人发现关注队列
	RobotFollowClocking = "robot_follow_clocking"

	// 机器人视频点赞队列
	RobotLikeVideo = "robot_like_video"

	// 机器人直播点赞队列
	RobotLikeLive = "robot_like_live"

	// 机器人发现点赞队列
	RobotLikeEncyclopedia = "robot_like_encyclopedia"

	// 机器人打卡点赞队列
	RobotLikeClocking = "robot_like_clocking"

	// 足迹置顶时间任务
	QueueFootprintTopTask = "footprint_top_task"
)

const (
	//分享直播间点击记录
	ShareClick = "share_click"
)

//拉新活动队列名称
const MiniAppActivityStartNotify = "miniapp_activity_start_notify" //小程序活动开始通知
const AppActivityNotify = "app_activity_notify"                    //app活动通知（活动开始提醒和中奖提醒）
const LuckDraw = "luck_draw"                                       //抽奖任务
const SyncVideoWatch = "recommend_sync_video_watch"                //视频观看同步推荐系统
const SyncVideoPvUv = "sync_video_pv_uv"                           //上报时同步视频pv、uv
