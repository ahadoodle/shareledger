// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gentlemint/level_fee.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type LevelFee struct {
	Level   string        `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Fee     types.DecCoin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee"`
	Creator string        `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *LevelFee) Reset()         { *m = LevelFee{} }
func (m *LevelFee) String() string { return proto.CompactTextString(m) }
func (*LevelFee) ProtoMessage()    {}
func (*LevelFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_b729e670765c5582, []int{0}
}
func (m *LevelFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LevelFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LevelFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LevelFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LevelFee.Merge(m, src)
}
func (m *LevelFee) XXX_Size() int {
	return m.Size()
}
func (m *LevelFee) XXX_DiscardUnknown() {
	xxx_messageInfo_LevelFee.DiscardUnknown(m)
}

var xxx_messageInfo_LevelFee proto.InternalMessageInfo

func (m *LevelFee) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LevelFee) GetFee() types.DecCoin {
	if m != nil {
		return m.Fee
	}
	return types.DecCoin{}
}

func (m *LevelFee) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type LevelFeeDetail struct {
	Level        string        `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Creator      string        `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	OriginalFee  types.DecCoin `protobuf:"bytes,2,opt,name=originalFee,proto3" json:"originalFee"`
	ConvertedFee *types.Coin   `protobuf:"bytes,4,opt,name=convertedFee,proto3" json:"convertedFee,omitempty"`
}

func (m *LevelFeeDetail) Reset()         { *m = LevelFeeDetail{} }
func (m *LevelFeeDetail) String() string { return proto.CompactTextString(m) }
func (*LevelFeeDetail) ProtoMessage()    {}
func (*LevelFeeDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_b729e670765c5582, []int{1}
}
func (m *LevelFeeDetail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LevelFeeDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LevelFeeDetail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LevelFeeDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LevelFeeDetail.Merge(m, src)
}
func (m *LevelFeeDetail) XXX_Size() int {
	return m.Size()
}
func (m *LevelFeeDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_LevelFeeDetail.DiscardUnknown(m)
}

var xxx_messageInfo_LevelFeeDetail proto.InternalMessageInfo

func (m *LevelFeeDetail) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LevelFeeDetail) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *LevelFeeDetail) GetOriginalFee() types.DecCoin {
	if m != nil {
		return m.OriginalFee
	}
	return types.DecCoin{}
}

func (m *LevelFeeDetail) GetConvertedFee() *types.Coin {
	if m != nil {
		return m.ConvertedFee
	}
	return nil
}

func init() {
	proto.RegisterType((*LevelFee)(nil), "shareledger.gentlemint.LevelFee")
	proto.RegisterType((*LevelFeeDetail)(nil), "shareledger.gentlemint.LevelFeeDetail")
}

func init() { proto.RegisterFile("gentlemint/level_fee.proto", fileDescriptor_b729e670765c5582) }

var fileDescriptor_b729e670765c5582 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0x86, 0xa7, 0x1f, 0x7c, 0xfe, 0x14, 0xe3, 0x62, 0x42, 0xcc, 0x48, 0x4c, 0x25, 0xac, 0x58,
	0xb5, 0x41, 0xdc, 0xba, 0x41, 0xe2, 0x4a, 0x37, 0x2c, 0xdd, 0x98, 0x99, 0x72, 0x28, 0x4d, 0x86,
	0x1e, 0xd2, 0xa9, 0x44, 0xef, 0xc2, 0x9b, 0x32, 0x61, 0xc9, 0xd2, 0x95, 0x31, 0x70, 0x23, 0xa6,
	0x1d, 0x89, 0x63, 0x22, 0x0b, 0x77, 0xe7, 0x34, 0xef, 0x79, 0x9e, 0x26, 0x2f, 0x6d, 0x29, 0x30,
	0x2e, 0x87, 0x99, 0x36, 0x4e, 0xe4, 0xb0, 0x80, 0xfc, 0x61, 0x02, 0xc0, 0xe7, 0x16, 0x1d, 0xc6,
	0x27, 0xc5, 0x34, 0xb5, 0x90, 0xc3, 0x58, 0x81, 0xe5, 0xdf, 0xb9, 0x16, 0x93, 0x58, 0xcc, 0xb0,
	0x10, 0x59, 0x5a, 0x80, 0x58, 0xf4, 0x32, 0x70, 0x69, 0x4f, 0x48, 0xd4, 0xa6, 0xbc, 0x6b, 0x35,
	0x15, 0x2a, 0x0c, 0xa3, 0xf0, 0x53, 0xf9, 0xda, 0x99, 0xd3, 0x83, 0x5b, 0x2f, 0xb8, 0x01, 0x88,
	0x9b, 0xf4, 0x7f, 0x90, 0x25, 0xa4, 0x4d, 0xba, 0x87, 0xa3, 0x72, 0x89, 0x2f, 0x69, 0x6d, 0x02,
	0x90, 0xfc, 0x6b, 0x93, 0x6e, 0xe3, 0xe2, 0x8c, 0x97, 0x16, 0xee, 0x2d, 0xfc, 0xcb, 0xc2, 0x87,
	0x20, 0xaf, 0x51, 0x9b, 0x41, 0x7d, 0xf9, 0x7e, 0x1e, 0x8d, 0x7c, 0x3c, 0x4e, 0xe8, 0xbe, 0xb4,
	0x90, 0x3a, 0xb4, 0x49, 0x2d, 0xd0, 0xb6, 0x6b, 0xe7, 0x95, 0xd0, 0xe3, 0xad, 0x72, 0x08, 0x2e,
	0xd5, 0xf9, 0x0e, 0xf1, 0x4e, 0x44, 0x3c, 0xa4, 0x0d, 0xb4, 0x5a, 0x69, 0x93, 0x7a, 0xc8, 0x1f,
	0xbe, 0x56, 0x3d, 0x8b, 0xaf, 0xe8, 0x91, 0x44, 0xb3, 0x00, 0xeb, 0x60, 0xec, 0x31, 0xf5, 0x80,
	0x39, 0xfd, 0x15, 0xe3, 0x19, 0xa3, 0x1f, 0xf1, 0xc1, 0xdd, 0x72, 0xcd, 0xc8, 0x6a, 0xcd, 0xc8,
	0xc7, 0x9a, 0x91, 0x97, 0x0d, 0x8b, 0x56, 0x1b, 0x16, 0xbd, 0x6d, 0x58, 0x74, 0xdf, 0x57, 0xda,
	0x4d, 0x1f, 0x33, 0x2e, 0x71, 0x26, 0x42, 0x59, 0x56, 0x1b, 0x25, 0x2a, 0xb5, 0x89, 0x27, 0x51,
	0x29, 0xd8, 0x3d, 0xcf, 0xa1, 0xc8, 0xf6, 0x42, 0x1f, 0xfd, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x64, 0x60, 0x32, 0xc1, 0xfb, 0x01, 0x00, 0x00,
}

func (m *LevelFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LevelFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LevelFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLevelFee(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLevelFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Level) > 0 {
		i -= len(m.Level)
		copy(dAtA[i:], m.Level)
		i = encodeVarintLevelFee(dAtA, i, uint64(len(m.Level)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LevelFeeDetail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LevelFeeDetail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LevelFeeDetail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConvertedFee != nil {
		{
			size, err := m.ConvertedFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLevelFee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLevelFee(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.OriginalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLevelFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Level) > 0 {
		i -= len(m.Level)
		copy(dAtA[i:], m.Level)
		i = encodeVarintLevelFee(dAtA, i, uint64(len(m.Level)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLevelFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovLevelFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LevelFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovLevelFee(uint64(l))
	}
	l = m.Fee.Size()
	n += 1 + l + sovLevelFee(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLevelFee(uint64(l))
	}
	return n
}

func (m *LevelFeeDetail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovLevelFee(uint64(l))
	}
	l = m.OriginalFee.Size()
	n += 1 + l + sovLevelFee(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLevelFee(uint64(l))
	}
	if m.ConvertedFee != nil {
		l = m.ConvertedFee.Size()
		n += 1 + l + sovLevelFee(uint64(l))
	}
	return n
}

func sovLevelFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLevelFee(x uint64) (n int) {
	return sovLevelFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LevelFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLevelFee
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
			return fmt.Errorf("proto: LevelFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LevelFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLevelFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLevelFee
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
func (m *LevelFeeDetail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLevelFee
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
			return fmt.Errorf("proto: LevelFeeDetail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LevelFeeDetail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OriginalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConvertedFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLevelFee
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
				return ErrInvalidLengthLevelFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLevelFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConvertedFee == nil {
				m.ConvertedFee = &types.Coin{}
			}
			if err := m.ConvertedFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLevelFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLevelFee
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
func skipLevelFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLevelFee
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
					return 0, ErrIntOverflowLevelFee
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
					return 0, ErrIntOverflowLevelFee
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
				return 0, ErrInvalidLengthLevelFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLevelFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLevelFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLevelFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLevelFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLevelFee = fmt.Errorf("proto: unexpected end of group")
)