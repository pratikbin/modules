// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/orders/internal/transactions/deputize/message.proto

package deputize

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
	From                 github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required field From missing"`
	FromID               github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d" valid:"required~required field FromID missing"`
	ToID                 github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,3,opt,name=to_i_d,json=toID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"to_i_d" valid:"required~required field ToID missing"`
	ClassificationID     github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,4,opt,name=classification_i_d,json=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classification_i_d" valid:"required~required field ClassificationID missing"`
	MaintainedProperties github_com_persistenceOne_persistenceSDK_schema_types.Properties      `protobuf:"bytes,5,opt,name=maintained_properties,json=maintainedProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"maintained_properties" valid:"required~required field MaintainedProperties missing"`
	AddMaintainer        bool                                                                  `protobuf:"varint,6,opt,name=add_maintainer,json=addMaintainer,proto3" json:"add_maintainer,omitempty"`
	RemoveMaintainer     bool                                                                  `protobuf:"varint,7,opt,name=remove_maintainer,json=removeMaintainer,proto3" json:"remove_maintainer,omitempty"`
	MutateMaintainer     bool                                                                  `protobuf:"varint,8,opt,name=mutate_maintainer,json=mutateMaintainer,proto3" json:"mutate_maintainer,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_65df680d26cf9839, []int{0}
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
	proto.RegisterType((*Message)(nil), "modules.orders.internal.transactions.deputize.Message")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/orders/internal/transactions/deputize/message.proto", fileDescriptor_65df680d26cf9839)
}

var fileDescriptor_65df680d26cf9839 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd4, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0x07, 0xf0, 0x8c, 0xae, 0xbb, 0x6b, 0x40, 0xa9, 0xa1, 0xc2, 0xe2, 0x21, 0x09, 0x0b, 0x42,
	0x41, 0xcd, 0x1c, 0xbc, 0x94, 0xe2, 0xc1, 0xd6, 0x55, 0x58, 0x4a, 0x51, 0xaa, 0x78, 0xf0, 0x12,
	0x66, 0x33, 0x6f, 0xd3, 0xc1, 0x4c, 0x26, 0xce, 0x4c, 0x0a, 0x8a, 0x7a, 0x94, 0x1e, 0x3c, 0xe8,
	0x7f, 0xd0, 0xb3, 0x07, 0xff, 0x04, 0xcf, 0x3d, 0xf6, 0x28, 0x1e, 0x82, 0xec, 0x5e, 0x3c, 0xef,
	0x5f, 0x20, 0x3b, 0xd9, 0x1f, 0x59, 0x11, 0xb6, 0x2d, 0x7b, 0x7b, 0xc9, 0x7b, 0xdf, 0x97, 0x0f,
	0x13, 0x18, 0x7b, 0x37, 0x03, 0xa9, 0x98, 0xd2, 0x90, 0x46, 0x10, 0x2a, 0xfa, 0x1a, 0x73, 0x41,
	0xf3, 0x04, 0x14, 0x16, 0x92, 0x82, 0x54, 0x98, 0xa5, 0x1a, 0x64, 0x4a, 0x12, 0xac, 0x25, 0x49,
	0x15, 0x89, 0x34, 0x13, 0xa9, 0xc2, 0x14, 0xb2, 0x5c, 0xb3, 0x77, 0x80, 0x39, 0x28, 0x45, 0x62,
	0x08, 0x32, 0x29, 0xb4, 0x70, 0xee, 0x4d, 0xc2, 0x41, 0x19, 0x0e, 0xa6, 0xe1, 0xa0, 0x1a, 0x0e,
	0xa6, 0xe1, 0x5b, 0xeb, 0xb1, 0x88, 0x85, 0x49, 0xe2, 0x71, 0x55, 0x2e, 0x69, 0x7f, 0x6f, 0xd8,
	0x8d, 0xbd, 0x72, 0xad, 0xf3, 0x15, 0xd9, 0xb5, 0xbe, 0x14, 0xbc, 0x85, 0x7c, 0xb4, 0x71, 0x75,
	0xe7, 0xc3, 0x49, 0xe1, 0x59, 0xbf, 0x0a, 0xef, 0x71, 0xcc, 0xf4, 0x41, 0xde, 0x0b, 0x22, 0xc1,
	0x71, 0xc5, 0xff, 0x34, 0x85, 0xea, 0xe3, 0xf3, 0xce, 0x2e, 0x56, 0xd1, 0x01, 0x70, 0x82, 0xf5,
	0xdb, 0x0c, 0x14, 0xee, 0x11, 0x05, 0xc1, 0x76, 0x14, 0x6d, 0x53, 0x2a, 0x41, 0xa9, 0x51, 0xe1,
	0xdd, 0x3d, 0x24, 0x09, 0xa3, 0x5b, 0x6d, 0x09, 0x6f, 0x72, 0x26, 0x81, 0x7e, 0x9c, 0x16, 0x7e,
	0x9f, 0x41, 0x42, 0xfd, 0x27, 0x52, 0x70, 0x9f, 0x33, 0xa5, 0x58, 0x1a, 0xb7, 0xf7, 0x0d, 0xc5,
	0xf9, 0x8c, 0xec, 0xe6, 0xb8, 0x08, 0x59, 0x48, 0x5b, 0x97, 0x8c, 0x4b, 0x4e, 0x5c, 0x9b, 0x17,
	0x72, 0x05, 0xdd, 0xce, 0xa8, 0xf0, 0x82, 0x33, 0x50, 0xba, 0x9d, 0x39, 0xa6, 0xde, 0x37, 0x2f,
	0x9c, 0x4f, 0xc8, 0xae, 0x6b, 0x61, 0x30, 0x97, 0x0d, 0x26, 0x5b, 0x01, 0x66, 0xd9, 0xb9, 0xbc,
	0x10, 0x55, 0x4a, 0x4d, 0x8b, 0x6e, 0xc7, 0xf9, 0x86, 0x6c, 0x27, 0x4a, 0x88, 0x52, 0xac, 0xcf,
	0x22, 0x32, 0xfe, 0xd5, 0x06, 0x55, 0x33, 0xa8, 0xf7, 0x2b, 0x40, 0x6d, 0x2e, 0x41, 0x3d, 0x5a,
	0xf8, 0x78, 0x15, 0xb8, 0x16, 0xfd, 0xd3, 0x72, 0x7e, 0x20, 0xfb, 0x26, 0x27, 0x2c, 0xd5, 0x84,
	0xa5, 0x40, 0xc3, 0x4c, 0x8a, 0x0c, 0xa4, 0x66, 0xa0, 0x5a, 0x57, 0x8c, 0xf7, 0x08, 0x4d, 0xc0,
	0x0f, 0x2f, 0x06, 0x7e, 0x36, 0x5b, 0x38, 0x2a, 0xbc, 0x07, 0x4b, 0xe0, 0x7b, 0x33, 0xc8, 0x3c,
	0x36, 0xc7, 0xaf, 0xf3, 0xff, 0xb4, 0x9d, 0xdb, 0xf6, 0x75, 0x42, 0x69, 0x38, 0xeb, 0xc9, 0x56,
	0xdd, 0x47, 0x1b, 0xcd, 0xfd, 0x6b, 0x84, 0xd2, 0xd9, 0x3e, 0xe9, 0xdc, 0xb1, 0x6f, 0x48, 0xe0,
	0xe2, 0x10, 0xaa, 0x93, 0x0d, 0x33, 0xb9, 0x56, 0x36, 0x16, 0x87, 0x79, 0xae, 0x89, 0x5e, 0x18,
	0x6e, 0x96, 0xc3, 0x65, 0x63, 0x3e, 0xbc, 0xd5, 0x3c, 0x3a, 0xf6, 0xac, 0x3f, 0xc7, 0x9e, 0xb5,
	0x93, 0x9d, 0x0c, 0x5c, 0x74, 0x3a, 0x70, 0xd1, 0xef, 0x81, 0x8b, 0xbe, 0x0c, 0x5d, 0xeb, 0x74,
	0xe8, 0x5a, 0x3f, 0x87, 0xae, 0xf5, 0xea, 0xe5, 0x99, 0x0f, 0xef, 0x5c, 0xb7, 0x4e, 0xaf, 0x6e,
	0x6e, 0x8a, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x7b, 0x40, 0xdf, 0xbd, 0x04, 0x00,
	0x00,
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
	if m.MutateMaintainer {
		i--
		if m.MutateMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.RemoveMaintainer {
		i--
		if m.RemoveMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.AddMaintainer {
		i--
		if m.AddMaintainer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.MaintainedProperties.Size()
		i -= size
		if _, err := m.MaintainedProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
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
	l = m.MaintainedProperties.Size()
	n += 1 + l + sovMessage(uint64(l))
	if m.AddMaintainer {
		n += 2
	}
	if m.RemoveMaintainer {
		n += 2
	}
	if m.MutateMaintainer {
		n += 2
	}
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainedProperties", wireType)
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
			if err := m.MaintainedProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddMaintainer = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoveMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RemoveMaintainer = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutateMaintainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MutateMaintainer = bool(v != 0)
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
