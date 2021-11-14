package rpccli

import (
	"context"
	"fmt"
	"liveearth/infrastructure/component/rpccli/balancer"
	"liveearth/infrastructure/consts"
	"liveearth/infrastructure/protos/data_platform"
	"liveearth/infrastructure/protos/push_stream"
	"strings"
	"time"

	"liveearth/infrastructure/protos/geofence"
	"liveearth/infrastructure/protos/guide"
	"liveearth/infrastructure/protos/recommend"
	"liveearth/infrastructure/protos/wetoken"

	"google.golang.org/grpc/resolver"

	"liveearth/infrastructure/protos/comment"
	"liveearth/infrastructure/protos/stream_sync"

	"github.com/sereiner/library/concurrent/cmap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"liveearth/infrastructure/config"
	"liveearth/infrastructure/protos/content"
	"liveearth/infrastructure/protos/footprint"
	"liveearth/infrastructure/protos/im"
	"liveearth/infrastructure/protos/message_push"
	"liveearth/infrastructure/protos/search"
	"liveearth/infrastructure/protos/user"
)

type IComponentRpcClient interface {
	GetUserServiceClient() user.UserServiceClient
	GetIMServiceClient() im.IMServiceClient
	GetContentServiceClient() content.ContentServiceClient
	GetSearchServiceClient() search.SearchServiceClient
	GetMessagePushServiceClient() message_push.MessagePushServiceClient
	GetFootprintServiceClient() footprint.FootprintServiceClient
	GetCommentServiceClient() comment.CommentServiceClient
	GetStreamSyncServiceClient() stream_sync.StreamSyncServiceClient
	GetWeTokenServiceClient() wetoken.WeTokenServiceClient
	GetRecommendServiceClient() recommend.RecommendServiceClient
	GetGeofenceServiceClient() geofence.GeofenceServiceClient
	GetGuideServiceClient() guide.GuideServiceClient
	GetDataPlatformServiceClient() data_platform.DataPlatformServiceClient
	GetPushStreamServiceClient() push_stream.PushStreamServiceClient

	GetClient(names ...string) (r interface{}, err error)
	GetClientBy(name string) (r interface{}, err error)
	SaveClientObject(name string, f func(conf config.Service) (interface{}, error)) (bool, interface{}, error)
	Close() error
}

type StandardRpcClient struct {
	name   string
	rpcCli cmap.ConcurrentMap
	ccMap  map[string]*grpc.ClientConn
}

func NewStandardRpcClient(name ...string) IComponentRpcClient {
	if len(name) > 0 {
		return &StandardRpcClient{
			name:   name[0],
			rpcCli: cmap.New(2),
			ccMap:  map[string]*grpc.ClientConn{},
		}
	}
	return &StandardRpcClient{
		name:   "default",
		rpcCli: cmap.New(2),
		ccMap:  map[string]*grpc.ClientConn{},
	}
}

