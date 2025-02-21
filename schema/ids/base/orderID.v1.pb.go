// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/ids/base/orderID.v1.proto

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

type OrderID struct {
	OrderID *HashID `protobuf:"bytes,1,opt,name=order_i_d,json=orderID,proto3" json:"order_i_d,omitempty"`
}

func (m *OrderID) Reset()         { *m = OrderID{} }
func (m *OrderID) String() string { return proto.CompactTextString(m) }
func (*OrderID) ProtoMessage()    {}
func (*OrderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8b52739d625e944, []int{0}
}
func (m *OrderID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderID.Merge(m, src)
}
func (m *OrderID) XXX_Size() int {
	return m.Size()
}
func (m *OrderID) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderID.DiscardUnknown(m)
}

var xxx_messageInfo_OrderID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OrderID)(nil), "ids.OrderID")
}

func init() { proto.RegisterFile("schema/ids/base/orderID.v1.proto", fileDescriptor_e8b52739d625e944) }

var fileDescriptor_e8b52739d625e944 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0x4c, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x2f, 0x4a, 0x49,
	0x2d, 0xf2, 0x74, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c,
	0x29, 0x96, 0x92, 0x47, 0x57, 0x96, 0x91, 0x58, 0x9c, 0x81, 0xa4, 0x4a, 0x4a, 0x24, 0x3d, 0x3f,
	0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0x16, 0x5c, 0xec, 0xfe, 0x10, 0xf3, 0x84,
	0xd4, 0xb9, 0x38, 0xc1, 0x46, 0xc7, 0x67, 0xc6, 0xa7, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b,
	0x71, 0xeb, 0x65, 0xa6, 0x14, 0xeb, 0x79, 0x80, 0x4d, 0x0a, 0x62, 0x87, 0x5a, 0x6c, 0xc5, 0xd2,
	0xb1, 0x40, 0x9e, 0xc1, 0xa9, 0x97, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xb8, 0xd8, 0x93, 0xf3, 0x73, 0x41, 0x3a, 0x9d, 0xf8, 0xa0, 0x66, 0x87, 0x19, 0x06, 0x80, 0x6c,
	0x0b, 0x60, 0x8c, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77,
	0x2c, 0x2e, 0x4e, 0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0xcf, 0xcd, 0x4f, 0x29, 0xcd,
	0x49, 0x2d, 0xd6, 0x47, 0xf3, 0xc5, 0x22, 0x26, 0x66, 0xcf, 0x88, 0x88, 0x55, 0x4c, 0xcc, 0x9e,
	0x29, 0xc5, 0xa7, 0xc0, 0xe4, 0x23, 0x26, 0x7e, 0xcf, 0x94, 0xe2, 0x18, 0xf7, 0x00, 0x27, 0xdf,
	0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x57, 0x60, 0xf1, 0x24, 0x36, 0xb0, 0x87, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x00, 0x2d, 0x3a, 0x30, 0x01, 0x00, 0x00,
}

func (m *OrderID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderID != nil {
		{
			size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrderIDV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderIDV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderIDV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderID != nil {
		l = m.OrderID.Size()
		n += 1 + l + sovOrderIDV1(uint64(l))
	}
	return n
}

func sovOrderIDV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderIDV1(x uint64) (n int) {
	return sovOrderIDV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderIDV1
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
			return fmt.Errorf("proto: OrderID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderIDV1
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
				return ErrInvalidLengthOrderIDV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrderIDV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OrderID == nil {
				m.OrderID = &HashID{}
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrderIDV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderIDV1
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
func skipOrderIDV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderIDV1
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
					return 0, ErrIntOverflowOrderIDV1
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
					return 0, ErrIntOverflowOrderIDV1
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
				return 0, ErrInvalidLengthOrderIDV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderIDV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderIDV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderIDV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderIDV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderIDV1 = fmt.Errorf("proto: unexpected end of group")
)
