// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/metas/internal/queries/meta/queryRequest.v1.proto

package meta

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
	DataId *base.DataID `protobuf:"bytes,1,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_525348278d0b15ba, []int{0}
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
	proto.RegisterType((*QueryRequest)(nil), "meta.QueryRequest")
}

func init() {
	proto.RegisterFile("modules/metas/internal/queries/meta/queryRequest.v1.proto", fileDescriptor_525348278d0b15ba)
}

var fileDescriptor_525348278d0b15ba = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcc, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca,
	0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x84, 0x0a, 0x83, 0x39, 0x95, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x7a, 0x65, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20,
	0x39, 0x29, 0xf9, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0xfd, 0xcc, 0x94, 0x62, 0xfd, 0xa4, 0xc4,
	0xe2, 0x54, 0xfd, 0x94, 0xc4, 0x92, 0x44, 0x4f, 0x17, 0xb8, 0x32, 0x29, 0x91, 0xf4, 0xfc, 0xf4,
	0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x59, 0x71, 0xf1, 0x04, 0x22, 0x99, 0x2a, 0xa4,
	0xc2, 0xc5, 0x0e, 0xd2, 0x18, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xad,
	0x97, 0x99, 0x52, 0xac, 0xe7, 0x02, 0x36, 0x2c, 0x88, 0x0d, 0x6c, 0x68, 0x8a, 0x15, 0x4b, 0xc7,
	0x02, 0x79, 0x06, 0xa7, 0xd5, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0,
	0xc5, 0x91, 0x9c, 0x9f, 0xab, 0x07, 0x72, 0x97, 0x93, 0x30, 0xb2, 0xf1, 0x61, 0x86, 0x01, 0x20,
	0x5b, 0x03, 0x18, 0xa3, 0x9c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x1d, 0x8b, 0x8b, 0x53, 0x4b, 0x7c, 0x13, 0xf3, 0x4a, 0x72, 0x52, 0xf5, 0xe1, 0xc1, 0x40, 0x38,
	0x38, 0x16, 0x31, 0x31, 0xfb, 0x46, 0x44, 0xac, 0x62, 0x62, 0xf1, 0x4d, 0x2d, 0x49, 0x3c, 0x05,
	0xa1, 0x1e, 0x31, 0x09, 0x80, 0xa8, 0x18, 0xf7, 0x00, 0x27, 0x10, 0x0d, 0x72, 0xee, 0x2b, 0x88,
	0x4c, 0x12, 0x1b, 0xd8, 0xc3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x95, 0xea, 0xec,
	0x6a, 0x01, 0x00, 0x00,
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
	if m.DataId != nil {
		{
			size, err := m.DataId.MarshalToSizedBuffer(dAtA[:i])
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
	if m.DataId != nil {
		l = m.DataId.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field DataId", wireType)
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
			if m.DataId == nil {
				m.DataId = &base.DataID{}
			}
			if err := m.DataId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