func (s *StandardRpcClient) GetUserServiceClient() user.UserServiceClient {

	r, err := s.GetClient(consts.UserCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(user.UserServiceClient)
	if !ok {
		panic("UserServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetWeTokenServiceClient() wetoken.WeTokenServiceClient {

	r, err := s.GetClient(consts.Wetoken)
	if err != nil {
		panic(err)
	}
	v, ok := r.(wetoken.WeTokenServiceClient)
	if !ok {
		panic("WeTokenServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetRecommendServiceClient() recommend.RecommendServiceClient {

	r, err := s.GetClient(consts.Recommend)
	if err != nil {
		panic(err)
	}
	v, ok := r.(recommend.RecommendServiceClient)
	if !ok {
		panic("RecommendServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetIMServiceClient() im.IMServiceClient {
	r, err := s.GetClient(consts.ImCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(im.IMServiceClient)
	if !ok {
		panic("IMServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetContentServiceClient() content.ContentServiceClient {
	r, err := s.GetClient(consts.ContentCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(content.ContentServiceClient)
	if !ok {
		panic("ContentServiceClient not found")
	}

	return v
}

func (s *StandardRpcClient) GetSearchServiceClient() search.SearchServiceClient {
	r, err := s.GetClient(consts.SearchCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(search.SearchServiceClient)
	if !ok {
		panic("SearchServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetMessagePushServiceClient() message_push.MessagePushServiceClient {

	r, err := s.GetClient(consts.MessagePushCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(message_push.MessagePushServiceClient)
	if !ok {
		panic("MessagePushServiceClient not found")
	}

	return v
}

func (s *StandardRpcClient) GetFootprintServiceClient() footprint.FootprintServiceClient {
	r, err := s.GetClient(consts.FootprintCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(footprint.FootprintServiceClient)
	if !ok {
		panic("FootprintServiceClient not found")
	}

	return v
}

func (s *StandardRpcClient) GetCommentServiceClient() comment.CommentServiceClient {
	r, err := s.GetClient(consts.CommentCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(comment.CommentServiceClient)
	if !ok {
		panic("CommentServiceClient not found")
	}

	return v
}

func (s *StandardRpcClient) GetStreamSyncServiceClient() stream_sync.StreamSyncServiceClient {
	r, err := s.GetClient(consts.StreamSyncCliName)
	if err != nil {
		panic(err)
	}
	v, ok := r.(stream_sync.StreamSyncServiceClient)
	if !ok {
		panic("StreamSyncServiceClient not found")
	}

	return v
}

func (s *StandardRpcClient) GetGeofenceServiceClient() geofence.GeofenceServiceClient {

	r, err := s.GetClient(consts.Geofence)
	if err != nil {
		panic(err)
	}
	v, ok := r.(geofence.GeofenceServiceClient)
	if !ok {
		panic("GeofenceServiceClient not found")
	}

	return v
}

func (s *StandardRpcClient) GetGuideServiceClient() guide.GuideServiceClient {
	r, err := s.GetClient(consts.Guide)
	if err != nil {
		panic(err)
	}
	v, ok := r.(guide.GuideServiceClient)
	if !ok {
		panic("GuideServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetDataPlatformServiceClient() data_platform.DataPlatformServiceClient {
	r, err := s.GetClient(consts.DataPlatform)
	if err != nil {
		panic(err)
	}
	v, ok := r.(data_platform.DataPlatformServiceClient)
	if !ok {
		panic("DataPlatformServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetPushStreamServiceClient() push_stream.PushStreamServiceClient {
	var (
		r   interface{}
		err error
	)
	r, err = s.GetClient(consts.PushStream)
	if err != nil {
		panic(err)
	}

	v, ok := r.(push_stream.PushStreamServiceClient)
	if !ok {
		panic("PushStreamServiceClient not found")
	}
	return v
}

func (s *StandardRpcClient) GetClient(names ...string) (r interface{}, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetClientBy(name)
}

func (s *StandardRpcClient) GetClientBy(name string) (r interface{}, err error) {

	_, c, err := s.SaveClientObject(name, func(conf config.Service) (interface{}, error) {

		re := balancer.NewResolver("live", strings.Split(conf.Addr, ","))
		resolver.Register(re)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		cc, err := grpc.DialContext(
			ctx,
			re.Scheme()+"://authority/",
			grpc.WithBalancerName(roundrobin.Name),
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			return nil, err
		}
		s.ccMap[name] = cc
		var r interface{}
		switch name {
		case consts.UserCliName:
			r = user.NewUserServiceClient(cc)
		case consts.ContentCliName:
			r = content.NewContentServiceClient(cc)
		case consts.ImCliName:
			r = im.NewIMServiceClient(cc)
		case consts.SearchCliName:
			r = search.NewSearchServiceClient(cc)
		case consts.MessagePushCliName:
			r = message_push.NewMessagePushServiceClient(cc)
		case consts.FootprintCliName:
			r = footprint.NewFootprintServiceClient(cc)
		case consts.CommentCliName:
			r = comment.NewCommentServiceClient(cc)
		case consts.StreamSyncCliName:
			r = stream_sync.NewStreamSyncServiceClient(cc)
		case consts.Wetoken:
			r = wetoken.NewWeTokenServiceClient(cc)
		case consts.Recommend:
			r = recommend.NewRecommendServiceClient(cc)
		case consts.Geofence:
			r = geofence.NewGeofenceServiceClient(cc)
		case consts.Guide:
			r = guide.NewGuideServiceClient(cc)
		case consts.DataPlatform:
			r = data_platform.NewDataPlatformServiceClient(cc)
		case consts.PushStream:
			r = push_stream.NewPushStreamServiceClient(cc)
		}

		return r, nil
	})

	return c, err
}

func (s *StandardRpcClient) SaveClientObject(name string, f func(conf config.Service) (interface{}, error)) (bool, interface{}, error) {

	key := fmt.Sprintf("%s/%s", "rpc", name)

	ok, ch, err := s.rpcCli.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {

		clientConf, ok := config.C.Service[name]
		if !ok {
			return nil, fmt.Errorf("grpc客户端配置不存在 name:%s", name)
		}

		return f(clientConf)
	})
	if err != nil {
		err = fmt.Errorf("创建rpc客户端失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch, err
}

func (s *StandardRpcClient) Close() error {
	for _, v := range s.ccMap {
		return v.Close()
	}

	return nil
}
