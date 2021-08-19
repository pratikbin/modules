// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/splits/internal/transactions/unwrap/message.proto

package unwrap

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	github_com_persistenceOne_persistenceSDK_schema_types_base "github.com/persistenceOne/persistenceSDK/schema/types/base"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	From      github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required field From missing"`
	FromID    github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d" valid:"required~required field FromID missing"`
	OwnableID github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,3,opt,name=ownable_i_d,json=ownableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"ownable_i_d" valid:"required~required field OwnableID missing"`
	Value     github_com_cosmos_cosmos_sdk_types.Int                                `protobuf:"bytes,4,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"value" valid:"required~required field Value missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6898651feeee6a63, []int{0}
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

type TransactionResponse struct {
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6898651feeee6a63, []int{1}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "persistence_sdk.modules.splits.internal.transactions.unwrap.Message")
	proto.RegisterType((*TransactionResponse)(nil), "persistence_sdk.modules.splits.internal.transactions.unwrap.TransactionResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/splits/internal/transactions/unwrap/message.proto", fileDescriptor_6898651feeee6a63)
}

var fileDescriptor_6898651feeee6a63 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x34, 0x84, 0xf6, 0xd8, 0x0c, 0x48, 0x51, 0x07, 0x3b, 0xca, 0x80, 0x18, 0xe8,
	0x9d, 0x80, 0x05, 0x95, 0xa9, 0x55, 0x40, 0x8a, 0x50, 0x55, 0x54, 0x0a, 0x03, 0x4b, 0x74, 0xb1,
	0x5f, 0xdc, 0x13, 0xf6, 0x9d, 0xb9, 0x77, 0x6e, 0xc5, 0x00, 0x33, 0x03, 0x03, 0x0c, 0xec, 0x15,
	0x23, 0xdf, 0x03, 0xa9, 0x63, 0x47, 0xc4, 0x60, 0xa1, 0x64, 0x61, 0xce, 0x27, 0x40, 0x39, 0x3b,
	0x8a, 0x85, 0x90, 0x5a, 0x28, 0x93, 0x9f, 0xdf, 0xf0, 0xfb, 0xff, 0xa4, 0xfb, 0x3f, 0x3a, 0xc8,
	0xc1, 0xa0, 0x44, 0x0b, 0x2a, 0x82, 0x21, 0xc6, 0x2f, 0x79, 0xa6, 0xe3, 0x22, 0x05, 0xe4, 0x98,
	0xa7, 0xd2, 0x22, 0x97, 0xca, 0x82, 0x51, 0x22, 0xe5, 0xd6, 0x08, 0x85, 0x22, 0xb2, 0x52, 0x2b,
	0xe4, 0x85, 0x3a, 0x32, 0x22, 0xe7, 0x19, 0x20, 0x8a, 0x04, 0x58, 0x6e, 0xb4, 0xd5, 0xfe, 0x83,
	0xdf, 0x50, 0xac, 0x46, 0xb1, 0x0a, 0xc5, 0x16, 0x28, 0xd6, 0x44, 0xb1, 0x0a, 0xb5, 0x7e, 0x3d,
	0xd1, 0x89, 0x76, 0x1c, 0x3e, 0x9f, 0x2a, 0x64, 0xef, 0x6b, 0x8b, 0x5e, 0xd9, 0xa9, 0x42, 0xfc,
	0x8f, 0x84, 0xb6, 0xc6, 0x46, 0x67, 0x1d, 0xd2, 0x25, 0xb7, 0xd6, 0xb6, 0xdf, 0x9c, 0x94, 0xa1,
	0xf7, 0xbd, 0x0c, 0x1f, 0x26, 0xd2, 0x1e, 0x14, 0x23, 0x16, 0xe9, 0x8c, 0x37, 0x04, 0x76, 0x15,
	0x34, 0x7f, 0x9f, 0xf6, 0x1f, 0x73, 0x8c, 0x0e, 0x20, 0x13, 0xdc, 0xbe, 0xce, 0x01, 0xf9, 0x48,
	0x20, 0xb0, 0xad, 0x28, 0xda, 0x8a, 0x63, 0x03, 0x88, 0xb3, 0x32, 0xbc, 0x7d, 0x28, 0x52, 0x19,
	0x6f, 0xf6, 0x0c, 0xbc, 0x2a, 0xa4, 0x81, 0xf8, 0xed, 0x62, 0xe8, 0x8e, 0x25, 0xa4, 0x71, 0xf7,
	0x91, 0xd1, 0x59, 0x37, 0x93, 0x88, 0x52, 0x25, 0xbd, 0x3d, 0xa7, 0xe2, 0xbf, 0x27, 0x74, 0x75,
	0x3e, 0x0c, 0xe5, 0x30, 0xee, 0x5c, 0x72, 0x5e, 0xa6, 0xf6, 0xba, 0xff, 0x4f, 0x5e, 0x6c, 0xd0,
	0x9f, 0x95, 0x21, 0x3b, 0x87, 0xca, 0xa0, 0xbf, 0x94, 0x69, 0x8f, 0xdd, 0xc2, 0xff, 0x44, 0xe8,
	0x55, 0x7d, 0xa4, 0xc4, 0x28, 0x05, 0x67, 0xb4, 0xe2, 0x8c, 0x8a, 0xff, 0x60, 0x74, 0xe7, 0x0c,
	0xa3, 0xdd, 0x2a, 0xb5, 0x29, 0xb5, 0xa6, 0x17, 0x3b, 0xbf, 0xa0, 0x97, 0x0f, 0x45, 0x5a, 0x40,
	0xa7, 0xe5, 0x84, 0x86, 0xb5, 0xd0, 0xcd, 0x86, 0x50, 0xa4, 0x31, 0xd3, 0x58, 0x7f, 0x36, 0xe6,
	0x45, 0xac, 0xd3, 0x95, 0x9d, 0x95, 0xe1, 0xc6, 0x19, 0xf1, 0xcf, 0xe7, 0xf4, 0x65, 0x74, 0x95,
	0xb6, 0xb9, 0xfa, 0xee, 0x38, 0xf4, 0x7e, 0x1e, 0x87, 0x5e, 0xef, 0x06, 0xbd, 0xb6, 0xbf, 0x2c,
	0xdd, 0x1e, 0x60, 0xae, 0x15, 0xc2, 0xdd, 0x2f, 0x84, 0xae, 0xec, 0x60, 0xe2, 0x7f, 0x26, 0xb4,
	0xfd, 0xcc, 0xf5, 0xd0, 0xef, 0xb3, 0x0b, 0xb4, 0x98, 0xd5, 0x5d, 0x5d, 0x7f, 0x72, 0x21, 0xca,
	0x1f, 0x54, 0x7b, 0xde, 0xb6, 0x3a, 0x99, 0x04, 0xe4, 0x74, 0x12, 0x90, 0x1f, 0x93, 0x80, 0x7c,
	0x98, 0x06, 0xde, 0xe9, 0x34, 0xf0, 0xbe, 0x4d, 0x03, 0xef, 0xc5, 0xfe, 0xb9, 0x1f, 0xf6, 0x2f,
	0x8e, 0x7b, 0xd4, 0x76, 0x27, 0x78, 0xef, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x4f, 0x66,
	0xc1, 0x22, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Unwrap(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Unwrap(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.splits.internal.transactions.unwrap.Msg/Unwrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Unwrap(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Unwrap(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unwrap not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Unwrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unwrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.splits.internal.transactions.unwrap.Msg/Unwrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unwrap(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.splits.internal.transactions.unwrap.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unwrap",
			Handler:    _Msg_Unwrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/splits/internal/transactions/unwrap/message.proto",
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
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.OwnableID.Size()
		i -= size
		if _, err := m.OwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.From.Size()
		i -= size
		if _, err := m.From.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
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
	l = m.From.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.OwnableID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.Value.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *TransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
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
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: TransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
