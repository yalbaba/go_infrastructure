/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type Clocking struct {
	Id               int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"`
	UserId           string    `json:"user_id" gorm:"column:user_id"`               // 用户id
	ClockingTitle    string    `json:"clocking_title" gorm:"column:clocking_title"` // 打卡标题
	ClockingDesc     string    `json:"clocking_desc" gorm:"column:clocking_desc"`
	CoverUrl         string    `json:"cover_url" gorm:"column:cover_url"`       // 封面图
	CoverWidth       int       `json:"cover_width" gorm:"column:cover_width"`   // 封面的宽
	CoverHeight      int       `json:"cover_height" gorm:"column:cover_height"` // 封面的高
	MediaType        int8      `json:"media_type" gorm:"column:media_type"`
	MediaCount       int       `json:"media_count" gorm:"column:media_count"`               // 媒体总数
	Continent        string    `json:"continent" gorm:"column:continent"`                   // 州
	Nation           string    `json:"nation" gorm:"column:nation"`                         // 国
	Province         string    `json:"province" gorm:"column:province"`                     // 省
	City             string    `json:"city" gorm:"column:city"`                             // 市
	County           string    `json:"county" gorm:"column:county"`                         // 区县
	Town             string    `json:"town" gorm:"column:town"`                             // 镇
	Address          string    `json:"address" gorm:"column:address"`                       // 详细地址
	Location         Point     `json:"location" gorm:"column:location"`                     // 定位点
	Altitude         int       `json:"altitude" gorm:"column:altitude"`                     // 海拔(米)
	PositionType     int8      `json:"position_type" gorm:"column:position_type"`           // 地址类型（10：区域，20：点位）
	FirstTagId       int       `json:"first_tag_id" gorm:"column:first_tag_id"`             // 一级tag
	SecondTagIds     string    `json:"second_tag_ids" gorm:"column:second_tag_ids"`         // 二级tagId列表, 用英文逗号连接
	KeywordIds       string    `json:"keyword_ids"`                                         // 关键词id列表
	MatchUserIds     string    `json:"match_user_ids" gorm:"column:match_user_ids"`         // 匹配用户id, json
	MatchDistrictIds string    `json:"match_district_ids" gorm:"column:match_district_ids"` // 匹配区域id, json
	VisibleStartTime time.Time `json:"visible_start_time" gorm:"column:visible_start_time"` // 可见开始时间
	VisibleEndTime   time.Time `json:"visible_end_time" gorm:"column:visible_end_time"`     // 可见结束时间
	PublishState     int8      `json:"publish_state" gorm:"column:publish_state"`           // 发布状态; 10=不发布; 20=发布
	State            int8      `json:"state" gorm:"column:state"`                           // 状态: 10 = 待审核; 20 = 审核通过; 30 = 未通过; 40=草稿; 90 = 下架
	FirstAuditTime   time.Time `json:"first_audit_time"`
	PublishTime      time.Time `json:"publish_time"`
	Deleted          int8      `json:"deleted" gorm:"column:deleted"`   // 删除标志: 10 = 未删除; 20 = 已经删除
	Operator         string    `json:"operator" gorm:"column:operator"` // 操作人
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime       time.Time `json:"update_time" gorm:"column:update_time"`
	FetchTime        time.Time `json:"fetch_time" gorm:"column:fetch_time"` // 爬虫抓取时间
	FetchOrigin      string    `json:"fetch_origin"`
	BaseViewNum      int       `json:"base_view_num"`
}

func (m *Clocking) TableName() string {
	return "guide.clocking"
}
