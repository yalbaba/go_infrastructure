/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/22
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type AppUpdate struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	DeviceType    int       `json:"device_type" xorm:"not null default 10 comment('设备类型: 20=安卓; 30=ios') index TINYINT(3)"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Operator      int       `json:"operator" xorm:"not null default 0 comment('操作人id') INT(10)"`
	OperatorName  string    `json:"operator_name" xorm:"not null default '' comment('操作人名字') VARCHAR(32)"`
	Deleted       int       `json:"deleted" xorm:"not null default 10 comment('删除标记 10 = 未删除; 20 = 删除') TINYINT(3)"`
	DeleteTime    time.Time `json:"delete_time" xorm:"comment('删除时间') TIMESTAMP"`
	VersionName   string    `json:"version_name" xorm:"not null default '' comment('版本名') VARCHAR(32)"`
	VersionNum    int       `json:"version_num" xorm:"not null default 0 comment('内部版本号') INT(10)"`
	DownloadUrl   string    `json:"download_url" xorm:"not null default '' comment('下载链接') VARCHAR(1024)"`
	UpdateLogic   int       `json:"update_logic" xorm:"not null default 20 comment('更新逻辑: 10=最新版本; 20=不提示; 30=提示; 40=强制') TINYINT(3)"`
	UpdateContent string    `json:"update_content" xorm:"not null default '' comment('更新内容') VARCHAR(32)"`
	State         int       `json:"state" xorm:"not null default 10 comment('开启状态: 10=关闭; 20=开启') TINYINT(3)"`
}

func (*AppUpdate) TableName() string {
	return "liveearth_primary.app_update"
}
