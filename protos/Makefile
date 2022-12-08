.PHONY:myrpc

all: test


#push_stream:
#	protoc -I. --go_out=plugins=grpc:. ./push_stream/push_stream.proto

test_rpc_server:
	protoc -I. --go_out=. ./test_rpc_server.proto