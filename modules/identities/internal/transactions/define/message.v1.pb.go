// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/define/message.v1.proto

package define

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
	base1 "github.com/AssetMantle/modules/schema/lists/base"
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
	From                    string              `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                  *base.IdentityID    `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	ImmutableMetaProperties *base1.PropertyList `protobuf:"bytes,3,opt,name=immutable_meta_properties,json=immutableMetaProperties,proto3" json:"immutable_meta_properties,omitempty"`
	ImmutableProperties     *base1.PropertyList `protobuf:"bytes,4,opt,name=immutable_properties,json=immutableProperties,proto3" json:"immutable_properties,omitempty"`
	MutableMetaProperties   *base1.PropertyList `protobuf:"bytes,5,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties       *base1.PropertyList `protobuf:"bytes,6,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bc54f84b92f4349, []int{0}
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

func (m *Message) GetFromID() *base.IdentityID {
	if m != nil {
		return m.FromID
	}
	return nil
}

func (m *Message) GetImmutableMetaProperties() *base1.PropertyList {
	if m != nil {
		return m.ImmutableMetaProperties
	}
	return nil
}

func (m *Message) GetImmutableProperties() *base1.PropertyList {
	if m != nil {
		return m.ImmutableProperties
	}
	return nil
}

func (m *Message) GetMutableMetaProperties() *base1.PropertyList {
	if m != nil {
		return m.MutableMetaProperties
	}
	return nil
}

func (m *Message) GetMutableProperties() *base1.PropertyList {
	if m != nil {
		return m.MutableProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "identities.transactions.define.Message")
}

func init() {
	proto.RegisterFile("modules/identities/internal/transactions/define/message.v1.proto", fileDescriptor_5bc54f84b92f4349)
}

