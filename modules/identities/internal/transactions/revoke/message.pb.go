// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/identities/internal/transactions/revoke/message.proto

package revoke

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	github_com_persistenceOne_persistenceSDK_schema_types_base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type Message struct {
	From             github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required From missing"`
	FromID           github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d" valid:"required~required FromID missing"`
	ToID             github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"to_i_d" valid:"required~required ToID missing"`
	ClassificationID github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classification_i_d" valid:"required~required ClassificationID missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c84d09bfd503118, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "modules.identities.internal.transactions.revoke.Message")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/identities/internal/transactions/revoke/message.proto", fileDescriptor_9c84d09bfd503118)
}

var fileDescriptor_9c84d09bfd503118 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd3, 0xb1, 0xab, 0xd3, 0x40,
	0x1c, 0x07, 0xf0, 0x9c, 0x86, 0xfa, 0xcc, 0x24, 0xc1, 0xe1, 0xe1, 0x90, 0x94, 0x0e, 0xf2, 0xa6,
	0x3b, 0xd0, 0x45, 0xde, 0xf6, 0x9e, 0x51, 0x28, 0xf2, 0x10, 0xd4, 0x41, 0x5c, 0xca, 0xf5, 0xee,
	0xd7, 0xf4, 0x6c, 0x72, 0x57, 0xef, 0x77, 0xad, 0xb8, 0x88, 0x93, 0x08, 0x2e, 0xfe, 0x09, 0x1d,
	0xfc, 0x63, 0x3a, 0x76, 0x70, 0x10, 0x87, 0x20, 0xed, 0xe2, 0xdc, 0xbf, 0x40, 0x72, 0xb5, 0x34,
	0x0a, 0x15, 0x7d, 0x74, 0xfb, 0x26, 0xe1, 0xfb, 0xcb, 0x27, 0xb9, 0xbb, 0xe8, 0x62, 0x0c, 0x16,
	0x15, 0x3a, 0xd0, 0x02, 0x7a, 0x28, 0x47, 0xac, 0x34, 0x72, 0x52, 0x00, 0x32, 0x25, 0x41, 0x3b,
	0xe5, 0x54, 0x1d, 0xb5, 0x03, 0xab, 0x79, 0xc1, 0x9c, 0xe5, 0x1a, 0xb9, 0x70, 0xca, 0x68, 0x64,
	0x16, 0xa6, 0x66, 0x04, 0xac, 0x04, 0x44, 0x9e, 0x03, 0x1d, 0x5b, 0xe3, 0x4c, 0xbc, 0xad, 0xd3,
	0x5d, 0x9d, 0x6e, 0xeb, 0xb4, 0x59, 0xa7, 0x9b, 0xfa, 0xad, 0x9b, 0xb9, 0xc9, 0x8d, 0xef, 0xb2,
	0x3a, 0x6d, 0xc6, 0x74, 0xbe, 0x84, 0xd1, 0xb5, 0x8b, 0xcd, 0xe0, 0xf8, 0x23, 0x89, 0xc2, 0x81,
	0x35, 0xe5, 0x31, 0x69, 0x93, 0x93, 0xeb, 0xe7, 0xaf, 0xe7, 0x55, 0x1a, 0x7c, 0xab, 0xd2, 0x07,
	0xb9, 0x72, 0xc3, 0x49, 0x9f, 0x0a, 0x53, 0xb2, 0xc6, 0x37, 0x3c, 0xd6, 0xd0, 0xbc, 0x7c, 0x9a,
	0x3d, 0x62, 0x28, 0x86, 0x50, 0x72, 0xe6, 0xde, 0x8c, 0x01, 0x59, 0x9f, 0x23, 0xd0, 0x33, 0x21,
	0xce, 0xa4, 0xb4, 0x80, 0xb8, 0xae, 0xd2, 0xdb, 0x53, 0x5e, 0x28, 0x79, 0xda, 0xb1, 0xf0, 0x6a,
	0xa2, 0x2c, 0xc8, 0xb7, 0xdb, 0xd0, 0x7e, 0x68, 0x4d, 0xd9, 0x2e, 0x15, 0xa2, 0xd2, 0x79, 0xe7,
	0x89, 0x47, 0xc4, 0xef, 0x49, 0x74, 0x54, 0x87, 0x9e, 0xea, 0xc9, 0xe3, 0x2b, 0x5e, 0x54, 0xfc,
	0x12, 0xdd, 0xbb, 0x94, 0x88, 0x76, 0xb3, 0x75, 0x95, 0x9e, 0xfc, 0x15, 0xd1, 0xcd, 0x76, 0x8c,
	0xd6, 0xc0, 0xdf, 0x88, 0xdf, 0x91, 0xa8, 0xe5, 0x8c, 0x67, 0x5c, 0xf5, 0x8c, 0x97, 0x07, 0x60,
	0xec, 0xff, 0x17, 0xcf, 0x4c, 0x13, 0x11, 0x3a, 0xd3, 0xcd, 0xe2, 0xcf, 0x24, 0x8a, 0x45, 0xc1,
	0x11, 0xd5, 0x40, 0x09, 0x5e, 0x2f, 0xab, 0xe7, 0x84, 0x9e, 0x33, 0x3d, 0x00, 0xe7, 0xce, 0x5e,
	0xce, 0xfd, 0xdf, 0x5e, 0xdb, 0xa4, 0xdd, 0x10, 0x7f, 0x3c, 0x3a, 0x3d, 0xfa, 0x30, 0x4b, 0x83,
	0x1f, 0xb3, 0x34, 0x38, 0xb7, 0xf3, 0x65, 0x42, 0x16, 0xcb, 0x84, 0x7c, 0x5f, 0x26, 0xe4, 0xd3,
	0x2a, 0x09, 0x16, 0xab, 0x24, 0xf8, 0xba, 0x4a, 0x82, 0x17, 0xcf, 0xff, 0x59, 0xf9, 0x9f, 0xe7,
	0xa3, 0xdf, 0xf2, 0x3b, 0xfa, 0xee, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xc8, 0x26, 0x8e,
	0x69, 0x03, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClassificationID.Size()
		i -= size
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ToID.Size()
		i -= size
		if _, err := m.ToID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.From.Size()
		i -= size
		if _, err := m.From.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.From.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.ToID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.ClassificationID.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
