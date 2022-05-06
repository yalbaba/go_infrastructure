package cron

import (
	"fmt"
	"go_infrastructure/pkg/iris"
	"math"
	"time"

	logger "github.com/sereiner/library/log"

	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"
)

type Handler func(ctx iris.Context) (err error)

type ICronTask interface {
	GetName() string
	ReduceRound(int)
	GetRound() int
	SetRound(int)
	GetExecuted() int
	AddExecuted()
	NextTime(time.Time) time.Time
	GetHandler() Handler
	Enable() bool
	SetDisable()
	GetTaskExecutionRecord() (string, error)
}

type Task struct {
	Name    string `json:"name"`
	Cron    string `json:"cron"`
	Disable bool   `json:"disable"`
	Handler Handler
}

type CronTask struct {
	*Task
	schedule cron.Schedule
	Executed int
	round    int
	result   string
	logger.ILogger
}

func NewCronTask(task *Task, l logger.ILogger) (r *CronTask) {

	r = &CronTask{
		Task:    task,
		ILogger: l,
	}

	s, err := cron.ParseStandard(task.Cron)
	if err != nil {
		panic(fmt.Errorf("%s的cron表达式(%s)配置有误", task.Name, task.Cron))
	}
	r.schedule = s

	return
}

func (m *CronTask) GetName() string {
	return m.Task.Name
}

func (m *CronTask) ReduceRound(v int) {
	m.round -= v
}

func (m *CronTask) GetRound() int {
	return m.round
}

func (m *CronTask) SetRound(v int) {
	m.round = v
}

func (m *CronTask) GetExecuted() int {
	return m.Executed
}

func (m *CronTask) AddExecuted() {
	if m.Executed >= math.MaxInt32 {
		m.Executed = 1
	} else {
		m.Executed++
	}
}

func (m *CronTask) GetHandler() Handler {

	return func(ctx iris.Context) (err error) {
		now := time.Now()
		ctx.Info("-----开始执行任务------", m.Name)
		err = m.Handler(ctx)
		if err != nil {
			ctx.Error("执行任务失败", m.Name, err, "耗时", time.Since(now).String())
			return
		}
		ctx.Info("执行任务成功", m.Name, "耗时", time.Since(now).String())
		return
	}

}

func (m *CronTask) NextTime(t time.Time) time.Time {
	return m.schedule.Next(t)
}

func (m *CronTask) Enable() bool {
	return !m.Disable
}

func (m *CronTask) SetDisable() {
	m.Disable = true
}

func (m *CronTask) GetTaskExecutionRecord() (string, error) {
	data := map[string]interface{}{
		"name":     m.Name,
		"cron":     m.Cron,
		"executed": m.Executed,
		"result":   m.result,
		"next":     m.NextTime(time.Now()).Format("20060102150405"),
		"last":     time.Now().Format("20060102150405"),
	}

	s, err := jsoniter.MarshalToString(data)
	if err != nil {
		return "", err
	}

	return s, nil
}
