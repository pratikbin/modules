// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/qualified/base/mutables.v1.proto

package base

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/lists/base"
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

type Mutables struct {
	PropertyList *base.List `protobuf:"bytes,1,opt,name=property_list,json=propertyList,proto3" json:"property_list,omitempty"`
}

func (m *Mutables) Reset()         { *m = Mutables{} }
func (m *Mutables) String() string { return proto.CompactTextString(m) }
func (*Mutables) ProtoMessage()    {}
func (*Mutables) Descriptor() ([]byte, []int) {
	return fileDescriptor_15cd80a375dd2556, []int{0}
}
func (m *Mutables) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mutables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mutables.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mutables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mutables.Merge(m, src)
}
func (m *Mutables) XXX_Size() int {
	return m.Size()
}
func (m *Mutables) XXX_DiscardUnknown() {
	xxx_messageInfo_Mutables.DiscardUnknown(m)
}

var xxx_messageInfo_Mutables proto.InternalMessageInfo

func (m *Mutables) GetPropertyList() *base.List {
	if m != nil {
		return m.PropertyList
	}
	return nil
}

func init() {
	proto.RegisterType((*Mutables)(nil), "qualified.Mutables")
}

func init() {
	proto.RegisterFile("schema/qualified/base/mutables.v1.proto", fileDescriptor_15cd80a375dd2556)
}

var fileDescriptor_15cd80a375dd2556 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x2c, 0x4d, 0xcc, 0xc9, 0x4c, 0xcb, 0x4c, 0x4d, 0xd1, 0x4f, 0x4a, 0x2c,
	0x4e, 0xd5, 0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0xca, 0x49, 0x2d, 0xd6, 0x2b, 0x33, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xab, 0x90, 0x92, 0x87, 0xea, 0xc9, 0xc9, 0x2c, 0x2e, 0x29,
	0x86, 0xa8, 0xf7, 0xc9, 0x2c, 0x2e, 0x81, 0xab, 0x55, 0xb2, 0xe1, 0xe2, 0xf0, 0x85, 0x1a, 0x20,
	0x64, 0xc0, 0xc5, 0x5b, 0x50, 0x94, 0x5f, 0x90, 0x5a, 0x54, 0x52, 0x19, 0x0f, 0xd2, 0x20, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xad, 0x07, 0xd6, 0xad, 0x07, 0xd2, 0x18, 0xc4, 0x03, 0x53,
	0x01, 0xe2, 0x39, 0x6d, 0x62, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e,
	0xde, 0xe4, 0xfc, 0x5c, 0x3d, 0xb8, 0x43, 0x9c, 0xf8, 0x61, 0xb6, 0x84, 0x19, 0x06, 0x80, 0x2c,
	0x0e, 0x60, 0x8c, 0x32, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77,
	0x2c, 0x2e, 0x4e, 0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0xcf, 0xcd, 0x4f, 0x29, 0xcd,
	0x49, 0x2d, 0xd6, 0xc7, 0xea, 0xdd, 0x45, 0x4c, 0xcc, 0x81, 0x11, 0x11, 0xab, 0x98, 0x38, 0x03,
	0x61, 0xc2, 0xa7, 0x90, 0xd8, 0x8f, 0x98, 0x44, 0xe1, 0xec, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4,
	0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x57, 0x48, 0x6a, 0x92, 0xd8, 0xc0, 0x3e, 0x37, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x35, 0x36, 0xf0, 0x34, 0x50, 0x01, 0x00, 0x00,
}

func (m *Mutables) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mutables) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mutables) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PropertyList != nil {
		{
			size, err := m.PropertyList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMutablesV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMutablesV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMutablesV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mutables) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PropertyList != nil {
		l = m.PropertyList.Size()
		n += 1 + l + sovMutablesV1(uint64(l))
	}
	return n
}

func sovMutablesV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMutablesV1(x uint64) (n int) {
	return sovMutablesV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mutables) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMutablesV1
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
			return fmt.Errorf("proto: Mutables: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mutables: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutablesV1
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
				return ErrInvalidLengthMutablesV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMutablesV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PropertyList == nil {
				m.PropertyList = &base.List{}
			}
			if err := m.PropertyList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMutablesV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMutablesV1
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
func skipMutablesV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMutablesV1
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
					return 0, ErrIntOverflowMutablesV1
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
					return 0, ErrIntOverflowMutablesV1
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
				return 0, ErrInvalidLengthMutablesV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMutablesV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMutablesV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMutablesV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMutablesV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMutablesV1 = fmt.Errorf("proto: unexpected end of group")
)
