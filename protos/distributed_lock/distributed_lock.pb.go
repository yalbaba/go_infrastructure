// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distributed_lock/distributed_lock.proto

package distributed_lock

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LockReq struct {
	StreamName           string   `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockReq) Reset()         { *m = LockReq{} }
func (m *LockReq) String() string { return proto.CompactTextString(m) }
func (*LockReq) ProtoMessage()    {}
func (*LockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b13054010f28a548, []int{0}
}

func (m *LockReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockReq.Unmarshal(m, b)
}
func (m *LockReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockReq.Marshal(b, m, deterministic)
}
func (m *LockReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockReq.Merge(m, src)
}
func (m *LockReq) XXX_Size() int {
	return xxx_messageInfo_LockReq.Size(m)
}
func (m *LockReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LockReq.DiscardUnknown(m)
}

var xxx_messageInfo_LockReq proto.InternalMessageInfo

func (m *LockReq) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

type LockResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	CheckValue           string   `protobuf:"bytes,3,opt,name=check_value,json=checkValue,proto3" json:"check_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockResp) Reset()         { *m = LockResp{} }
func (m *LockResp) String() string { return proto.CompactTextString(m) }
func (*LockResp) ProtoMessage()    {}
func (*LockResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b13054010f28a548, []int{1}
}

func (m *LockResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockResp.Unmarshal(m, b)
}
func (m *LockResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockResp.Marshal(b, m, deterministic)
}
func (m *LockResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockResp.Merge(m, src)
}
func (m *LockResp) XXX_Size() int {
	return xxx_messageInfo_LockResp.Size(m)
}
func (m *LockResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LockResp.DiscardUnknown(m)
}

var xxx_messageInfo_LockResp proto.InternalMessageInfo

func (m *LockResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LockResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LockResp) GetCheckValue() string {
	if m != nil {
		return m.CheckValue
	}
	return ""
}

func init() {
	proto.RegisterType((*LockReq)(nil), "LockReq")
	proto.RegisterType((*LockResp)(nil), "LockResp")
}

func init() {
	proto.RegisterFile("distributed_lock/distributed_lock.proto", fileDescriptor_b13054010f28a548)
}

var fileDescriptor_b13054010f28a548 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x89, 0xcf, 0xc9, 0x4f, 0xce, 0xd6, 0x47, 0x17, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe2, 0x62, 0xf7, 0xc9, 0x4f, 0xce, 0x0e, 0x4a, 0x2d,
	0x14, 0x92, 0xe7, 0xe2, 0x2e, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x8d, 0xcf, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x82, 0x08, 0xf9, 0x25, 0xe6, 0xa6, 0x2a, 0x45, 0x72,
	0x71, 0x40, 0xd4, 0x16, 0x17, 0x08, 0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0x40, 0x54, 0xb1, 0x06,
	0x81, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60,
	0xcd, 0x30, 0x2e, 0xc8, 0xe8, 0xe4, 0x8c, 0xd4, 0xe4, 0xec, 0xf8, 0xb2, 0xc4, 0x9c, 0xd2, 0x54,
	0x09, 0x66, 0x88, 0xd1, 0x60, 0xa1, 0x30, 0x90, 0x88, 0x51, 0x08, 0x97, 0x98, 0x0b, 0xc2, 0x81,
	0x20, 0x5b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xa4, 0xb9, 0x58, 0x40, 0x5c, 0x21,
	0x0e, 0x3d, 0xa8, 0x3b, 0xa5, 0x38, 0xf5, 0xe0, 0xae, 0x90, 0xe5, 0x62, 0x0b, 0xcd, 0xc3, 0x29,
	0xed, 0x24, 0x12, 0x25, 0xa4, 0x87, 0xe1, 0xf1, 0x24, 0x36, 0xb0, 0xcf, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0xe8, 0x71, 0xf9, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DistributedLockServiceClient is the client API for DistributedLockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistributedLockServiceClient interface {
	Lock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockResp, error)
	UnLock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockResp, error)
}

type distributedLockServiceClient struct {
	cc *grpc.ClientConn
}

func NewDistributedLockServiceClient(cc *grpc.ClientConn) DistributedLockServiceClient {
	return &distributedLockServiceClient{cc}
}

func (c *distributedLockServiceClient) Lock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockResp, error) {
	out := new(LockResp)
	err := c.cc.Invoke(ctx, "/DistributedLockService/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedLockServiceClient) UnLock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockResp, error) {
	out := new(LockResp)
	err := c.cc.Invoke(ctx, "/DistributedLockService/UnLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributedLockServiceServer is the server API for DistributedLockService service.
type DistributedLockServiceServer interface {
	Lock(context.Context, *LockReq) (*LockResp, error)
	UnLock(context.Context, *LockReq) (*LockResp, error)
}

// UnimplementedDistributedLockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDistributedLockServiceServer struct {
}

func (*UnimplementedDistributedLockServiceServer) Lock(ctx context.Context, req *LockReq) (*LockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (*UnimplementedDistributedLockServiceServer) UnLock(ctx context.Context, req *LockReq) (*LockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnLock not implemented")
}

func RegisterDistributedLockServiceServer(s *grpc.Server, srv DistributedLockServiceServer) {
	s.RegisterService(&_DistributedLockService_serviceDesc, srv)
}

func _DistributedLockService_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedLockServiceServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedLockService/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedLockServiceServer).Lock(ctx, req.(*LockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedLockService_UnLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedLockServiceServer).UnLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributedLockService/UnLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedLockServiceServer).UnLock(ctx, req.(*LockReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistributedLockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DistributedLockService",
	HandlerType: (*DistributedLockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lock",
			Handler:    _DistributedLockService_Lock_Handler,
		},
		{
			MethodName: "UnLock",
			Handler:    _DistributedLockService_UnLock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributed_lock/distributed_lock.proto",
}