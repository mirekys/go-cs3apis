// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/authprovider/v0alpha/authprovider.proto

package authproviderv0alphapb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc"
	types "github.com/cs3org/go-cs3apis/cs3/types"
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

type AuthenticateRequest struct {
	Opaque               *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	ClientId             string        `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string        `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e542c8cad6d0bb, []int{0}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *AuthenticateRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AuthenticateRequest) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

type AuthenticateResponse struct {
	Status               *rpc.Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId               *types.UserId `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e542c8cad6d0bb, []int{1}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AuthenticateResponse) GetUserId() *types.UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticateRequest)(nil), "cs3.authproviderv0alpha.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "cs3.authproviderv0alpha.AuthenticateResponse")
}

func init() {
	proto.RegisterFile("cs3/authprovider/v0alpha/authprovider.proto", fileDescriptor_08e542c8cad6d0bb)
}

var fileDescriptor_08e542c8cad6d0bb = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6a, 0xfa, 0x40,
	0x14, 0xc5, 0x89, 0x7f, 0xc8, 0xbf, 0x4e, 0x2d, 0xc5, 0xa9, 0xa2, 0xe8, 0x46, 0xec, 0xa2, 0xf6,
	0x6b, 0x14, 0xf3, 0x04, 0x89, 0x2b, 0x57, 0x4a, 0x42, 0x4b, 0x29, 0x85, 0x12, 0x27, 0x17, 0x0c,
	0x5a, 0x33, 0xce, 0x87, 0xd0, 0x6d, 0x17, 0x7d, 0x90, 0x2e, 0xfb, 0x28, 0x7d, 0xaa, 0x92, 0x99,
	0x11, 0x22, 0x44, 0xe8, 0x26, 0x81, 0xdf, 0x39, 0x37, 0x39, 0xe7, 0xce, 0xa0, 0x5b, 0x2a, 0xbc,
	0x61, 0xac, 0xe4, 0x92, 0xf1, 0x6c, 0x97, 0x26, 0xc0, 0x87, 0xbb, 0x51, 0xbc, 0x66, 0xcb, 0xf8,
	0x00, 0x12, 0xc6, 0x33, 0x99, 0xe1, 0x16, 0x15, 0x1e, 0x29, 0x72, 0xeb, 0xed, 0x34, 0xf2, 0xaf,
	0x70, 0x46, 0x87, 0x42, 0xc6, 0x52, 0x09, 0x63, 0xef, 0x34, 0x73, 0x2a, 0xdf, 0x19, 0x08, 0xf3,
	0x34, 0xb8, 0xff, 0xe1, 0xa0, 0x0b, 0x5f, 0xc9, 0x25, 0x6c, 0x64, 0x4a, 0x63, 0x09, 0x21, 0x6c,
	0x15, 0x08, 0x89, 0xaf, 0x91, 0x9b, 0xb1, 0x78, 0xab, 0xa0, 0xed, 0xf4, 0x9c, 0xc1, 0xe9, 0xb8,
	0x4e, 0xf2, 0xdf, 0x99, 0xc9, 0x99, 0x16, 0x42, 0x6b, 0xc0, 0x5d, 0x54, 0xa5, 0xeb, 0x14, 0x36,
	0xf2, 0x35, 0x4d, 0xda, 0x95, 0x9e, 0x33, 0xa8, 0x86, 0x27, 0x06, 0x4c, 0x13, 0x7c, 0x89, 0xce,
	0xac, 0x28, 0x80, 0x72, 0x90, 0xed, 0x7f, 0xda, 0x50, 0x33, 0x30, 0xd2, 0xac, 0xbf, 0x42, 0x8d,
	0xc3, 0x0c, 0x82, 0x65, 0x1b, 0x01, 0xf8, 0x0a, 0xb9, 0xa6, 0x83, 0x0d, 0x71, 0xae, 0x43, 0x70,
	0x46, 0x49, 0xa4, 0x71, 0x68, 0x65, 0x7c, 0x83, 0xfe, 0x2b, 0x01, 0x7c, 0x1f, 0xe0, 0x30, 0xee,
	0x83, 0x00, 0x3e, 0x4d, 0x42, 0x57, 0xe9, 0xf7, 0x78, 0xdf, 0x78, 0x6e, 0xd7, 0x16, 0x01, 0xdf,
	0xa5, 0x14, 0xf0, 0x0a, 0xd5, 0x8a, 0x21, 0xf0, 0x1d, 0x39, 0xb2, 0x60, 0x52, 0xb2, 0xaf, 0xce,
	0xfd, 0x1f, 0xdd, 0xa6, 0x59, 0xf0, 0xe9, 0xa0, 0x2e, 0xcd, 0xde, 0x8e, 0x0d, 0x05, 0x75, 0xbf,
	0x00, 0xe7, 0xf9, 0x49, 0xcd, 0x9d, 0xe7, 0x66, 0x89, 0x93, 0x2d, 0xbe, 0x2a, 0xee, 0x24, 0x98,
	0x3d, 0xf9, 0xc1, 0x77, 0xa5, 0x35, 0x89, 0x3c, 0x52, 0xec, 0xf6, 0x38, 0xf2, 0x73, 0xd3, 0x8f,
	0x56, 0x5e, 0x4a, 0x94, 0x85, 0xab, 0xaf, 0x81, 0xf7, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x0e,
	0x3f, 0xb9, 0x7b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthProviderServiceClient is the client API for AuthProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthProviderServiceClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type authProviderServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthProviderServiceClient(cc *grpc.ClientConn) AuthProviderServiceClient {
	return &authProviderServiceClient{cc}
}

func (c *authProviderServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/cs3.authproviderv0alpha.AuthProviderService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthProviderServiceServer is the server API for AuthProviderService service.
type AuthProviderServiceServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

// UnimplementedAuthProviderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthProviderServiceServer struct {
}

func (*UnimplementedAuthProviderServiceServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func RegisterAuthProviderServiceServer(s *grpc.Server, srv AuthProviderServiceServer) {
	s.RegisterService(&_AuthProviderService_serviceDesc, srv)
}

func _AuthProviderService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthProviderServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authproviderv0alpha.AuthProviderService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthProviderServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthProviderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.authproviderv0alpha.AuthProviderService",
	HandlerType: (*AuthProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthProviderService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/authprovider/v0alpha/authprovider.proto",
}
