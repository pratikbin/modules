// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/ids/base/dataID.v1.proto

package base

import (
	fmt "fmt"
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

type DataID struct {
	Type   *StringID `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	HashID *HashID   `protobuf:"bytes,2,opt,name=hash_i_d,json=hashID,proto3" json:"hash_i_d,omitempty"`
}

func (m *DataID) Reset()         { *m = DataID{} }
func (m *DataID) String() string { return proto.CompactTextString(m) }
func (*DataID) ProtoMessage()    {}
func (*DataID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eec6336bce48983, []int{0}
}
func (m *DataID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataID.Merge(m, src)
}
func (m *DataID) XXX_Size() int {
	return m.Size()
}
func (m *DataID) XXX_DiscardUnknown() {
	xxx_messageInfo_DataID.DiscardUnknown(m)
}

var xxx_messageInfo_DataID proto.InternalMessageInfo

func (m *DataID) GetType() *StringID {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DataID) GetHashID() *HashID {
	if m != nil {
		return m.HashID
	}
	return nil
}

func init() {
	proto.RegisterType((*DataID)(nil), "ids.DataID")
}

func init() { proto.RegisterFile("schema/ids/base/dataID.v1.proto", fileDescriptor_3eec6336bce48983) }

var fileDescriptor_3eec6336bce48983 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0x4c, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0x49, 0x2c, 0x49,
	0xf4, 0x74, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29,
	0x96, 0xc2, 0x50, 0x95, 0x91, 0x58, 0x9c, 0x81, 0xa4, 0x4a, 0x4a, 0x11, 0x5d, 0x41, 0x71, 0x49,
	0x51, 0x66, 0x5e, 0x3a, 0x92, 0x12, 0xa5, 0x20, 0x2e, 0x36, 0x17, 0xb0, 0xd9, 0x42, 0x8a, 0x5c,
	0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xbc, 0x7a, 0x99, 0x29,
	0xc5, 0x7a, 0xc1, 0x50, 0xf5, 0x41, 0x60, 0x29, 0x21, 0x55, 0x2e, 0x0e, 0x90, 0x15, 0xf1, 0x99,
	0xf1, 0x29, 0x12, 0x4c, 0x60, 0x65, 0xdc, 0x60, 0x65, 0x1e, 0x60, 0x7b, 0x83, 0xd8, 0x20, 0xf6,
	0x3b, 0xf5, 0x30, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x03, 0x17, 0x7b, 0x72,
	0x7e, 0x2e, 0x48, 0x8b, 0x13, 0x2f, 0xc4, 0xd6, 0x30, 0xc3, 0x00, 0x90, 0x33, 0x02, 0x18, 0xa3,
	0xf4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x1d, 0x8b, 0x8b, 0x53,
	0x4b, 0x7c, 0x13, 0xf3, 0x4a, 0x72, 0x52, 0xf5, 0x73, 0xf3, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0xf5,
	0xd1, 0xbc, 0xb2, 0x88, 0x89, 0xd9, 0x33, 0x22, 0x62, 0x15, 0x13, 0xb3, 0x67, 0x4a, 0xf1, 0x29,
	0x30, 0xf9, 0x88, 0x89, 0xdf, 0x33, 0xa5, 0x38, 0xc6, 0x3d, 0xc0, 0xc9, 0x37, 0xb5, 0x24, 0x11,
	0x14, 0x60, 0xaf, 0xc0, 0xe2, 0x49, 0x6c, 0x60, 0x9f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0x31, 0xa7, 0x93, 0x55, 0x01, 0x00, 0x00,
}

func (m *DataID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintDataIDV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Type != nil {
		{
			size, err := m.Type.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataIDV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataIDV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataIDV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		l = m.Type.Size()
		n += 1 + l + sovDataIDV1(uint64(l))
	}
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovDataIDV1(uint64(l))
	}
	return n
}

func sovDataIDV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataIDV1(x uint64) (n int) {
	return sovDataIDV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataIDV1
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
			return fmt.Errorf("proto: DataID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataIDV1
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
				return ErrInvalidLengthDataIDV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataIDV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Type == nil {
				m.Type = &StringID{}
			}
			if err := m.Type.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataIDV1
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
				return ErrInvalidLengthDataIDV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataIDV1
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
			skippy, err := skipDataIDV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataIDV1
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
func skipDataIDV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataIDV1
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
					return 0, ErrIntOverflowDataIDV1
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
					return 0, ErrIntOverflowDataIDV1
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
				return 0, ErrInvalidLengthDataIDV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataIDV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataIDV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataIDV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataIDV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataIDV1 = fmt.Errorf("proto: unexpected end of group")
)
