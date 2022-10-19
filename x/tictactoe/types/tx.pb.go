// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tictactoe/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgInvite struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Opponent string `protobuf:"bytes,2,opt,name=opponent,proto3" json:"opponent,omitempty"`
}

func (m *MsgInvite) Reset()         { *m = MsgInvite{} }
func (m *MsgInvite) String() string { return proto.CompactTextString(m) }
func (*MsgInvite) ProtoMessage()    {}
func (*MsgInvite) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e04dc317c099c4, []int{0}
}
func (m *MsgInvite) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInvite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInvite.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInvite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInvite.Merge(m, src)
}
func (m *MsgInvite) XXX_Size() int {
	return m.Size()
}
func (m *MsgInvite) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInvite.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInvite proto.InternalMessageInfo

func (m *MsgInvite) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgInvite) GetOpponent() string {
	if m != nil {
		return m.Opponent
	}
	return ""
}

type MsgInviteResponse struct {
	GameIndex uint64 `protobuf:"varint,1,opt,name=gameIndex,proto3" json:"gameIndex,omitempty"`
	X         string `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	O         string `protobuf:"bytes,3,opt,name=o,proto3" json:"o,omitempty"`
}

func (m *MsgInviteResponse) Reset()         { *m = MsgInviteResponse{} }
func (m *MsgInviteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInviteResponse) ProtoMessage()    {}
func (*MsgInviteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e04dc317c099c4, []int{1}
}
func (m *MsgInviteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInviteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInviteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInviteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInviteResponse.Merge(m, src)
}
func (m *MsgInviteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInviteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInviteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInviteResponse proto.InternalMessageInfo

func (m *MsgInviteResponse) GetGameIndex() uint64 {
	if m != nil {
		return m.GameIndex
	}
	return 0
}

func (m *MsgInviteResponse) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *MsgInviteResponse) GetO() string {
	if m != nil {
		return m.O
	}
	return ""
}

type MsgAccept struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameIndex uint64 `protobuf:"varint,2,opt,name=gameIndex,proto3" json:"gameIndex,omitempty"`
}

func (m *MsgAccept) Reset()         { *m = MsgAccept{} }
func (m *MsgAccept) String() string { return proto.CompactTextString(m) }
func (*MsgAccept) ProtoMessage()    {}
func (*MsgAccept) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e04dc317c099c4, []int{2}
}
func (m *MsgAccept) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAccept) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAccept.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAccept) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAccept.Merge(m, src)
}
func (m *MsgAccept) XXX_Size() int {
	return m.Size()
}
func (m *MsgAccept) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAccept.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAccept proto.InternalMessageInfo

func (m *MsgAccept) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAccept) GetGameIndex() uint64 {
	if m != nil {
		return m.GameIndex
	}
	return 0
}

type MsgAcceptResponse struct {
}

func (m *MsgAcceptResponse) Reset()         { *m = MsgAcceptResponse{} }
func (m *MsgAcceptResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAcceptResponse) ProtoMessage()    {}
func (*MsgAcceptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e04dc317c099c4, []int{3}
}
func (m *MsgAcceptResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAcceptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAcceptResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAcceptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAcceptResponse.Merge(m, src)
}
func (m *MsgAcceptResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAcceptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAcceptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAcceptResponse proto.InternalMessageInfo

type MsgMove struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameIndex uint64 `protobuf:"varint,2,opt,name=gameIndex,proto3" json:"gameIndex,omitempty"`
	X         uint64 `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y         uint64 `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *MsgMove) Reset()         { *m = MsgMove{} }
func (m *MsgMove) String() string { return proto.CompactTextString(m) }
func (*MsgMove) ProtoMessage()    {}
func (*MsgMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e04dc317c099c4, []int{4}
}
func (m *MsgMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMove.Merge(m, src)
}
func (m *MsgMove) XXX_Size() int {
	return m.Size()
}
func (m *MsgMove) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMove.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMove proto.InternalMessageInfo

func (m *MsgMove) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgMove) GetGameIndex() uint64 {
	if m != nil {
		return m.GameIndex
	}
	return 0
}

func (m *MsgMove) GetX() uint64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *MsgMove) GetY() uint64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type MsgMoveResponse struct {
	Winner string `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (m *MsgMoveResponse) Reset()         { *m = MsgMoveResponse{} }
func (m *MsgMoveResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMoveResponse) ProtoMessage()    {}
func (*MsgMoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e04dc317c099c4, []int{5}
}
func (m *MsgMoveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMoveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMoveResponse.Merge(m, src)
}
func (m *MsgMoveResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMoveResponse proto.InternalMessageInfo

func (m *MsgMoveResponse) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgInvite)(nil), "cjcobb23.tictactoe.tictactoe.MsgInvite")
	proto.RegisterType((*MsgInviteResponse)(nil), "cjcobb23.tictactoe.tictactoe.MsgInviteResponse")
	proto.RegisterType((*MsgAccept)(nil), "cjcobb23.tictactoe.tictactoe.MsgAccept")
	proto.RegisterType((*MsgAcceptResponse)(nil), "cjcobb23.tictactoe.tictactoe.MsgAcceptResponse")
	proto.RegisterType((*MsgMove)(nil), "cjcobb23.tictactoe.tictactoe.MsgMove")
	proto.RegisterType((*MsgMoveResponse)(nil), "cjcobb23.tictactoe.tictactoe.MsgMoveResponse")
}

