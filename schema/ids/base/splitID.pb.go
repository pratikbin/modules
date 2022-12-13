// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/ids/base/splitID.proto

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

type SplitID struct {
	OwnerId   *IdentityID `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnableId *OwnableID  `protobuf:"bytes,2,opt,name=ownable_id,json=ownableId,proto3" json:"ownable_id,omitempty"`
}

func (m *SplitID) Reset()         { *m = SplitID{} }
func (m *SplitID) String() string { return proto.CompactTextString(m) }
func (*SplitID) ProtoMessage()    {}
func (*SplitID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8594a36c5a7a51, []int{0}
}
func (m *SplitID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SplitID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SplitID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SplitID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitID.Merge(m, src)
}
func (m *SplitID) XXX_Size() int {
	return m.Size()
}
func (m *SplitID) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitID.DiscardUnknown(m)
}

var xxx_messageInfo_SplitID proto.InternalMessageInfo

func (m *SplitID) GetOwnerId() *IdentityID {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

func (m *SplitID) GetOwnableId() *OwnableID {
	if m != nil {
		return m.OwnableId
	}
	return nil
}

func init() {
	proto.RegisterType((*SplitID)(nil), "base.SplitID")
}

func init() { proto.RegisterFile("schema/ids/base/splitID.proto", fileDescriptor_fe8594a36c5a7a51) }

var fileDescriptor_fe8594a36c5a7a51 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0x4c, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e, 0xc8, 0xc9,
	0x2c, 0xf1, 0x74, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89, 0x49, 0x29, 0xa0,
	0x2b, 0xca, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x84, 0xa9, 0x93, 0x92, 0x47, 0x57, 0x91,
	0x5f, 0x9e, 0x97, 0x98, 0x94, 0x93, 0x0a, 0x53, 0xa0, 0x94, 0xc6, 0xc5, 0x1e, 0x0c, 0x31, 0x59,
	0x48, 0x9b, 0x8b, 0x23, 0xbf, 0x3c, 0x2f, 0xb5, 0x28, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51,
	0x83, 0xdb, 0x48, 0x40, 0x0f, 0xa4, 0x47, 0xcf, 0x13, 0x6e, 0x6a, 0x10, 0x3b, 0x58, 0x85, 0x67,
	0x8a, 0x90, 0x1e, 0x17, 0x17, 0xd4, 0x28, 0x90, 0x72, 0x26, 0xb0, 0x72, 0x7e, 0x88, 0x72, 0x7f,
	0x98, 0x15, 0x41, 0x9c, 0x30, 0xdb, 0x52, 0x9c, 0x26, 0x30, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x03, 0x17, 0x47, 0x72, 0x7e, 0x2e, 0x58, 0xa7, 0x13, 0x0f, 0xd4, 0x29, 0x01,
	0x20, 0xa7, 0x05, 0x30, 0x46, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x3b, 0x16, 0x17, 0xa7, 0x96, 0xf8, 0x26, 0xe6, 0x95, 0xe4, 0xa4, 0xea, 0xe7, 0xe6, 0xa7,
	0x94, 0xe6, 0xa4, 0x16, 0xeb, 0xa3, 0x79, 0x6e, 0x11, 0x13, 0xb3, 0x53, 0x44, 0xc4, 0x2a, 0x26,
	0x16, 0xa7, 0xc4, 0xe2, 0xd4, 0x53, 0x10, 0xea, 0x11, 0x93, 0x00, 0x88, 0x8a, 0x71, 0x0f, 0x70,
	0xf2, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0x05, 0x91, 0x49, 0x62, 0x03, 0x87, 0x80,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x34, 0xda, 0x2a, 0x6b, 0x01, 0x00, 0x00,
}

func (m *SplitID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SplitID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SplitID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OwnableId != nil {
		{
			size, err := m.OwnableId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplitID(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.OwnerId != nil {
		{
			size, err := m.OwnerId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSplitID(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSplitID(dAtA []byte, offset int, v uint64) int {
	offset -= sovSplitID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SplitID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnerId != nil {
		l = m.OwnerId.Size()
		n += 1 + l + sovSplitID(uint64(l))
	}
	if m.OwnableId != nil {
		l = m.OwnableId.Size()
		n += 1 + l + sovSplitID(uint64(l))
	}
	return n
}

func sovSplitID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSplitID(x uint64) (n int) {
	return sovSplitID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SplitID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSplitID
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
			return fmt.Errorf("proto: SplitID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SplitID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplitID
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
				return ErrInvalidLengthSplitID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplitID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerId == nil {
				m.OwnerId = &IdentityID{}
			}
			if err := m.OwnerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplitID
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
				return ErrInvalidLengthSplitID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplitID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnableId == nil {
				m.OwnableId = &OwnableID{}
			}
			if err := m.OwnableId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSplitID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSplitID
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
func skipSplitID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSplitID
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
					return 0, ErrIntOverflowSplitID
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
					return 0, ErrIntOverflowSplitID
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
				return 0, ErrInvalidLengthSplitID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSplitID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSplitID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSplitID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSplitID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSplitID = fmt.Errorf("proto: unexpected end of group")
)
