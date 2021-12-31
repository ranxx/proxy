// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfers.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1 "github.com/ranxx/proxy/proto/msg/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Transfer struct {
	Network v1.NetworkType `protobuf:"varint,2,opt,name=network,proto3,enum=v1.NetworkType" json:"network,omitempty"`
	Laddr   []*v1.Addr     `protobuf:"bytes,3,rep,name=laddr,proto3" json:"laddr,omitempty"`
	Raddr   *v1.Addr       `protobuf:"bytes,4,opt,name=raddr,proto3" json:"raddr,omitempty"`
	Machine string         `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_61401ad0522767fa, []int{0}
}
func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return m.Size()
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetNetwork() v1.NetworkType {
	if m != nil {
		return m.Network
	}
	return v1.NetworkType_HTTP
}

func (m *Transfer) GetLaddr() []*v1.Addr {
	if m != nil {
		return m.Laddr
	}
	return nil
}

func (m *Transfer) GetRaddr() *v1.Addr {
	if m != nil {
		return m.Raddr
	}
	return nil
}

func (m *Transfer) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

type ListTransferReq struct {
}

func (m *ListTransferReq) Reset()         { *m = ListTransferReq{} }
func (m *ListTransferReq) String() string { return proto.CompactTextString(m) }
func (*ListTransferReq) ProtoMessage()    {}
func (*ListTransferReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61401ad0522767fa, []int{1}
}
func (m *ListTransferReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListTransferReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListTransferReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListTransferReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransferReq.Merge(m, src)
}
func (m *ListTransferReq) XXX_Size() int {
	return m.Size()
}
func (m *ListTransferReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransferReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransferReq proto.InternalMessageInfo

type ListTransferRsp struct {
	Items []*Transfer `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *ListTransferRsp) Reset()         { *m = ListTransferRsp{} }