func init() { proto.RegisterFile("tictactoe/tx.proto", fileDescriptor_82e04dc317c099c4) }

var fileDescriptor_82e04dc317c099c4 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x3b, 0x6d, 0x49, 0xed, 0x43, 0x10, 0x47, 0x90, 0x50, 0x4a, 0x90, 0x80, 0xa8, 0x0b,
	0x13, 0x68, 0x4f, 0x50, 0x5d, 0x75, 0x91, 0x4d, 0x97, 0xa2, 0x8b, 0x66, 0x7c, 0xc4, 0x08, 0x9d,
	0x17, 0x9a, 0xb1, 0xa6, 0xb7, 0xf0, 0x1e, 0x5e, 0xc4, 0x65, 0x97, 0x2e, 0xa5, 0xbd, 0x88, 0x38,
	0x93, 0x4c, 0xad, 0x8b, 0xb6, 0xb8, 0x9b, 0x7f, 0xe6, 0xe5, 0xfb, 0x7f, 0xfe, 0xf0, 0x80, 0xab,
	0x54, 0xa8, 0xb1, 0x50, 0x84, 0xa1, 0x2a, 0x82, 0x6c, 0x4a, 0x8a, 0x78, 0x57, 0x3c, 0x0b, 0x8a,
	0xe3, 0x5e, 0x3f, 0xb0, 0x8f, 0xeb, 0x93, 0x3f, 0x80, 0x76, 0x94, 0x27, 0x43, 0x39, 0x4b, 0x15,
	0x72, 0x17, 0x5a, 0x62, 0x8a, 0x63, 0x45, 0x53, 0x97, 0x9d, 0xb1, 0xcb, 0xf6, 0xa8, 0x92, 0xbc,
	0x03, 0x07, 0x94, 0x65, 0x24, 0x51, 0x2a, 0xb7, 0xae, 0x9f, 0xac, 0xf6, 0x23, 0x38, 0xb6, 0x88,
	0x11, 0xe6, 0x19, 0xc9, 0x1c, 0x79, 0x17, 0xda, 0xc9, 0x78, 0x82, 0x43, 0xf9, 0x88, 0x85, 0x86,
	0x35, 0x47, 0xeb, 0x0b, 0x7e, 0x08, 0xac, 0x28, 0x39, 0x4c, 0x2b, 0x72, 0x1b, 0x46, 0x91, 0x7f,
	0xab, 0x13, 0x0d, 0x84, 0xc0, 0x4c, 0x6d, 0x49, 0xb4, 0x61, 0x50, 0xff, 0x63, 0xe0, 0x9f, 0xe8,
	0x4c, 0x06, 0x52, 0x65, 0xf2, 0x1f, 0xa0, 0x15, 0xe5, 0x49, 0x44, 0x33, 0xfc, 0x2f, 0xd7, 0x04,
	0x6f, 0xe8, 0x5b, 0x13, 0x7c, 0xee, 0x36, 0x8d, 0x9a, 0xfb, 0x57, 0x70, 0x54, 0xe2, 0x6d, 0x0b,
	0xa7, 0xe0, 0xbc, 0xa6, 0x52, 0x62, 0xe5, 0x52, 0xaa, 0xde, 0x7b, 0x1d, 0x1a, 0x51, 0x9e, 0xf0,
	0x18, 0x9c, 0xb2, 0xfa, 0x8b, 0x60, 0xdb, 0x6f, 0x0a, 0x6c, 0xc1, 0x9d, 0x70, 0xcf, 0x41, 0x9b,
	0x21, 0x06, 0xa7, 0x2c, 0x73, 0xb7, 0x87, 0x19, 0xdc, 0xc3, 0x63, 0xb3, 0x59, 0x7e, 0x0f, 0x4d,
	0x5d, 0xeb, 0xf9, 0xce, 0x0f, 0x7f, 0xc6, 0x3a, 0xd7, 0x7b, 0x8d, 0x55, 0xf4, 0x9b, 0xe1, 0xc7,
	0xd2, 0x63, 0x8b, 0xa5, 0xc7, 0xbe, 0x96, 0x1e, 0x7b, 0x5b, 0x79, 0xb5, 0xc5, 0xca, 0xab, 0x7d,
	0xae, 0xbc, 0xda, 0x5d, 0x98, 0xa4, 0xea, 0xe9, 0x25, 0x0e, 0x04, 0x4d, 0xc2, 0x0a, 0x19, 0xae,
	0x77, 0xa0, 0xf8, 0x75, 0x56, 0xf3, 0x0c, 0xf3, 0xd8, 0xd1, 0x3b, 0xd1, 0xff, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xce, 0xdc, 0x28, 0x63, 0x29, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Invite(ctx context.Context, in *MsgInvite, opts ...grpc.CallOption) (*MsgInviteResponse, error)
	Accept(ctx context.Context, in *MsgAccept, opts ...grpc.CallOption) (*MsgAcceptResponse, error)
	Move(ctx context.Context, in *MsgMove, opts ...grpc.CallOption) (*MsgMoveResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Invite(ctx context.Context, in *MsgInvite, opts ...grpc.CallOption) (*MsgInviteResponse, error) {
	out := new(MsgInviteResponse)
	err := c.cc.Invoke(ctx, "/cjcobb23.tictactoe.tictactoe.Msg/Invite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Accept(ctx context.Context, in *MsgAccept, opts ...grpc.CallOption) (*MsgAcceptResponse, error) {
	out := new(MsgAcceptResponse)
	err := c.cc.Invoke(ctx, "/cjcobb23.tictactoe.tictactoe.Msg/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Move(ctx context.Context, in *MsgMove, opts ...grpc.CallOption) (*MsgMoveResponse, error) {
	out := new(MsgMoveResponse)
	err := c.cc.Invoke(ctx, "/cjcobb23.tictactoe.tictactoe.Msg/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Invite(context.Context, *MsgInvite) (*MsgInviteResponse, error)
	Accept(context.Context, *MsgAccept) (*MsgAcceptResponse, error)
	Move(context.Context, *MsgMove) (*MsgMoveResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Invite(ctx context.Context, req *MsgInvite) (*MsgInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invite not implemented")
}
func (*UnimplementedMsgServer) Accept(ctx context.Context, req *MsgAccept) (*MsgAcceptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (*UnimplementedMsgServer) Move(ctx context.Context, req *MsgMove) (*MsgMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Invite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInvite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Invite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cjcobb23.tictactoe.tictactoe.Msg/Invite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Invite(ctx, req.(*MsgInvite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAccept)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cjcobb23.tictactoe.tictactoe.Msg/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Accept(ctx, req.(*MsgAccept))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cjcobb23.tictactoe.tictactoe.Msg/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Move(ctx, req.(*MsgMove))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cjcobb23.tictactoe.tictactoe.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invite",
			Handler:    _Msg_Invite_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _Msg_Accept_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _Msg_Move_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tictactoe/tx.proto",
}

func (m *MsgInvite) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInvite) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInvite) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Opponent) > 0 {
		i -= len(m.Opponent)
		copy(dAtA[i:], m.Opponent)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Opponent)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInviteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInviteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInviteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.O) > 0 {
		i -= len(m.O)
		copy(dAtA[i:], m.O)
		i = encodeVarintTx(dAtA, i, uint64(len(m.O)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarintTx(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgAccept) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAccept) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAccept) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAcceptResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAcceptResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAcceptResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Y != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x20
	}
	if m.X != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x18
	}
	if m.GameIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMoveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMoveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMoveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInvite) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Opponent)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInviteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameIndex != 0 {
		n += 1 + sovTx(uint64(m.GameIndex))
	}
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.O)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAccept) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameIndex != 0 {
		n += 1 + sovTx(uint64(m.GameIndex))
	}
	return n
}

func (m *MsgAcceptResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameIndex != 0 {
		n += 1 + sovTx(uint64(m.GameIndex))
	}
	if m.X != 0 {
		n += 1 + sovTx(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovTx(uint64(m.Y))
	}
	return n
}

func (m *MsgMoveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInvite) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInvite: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInvite: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opponent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Opponent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgInviteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInviteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInviteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			m.GameIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field O", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.O = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAccept) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAccept: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAccept: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			m.GameIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAcceptResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAcceptResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAcceptResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			m.GameIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMoveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMoveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMoveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
