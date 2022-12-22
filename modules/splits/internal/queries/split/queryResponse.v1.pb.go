// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/queries/asset/queryResponse.v1.proto

package split

import (
	fmt "fmt"
	github_com_AssetMantle_modules_schema_helpers "github.com/AssetMantle/modules/schema/helpers"
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
	Success bool                                                     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string                                                   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	List    []github_com_AssetMantle_modules_schema_helpers.Mappable `protobuf:"bytes,3,rep,name=list,proto3,customtype=github.com/AssetMantle/modules/schema/helpers.Mappable" json:"list,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8e64e17e5c4c3c, []int{0}
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
	proto.RegisterType((*QueryResponse)(nil), "asset.QueryResponse")
}

func init() {
	proto.RegisterFile("modules/assets/internal/queries/asset/queryResponse.v1.proto", fileDescriptor_7a8e64e17e5c4c3c)
}

var fileDescriptor_7a8e64e17e5c4c3c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc9, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x84, 0x89, 0x83, 0x79, 0x95, 0x41, 0xa9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x7a, 0x65, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0xac, 0x60, 0x59, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x88, 0x3e, 0x88, 0x05, 0x91, 0x54,
	0x9a, 0xca, 0xc8, 0xc5, 0x1b, 0x88, 0xac, 0x4f, 0x48, 0x82, 0x8b, 0xbd, 0xb8, 0x34, 0x39, 0x39,
	0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x23, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0x4d,
	0x2d, 0x2a, 0xca, 0x2f, 0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xfc, 0xb8,
	0x58, 0x72, 0x32, 0x8b, 0x4b, 0x24, 0x98, 0x15, 0x98, 0x35, 0x38, 0x9d, 0xac, 0x6e, 0xdd, 0x93,
	0x37, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x04, 0xd9, 0xed,
	0x9b, 0x98, 0x57, 0x92, 0x93, 0xaa, 0x0f, 0xf3, 0x45, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0xa2, 0x7e,
	0x46, 0x6a, 0x4e, 0x41, 0x6a, 0x51, 0xb1, 0x9e, 0x6f, 0x62, 0x41, 0x41, 0x62, 0x52, 0x4e, 0x6a,
	0x10, 0xd8, 0x1c, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4e, 0x9b, 0x19, 0x4f, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x8b, 0x33, 0x39, 0x3f, 0x57, 0x0f, 0xec, 0x27, 0x27, 0x11,
	0x14, 0xa7, 0x87, 0x19, 0x06, 0x80, 0xfc, 0x14, 0xc0, 0x18, 0xe5, 0x42, 0xc0, 0x05, 0x44, 0x85,
	0xe7, 0x22, 0x26, 0x66, 0xc7, 0x88, 0x88, 0x55, 0x4c, 0xac, 0x60, 0x13, 0x4e, 0x41, 0xe9, 0x47,
	0x4c, 0x82, 0x60, 0x3a, 0xc6, 0x3d, 0xc0, 0xc9, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0xf1,
	0x15, 0x54, 0x2e, 0x89, 0x0d, 0x1c, 0xa8, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x22,
	0x1a, 0xff, 0xb1, 0x01, 0x00, 0x00,
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
				size := m.List[iNdEx].Size()
				i -= size
				if _, err := m.List[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
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
			var v github_com_AssetMantle_modules_schema_helpers.Mappable
			m.List = append(m.List, v)
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
