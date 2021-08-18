// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/orders/internal/transactions/take/message.proto

package take

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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
	From              github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required field From missing"`
	FromID            github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d" valid:"required~required field FromID missing"`
	TakerOwnableSplit github_com_cosmos_cosmos_sdk_types.Dec                                `protobuf:"bytes,3,opt,name=taker_ownable_split,json=takerOwnableSplit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"taker_ownable_split" valid:"required~required field TakerOwnableSplit missing"`
	OrderID           github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,4,opt,name=order_i_d,json=orderID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"order_i_d" valid:"required~required field OrderID missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e4f159f493b6506, []int{0}
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
	proto.RegisterType((*Message)(nil), "persistence_sdk.modules.orders.internal.transactions.take.Message")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/orders/internal/transactions/take/message.proto", fileDescriptor_4e4f159f493b6506)
}

var fileDescriptor_4e4f159f493b6506 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x10, 0xf5, 0x87, 0x37, 0x0c, 0x43, 0xc4, 0x60, 0x47, 0x19, 0x10, 0x03, 0xdc,
	0x0d, 0x2c, 0xb4, 0x5b, 0xab, 0x00, 0xaa, 0x10, 0x8a, 0xd4, 0x32, 0xb1, 0x58, 0x17, 0xdf, 0x8b,
	0x7b, 0x8a, 0xef, 0x2e, 0xdc, 0xbb, 0x80, 0x18, 0x60, 0x66, 0x40, 0x02, 0x66, 0x96, 0xfe, 0x39,
	0x1d, 0x3b, 0x22, 0x06, 0x0b, 0x25, 0x12, 0x62, 0xce, 0x5f, 0x80, 0xee, 0x9c, 0xaa, 0x56, 0x97,
	0x44, 0xa8, 0x93, 0x9f, 0xcf, 0xd2, 0xe7, 0xfb, 0xd1, 0xd7, 0xef, 0xe2, 0x17, 0x53, 0xb0, 0x28,
	0xd1, 0x81, 0x2e, 0x20, 0x47, 0x31, 0x61, 0xca, 0x88, 0x59, 0x05, 0xc8, 0x8c, 0x15, 0x60, 0x91,
	0x49, 0xed, 0xc0, 0x6a, 0x5e, 0x31, 0x67, 0xb9, 0x46, 0x5e, 0x38, 0x69, 0x34, 0x32, 0xc7, 0x27,
	0xc0, 0x14, 0x20, 0xf2, 0x12, 0xe8, 0xd4, 0x1a, 0x67, 0x92, 0xbd, 0x6b, 0x20, 0xba, 0x02, 0xd1,
	0x06, 0x44, 0x2f, 0x41, 0xb4, 0x0d, 0xa2, 0x1e, 0x74, 0xff, 0x5e, 0x69, 0x4a, 0x13, 0x28, 0xcc,
	0x4f, 0x0d, 0xb0, 0xff, 0xa7, 0x13, 0x6f, 0xbf, 0x6a, 0x22, 0x92, 0xef, 0x24, 0xee, 0x8c, 0xad,
	0x51, 0x5d, 0xd2, 0x23, 0x0f, 0x77, 0x0f, 0x3f, 0x9e, 0xd7, 0x59, 0xf4, 0xab, 0xce, 0x9e, 0x95,
	0xd2, 0x9d, 0xce, 0x46, 0xb4, 0x30, 0x8a, 0xb5, 0xe2, 0x87, 0x1a, 0xda, 0xaf, 0x27, 0x83, 0x97,
	0x0c, 0x8b, 0x53, 0x50, 0x9c, 0xb9, 0x0f, 0x53, 0x40, 0x36, 0xe2, 0x08, 0xf4, 0xa0, 0x28, 0x0e,
	0x84, 0xb0, 0x80, 0xb8, 0xac, 0xb3, 0x47, 0xef, 0x78, 0x25, 0xc5, 0x7e, 0xdf, 0xc2, 0xdb, 0x99,
	0xb4, 0x20, 0x3e, 0x5d, 0x0e, 0xbd, 0xb1, 0x84, 0x4a, 0xf4, 0x9e, 0x5b, 0xa3, 0x7a, 0x4a, 0x22,
	0x4a, 0x5d, 0xf6, 0x8f, 0x83, 0x4a, 0xf2, 0x85, 0xc4, 0x3b, 0x7e, 0xc8, 0x65, 0x2e, 0xba, 0xb7,
	0x82, 0x97, 0x5d, 0x79, 0x3d, 0xfd, 0x2f, 0x2f, 0x7a, 0x34, 0x58, 0xd6, 0x19, 0xdd, 0x40, 0xe5,
	0x68, 0x70, 0x25, 0xb3, 0x35, 0x0e, 0x07, 0xc9, 0x0f, 0x12, 0xdf, 0xf5, 0x6d, 0xda, 0xdc, 0xbc,
	0xd7, 0x7c, 0x54, 0x41, 0x8e, 0xd3, 0x4a, 0xba, 0xee, 0xed, 0x60, 0x36, 0x59, 0x99, 0x3d, 0x68,
	0x99, 0x15, 0x06, 0x95, 0xc1, 0xd5, 0xe3, 0xb1, 0xff, 0xf7, 0x8d, 0xc6, 0x00, 0x8a, 0x65, 0x9d,
	0xed, 0xad, 0xf1, 0x78, 0xed, 0xb3, 0x86, 0x4d, 0xd4, 0x89, 0x4f, 0xba, 0x52, 0xba, 0xe3, 0xae,
	0x7f, 0x4b, 0xbe, 0x92, 0x78, 0x37, 0x2c, 0x42, 0x68, 0xab, 0x13, 0x9c, 0xf0, 0x06, 0xda, 0x62,
	0x6b, 0x2c, 0x87, 0x3e, 0xb3, 0x5d, 0xd7, 0xb6, 0x69, 0x4e, 0xf6, 0x77, 0x3e, 0x9f, 0x65, 0xd1,
	0xdf, 0xb3, 0x2c, 0x3a, 0xac, 0xce, 0xe7, 0x29, 0xb9, 0x98, 0xa7, 0xe4, 0xf7, 0x3c, 0x25, 0xdf,
	0x16, 0x69, 0x74, 0xb1, 0x48, 0xa3, 0x9f, 0x8b, 0x34, 0x7a, 0x73, 0xbc, 0xb1, 0xd9, 0xc6, 0xb7,
	0x66, 0xb4, 0x15, 0xb6, 0xfb, 0xc9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x88, 0x75, 0xea,
	0x79, 0x03, 0x00, 0x00,
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
		size := m.OrderID.Size()
		i -= size
		if _, err := m.OrderID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TakerOwnableSplit.Size()
		i -= size
		if _, err := m.TakerOwnableSplit.MarshalTo(dAtA[i:]); err != nil {
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
	l = m.TakerOwnableSplit.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.OrderID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field TakerOwnableSplit", wireType)
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
			if err := m.TakerOwnableSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
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
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
