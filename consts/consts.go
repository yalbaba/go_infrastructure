package consts

type LikeOp uint8

const (
	Like LikeOp = iota + 1
	CancelLike
)

//消息推送mqc队列名
const (
	NotificationTopicName = "message_push-notification"
	OperatePushTopic      = "operate_message_push"     //推送运营消息队列名
	AppOperatePushTopic   = "operate_app_message_push" //站内信消息队列名
	UserRemindPushTopic   = "user_remind_message_push" //用户唤活消息队列名
	PhoneMsgPushTopic     = "phone_message_push"       //短信推送消息队列名
)

const DownLoadVideoTopic = "live_video_download"

type ServerType uint8

const (
	HttpServer ServerType = iota + 1
	RpcServer
	MqcServer
	CronServer
	NsqConsumeServer
)

func (s ServerType) String() string {
	switch s {
	case HttpServer:
		return "http"
	case RpcServer:
		return "grpc"
	case MqcServer:
		return "mqc"
	case CronServer:
		return "cron"
	case NsqConsumeServer:
		return "nsq_consume"
	}

	return ""
}

// 数据类型
const (
	Live           = 30  // 直播
	Footprint      = 40  // 足迹
	Topic          = 60  // 话题
	Article        = 71  // 文章
	Video          = 80  // 视频
	SpecialSubject = 90  // 专题
	Landscape      = 100 // 景区
	Clocking       = 110 // 打卡
)

//通知跳转类型
const (
	WikiNotice     int8 = 40  //百科
	StoryNotice    int8 = 20  //故事
	LiveNotice     int8 = 10  //直播
	VideoNotice    int8 = 60  //视频
	ActivityNotice int8 = 120 //拉新活动
)

//通知的附件媒体类型
const (
	ImageFile = 20
	VideoFile = 30
)

// 通用审核状态
const (
	ContentWaitAudit   = 10 // 等待审核
	ContentAuditPass   = 20 // 审核通过
	ContentAuditNoPass = 30 // 审核不通过
	ContentDraft       = 40 // 草稿
	ContentSoldOut     = 90 // 下架
)

//五一活动id写死的
const WuyiActivityId = 1

//短信模板code
const (
	LookOverResult = "SMS_215802388" //五一活动参与人通知
	LuckyNotify    = "SMS_215820362" //活动结果通知
)
