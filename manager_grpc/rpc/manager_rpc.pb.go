// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager_rpc.proto

package manager_rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type UpdateCounterRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCounterRequest) Reset()         { *m = UpdateCounterRequest{} }
func (m *UpdateCounterRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCounterRequest) ProtoMessage()    {}
func (*UpdateCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3142b9f989b401f, []int{0}
}

func (m *UpdateCounterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCounterRequest.Unmarshal(m, b)
}
func (m *UpdateCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCounterRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCounterRequest.Merge(m, src)
}
func (m *UpdateCounterRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCounterRequest.Size(m)
}
func (m *UpdateCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCounterRequest proto.InternalMessageInfo

func (m *UpdateCounterRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateCounterRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type NewTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTokenResponse) Reset()         { *m = NewTokenResponse{} }
func (m *NewTokenResponse) String() string { return proto.CompactTextString(m) }
func (*NewTokenResponse) ProtoMessage()    {}
func (*NewTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3142b9f989b401f, []int{1}
}

func (m *NewTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTokenResponse.Unmarshal(m, b)
}
func (m *NewTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTokenResponse.Marshal(b, m, deterministic)
}
func (m *NewTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTokenResponse.Merge(m, src)
}
func (m *NewTokenResponse) XXX_Size() int {
	return xxx_messageInfo_NewTokenResponse.Size(m)
}
func (m *NewTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewTokenResponse proto.InternalMessageInfo

func (m *NewTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateCounterResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Counter              int64    `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCounterResponse) Reset()         { *m = UpdateCounterResponse{} }
func (m *UpdateCounterResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCounterResponse) ProtoMessage()    {}
func (*UpdateCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3142b9f989b401f, []int{2}
}

func (m *UpdateCounterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCounterResponse.Unmarshal(m, b)
}
func (m *UpdateCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCounterResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCounterResponse.Merge(m, src)
}
func (m *UpdateCounterResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCounterResponse.Size(m)
}
func (m *UpdateCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCounterResponse proto.InternalMessageInfo

func (m *UpdateCounterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateCounterResponse) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*UpdateCounterRequest)(nil), "manager_rpc.UpdateCounterRequest")
	proto.RegisterType((*NewTokenResponse)(nil), "manager_rpc.NewTokenResponse")
	proto.RegisterType((*UpdateCounterResponse)(nil), "manager_rpc.UpdateCounterResponse")
}

func init() {
	proto.RegisterFile("manager_rpc.proto", fileDescriptor_e3142b9f989b401f)
}

var fileDescriptor_e3142b9f989b401f = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x2d, 0x8a, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x92, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x25, 0x95, 0xa6, 0xe9,
	0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x54, 0x2a, 0x39, 0x70, 0x89, 0x84, 0x16, 0xa4, 0x24, 0x96,
	0xa4, 0x3a, 0xe7, 0x97, 0xe6, 0x95, 0xa4, 0x16, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x89, 0x70, 0xb1, 0x96, 0xe4, 0x67, 0xa7, 0xe6, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41,
	0x38, 0x42, 0x42, 0x5c, 0x2c, 0x05, 0x89, 0x25, 0x19, 0x12, 0x4c, 0x60, 0x41, 0x30, 0x5b, 0x49,
	0x83, 0x4b, 0xc0, 0x2f, 0xb5, 0x3c, 0x04, 0x24, 0x1f, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x8a, 0x5d, 0xb7, 0x92, 0x37, 0x97, 0x28, 0x9a, 0x5d, 0x50, 0xe5, 0x12, 0x5c, 0xec, 0xc5, 0xa5,
	0xc9, 0xc9, 0xa9, 0xc5, 0xc5, 0x60, 0x0d, 0x1c, 0x41, 0x30, 0x2e, 0x48, 0x26, 0x19, 0xa2, 0x18,
	0x6c, 0x27, 0x73, 0x10, 0x8c, 0x6b, 0xb4, 0x9a, 0x91, 0x8b, 0x0f, 0x6a, 0x4e, 0x70, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0xaa, 0x90, 0x23, 0x17, 0x07, 0xcc, 0x25, 0x42, 0x62, 0x7a, 0x10, 0x5f, 0xeb,
	0xc1, 0x7c, 0xad, 0xe7, 0x0a, 0xf2, 0xb5, 0x94, 0xac, 0x1e, 0x72, 0x68, 0x61, 0x38, 0x3c, 0x8c,
	0x8b, 0x17, 0xc5, 0x89, 0x42, 0x8a, 0x28, 0xea, 0xb1, 0x05, 0x95, 0x94, 0x12, 0x3e, 0x25, 0x10,
	0x73, 0x93, 0xd8, 0xc0, 0xce, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0x8a, 0xe6, 0x1e,
	0xac, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterServiceClient interface {
	NewToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NewTokenResponse, error)
	UpdateCounter(ctx context.Context, in *UpdateCounterRequest, opts ...grpc.CallOption) (*UpdateCounterResponse, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) NewToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NewTokenResponse, error) {
	out := new(NewTokenResponse)
	err := c.cc.Invoke(ctx, "/manager_rpc.CounterService/NewToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) UpdateCounter(ctx context.Context, in *UpdateCounterRequest, opts ...grpc.CallOption) (*UpdateCounterResponse, error) {
	out := new(UpdateCounterResponse)
	err := c.cc.Invoke(ctx, "/manager_rpc.CounterService/UpdateCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
type CounterServiceServer interface {
	NewToken(context.Context, *empty.Empty) (*NewTokenResponse, error)
	UpdateCounter(context.Context, *UpdateCounterRequest) (*UpdateCounterResponse, error)
}

// UnimplementedCounterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (*UnimplementedCounterServiceServer) NewToken(ctx context.Context, req *empty.Empty) (*NewTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewToken not implemented")
}
func (*UnimplementedCounterServiceServer) UpdateCounter(ctx context.Context, req *UpdateCounterRequest) (*UpdateCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCounter not implemented")
}

func RegisterCounterServiceServer(s *grpc.Server, srv CounterServiceServer) {
	s.RegisterService(&_CounterService_serviceDesc, srv)
}

func _CounterService_NewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).NewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_rpc.CounterService/NewToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).NewToken(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_UpdateCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).UpdateCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_rpc.CounterService/UpdateCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).UpdateCounter(ctx, req.(*UpdateCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CounterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manager_rpc.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewToken",
			Handler:    _CounterService_NewToken_Handler,
		},
		{
			MethodName: "UpdateCounter",
			Handler:    _CounterService_UpdateCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager_rpc.proto",
}