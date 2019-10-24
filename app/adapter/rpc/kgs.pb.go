// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/adapter/rpc/kgs.proto

package rpc

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PopulateKeysRequest struct {
	KeyLength            uint32   `protobuf:"varint,1,opt,name=key_length,json=keyLength,proto3" json:"key_length,omitempty"`
	RequesterEmail       string   `protobuf:"bytes,2,opt,name=requester_email,json=requesterEmail,proto3" json:"requester_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopulateKeysRequest) Reset()         { *m = PopulateKeysRequest{} }
func (m *PopulateKeysRequest) String() string { return proto.CompactTextString(m) }
func (*PopulateKeysRequest) ProtoMessage()    {}
func (*PopulateKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb3e7329511ae677, []int{0}
}

func (m *PopulateKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopulateKeysRequest.Unmarshal(m, b)
}
func (m *PopulateKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopulateKeysRequest.Marshal(b, m, deterministic)
}
func (m *PopulateKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopulateKeysRequest.Merge(m, src)
}
func (m *PopulateKeysRequest) XXX_Size() int {
	return xxx_messageInfo_PopulateKeysRequest.Size(m)
}
func (m *PopulateKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PopulateKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PopulateKeysRequest proto.InternalMessageInfo

func (m *PopulateKeysRequest) GetKeyLength() uint32 {
	if m != nil {
		return m.KeyLength
	}
	return 0
}

func (m *PopulateKeysRequest) GetRequesterEmail() string {
	if m != nil {
		return m.RequesterEmail
	}
	return ""
}

type AllocateKeysRequest struct {
	MaxKeyCount          uint32   `protobuf:"varint,1,opt,name=max_key_count,json=maxKeyCount,proto3" json:"max_key_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocateKeysRequest) Reset()         { *m = AllocateKeysRequest{} }
func (m *AllocateKeysRequest) String() string { return proto.CompactTextString(m) }
func (*AllocateKeysRequest) ProtoMessage()    {}
func (*AllocateKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb3e7329511ae677, []int{1}
}

func (m *AllocateKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateKeysRequest.Unmarshal(m, b)
}
func (m *AllocateKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateKeysRequest.Marshal(b, m, deterministic)
}
func (m *AllocateKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateKeysRequest.Merge(m, src)
}
func (m *AllocateKeysRequest) XXX_Size() int {
	return xxx_messageInfo_AllocateKeysRequest.Size(m)
}
func (m *AllocateKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateKeysRequest proto.InternalMessageInfo

func (m *AllocateKeysRequest) GetMaxKeyCount() uint32 {
	if m != nil {
		return m.MaxKeyCount
	}
	return 0
}

type AllocateKeysResponse struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocateKeysResponse) Reset()         { *m = AllocateKeysResponse{} }
func (m *AllocateKeysResponse) String() string { return proto.CompactTextString(m) }
func (*AllocateKeysResponse) ProtoMessage()    {}
func (*AllocateKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb3e7329511ae677, []int{2}
}

func (m *AllocateKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateKeysResponse.Unmarshal(m, b)
}
func (m *AllocateKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateKeysResponse.Marshal(b, m, deterministic)
}
func (m *AllocateKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateKeysResponse.Merge(m, src)
}
func (m *AllocateKeysResponse) XXX_Size() int {
	return xxx_messageInfo_AllocateKeysResponse.Size(m)
}
func (m *AllocateKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateKeysResponse proto.InternalMessageInfo

func (m *AllocateKeysResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*PopulateKeysRequest)(nil), "rpc.PopulateKeysRequest")
	proto.RegisterType((*AllocateKeysRequest)(nil), "rpc.AllocateKeysRequest")
	proto.RegisterType((*AllocateKeysResponse)(nil), "rpc.AllocateKeysResponse")
}

func init() { proto.RegisterFile("app/adapter/rpc/kgs.proto", fileDescriptor_eb3e7329511ae677) }

