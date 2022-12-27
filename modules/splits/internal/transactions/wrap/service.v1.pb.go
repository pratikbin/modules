// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/splits/internal/transactions/wrap/service.v1.proto

package wrap

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
	proto.RegisterFile("modules/splits/internal/transactions/wrap/service.v1.proto", fileDescriptor_e551e2cb0f856d6f)
}

var fileDescriptor_e551e2cb0f856d6f = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x9b, 0x08, 0x15, 0x32, 0x38, 0x54, 0x84, 0x12, 0x4b, 0xc0, 0xe2, 0x7c, 0x47, 0x75,
	0x8b, 0x93, 0x5d, 0xec, 0x52, 0x28, 0xb6, 0x58, 0x90, 0x2e, 0xaf, 0xe9, 0x11, 0x03, 0x97, 0xbb,
	0xe3, 0xde, 0xb5, 0xdd, 0xfd, 0x04, 0xa2, 0xdf, 0xc0, 0xd1, 0x4f, 0x22, 0x4e, 0x05, 0x17, 0x47,
	0x49, 0x9d, 0xfc, 0x14, 0x92, 0x5e, 0x43, 0xb3, 0x04, 0xea, 0xfc, 0xff, 0xfd, 0xef, 0xf7, 0xde,
	0x3b, 0x2f, 0x4c, 0xe5, 0x6c, 0xce, 0x19, 0x52, 0x54, 0x3c, 0x31, 0x48, 0x13, 0x61, 0x98, 0x16,
	0xc0, 0xa9, 0xd1, 0x20, 0x10, 0x22, 0x93, 0x48, 0x81, 0x74, 0xa9, 0x41, 0x51, 0x64, 0x7a, 0x91,
	0x44, 0x8c, 0x2c, 0x3a, 0x44, 0x69, 0x69, 0x64, 0xa3, 0x69, 0x3b, 0xa4, 0x8c, 0x92, 0x1c, 0xf5,
	0x5b, 0xb1, 0x94, 0x31, 0x67, 0x14, 0x54, 0x42, 0x41, 0x08, 0x69, 0xc0, 0x86, 0x9b, 0x9e, 0xff,
	0x0f, 0x67, 0xca, 0x10, 0x21, 0xde, 0x39, 0xfd, 0xab, 0xfd, 0xbb, 0x9a, 0xa1, 0x92, 0x02, 0x77,
	0xe5, 0x8b, 0xa5, 0x77, 0x38, 0xb4, 0x4b, 0x34, 0xb8, 0x57, 0xef, 0x81, 0x98, 0x71, 0xd6, 0x38,
	0x23, 0x55, 0x6b, 0x90, 0xbe, 0xb5, 0xfb, 0xed, 0x6a, 0xe4, 0x76, 0x2b, 0x69, 0x9f, 0x3e, 0x7e,
	0xfe, 0xbc, 0xb8, 0x27, 0xed, 0x63, 0x9a, 0x82, 0x30, 0x9c, 0x15, 0x03, 0xe6, 0x54, 0xf7, 0xd9,
	0x7d, 0xcf, 0x02, 0x67, 0x95, 0x05, 0xce, 0x77, 0x16, 0x38, 0x4f, 0xeb, 0xa0, 0xb6, 0x5a, 0x07,
	0xb5, 0xaf, 0x75, 0x50, 0xf3, 0x5a, 0x91, 0x4c, 0x2b, 0x9f, 0xef, 0x1e, 0x6d, 0xe7, 0xbd, 0xeb,
	0x0c, 0xf2, 0x0d, 0x06, 0xce, 0x7d, 0x2f, 0x4e, 0xcc, 0xc3, 0x7c, 0x4a, 0x22, 0x99, 0xd2, 0x6b,
	0x44, 0x66, 0xfa, 0x56, 0x57, 0xdc, 0x65, 0xef, 0xfb, 0xbc, 0xba, 0x07, 0xc3, 0xd1, 0xf8, 0xcd,
	0x6d, 0x0e, 0xad, 0x7c, 0x54, 0x96, 0x8f, 0x35, 0xa8, 0x8f, 0x22, 0x9a, 0x94, 0xa3, 0x49, 0x1e,
	0x65, 0xee, 0x79, 0x55, 0x34, 0xb9, 0x19, 0x74, 0xfb, 0xcc, 0xc0, 0x0c, 0x0c, 0xfc, 0xba, 0xbe,
	0xc5, 0xc2, 0xb0, 0xcc, 0x85, 0x61, 0x0e, 0x4e, 0xeb, 0x9b, 0x4f, 0xb9, 0xfc, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xce, 0x28, 0x48, 0x74, 0x83, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/splits.transactions.wrap.Service/Handle", in, out, opts...)
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
		FullMethod: "/splits.transactions.wrap.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "splits.transactions.wrap.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/splits/internal/transactions/wrap/service.v1.proto",
}
