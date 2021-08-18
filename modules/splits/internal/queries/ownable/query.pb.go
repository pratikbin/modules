// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/splits/internal/queries/ownable/query.proto

package ownable

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryRequest struct {
	OwnableID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=ownable_i_d,json=ownableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"ownable_i_d" valid:"required~required field OwnableID missing"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_845976b6477d0a88, []int{0}
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
	Success bool                                    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string                                  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" swaggertype:"string"`
	Value   *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"value,omitempty" swaggertype:"string"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_845976b6477d0a88, []int{1}
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
	proto.RegisterType((*QueryRequest)(nil), "persistence_sdk.modules.splits.internal.queries.ownable.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "persistence_sdk.modules.splits.internal.queries.ownable.QueryResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/splits/internal/queries/ownable/query.proto", fileDescriptor_845976b6477d0a88)
}

var fileDescriptor_845976b6477d0a88 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0x6f, 0x83, 0x42, 0x88, 0x81, 0xc6, 0x4a, 0x61, 0x9d, 0x90, 0x6d, 0xb9, 0x88, 0x68,
	0xd8, 0x15, 0x50, 0x80, 0x52, 0x1e, 0x07, 0x52, 0x44, 0x11, 0x38, 0x1a, 0x44, 0x73, 0xda, 0xb3,
	0x07, 0x67, 0xc1, 0xde, 0xf1, 0xed, 0xac, 0x13, 0xa5, 0xa1, 0xe0, 0x09, 0x90, 0x10, 0xaf, 0xc0,
	0x0b, 0xf0, 0x02, 0x94, 0x29, 0x28, 0x22, 0xd1, 0xa0, 0x14, 0x16, 0xba, 0xe3, 0x09, 0xf2, 0x04,
	0xe8, 0xfc, 0x47, 0xb2, 0x40, 0x48, 0x27, 0x52, 0x79, 0xc6, 0x1e, 0xff, 0xe6, 0x9b, 0xfd, 0x66,
	0x9d, 0x47, 0x05, 0x18, 0x52, 0x64, 0x41, 0xc7, 0x30, 0xa5, 0xe4, 0xad, 0xc8, 0x31, 0x29, 0x33,
	0x20, 0x41, 0x45, 0xa6, 0x2c, 0x09, 0xa5, 0x2d, 0x18, 0x2d, 0x33, 0x31, 0x2f, 0xc1, 0x28, 0x20,
	0x81, 0xc7, 0x5a, 0xce, 0x32, 0xa8, 0xf3, 0x13, 0x5e, 0x18, 0xb4, 0xe8, 0x3e, 0xf8, 0x03, 0xc2,
	0x5b, 0x08, 0x6f, 0x20, 0xbc, 0x83, 0xf0, 0x16, 0xc2, 0x5b, 0xc8, 0x70, 0x27, 0xc5, 0x14, 0x6b,
	0x86, 0x58, 0x45, 0x0d, 0x6e, 0x78, 0x2b, 0x45, 0x4c, 0x33, 0x10, 0xb2, 0x50, 0x42, 0x6a, 0x8d,
	0x56, 0x5a, 0x85, 0x9a, 0x9a, 0xaf, 0xd1, 0x67, 0xe6, 0xdc, 0x78, 0xbe, 0x6a, 0x3e, 0x81, 0x79,
	0x09, 0x64, 0xdd, 0x4f, 0xcc, 0xb9, 0xde, 0x02, 0xa7, 0x6a, 0x9a, 0x78, 0x2c, 0x64, 0xb7, 0xb7,
	0x47, 0xe5, 0x69, 0x15, 0x0c, 0xce, 0xab, 0xe0, 0x61, 0xaa, 0xec, 0x61, 0x39, 0xe3, 0x31, 0xe6,
	0xa2, 0x27, 0xf3, 0x40, 0x43, 0x3f, 0x7d, 0x31, 0x7e, 0x2a, 0x28, 0x3e, 0x84, 0x5c, 0x0a, 0x7b,
	0x52, 0x00, 0xf1, 0xfd, 0xf1, 0x45, 0x15, 0xdc, 0x3d, 0x92, 0x99, 0x4a, 0xf6, 0x22, 0x03, 0xf3,
	0x52, 0x19, 0x48, 0xde, 0x75, 0x41, 0xf8, 0x5a, 0x41, 0x96, 0x84, 0x07, 0x4d, 0xd7, 0xfd, 0x71,
	0x98, 0x2b, 0x22, 0xa5, 0xd3, 0x68, 0xb2, 0x8d, 0xdd, 0xbb, 0xe8, 0x0b, 0x73, 0x6e, 0xb6, 0x42,
	0xa9, 0x40, 0x4d, 0xe0, 0x7a, 0xce, 0x16, 0x95, 0x71, 0x0c, 0x44, 0xb5, 0xc8, 0x6b, 0x93, 0x2e,
	0x75, 0xb9, 0xb3, 0x09, 0xc6, 0xa0, 0xf1, 0x36, 0x6a, 0xf1, 0xde, 0x45, 0x15, 0xec, 0xd0, 0xb1,
	0x4c, 0x53, 0x30, 0x2b, 0x3d, 0x7b, 0x11, 0x59, 0x53, 0xf3, 0x9b, 0x32, 0xf7, 0xa5, 0xb3, 0x79,
	0x24, 0xb3, 0x12, 0xbc, 0x2b, 0x75, 0xfd, 0xe8, 0xbc, 0x0a, 0x76, 0x7b, 0x83, 0xc6, 0x48, 0x39,
	0x52, 0xfb, 0xb8, 0xb3, 0xb2, 0xb5, 0x99, 0x6a, 0x0c, 0xf1, 0xbf, 0xc9, 0x35, 0xf0, 0xde, 0x37,
	0xe6, 0x6c, 0xb5, 0x73, 0xb9, 0x5f, 0x7b, 0xf1, 0x63, 0xfe, 0x9f, 0x26, 0xf3, 0xbe, 0x59, 0xc3,
	0x27, 0x97, 0xc5, 0x34, 0x47, 0x19, 0xed, 0xbe, 0xff, 0xfe, 0xeb, 0xe3, 0x46, 0xe8, 0xfa, 0x7f,
	0xb9, 0xd8, 0xec, 0x6d, 0xfb, 0xd7, 0xe8, 0xcd, 0xe9, 0xc2, 0x67, 0x67, 0x0b, 0x9f, 0xfd, 0x5c,
	0xf8, 0xec, 0xc3, 0xd2, 0x1f, 0x9c, 0x2d, 0xfd, 0xc1, 0x8f, 0xa5, 0x3f, 0x78, 0xf5, 0x6c, 0xed,
	0xc5, 0x58, 0xf3, 0x4a, 0xcc, 0xae, 0xd6, 0x0b, 0x7a, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0xb7, 0xb1, 0x10, 0x54, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OwnableClient is the client API for Ownable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OwnableClient interface {
	Ownable(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type ownableClient struct {
	cc grpc1.ClientConn
}

func NewOwnableClient(cc grpc1.ClientConn) OwnableClient {
	return &ownableClient{cc}
}

func (c *ownableClient) Ownable(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.splits.internal.queries.ownable.Ownable/Ownable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnableServer is the server API for Ownable service.
type OwnableServer interface {
	Ownable(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedOwnableServer can be embedded to have forward compatible implementations.
type UnimplementedOwnableServer struct {
}

func (*UnimplementedOwnableServer) Ownable(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ownable not implemented")
}

func RegisterOwnableServer(s grpc1.Server, srv OwnableServer) {
	s.RegisterService(&_Ownable_serviceDesc, srv)
}

func _Ownable_Ownable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnableServer).Ownable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.splits.internal.queries.ownable.Ownable/Ownable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnableServer).Ownable(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ownable_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.splits.internal.queries.ownable.Ownable",
	HandlerType: (*OwnableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ownable",
			Handler:    _Ownable_Ownable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/splits/internal/queries/ownable/query.proto",
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
		size := m.OwnableID.Size()
		i -= size
		if _, err := m.OwnableID.MarshalTo(dAtA[i:]); err != nil {
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
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
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
	l = m.OwnableID.Size()
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
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovQuery(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
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
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Value = &v
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