func (m *ListTransferRsp) String() string { return proto.CompactTextString(m) }
func (*ListTransferRsp) ProtoMessage()    {}
func (*ListTransferRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61401ad0522767fa, []int{2}
}
func (m *ListTransferRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListTransferRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListTransferRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListTransferRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransferRsp.Merge(m, src)
}
func (m *ListTransferRsp) XXX_Size() int {
	return m.Size()
}
func (m *ListTransferRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransferRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransferRsp proto.InternalMessageInfo

func (m *ListTransferRsp) GetItems() []*Transfer {
	if m != nil {
		return m.Items
	}
	return nil
}

type AddTransferReq struct {
	Transfers []*Transfer `protobuf:"bytes,1,rep,name=transfers,proto3" json:"transfers,omitempty"`
}

func (m *AddTransferReq) Reset()         { *m = AddTransferReq{} }
func (m *AddTransferReq) String() string { return proto.CompactTextString(m) }
func (*AddTransferReq) ProtoMessage()    {}
func (*AddTransferReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61401ad0522767fa, []int{3}
}
func (m *AddTransferReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddTransferReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddTransferReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddTransferReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTransferReq.Merge(m, src)
}
func (m *AddTransferReq) XXX_Size() int {
	return m.Size()
}
func (m *AddTransferReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTransferReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddTransferReq proto.InternalMessageInfo

func (m *AddTransferReq) GetTransfers() []*Transfer {
	if m != nil {
		return m.Transfers
	}
	return nil
}

type AddTransferRsp struct {
	Item *v1.ClientBind `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (m *AddTransferRsp) Reset()         { *m = AddTransferRsp{} }
func (m *AddTransferRsp) String() string { return proto.CompactTextString(m) }
func (*AddTransferRsp) ProtoMessage()    {}
func (*AddTransferRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61401ad0522767fa, []int{4}
}
func (m *AddTransferRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddTransferRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddTransferRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddTransferRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTransferRsp.Merge(m, src)
}
func (m *AddTransferRsp) XXX_Size() int {
	return m.Size()
}
func (m *AddTransferRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTransferRsp.DiscardUnknown(m)
}

var xxx_messageInfo_AddTransferRsp proto.InternalMessageInfo

func (m *AddTransferRsp) GetItem() *v1.ClientBind {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*Transfer)(nil), "v1.Transfer")
	proto.RegisterType((*ListTransferReq)(nil), "v1.ListTransferReq")
	proto.RegisterType((*ListTransferRsp)(nil), "v1.ListTransferRsp")
	proto.RegisterType((*AddTransferReq)(nil), "v1.AddTransferReq")
	proto.RegisterType((*AddTransferRsp)(nil), "v1.AddTransferRsp")
}

func init() { proto.RegisterFile("transfers.proto", fileDescriptor_61401ad0522767fa) }

var fileDescriptor_61401ad0522767fa = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xbf, 0x6e, 0x13, 0x31,
	0x1c, 0x8e, 0xd3, 0xa6, 0x4d, 0x9c, 0x2a, 0xa1, 0x86, 0xe1, 0x14, 0x85, 0x23, 0xf2, 0x94, 0x56,
	0xea, 0x39, 0x77, 0xc0, 0x52, 0x31, 0x90, 0xb2, 0x22, 0x86, 0xa8, 0xea, 0xc0, 0xe6, 0xe6, 0xdc,
	0xd4, 0x90, 0xd8, 0xc6, 0x36, 0x47, 0xba, 0x32, 0xb0, 0x52, 0x95, 0x47, 0xe0, 0x01, 0x78, 0x80,
	0xbe, 0x00, 0x63, 0x25, 0x16, 0x46, 0x94, 0x54, 0xe2, 0x35, 0xaa, 0x73, 0x72, 0xc9, 0x35, 0xe9,
	0x74, 0xf2, 0xf7, 0xfd, 0xbe, 0x3f, 0xf6, 0xfd, 0x60, 0xdd, 0x6a, 0x2a, 0xcc, 0x19, 0xd3, 0x26,
	0x50, 0x5a, 0x5a, 0x89, 0x8a, 0x49, 0xd8, 0x68, 0x0e, 0xa4, 0x1c, 0x0c, 0x19, 0xa1, 0x8a, 0x13,
	0x2a, 0x84, 0xb4, 0xd4, 0x72, 0x29, 0xe6, 0x13, 0x8d, 0xd9, 0xa7, 0x7f, 0x30, 0x60, 0xe2, 0x40,
	0x2a, 0x26, 0xa8, 0xe2, 0x49, 0x44, 0xa4, 0x72, 0x33, 0x0f, 0xcc, 0x37, 0x95, 0x96, 0xe3, 0x0b,
	0xe2, 0x0e, 0x64, 0x64, 0x06, 0x24, 0x09, 0xd3, 0xcf, 0x8c, 0xc5, 0xdf, 0x01, 0x2c, 0x1f, 0xcf,
	0x3b, 0xa0, 0x3d, 0xb8, 0x2d, 0x98, 0xfd, 0x22, 0xf5, 0x47, 0xaf, 0xd8, 0x02, 0xed, 0x5a, 0x54,
	0x0f, 0x92, 0x30, 0x78, 0x37, 0x83, 0x8e, 0x2f, 0x14, 0xeb, 0x65, 0x3c, 0xf2, 0x61, 0x69, 0x48,
	0xe3, 0x58, 0x7b, 0x1b, 0xad, 0x8d, 0x76, 0x35, 0x2a, 0xa7, 0x83, 0xdd, 0x38, 0xd6, 0xbd, 0x19,
	0x9c, 0xf2, 0xda, 0xf1, 0x9b, 0x2d, 0x70, 0x9f, 0x77, 0x30, 0xf2, 0xe0, 0xf6, 0x88, 0xf6, 0xcf,
	0xb9, 0x60, 0x5e, 0xa9, 0x05, 0xda, 0x95, 0x5e, 0x76, 0xc4, 0xbb, 0xb0, 0xfe, 0x96, 0x1b, 0x9b,
	0x95, 0xea, 0xb1, 0x4f, 0xf8, 0xe5, 0x0a, 0x64, 0x14, 0xc2, 0xb0, 0xc4, 0x2d, 0x1b, 0x19, 0x0f,
	0xb8, 0xfc, 0x9d, 0xd4, 0x7f, 0xc1, 0xcf, 0x28, 0xfc, 0x0a, 0xd6, 0xba, 0x71, 0x9c, 0x33, 0x42,
	0xfb, 0xb0, 0xb2, 0x78, 0xf0, 0x07, 0x95, 0x4b, 0x1a, 0xbf, 0xb8, 0xaf, 0x76, 0x99, 0x9b, 0xa9,
	0xb1, 0x07, 0xdc, 0x95, 0x6a, 0xa9, 0xf0, 0xcd, 0x90, 0x33, 0x61, 0x8f, 0xb8, 0x88, 0x7b, 0x8e,
	0x8b, 0xae, 0x01, 0xac, 0x64, 0x1a, 0x83, 0x4e, 0xe0, 0x4e, 0xbe, 0x38, 0x7a, 0x9c, 0x6a, 0x56,
	0x6e, 0xd7, 0x58, 0x07, 0x8d, 0xc2, 0xcd, 0xab, 0x6e, 0x19, 0x6e, 0xdd, 0x5e, 0xff, 0xfc, 0x7f,
	0xf9, 0xed, 0xeb, 0x9f, 0xdb, 0x1f, 0xc5, 0x2a, 0xaa, 0x90, 0xac, 0x1c, 0x3a, 0x81, 0xd5, 0x5c,
	0x37, 0x84, 0xe6, 0xaf, 0x9b, 0x77, 0x5d, 0xc3, 0x8c, 0xc2, 0xcf, 0x56, 0x4d, 0x6b, 0x78, 0x69,
	0x7a, 0x08, 0xf6, 0x8f, 0x7e, 0x81, 0xab, 0x2e, 0x47, 0x9d, 0xe5, 0x4a, 0xe0, 0xa7, 0xe9, 0x7f,
	0x14, 0xe3, 0x71, 0xe3, 0x89, 0xa6, 0x22, 0x30, 0x96, 0xea, 0xd7, 0x67, 0x72, 0x3c, 0xa2, 0x7c,
	0x18, 0xf4, 0xe5, 0x28, 0x2a, 0x85, 0x41, 0x27, 0xe8, 0xe0, 0x5d, 0xb7, 0xab, 0x49, 0xb8, 0x30,
	0x33, 0xd1, 0x23, 0xaa, 0xd4, 0x90, 0xf7, 0xdd, 0x2a, 0x92, 0x0f, 0x46, 0x8a, 0xc3, 0x35, 0xe4,
	0xf7, 0xc4, 0x07, 0x37, 0x13, 0x1f, 0xfc, 0x9b, 0xf8, 0xe0, 0x72, 0xea, 0x17, 0x6e, 0xa6, 0x7e,
	0xe1, 0xef, 0xd4, 0x2f, 0xbc, 0xdf, 0x1b, 0x70, 0x7b, 0xfe, 0xf9, 0x34, 0x4d, 0x20, 0x2e, 0x9d,
	0xe4, 0x77, 0x78, 0x91, 0x41, 0x92, 0xf0, 0x74, 0xcb, 0x61, 0xcf, 0xef, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0xac, 0xe1, 0x48, 0x49, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransfersClient is the client API for Transfers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransfersClient interface {
	// 代理端口列表
	ListTransfer(ctx context.Context, in *ListTransferReq, opts ...grpc.CallOption) (*ListTransferRsp, error)
	// 新增代理端口
	AddTransfer(ctx context.Context, in *AddTransferReq, opts ...grpc.CallOption) (*AddTransferRsp, error)
}

type transfersClient struct {
	cc *grpc.ClientConn
}

func NewTransfersClient(cc *grpc.ClientConn) TransfersClient {
	return &transfersClient{cc}
}

func (c *transfersClient) ListTransfer(ctx context.Context, in *ListTransferReq, opts ...grpc.CallOption) (*ListTransferRsp, error) {
	out := new(ListTransferRsp)
	err := c.cc.Invoke(ctx, "/v1.Transfers/ListTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfersClient) AddTransfer(ctx context.Context, in *AddTransferReq, opts ...grpc.CallOption) (*AddTransferRsp, error) {
	out := new(AddTransferRsp)
	err := c.cc.Invoke(ctx, "/v1.Transfers/AddTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfersServer is the server API for Transfers service.
type TransfersServer interface {
	// 代理端口列表
	ListTransfer(context.Context, *ListTransferReq) (*ListTransferRsp, error)
	// 新增代理端口
	AddTransfer(context.Context, *AddTransferReq) (*AddTransferRsp, error)
}

// UnimplementedTransfersServer can be embedded to have forward compatible implementations.
type UnimplementedTransfersServer struct {
}

func (*UnimplementedTransfersServer) ListTransfer(ctx context.Context, req *ListTransferReq) (*ListTransferRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransfer not implemented")
}
func (*UnimplementedTransfersServer) AddTransfer(ctx context.Context, req *AddTransferReq) (*AddTransferRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransfer not implemented")
}

func RegisterTransfersServer(s *grpc.Server, srv TransfersServer) {
	s.RegisterService(&_Transfers_serviceDesc, srv)
}

func _Transfers_ListTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).ListTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Transfers/ListTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).ListTransfer(ctx, req.(*ListTransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfers_AddTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServer).AddTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Transfers/AddTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServer).AddTransfer(ctx, req.(*AddTransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transfers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Transfers",
	HandlerType: (*TransfersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransfer",
			Handler:    _Transfers_ListTransfer_Handler,
		},
		{
			MethodName: "AddTransfer",
			Handler:    _Transfers_AddTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfers.proto",
}

func (m *Transfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transfer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Network != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTransfers(dAtA, i, uint64(m.Network))
	}
	if len(m.Laddr) > 0 {
		for _, msg := range m.Laddr {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintTransfers(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Raddr != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTransfers(dAtA, i, uint64(m.Raddr.Size()))
		n1, err1 := m.Raddr.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if len(m.Machine) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTransfers(dAtA, i, uint64(len(m.Machine)))
		i += copy(dAtA[i:], m.Machine)
	}
	return i, nil
}

func (m *ListTransferReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListTransferReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ListTransferRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListTransferRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTransfers(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *AddTransferReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddTransferReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Transfers) > 0 {
		for _, msg := range m.Transfers {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTransfers(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *AddTransferRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddTransferRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Item != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTransfers(dAtA, i, uint64(m.Item.Size()))
		n2, err2 := m.Item.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}

func encodeVarintTransfers(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Transfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Network != 0 {
		n += 1 + sovTransfers(uint64(m.Network))
	}
	if len(m.Laddr) > 0 {
		for _, e := range m.Laddr {
			l = e.Size()
			n += 1 + l + sovTransfers(uint64(l))
		}
	}
	if m.Raddr != nil {
		l = m.Raddr.Size()
		n += 1 + l + sovTransfers(uint64(l))
	}
	l = len(m.Machine)
	if l > 0 {
		n += 1 + l + sovTransfers(uint64(l))
	}
	return n
}

func (m *ListTransferReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListTransferRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovTransfers(uint64(l))
		}
	}
	return n
}

func (m *AddTransferReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Transfers) > 0 {
		for _, e := range m.Transfers {
			l = e.Size()
			n += 1 + l + sovTransfers(uint64(l))
		}
	}
	return n
}

func (m *AddTransferRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Item != nil {
		l = m.Item.Size()
		n += 1 + l + sovTransfers(uint64(l))
	}
	return n
}

func sovTransfers(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransfers(x uint64) (n int) {
	return sovTransfers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfers
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
			return fmt.Errorf("proto: Transfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			m.Network = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Network |= v1.NetworkType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Laddr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Laddr = append(m.Laddr, &v1.Addr{})
			if err := m.Laddr[len(m.Laddr)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raddr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Raddr == nil {
				m.Raddr = &v1.Addr{}
			}
			if err := m.Raddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Machine", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
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
				return ErrInvalidLengthTransfers
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Machine = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfers
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
func (m *ListTransferReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfers
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
			return fmt.Errorf("proto: ListTransferReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListTransferReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTransfers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfers
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
func (m *ListTransferRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfers
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
			return fmt.Errorf("proto: ListTransferRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListTransferRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Transfer{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfers
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
func (m *AddTransferReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfers
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
			return fmt.Errorf("proto: AddTransferReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddTransferReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transfers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transfers = append(m.Transfers, &Transfer{})
			if err := m.Transfers[len(m.Transfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfers
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
func (m *AddTransferRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfers
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
			return fmt.Errorf("proto: AddTransferRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddTransferRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Item", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Item == nil {
				m.Item = &v1.ClientBind{}
			}
			if err := m.Item.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransfers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTransfers
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
func skipTransfers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransfers
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
					return 0, ErrIntOverflowTransfers
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTransfers
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
				return 0, ErrInvalidLengthTransfers
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTransfers
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTransfers
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTransfers(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTransfers
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTransfers = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransfers   = fmt.Errorf("proto: integer overflow")
)