// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: booking/tx.proto

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

type MsgCreateBooking struct {
	Booker   string `protobuf:"bytes,1,opt,name=booker,proto3" json:"booker,omitempty"`
	UUID     string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Duration int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *MsgCreateBooking) Reset()         { *m = MsgCreateBooking{} }
func (m *MsgCreateBooking) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBooking) ProtoMessage()    {}
func (*MsgCreateBooking) Descriptor() ([]byte, []int) {
	return fileDescriptor_e243c29ad9ae20da, []int{0}
}
func (m *MsgCreateBooking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBooking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBooking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBooking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBooking.Merge(m, src)
}
func (m *MsgCreateBooking) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBooking) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBooking.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBooking proto.InternalMessageInfo

func (m *MsgCreateBooking) GetBooker() string {
	if m != nil {
		return m.Booker
	}
	return ""
}

func (m *MsgCreateBooking) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *MsgCreateBooking) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type MsgCreateBookingResponse struct {
}

func (m *MsgCreateBookingResponse) Reset()         { *m = MsgCreateBookingResponse{} }
func (m *MsgCreateBookingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBookingResponse) ProtoMessage()    {}
func (*MsgCreateBookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e243c29ad9ae20da, []int{1}
}
func (m *MsgCreateBookingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBookingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBookingResponse.Merge(m, src)
}
func (m *MsgCreateBookingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBookingResponse proto.InternalMessageInfo

type MsgCompleteBooking struct {
	Booker string `protobuf:"bytes,1,opt,name=booker,proto3" json:"booker,omitempty"`
	BookID string `protobuf:"bytes,2,opt,name=bookID,proto3" json:"bookID,omitempty"`
}

func (m *MsgCompleteBooking) Reset()         { *m = MsgCompleteBooking{} }
func (m *MsgCompleteBooking) String() string { return proto.CompactTextString(m) }
func (*MsgCompleteBooking) ProtoMessage()    {}
func (*MsgCompleteBooking) Descriptor() ([]byte, []int) {
	return fileDescriptor_e243c29ad9ae20da, []int{2}
}
func (m *MsgCompleteBooking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCompleteBooking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCompleteBooking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCompleteBooking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCompleteBooking.Merge(m, src)
}
func (m *MsgCompleteBooking) XXX_Size() int {
	return m.Size()
}
func (m *MsgCompleteBooking) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCompleteBooking.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCompleteBooking proto.InternalMessageInfo

func (m *MsgCompleteBooking) GetBooker() string {
	if m != nil {
		return m.Booker
	}
	return ""
}

func (m *MsgCompleteBooking) GetBookID() string {
	if m != nil {
		return m.BookID
	}
	return ""
}

type MsgCompleteBookingResponse struct {
}

func (m *MsgCompleteBookingResponse) Reset()         { *m = MsgCompleteBookingResponse{} }
func (m *MsgCompleteBookingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCompleteBookingResponse) ProtoMessage()    {}
func (*MsgCompleteBookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e243c29ad9ae20da, []int{3}
}
func (m *MsgCompleteBookingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCompleteBookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCompleteBookingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCompleteBookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCompleteBookingResponse.Merge(m, src)
}
func (m *MsgCompleteBookingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCompleteBookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCompleteBookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCompleteBookingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateBooking)(nil), "shareledger.booking.MsgCreateBooking")
	proto.RegisterType((*MsgCreateBookingResponse)(nil), "shareledger.booking.MsgCreateBookingResponse")
	proto.RegisterType((*MsgCompleteBooking)(nil), "shareledger.booking.MsgCompleteBooking")
	proto.RegisterType((*MsgCompleteBookingResponse)(nil), "shareledger.booking.MsgCompleteBookingResponse")
}

func init() { proto.RegisterFile("booking/tx.proto", fileDescriptor_e243c29ad9ae20da) }

