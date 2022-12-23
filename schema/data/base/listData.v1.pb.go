// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/data/base/listData.v1.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/lists/base"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type ListData struct {
	Value *base.AnyDataList `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ListData) Reset()         { *m = ListData{} }
func (m *ListData) String() string { return proto.CompactTextString(m) }
func (*ListData) ProtoMessage()    {}
func (*ListData) Descriptor() ([]byte, []int) {
	return fileDescriptor_543f32b22fd81474, []int{0}
}
func (m *ListData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListData.Merge(m, src)
}
func (m *ListData) XXX_Size() int {
	return m.Size()
}
func (m *ListData) XXX_DiscardUnknown() {
	xxx_messageInfo_ListData.DiscardUnknown(m)
}

var xxx_messageInfo_ListData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListData)(nil), "data.ListData")
}

func init() {
	proto.RegisterFile("schema/data/base/listData.v1.proto", fileDescriptor_543f32b22fd81474)
}

var fileDescriptor_543f32b22fd81474 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0xc9, 0x2c,
	0x2e, 0x71, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x49, 0x4a, 0x29, 0x43, 0x55, 0x82, 0xe4, 0x8b, 0x21, 0x4a, 0x13, 0xf3, 0x2a, 0x41, 0x2a,
	0x7d, 0x32, 0x8b, 0x4b, 0x20, 0x4a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10,
	0x0b, 0x22, 0xaa, 0x64, 0xc5, 0xc5, 0xe1, 0x03, 0x35, 0x55, 0x48, 0x83, 0x8b, 0xb5, 0x2c, 0x31,
	0xa7, 0x34, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x48, 0x0f, 0x6c, 0x9e, 0x9e, 0x23,
	0xc2, 0xa8, 0x20, 0x88, 0x02, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0x53, 0x18, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b, 0x23, 0x39, 0x3f, 0x57, 0x0f, 0xe4, 0x36,
	0x27, 0x7e, 0x98, 0xf1, 0x61, 0x86, 0x01, 0x20, 0x1b, 0x03, 0x18, 0xa3, 0xf4, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x1d, 0x8b, 0x8b, 0x53, 0x4b, 0x7c, 0x13, 0xf3,
	0x4a, 0x72, 0x52, 0xf5, 0x73, 0xf3, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0xf5, 0xd1, 0xfd, 0xbd, 0x88,
	0x89, 0xd9, 0x25, 0x22, 0x62, 0x15, 0x13, 0x0b, 0xc8, 0x98, 0x53, 0x10, 0xea, 0x11, 0x93, 0x00,
	0x88, 0x8a, 0x71, 0x0f, 0x70, 0xf2, 0x4d, 0x2d, 0x49, 0x04, 0xa9, 0x7d, 0x05, 0x91, 0x49, 0x62,
	0x03, 0xfb, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x75, 0x6e, 0xda, 0x8d, 0x40, 0x01, 0x00,
	0x00,
}

func (m *ListData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListDataV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListDataV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovListDataV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovListDataV1(uint64(l))
	}
	return n
}

func sovListDataV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListDataV1(x uint64) (n int) {
	return sovListDataV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListDataV1
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
			return fmt.Errorf("proto: ListData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListDataV1
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
				return ErrInvalidLengthListDataV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListDataV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &base.AnyDataList{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListDataV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListDataV1
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
func skipListDataV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListDataV1
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
					return 0, ErrIntOverflowListDataV1
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
					return 0, ErrIntOverflowListDataV1
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
				return 0, ErrInvalidLengthListDataV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListDataV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListDataV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListDataV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListDataV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListDataV1 = fmt.Errorf("proto: unexpected end of group")
)
