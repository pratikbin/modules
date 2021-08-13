// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/orders/internal/mappable/order.proto

package mappable

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/traits/base"
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

type Order struct {
	ID            github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"i_d"`
	HasImmutables base.HasImmutables                                       `protobuf:"bytes,2,opt,name=has_immutables,json=hasImmutables,proto3" json:"has_immutables"`
	HasMutables   base.HasMutables                                         `protobuf:"bytes,3,opt,name=has_mutables,json=hasMutables,proto3" json:"has_mutables"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_619aa3db73f03a59, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetHasImmutables() base.HasImmutables {
	if m != nil {
		return m.HasImmutables
	}
	return base.HasImmutables{}
}

func (m *Order) GetHasMutables() base.HasMutables {
	if m != nil {
		return m.HasMutables
	}
	return base.HasMutables{}
}

func init() {
	proto.RegisterType((*Order)(nil), "modules.orders.internal.mappable.Order")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/orders/internal/mappable/order.proto", fileDescriptor_619aa3db73f03a59)
}

var fileDescriptor_619aa3db73f03a59 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4d, 0x2c, 0x28, 0x48, 0x4c, 0xca, 0x49, 0x85, 0x48, 0xe8, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x29, 0x40, 0x55, 0xeb, 0x41, 0x54, 0xeb, 0xc1, 0x54, 0xeb, 0xc1,
	0x54, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x15, 0xeb, 0x83, 0x58, 0x10, 0x7d, 0x52, 0x66,
	0xe8, 0xb6, 0x16, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea, 0x97, 0x14, 0x25, 0x66, 0x96, 0x14, 0xeb,
	0x27, 0x25, 0x16, 0xa7, 0xea, 0x67, 0x24, 0x16, 0x7b, 0xe6, 0xe6, 0x96, 0x96, 0x80, 0x8c, 0x29,
	0x86, 0xea, 0x33, 0x21, 0x4e, 0x9f, 0x2f, 0x8a, 0x2e, 0xa5, 0x2f, 0x8c, 0x5c, 0xac, 0xfe, 0x20,
	0x07, 0x0a, 0x05, 0x72, 0x31, 0x67, 0xc6, 0xa7, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x39,
	0x9c, 0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x45, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0x92, 0xf9, 0xfe, 0x79, 0xa9, 0xc8, 0xdc, 0x60, 0x17, 0x6f, 0xb8, 0x6d,
	0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x9e, 0x2e, 0x41, 0x4c, 0x99, 0x2e, 0x42, 0x7e, 0x5c, 0x7c, 0x19,
	0x89, 0xc5, 0xf1, 0x99, 0x70, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0xea, 0x41,
	0x54, 0xeb, 0x41, 0xdc, 0xa6, 0x07, 0x72, 0x9b, 0x9e, 0x07, 0xb2, 0x9f, 0x9c, 0x58, 0x40, 0x0e,
	0x08, 0xe2, 0x45, 0xf1, 0xa8, 0x90, 0x07, 0x17, 0x0f, 0xc8, 0x3c, 0xb8, 0x69, 0xcc, 0x60, 0xd3,
	0xe4, 0x71, 0x98, 0xe6, 0x8b, 0x6a, 0x16, 0x37, 0x92, 0xe7, 0x9d, 0x92, 0x4f, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x93, 0x68, 0x1f, 0x13, 0x4a, 0x0d, 0x49, 0x6c, 0xe0,
	0x20, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x9b, 0x0a, 0x61, 0x48, 0x02, 0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HasMutables.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.HasImmutables.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.HasImmutables.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.HasMutables.Size()
	n += 1 + l + sovOrder(uint64(l))
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasImmutables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HasImmutables.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasMutables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HasMutables.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
