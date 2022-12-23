// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/ids/base/maintainerID.v1.proto

package base

import (
	fmt "fmt"
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

type MaintainerID struct {
	HashID *HashID `protobuf:"bytes,1,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d,omitempty"`
}

func (m *MaintainerID) Reset()         { *m = MaintainerID{} }
func (m *MaintainerID) String() string { return proto.CompactTextString(m) }
func (*MaintainerID) ProtoMessage()    {}
func (*MaintainerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bbbc14d866ec4c, []int{0}
}
func (m *MaintainerID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MaintainerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MaintainerID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MaintainerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintainerID.Merge(m, src)
}
func (m *MaintainerID) XXX_Size() int {
	return m.Size()
}
func (m *MaintainerID) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintainerID.DiscardUnknown(m)
}

var xxx_messageInfo_MaintainerID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MaintainerID)(nil), "ids.MaintainerID")
}

func init() {
	proto.RegisterFile("schema/ids/base/maintainerID.v1.proto", fileDescriptor_e3bbbc14d866ec4c)
}

var fileDescriptor_e3bbbc14d866ec4c = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0x4c, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4d, 0xcc, 0xcc,
	0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0xf2, 0x74, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x92, 0x47, 0x57, 0x9b, 0x91, 0x58, 0x9c, 0x81, 0xa4,
	0x4a, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0xd6, 0x5c,
	0x3c, 0xbe, 0x48, 0x86, 0x0a, 0xa9, 0x72, 0x71, 0x80, 0x34, 0xc6, 0x67, 0xc6, 0xa7, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xeb, 0x65, 0xa6, 0x14, 0xeb, 0x79, 0x80, 0x4d, 0x0b, 0x62,
	0x83, 0x98, 0x6a, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0x69, 0x12, 0xe3, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x82, 0x34, 0x3a, 0x09, 0x23, 0x1b,
	0x1f, 0x66, 0x18, 0x00, 0xb2, 0x35, 0x80, 0x31, 0x4a, 0x2f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xb1, 0xb8, 0x38, 0xb5, 0xc4, 0x37, 0x31, 0xaf, 0x24, 0x27, 0x55,
	0x3f, 0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x58, 0x1f, 0xcd, 0x37, 0x8b, 0x98, 0x98, 0x3d, 0x23,
	0x22, 0x56, 0x31, 0x31, 0x7b, 0xa6, 0x14, 0x9f, 0x02, 0x93, 0x8f, 0x98, 0xf8, 0x3d, 0x53, 0x8a,
	0x63, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x5f, 0x81, 0xc5, 0x93,
	0xd8, 0xc0, 0x1e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x06, 0xf0, 0x3f, 0x3d, 0x01,
	0x00, 0x00,
}

func (m *MaintainerID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MaintainerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MaintainerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HashID != nil {
		{
			size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMaintainerIDV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMaintainerIDV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMaintainerIDV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MaintainerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovMaintainerIDV1(uint64(l))
	}
	return n
}

func sovMaintainerIDV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMaintainerIDV1(x uint64) (n int) {
	return sovMaintainerIDV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MaintainerID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintainerIDV1
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
			return fmt.Errorf("proto: MaintainerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MaintainerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintainerIDV1
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
				return ErrInvalidLengthMaintainerIDV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMaintainerIDV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HashID == nil {
				m.HashID = &HashID{}
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintainerIDV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMaintainerIDV1
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
func skipMaintainerIDV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMaintainerIDV1
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
					return 0, ErrIntOverflowMaintainerIDV1
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
					return 0, ErrIntOverflowMaintainerIDV1
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
				return 0, ErrInvalidLengthMaintainerIDV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMaintainerIDV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMaintainerIDV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMaintainerIDV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMaintainerIDV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMaintainerIDV1 = fmt.Errorf("proto: unexpected end of group")
)
