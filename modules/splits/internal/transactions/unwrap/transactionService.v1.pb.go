// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/splits/internal/transactions/unwrap/transactionService.v1.proto

package unwrap

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() {
	proto.RegisterFile("modules/splits/internal/transactions/unwrap/transactionService.v1.proto", fileDescriptor_09db33f10ab80470)
}

var fileDescriptor_09db33f10ab80470 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x3f, 0x4b, 0xf3, 0x40,
	0x1c, 0x6e, 0xee, 0x85, 0x0e, 0x79, 0xb7, 0x82, 0x88, 0x47, 0xbd, 0xa1, 0x82, 0xe3, 0x1d, 0xd5,
	0x2d, 0xb8, 0xd8, 0xa5, 0x22, 0x14, 0x8a, 0x4d, 0x1d, 0xa4, 0xcb, 0xaf, 0xe9, 0x11, 0x0f, 0x92,
	0xbb, 0x90, 0xbb, 0xd6, 0xc9, 0xc5, 0xd9, 0x41, 0xf0, 0x1b, 0xb8, 0xe9, 0x27, 0x11, 0xa7, 0x82,
	0x8b, 0xa3, 0x24, 0x4e, 0x7e, 0x0a, 0x31, 0x47, 0x48, 0xa0, 0xb4, 0x50, 0xd7, 0x7b, 0x9e, 0xe7,
	0x9e, 0x3f, 0x3f, 0xb7, 0x1f, 0xab, 0xd9, 0x3c, 0xe2, 0x9a, 0xe9, 0x24, 0x12, 0x46, 0x33, 0x21,
	0x0d, 0x4f, 0x25, 0x44, 0xcc, 0xa4, 0x20, 0x35, 0x04, 0x46, 0x28, 0xa9, 0xd9, 0x5c, 0xde, 0xa4,
	0x90, 0xd4, 0xdf, 0x46, 0x3c, 0x5d, 0x88, 0x80, 0xd3, 0x45, 0x97, 0x26, 0xa9, 0x32, 0xaa, 0x85,
	0xed, 0x07, 0xb4, 0xae, 0xa3, 0x56, 0x87, 0xdb, 0xa1, 0x52, 0x61, 0xc4, 0x19, 0x24, 0x82, 0x81,
	0x94, 0xca, 0x80, 0x85, 0x0b, 0x25, 0x3e, 0xd9, 0x26, 0x42, 0xcc, 0xb5, 0x86, 0xb0, 0xf2, 0xc5,
	0x67, 0x7f, 0x2c, 0x70, 0xc1, 0x75, 0xa2, 0xa4, 0xae, 0x7e, 0x3a, 0xba, 0x77, 0xdc, 0xff, 0x7e,
	0x45, 0x68, 0xdd, 0xba, 0xcd, 0x71, 0x21, 0x6b, 0x1d, 0xd0, 0xf5, 0xe5, 0xe8, 0xc0, 0x26, 0xc2,
	0x6c, 0x13, 0xc9, 0x5f, 0x35, 0xee, 0xec, 0xdf, 0xbd, 0x7f, 0x3d, 0xa2, 0xdd, 0xce, 0x0e, 0x8b,
	0x41, 0x9a, 0x88, 0x97, 0x0d, 0xac, 0xa4, 0xf7, 0x8c, 0x5e, 0x33, 0xe2, 0x2c, 0x33, 0xe2, 0x7c,
	0x66, 0xc4, 0x79, 0xc8, 0x49, 0x63, 0x99, 0x93, 0xc6, 0x47, 0x4e, 0x1a, 0x2e, 0x09, 0x54, 0xbc,
	0xc1, 0xad, 0xb7, 0xe7, 0xaf, 0x1c, 0xea, 0xb2, 0x3b, 0xfc, 0x2d, 0x39, 0x74, 0xae, 0xce, 0x43,
	0x61, 0xae, 0xe7, 0x53, 0x1a, 0xa8, 0x98, 0x9d, 0x6a, 0xcd, 0xcd, 0xc0, 0xba, 0x97, 0x3b, 0x6e,
	0xb1, 0xe7, 0x13, 0xfa, 0x37, 0xf2, 0xc7, 0x2f, 0x08, 0x8f, 0x6c, 0x16, 0xbf, 0x9e, 0xc5, 0x6e,
	0xf7, 0x56, 0x82, 0x93, 0x3a, 0x38, 0xb1, 0x60, 0x86, 0x0e, 0xd7, 0x83, 0x93, 0xfe, 0xb0, 0x37,
	0xe0, 0x06, 0x66, 0x60, 0xe0, 0x1b, 0xb5, 0x2d, 0xd1, 0xf3, 0xea, 0x4c, 0xcf, 0xb3, 0xd4, 0x69,
	0xb3, 0xb8, 0xe0, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x8d, 0xc7, 0xb4, 0xce, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	Unwrap(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionClient struct {
	cc grpc1.ClientConn
}

func NewTransactionClient(cc grpc1.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) Unwrap(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/splits.transactions.unwrap.Transaction/Unwrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	Unwrap(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) Unwrap(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unwrap not implemented")
}

func RegisterTransactionServer(s grpc1.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_Unwrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Unwrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/splits.transactions.unwrap.Transaction/Unwrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Unwrap(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "splits.transactions.unwrap.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unwrap",
			Handler:    _Transaction_Unwrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/splits/internal/transactions/unwrap/transactionService.v1.proto",
}
