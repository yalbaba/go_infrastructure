/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/4/8
   Description :
-------------------------------------------------
*/

package guide

import (
	"time"
)

type LandscapeNotice struct {
	Id              int       `json:"id" xorm:"autoincr"`
	LandscapeViewId int       `json:"landscape_view_id"`
	NoticeTitle     string    `json:"notice_title"` // 公告标题
	NoticeUrl       string    `json:"notice_url"`   // 公告地址
	CreateTime      time.Time `json:"create_time"`
}

func (m *LandscapeNotice) TableName() string {
	return "guide.landscape_notice"
}
