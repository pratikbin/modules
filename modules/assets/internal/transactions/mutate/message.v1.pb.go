// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/mutate/message.v1.proto

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
	AssetID               *base.AssetID       `protobuf:"bytes,3,opt,name=asset_i_d,json=assetID,proto3" json:"asset_i_d,omitempty"`
	MutableMetaProperties *base1.PropertyList `protobuf:"bytes,4,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3" json:"mutable_meta_properties,omitempty"`
	MutableProperties     *base1.PropertyList `protobuf:"bytes,5,opt,name=mutable_properties,json=mutableProperties,proto3" json:"mutable_properties,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_608bf2813b136bf2, []int{0}
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

func (m *Message) GetAssetID() *base.AssetID {
	if m != nil {
		return m.AssetID
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
	proto.RegisterType((*Message)(nil), "assets.transactions.mutate.Message")
}

func init() {
	proto.RegisterFile("modules/assets/internal/transactions/mutate/message.v1.proto", fileDescriptor_608bf2813b136bf2)
}

var fileDescriptor_608bf2813b136bf2 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0xec, 0xba, 0xeb, 0x8e, 0xa2, 0x38, 0x22, 0x96, 0x20, 0xa1, 0x28, 0x68, 0xbd,
	0xcc, 0xb0, 0x7a, 0x0b, 0x5e, 0x1a, 0x16, 0xa4, 0x6a, 0x20, 0x84, 0xc5, 0x83, 0x14, 0xc2, 0x24,
	0x19, 0xb7, 0x03, 0x99, 0x4c, 0xc9, 0xbc, 0x0a, 0x3d, 0xfb, 0x05, 0xfc, 0x04, 0x1e, 0x3c, 0xfa,
	0x49, 0xc4, 0x53, 0x8f, 0x1e, 0x25, 0xbd, 0xf9, 0x29, 0x24, 0x33, 0x69, 0x0c, 0x42, 0x0f, 0x7b,
	0x4a, 0xe0, 0xfd, 0xfe, 0xef, 0xfd, 0xff, 0x6f, 0x1e, 0x7a, 0x25, 0x55, 0xb1, 0x2e, 0xb9, 0xa6,
	0x4c, 0x6b, 0x0e, 0x9a, 0x8a, 0x0a, 0x78, 0x5d, 0xb1, 0x92, 0x42, 0xcd, 0x2a, 0xcd, 0x72, 0x10,
	0xaa, 0xd2, 0x54, 0xae, 0x81, 0x01, 0xa7, 0x92, 0x6b, 0xcd, 0xae, 0x38, 0xf9, 0x74, 0x4e, 0x56,
	0xb5, 0x02, 0x85, 0x3d, 0xab, 0x22, 0x43, 0x98, 0x58, 0xd8, 0x9b, 0xe8, 0x7c, 0xc9, 0x25, 0xa3,
	0xa2, 0xd0, 0x34, 0x63, 0x9a, 0xdb, 0x09, 0xf3, 0x8b, 0x5e, 0xed, 0x3d, 0xf9, 0x9f, 0x10, 0x05,
	0xaf, 0x40, 0xc0, 0x66, 0x08, 0x3d, 0xeb, 0xa0, 0x52, 0x68, 0xe8, 0xb0, 0x55, 0xad, 0x56, 0xbc,
	0x86, 0xcd, 0x3b, 0xa1, 0xa1, 0x07, 0x1f, 0x7f, 0x76, 0xd1, 0x69, 0x64, 0x0d, 0x62, 0x8c, 0x8e,
	0x3f, 0xd6, 0x4a, 0x8e, 0x9d, 0x89, 0x33, 0x3d, 0x4b, 0xcc, 0x3f, 0x7e, 0x8e, 0x6e, 0xb6, 0xdf,
	0x54, 0xa4, 0xc5, 0xd8, 0x9d, 0x38, 0xd3, 0x5b, 0x2f, 0xee, 0x12, 0x51, 0x68, 0x32, 0xef, 0x87,
	0x26, 0x27, 0x2d, 0x30, 0xbf, 0xc0, 0x53, 0x74, 0x66, 0xcc, 0x1a, 0xf6, 0xc8, 0xb0, 0xb7, 0x0d,
	0x3b, 0xb3, 0x11, 0x92, 0xd3, 0x2e, 0x0b, 0x7e, 0x8b, 0x1e, 0xb6, 0x71, 0xb3, 0x92, 0xa7, 0x92,
	0x03, 0x4b, 0x3b, 0x6b, 0x82, 0xeb, 0xf1, 0xb1, 0xd1, 0xdd, 0x27, 0xc6, 0x38, 0x89, 0x07, 0x9e,
	0x93, 0x07, 0x9d, 0x26, 0xe2, 0xc0, 0xe2, 0x5e, 0x81, 0x43, 0x84, 0xf7, 0xcd, 0x06, 0x7d, 0x6e,
	0x1c, 0xee, 0x73, 0xaf, 0xc3, 0xff, 0xf5, 0x08, 0xbf, 0xba, 0x3f, 0x1a, 0xdf, 0xd9, 0x36, 0xbe,
	0xf3, 0xbb, 0xf1, 0x9d, 0x2f, 0x3b, 0x7f, 0xb4, 0xdd, 0xf9, 0xa3, 0x5f, 0x3b, 0x7f, 0x84, 0xfc,
	0x5c, 0x49, 0x72, 0xf8, 0xc1, 0xc2, 0x3b, 0xdd, 0xf6, 0xde, 0x9f, 0xc7, 0xed, 0x42, 0x63, 0xe7,
	0xc3, 0x9b, 0x2b, 0x01, 0xcb, 0x75, 0x46, 0x72, 0x25, 0xa9, 0x89, 0x1e, 0xb1, 0x0a, 0x4a, 0x4e,
	0xf7, 0x37, 0x73, 0x8d, 0xdb, 0xf9, 0xe6, 0x1e, 0xcd, 0x2e, 0xa3, 0xef, 0xae, 0x37, 0xb3, 0x06,
	0x2e, 0x87, 0x06, 0x22, 0x83, 0xfc, 0xdc, 0x17, 0x17, 0xc3, 0xe2, 0xc2, 0x16, 0x1b, 0xf7, 0xe9,
	0xe1, 0xe2, 0xe2, 0x75, 0x1c, 0xb6, 0xbb, 0x2c, 0x18, 0xb0, 0x3f, 0xee, 0x23, 0x0b, 0x06, 0xc1,
	0x90, 0x0c, 0x02, 0x8b, 0x66, 0x27, 0xe6, 0x5a, 0x5e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x06,
	0xdd, 0xa7, 0xf5, 0xf9, 0x02, 0x00, 0x00,
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
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.AssetID != nil {
		l = m.AssetID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
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
			if m.AssetID == nil {
				m.AssetID = &base.AssetID{}
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
