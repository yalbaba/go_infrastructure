package logfilter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"liveearth/infrastructure/config"
	"strings"
	"time"

	"github.com/cyanBone/dingtalk_robot"
	"github.com/cyanBone/dingtalk_robot/message"
	jsoniter "github.com/json-iterator/go"
	"github.com/patrickmn/go-cache"
)

type rpcWriter struct {
	service     string
	platName    string
	systemName  string
	serverTypes []string
	clusterName string
	report      *dingtalk_robot.Client
	cache       *cache.Cache
}

func newRPCWriter(service string, platName string, systemName string, clusterName string, serverTypes []string) (r *rpcWriter) {
	client, _ := dingtalk_robot.New(
		"https://oapi.dingtalk.com/robot/send?access_token=f7463495345f72deeb45f7c42b76dc03127d86d06b89591137242ddd4e7a347b",
		"SEC663d96395bf720fff39765aa88d2a09ba6eac1965e29a72aa4fb620fdfaf997e",
	)

	if config.C.Debug {
		systemName = systemName + "-develop"
	}

	return &rpcWriter{
		service:     service,
		platName:    platName,
		systemName:  systemName,
		clusterName: clusterName,
		serverTypes: serverTypes,
		report:      client,
		cache:       cache.New(10*time.Minute, 15*time.Minute),
	}
}

func (r *rpcWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	p[0] = byte('[')
	p = append(p, byte(']'))
	var buff bytes.Buffer
	if err := json.Compact(&buff, p); err != nil {
		err = fmt.Errorf("json.compact.err:%v", err)
		return 0, err
	}

	// todo://发送到远程日志收集器
	var ls []*Layout
	_ = jsoniter.Unmarshal(buff.Bytes(), &ls)
	for _, v := range ls {
		if v.Level == "e" {
			go r.reportDingDing(v)
		}
	}

	return len(p) - 1, nil
}

func (r *rpcWriter) Close() error {

	return nil
}

type Layout struct {
	Time    string `json:"time"`
	Content string `json:"content"`
	Level   string `json:"level"`
	Session string `json:"session"`
}

func (r *rpcWriter) reportDingDing(l *Layout) {

	_, found := r.cache.Get(r.systemName)
	if found {
		return
	}

	if !strings.Contains(l.Content, "panic") {
		return
	}

	markdownMessage := message.NewMarkdownMessage()
	markdownMessage.Title = "错误警告"
	markdownMessage.Text = "### 平台 \n" +
		"> **" + r.platName + "--" + r.systemName + "** \n" +
		"### 时间 \n" +
		"> **" + l.Time + "** \n" +
		"### session \n" +
		"> **" + l.Session + "** \n" +
		"### 问题 \n" +
		"> **" + l.Content + "** \n"

	markdownMessage.AtAll(true)

	_ = r.report.Send(markdownMessage)

	r.cache.Set(r.systemName, struct{}{}, cache.DefaultExpiration)
}
