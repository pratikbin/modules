// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/deputize/service.v1.proto

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
	proto.RegisterFile("modules/identities/internal/transactions/deputize/service.v1.proto", fileDescriptor_ecd25bec3a6f017f)
}

var fileDescriptor_ecd25bec3a6f017f = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x9b, 0x08, 0x15, 0x32, 0x38, 0x74, 0x8c, 0x1a, 0x4a, 0x75, 0x51, 0xe4, 0x8e, 0xea,
	0x96, 0xcd, 0x58, 0xd0, 0x0e, 0x85, 0xa0, 0xc5, 0x41, 0xba, 0x5c, 0x93, 0x47, 0x3c, 0x48, 0xee,
	0x42, 0xee, 0xb5, 0x83, 0xa3, 0xf8, 0x01, 0x0a, 0x7e, 0x03, 0x47, 0x3f, 0x86, 0x93, 0x38, 0x15,
	0x5c, 0x1c, 0x25, 0x75, 0xf2, 0x53, 0x48, 0x9b, 0x86, 0xdc, 0x16, 0xea, 0x76, 0xc3, 0xef, 0xff,
	0xde, 0xef, 0xff, 0xce, 0xf2, 0x12, 0x19, 0x4e, 0x62, 0x50, 0x94, 0x87, 0x20, 0x90, 0x23, 0x5f,
	0x3e, 0x05, 0x42, 0x26, 0x58, 0x4c, 0x31, 0x63, 0x42, 0xb1, 0x00, 0xb9, 0x14, 0x8a, 0x86, 0x90,
	0x4e, 0x90, 0x3f, 0x00, 0x55, 0x90, 0x4d, 0x79, 0x00, 0x64, 0xda, 0x25, 0x69, 0x26, 0x51, 0xb6,
	0xda, 0x55, 0x96, 0xe8, 0x11, 0x52, 0x46, 0xec, 0xbd, 0x48, 0xca, 0x28, 0x06, 0xca, 0x52, 0x4e,
	0x99, 0x10, 0x12, 0x59, 0x01, 0xac, 0xf2, 0xf6, 0x3f, 0x1c, 0x12, 0x50, 0x8a, 0x45, 0x95, 0x83,
	0x7d, 0xb1, 0xf9, 0x8c, 0x0c, 0x54, 0x2a, 0x85, 0xaa, 0x86, 0x9c, 0xce, 0x0c, 0x6b, 0xfb, 0xa6,
	0x68, 0xd7, 0x7a, 0x32, 0xac, 0xe6, 0x15, 0x13, 0x61, 0x0c, 0xad, 0x23, 0x52, 0x57, 0x90, 0x0c,
	0x0a, 0x1f, 0xfb, 0xb8, 0x1e, 0xbd, 0x5e, 0xaf, 0xed, 0x1c, 0x3c, 0x7e, 0xfe, 0x3c, 0x9b, 0xfb,
	0x9d, 0x5d, 0x9a, 0x30, 0x81, 0x31, 0xe8, 0xea, 0x25, 0xed, 0xbd, 0x99, 0xef, 0xb9, 0x63, 0xcc,
	0x73, 0xc7, 0xf8, 0xce, 0x1d, 0x63, 0xb6, 0x70, 0x1a, 0xf3, 0x85, 0xd3, 0xf8, 0x5a, 0x38, 0x0d,
	0xeb, 0x30, 0x90, 0x49, 0xed, 0x3a, 0x6f, 0x67, 0x5d, 0xe8, 0xb6, 0xeb, 0x2f, 0x3b, 0xfa, 0xc6,
	0x9d, 0x1f, 0x71, 0xbc, 0x9f, 0x8c, 0x49, 0x20, 0x13, 0x7a, 0xae, 0x14, 0xe0, 0xa0, 0x58, 0x5f,
	0x5e, 0x70, 0xe3, 0x4b, 0xbe, 0x98, 0x5b, 0xfd, 0x61, 0xef, 0xd5, 0x6c, 0xf7, 0x2b, 0x99, 0xa1,
	0x2e, 0xd3, 0x5b, 0x83, 0x1f, 0x3a, 0x32, 0xd2, 0x91, 0x51, 0x89, 0xe4, 0xe6, 0x49, 0x1d, 0x32,
	0xba, 0xf4, 0xbd, 0x01, 0x20, 0x0b, 0x19, 0xb2, 0x5f, 0xb3, 0x53, 0xe1, 0xae, 0xab, 0xf3, 0xae,
	0x5b, 0x06, 0xc6, 0xcd, 0xd5, 0xf7, 0x9e, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xae, 0xde, 0x9d,
	0xa7, 0xed, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.deputize.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.deputize.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.deputize.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/identities/internal/transactions/deputize/service.v1.proto",
}
