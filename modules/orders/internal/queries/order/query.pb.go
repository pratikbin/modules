// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/orders/internal/queries/order/query.proto

package order

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_helpers "github.com/persistenceOne/persistenceSDK/schema/helpers"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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
	OrderID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=order_i_d,json=orderID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"order_i_d" valid:"required~required field OrderID missing"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51602469a6390915, []int{0}
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

type QueryResponse struct {
	Success bool                                                               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string                                                             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" swaggertype:"string"`
	List    []github_com_persistenceOne_persistenceSDK_schema_helpers.Mappable `protobuf:"bytes,3,rep,name=list,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/helpers.Mappable" json:"list,omitempty"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51602469a6390915, []int{1}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *QueryResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "persistence_sdk.persistence_sdk.modules.orders.internal.queries.order.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "persistence_sdk.persistence_sdk.modules.orders.internal.queries.order.QueryResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/orders/internal/queries/order/query.proto", fileDescriptor_51602469a6390915)
}

var fileDescriptor_51602469a6390915 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x18, 0x85, 0x63, 0x0a, 0x94, 0x46, 0xb0, 0x44, 0x1d, 0x22, 0x86, 0x24, 0xca, 0xd4, 0xc9, 0x1e,
	0x58, 0x50, 0x27, 0xa8, 0xc2, 0x50, 0x21, 0xa8, 0x08, 0x0b, 0x62, 0xa9, 0xd2, 0xe4, 0x27, 0xb1,
	0x48, 0xe2, 0xd4, 0xbf, 0x03, 0xea, 0xc2, 0x2b, 0xc0, 0x23, 0xf0, 0x0e, 0xbc, 0x44, 0xc7, 0x8e,
	0xa8, 0x43, 0x74, 0xd5, 0xbe, 0x41, 0x9f, 0xe0, 0x2a, 0x4e, 0x23, 0x55, 0x9d, 0x7a, 0xb7, 0xff,
	0x1c, 0xdb, 0xdf, 0x39, 0x96, 0x6d, 0xbe, 0xad, 0x40, 0x22, 0x47, 0x05, 0x65, 0x0c, 0x4b, 0x4c,
	0xbe, 0xb3, 0x42, 0x24, 0x75, 0x0e, 0xc8, 0x84, 0x4c, 0x40, 0x22, 0xe3, 0xa5, 0x02, 0x59, 0x46,
	0x39, 0x5b, 0xd7, 0x20, 0x79, 0xef, 0x6b, 0xb5, 0xa1, 0x95, 0x14, 0x4a, 0x58, 0xef, 0xae, 0x10,
	0xf4, 0x5a, 0x9f, 0x91, 0xb4, 0x43, 0xd2, 0x1e, 0x49, 0xcf, 0xc8, 0xce, 0x7f, 0x39, 0x4e, 0x45,
	0x2a, 0x34, 0x91, 0xb5, 0x53, 0x07, 0xf7, 0xff, 0x12, 0xf3, 0xf9, 0xa7, 0x36, 0x2c, 0x84, 0x75,
	0x0d, 0xa8, 0xac, 0xdf, 0xc4, 0x1c, 0xe9, 0x03, 0x4b, 0xbe, 0x4c, 0x6c, 0xe2, 0x91, 0xc9, 0x68,
	0x86, 0xdb, 0xc6, 0x35, 0xf6, 0x8d, 0xfb, 0x3a, 0xe5, 0x2a, 0xab, 0x57, 0x34, 0x16, 0x05, 0xbb,
	0x28, 0xb1, 0x28, 0xe1, 0x52, 0x7e, 0x0e, 0xde, 0x33, 0x8c, 0x33, 0x28, 0x22, 0xa6, 0x36, 0x15,
	0x20, 0x9d, 0x07, 0xa7, 0xc6, 0x65, 0x3f, 0xa2, 0x9c, 0x27, 0x53, 0x5f, 0xc2, 0xba, 0xe6, 0x12,
	0x92, 0x5f, 0xfd, 0xe0, 0x7d, 0xe3, 0x90, 0x27, 0xde, 0xa2, 0xcd, 0x9c, 0x07, 0x5e, 0xc1, 0x11,
	0x79, 0x99, 0xfa, 0xe1, 0x50, 0x74, 0x8e, 0xff, 0x8f, 0x98, 0x2f, 0xce, 0x15, 0xb1, 0x12, 0x25,
	0x82, 0x65, 0x9b, 0x43, 0xac, 0xe3, 0x18, 0x10, 0x75, 0xc1, 0x67, 0x61, 0x2f, 0x2d, 0x6a, 0x3e,
	0x01, 0x29, 0x85, 0xb4, 0x1f, 0xe9, 0xe2, 0xf6, 0xa9, 0x71, 0xc7, 0xf8, 0x33, 0x4a, 0x53, 0x90,
	0x6d, 0x97, 0xa9, 0x8f, 0x4a, 0x6a, 0x7a, 0xb7, 0xcd, 0xfa, 0x62, 0x3e, 0xce, 0x39, 0x2a, 0x7b,
	0xe0, 0x0d, 0x26, 0xa3, 0x59, 0xb0, 0x6f, 0xdc, 0x37, 0x0f, 0xbd, 0x63, 0x06, 0x79, 0xbb, 0x40,
	0x3f, 0x44, 0x55, 0x15, 0xad, 0x72, 0x08, 0x35, 0x71, 0x96, 0x6d, 0x0f, 0x0e, 0xd9, 0x1d, 0x1c,
	0x72, 0x77, 0x70, 0xc8, 0x9f, 0xa3, 0x63, 0xec, 0x8e, 0x8e, 0xf1, 0xff, 0xe8, 0x18, 0x5f, 0x3f,
	0xde, 0x9c, 0x70, 0xd3, 0x5f, 0x59, 0x3d, 0xd5, 0x2f, 0xf9, 0xea, 0x3e, 0x00, 0x00, 0xff, 0xff,
	0x9b, 0xa6, 0xa1, 0xd6, 0x6b, 0x02, 0x00, 0x00,
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
	{
		size := m.OrderID.Size()
		i -= size
		if _, err := m.OrderID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.List[iNdEx].Size()
				i -= size
				if _, err := m.List[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
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
	l = m.OrderID.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_persistenceOne_persistenceSDK_schema_helpers.Mappable
			m.List = append(m.List, v)
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
