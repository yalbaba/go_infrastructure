package guide

import "time"

type ActivityRecord struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId            string    `json:"user_id" `             // 参与活动的用户id
	VerifyCode        string    `json:"verify_code" `         // 分享用户的标识（用于活动分享关联）
	ActivityId        int       `json:"activity_id" `         // 活动id
	CurrentStepId     int       `json:"current_step_id" `     // 当前进行的活动阶段id
	GrandTotalAmount  int       `json:"grand_total_amount" `  // 累计领取红包金额（单位：分）
	GrandTotalInvited int       `json:"grand_total_invited" ` // 累计已邀请人数
	IsEnterLucky      int8      `json:"is_enter_lucky" `      // 是否报名锦鲤：10：否，20：是
	HasConverted      int8      `json:"has_converted" `       // 是否已兑换：10，否，20，是
	Deleted           int8      `json:"deleted" `             // 删除标记：10否，20是
	InviteDeadline    time.Time `json:"invite_deadline" `     // 邀请截止时间
	CreateTime        time.Time `json:"create_time" `         // 创建时间
}

func (m *ActivityRecord) TableName() string {
	return "guide.activity_record"
}
