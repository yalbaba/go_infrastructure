.PHONY:content
.PHONY:im
.PHONY:message_push
.PHONY:search
.PHONY:user
.PHONY:footprint
.PHONY:comment
.PHONY:wetoken
.PHONY:recommend
.PHONY:geofence
.PHONY:guide
.PHONY:data_platform
.PHONY:push_stream

all: content im message_push search user footprint comment recommend push_stream


content:
	protoc -I. --go_out=plugins=grpc:. ./content/content.proto

im:
	protoc -I. --go_out=plugins=grpc:. ./im/im.proto

message_push:
	protoc -I. --go_out=plugins=grpc:. ./message_push/message_push.proto

search:
	protoc -I. --go_out=plugins=grpc:. ./search/search.proto


user:
	protoc -I. --go_out=plugins=grpc:. ./user/user.proto

footprint:
	protoc -I. --go_out=plugins=grpc:. ./footprint/footprint.proto

comment:
	protoc -I. --go_out=plugins=grpc:. ./comment/comment.proto

stream_sync:
	protoc -I. --go_out=plugins=grpc:. ./stream_sync/stream_sync.proto

wetoken:
	protoc -I. --go_out=plugins=grpc:. ./wetoken/wetoken.proto

recommend:
	protoc -I. --go_out=plugins=grpc:. ./recommend/recommend.proto

geofence:
	protoc -I. --go_out=plugins=grpc:. ./geofence/geofence.proto

guide:
	protoc -I. --go_out=plugins=grpc:. ./guide/guide.proto

data_platform:
	protoc -I. --go_out=plugins=grpc:. ./data_platform/data_platform.proto

push_stream:
	protoc -I. --go_out=plugins=grpc:. ./push_stream/push_stream.proto