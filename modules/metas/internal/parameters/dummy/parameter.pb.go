// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/metas/internal/parameters/dummy/parameter.proto

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
	BaseParameter base.Parameter `protobuf:"bytes,1,opt,name=base_parameter,json=baseParameter,proto3" json:"base_parameter"`
}

func (m *DummyParameter) Reset()      { *m = DummyParameter{} }
func (*DummyParameter) ProtoMessage() {}
func (*DummyParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a688f7c03d55b51, []int{0}
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

func (m *DummyParameter) GetBaseParameter() base.Parameter {
	if m != nil {
		return m.BaseParameter
	}
	return base.Parameter{}
}

func init() {
	proto.RegisterType((*DummyParameter)(nil), "persistence_sdk.modules.metas.internal.parameters.dummy.DummyParameter")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/metas/internal/parameters/dummy/parameter.proto", fileDescriptor_7a688f7c03d55b51)
}

var fileDescriptor_7a688f7c03d55b51 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2f, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca,
	0x4b, 0xcc, 0xd1, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x2a, 0xd6, 0x4f, 0x29,
	0xcd, 0xcd, 0xad, 0x44, 0x08, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x99, 0xa3, 0x19, 0xa4,
	0x07, 0x35, 0x48, 0x0f, 0x6c, 0x90, 0x1e, 0xcc, 0x20, 0x3d, 0x84, 0x41, 0x7a, 0x60, 0x83, 0xa4,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x66, 0xe8, 0x83, 0x58, 0x10, 0xe3, 0xa4, 0x0c, 0xd1, 0xdd,
	0x55, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x9f, 0x94, 0x58,
	0x9c, 0x8a, 0xee, 0x02, 0xa5, 0x42, 0x2e, 0x3e, 0x17, 0x90, 0x89, 0x01, 0x30, 0x71, 0xa1, 0x48,
	0x2e, 0x3e, 0x90, 0xca, 0x78, 0xb8, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x1d, 0x3d,
	0x74, 0xc7, 0x42, 0x4c, 0xd7, 0x03, 0x9b, 0xae, 0x07, 0xd2, 0xa3, 0x07, 0x37, 0xc5, 0x89, 0xe5,
	0xc4, 0x3d, 0x79, 0x86, 0x20, 0x5e, 0x90, 0x28, 0x5c, 0xd0, 0x8a, 0x65, 0xc6, 0x02, 0x79, 0x06,
	0xa7, 0xac, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x0a, 0x48, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xb2, 0xcc, 0x3f, 0x2f, 0x15, 0x99, 0x1b,
	0xec, 0xe2, 0x4d, 0x6c, 0x80, 0x27, 0xb1, 0x81, 0x7d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x37, 0x1a, 0xe3, 0xb2, 0x01, 0x00, 0x00,
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
		size, err := m.BaseParameter.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.BaseParameter.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field BaseParameter", wireType)
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
			if err := m.BaseParameter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
