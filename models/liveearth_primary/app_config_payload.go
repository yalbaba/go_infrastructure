/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/21
   Description :
-------------------------------------------------
*/

package liveearth_primary

// 启动页负载
type AppConfigStartPagePayload struct {
	MediaUrls []string `json:"media_urls"`
}

// 文本资料负载
type AppConfigTextDataPayload struct {
	Terms               string `json:"terms"`                 // 使用条款
	Policy              string `json:"policy"`                // 隐私政策
	UnregisterProtocol  string `json:"unregister_protocol"`   // 注销协议
	ChinaMobileProtocol string `json:"china_mobile_protocol"` // 中国移动使用协议
	ThirdPartyProtocol  string `json:"third_party_protocol"`  // 第三方协议
	UserProtocol        string `json:"user_protocol"`         // 用户协议
	Consociation        string `json:"consociation"`          // 商务合作
	About               string `json:"about"`                 // 关于
}
