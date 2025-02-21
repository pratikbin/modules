// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/deputize/service.v1.proto

package deputize

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
	proto.RegisterFile("modules/assets/internal/transactions/deputize/service.v1.proto", fileDescriptor_0b9f2fc670836f8b)
}

var fileDescriptor_0b9f2fc670836f8b = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xfc, 0xa1, 0x7f, 0xc8, 0xe0, 0xd0, 0x45, 0x0c, 0xe5, 0x94, 0x82, 0xe2, 0x74,
	0x47, 0x75, 0xcb, 0xa0, 0xb4, 0x14, 0x74, 0xb0, 0x50, 0xb4, 0x38, 0x48, 0x97, 0xb7, 0xc9, 0x4b,
	0x3c, 0x48, 0xee, 0x42, 0xee, 0x5a, 0xc4, 0x51, 0xbf, 0x80, 0xe0, 0x37, 0x70, 0xec, 0x27, 0x11,
	0xa7, 0x82, 0x8b, 0xa3, 0xa4, 0x4e, 0x7e, 0x0a, 0x69, 0xaf, 0x47, 0x33, 0x05, 0x3a, 0x3f, 0xbf,
	0xe7, 0x9e, 0xe7, 0x7d, 0xce, 0x3b, 0x4b, 0x65, 0x34, 0x49, 0x50, 0x31, 0x50, 0x0a, 0xb5, 0x62,
	0x5c, 0x68, 0xcc, 0x05, 0x24, 0x4c, 0xe7, 0x20, 0x14, 0x84, 0x9a, 0x4b, 0xa1, 0x58, 0x84, 0xd9,
	0x44, 0xf3, 0x47, 0x64, 0x0a, 0xf3, 0x29, 0x0f, 0x91, 0x4e, 0xdb, 0x34, 0xcb, 0xa5, 0x96, 0x8d,
	0xa6, 0xf1, 0xd1, 0x32, 0x4e, 0x2d, 0xee, 0x37, 0x63, 0x29, 0xe3, 0x04, 0x19, 0x64, 0x9c, 0x81,
	0x10, 0x52, 0x83, 0x01, 0x56, 0x5e, 0x7f, 0xcb, 0xec, 0x14, 0x95, 0x82, 0x78, 0x93, 0xed, 0x9f,
	0x6f, 0xe7, 0xcf, 0x51, 0x65, 0x52, 0xa8, 0xcd, 0x03, 0x27, 0xcf, 0x8e, 0xf7, 0xff, 0xc6, 0x5c,
	0xd4, 0x78, 0xf0, 0xea, 0x97, 0x20, 0xa2, 0x04, 0x1b, 0x87, 0xb4, 0xea, 0x26, 0xda, 0x37, 0x35,
	0xfc, 0xa3, 0x6a, 0xec, 0x7a, 0x9d, 0xd6, 0xda, 0x7f, 0xfa, 0xfc, 0x79, 0x75, 0xf7, 0x5a, 0xbb,
	0x2c, 0x05, 0xa1, 0x97, 0x63, 0x98, 0xb6, 0x96, 0xec, 0xce, 0xdc, 0xf7, 0x82, 0x38, 0xf3, 0x82,
	0x38, 0xdf, 0x05, 0x71, 0x5e, 0x16, 0xa4, 0x36, 0x5f, 0x90, 0xda, 0xd7, 0x82, 0xd4, 0xbc, 0x83,
	0x50, 0xa6, 0x95, 0x31, 0xdd, 0x9d, 0x75, 0xff, 0xdb, 0xf6, 0x60, 0x79, 0xd2, 0xc0, 0xb9, 0xbb,
	0x8a, 0xb9, 0xbe, 0x9f, 0x8c, 0x69, 0x28, 0x53, 0xd6, 0x59, 0x5a, 0xfb, 0x26, 0xd6, 0x8e, 0xb5,
	0xd5, 0x68, 0x6f, 0xee, 0xbf, 0xce, 0xb0, 0x37, 0x73, 0x9b, 0x1d, 0x53, 0x62, 0x58, 0x2e, 0xd1,
	0x5b, 0x43, 0x1f, 0x56, 0x1e, 0x95, 0xe5, 0x91, 0x95, 0x0b, 0xf7, 0xb8, 0x4a, 0x1e, 0x5d, 0x0c,
	0xba, 0x7d, 0xd4, 0x10, 0x81, 0x86, 0x5f, 0x97, 0x18, 0x34, 0x08, 0xca, 0x6c, 0x10, 0x58, 0x78,
	0x5c, 0x5f, 0xfd, 0xdc, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x17, 0x4e, 0xc6, 0xb8,
	0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assets.transactions.deputize.Service/Handle", in, out, opts...)
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
		FullMethod: "/assets.transactions.deputize.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.deputize.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/deputize/service.v1.proto",
}
