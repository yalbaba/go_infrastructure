/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/23
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AppUpdateChannelCache struct {
	Id            int    `json:"id"`
	AppUpdateId   int    `json:"app_update_id"`
	ChannelId     int    `json:"channel_id"`
	VersionNum    int    `json:"version_num"`
	UpdateContent string `json:"update_content"`
	DownloadUrl   string `json:"download_url"`
	PackageSize   int    `json:"package_size"`
	PackageMd5    string `json:"package_md_5"`
	VersionName   string `json:"version_name"`
	ChannelName   string `json:"channel_name"`
	ChannelEnName string `json:"channel_en_name"`
}
