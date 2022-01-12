// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: clients.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ClientStatus int32

const (
	ClientStatus__cilents ClientStatus = 0
	// 停止
	ClientStatus_Stop ClientStatus = 1
	// 运行中
	ClientStatus_Run ClientStatus = 2
)

var ClientStatus_name = map[int32]string{
	0: "_cilents",
	1: "Stop",
	2: "Run",
}

var ClientStatus_value = map[string]int32{
	"_cilents": 0,
	"Stop":     1,
	"Run":      2,
}

func (x ClientStatus) String() string {
	return proto.EnumName(ClientStatus_name, int32(x))
}

func (ClientStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{0}
}

type Client struct {
	Id      int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Machine string       `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	Status  ClientStatus `protobuf:"varint,3,opt,name=status,proto3,enum=v1.ClientStatus" json:"status,omitempty"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{0}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Client.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return m.Size()
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Client) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *Client) GetStatus() ClientStatus {
	if m != nil {
		return m.Status
	}
	return ClientStatus__cilents
}

type ListClientReq struct {
	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *ListClientReq) Reset()         { *m = ListClientReq{} }
func (m *ListClientReq) String() string { return proto.CompactTextString(m) }
func (*ListClientReq) ProtoMessage()    {}
func (*ListClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{1}
}
func (m *ListClientReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListClientReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClientReq.Merge(m, src)
}
func (m *ListClientReq) XXX_Size() int {
	return m.Size()
}
func (m *ListClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListClientReq proto.InternalMessageInfo

func (m *ListClientReq) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListClientReq) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListClientRsp struct {
	Total int64     `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Client `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *ListClientRsp) Reset()         { *m = ListClientRsp{} }
