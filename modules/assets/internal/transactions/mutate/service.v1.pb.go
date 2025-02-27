// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/mutate/service.v1.proto

package mutate

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
	proto.RegisterFile("modules/assets/internal/transactions/mutate/service.v1.proto", fileDescriptor_0cf26871a6f7e232)
}

var fileDescriptor_0cf26871a6f7e232 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x73, 0x27, 0x44, 0xb8, 0xc2, 0x22, 0x20, 0xc2, 0x11, 0xb7, 0x88, 0x62, 0xb9, 0x4b,
	0xb4, 0x3b, 0xb4, 0x48, 0x1a, 0x45, 0x38, 0x08, 0x1a, 0x2c, 0x24, 0xcd, 0xe4, 0xb2, 0x9c, 0x07,
	0xb7, 0xbb, 0x61, 0x77, 0x92, 0xce, 0xc6, 0x5f, 0x20, 0xd8, 0x5b, 0x58, 0xfa, 0x4b, 0xc4, 0x2a,
	0x60, 0x63, 0x29, 0x17, 0x2b, 0x7f, 0x85, 0x24, 0x93, 0x90, 0x6b, 0x12, 0x48, 0xfd, 0xbe, 0xc7,
	0x7b, 0xf3, 0x26, 0x38, 0x57, 0x66, 0x30, 0xca, 0xa5, 0x13, 0xe0, 0x9c, 0x44, 0x27, 0x32, 0x8d,
	0xd2, 0x6a, 0xc8, 0x05, 0x5a, 0xd0, 0x0e, 0x12, 0xcc, 0x8c, 0x76, 0x42, 0x8d, 0x10, 0x50, 0x0a,
	0x27, 0xed, 0x38, 0x4b, 0x24, 0x1f, 0x37, 0xf9, 0xd0, 0x1a, 0x34, 0xb5, 0x90, 0x5c, 0xbc, 0x0c,
	0x73, 0x82, 0xc3, 0x7a, 0x6a, 0x4c, 0x9a, 0x4b, 0x01, 0xc3, 0x4c, 0x80, 0xd6, 0x06, 0x81, 0xe4,
	0xb9, 0x33, 0xdc, 0x2a, 0x57, 0x49, 0xe7, 0x20, 0x5d, 0xe5, 0x86, 0x17, 0xdb, 0xb8, 0xad, 0x74,
	0x43, 0xa3, 0xdd, 0xca, 0x7e, 0xfa, 0x18, 0xec, 0xde, 0xd2, 0x29, 0x35, 0x1b, 0x54, 0xaf, 0x40,
	0x0f, 0x72, 0x59, 0x3b, 0xe2, 0xeb, 0x8f, 0xe1, 0x31, 0x35, 0x08, 0x8f, 0x37, 0x41, 0x37, 0x8b,
	0xa0, 0xc6, 0xe1, 0xd3, 0xd7, 0xef, 0x8b, 0x7f, 0xd0, 0xd8, 0x17, 0x0a, 0x34, 0xce, 0x36, 0xa0,
	0x9a, 0xc4, 0xb5, 0x5f, 0xfd, 0x8f, 0x82, 0x79, 0x93, 0x82, 0x79, 0x3f, 0x05, 0xf3, 0x9e, 0xa7,
	0xac, 0x32, 0x99, 0xb2, 0xca, 0xf7, 0x94, 0x55, 0x02, 0x96, 0x18, 0xb5, 0x21, 0xa2, 0xbd, 0xb7,
	0xe8, 0x7d, 0xd7, 0xec, 0xcc, 0x2e, 0xe9, 0x78, 0xf7, 0xd7, 0x69, 0x86, 0x0f, 0xa3, 0x3e, 0x4f,
	0x8c, 0x12, 0xad, 0x99, 0x31, 0xa6, 0xc8, 0xe5, 0x42, 0x5b, 0x2c, 0xf5, 0xe6, 0xef, 0xb4, 0xba,
	0xf1, 0xbb, 0x1f, 0xb6, 0xa8, 0x40, 0xb7, 0x5c, 0x20, 0x9e, 0x23, 0x9f, 0x4b, 0xb1, 0x57, 0x16,
	0x7b, 0x24, 0x16, 0xfe, 0xc9, 0x7a, 0xb1, 0x77, 0xd9, 0x69, 0xc7, 0x12, 0x61, 0x00, 0x08, 0x7f,
	0x7e, 0x9d, 0xc0, 0x28, 0x2a, 0x93, 0x51, 0x44, 0x68, 0xbf, 0x3a, 0x7f, 0xd3, 0xd9, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x20, 0xfb, 0x7c, 0x66, 0x9d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/assets.transactions.mutate.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*Response, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assets.transactions.mutate.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.mutate.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/mutate/service.v1.proto",
}
