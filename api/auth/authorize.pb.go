// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/auth/authorize.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AuthorizeRequest struct {
	Token                string               `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthorizeRequest) Reset()         { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()    {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8db002294e15fb6c, []int{0}
}

func (m *AuthorizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeRequest.Unmarshal(m, b)
}
func (m *AuthorizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeRequest.Merge(m, src)
}
func (m *AuthorizeRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeRequest.Size(m)
}
func (m *AuthorizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeRequest proto.InternalMessageInfo

func (m *AuthorizeRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthorizeRequest) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

type AuthorizeResponse struct {
	Authorized           bool                 `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	Token                string               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthorizeResponse) Reset()         { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()    {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8db002294e15fb6c, []int{1}
}

func (m *AuthorizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeResponse.Unmarshal(m, b)
}
func (m *AuthorizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeResponse.Merge(m, src)
}
func (m *AuthorizeResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizeResponse.Size(m)
}
func (m *AuthorizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeResponse proto.InternalMessageInfo

func (m *AuthorizeResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthorizeResponse) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

func (m *AuthorizeResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthorizeRequest)(nil), "auth.AuthorizeRequest")
	proto.RegisterType((*AuthorizeResponse)(nil), "auth.AuthorizeResponse")
}

func init() { proto.RegisterFile("api/auth/authorize.proto", fileDescriptor_8db002294e15fb6c) }

var fileDescriptor_8db002294e15fb6c = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0x13, 0xf9, 0x45, 0x99, 0x55, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x2c, 0x20, 0x01, 0x29, 0xf9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0xb0, 0x58,
	0x52, 0x69, 0x9a, 0x7e, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x01, 0x44, 0x99, 0x52,
	0x12, 0x97, 0x80, 0x23, 0x4c, 0x67, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08, 0x17,
	0x6b, 0x49, 0x7e, 0x76, 0x6a, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0x64,
	0xc1, 0xc5, 0x99, 0x5a, 0x51, 0x90, 0x59, 0x94, 0x9a, 0xe2, 0x58, 0x22, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x31, 0x5e, 0x0f, 0x66, 0xbc, 0x5e, 0x08, 0xcc, 0xf8, 0x20, 0x84,
	0x62, 0xa5, 0x66, 0x46, 0x2e, 0x41, 0x24, 0x4b, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xe4,
	0xb8, 0xb8, 0xe0, 0x6e, 0x4e, 0x01, 0x5b, 0xc5, 0x11, 0x84, 0x24, 0x42, 0xbe, 0x7d, 0x08, 0xf7,
	0x33, 0x23, 0xb9, 0xdf, 0x28, 0x00, 0xc9, 0xa7, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42,
	0x36, 0x5c, 0x9c, 0x70, 0x31, 0x21, 0x31, 0x3d, 0x90, 0xed, 0x7a, 0xe8, 0xc1, 0x21, 0x25, 0x8e,
	0x21, 0x0e, 0xf1, 0x41, 0x12, 0x1b, 0xd8, 0x19, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b,
	0x10, 0x60, 0xc4, 0x85, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizeServiceClient is the client API for AuthorizeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizeServiceClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

type authorizeServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizeServiceClient(cc *grpc.ClientConn) AuthorizeServiceClient {
	return &authorizeServiceClient{cc}
}

func (c *authorizeServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthorizeService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizeServiceServer is the server API for AuthorizeService service.
type AuthorizeServiceServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
}

// UnimplementedAuthorizeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizeServiceServer struct {
}

func (*UnimplementedAuthorizeServiceServer) Authorize(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}

func RegisterAuthorizeServiceServer(s *grpc.Server, srv AuthorizeServiceServer) {
	s.RegisterService(&_AuthorizeService_serviceDesc, srv)
}

func _AuthorizeService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizeServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthorizeService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizeServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthorizeService",
	HandlerType: (*AuthorizeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _AuthorizeService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/auth/authorize.proto",
}
