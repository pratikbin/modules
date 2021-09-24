// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/splits/internal/key/splitID.proto

package key

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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
	OwnerID   github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=owner_i_d,json=ownerID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"owner_i_d"`
	OwnableID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,2,opt,name=ownable_i_d,json=ownableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"ownable_i_d"`
}

func (m *SplitID) Reset()      { *m = SplitID{} }
func (*SplitID) ProtoMessage() {}
func (*SplitID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6779bb553020d011, []int{0}
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

func init() {
	proto.RegisterType((*SplitID)(nil), "persistence_sdk.modules.splits.internal.key.SplitID")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/splits/internal/key/splitID.proto", fileDescriptor_6779bb553020d011)
}

var fileDescriptor_6779bb553020d011 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2c, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4e, 0xad, 0x84, 0x88, 0x79, 0xba, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x69, 0xa3, 0x69, 0xd5, 0x83, 0x6a, 0xd5, 0x83, 0x68, 0xd5, 0x83, 0x69, 0xd5, 0xcb,
	0x4e, 0xad, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xeb, 0xd3, 0x07, 0xb1, 0x20, 0x46, 0x28,
	0x9d, 0x67, 0xe4, 0x62, 0x0f, 0x86, 0x18, 0x2a, 0x14, 0xc3, 0xc5, 0x99, 0x5f, 0x9e, 0x97, 0x5a,
	0x14, 0x9f, 0x19, 0x9f, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xe4, 0x70, 0xe2, 0x9e, 0x3c,
	0xc3, 0xad, 0x7b, 0xf2, 0x16, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x48, 0x96, 0xfa, 0xe7, 0xa5, 0x22, 0x73, 0x83, 0x5d, 0xbc, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73,
	0x13, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x3c, 0x5d, 0x82, 0xd8, 0xc1, 0x46, 0x7a, 0xba,
	0x08, 0x25, 0x70, 0x71, 0xe7, 0x97, 0xe7, 0x25, 0x26, 0xe5, 0xa4, 0x82, 0xcd, 0x67, 0xa2, 0x92,
	0xf9, 0x9c, 0x50, 0x43, 0x3d, 0x5d, 0xac, 0x58, 0x66, 0x2c, 0x90, 0x67, 0x70, 0x8a, 0x3b, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x17, 0xa2, 0x2d, 0xc1, 0x13, 0x05, 0x49,
	0x6c, 0xe0, 0x80, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0xac, 0x6b, 0xbe, 0xb8, 0x01,
	0x00, 0x00,
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
	{
		size := m.OwnableID.Size()
		i -= size
		if _, err := m.OwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSplitID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.OwnerID.Size()
		i -= size
		if _, err := m.OwnerID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSplitID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.OwnerID.Size()
	n += 1 + l + sovSplitID(uint64(l))
	l = m.OwnableID.Size()
	n += 1 + l + sovSplitID(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplitID
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
				return ErrInvalidLengthSplitID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSplitID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplitID
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
				return ErrInvalidLengthSplitID
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSplitID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
