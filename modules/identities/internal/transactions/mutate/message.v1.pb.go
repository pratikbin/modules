// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/mutate/message.v1.proto

package mutate

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
	From                  string              `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID                *base.IdentityID    `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	IdentityID            *base.IdentityID    `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3" json:"identity_i_d,omitempty"`
	MutableMetaProperties *base1.PropertyList `protobuf:"bytes,4,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties     *base1.PropertyList `protobuf:"bytes,5,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b226ac156422112, []int{0}
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

func (m *Message) GetIdentityID() *base.IdentityID {
	if m != nil {
		return m.IdentityID
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
	proto.RegisterType((*Message)(nil), "identities.transactions.mutate.Message")
}

func init() {
	proto.RegisterFile("modules/identities/internal/transactions/mutate/message.v1.proto", fileDescriptor_9b226ac156422112)
}

var fileDescriptor_9b226ac156422112 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x9b, 0xdc, 0xeb, 0x55, 0x47, 0x51, 0x1c, 0x11, 0xcb, 0x5d, 0x0c, 0xe5, 0xba, 0xb0,
	0xba, 0x98, 0xa1, 0xba, 0xcb, 0x4a, 0x43, 0x41, 0x82, 0x46, 0x42, 0x29, 0x2e, 0xa4, 0x10, 0x26,
	0xc9, 0xd8, 0x0e, 0x64, 0x32, 0x25, 0x73, 0x2a, 0xf4, 0x1d, 0x5c, 0xf8, 0x0c, 0x2e, 0x7d, 0x00,
	0x9f, 0x41, 0x5c, 0x75, 0xe9, 0x52, 0xd2, 0x9d, 0x4f, 0x21, 0xf9, 0x9f, 0x85, 0x16, 0x5c, 0xcd,
	0x59, 0xfc, 0xbe, 0xef, 0x7c, 0xe7, 0xcc, 0x41, 0x2f, 0x94, 0x4e, 0x76, 0xa9, 0x30, 0x4c, 0x26,
	0x22, 0x03, 0x09, 0xb2, 0x2c, 0x33, 0x10, 0x79, 0xc6, 0x53, 0x06, 0x39, 0xcf, 0x0c, 0x8f, 0x41,
	0xea, 0xcc, 0x30, 0xb5, 0x03, 0x0e, 0x82, 0x29, 0x61, 0x0c, 0x5f, 0x0b, 0xfa, 0x71, 0x46, 0xb7,
	0xb9, 0x06, 0x8d, 0x49, 0xaf, 0xa4, 0x43, 0x01, 0xad, 0x05, 0x97, 0x8f, 0x4c, 0xbc, 0x11, 0x8a,
	0x33, 0x99, 0x18, 0x16, 0x71, 0x23, 0xda, 0x4e, 0x7b, 0x6f, 0xde, 0x99, 0x5c, 0x3e, 0x6e, 0xa0,
	0x54, 0x1a, 0x68, 0xb0, 0x6d, 0xae, 0xb7, 0x22, 0x87, 0xfd, 0x1b, 0x69, 0xa0, 0x03, 0xaf, 0x3e,
	0xd9, 0xe8, 0xba, 0x5f, 0x47, 0xc0, 0x18, 0x9d, 0x7f, 0xc8, 0xb5, 0x1a, 0x5b, 0x13, 0x6b, 0x7a,
	0x73, 0x51, 0xd5, 0xf8, 0x09, 0xba, 0x51, 0xbe, 0xa1, 0x0c, 0x93, 0xb1, 0x3d, 0xb1, 0xa6, 0xb7,
	0x9e, 0xdd, 0xa5, 0x32, 0x31, 0xd4, 0xeb, 0x9a, 0x2e, 0x2e, 0x4a, 0xc0, 0x9b, 0xe3, 0x19, 0xba,
	0xdd, 0x46, 0xa9, 0xf0, 0xb3, 0xbf, 0xe3, 0xa8, 0xcf, 0x8b, 0x5f, 0xa3, 0x87, 0xe5, 0x54, 0x51,
	0x2a, 0x42, 0x25, 0x80, 0x87, 0x4d, 0x46, 0x29, 0xcc, 0xf8, 0xbc, 0x52, 0xdf, 0xa7, 0xd5, 0x04,
	0x34, 0x18, 0x84, 0x5f, 0x3c, 0x68, 0x34, 0xbe, 0x00, 0x1e, 0x74, 0x0a, 0xec, 0x22, 0xdc, 0x9a,
	0x0d, 0x7c, 0xae, 0xfd, 0xdb, 0xe7, 0x5e, 0x83, 0xf7, 0x1e, 0xee, 0x37, 0xfb, 0x7b, 0x41, 0xac,
	0x43, 0x41, 0xac, 0x5f, 0x05, 0xb1, 0x3e, 0x1f, 0xc9, 0xe8, 0x70, 0x24, 0xa3, 0x9f, 0x47, 0x32,
	0x42, 0x57, 0xb1, 0x56, 0xf4, 0xf4, 0xdf, 0xb8, 0x77, 0x9a, 0x55, 0xbe, 0x9b, 0x05, 0xe5, 0x76,
	0x03, 0xeb, 0xfd, 0xdb, 0xb5, 0x84, 0xcd, 0x2e, 0xa2, 0xb1, 0x56, 0xec, 0xa5, 0x31, 0x02, 0x7c,
	0x9e, 0x41, 0x2a, 0x58, 0x7b, 0x26, 0xff, 0x79, 0x2e, 0x5f, 0xec, 0x33, 0x6f, 0xe9, 0x7f, 0xb5,
	0x89, 0xd7, 0x07, 0x59, 0x0e, 0x83, 0xf8, 0x15, 0xf6, 0x63, 0x08, 0xac, 0x86, 0xc0, 0xaa, 0x06,
	0x0a, 0xfb, 0xe9, 0x69, 0x60, 0xf5, 0x2a, 0x70, 0xcb, 0xfd, 0x26, 0x1c, 0xf8, 0x6f, 0x7b, 0xd2,
	0xc3, 0x8e, 0x33, 0xa4, 0x1d, 0xa7, 0xc6, 0xa3, 0x8b, 0xea, 0x9c, 0x9e, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0xd6, 0x2d, 0xfd, 0x00, 0x03, 0x00, 0x00,
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
		dAtA[i] = 0x2a
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
		dAtA[i] = 0x22
	}
	if m.IdentityID != nil {
		{
			size, err := m.IdentityID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
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
			if m.IdentityID == nil {
				m.IdentityID = &base.IdentityID{}
			}
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
		case 5:
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
