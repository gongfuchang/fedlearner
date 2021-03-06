// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pairing_service.proto

package proto

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

type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type RegisterRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,3,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RegisterRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RegisterRequest) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type PairRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairRequest) Reset()         { *m = PairRequest{} }
func (m *PairRequest) String() string { return proto.CompactTextString(m) }
func (*PairRequest) ProtoMessage()    {}
func (*PairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{2}
}

func (m *PairRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairRequest.Unmarshal(m, b)
}
func (m *PairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairRequest.Marshal(b, m, deterministic)
}
func (m *PairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairRequest.Merge(m, src)
}
func (m *PairRequest) XXX_Size() int {
	return xxx_messageInfo_PairRequest.Size(m)
}
func (m *PairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PairRequest proto.InternalMessageInfo

func (m *PairRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *PairRequest) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type Pair struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	LeaderIds            []string `protobuf:"bytes,2,rep,name=leader_ids,json=leaderIds,proto3" json:"leader_ids,omitempty"`
	FollowerIds          []string `protobuf:"bytes,3,rep,name=follower_ids,json=followerIds,proto3" json:"follower_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{3}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Pair) GetLeaderIds() []string {
	if m != nil {
		return m.LeaderIds
	}
	return nil
}

func (m *Pair) GetFollowerIds() []string {
	if m != nil {
		return m.FollowerIds
	}
	return nil
}

type FinishRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FinishRequest) Reset()         { *m = FinishRequest{} }
func (m *FinishRequest) String() string { return proto.CompactTextString(m) }
func (*FinishRequest) ProtoMessage()    {}
func (*FinishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{4}
}

func (m *FinishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishRequest.Unmarshal(m, b)
}
func (m *FinishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishRequest.Marshal(b, m, deterministic)
}
func (m *FinishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishRequest.Merge(m, src)
}
func (m *FinishRequest) XXX_Size() int {
	return xxx_messageInfo_FinishRequest.Size(m)
}
func (m *FinishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FinishRequest proto.InternalMessageInfo

func (m *FinishRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *FinishRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type ShutDownRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutDownRequest) Reset()         { *m = ShutDownRequest{} }
func (m *ShutDownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutDownRequest) ProtoMessage()    {}
func (*ShutDownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{5}
}

func (m *ShutDownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutDownRequest.Unmarshal(m, b)
}
func (m *ShutDownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutDownRequest.Marshal(b, m, deterministic)
}
func (m *ShutDownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutDownRequest.Merge(m, src)
}
func (m *ShutDownRequest) XXX_Size() int {
	return xxx_messageInfo_ShutDownRequest.Size(m)
}
func (m *ShutDownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutDownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutDownRequest proto.InternalMessageInfo

func (m *ShutDownRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ShutDownRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "fedlearner.Status")
	proto.RegisterType((*RegisterRequest)(nil), "fedlearner.RegisterRequest")
	proto.RegisterType((*PairRequest)(nil), "fedlearner.PairRequest")
	proto.RegisterType((*Pair)(nil), "fedlearner.Pair")
	proto.RegisterType((*FinishRequest)(nil), "fedlearner.FinishRequest")
	proto.RegisterType((*ShutDownRequest)(nil), "fedlearner.ShutDownRequest")
}

func init() { proto.RegisterFile("pairing_service.proto", fileDescriptor_508234a4704e8446) }

var fileDescriptor_508234a4704e8446 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x86, 0xe9, 0x57, 0x30, 0xa7, 0xbb, 0xae, 0x0c, 0xac, 0xd6, 0x5d, 0x84, 0x1a, 0x41, 0x7a,
	0x95, 0x40, 0xf7, 0x42, 0x70, 0xbd, 0x51, 0x74, 0xb1, 0xa0, 0x50, 0xd2, 0x3b, 0x11, 0xc2, 0x24,
	0x73, 0x9a, 0x0e, 0xa6, 0x99, 0x71, 0x66, 0xe2, 0x92, 0xff, 0xe4, 0x8f, 0x94, 0xcc, 0x24, 0x6e,
	0x77, 0xad, 0x85, 0x5e, 0xf5, 0xcc, 0xf9, 0x78, 0xe7, 0xf4, 0x99, 0x37, 0x70, 0x2e, 0x29, 0x57,
	0xbc, 0xcc, 0x13, 0x8d, 0xea, 0x17, 0xcf, 0x30, 0x94, 0x4a, 0x18, 0x41, 0x60, 0x8d, 0xac, 0x40,
	0xaa, 0x4a, 0x54, 0x17, 0x97, 0xb9, 0x10, 0x79, 0x81, 0x91, 0xad, 0xa4, 0xd5, 0x3a, 0xc2, 0xad,
	0x34, 0xb5, 0x6b, 0x0c, 0xde, 0x83, 0xb7, 0x32, 0xd4, 0x54, 0x9a, 0x10, 0x18, 0x66, 0x82, 0xe1,
	0xa4, 0x37, 0xed, 0xcd, 0x46, 0xb1, 0x8d, 0xc9, 0x2b, 0x38, 0x45, 0xa5, 0x84, 0x4a, 0xb6, 0xa8,
	0x35, 0xcd, 0x71, 0xd2, 0x9f, 0xf6, 0x66, 0x7e, 0x7c, 0x62, 0x93, 0x5f, 0x5d, 0x2e, 0x60, 0x70,
	0x16, 0x63, 0xce, 0xb5, 0x41, 0x15, 0xe3, 0xcf, 0x0a, 0xb5, 0x21, 0xe7, 0xe0, 0x51, 0x29, 0x13,
	0xce, 0xac, 0x9a, 0x1f, 0x8f, 0xa8, 0x94, 0x0b, 0xd6, 0x5c, 0xa1, 0x44, 0xd1, 0xa9, 0xd8, 0x98,
	0xbc, 0x86, 0x51, 0xf3, 0x17, 0xf4, 0x64, 0x30, 0x1d, 0xcc, 0xc6, 0xf3, 0x27, 0xe1, 0xdd, 0xe6,
	0xe1, 0x92, 0x72, 0x15, 0xbb, 0x72, 0xf0, 0x05, 0xc6, 0xf6, 0x78, 0xf8, 0x86, 0xbf, 0x6a, 0xfd,
	0xc3, 0x6a, 0xdf, 0x61, 0xd8, 0x1c, 0x9b, 0x8d, 0x4c, 0x2d, 0xb1, 0x15, 0xb1, 0x31, 0x79, 0x01,
	0x50, 0x20, 0x65, 0xa8, 0x12, 0xce, 0x9c, 0x90, 0x1f, 0xfb, 0x2e, 0xb3, 0x60, 0x9a, 0xbc, 0x84,
	0x93, 0xb5, 0x28, 0x0a, 0x71, 0xdb, 0x36, 0x0c, 0x6c, 0xc3, 0xb8, 0xcb, 0x2d, 0x98, 0x0e, 0xde,
	0xc2, 0xe9, 0x0d, 0x2f, 0xb9, 0xde, 0x1c, 0xcf, 0x23, 0x78, 0x07, 0x67, 0xab, 0x4d, 0x65, 0x3e,
	0x8a, 0xdb, 0xf2, 0xf8, 0xe9, 0xf9, 0xef, 0x3e, 0x3c, 0x5e, 0x3a, 0x47, 0xac, 0x9c, 0x21, 0xc8,
	0x1c, 0x86, 0x4b, 0x5e, 0xe6, 0xe4, 0x69, 0xe8, 0x7c, 0x10, 0x76, 0x3e, 0x08, 0x3f, 0x35, 0x3e,
	0xb8, 0x20, 0xbb, 0x8c, 0x5a, 0x2f, 0x5c, 0xc3, 0xa3, 0xee, 0x49, 0xc9, 0xe5, 0x6e, 0xfd, 0xc1,
	0x43, 0xef, 0x1d, 0xbe, 0x6a, 0xd9, 0x3e, 0xfb, 0x07, 0xfe, 0x81, 0xa1, 0x37, 0xe0, 0x39, 0x64,
	0xe4, 0xf9, 0x6e, 0xf5, 0x1e, 0xc6, 0xff, 0xad, 0xda, 0xf1, 0xba, 0xbf, 0xea, 0x03, 0x8a, 0xfb,
	0x86, 0x3f, 0x7c, 0xfe, 0x76, 0x93, 0x73, 0xb3, 0xa9, 0xd2, 0x30, 0x13, 0xdb, 0x28, 0xad, 0x0d,
	0x32, 0x5a, 0x66, 0x18, 0xdd, 0x75, 0x46, 0x0c, 0x65, 0x21, 0xea, 0xe8, 0x47, 0x95, 0xa2, 0x2a,
	0xd1, 0xa0, 0x4e, 0x84, 0x44, 0x45, 0x8d, 0x50, 0xee, 0x83, 0xba, 0x76, 0x38, 0x3d, 0xfb, 0x73,
	0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x5e, 0x80, 0xbd, 0x90, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PairingServiceClient is the client API for PairingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PairingServiceClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Status, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Status, error)
	Pair(ctx context.Context, in *PairRequest, opts ...grpc.CallOption) (*Status, error)
	Finish(ctx context.Context, in *FinishRequest, opts ...grpc.CallOption) (*Status, error)
	ShutDown(ctx context.Context, in *ShutDownRequest, opts ...grpc.CallOption) (*Status, error)
}

type pairingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPairingServiceClient(cc *grpc.ClientConn) PairingServiceClient {
	return &pairingServiceClient{cc}
}

func (c *pairingServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/fedlearner.PairingService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/fedlearner.PairingService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) Pair(ctx context.Context, in *PairRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/fedlearner.PairingService/Pair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) Finish(ctx context.Context, in *FinishRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/fedlearner.PairingService/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) ShutDown(ctx context.Context, in *ShutDownRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/fedlearner.PairingService/ShutDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PairingServiceServer is the server API for PairingService service.
type PairingServiceServer interface {
	Ping(context.Context, *empty.Empty) (*Status, error)
	Register(context.Context, *RegisterRequest) (*Status, error)
	Pair(context.Context, *PairRequest) (*Status, error)
	Finish(context.Context, *FinishRequest) (*Status, error)
	ShutDown(context.Context, *ShutDownRequest) (*Status, error)
}

// UnimplementedPairingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPairingServiceServer struct {
}

func (*UnimplementedPairingServiceServer) Ping(ctx context.Context, req *empty.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedPairingServiceServer) Register(ctx context.Context, req *RegisterRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedPairingServiceServer) Pair(ctx context.Context, req *PairRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pair not implemented")
}
func (*UnimplementedPairingServiceServer) Finish(ctx context.Context, req *FinishRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}
func (*UnimplementedPairingServiceServer) ShutDown(ctx context.Context, req *ShutDownRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutDown not implemented")
}

func RegisterPairingServiceServer(s *grpc.Server, srv PairingServiceServer) {
	s.RegisterService(&_PairingService_serviceDesc, srv)
}

func _PairingService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fedlearner.PairingService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fedlearner.PairingService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_Pair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Pair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fedlearner.PairingService/Pair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Pair(ctx, req.(*PairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fedlearner.PairingService/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Finish(ctx, req.(*FinishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_ShutDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).ShutDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fedlearner.PairingService/ShutDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).ShutDown(ctx, req.(*ShutDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PairingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fedlearner.PairingService",
	HandlerType: (*PairingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PairingService_Ping_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PairingService_Register_Handler,
		},
		{
			MethodName: "Pair",
			Handler:    _PairingService_Pair_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _PairingService_Finish_Handler,
		},
		{
			MethodName: "ShutDown",
			Handler:    _PairingService_ShutDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pairing_service.proto",
}
