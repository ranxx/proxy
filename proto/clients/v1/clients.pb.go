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

type Client struct {
	Id      int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Machine string   `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	Account []string `protobuf:"bytes,3,rep,name=account,proto3" json:"account,omitempty"`
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

func (m *Client) GetId() int32 {
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

func (m *Client) GetAccount() []string {
	if m != nil {
		return m.Account
	}
	return nil
}

type ListClientReq struct {
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

type ListClientRsp struct {
	Items []*Client `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
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

func (m *ListClientRsp) GetItems() []*Client {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "v1.Client")
	proto.RegisterType((*ListClientReq)(nil), "v1.ListClientReq")
	proto.RegisterType((*ListClientRsp)(nil), "v1.ListClientRsp")
}

func init() { proto.RegisterFile("clients.proto", fileDescriptor_6c7b36ecb5ad4a28) }

var fileDescriptor_6c7b36ecb5ad4a28 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbd, 0x8e, 0xda, 0x40,
	0x1c, 0xc4, 0x59, 0x23, 0x83, 0x58, 0x44, 0x48, 0x56, 0x29, 0x2c, 0x2b, 0x58, 0x96, 0x9b, 0xd0,
	0xe0, 0xc5, 0x4e, 0x97, 0x2a, 0x24, 0x2d, 0x4d, 0x28, 0xa3, 0x34, 0x8b, 0x31, 0x66, 0x23, 0x7b,
	0x77, 0xe3, 0x5d, 0x2c, 0xa7, 0x4d, 0x95, 0x32, 0x22, 0x6f, 0x90, 0x32, 0x91, 0x92, 0xf2, 0x5e,
	0xe1, 0x4a, 0xa4, 0x6b, 0xae, 0x3c, 0xc1, 0x7d, 0xbc, 0xc6, 0xc9, 0x1f, 0xdc, 0x89, 0xa3, 0xb2,
	0xfe, 0xbf, 0x19, 0xcf, 0x48, 0xb3, 0xb0, 0x17, 0xc4, 0x34, 0x64, 0x4a, 0xba, 0x22, 0xe5, 0x8a,
	0x23, 0x2d, 0xf3, 0xcc, 0x57, 0x11, 0xe7, 0x51, 0x1c, 0x62, 0x22, 0x28, 0x26, 0x8c, 0x71, 0x45,
	0x14, 0xe5, 0xac, 0x76, 0x98, 0xd5, 0x27, 0x18, 0x45, 0x21, 0x1b, 0x71, 0x11, 0x32, 0x22, 0x68,
	0xe6, 0x63, 0x2e, 0x4a, 0xcf, 0xa9, 0xdf, 0x99, 0xc2, 0xd6, 0x87, 0xb2, 0x02, 0x3d, 0x83, 0x1a,
	0x5d, 0x18, 0xc0, 0x06, 0x43, 0x7d, 0xa6, 0xd1, 0x05, 0x32, 0x60, 0x3b, 0x21, 0xc1, 0x8a, 0xb2,
	0xd0, 0xd0, 0x6c, 0x30, 0xec, 0xcc, 0x0e, 0x67, 0xa1, 0x90, 0x20, 0xe0, 0x6b, 0xa6, 0x8c, 0xa6,
	0xdd, 0x2c, 0x94, 0xfa, 0x74, 0xfa, 0xb0, 0x37, 0xa5, 0x52, 0x55, 0x89, 0xb3, 0xf0, 0xab, 0xe3,
	0x1d, 0x01, 0x29, 0x90, 0x0d, 0x75, 0xaa, 0xc2, 0x44, 0x1a, 0xc0, 0x6e, 0x0e, 0xbb, 0x3e, 0x74,
	0x33, 0xcf, 0xad, 0xd5, 0x4a, 0xf0, 0x3f, 0xc3, 0x76, 0x05, 0x24, 0xfa, 0x08, 0xe1, 0xe3, 0xdf,
	0xe8, 0x45, 0xe1, 0x3d, 0x8a, 0x37, 0x9f, 0x22, 0x29, 0x9c, 0xc1, 0x66, 0xd2, 0x85, 0x9d, 0x9b,
	0x7f, 0xbf, 0x6f, 0x7f, 0x9c, 0xdd, 0xfd, 0xfd, 0xff, 0xfd, 0xe2, 0xfa, 0x97, 0xd6, 0x41, 0x6d,
	0x5c, 0x0d, 0xf9, 0xfe, 0x0f, 0xd8, 0x4c, 0x96, 0x08, 0x3f, 0x74, 0x38, 0x03, 0xa8, 0xa7, 0x84,
	0xe5, 0xb9, 0xf9, 0x32, 0x25, 0xcc, 0x95, 0x8a, 0xa4, 0xef, 0x96, 0x3c, 0x4f, 0x08, 0x8d, 0xdd,
	0x80, 0x27, 0xbe, 0xee, 0xb9, 0x63, 0x77, 0xec, 0xf4, 0xcb, 0xb9, 0x33, 0xaf, 0x8e, 0x91, 0xfe,
	0x73, 0x22, 0x44, 0x4c, 0x83, 0x72, 0x4b, 0xfc, 0x45, 0x72, 0xf6, 0xf6, 0x84, 0x9c, 0xef, 0x2c,
	0xb0, 0xdd, 0x59, 0xe0, 0x6a, 0x67, 0x81, 0x9f, 0x7b, 0xab, 0xb1, 0xdd, 0x5b, 0x8d, 0xcb, 0xbd,
	0xd5, 0xf8, 0xf4, 0x3a, 0xa2, 0x6a, 0xb5, 0x9e, 0x17, 0xf9, 0xb8, 0xec, 0xc6, 0x22, 0xe5, 0xf9,
	0x37, 0x5c, 0xbe, 0xc8, 0xa1, 0x01, 0x67, 0xde, 0xbc, 0x55, 0x92, 0x37, 0xf7, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x36, 0x89, 0x08, 0x3f, 0x06, 0x02, 0x00, 0x00,
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

// ClientsServer is the server API for Clients service.
type ClientsServer interface {
	// 客户端列表
	ListClient(context.Context, *ListClientReq) (*ListClientRsp, error)
}

// UnimplementedClientsServer can be embedded to have forward compatible implementations.
type UnimplementedClientsServer struct {
}

func (*UnimplementedClientsServer) ListClient(ctx context.Context, req *ListClientReq) (*ListClientRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClient not implemented")
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

var _Clients_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Clients",
	HandlerType: (*ClientsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClient",
			Handler:    _Clients_ListClient_Handler,
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
	if len(m.Account) > 0 {
		for _, s := range m.Account {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
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
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
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
	if len(m.Account) > 0 {
		for _, s := range m.Account {
			l = len(s)
			n += 1 + l + sovClients(uint64(l))
		}
	}
	return n
}

func (m *ListClientReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListClientRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovClients(uint64(l))
		}
	}
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
				m.Id |= int32(b&0x7F) << shift
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = append(m.Account, string(dAtA[iNdEx:postIndex]))
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