func (m *ListClientRsp) String() string { return proto.CompactTextString(m) }
func (*ListClientRsp) ProtoMessage()    {}
func (*ListClientRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2}
}
func (m *ListClientRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListClientRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListClientRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListClientRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClientRsp.Merge(m, src)
}
func (m *ListClientRsp) XXX_Size() int {
	return m.Size()
}
func (m *ListClientRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClientRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ListClientRsp proto.InternalMessageInfo

func (m *ListClientRsp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListClientRsp) GetItems() []*Client {
	if m != nil {
		return m.Items
	}
	return nil
}

type StopClientReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *StopClientReq) Reset()         { *m = StopClientReq{} }
func (m *StopClientReq) String() string { return proto.CompactTextString(m) }
func (*StopClientReq) ProtoMessage()    {}
func (*StopClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3}
}
func (m *StopClientReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StopClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StopClientReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StopClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopClientReq.Merge(m, src)
}
func (m *StopClientReq) XXX_Size() int {
	return m.Size()
}
func (m *StopClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopClientReq proto.InternalMessageInfo

func (m *StopClientReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type StopClientRsp struct {
}

func (m *StopClientRsp) Reset()         { *m = StopClientRsp{} }
func (m *StopClientRsp) String() string { return proto.CompactTextString(m) }
func (*StopClientRsp) ProtoMessage()    {}
func (*StopClientRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{4}
}
func (m *StopClientRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StopClientRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StopClientRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StopClientRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopClientRsp.Merge(m, src)
}
func (m *StopClientRsp) XXX_Size() int {
	return m.Size()
}
func (m *StopClientRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_StopClientRsp.DiscardUnknown(m)
}

var xxx_messageInfo_StopClientRsp proto.InternalMessageInfo

type DeleteClientReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *DeleteClientReq) Reset()         { *m = DeleteClientReq{} }
func (m *DeleteClientReq) String() string { return proto.CompactTextString(m) }
func (*DeleteClientReq) ProtoMessage()    {}
func (*DeleteClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{5}
}
func (m *DeleteClientReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteClientReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeleteClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClientReq.Merge(m, src)
}
func (m *DeleteClientReq) XXX_Size() int {
	return m.Size()
}
func (m *DeleteClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClientReq proto.InternalMessageInfo

func (m *DeleteClientReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteClientRsp struct {
}

func (m *DeleteClientRsp) Reset()         { *m = DeleteClientRsp{} }
func (m *DeleteClientRsp) String() string { return proto.CompactTextString(m) }
func (*DeleteClientRsp) ProtoMessage()    {}
func (*DeleteClientRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{6}
}
func (m *DeleteClientRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteClientRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteClientRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeleteClientRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClientRsp.Merge(m, src)
}
func (m *DeleteClientRsp) XXX_Size() int {
	return m.Size()
}
func (m *DeleteClientRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClientRsp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClientRsp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("v1.ClientStatus", ClientStatus_name, ClientStatus_value)
	proto.RegisterType((*Client)(nil), "v1.Client")
	proto.RegisterType((*ListClientReq)(nil), "v1.ListClientReq")
	proto.RegisterType((*ListClientRsp)(nil), "v1.ListClientRsp")
	proto.RegisterType((*StopClientReq)(nil), "v1.StopClientReq")
	proto.RegisterType((*StopClientRsp)(nil), "v1.StopClientRsp")
	proto.RegisterType((*DeleteClientReq)(nil), "v1.DeleteClientReq")
	proto.RegisterType((*DeleteClientRsp)(nil), "v1.DeleteClientRsp")
}

func init() { proto.RegisterFile("clients.proto", fileDescriptor_6c7b36ecb5ad4a28) }

var fileDescriptor_6c7b36ecb5ad4a28 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4f, 0x8b, 0xd3, 0x4e,
	0x18, 0x6e, 0xd2, 0x5f, 0xdb, 0xfd, 0xbd, 0x76, 0xfb, 0x67, 0x5c, 0xa4, 0x04, 0x37, 0x5b, 0x73,
	0xb1, 0x2c, 0x6c, 0x66, 0x1b, 0x6f, 0x82, 0x60, 0x55, 0xf0, 0xe2, 0xc5, 0x2c, 0x1e, 0x14, 0x41,
	0x66, 0xd3, 0xb4, 0x3b, 0x92, 0x66, 0xc6, 0xce, 0xb4, 0x54, 0xc4, 0x83, 0x9e, 0x7a, 0x12, 0xa9,
	0x07, 0x11, 0x3c, 0xe8, 0xe2, 0x41, 0x14, 0xf5, 0xe8, 0x57, 0xf0, 0xb8, 0xe0, 0xc5, 0xa3, 0xb4,
	0x76, 0xd5, 0x6f, 0x21, 0x99, 0x64, 0xed, 0xb6, 0xc5, 0x53, 0x78, 0x9f, 0x79, 0xde, 0xe7, 0x79,
	0xdf, 0x27, 0x33, 0xb0, 0xea, 0x05, 0xd4, 0x0f, 0xa5, 0xb0, 0x79, 0x97, 0x49, 0x86, 0xf4, 0x7e,
	0xdd, 0x38, 0xdd, 0x66, 0xac, 0x1d, 0xf8, 0x98, 0x70, 0x8a, 0x49, 0x18, 0x32, 0x49, 0x24, 0x65,
	0x61, 0xc2, 0x30, 0xe2, 0x8f, 0xb7, 0xd5, 0xf6, 0xc3, 0x2d, 0xc6, 0xfd, 0x90, 0x70, 0xda, 0x77,
	0x30, 0xe3, 0x8a, 0xb3, 0xcc, 0xb7, 0x6e, 0x43, 0xf6, 0xb2, 0xb2, 0x40, 0x05, 0xd0, 0x69, 0xb3,
	0xa2, 0x55, 0xb5, 0x5a, 0xda, 0xd5, 0x69, 0x13, 0x55, 0x20, 0xd7, 0x21, 0xde, 0x1e, 0x0d, 0xfd,
	0x8a, 0x5e, 0xd5, 0x6a, 0xff, 0xbb, 0x47, 0x25, 0xaa, 0x41, 0x56, 0x48, 0x22, 0x7b, 0xa2, 0x92,
	0xae, 0x6a, 0xb5, 0x82, 0x53, 0xb2, 0xfb, 0x75, 0x3b, 0x56, 0xd9, 0x51, 0xb8, 0x9b, 0x9c, 0x5b,
	0x17, 0x60, 0xf5, 0x1a, 0x15, 0x32, 0x3e, 0x73, 0xfd, 0x7b, 0xe8, 0x14, 0x64, 0x59, 0xab, 0x25,
	0x7c, 0x99, 0x18, 0x25, 0x15, 0x5a, 0x83, 0x4c, 0x40, 0x3b, 0x54, 0x2a, 0xab, 0xb4, 0x1b, 0x17,
	0xd6, 0xd5, 0xb9, 0x76, 0xc1, 0x23, 0x9a, 0x64, 0x92, 0x04, 0x49, 0x77, 0x5c, 0xa0, 0x2a, 0x64,
	0xa8, 0xf4, 0x3b, 0xa2, 0xa2, 0x57, 0xd3, 0xb5, 0x13, 0x0e, 0xcc, 0xc6, 0x71, 0xe3, 0x03, 0x6b,
	0x03, 0x56, 0x77, 0x24, 0xe3, 0xb3, 0x39, 0x16, 0x96, 0xb5, 0x8a, 0x73, 0x04, 0xc1, 0xad, 0x33,
	0x50, 0xbc, 0xe2, 0x07, 0xbe, 0xf4, 0xff, 0xdd, 0x53, 0x5e, 0xa0, 0x08, 0xbe, 0x89, 0x21, 0x7f,
	0x3c, 0x07, 0x94, 0x87, 0x95, 0x3b, 0x1e, 0x0d, 0xa2, 0x3f, 0x58, 0x4a, 0xa1, 0x15, 0xf8, 0x2f,
	0x32, 0x29, 0x69, 0x28, 0x07, 0x69, 0xb7, 0x17, 0x96, 0x74, 0xe7, 0x89, 0x0e, 0xb9, 0xb8, 0x43,
	0xa0, 0x1b, 0x00, 0xb3, 0x6d, 0x51, 0x39, 0xda, 0x62, 0x2e, 0x3c, 0x63, 0x11, 0x8a, 0xc6, 0x1c,
	0x35, 0xca, 0x50, 0x9c, 0x7e, 0xdc, 0x3f, 0x1c, 0x7e, 0xfe, 0xf9, 0xfe, 0xd3, 0x74, 0xf8, 0xfc,
	0xd7, 0xeb, 0xb7, 0x8f, 0xbf, 0xfe, 0x78, 0xa6, 0xe7, 0x50, 0x06, 0x07, 0x54, 0x48, 0x74, 0x1d,
	0x60, 0xb6, 0x5a, 0x2c, 0x3b, 0x97, 0x85, 0xb1, 0x08, 0x09, 0x6e, 0xad, 0x8f, 0x1a, 0x2b, 0x90,
	0x9d, 0x3e, 0x7a, 0x79, 0xf8, 0x61, 0x5f, 0xa9, 0xe5, 0x0d, 0xc0, 0x42, 0x32, 0x8e, 0x1f, 0xd0,
	0xe6, 0x43, 0x74, 0x13, 0xf2, 0xc7, 0x37, 0x47, 0x27, 0x23, 0x85, 0x85, 0xb8, 0x8c, 0x65, 0x50,
	0x70, 0x6b, 0x23, 0x16, 0x1e, 0xbe, 0xfa, 0xfd, 0xe2, 0x8d, 0x12, 0x2e, 0x6c, 0xe6, 0x71, 0x7c,
	0xc7, 0x95, 0xf4, 0xa5, 0x77, 0xda, 0xa8, 0xd1, 0x42, 0xf8, 0x6f, 0x2a, 0xd6, 0x3a, 0x64, 0xba,
	0x24, 0x1c, 0x0c, 0x8c, 0xb5, 0x2e, 0x09, 0x6d, 0x21, 0x49, 0xf7, 0x62, 0x8b, 0x0d, 0x3a, 0x84,
	0x06, 0xb6, 0xc7, 0x3a, 0x4e, 0xa6, 0x6e, 0x6f, 0xdb, 0xdb, 0x56, 0x51, 0x3d, 0x87, 0x7e, 0x3d,
	0xd1, 0x12, 0x4e, 0x89, 0x70, 0x1e, 0x50, 0x4f, 0xdd, 0x75, 0x7c, 0x57, 0xb0, 0xf0, 0xfc, 0x12,
	0xf2, 0x65, 0x6c, 0x6a, 0x07, 0x63, 0x53, 0xfb, 0x3e, 0x36, 0xb5, 0xa7, 0x13, 0x33, 0x75, 0x30,
	0x31, 0x53, 0xdf, 0x26, 0x66, 0xea, 0xd6, 0xd9, 0x36, 0x95, 0x7b, 0xbd, 0xdd, 0x48, 0x1f, 0x2b,
	0x6f, 0xcc, 0xbb, 0x6c, 0x70, 0x1f, 0xab, 0x17, 0x73, 0xe4, 0x80, 0xfb, 0xf5, 0xdd, 0xac, 0x42,
	0xce, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x94, 0x76, 0xf8, 0xbe, 0xa6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientsClient is the client API for Clients service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientsClient interface {
	// 客户端列表
	ListClient(ctx context.Context, in *ListClientReq, opts ...grpc.CallOption) (*ListClientRsp, error)
	// 停止
	StopClient(ctx context.Context, in *StopClientReq, opts ...grpc.CallOption) (*StopClientRsp, error)
	// 删除
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientRsp, error)
}

type clientsClient struct {
	cc *grpc.ClientConn
}

func NewClientsClient(cc *grpc.ClientConn) ClientsClient {
	return &clientsClient{cc}
}

func (c *clientsClient) ListClient(ctx context.Context, in *ListClientReq, opts ...grpc.CallOption) (*ListClientRsp, error) {
	out := new(ListClientRsp)
	err := c.cc.Invoke(ctx, "/v1.Clients/ListClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsClient) StopClient(ctx context.Context, in *StopClientReq, opts ...grpc.CallOption) (*StopClientRsp, error) {
	out := new(StopClientRsp)
	err := c.cc.Invoke(ctx, "/v1.Clients/StopClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientRsp, error) {
	out := new(DeleteClientRsp)
	err := c.cc.Invoke(ctx, "/v1.Clients/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientsServer is the server API for Clients service.
type ClientsServer interface {
	// 客户端列表
	ListClient(context.Context, *ListClientReq) (*ListClientRsp, error)
	// 停止
	StopClient(context.Context, *StopClientReq) (*StopClientRsp, error)
	// 删除
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientRsp, error)
}

// UnimplementedClientsServer can be embedded to have forward compatible implementations.
type UnimplementedClientsServer struct {
}

func (*UnimplementedClientsServer) ListClient(ctx context.Context, req *ListClientReq) (*ListClientRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClient not implemented")
}
func (*UnimplementedClientsServer) StopClient(ctx context.Context, req *StopClientReq) (*StopClientRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopClient not implemented")
}
func (*UnimplementedClientsServer) DeleteClient(ctx context.Context, req *DeleteClientReq) (*DeleteClientRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}

func RegisterClientsServer(s *grpc.Server, srv ClientsServer) {
	s.RegisterService(&_Clients_serviceDesc, srv)
}

func _Clients_ListClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).ListClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Clients/ListClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).ListClient(ctx, req.(*ListClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clients_StopClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).StopClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Clients/StopClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).StopClient(ctx, req.(*StopClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clients_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Clients/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clients_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Clients",
	HandlerType: (*ClientsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClient",
			Handler:    _Clients_ListClient_Handler,
		},
		{
			MethodName: "StopClient",
			Handler:    _Clients_StopClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Clients_DeleteClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clients.proto",
}

func (m *Client) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Client) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Id))
	}
	if len(m.Machine) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClients(dAtA, i, uint64(len(m.Machine)))
		i += copy(dAtA[i:], m.Machine)
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *ListClientReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListClientReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Offset))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Limit))
	}
	return i, nil
}

