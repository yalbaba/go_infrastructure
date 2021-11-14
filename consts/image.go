package consts

import (
	"strings"
)

const (
	cover = "image/auto-orient,1/resize,p_70/quality,q_90"
)

type Image struct {
	URL    string `json:"url"`   // 如果是视频, 这里放的是原始视频地址
	Thumb  string `json:"thumb"` // 如果是视频, 这里放的是加验证后的视频地址
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func NewImage(url string, Width, Height int) *Image {
	thumb := url
	if !strings.Contains(url, ".mafengwo.net") { // 马蜂窝的图片不能加后缀
		thumb = url + "?x-oss-process=" + cover
	}
	return &Image{
		URL:    url,
		Thumb:  thumb,
		Width:  Width,
		Height: Height,
	}
}
