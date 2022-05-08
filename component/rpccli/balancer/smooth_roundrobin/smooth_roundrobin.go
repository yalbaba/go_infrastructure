package smooth_roundrobin

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/yalbaba/go_infrastructure/component/rpccli/balancer"
	"github.com/yalbaba/go_infrastructure/utils"

	"github.com/coreos/etcd/mvcc/mvccpb"
	jsoniter "github.com/json-iterator/go"
	"github.com/ozonru/etcd/clientv3"
	logger "github.com/sereiner/library/log"
	"google.golang.org/grpc"
)

/*
平滑轮询负载均衡算法：
1、准备好初始权重列表
2、准备好轮询过程中的当前权重列表（初始权重为0）
3、每次轮询时，当前权重列表的每个权重加上自己的初始权重
4、查找当前权重列表中权重最大的元素为目标元素返回
5、最后将最大的元素的权重-总的初始权重
*/

const (
	PushStream = "push_stream"
)

type SmoothRoundRobin struct {
	cli  *clientv3.Client
	lock *sync.Mutex
	*logger.Logger
	hasStarted bool

	base_servers map[string][]Item //key：服务的名称，value：基础的服务权重列表
	cur_servers  map[string][]Item //key：服务的名称，value：开始轮询后的权重列表
}

type Item struct {
	name   string //服务完整名字
	weight int    //权重
	addr   string //服务ip地址
}

func NewSmoothRoundRobin(cli *clientv3.Client, l *logger.Logger) {
	balancer.B = &SmoothRoundRobin{
		cli:    cli,
		lock:   &sync.Mutex{},
		Logger: l,
	}
}

func (b *SmoothRoundRobin) GetConn(target string) (cc *grpc.ClientConn, err error) {

	if b.hasStarted && len(b.base_servers[target]) < 2 {
		err = utils.DingTalk(target)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	if _, ok := b.base_servers[target]; !ok {
		err := b.Start(target)
		if err != nil {
			return nil, err
		}
	}

	return b.getConn(target)
}

func (b *SmoothRoundRobin) Start(target string) error {

	resp, err := b.cli.Get(context.Background(), target, clientv3.WithPrefix())
	if err != nil {
		return fmt.Errorf("获取target服务失败,err:%v,target:%s", err, target)
	}

	b.base_servers = make(map[string][]Item)
	b.base_servers[target] = make([]Item, len(resp.Kvs))
	b.cur_servers = make(map[string][]Item)
	b.cur_servers[target] = make([]Item, len(resp.Kvs))

	switch target {
	case PushStream:
		//初始化权重列表
		err = b.initPushStreamServers(resp)
		if err != nil {
			return err
		}
		//开启协程监控
		b.watchPushStreamServers(resp)
	}

	b.hasStarted = true

	return nil
}

//更新或者添加服务节点的操作
func (b *SmoothRoundRobin) update(target, serverName string, rest int, addr string) {

	//如果权重为0，下线服务or不注册服务
	if rest <= 0 {
		b.delete(target, serverName)
		return
	}

	var exist bool
	for i := 0; i < len(b.base_servers[target]); i++ {
		if b.base_servers[target][i].name == serverName {
			exist = true
			b.base_servers[target][i].weight = rest
			break
		}
	}

	if !exist { //不存在，新注册服务
		if rest <= 0 {
			return
		}
		b.base_servers[target] = append(b.base_servers[target], Item{
			name:   serverName,
			weight: rest,
			addr:   addr,
		})
		b.cur_servers[target] = append(b.cur_servers[target], Item{
			name: serverName,
			addr: addr,
		})
	}
}

//删除服务或者服务下线了的操作
func (b *SmoothRoundRobin) delete(target, serverName string) {
	var newServers []Item
	var newCurServers []Item

	for i := 0; i < len(b.base_servers[target]); i++ {
		if b.base_servers[target][i].name != serverName {
			newServers = append(newServers, b.base_servers[target][i])
		}
	}

	for j := 0; j < len(b.cur_servers[target]); j++ {
		if b.cur_servers[target][j].name != serverName {
			newCurServers = append(newCurServers, b.cur_servers[target][j])
		}
	}

	b.base_servers[target] = newServers
	b.cur_servers[target] = newCurServers
}

//根据平滑轮询获取连接
func (b *SmoothRoundRobin) getConn(target string) (*grpc.ClientConn, error) {

	//计算总权重
	var count int
	for _, v := range b.base_servers[target] {
		count += v.weight
	}

	//将cur_servers列表加上对应的初始权重
	for i := range b.base_servers[target] {
		b.cur_servers[target][i].weight += b.base_servers[target][i].weight
	}

	//在cur_servers列表中找到权重最大的
	var (
		temp int
		idx  int
	)
	for i := range b.cur_servers[target] {
		if b.cur_servers[target][i].weight > temp {
			temp = b.cur_servers[target][i].weight
			idx = i
		}
	}

	//将最大的对象的权重-总权重
	b.cur_servers[target][idx].weight -= count

	//获取rpc连接返回
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx,
		b.cur_servers[target][idx].addr,
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("获取rpc连接失败,err:%v,addr:%s", err, b.cur_servers[target][idx].addr)
	}

	return conn, nil
}