var fileDescriptor_eb3e7329511ae677 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xa9, 0x93, 0x41, 0xff, 0x6e, 0x0a, 0x99, 0x48, 0x57, 0x11, 0x4a, 0x2f, 0x16, 0x0f,
	0x29, 0xe8, 0xc9, 0x9b, 0x32, 0x86, 0x87, 0x7a, 0x90, 0xde, 0xa5, 0x64, 0xf1, 0x6f, 0x95, 0xa6,
	0x4d, 0x4c, 0x52, 0x58, 0x3e, 0x85, 0x5f, 0x59, 0xda, 0x4e, 0x99, 0xa3, 0xb7, 0xf0, 0xcb, 0xcb,
	0x7b, 0x79, 0x0f, 0x96, 0x4c, 0xa9, 0x94, 0xbd, 0x31, 0x65, 0x51, 0xa7, 0x5a, 0xf1, 0xb4, 0x2a,
	0x0d, 0x55, 0x5a, 0x5a, 0x49, 0x26, 0x5a, 0xf1, 0xf0, 0xb2, 0x94, 0xb2, 0x14, 0x98, 0xf6, 0x68,
	0xd3, 0xbe, 0xa7, 0x58, 0x2b, 0xeb, 0x06, 0x45, 0xfc, 0x0a, 0x8b, 0x17, 0xa9, 0x5a, 0xc1, 0x2c,
	0x66, 0xe8, 0x4c, 0x8e, 0x5f, 0x2d, 0x1a, 0x4b, 0xae, 0x00, 0x2a, 0x74, 0x85, 0xc0, 0xa6, 0xb4,
	0x1f, 0x81, 0x17, 0x79, 0xc9, 0x3c, 0xf7, 0x2b, 0x74, 0xcf, 0x3d, 0x20, 0xd7, 0x70, 0xa6, 0x07,
	0x25, 0xea, 0x02, 0x6b, 0xf6, 0x29, 0x82, 0xa3, 0xc8, 0x4b, 0xfc, 0xfc, 0xf4, 0x0f, 0xaf, 0x3b,
	0x1a, 0xdf, 0xc3, 0xe2, 0x51, 0x08, 0xc9, 0x0f, 0xec, 0x63, 0x98, 0xd7, 0x6c, 0x5b, 0x74, 0x11,
	0x5c, 0xb6, 0x8d, 0xdd, 0x25, 0x9c, 0xd4, 0x6c, 0x9b, 0xa1, 0x5b, 0x75, 0x28, 0xbe, 0x81, 0xf3,
	0xff, 0x4f, 0x8d, 0x92, 0x8d, 0x41, 0x42, 0xe0, 0xb8, 0x42, 0x67, 0x02, 0x2f, 0x9a, 0x24, 0x7e,
	0xde, 0x9f, 0x6f, 0xbf, 0x3d, 0x98, 0x66, 0xe8, 0x9e, 0xb0, 0x21, 0x0f, 0x30, 0xdb, 0x2f, 0x44,
	0x02, 0xaa, 0x15, 0xa7, 0x23, 0x1d, 0xc3, 0x0b, 0x3a, 0x0c, 0x43, 0x7f, 0x87, 0xa1, 0xeb, 0x6e,
	0x18, 0xb2, 0x82, 0xd9, 0x7e, 0xf0, 0xce, 0x61, 0xa4, 0x46, 0xb8, 0x1c, 0xb9, 0x19, 0x7e, 0xb9,
	0x99, 0xf6, 0xa6, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xfb, 0xd8, 0x5b, 0x9d, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyGenClient is the client API for KeyGen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyGenClient interface {
	PopulateKeys(ctx context.Context, in *PopulateKeysRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AllocateKeys(ctx context.Context, in *AllocateKeysRequest, opts ...grpc.CallOption) (*AllocateKeysResponse, error)
}

type keyGenClient struct {
	cc *grpc.ClientConn
}

func NewKeyGenClient(cc *grpc.ClientConn) KeyGenClient {
	return &keyGenClient{cc}
}

func (c *keyGenClient) PopulateKeys(ctx context.Context, in *PopulateKeysRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpc.KeyGen/PopulateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyGenClient) AllocateKeys(ctx context.Context, in *AllocateKeysRequest, opts ...grpc.CallOption) (*AllocateKeysResponse, error) {
	out := new(AllocateKeysResponse)
	err := c.cc.Invoke(ctx, "/rpc.KeyGen/AllocateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyGenServer is the server API for KeyGen service.
type KeyGenServer interface {
	PopulateKeys(context.Context, *PopulateKeysRequest) (*empty.Empty, error)
	AllocateKeys(context.Context, *AllocateKeysRequest) (*AllocateKeysResponse, error)
}

// UnimplementedKeyGenServer can be embedded to have forward compatible implementations.
type UnimplementedKeyGenServer struct {
}

func (*UnimplementedKeyGenServer) PopulateKeys(ctx context.Context, req *PopulateKeysRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopulateKeys not implemented")
}
func (*UnimplementedKeyGenServer) AllocateKeys(ctx context.Context, req *AllocateKeysRequest) (*AllocateKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateKeys not implemented")
}

func RegisterKeyGenServer(s *grpc.Server, srv KeyGenServer) {
	s.RegisterService(&_KeyGen_serviceDesc, srv)
}

func _KeyGen_PopulateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopulateKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyGenServer).PopulateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KeyGen/PopulateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyGenServer).PopulateKeys(ctx, req.(*PopulateKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyGen_AllocateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyGenServer).AllocateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.KeyGen/AllocateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyGenServer).AllocateKeys(ctx, req.(*AllocateKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyGen_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.KeyGen",
	HandlerType: (*KeyGenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PopulateKeys",
			Handler:    _KeyGen_PopulateKeys_Handler,
		},
		{
			MethodName: "AllocateKeys",
			Handler:    _KeyGen_AllocateKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/adapter/rpc/kgs.proto",
}
