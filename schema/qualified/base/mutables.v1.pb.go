// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/qualified/base/mutables.v1.proto

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

type Mutables struct {
	PropertyList *base.PropertyList `protobuf:"bytes,1,opt,name=property_list,json=propertyList,proto3" json:"property_list,omitempty"`
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

func init() {
	proto.RegisterType((*Mutables)(nil), "qualified.Mutables")
}

func init() {
	proto.RegisterFile("schema/qualified/base/mutables.v1.proto", fileDescriptor_15cd80a375dd2556)
}

var fileDescriptor_15cd80a375dd2556 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x2c, 0x4d, 0xcc, 0xc9, 0x4c, 0xcb, 0x4c, 0x4d, 0xd1, 0x4f, 0x4a, 0x2c,
	0x4e, 0xd5, 0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0xca, 0x49, 0x2d, 0xd6, 0x2b, 0x33, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xab, 0x90, 0x82, 0xe9, 0xc9, 0xc9, 0x2c, 0x2e, 0x29, 0x86,
	0xa8, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xf4, 0xc9, 0x2c, 0x2e, 0x81, 0xeb, 0x91,
	0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x92, 0x17, 0x17, 0x87,
	0x2f, 0xd4, 0x78, 0x21, 0x0b, 0x2e, 0x5e, 0x98, 0xd6, 0x78, 0x90, 0x71, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x60, 0xb3, 0xf5, 0x02, 0x90, 0x8c, 0x0d, 0xe2, 0x41, 0xb6, 0xc4,
	0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0xd3, 0x26, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0xe0, 0xe2, 0x4d, 0xce, 0xcf, 0xd5, 0x83, 0x3b, 0xda, 0x89, 0x1f, 0x66, 0x67, 0x98,
	0x61, 0x00, 0xc8, 0x19, 0x01, 0x8c, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x8e, 0xc5, 0xc5, 0xa9, 0x25, 0xbe, 0x89, 0x79, 0x25, 0x39, 0xa9, 0xfa, 0xb9,
	0xf9, 0x29, 0xa5, 0x39, 0xa9, 0xc5, 0xfa, 0x58, 0x83, 0x66, 0x11, 0x13, 0x73, 0x60, 0x44, 0xc4,
	0x2a, 0x26, 0xce, 0x40, 0x98, 0xf0, 0x29, 0x24, 0xf6, 0x23, 0x26, 0x51, 0x38, 0x3b, 0xc6, 0x3d,
	0xc0, 0xc9, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0xf1, 0x15, 0x92, 0x9a, 0x24, 0x36, 0x70,
	0x38, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0xcb, 0x7b, 0xf2, 0x7c, 0x01, 0x00, 0x00,
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
				m.PropertyList = &base.PropertyList{}
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
