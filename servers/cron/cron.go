package cron

import (
	"go_infrastructure/component"
	"go_infrastructure/consts"
	"go_infrastructure/pkg/iris"
	"go_infrastructure/pkg/iris/context"
	"go_infrastructure/servers"
	"net/http/httptest"
	"sync"

	"time"

	"github.com/sereiner/library/concurrent/cmap"
	"github.com/sereiner/library/utility"
)

//Processor 任务处理程序
type CronServer struct {
	lock        sync.Mutex
	once        sync.Once
	done        bool
	closeChan   chan struct{}
	length      int
	index       int
	span        time.Duration
	slots       []cmap.ConcurrentMap
	startTime   time.Time
	isPause     bool
	c           component.Container
	ContextPool *context.Pool
}

func (s *CronServer) GetServerType() consts.ServerType {
	return consts.CronServer
}

// NewProcessor 创建processor
func NewProcessor(c component.Container) (p *CronServer) {

	p = &CronServer{
		closeChan: make(chan struct{}),
		span:      time.Second,
		length:    60,
		startTime: time.Now(),
		c:         c,
	}

	p.ContextPool = context.New(func() context.Context {
		return context.NewContext(&iris.Application{})
	})

	p.slots = make([]cmap.ConcurrentMap, p.length, p.length)
	for i := 0; i < p.length; i++ {
		p.slots[i] = cmap.New(2)
	}

	return p
}

func (s *CronServer) Start() error {

	s.c.Debug("开始启动 CRON 服务器...")

	go func() {
	START:
		for {
			select {
			case <-s.closeChan:
				break START
			case <-time.After(s.span):
				s.execute()
			}
		}
		return
	}()

	s.c.Debug("CRON 服务器启动成功")

	return nil
}

func (s *CronServer) execute() {

	s.startTime = time.Now()

	s.lock.Lock()
	defer s.lock.Unlock()

	s.index = (s.index + 1) % s.length
	current := s.slots[s.index]
	current.RemoveIterCb(func(k string, value interface{}) bool {
		task := value.(ICronTask)
		task.ReduceRound(1)
		if task.GetRound() < 0 {
			go s.handle(task)
			return true
		}
		return false
	})

}
func (s *CronServer) handle(task ICronTask) error {

	if s.done || !task.Enable() {
		return nil
	}

	if !s.isPause {
		task.AddExecuted()
		handler := task.GetHandler()
		if handler != nil {
			ctx := s.ContextPool.Acquire(httptest.NewRecorder(), nil)
			_ = handler(ctx)
			s.ContextPool.Release(ctx)
		}
	}

	s.RegisterService(task)

	return nil

}
func (s *CronServer) Remove(name string) {

	if name == "" {
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	for _, slot := range s.slots {
		slot.RemoveIterCb(func(k string, value interface{}) bool {
			task := value.(ICronTask)
			task.SetDisable()
			return task.GetName() == name
		})
	}

}

// Add 添加任务
func (s *CronServer) RegisterService(sc ...interface{}) {

	for _, v := range sc {

		s.lock.Lock()

		task, ok := v.(ICronTask)
		if !ok {
			s.lock.Unlock()
			s.c.Error("task type error")
		}

		if s.done {
			s.lock.Unlock()
			return
		}

		now := time.Now()
		nextTime := task.NextTime(now)
		if nextTime.Sub(now) < 0 {
			s.lock.Unlock()
			s.c.Error("next time less than now.1")
			return
		}

		offset, round := s.getOffset(now, nextTime)
		if offset < 0 || round < 0 {
			s.lock.Unlock()
			s.c.Error("next time less than now.2")
			return
		}

		task.SetRound(round)
		s.slots[offset].Set(utility.GetGUID(), task)

		s.lock.Unlock()
	}

	return
}

func (s *CronServer) getOffset(now time.Time, next time.Time) (pos int, circle int) {

	// 计算剩余时间
	d := next.Sub(now)
	delaySeconds := int(d/1e9) + 1
	intervalSeconds := int(s.span.Seconds())
	circle = delaySeconds / intervalSeconds / s.length
	pos = (s.index + delaySeconds/intervalSeconds) % s.length
	return
}

// Pause 暂停所有任务
func (s *CronServer) Pause() error {
	s.isPause = true
	return nil
}

// Resume 恢复所有任务
func (s *CronServer) Resume() error {
	s.isPause = false
	return nil
}

// Close 退出
func (s *CronServer) Close() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.done = true
	s.once.Do(func() {
		close(s.closeChan)
	})

	return nil
}

type cronServerAdapter struct {
}

func (h *cronServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewProcessor(c)
}

func init() {
	servers.Register(consts.CronServer, &cronServerAdapter{})
}