var fileDescriptor_e243c29ad9ae20da = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xca, 0xcf, 0xcf,
	0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2e, 0xce,
	0x48, 0x2c, 0x4a, 0xcd, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd2, 0x83, 0xca, 0x2a, 0x45, 0x71, 0x09,
	0xf8, 0x16, 0xa7, 0x3b, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x3a, 0x41, 0xc4, 0x84, 0xc4, 0xb8, 0xd8,
	0x40, 0xd2, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x10, 0x17,
	0x4b, 0x68, 0xa8, 0xa7, 0x8b, 0x04, 0x13, 0x58, 0x14, 0xcc, 0x16, 0x92, 0xe2, 0xe2, 0x48, 0x29,
	0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf3, 0x95,
	0xa4, 0xb8, 0x24, 0xd0, 0xcd, 0x0e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x72, 0xe1,
	0x12, 0x02, 0xc9, 0xe5, 0xe7, 0x16, 0xe4, 0xa4, 0x12, 0xb6, 0x19, 0x2a, 0x0e, 0xb7, 0x1b, 0xca,
	0x53, 0x92, 0xe1, 0x92, 0xc2, 0x34, 0x05, 0x66, 0x87, 0xd1, 0x4d, 0x46, 0x2e, 0x66, 0xdf, 0xe2,
	0x74, 0xa1, 0x54, 0x2e, 0x5e, 0x54, 0x0f, 0xaa, 0xea, 0x61, 0x09, 0x0a, 0x3d, 0x74, 0xb7, 0x4a,
	0xe9, 0x12, 0xa5, 0x0c, 0x66, 0x9d, 0x50, 0x36, 0x17, 0x3f, 0xba, 0x7f, 0xd4, 0x71, 0x9a, 0x80,
	0xaa, 0x50, 0x4a, 0x9f, 0x48, 0x85, 0x30, 0xcb, 0x9c, 0xbc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0x6c, 0x68, 0x11, 0x28, 0x15, 0x20, 0x19, 0xaf, 0x5f, 0xa1, 0x0f, 0x4f, 0x1b, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xf4, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x80, 0x79, 0x0b,
	0xdc, 0x33, 0x02, 0x00, 0x00,
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
	CreateBooking(ctx context.Context, in *MsgCreateBooking, opts ...grpc.CallOption) (*MsgCreateBookingResponse, error)
	CompleteBooking(ctx context.Context, in *MsgCompleteBooking, opts ...grpc.CallOption) (*MsgCompleteBookingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBooking(ctx context.Context, in *MsgCreateBooking, opts ...grpc.CallOption) (*MsgCreateBookingResponse, error) {
	out := new(MsgCreateBookingResponse)
	err := c.cc.Invoke(ctx, "/shareledger.booking.Msg/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CompleteBooking(ctx context.Context, in *MsgCompleteBooking, opts ...grpc.CallOption) (*MsgCompleteBookingResponse, error) {
	out := new(MsgCompleteBookingResponse)
	err := c.cc.Invoke(ctx, "/shareledger.booking.Msg/CompleteBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateBooking(context.Context, *MsgCreateBooking) (*MsgCreateBookingResponse, error)
	CompleteBooking(context.Context, *MsgCompleteBooking) (*MsgCompleteBookingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBooking(ctx context.Context, req *MsgCreateBooking) (*MsgCreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (*UnimplementedMsgServer) CompleteBooking(ctx context.Context, req *MsgCompleteBooking) (*MsgCompleteBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteBooking not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.booking.Msg/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBooking(ctx, req.(*MsgCreateBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CompleteBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCompleteBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CompleteBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.booking.Msg/CompleteBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CompleteBooking(ctx, req.(*MsgCompleteBooking))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shareledger.booking.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _Msg_CreateBooking_Handler,
		},
		{
			MethodName: "CompleteBooking",
			Handler:    _Msg_CompleteBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking/tx.proto",
}

func (m *MsgCreateBooking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBooking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBooking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Duration != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Booker) > 0 {
		i -= len(m.Booker)
		copy(dAtA[i:], m.Booker)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Booker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBookingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBookingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBookingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCompleteBooking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCompleteBooking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCompleteBooking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BookID) > 0 {
		i -= len(m.BookID)
		copy(dAtA[i:], m.BookID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BookID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Booker) > 0 {
		i -= len(m.Booker)
		copy(dAtA[i:], m.Booker)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Booker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCompleteBookingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCompleteBookingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCompleteBookingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreateBooking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Booker)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovTx(uint64(m.Duration))
	}
	return n
}

func (m *MsgCreateBookingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCompleteBooking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Booker)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BookID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCompleteBookingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateBooking) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBooking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBooking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Booker", wireType)
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
			m.Booker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
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
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int64(b&0x7F) << shift
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
func (m *MsgCreateBookingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBookingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBookingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgCompleteBooking) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCompleteBooking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCompleteBooking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Booker", wireType)
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
			m.Booker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookID", wireType)
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
			m.BookID = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCompleteBookingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCompleteBookingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCompleteBookingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
