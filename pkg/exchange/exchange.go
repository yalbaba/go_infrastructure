package exchange

import (
	"fmt"
	"sync"
)

//WSExchange web socket exchange
var WSExchange = NewExchange()

//Exchange 数据交换中心
type Exchange struct {
	uuid      map[string]func(buff []byte) error
	lRelation map[string]string
	rRelation map[string]string
	lock      sync.RWMutex
}

//NewExchange 构建数据交换中心
func NewExchange() *Exchange {
	return &Exchange{
		uuid:      make(map[string]func(buff []byte) error),
		lRelation: make(map[string]string),
		rRelation: make(map[string]string),
	}
}

//Subscribe 订阅消息通知
func (e *Exchange) Subscribe(uuid string, f func([]byte) error) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.uuid[uuid]; !ok {
		e.uuid[uuid] = f
		e.rRelation[uuid] = uuid
		return nil
	}
	return fmt.Errorf("重复的消息订阅：%s", uuid)
}

//Unsubscribe 取消订阅
func (e *Exchange) Unsubscribe(uuid string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.uuid[uuid]; ok {
		delete(e.uuid, uuid)
	}
	if _, ok := e.rRelation[uuid]; ok {
		delete(e.rRelation, uuid)
	}
	if v, ok := e.lRelation[uuid]; ok {
		delete(e.rRelation, v)
		delete(e.lRelation, uuid)
	}
}

//Relate 关联别名
func (e *Exchange) Relate(uuid string, name string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.lRelation[uuid] = name
	e.rRelation[name] = uuid
}

//Clear 清除所有订阅者
func (e *Exchange) Clear() {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.uuid = make(map[string]func(buff []byte) error)
	e.lRelation = make(map[string]string)
	e.rRelation = make(map[string]string)
}

//Notify 消息通知
func (e *Exchange) Notify(name string, buff []byte) error {
	e.lock.RLock()
	defer e.lock.RUnlock()
	uuid := e.rRelation[name]
	if v, ok := e.uuid[uuid]; ok {
		return v(buff)
	}
	return fmt.Errorf("未找到消息订阅者:%s %v", name, e.uuid)
}

//Broadcast 发送广播消息
func (e *Exchange) Broadcast(buff []byte) error {
	e.lock.RLock()
	defer e.lock.RUnlock()
	for _, f := range e.uuid {
		if err := f(buff); err != nil {
			return err
		}
	}
	return nil
}