var fileDescriptor_5bc54f84b92f4349 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x37, 0x69, 0x5d, 0x75, 0x04, 0xc5, 0x54, 0xe9, 0xda, 0xc3, 0xb0, 0xd4, 0x83, 0xd5,
	0xc3, 0x0c, 0xd5, 0xdb, 0x9e, 0x74, 0x59, 0x94, 0xa0, 0xab, 0x61, 0x29, 0x1e, 0x64, 0x21, 0x4c,
	0x92, 0xb7, 0xed, 0x40, 0x26, 0xb3, 0x64, 0xde, 0x15, 0xfa, 0x2d, 0xfc, 0x0c, 0x1e, 0xfd, 0x00,
	0x7e, 0x06, 0xf1, 0xd4, 0xa3, 0x47, 0xc9, 0xde, 0x3c, 0xfa, 0x09, 0x24, 0x7f, 0x9a, 0x0c, 0xd8,
	0x14, 0x7a, 0xca, 0x1c, 0x7e, 0xcf, 0xf3, 0xfe, 0xc8, 0xfb, 0x92, 0x97, 0x4a, 0x27, 0xeb, 0x14,
	0x0c, 0x97, 0x09, 0x64, 0x28, 0x51, 0x96, 0xcf, 0x0c, 0x21, 0xcf, 0x44, 0xca, 0x31, 0x17, 0x99,
	0x11, 0x31, 0x4a, 0x9d, 0x19, 0x9e, 0xc0, 0xb1, 0xcc, 0x80, 0x2b, 0x30, 0x46, 0x9c, 0x00, 0xfb,
	0x7c, 0xc8, 0x56, 0xb9, 0x46, 0xed, 0xd1, 0x2e, 0xc9, 0xec, 0x00, 0xab, 0x03, 0x7b, 0x8f, 0x4d,
	0x7c, 0x0a, 0x4a, 0x70, 0x99, 0x18, 0x1e, 0x09, 0x03, 0x17, 0x93, 0xce, 0xfc, 0x59, 0x5b, 0xb2,
	0xf7, 0xa4, 0x81, 0x52, 0x69, 0xb0, 0xc1, 0x56, 0xb9, 0x5e, 0x41, 0x8e, 0x67, 0xef, 0xa4, 0xc1,
	0x16, 0xdc, 0xff, 0xeb, 0x92, 0x9b, 0xf3, 0x5a, 0xc1, 0xf3, 0xc8, 0xf6, 0x71, 0xae, 0xd5, 0xc8,
	0x19, 0x3b, 0x07, 0xb7, 0x17, 0xd5, 0xdb, 0x7b, 0x4a, 0x6e, 0x95, 0xdf, 0x50, 0x86, 0xc9, 0xc8,
	0x1d, 0x3b, 0x07, 0x77, 0x9e, 0xdf, 0x63, 0x32, 0x31, 0xcc, 0x6f, 0x87, 0x2e, 0x86, 0x25, 0xe0,
	0xcf, 0xbc, 0x0f, 0xe4, 0x91, 0x54, 0x6a, 0x8d, 0x22, 0x4a, 0x21, 0x54, 0x80, 0x22, 0x6c, 0x46,
	0x4a, 0x30, 0xa3, 0xad, 0x2a, 0xbb, 0xc3, 0x2a, 0x21, 0x16, 0x58, 0x2e, 0x8b, 0xdd, 0x36, 0x35,
	0x07, 0x14, 0x41, 0x9b, 0xf1, 0x5e, 0x93, 0x07, 0x5d, 0xa1, 0xd5, 0xb5, 0xdd, 0xdf, 0xb5, 0xd3,
	0x06, 0xac, 0x9e, 0xb7, 0x64, 0xb7, 0x4f, 0xeb, 0x46, 0x7f, 0xd5, 0xc3, 0xcb, 0xa5, 0xa6, 0xc4,
	0xbb, 0x44, 0x69, 0xd8, 0xdf, 0x73, 0xff, 0x3f, 0xa1, 0xe9, 0x77, 0xf7, 0x47, 0x41, 0x9d, 0xf3,
	0x82, 0x3a, 0xbf, 0x0b, 0xea, 0x7c, 0xd9, 0xd0, 0xc1, 0xf9, 0x86, 0x0e, 0x7e, 0x6d, 0xe8, 0x80,
	0xec, 0xc7, 0x5a, 0xb1, 0xab, 0x2f, 0x60, 0x7a, 0xb7, 0x59, 0xd8, 0xc7, 0xc3, 0xa0, 0xdc, 0x61,
	0xe0, 0x7c, 0x7a, 0x7f, 0x22, 0xf1, 0x74, 0x1d, 0xb1, 0x58, 0x2b, 0xfe, 0xca, 0x18, 0xc0, 0xb9,
	0xc8, 0x30, 0x05, 0x7e, 0x71, 0x8c, 0xd7, 0x3c, 0xca, 0xaf, 0xee, 0x96, 0x7f, 0x34, 0xfb, 0xe6,
	0x52, 0xbf, 0x13, 0x39, 0xb2, 0x45, 0x66, 0x15, 0xf6, 0xd3, 0x06, 0x96, 0x36, 0xb0, 0xac, 0x81,
	0xc2, 0x7d, 0x76, 0x35, 0xb0, 0x7c, 0x13, 0x4c, 0xcb, 0xff, 0x9b, 0x08, 0x14, 0x7f, 0xdc, 0x71,
	0x07, 0x4f, 0x26, 0x36, 0x3d, 0x99, 0xd4, 0x78, 0x34, 0xac, 0x8e, 0xf6, 0xc5, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x89, 0x00, 0xde, 0x66, 0x03, 0x00, 0x00,
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
	if m.MutableProperties != nil {
		{
			size, err := m.MutableProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.MutableMetaProperties != nil {
		{
			size, err := m.MutableMetaProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ImmutableProperties != nil {
		{
			size, err := m.ImmutableProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ImmutableMetaProperties != nil {
		{
			size, err := m.ImmutableMetaProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessageV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FromID != nil {
		{
			size, err := m.FromID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.FromID != nil {
		l = m.FromID.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.ImmutableMetaProperties != nil {
		l = m.ImmutableMetaProperties.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.ImmutableProperties != nil {
		l = m.ImmutableProperties.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.MutableMetaProperties != nil {
		l = m.MutableMetaProperties.Size()
		n += 1 + l + sovMessageV1(uint64(l))
	}
	if m.MutableProperties != nil {
		l = m.MutableProperties.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
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
			if m.FromID == nil {
				m.FromID = &base.IdentityID{}
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
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
			if m.ImmutableMetaProperties == nil {
				m.ImmutableMetaProperties = &base1.PropertyList{}
			}
			if err := m.ImmutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
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
			if m.ImmutableProperties == nil {
				m.ImmutableProperties = &base1.PropertyList{}
			}
			if err := m.ImmutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
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
			if m.MutableMetaProperties == nil {
				m.MutableMetaProperties = &base1.PropertyList{}
			}
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
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
			if m.MutableProperties == nil {
				m.MutableProperties = &base1.PropertyList{}
			}
			if err := m.MutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
