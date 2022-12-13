// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/classifications/internal/parameters/dummy/parameter.proto

package dummy

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type DummyParameter struct {
	ID   base.ID      `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3" json:"i_d"`
	Data base.DecData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (m *DummyParameter) Reset()      { *m = DummyParameter{} }
func (*DummyParameter) ProtoMessage() {}
func (*DummyParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa6cc99a8174b09, []int{0}
}
func (m *DummyParameter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DummyParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DummyParameter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DummyParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyParameter.Merge(m, src)
}
func (m *DummyParameter) XXX_Size() int {
	return m.Size()
}
func (m *DummyParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyParameter.DiscardUnknown(m)
}

var xxx_messageInfo_DummyParameter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DummyParameter)(nil), "persistence_sdk.modules.classifications.internal.parameters.dummy.DummyParameter")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/classifications/internal/parameters/dummy/parameter.proto", fileDescriptor_9fa6cc99a8174b09)
}

var fileDescriptor_9fa6cc99a8174b09 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd1, 0x3f, 0x4b, 0x3b, 0x31,
	0x18, 0x07, 0xf0, 0xa4, 0xbf, 0xf2, 0x43, 0x4e, 0x70, 0x28, 0x0e, 0xa5, 0x43, 0x2a, 0x82, 0x20,
	0x1d, 0x12, 0xd0, 0x4d, 0x5c, 0x2c, 0x59, 0xc4, 0xc1, 0x7f, 0x5b, 0x97, 0x92, 0x26, 0x8f, 0x6d,
	0xb0, 0xb9, 0x1c, 0x97, 0xdc, 0xd0, 0x77, 0xe0, 0xe8, 0xe8, 0x24, 0x7d, 0x39, 0x1d, 0x3b, 0x3a,
	0x89, 0xdc, 0xbd, 0x11, 0xb9, 0xdc, 0x95, 0x96, 0x5b, 0xd4, 0x2d, 0x09, 0xcf, 0xf7, 0x93, 0xe7,
	0x49, 0xa2, 0xfb, 0x04, 0x52, 0xa7, 0x9d, 0x87, 0x58, 0xc2, 0xd8, 0xa9, 0x67, 0x66, 0xac, 0xca,
	0xe6, 0xe0, 0x98, 0x9c, 0x0b, 0xe7, 0xf4, 0x93, 0x96, 0xc2, 0x6b, 0x1b, 0x3b, 0xa6, 0x63, 0x0f,
	0x69, 0x2c, 0xe6, 0x2c, 0x11, 0xa9, 0x30, 0xe0, 0x21, 0x75, 0x4c, 0x65, 0xc6, 0x2c, 0xb6, 0x07,
	0x34, 0x49, 0xad, 0xb7, 0x9d, 0xab, 0x06, 0x49, 0x6b, 0x92, 0x36, 0x48, 0xba, 0x21, 0xe9, 0x96,
	0xa4, 0x81, 0xec, 0x1d, 0x4e, 0xed, 0xd4, 0x06, 0x8d, 0x95, 0xab, 0x0a, 0xee, 0x0d, 0x9a, 0xbd,
	0x3a, 0x39, 0x03, 0x23, 0x98, 0x5f, 0x24, 0xe0, 0xd8, 0x44, 0x38, 0x60, 0x5a, 0xd5, 0xb5, 0xec,
	0xe7, 0x5a, 0x05, 0x92, 0x0b, 0x2f, 0xaa, 0xc0, 0xf1, 0x3b, 0x8e, 0x0e, 0x78, 0x79, 0xf9, 0xdd,
	0xa6, 0x99, 0xce, 0x65, 0xf4, 0x4f, 0x8f, 0x55, 0x17, 0x1f, 0xe1, 0xd3, 0xfd, 0xb3, 0x13, 0xda,
	0x1c, 0xab, 0x12, 0x69, 0x10, 0x69, 0x29, 0xd2, 0x6b, 0x3e, 0x6c, 0xaf, 0x3e, 0xfb, 0xe8, 0xa1,
	0xa5, 0x79, 0x87, 0x47, 0x6d, 0x25, 0xbc, 0xe8, 0xb6, 0x42, 0x7c, 0xf0, 0x8b, 0x38, 0xaf, 0x1a,
	0xaa, 0x8d, 0x90, 0xbe, 0xd8, 0x7b, 0x59, 0xf6, 0xd1, 0xdb, 0xb2, 0x8f, 0x86, 0x7e, 0x95, 0x13,
	0xbc, 0xce, 0x09, 0xfe, 0xca, 0x09, 0x7e, 0x2d, 0x08, 0x5a, 0x17, 0x04, 0x7d, 0x14, 0x04, 0x8d,
	0x46, 0x53, 0xed, 0x67, 0xd9, 0x84, 0x4a, 0x6b, 0x76, 0xc7, 0xbe, 0x8d, 0x61, 0x77, 0xfb, 0xc8,
	0x6f, 0xfe, 0xfe, 0xb9, 0x93, 0xff, 0xe1, 0x75, 0xce, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe0,
	0xa3, 0x2c, 0x15, 0x28, 0x02, 0x00, 0x00,
}

func (m *DummyParameter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DummyParameter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DummyParameter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParameter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParameter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParameter(dAtA []byte, offset int, v uint64) int {
	offset -= sovParameter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DummyParameter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovParameter(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovParameter(uint64(l))
	return n
}

func sovParameter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParameter(x uint64) (n int) {
	return sovParameter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DummyParameter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParameter
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
			return fmt.Errorf("proto: DummyParameter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DummyParameter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameter
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
				return ErrInvalidLengthParameter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParameter
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParameter
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
				return ErrInvalidLengthParameter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParameter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParameter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParameter
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
func skipParameter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParameter
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
					return 0, ErrIntOverflowParameter
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
					return 0, ErrIntOverflowParameter
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
				return 0, ErrInvalidLengthParameter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParameter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParameter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParameter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParameter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParameter = fmt.Errorf("proto: unexpected end of group")
)
