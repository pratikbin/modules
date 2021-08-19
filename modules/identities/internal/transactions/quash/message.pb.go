// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/identities/internal/transactions/quash/message.proto

package quash

import (
	context "context"
	fmt "fmt"
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
	From       github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required From missing"`
	FromID     github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d" valid:"required~required FromID missing"`
	IdentityID github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,3,opt,name=identity_i_d,json=identityID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"identity_i_d" valid:"required~required IdentityID missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0477d664af2a02, []int{0}
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
	return fileDescriptor_7a0477d664af2a02, []int{1}
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
	proto.RegisterType((*Message)(nil), "persistence_sdk.modules.identities.internal.transactions.quash.Message")
	proto.RegisterType((*TransactionResponse)(nil), "persistence_sdk.modules.identities.internal.transactions.quash.TransactionResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/identities/internal/transactions/quash/message.proto", fileDescriptor_7a0477d664af2a02)
}

var fileDescriptor_7a0477d664af2a02 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6b, 0x14, 0x41,
	0x14, 0xc7, 0x77, 0x3c, 0x89, 0xc9, 0x60, 0xb5, 0x2a, 0x84, 0x14, 0xbb, 0xc7, 0x16, 0x92, 0x42,
	0x66, 0x40, 0x1b, 0x49, 0x21, 0x24, 0x9c, 0xca, 0xa1, 0x41, 0x49, 0x04, 0xc1, 0xe6, 0x98, 0xdb,
	0x79, 0xd9, 0x1b, 0xdc, 0x9d, 0xd9, 0x9b, 0x37, 0x1b, 0x49, 0x63, 0x29, 0x82, 0x8d, 0xf8, 0x09,
	0xd2, 0x5a, 0xfa, 0x09, 0x6c, 0x53, 0xa6, 0x14, 0x8b, 0x45, 0xee, 0x1a, 0xeb, 0x7c, 0x02, 0xb9,
	0xc9, 0x9e, 0xb7, 0x88, 0x8a, 0x7a, 0x76, 0xff, 0x79, 0xc5, 0xef, 0xfd, 0x78, 0xf3, 0x1e, 0x7d,
	0x58, 0x82, 0x45, 0x85, 0x0e, 0x74, 0x0a, 0x03, 0x94, 0xcf, 0x79, 0x61, 0x64, 0x95, 0x03, 0x72,
	0x25, 0x41, 0x3b, 0xe5, 0xd4, 0x2c, 0x6a, 0x07, 0x56, 0x8b, 0x9c, 0x3b, 0x2b, 0x34, 0x8a, 0xd4,
	0x29, 0xa3, 0x91, 0x8f, 0x2b, 0x81, 0x23, 0x5e, 0x00, 0xa2, 0xc8, 0x80, 0x95, 0xd6, 0x38, 0x13,
	0xde, 0xf9, 0x81, 0xc6, 0x1a, 0x1a, 0x5b, 0xd0, 0xd8, 0x9c, 0xc6, 0xda, 0x34, 0xe6, 0x69, 0x1b,
	0x57, 0x33, 0x93, 0x19, 0x8f, 0xe2, 0xb3, 0x74, 0x4e, 0x4d, 0x3e, 0x76, 0xe8, 0xa5, 0xdd, 0xf3,
	0x3e, 0xe1, 0x1b, 0x42, 0x2f, 0x1e, 0x58, 0x53, 0xac, 0x93, 0x2e, 0xd9, 0x5c, 0xdb, 0x79, 0x71,
	0x52, 0xc7, 0xc1, 0xe7, 0x3a, 0xbe, 0x9b, 0x29, 0x37, 0xaa, 0x86, 0x2c, 0x35, 0x05, 0x6f, 0x39,
	0x3c, 0xd2, 0xd0, 0x7e, 0xee, 0xf7, 0x1e, 0x70, 0x4c, 0x47, 0x50, 0x08, 0xee, 0x8e, 0x4a, 0x40,
	0x3e, 0x14, 0x08, 0x6c, 0x3b, 0x4d, 0xb7, 0xa5, 0xb4, 0x80, 0x78, 0x56, 0xc7, 0xd7, 0x0f, 0x45,
	0xae, 0xe4, 0x56, 0x62, 0x61, 0x5c, 0x29, 0x0b, 0xf2, 0xe5, 0x3c, 0x74, 0xef, 0x59, 0x53, 0x74,
	0x0b, 0x85, 0xa8, 0x74, 0x96, 0xec, 0x79, 0x89, 0xf0, 0x15, 0xa1, 0xab, 0xb3, 0x30, 0x50, 0x03,
	0xb9, 0x7e, 0xc1, 0x1b, 0xe5, 0x8d, 0xd1, 0xed, 0x7f, 0x32, 0x62, 0xfd, 0xde, 0x59, 0x1d, 0x6f,
	0xfe, 0x56, 0xa2, 0xdf, 0x5b, 0x68, 0xac, 0x1c, 0xf8, 0x42, 0xf8, 0x8e, 0xd0, 0xcb, 0xcd, 0x8c,
	0x8f, 0xbc, 0x4c, 0xc7, 0xcb, 0x94, 0xff, 0x41, 0xe6, 0xc6, 0x2f, 0x65, 0xfa, 0x4d, 0xc3, 0xb6,
	0x10, 0x55, 0xdf, 0x8b, 0x5b, 0xab, 0xaf, 0x8f, 0xe3, 0xe0, 0xeb, 0x71, 0x1c, 0x24, 0xd7, 0xe8,
	0x95, 0x27, 0x8b, 0xdf, 0xde, 0x03, 0x2c, 0x8d, 0x46, 0xb8, 0xf9, 0x81, 0xd0, 0xce, 0x2e, 0x66,
	0xe1, 0x7b, 0x42, 0xd7, 0x1e, 0x5b, 0x73, 0xa8, 0x50, 0x19, 0x1d, 0xde, 0x67, 0xcb, 0x6d, 0x11,
	0x6b, 0x76, 0x65, 0x63, 0x7f, 0x59, 0xd0, 0x4f, 0x9c, 0x77, 0xc6, 0x27, 0x93, 0x88, 0x9c, 0x4e,
	0x22, 0xf2, 0x65, 0x12, 0x91, 0xb7, 0xd3, 0x28, 0x38, 0x9d, 0x46, 0xc1, 0xa7, 0x69, 0x14, 0x3c,
	0x7b, 0xfa, 0xc7, 0x43, 0xfe, 0xbb, 0x1b, 0x1b, 0xae, 0xf8, 0x33, 0xb8, 0xf5, 0x2d, 0x00, 0x00,
	0xff, 0xff, 0xa1, 0xc2, 0xd9, 0xfd, 0xac, 0x03, 0x00, 0x00,
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
	Provision(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Provision(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.identities.internal.transactions.quash.Msg/Provision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Provision(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Provision(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provision not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Provision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Provision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.identities.internal.transactions.quash.Msg/Provision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Provision(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.identities.internal.transactions.quash.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provision",
			Handler:    _Msg_Provision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/identities/internal/transactions/quash/message.proto",
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
		size := m.IdentityID.Size()
		i -= size
		if _, err := m.IdentityID.MarshalTo(dAtA[i:]); err != nil {
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
	l = m.IdentityID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
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
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
