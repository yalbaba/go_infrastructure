package guide

import "time"

type CommentPkg struct {
	Id              int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"`
	CommentPkgName  string    `json:"comment_pkg_name" gorm:"column:comment_pkg_name"`   // 解说包名称
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区视图id
	LandscapeId     int       `json:"landscape_id" gorm:"column:landscape_id"`           // 景区id
	ViewpointNum    int       `json:"viewpoint_num" gorm:"column:viewpoint_num"`         // 景点数
	ListenNum       int       `json:"listen_num" gorm:"column:listen_num"`               // 收听数
	BaseListenNum   int       `json:"base_listen_num" gorm:"column:base_listen_num"`     // 基础收听数
	State           int8      `json:"state" gorm:"column:state"`                         // 状态; 10 = 关 ; 20 = 开
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime      time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *CommentPkg) TableName() string {
	return "guide.comment_pkg"
}
