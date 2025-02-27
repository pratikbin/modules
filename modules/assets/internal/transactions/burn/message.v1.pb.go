// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/burn/message.v1.proto

package burn

import (
	fmt "fmt"
	base "github.com/AssetMantle/modules/schema/ids/base"
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
	From    string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID  *base.IdentityID `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3" json:"from_i_d,omitempty"`
	AssetID *base.AssetID    `protobuf:"bytes,3,opt,name=asset_i_d,json=assetID,proto3" json:"asset_i_d,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ad015adf6b2fae, []int{0}
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

func init() {
	proto.RegisterType((*Message)(nil), "assets.transactions.burn.Message")
}

func init() {
	proto.RegisterFile("modules/assets/internal/transactions/burn/message.v1.proto", fileDescriptor_83ad015adf6b2fae)
}

var fileDescriptor_83ad015adf6b2fae = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0x97, 0xee, 0x65, 0x7b, 0x57, 0x45, 0xa1, 0xa7, 0x32, 0xa4, 0x0c, 0xf5, 0x30, 0x2f,
	0x09, 0xd3, 0x5b, 0x6f, 0x2b, 0x03, 0xdd, 0x61, 0x30, 0xc6, 0xf0, 0x20, 0x83, 0x91, 0xb6, 0x71,
	0x0b, 0xac, 0x89, 0xe4, 0x49, 0x05, 0xbf, 0x85, 0xf8, 0x11, 0x3c, 0xfa, 0x49, 0xc4, 0xd3, 0x8e,
	0x1e, 0xa5, 0xbd, 0xf9, 0x29, 0xa4, 0xe9, 0x36, 0x8a, 0x30, 0xf0, 0x94, 0xc0, 0xff, 0xf7, 0x3c,
	0xbf, 0x27, 0x4f, 0x6c, 0x3f, 0x91, 0x71, 0xba, 0x62, 0x40, 0x28, 0x00, 0xd3, 0x40, 0xb8, 0xd0,
	0x4c, 0x09, 0xba, 0x22, 0x5a, 0x51, 0x01, 0x34, 0xd2, 0x5c, 0x0a, 0x20, 0x61, 0xaa, 0x04, 0x49,
	0x18, 0x00, 0x5d, 0x30, 0xfc, 0xd8, 0xc3, 0x0f, 0x4a, 0x6a, 0xe9, 0xb8, 0x65, 0x0d, 0xae, 0xa2,
	0xb8, 0x40, 0xdb, 0x67, 0x10, 0x2d, 0x59, 0x42, 0x09, 0x8f, 0x81, 0x84, 0x14, 0x18, 0xe1, 0x31,
	0x13, 0x9a, 0xeb, 0xa7, 0xe1, 0x60, 0x57, 0xde, 0xee, 0xfc, 0x86, 0x4c, 0xbb, 0x0a, 0x71, 0xaa,
	0xec, 0xe6, 0xa8, 0x94, 0x3a, 0x8e, 0xfd, 0xef, 0x5e, 0xc9, 0xc4, 0x45, 0x1d, 0xd4, 0x6d, 0x4d,
	0xcc, 0xdd, 0xb9, 0xb0, 0xff, 0x17, 0xe7, 0x9c, 0xcf, 0x63, 0xd7, 0xea, 0xa0, 0xee, 0xc1, 0xe5,
	0x31, 0xe6, 0x31, 0xe0, 0xe1, 0x4e, 0x36, 0x69, 0x14, 0xc0, 0x70, 0xe0, 0x74, 0xed, 0x96, 0xe9,
	0x6e, 0xd8, 0xba, 0x61, 0x0f, 0x0d, 0xdb, 0x2f, 0x9d, 0x93, 0xe6, 0x46, 0x1e, 0xbc, 0x58, 0xef,
	0x99, 0x87, 0xd6, 0x99, 0x87, 0xbe, 0x32, 0x0f, 0x3d, 0xe7, 0x5e, 0x6d, 0x9d, 0x7b, 0xb5, 0xcf,
	0xdc, 0xab, 0xd9, 0x27, 0x91, 0x4c, 0xf0, 0xbe, 0x37, 0x07, 0x47, 0x9b, 0x51, 0x6f, 0x7b, 0xe3,
	0x62, 0xf8, 0x31, 0xba, 0xbb, 0x59, 0x70, 0xbd, 0x4c, 0x43, 0x1c, 0xc9, 0x84, 0x18, 0xcf, 0x88,
	0x0a, 0xbd, 0x62, 0x64, 0xbb, 0xf2, 0x3f, 0xaf, 0xfe, 0xd5, 0xaa, 0xf7, 0xa7, 0xc1, 0x9b, 0xe5,
	0xf6, 0x4b, 0xf9, 0xb4, 0x2a, 0x0f, 0x52, 0x25, 0x3e, 0xb6, 0xd1, 0xac, 0x1a, 0xcd, 0x8a, 0x28,
	0xb3, 0xce, 0xf7, 0x45, 0xb3, 0xeb, 0x71, 0x30, 0x62, 0x9a, 0xc6, 0x54, 0xd3, 0x6f, 0xab, 0x5d,
	0x62, 0xbe, 0x5f, 0xe5, 0x7c, 0xbf, 0x00, 0xc3, 0x86, 0xf9, 0x8f, 0xab, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x48, 0xff, 0xe7, 0xba, 0x2e, 0x02, 0x00, 0x00,
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