/*
推流服务的实现
*/
type PushServiceInfo struct {
	Using int    `json:"using"` //正在使用的线程数
	Rest  int    `json:"rest"`  //剩下的线程数
	Addr  string `json:"addr"`  //服务的地址：ip+port
}

//初始化推流服务的权重列表
func (b *SmoothRoundRobin) initPushStreamServers(resp *clientv3.GetResponse) error {
	for i, kv := range resp.Kvs {
		var serviceInfo PushServiceInfo
		if err := json.Unmarshal(kv.Value, &serviceInfo); err != nil {
			return fmt.Errorf("注册中心服务信息转结构体失败,err:%v,source:%s", err, kv.Value)
		}

		b.base_servers[PushStream][i] = Item{name: string(kv.Key), weight: serviceInfo.Rest, addr: serviceInfo.Addr}
		b.cur_servers[PushStream][i] = Item{name: string(kv.Key), addr: serviceInfo.Addr}
	}

	return nil
}

//推流服务监控
func (b *SmoothRoundRobin) watchPushStreamServers(resp *clientv3.GetResponse) {
	//开启协程监控
	opt := []clientv3.OpOption{clientv3.WithRev(resp.Header.Revision + 1), clientv3.WithPrefix(), clientv3.WithPrevKV()}
	watchChan := b.cli.Watch(context.TODO(), PushStream, opt...)

	go func() {
		for {
			wrsp := <-watchChan
			if wrsp.Events == nil { //管道已关闭
				panic("etcd has closed!")
			}
			if err := wrsp.Err(); err != nil {
				b.Errorf("watcher 事件通知错误,err:%v", err)
				continue
			}
			for _, event := range wrsp.Events {
				switch event.Type {
				case mvccpb.PUT:
					b.Warn("PUT key::", string(event.Kv.Key), "value::", string(event.Kv.Value))
					var data PushServiceInfo
					err := jsoniter.UnmarshalFromString(string(event.Kv.Value), &data)
					if err != nil {
						panic(fmt.Errorf("服务信息格式错误,err:%v,value:%s", err, string(event.Kv.Value)))
					}
					b.update(PushStream, string(event.Kv.Key), data.Rest, data.Addr)
					b.Warn("cur: ", b.cur_servers)
					b.Warn("base: ", b.base_servers)
				case mvccpb.DELETE:
					b.Warn("DELETE key::", string(event.Kv.Key), "value::", string(event.Kv.Value))
					b.delete(PushStream, string(event.Kv.Key))
					b.Warn("cur: ", b.cur_servers)
					b.Warn("base: ", b.base_servers)
				}

			}
		}
	}()
}
