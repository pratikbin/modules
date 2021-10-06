// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/idData.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type IDData struct {
	Value ID `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
}

func (m *IDData) Reset()      { *m = IDData{} }
func (*IDData) ProtoMessage() {}
func (*IDData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b059d48eb657315, []int{0}
}
func (m *IDData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDData.Merge(m, src)
}
func (m *IDData) XXX_Size() int {
	return m.Size()
}
func (m *IDData) XXX_DiscardUnknown() {
	xxx_messageInfo_IDData.DiscardUnknown(m)
}

var xxx_messageInfo_IDData proto.InternalMessageInfo

func (m *IDData) GetValue() ID {
	if m != nil {
		return m.Value
	}
	return ID{}
}

func init() {
	proto.RegisterType((*IDData)(nil), "persistence_sdk.schema.types.base.IDData")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/idData.proto", fileDescriptor_2b059d48eb657315)
}

var fileDescriptor_2b059d48eb657315 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4c,
	0x71, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x44, 0x53, 0xaf, 0x07,
	0x51, 0xaf, 0x07, 0x56, 0xaf, 0x07, 0x52, 0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xad,
	0x0f, 0x62, 0x41, 0x34, 0x4a, 0x69, 0x11, 0x63, 0x11, 0x44, 0xad, 0x52, 0x20, 0x17, 0x9b, 0xa7,
	0x0b, 0xc8, 0x52, 0x21, 0x47, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x55, 0x3d, 0x82, 0xd6, 0xeb, 0x79, 0xba, 0x38, 0xb1, 0x9c, 0xb8, 0x27, 0xcf,
	0x10, 0x04, 0xd1, 0x69, 0xc5, 0x32, 0x63, 0x81, 0x3c, 0x83, 0x53, 0xc8, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x23, 0x99, 0xee, 0x9f, 0x97, 0x8a, 0xcc, 0x0d, 0x76, 0xf1, 0xc6, 0x74, 0x71, 0x12,
	0x1b, 0xd8, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xa3, 0xd0, 0x30, 0x46, 0x01,
	0x00, 0x00,
}

func (m *IDData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIdData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIdData(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IDData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Value.Size()
	n += 1 + l + sovIdData(uint64(l))
	return n
}

func sovIdData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdData(x uint64) (n int) {
	return sovIdData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IDData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdData
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
			return fmt.Errorf("proto: IDData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdData
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
				return ErrInvalidLengthIdData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdData
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
func skipIdData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdData
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
					return 0, ErrIntOverflowIdData
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
					return 0, ErrIntOverflowIdData
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
				return 0, ErrInvalidLengthIdData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdData = fmt.Errorf("proto: unexpected end of group")
)
