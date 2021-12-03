.PHONY:myrpc

all: push_stream


push_stream:
	protoc -I. --go_out=plugins=grpc:. ./push_stream/push_stream.proto