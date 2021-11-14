package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/mts"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

const (
	OssRegion          = "oss-cn-hangzhou"
	OssAccessKeyID     = "LTAI4G125o6Yxa3WZswiRFgG"
	OssAccessKeySecret = "jUi25EZJbAjau4LJB90MyXk6QjQsSC"
	OssRoleArn         = "acs:ram::1762212371458479:role/oss-user"
	OssSessionName     = "earthlive_ugc_user"
	OssEndpoint        = "oss-accelerate.aliyuncs.com"
	OssHost            = "https://cdn.image.earthonline.com/"
)

type OssBaseConf struct {
	OssRegion          string `json:"oss_region"`
	OssAccessKeyID     string `json:"oss_access_key_id"`
	OssAccessKeySecret string `json:"oss_access_key_secret"`
	OssRoleArn         string `json:"oss_role_arn"`
	OssSessionName     string `json:"oss_session_name"`
	OssEndpoint        string `json:"oss_endpoint"`
}

type OssConf struct {
	OssBaseConf
	OssBucket string `json:"oss_bucket"`
	OssPath   string `json:"oss_path"`
	OssHost   string `json:"oss_host"`
}

var ossBaseConf = OssBaseConf{
	OssRegion:          "cn-hangzhou",
	OssAccessKeyID:     "LTAI4G125o6Yxa3WZswiRFgG",
	OssAccessKeySecret: "jUi25EZJbAjau4LJB90MyXk6QjQsSC",
	OssRoleArn:         "acs:ram::1762212371458479:role/oss-user",
	OssSessionName:     "earthlive_ugc_user",
	OssEndpoint:        "oss-accelerate.aliyuncs.com",
}

var ossConfMap = map[string]*OssConf{
	"liveearth-image": &OssConf{
		OssBaseConf: ossBaseConf,
		OssPath:     "default",
		OssBucket:   "liveearth-image",
		OssHost:     "https://cdn.image.earthonline.com/",
	},
	"liveearth-video": &OssConf{
		OssBaseConf: ossBaseConf,
		OssBucket:   "liveearth-video",
		OssPath:     "default",
		OssHost:     "http://cdn.video.earthonline.com/",
	},
}

// 返回oss对象
type OssAliyunResponse struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Expiration      string `json:"expiration"` // 超时时间
	SecurityToken   string `json:"security_token"`
	Path            string `json:"path"`
	Bucket          string `json:"bucket"`
	Bukect          string `json:"bukect"` // 这个单词之前的开发写错了,这里要兼容之前的版本
	Host            string `json:"host"`
}

// 获取阿里云oss的key值
func GetAliyunOssAccesskeyUtil(bucketName, path string) (*OssAliyunResponse, error) {

	client, err := sts.NewClientWithAccessKey(
		ossConfMap[bucketName].OssRegion,
		ossConfMap[bucketName].OssAccessKeyID,
		ossConfMap[bucketName].OssAccessKeySecret,
	)
	if err != nil {
		return nil, err
	}

	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = ossConfMap[bucketName].OssRoleArn
	request.RoleSessionName = ossConfMap[bucketName].OssSessionName
	request.DurationSeconds = "3600"
	resp, err := client.AssumeRole(request)
	if err != nil {
		return nil, err // 报错
	}

	res := &OssAliyunResponse{}
	res.Expiration = resp.Credentials.Expiration
	res.AccessKeyID = resp.Credentials.AccessKeyId
	res.AccessKeySecret = resp.Credentials.AccessKeySecret
	res.Path = path
	res.Endpoint = ossConfMap[bucketName].OssEndpoint
	res.SecurityToken = resp.Credentials.SecurityToken
	res.Bucket = bucketName
	res.Bukect = bucketName
	res.Host = ossConfMap[bucketName].OssHost

	return res, nil
}

const (
	PipelineId    = "a1fca4c881ea487282b2ed48bb842742"
	WaterId       = "da07a2722770418ebf760eb752b8b86e"
	TemplatteId   = "65e4ab2e72144a60be92a24fb80043fa"
	InputLocation = "oss-cn-hangzhou"
	InputBucket   = "liveearth-image"
	OutPutBucket  = "liveearth-image"
	//RegionId, accessKeyId, accessKeySecret
)

