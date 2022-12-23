// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/queries/order/queryResponse.v1.proto

package order

import (
	fmt "fmt"
	mappable "github.com/AssetMantle/modules/modules/orders/internal/mappable"
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

type QueryResponse struct {
	Success bool                 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string               `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	List    []*mappable.Mappable `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ab2cdd0000125b, []int{0}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryResponse)(nil), "orders.queries.order.QueryResponse")
}

func init() {
	proto.RegisterFile("modules/orders/internal/queries/order/queryResponse.v1.proto", fileDescriptor_e2ab2cdd0000125b)
}

var fileDescriptor_e2ab2cdd0000125b = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc9, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x84, 0x89, 0x83, 0x79, 0x95, 0x41, 0xa9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x7a, 0x65, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x22, 0x10, 0x5d, 0x7a, 0x50, 0xc5, 0x7a, 0x60, 0xae, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x81, 0x3e, 0x88, 0x05, 0x51, 0x2b, 0x65, 0x84, 0xcb, 0xa6, 0xdc, 0xc4, 0x82, 0x82, 0xc4, 0xa4,
	0x9c, 0x54, 0x38, 0x03, 0x6e, 0xbe, 0x52, 0x36, 0x17, 0x6f, 0x20, 0xb2, 0xcd, 0x42, 0x12, 0x5c,
	0xec, 0xc5, 0xa5, 0xc9, 0xc9, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x30,
	0xae, 0x90, 0x08, 0x17, 0x6b, 0x6a, 0x51, 0x51, 0x7e, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x84, 0x23, 0xa4, 0xc2, 0xc5, 0x92, 0x93, 0x59, 0x5c, 0x22, 0xc1, 0xac, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xa0, 0x07, 0x75, 0xaf, 0x2f, 0xd4, 0xa6, 0x20, 0xb0, 0xac, 0x15, 0x4b, 0xc7, 0x02,
	0x79, 0x06, 0xa7, 0x06, 0xa6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0x92,
	0x48, 0xce, 0xcf, 0xd5, 0xc3, 0xe6, 0x57, 0x27, 0x11, 0x14, 0xf7, 0x85, 0x19, 0x06, 0x80, 0xdc,
	0x1d, 0xc0, 0x18, 0xe5, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef,
	0x58, 0x5c, 0x9c, 0x5a, 0xe2, 0x9b, 0x98, 0x57, 0x02, 0xf2, 0x23, 0x34, 0x10, 0x88, 0x0a, 0xf6,
	0x45, 0x4c, 0xcc, 0xfe, 0x81, 0xfe, 0xab, 0x98, 0x44, 0xfc, 0x21, 0x56, 0x07, 0x42, 0xad, 0x06,
	0x73, 0x4f, 0xc1, 0x84, 0x63, 0xa0, 0xc2, 0x31, 0x60, 0xee, 0x23, 0x26, 0x05, 0x6c, 0xc2, 0x31,
	0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89, 0x29, 0x89, 0x25, 0x89, 0xaf, 0x98, 0xc4, 0x20, 0x4a,
	0xac, 0xac, 0xa0, 0x6a, 0xac, 0xac, 0xc0, 0x02, 0x49, 0x6c, 0xe0, 0x60, 0x37, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0xf6, 0xaf, 0x83, 0x16, 0x02, 0x00, 0x00,
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryResponseV1(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintQueryResponseV1(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryResponseV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryResponseV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovQueryResponseV1(uint64(l))
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovQueryResponseV1(uint64(l))
		}
	}
	return n
}

func sovQueryResponseV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryResponseV1(x uint64) (n int) {
	return sovQueryResponseV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryResponseV1
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
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponseV1
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
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponseV1
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
				return ErrInvalidLengthQueryResponseV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponseV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResponseV1
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
				return ErrInvalidLengthQueryResponseV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryResponseV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &mappable.Mappable{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryResponseV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryResponseV1
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
func skipQueryResponseV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryResponseV1
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
					return 0, ErrIntOverflowQueryResponseV1
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
					return 0, ErrIntOverflowQueryResponseV1
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
				return 0, ErrInvalidLengthQueryResponseV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryResponseV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryResponseV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryResponseV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryResponseV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryResponseV1 = fmt.Errorf("proto: unexpected end of group")
)