func (m *ListClientRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListClientRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Total))
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintClients(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StopClientReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StopClientReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *StopClientRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StopClientRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DeleteClientReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteClientReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClients(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *DeleteClientRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteClientRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintClients(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Client) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClients(uint64(m.Id))
	}
	l = len(m.Machine)
	if l > 0 {
		n += 1 + l + sovClients(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovClients(uint64(m.Status))
	}
	return n
}

func (m *ListClientReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovClients(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovClients(uint64(m.Limit))
	}
	return n
}

func (m *ListClientRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovClients(uint64(m.Total))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovClients(uint64(l))
		}
	}
	return n
}

func (m *StopClientReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClients(uint64(m.Id))
	}
	return n
}

func (m *StopClientRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DeleteClientReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClients(uint64(m.Id))
	}
	return n
}

func (m *DeleteClientRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovClients(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClients(x uint64) (n int) {
	return sovClients(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Client) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: Client: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Client: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Machine", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
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
				return ErrInvalidLengthClients
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClients
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Machine = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ClientStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func (m *ListClientReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: ListClientReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListClientReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func (m *ListClientRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: ListClientRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListClientRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
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
				return ErrInvalidLengthClients
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClients
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Client{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func (m *StopClientReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: StopClientReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StopClientReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func (m *StopClientRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: StopClientRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StopClientRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func (m *DeleteClientReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: DeleteClientReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteClientReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClients
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func (m *DeleteClientRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClients
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
			return fmt.Errorf("proto: DeleteClientRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteClientRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClients(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClients
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClients
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
func skipClients(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClients
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
					return 0, ErrIntOverflowClients
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
					return 0, ErrIntOverflowClients
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
				return 0, ErrInvalidLengthClients
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthClients
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClients
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
				next, err := skipClients(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthClients
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
	ErrInvalidLengthClients = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClients   = fmt.Errorf("proto: integer overflow")
)
