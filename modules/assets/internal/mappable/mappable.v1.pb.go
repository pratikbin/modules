// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/mappable/mappable.v1.proto

package mappable

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/documents/base"
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

type Mappable struct {
	Asset *base.Document `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (m *Mappable) Reset()         { *m = Mappable{} }
func (m *Mappable) String() string { return proto.CompactTextString(m) }
func (*Mappable) ProtoMessage()    {}
func (*Mappable) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df0aa53619403bb, []int{0}
}
func (m *Mappable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mappable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mappable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mappable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mappable.Merge(m, src)
}
func (m *Mappable) XXX_Size() int {
	return m.Size()
}
func (m *Mappable) XXX_DiscardUnknown() {
	xxx_messageInfo_Mappable.DiscardUnknown(m)
}

var xxx_messageInfo_Mappable proto.InternalMessageInfo

func (m *Mappable) GetAsset() *base.Document {
	if m != nil {
		return m.Asset
	}
	return nil
}

func init() {
	proto.RegisterType((*Mappable)(nil), "assets.Mappable")
}

func init() {
	proto.RegisterFile("modules/assets/internal/mappable/mappable.v1.proto", fileDescriptor_3df0aa53619403bb)
}

var fileDescriptor_3df0aa53619403bb = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xca, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4d, 0x2c, 0x28, 0x48, 0x4c, 0xca, 0x49, 0x85, 0x33, 0xf4, 0xca,
	0x0c, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x20, 0x6a, 0xa5, 0xd4, 0x8b, 0x93, 0x33,
	0x52, 0x73, 0x13, 0xf5, 0x53, 0xf2, 0x93, 0x4b, 0x73, 0x53, 0xf3, 0x4a, 0x8a, 0xf5, 0x93, 0x12,
	0x8b, 0x53, 0xe1, 0x5c, 0xb8, 0x06, 0x25, 0x53, 0x2e, 0x0e, 0x5f, 0xa8, 0x29, 0x42, 0x9a, 0x5c,
	0xac, 0x60, 0xed, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x70, 0xdd, 0x7a, 0x2e,
	0x50, 0x56, 0x10, 0x44, 0x85, 0xd3, 0x3a, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0xe0, 0xe2, 0x4a, 0xce, 0xcf, 0xd5, 0x83, 0x38, 0xc3, 0x89, 0x1f, 0x66, 0x76, 0x98, 0x61,
	0x00, 0xc8, 0xba, 0x00, 0xc6, 0x28, 0xfb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0x7d, 0x47, 0x90, 0x2a, 0xdf, 0xc4, 0xbc, 0x12, 0x90, 0x5f, 0xa0, 0x9e, 0x25, 0xe4, 0xe9,
	0x45, 0x4c, 0xcc, 0x8e, 0x11, 0x11, 0xab, 0x98, 0xd8, 0xc0, 0x9a, 0x8b, 0x4f, 0xc1, 0x18, 0x8f,
	0x98, 0x84, 0x20, 0x8c, 0x18, 0xf7, 0x00, 0x27, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4,
	0x57, 0x30, 0xd9, 0x24, 0x36, 0xb0, 0x77, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x1f,
	0x03, 0xa6, 0x55, 0x01, 0x00, 0x00,
}

func (m *Mappable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mappable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mappable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Asset != nil {
		{
			size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMappableV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMappableV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMappableV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mappable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Asset != nil {
		l = m.Asset.Size()
		n += 1 + l + sovMappableV1(uint64(l))
	}
	return n
}

func sovMappableV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMappableV1(x uint64) (n int) {
	return sovMappableV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mappable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMappableV1
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
			return fmt.Errorf("proto: Mappable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mappable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMappableV1
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
				return ErrInvalidLengthMappableV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMappableV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Asset == nil {
				m.Asset = &base.Document{}
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMappableV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMappableV1
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
func skipMappableV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMappableV1
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
					return 0, ErrIntOverflowMappableV1
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
					return 0, ErrIntOverflowMappableV1
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
				return 0, ErrInvalidLengthMappableV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMappableV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMappableV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMappableV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMappableV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMappableV1 = fmt.Errorf("proto: unexpected end of group")
)
