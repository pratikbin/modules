// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/define/service.v1.proto

package define

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
	proto.RegisterFile("modules/assets/internal/transactions/define/service.v1.proto", fileDescriptor_28d0d726edc7c833)
}

var fileDescriptor_28d0d726edc7c833 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xfb, 0x50,
	0x14, 0xc5, 0x9b, 0xfc, 0xa1, 0x7f, 0xc8, 0xe0, 0x50, 0x10, 0x21, 0xd4, 0x37, 0x54, 0x71, 0x7c,
	0x8f, 0xea, 0x16, 0x74, 0x68, 0x29, 0x28, 0x42, 0xa1, 0x68, 0x71, 0x90, 0x2e, 0xb7, 0xc9, 0x35,
	0x06, 0x92, 0xf7, 0xca, 0x7b, 0xaf, 0xdd, 0x5c, 0xfc, 0x04, 0x82, 0xbb, 0x83, 0xa3, 0x9f, 0x44,
	0x9c, 0x0a, 0x2e, 0x8e, 0x92, 0x3a, 0xf9, 0x29, 0xa4, 0xbd, 0x2d, 0xcd, 0xd2, 0x42, 0xe6, 0xf3,
	0x3b, 0x9c, 0x73, 0xcf, 0xf5, 0x4e, 0x33, 0x15, 0x8d, 0x53, 0x34, 0x02, 0x8c, 0x41, 0x6b, 0x44,
	0x22, 0x2d, 0x6a, 0x09, 0xa9, 0xb0, 0x1a, 0xa4, 0x81, 0xd0, 0x26, 0x4a, 0x1a, 0x11, 0xe1, 0x5d,
	0x22, 0x51, 0x18, 0xd4, 0x93, 0x24, 0x44, 0x3e, 0x69, 0xf2, 0x91, 0x56, 0x56, 0xd5, 0x7c, 0x72,
	0xf1, 0x22, 0xcc, 0x09, 0xf6, 0xeb, 0xb1, 0x52, 0x71, 0x8a, 0x02, 0x46, 0x89, 0x00, 0x29, 0x95,
	0x05, 0x92, 0x17, 0x4e, 0xbf, 0x54, 0x6e, 0x86, 0xc6, 0x40, 0xbc, 0xce, 0xf5, 0xcf, 0xca, 0xb8,
	0x35, 0x9a, 0x91, 0x92, 0x66, 0x6d, 0x3f, 0x7e, 0xf0, 0xfe, 0x5f, 0xd3, 0x29, 0x35, 0xed, 0x55,
	0x2f, 0x40, 0x46, 0x29, 0xd6, 0x0e, 0xf8, 0xe6, 0x63, 0x78, 0x97, 0x1a, 0xf8, 0x87, 0xdb, 0xa0,
	0xab, 0x65, 0x50, 0x63, 0xff, 0xf1, 0xf3, 0xe7, 0xd9, 0xdd, 0x6b, 0xec, 0x8a, 0x0c, 0xa4, 0x9d,
	0x6f, 0x40, 0x35, 0x89, 0x6b, 0xbf, 0xb8, 0xef, 0x39, 0x73, 0xa6, 0x39, 0x73, 0xbe, 0x73, 0xe6,
	0x3c, 0xcd, 0x58, 0x65, 0x3a, 0x63, 0x95, 0xaf, 0x19, 0xab, 0x78, 0x2c, 0x54, 0xd9, 0x96, 0x88,
	0xf6, 0xce, 0xb2, 0xf7, 0x4d, 0xb3, 0x37, 0xbf, 0xa4, 0xe7, 0xdc, 0x5e, 0xc6, 0x89, 0xbd, 0x1f,
	0x0f, 0x79, 0xa8, 0x32, 0xd1, 0x9a, 0x1b, 0xbb, 0x14, 0xb9, 0x5a, 0xa8, 0xc4, 0x52, 0xaf, 0xee,
	0xbf, 0x56, 0xbf, 0xf3, 0xe6, 0xfa, 0x2d, 0x2a, 0xd0, 0x2f, 0x16, 0xe8, 0x2c, 0x90, 0x8f, 0x95,
	0x38, 0x28, 0x8a, 0x03, 0x12, 0x73, 0xf7, 0x68, 0xb3, 0x38, 0x38, 0xef, 0xb5, 0xbb, 0x68, 0x21,
	0x02, 0x0b, 0xbf, 0x6e, 0x9d, 0xc0, 0x20, 0x28, 0x92, 0x41, 0x40, 0xe8, 0xb0, 0xba, 0x78, 0xd3,
	0xc9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x9d, 0x18, 0x36, 0x9d, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assets.transactions.define.Service/Handle", in, out, opts...)
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
		FullMethod: "/assets.transactions.define.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.define.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/define/service.v1.proto",
}
