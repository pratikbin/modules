// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/metas/internal/transactions/reveal/message.v1.proto

package reveal

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/data/base"
	_ "github.com/AssetMantle/modules/schema/ids/base"
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

type Message struct {
	From string        `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Data *base.AnyData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2150b3df21f6910, []int{0}
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

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetData() *base.AnyData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "metas.transactions.reveal.Message")
}

func init() {
	proto.RegisterFile("modules/metas/internal/transactions/reveal/message.v1.proto", fileDescriptor_b2150b3df21f6910)
}

var fileDescriptor_b2150b3df21f6910 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0x80, 0x7b, 0xb1, 0x28, 0x46, 0x74, 0xc8, 0xd4, 0x56, 0x0c, 0xad, 0x22, 0x74, 0xba, 0xa3,
	0xba, 0xc5, 0xc5, 0x96, 0x82, 0x74, 0x08, 0x94, 0x50, 0x1c, 0xa4, 0xcb, 0x6b, 0x72, 0xb6, 0x81,
	0xdc, 0x9d, 0xe4, 0x5e, 0x0b, 0xfd, 0x17, 0x8e, 0xce, 0x8e, 0xfe, 0x12, 0x71, 0xea, 0xe8, 0x28,
	0xe9, 0xe6, 0xaf, 0x90, 0x5c, 0x4a, 0xc8, 0x52, 0x70, 0x4b, 0xf8, 0xbe, 0xc7, 0x77, 0xef, 0xd9,
	0x77, 0x42, 0x45, 0xcb, 0x84, 0x6b, 0x26, 0x38, 0x82, 0x66, 0xb1, 0x44, 0x9e, 0x4a, 0x48, 0x18,
	0xa6, 0x20, 0x35, 0x84, 0x18, 0x2b, 0xa9, 0x59, 0xca, 0x57, 0x1c, 0x12, 0x26, 0xb8, 0xd6, 0x30,
	0xe7, 0x74, 0xd5, 0xa3, 0x2f, 0xa9, 0x42, 0xe5, 0x34, 0xcd, 0x10, 0xad, 0xba, 0xb4, 0x70, 0x5b,
	0x1d, 0x1d, 0x2e, 0xb8, 0x00, 0x16, 0x01, 0x02, 0x9b, 0x81, 0xe6, 0x0c, 0xe4, 0x7a, 0x08, 0x08,
	0xe5, 0x74, 0xeb, 0x6a, 0xa7, 0xc4, 0x91, 0x2e, 0x8c, 0x38, 0xe2, 0x12, 0x63, 0x5c, 0x8f, 0x86,
	0xa5, 0x74, 0x79, 0x6f, 0x1f, 0xf9, 0x45, 0xd6, 0x71, 0xec, 0xfa, 0x73, 0xaa, 0x44, 0x83, 0xb4,
	0x49, 0xf7, 0x38, 0x30, 0xdf, 0x4e, 0xc7, 0xae, 0xe7, 0x85, 0x86, 0xd5, 0x26, 0xdd, 0x93, 0x9b,
	0x53, 0x9a, 0xff, 0xd0, 0x7e, 0x51, 0x0a, 0x0c, 0x1a, 0xbc, 0x59, 0x9f, 0x99, 0x4b, 0x36, 0x99,
	0x4b, 0x7e, 0x32, 0x97, 0xbc, 0x6e, 0xdd, 0xda, 0x66, 0xeb, 0xd6, 0xbe, 0xb7, 0x6e, 0xcd, 0xbe,
	0x08, 0x95, 0xa0, 0x7b, 0x77, 0x18, 0x9c, 0xed, 0xca, 0x8f, 0xbd, 0x71, 0xfe, 0x96, 0x31, 0x79,
	0x1a, 0xcd, 0x63, 0x5c, 0x2c, 0x67, 0x34, 0x54, 0x82, 0xf5, 0xb5, 0xe6, 0xe8, 0x83, 0xc4, 0x84,
	0xb3, 0xf2, 0x88, 0xff, 0x3e, 0xe6, 0xbb, 0x75, 0xe0, 0x4f, 0x82, 0x0f, 0xab, 0xe9, 0x9b, 0xfc,
	0xa4, 0x9a, 0x0f, 0x8c, 0xf1, 0xb5, 0x63, 0xd3, 0x2a, 0x9b, 0x16, 0x2c, 0xb3, 0xae, 0xf7, 0xb2,
	0xe9, 0xc3, 0x78, 0x90, 0xc3, 0x7c, 0xf9, 0x5f, 0xeb, 0xdc, 0x78, 0x9e, 0x57, 0x15, 0x3d, 0xaf,
	0x30, 0x67, 0x87, 0xe6, 0xc6, 0xb7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0xda, 0x84, 0x61,
	0x05, 0x02, 0x00, 0x00,
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
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMessageV1(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessageV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessageV1(v)
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
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	return n
}

func sovMessageV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessageV1(x uint64) (n int) {
	return sovMessageV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageV1
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
					return ErrIntOverflowMessageV1
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
				return ErrInvalidLengthMessageV1
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageV1
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
				return ErrInvalidLengthMessageV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &base.AnyData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessageV1
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
func skipMessageV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessageV1
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
					return 0, ErrIntOverflowMessageV1
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
					return 0, ErrIntOverflowMessageV1
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
				return 0, ErrInvalidLengthMessageV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessageV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessageV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessageV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessageV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessageV1 = fmt.Errorf("proto: unexpected end of group")
)
