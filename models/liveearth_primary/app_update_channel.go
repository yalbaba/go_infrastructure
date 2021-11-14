/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/22
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AppUpdateChannel struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AppUpdateId   int    `json:"app_update_id" xorm:"not null default 0 comment('app更新id') INT(10)"`
	ChannelId     int    `json:"channel_id" xorm:"not null default 0 comment('渠道id') INT(10)"`
	VersionNum    int    `json:"version_num" xorm:"not null default 0 comment('内部版本号') INT(10)"`
	DownloadUrl   string `json:"download_url" xorm:"not null default '' comment('下载链接') VARCHAR(1024)"`
	PackageSize   int    `json:"package_size" xorm:"not null default 0 comment('下载包大小') INT(10)"`
	PackageMd5    string `json:"package_md5" xorm:"not null default '' comment('包md5校验值') VARCHAR(32)"`
	VersionName   string `json:"version_name" xorm:"not null default '' comment('版本名') VARCHAR(32)"`
	ChannelName   string `json:"channel_name" xorm:"not null default '' comment('渠道名') VARCHAR(32)"`
	ChannelEnName string `json:"channel_en_name" xorm:"not null default '' comment('渠道英文名') VARCHAR(32)"`
}

func (*AppUpdateChannel) TableName() string {
	return "liveearth_primary.app_update_channel"
}
