// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/maintainers/internal/queries/maintainer/queryRequest.v1.proto

package maintainer

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
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

type QueryRequest struct {
	MaintainerId *base.MaintainerID `protobuf:"bytes,1,opt,name=maintainer_id,json=maintainerId,proto3" json:"maintainer_id,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b0ae64c6d0435a, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryRequest)(nil), "maintainer.QueryRequest")
}

func init() {
	proto.RegisterFile("modules/maintainers/internal/queries/maintainer/queryRequest.v1.proto", fileDescriptor_73b0ae64c6d0435a)
}

var fileDescriptor_73b0ae64c6d0435a = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcd, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0xcc, 0xcc, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x2a, 0xd6,
	0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x44, 0x91,
	0x04, 0x0b, 0x55, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0xe8, 0x95, 0x19, 0xea, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x54, 0x48, 0xa9, 0x16, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea,
	0x67, 0xa6, 0x14, 0xeb, 0x27, 0x25, 0x16, 0xa7, 0x22, 0xe9, 0xf6, 0x74, 0x81, 0x6b, 0x91, 0x12,
	0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x92, 0x0f, 0x17, 0x4f, 0x20,
	0x92, 0x0d, 0x42, 0x66, 0x5c, 0xbc, 0x08, 0xed, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x82, 0x7a, 0x99, 0x29, 0xc5, 0x7a, 0xbe, 0x48, 0x06, 0x07, 0xf1, 0x20, 0x59, 0x93,
	0x62, 0xc5, 0xd2, 0xb1, 0x40, 0x9e, 0xc1, 0xe9, 0x2a, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x70, 0xf1, 0x25, 0xe7, 0xe7, 0xea, 0x21, 0xb4, 0x38, 0x09, 0x23, 0x5b, 0x1b,
	0x66, 0x18, 0x00, 0x72, 0x4d, 0x00, 0x63, 0x94, 0x5f, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0xbe, 0x63, 0x71, 0x71, 0x6a, 0x89, 0x6f, 0x62, 0x5e, 0x49, 0x4e, 0xaa, 0x3e,
	0x3c, 0xd8, 0x48, 0x0b, 0xbe, 0x45, 0x4c, 0xcc, 0xbe, 0x11, 0x11, 0xab, 0x98, 0xb8, 0x10, 0x4e,
	0x3f, 0x85, 0xcc, 0x79, 0xc4, 0x24, 0x86, 0xe0, 0xc4, 0xb8, 0x07, 0x38, 0xf9, 0xa6, 0x96, 0x24,
	0xa6, 0x24, 0x96, 0x24, 0xbe, 0x42, 0x56, 0x95, 0xc4, 0x06, 0x0e, 0x2c, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0x1e, 0x12, 0x87, 0xbe, 0x01, 0x00, 0x00,
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaintainerId != nil {
		{
			size, err := m.MaintainerId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRequestV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryRequestV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryRequestV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaintainerId != nil {
		l = m.MaintainerId.Size()
		n += 1 + l + sovQueryRequestV1(uint64(l))
	}
	return n
}

func sovQueryRequestV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRequestV1(x uint64) (n int) {
	return sovQueryRequestV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRequestV1
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
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRequestV1
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
				return ErrInvalidLengthQueryRequestV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRequestV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaintainerId == nil {
				m.MaintainerId = &base.MaintainerID{}
			}
			if err := m.MaintainerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRequestV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRequestV1
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
func skipQueryRequestV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryRequestV1
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
					return 0, ErrIntOverflowQueryRequestV1
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
					return 0, ErrIntOverflowQueryRequestV1
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
				return 0, ErrInvalidLengthQueryRequestV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryRequestV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryRequestV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryRequestV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryRequestV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryRequestV1 = fmt.Errorf("proto: unexpected end of group")
)
