/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

// 启动页
type AppConfigStartPageCache struct {
	DeviceType int      `json:"device_type,omitempty"` // 设备类型: 10=所有; 20=安卓; 30=ios
	MediaUrls  []string `json:"media_urls,omitempty"`  // 媒体url
	State      int      `json:"state,omitempty"`       // 开启状态: 10=关闭; 20=开启
}

// 兴趣页
type AppConfigInterestPageCache struct {
	HintProtocolDeviceType int      `json:"hint_protocol_device_type,omitempty"` // 提示协议设备类型; 10=所有; 20=安卓; 30=ios
	InterestTags           []string `json:"interest_tags,omitempty"`             // 兴趣标签
	InterestTagsName       []string `json:"interest_tag_names,omitempty"`        // 兴趣标签名
	AreaTags               []string `json:"area_tags,omitempty"`                 // 区域标签
	AreaTagsName           []string `json:"area_tag_names,omitempty"`            // 区域标签名
}

// 文本资料
type AppConfigTextDataCache struct {
	Terms               string `json:"terms,omitempty"`                 // 使用条款
	Policy              string `json:"policy,omitempty"`                // 隐私政策
	UnregisterProtocol  string `json:"unregister_protocol"`             // 注销协议
	ChinaMobileProtocol string `json:"china_mobile_protocol,omitempty"` // 中国移动使用协议
	ThirdPartyProtocol  string `json:"third_party_protocol,omitempty"`  // 第三方协议
	UserProtocol        string `json:"user_protocol,omitempty"`         // 用户协议
	Consociation        string `json:"consociation,omitempty"`          // 商务合作
	About               string `json:"about,omitempty"`                 // 关于
}
