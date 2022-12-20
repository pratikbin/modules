// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/lists/base/List.v1.proto

package base

import (
	fmt "fmt"
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

type List struct {
	// Types that are valid to be assigned to Impl:
	//	*List_DataList
	//	*List_IdList
	//	*List_PropertyList
	Impl isList_Impl `protobuf_oneof:"Impl"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_a577ea7e549d0e82, []int{0}
}
func (m *List) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_List.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return m.Size()
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

type isList_Impl interface {
	isList_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type List_DataList struct {
	DataList *DataList `protobuf:"bytes,1,opt,name=data_list,json=dataList,proto3,oneof" json:"data_list,omitempty"`
}
type List_IdList struct {
	IdList *IDList `protobuf:"bytes,2,opt,name=id_list,json=idList,proto3,oneof" json:"id_list,omitempty"`
}
type List_PropertyList struct {
	PropertyList *PropertyList `protobuf:"bytes,3,opt,name=property_list,json=propertyList,proto3,oneof" json:"property_list,omitempty"`
}

func (*List_DataList) isList_Impl()     {}
func (*List_IdList) isList_Impl()       {}
func (*List_PropertyList) isList_Impl() {}

func (m *List) GetImpl() isList_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *List) GetDataList() *DataList {
	if x, ok := m.GetImpl().(*List_DataList); ok {
		return x.DataList
	}
	return nil
}

func (m *List) GetIdList() *IDList {
	if x, ok := m.GetImpl().(*List_IdList); ok {
		return x.IdList
	}
	return nil
}

func (m *List) GetPropertyList() *PropertyList {
	if x, ok := m.GetImpl().(*List_PropertyList); ok {
		return x.PropertyList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*List) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*List_DataList)(nil),
		(*List_IdList)(nil),
		(*List_PropertyList)(nil),
	}
}

func init() {
	proto.RegisterType((*List)(nil), "lists.List")
}

func init() { proto.RegisterFile("schema/lists/base/List.v1.proto", fileDescriptor_a577ea7e549d0e82) }

var fileDescriptor_a577ea7e549d0e82 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0xcf, 0xc9, 0x2c, 0x2e, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xf7, 0xc9,
	0x2c, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xcb, 0x48,
	0x29, 0x63, 0xaa, 0x4b, 0x49, 0x2c, 0x49, 0x44, 0x51, 0x2b, 0xa5, 0x88, 0xa9, 0x28, 0x33, 0x05,
	0x55, 0x89, 0x3a, 0xa6, 0x92, 0x82, 0xa2, 0xfc, 0x82, 0xd4, 0xa2, 0x92, 0x4a, 0x14, 0x85, 0x4a,
	0x4b, 0x18, 0xb9, 0x58, 0x40, 0x22, 0x42, 0x7a, 0x5c, 0x9c, 0x20, 0x9b, 0xe2, 0x41, 0x3a, 0x24,
	0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xf5, 0xc0, 0xda, 0xf5, 0x5c, 0xa0, 0x2e, 0xf0, 0x60,
	0x08, 0xe2, 0x80, 0xb9, 0x46, 0x48, 0x83, 0x8b, 0x3d, 0x33, 0x05, 0xa2, 0x9a, 0x09, 0xac, 0x9a,
	0x17, 0xaa, 0xda, 0xd3, 0x05, 0xaa, 0x96, 0x0d, 0xe2, 0x28, 0x21, 0x2b, 0x2e, 0x5e, 0x98, 0xdd,
	0x10, 0xf5, 0xcc, 0x60, 0xf5, 0xc2, 0x50, 0xf5, 0x01, 0x48, 0xee, 0xf2, 0x60, 0x08, 0xe2, 0x41,
	0x76, 0xa7, 0x13, 0x1b, 0x17, 0x8b, 0x67, 0x6e, 0x41, 0x8e, 0xd3, 0x34, 0xc6, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0xe2, 0x4c, 0xce, 0xcf, 0x85, 0x18, 0xe5, 0xc4, 0x0d,
	0xd2, 0x13, 0x66, 0x18, 0x00, 0xf2, 0x59, 0x00, 0x63, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x63, 0x71, 0x71, 0x6a, 0x89, 0x6f, 0x62, 0x5e, 0x49, 0x4e,
	0xaa, 0x7e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0xb1, 0x3e, 0x46, 0x18, 0x2d, 0x62, 0x62, 0xf6,
	0x89, 0x88, 0x58, 0xc5, 0xc4, 0x0a, 0x32, 0xa7, 0xf8, 0x14, 0x94, 0x7e, 0xc4, 0x24, 0x08, 0xa6,
	0x63, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x41, 0x41, 0xf0, 0x0a, 0x2a, 0x97, 0xc4, 0x06,
	0x0e, 0x46, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x53, 0x5a, 0xab, 0xe1, 0x01, 0x00,
	0x00,
}

func (m *List) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *List) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *List) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Impl != nil {
		{
			size := m.Impl.Size()
			i -= size
			if _, err := m.Impl.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *List_DataList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *List_DataList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DataList != nil {
		{
			size, err := m.DataList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *List_IdList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *List_IdList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IdList != nil {
		{
			size, err := m.IdList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *List_PropertyList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *List_PropertyList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PropertyList != nil {
		{
			size, err := m.PropertyList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintListV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovListV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *List) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Impl != nil {
		n += m.Impl.Size()
	}
	return n
}

func (m *List_DataList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DataList != nil {
		l = m.DataList.Size()
		n += 1 + l + sovListV1(uint64(l))
	}
	return n
}
func (m *List_IdList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IdList != nil {
		l = m.IdList.Size()
		n += 1 + l + sovListV1(uint64(l))
	}
	return n
}
func (m *List_PropertyList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PropertyList != nil {
		l = m.PropertyList.Size()
		n += 1 + l + sovListV1(uint64(l))
	}
	return n
}

func sovListV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListV1(x uint64) (n int) {
	return sovListV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *List) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListV1
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
			return fmt.Errorf("proto: List: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: List: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListV1
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
				return ErrInvalidLengthListV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DataList{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &List_DataList{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListV1
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
				return ErrInvalidLengthListV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IDList{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &List_IdList{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListV1
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
				return ErrInvalidLengthListV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PropertyList{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &List_PropertyList{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListV1
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
func skipListV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListV1
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
					return 0, ErrIntOverflowListV1
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
					return 0, ErrIntOverflowListV1
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
				return 0, ErrInvalidLengthListV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListV1 = fmt.Errorf("proto: unexpected end of group")
)