type OutPut struct {
	OutputObject string        `json:"OutputObject"`
	TemplateId   string        `json:"TemplateId"`
	WaterMarks   []*WaterMarks `json:"WaterMarks"`
}
type WaterMarks struct {
	WaterMarkTemplateId string     `json:"WaterMarkTemplateId"`
	InputFile           *InputFile `json:"InputFile"`
}
type InputFile struct {
	URL string `json:"URL"`
}

//加水印接口参数
type AddWatermarkReq struct {
	WatermarkUrl    string `json:"watermark_url"`     //水印url
	InputBucket     string `json:"input_bucket"`      //输入oss的buket
	InputObjectName string `json:"input_object_name"` //输入的文件
}

type WatermarkResp struct {
	Url string `json:"url"` //加水印后的地址
}

func mtsInput(req *AddWatermarkReq) (string, error) {
	js := jsoniter.ConfigCompatibleWithStandardLibrary
	inputMap := map[string]string{"Bucket": req.InputBucket, "Location": InputLocation, "Object": req.InputObjectName}
	res, err := js.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func mtsOutPut(watermarkUrl, outPutFile string) (string, error) {
	waterFile := &InputFile{
		URL: watermarkUrl,
	}
	waterMarks := &WaterMarks{
		WaterMarkTemplateId: WaterId,
		InputFile:           waterFile,
	}
	waterMarkss := []*WaterMarks{waterMarks}

	outputs := &OutPut{
		OutputObject: "video/" + outPutFile,
		TemplateId:   TemplatteId,
		WaterMarks:   waterMarkss,
	}
	js := jsoniter.ConfigCompatibleWithStandardLibrary
	outputsList := []*OutPut{outputs}
	res, err := js.Marshal(outputsList)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

//生产用户id
func ProduceUserIdByUuid() string {
	userId := uuid.NewV4().String()
	userId = strings.ReplaceAll(userId, "-", "")
	return userId
}

func ossFileIsExist(BucketName, objectName string) (bool, error) {
	client, err := oss.New(OssEndpoint, OssAccessKeyID, OssAccessKeySecret)
	if err != nil {
		return false, err
	}
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return false, err
	}
	return bucket.IsObjectExist(objectName)
}
func VideoTranscodingAddWatermark(req *AddWatermarkReq) (*WatermarkResp, error) {

	//https://liveearth-video.oss-cn-hangzhou.aliyuncs.com/video/file_731052d1856f6e26d53c03d6b897fcd9.mp4
	if req.WatermarkUrl == "https://earthlive.oss-cn-hangzhou.aliyuncs.com/video/earthlive_log.png" || req.WatermarkUrl == "" {
		return &WatermarkResp{Url: "https://liveearth-video.oss-cn-hangzhou.aliyuncs.com/" + req.InputObjectName}, nil
	}

	var outputObjectFil string
	req.WatermarkUrl = strings.Replace(req.WatermarkUrl, "cdn.image.earthonline.com", "liveearth-image.oss-cn-hangzhou.aliyuncs.com", 1)
	if !strings.HasSuffix(req.WatermarkUrl, "png") {
		return nil, errors.New("水印不是png格式")
	}
	isExist, err := ossFileIsExist(req.InputBucket, req.InputObjectName)
	if err != nil {
		return nil, fmt.Errorf("检查文件是否存在失败,err:%v", err)
	}
	if !isExist {
		return nil, errors.New("视频文件在oss上不存在，或者路径错误")
	}
	input, err := mtsInput(req)
	if err != nil {
		return nil, err
	}
	outputObjectFil = ProduceUserIdByUuid() + ".mp4"
	output, err := mtsOutPut(req.WatermarkUrl, outputObjectFil)
	if err != nil {
		return nil, fmt.Errorf("创建水印地址失败，err:%v", err)
	}
	client, err := mts.NewClientWithAccessKey("cn-hangzhou", OssAccessKeyID, OssAccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("连接阿里云失败,err:%v", err)
	}
	request := mts.CreateSubmitJobsRequest()
	request.Scheme = "http"
	request.OutputBucket = OutPutBucket
	request.OutputLocation = "oss-cn-hangzhou"
	request.PipelineId = PipelineId
	request.Input = input
	request.Outputs = output
	request.Method = "POST"
	response, err := client.SubmitJobs(request)
	if err != nil {
		return nil, err
	}
	if response.BaseResponse.IsSuccess() {
		return &WatermarkResp{
			Url: OssHost + "video/" + outputObjectFil,
		}, nil
	} else {
		return nil, errors.New(response.BaseResponse.GetHttpContentString())
	}
}
