package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	TxKey     = "MqggNDhmJNK316fzlWQzuTrTCVpuVrzG"
	TxHost    = "cdn.live.earthonline.com"
	TxAppname = "live"

	// TxKey     = "EarthliveChaoShiKong23213Csklive"
	// TxHost    = "txlive.national-space.com"
	// TxAppname = "live"

	ALYKey     = "B!22K3L12LhhYNaaUEarthUca8"
	ALYHOST    = "live.national-space.com"
	ALYAPPNAME = "live"
)

const (
	ThirdPartyStreamValid = 0
	TencentStreamValid    = 300 * 60 * 1000
	AliyunStreamValid     = 300 * 60 * 1000
)

var Stream = &streamCli{
	appName: "live",
	playDomainMap: map[int]struct {
		Domain           string
		SecretKey        string
		ExpireTimeOffset int64
	}{
		// 直播地球用的
		0: {"cdn.live.earthonline.com", "MqggNDhmJNK316fzlWQzuTrTCVpuVrzG", 18000},
		// 直播源开放平台给用户转播用的
		1: {"deal.live.cdn.earthonline.com", "be3450d89bd287feddfe3cd52eee7001", 7200},
	},
	pushUrlFormat: "rtmp://%s/%s/%s",
	playUrlFormatMap: map[string]string{
		"rtmp": "rtmp://%s/%s/%s?txSecret=%s&txTime=%s",
		"flv":  "http://%s/%s/%s.flv?txSecret=%s&txTime=%s",
		"hls":  "http://%s/%s/%s.m3u8?txSecret=%s&txTime=%s",
		"udp":  "webrtc://%s/%s/%s?txSecret=%s&txTime=%s",
	},
}

type streamCli struct {
	appName       string // app名
	playDomainMap map[int]struct {
		Domain           string // 播流域名
		SecretKey        string // 播流key
		ExpireTimeOffset int64  // 播流有效时间偏移(秒)
	}
	pushUrlFormat    string            // 推流url格式
	playUrlFormatMap map[string]string // 播放url格式
}

// 构建推送地址
func (s *streamCli) MakePushUrl(streamName string, domain string) string {
	return fmt.Sprintf(s.pushUrlFormat, domain, s.appName, streamName)
}

// 构建播放地址
func (s *streamCli) MakePlayUrl(domainIndex int, streamName string, deadline time.Time, formatType string) string {
	playDomain, ok := s.playDomainMap[domainIndex]
	if !ok {
		panic(fmt.Errorf("未定义的域名索引: %d", domainIndex))
	}

	timeHexStr := strconv.FormatInt(deadline.Unix()-playDomain.ExpireTimeOffset, 16)
	secret := Md5(playDomain.SecretKey + streamName + timeHexStr)
	if format, ok := s.playUrlFormatMap[formatType]; ok {
		return fmt.Sprintf(format, playDomain.Domain, s.appName, streamName, secret, timeHexStr)
	}
	panic(fmt.Errorf("不支持的格式: %s", formatType))
}

type LiveRes struct {
	Url    string   `json:"url"`
	Valid  int      `json:"valid"` // 截至时间戳
	Source PlatType `json:"source"`
}

func Md5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}

type PlatType uint8

const (
	ThirdPartyStream PlatType = 10*iota + 10 // 直接返回url,主要来自三方直接可以播放的地址
	TencentStream                            // 腾讯云
	AliyunStream                             // 阿里云
)

func GetStreamObject(streamName string, streamSource PlatType) (*LiveRes, error) {
	res, err := DirectReturn(streamName) // 10状态 todo 后面会加redis缓存处理
	if err == nil {
		return res, nil
	}
	switch streamSource {
	case TencentStream: // 腾讯
		return getstreameObjectFromtx(streamName)
	case AliyunStream: // 阿里
		return getstreameObjectFromAliyun(streamName)
	default: // 默认返回腾讯云
		return getstreameObjectFromtx(streamName)
	}
}

func getstreameObjectFromtx(streamName string) (*LiveRes, error) {
	nowSec := time.Now().Unix()
	nowSecHexStr := strconv.FormatInt(nowSec, 16)
	secretDoc := strings.Join([]string{TxKey, streamName, nowSecHexStr}, "")
	secret := Md5(secretDoc)
	m3u8Url := fmt.Sprintf("http://%s/%s/%s.flv?txSecret=%s&txTime=%s", TxHost, TxAppname, streamName, secret, nowSecHexStr)
	return &LiveRes{Url: m3u8Url, Valid: TencentStreamValid, Source: TencentStream}, nil
}

func getstreameObjectFromAliyun(streamName string) (*LiveRes, error) {
	nowSec := time.Now().Unix()
	nowSecStr := strconv.FormatInt(nowSec, 10)
	secretDoc := fmt.Sprintf("/%s/%s.m3u8-%s-0-0-%s", ALYAPPNAME, streamName, nowSecStr, ALYKey)
	secret := Md5(secretDoc)
	m3u8Url := fmt.Sprintf("http://%s/%s/%s.m3u8?auth_key=%s-0-0-%s", ALYHOST, ALYAPPNAME, streamName, nowSecStr, secret)
	return &LiveRes{Url: m3u8Url, Valid: AliyunStreamValid, Source: AliyunStream}, nil
}

func DirectReturn(name string) (*LiveRes, error) {
	directs := map[string]string{
		"zmlm": "https://mobilelive-play.ysp.cctv.cn/ysp/860059EAA5D6E07336D6B69A602D2BA513557780C783A692BBC6475489FF62DF44D0DC5B6F751D9716A0605E7AECEA796F61EB5686D5FE8EC93777B4A110E36F700F7EDA3A9E6D3D669C740BAEE395B457F4D3F7A1E8F71D079A71145966274282BD7F2B1694C2213939D1969221CB4F3FE566C107B576503AF2F49FEC22AA0C/2003445701_shd.flv",
		// "a86r7j6wkfq3up_q":"https://pl.live.weibo.com/alicdn/be3c4115c6afb3a1579a2d22df32cdbe_wb720.m3u8",
		// "adaspace20200529":"http://streamtest-pull.national-space.com/live/adaspace20200529.flv",
	}

	if url, ok := directs[name]; ok {
		return &LiveRes{
			Url:    url,
			Valid:  ThirdPartyStreamValid,
			Source: ThirdPartyStream,
		}, nil
	}
	return nil, errors.New("not in directs")
}
