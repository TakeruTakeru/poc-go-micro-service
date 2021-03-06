// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/auth/authentication.proto

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

type AuthenticationRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a542fe5b729fe5, []int{0}
}

func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(m, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticationRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticationResponse struct {
	Authorized           bool                 `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	Token                string               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthenticationResponse) Reset()         { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()    {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a542fe5b729fe5, []int{1}
}

func (m *AuthenticationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResponse.Unmarshal(m, b)
}
func (m *AuthenticationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResponse.Merge(m, src)
}
func (m *AuthenticationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResponse.Size(m)
}
func (m *AuthenticationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResponse proto.InternalMessageInfo

func (m *AuthenticationResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthenticationResponse) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

func (m *AuthenticationResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthenticationRequest)(nil), "auth.AuthenticationRequest")
	proto.RegisterType((*AuthenticationResponse)(nil), "auth.AuthenticationResponse")
}

func init() { proto.RegisterFile("api/auth/authentication.proto", fileDescriptor_00a542fe5b729fe5) }

var fileDescriptor_00a542fe5b729fe5 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x3d, 0x4e, 0x03, 0x31,
	0x10, 0x85, 0xb5, 0xfc, 0x29, 0x19, 0xa8, 0x2c, 0x40, 0xab, 0xe5, 0x57, 0xa9, 0xa8, 0xbc, 0x52,
	0x68, 0x68, 0x53, 0x21, 0x2a, 0x24, 0x93, 0x0b, 0x38, 0xd9, 0x21, 0x58, 0xb0, 0x1e, 0x63, 0xcf,
	0x02, 0xe2, 0x04, 0xdc, 0x85, 0x4b, 0x22, 0xdb, 0x2c, 0x84, 0x28, 0x1d, 0x8d, 0xa5, 0x79, 0x6f,
	0xfc, 0xfc, 0x3e, 0xc3, 0x89, 0x76, 0xa6, 0xd6, 0x1d, 0x3f, 0xa4, 0x03, 0x2d, 0x9b, 0xb9, 0x66,
	0x43, 0x56, 0x3a, 0x4f, 0x4c, 0x62, 0x2b, 0xaa, 0xd5, 0xd9, 0x82, 0x68, 0xf1, 0x84, 0x75, 0xd2,
	0x66, 0xdd, 0x7d, 0xcd, 0xa6, 0xc5, 0xc0, 0xba, 0x75, 0x79, 0x6d, 0x74, 0x0b, 0x07, 0x93, 0x3f,
	0xd7, 0x15, 0x3e, 0x77, 0x18, 0x58, 0x54, 0x30, 0xe8, 0x02, 0x7a, 0xab, 0x5b, 0x2c, 0x8b, 0xf3,
	0xe2, 0x62, 0xa8, 0x7e, 0xe6, 0xe8, 0x39, 0x1d, 0xc2, 0x2b, 0xf9, 0xa6, 0xdc, 0xc8, 0x5e, 0x3f,
	0x8f, 0x3e, 0x0a, 0x38, 0x5c, 0x4d, 0x0c, 0x8e, 0x6c, 0x40, 0x71, 0x0a, 0x10, 0x4b, 0x91, 0x37,
	0xef, 0xd8, 0xa4, 0xd0, 0x81, 0x5a, 0x52, 0xc4, 0x15, 0x0c, 0xf1, 0xcd, 0x19, 0x8f, 0xcd, 0x84,
	0x53, 0xee, 0xee, 0xb8, 0x92, 0x19, 0x40, 0xf6, 0x00, 0x72, 0xda, 0x03, 0xa8, 0xdf, 0x65, 0xb1,
	0x0f, 0xdb, 0x4c, 0x8f, 0x68, 0xcb, 0xcd, 0xd4, 0x26, 0x0f, 0xe3, 0xcf, 0x62, 0x15, 0xee, 0x0e,
	0xfd, 0x8b, 0x99, 0xa3, 0xb8, 0x81, 0xbd, 0x25, 0x03, 0xc5, 0x91, 0x8c, 0x35, 0xe4, 0xda, 0x9f,
	0xa8, 0x8e, 0xd7, 0x9b, 0xdf, 0x50, 0xd7, 0x00, 0x0a, 0x03, 0xf2, 0x34, 0x3e, 0xf9, 0x8f, 0xa0,
	0xd9, 0x4e, 0x42, 0xbc, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xfe, 0xc2, 0xdf, 0xd8, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	ResetToken(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ResetToken(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationService/ResetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	ResetToken(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) Authenticate(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (*UnimplementedAuthenticationServiceServer) ResetToken(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetToken not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ResetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ResetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationService/ResetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ResetToken(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthenticationService_Authenticate_Handler,
		},
		{
			MethodName: "ResetToken",
			Handler:    _AuthenticationService_ResetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/auth/authentication.proto",
}
