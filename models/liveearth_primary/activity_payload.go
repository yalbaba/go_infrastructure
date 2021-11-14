/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/11
   Description :
-------------------------------------------------
*/

package liveearth_primary

// 活动表负载内容
type ActivityPayload struct {
	MediaUrls       []string `json:"media_urls,omitempty"` // 媒体url
	HeaderMediaUrls []string `json:"header_media_urls,omitempty"`

	SkipUrl string  `json:"skip_url,omitempty"` // 跳转url
	SkipLon float64 `json:"skip_lon,omitempty"` // 跳转经度
	SkipLat float64 `json:"skip_lat,omitempty"` // 跳转纬度

	ShowTime         int  `json:"show_time,omitempty"`           // 显示时间, 秒
	ShowSkipWait     bool `json:"show_skip_wait,omitempty"`      // 是否显示跳过等待
	ShowSkipWaitTime int  `json:"show_skip_wait_time,omitempty"` // 显示跳过等待时间, 秒, 表示经过多少秒后显示跳过等待按钮

	ArticleId       int                    `json:"article_id,omitempty"`        // 文章id
	ArticleTitle    string                 `json:"article_title,omitempty"`     // 文章标题
	ArticleType     int                    `json:"article_type,omitempty"`      // 文章类型
	ArticleCoverUrl []string               `json:"article_cover_url,omitempty"` // 文章封面图
	ArticleVideoUrl string                 `json:"article_video_url,omitempty"` // 文章video地址
	CustomData      map[string]interface{} `json:"custom_data"`
}
