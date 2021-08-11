// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/idData.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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

type IDData struct {
	Value github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=value,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"value"`
}

func (m *IDData) Reset()      { *m = IDData{} }
func (*IDData) ProtoMessage() {}
func (*IDData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b059d48eb657315, []int{0}
}
func (m *IDData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDData.Unmarshal(m, b)
}
func (m *IDData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDData.Marshal(b, m, deterministic)
}
func (m *IDData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDData.Merge(m, src)
}
func (m *IDData) XXX_Size() int {
	return xxx_messageInfo_IDData.Size(m)
}
func (m *IDData) XXX_DiscardUnknown() {
	xxx_messageInfo_IDData.DiscardUnknown(m)
}

var xxx_messageInfo_IDData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IDData)(nil), "schema.types.base.IDData")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/idData.proto", fileDescriptor_2b059d48eb657315)
}

var fileDescriptor_2b059d48eb657315 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4c,
	0x71, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xc8, 0xeb, 0x81,
	0xe5, 0xf5, 0x40, 0xf2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2,
	0x50, 0xa9, 0x88, 0x8b, 0xcd, 0xd3, 0x05, 0xa4, 0x51, 0x28, 0x8c, 0x8b, 0xb5, 0x2c, 0x31, 0xa7,
	0x34, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0xc9, 0xe1, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7,
	0xe4, 0x2d, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x91, 0x1c, 0xe1,
	0x9f, 0x97, 0x8a, 0xcc, 0x0d, 0x76, 0xf1, 0x46, 0x71, 0x92, 0x9e, 0xa7, 0x4b, 0x10, 0xc4, 0x38,
	0x2b, 0x81, 0x19, 0x0b, 0xe4, 0x19, 0x3a, 0x16, 0xca, 0x33, 0x4c, 0x58, 0x28, 0xcf, 0xb0, 0x60,
	0xa1, 0x3c, 0x83, 0x53, 0xc8, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59,
	0x91, 0x65, 0x19, 0xd8, 0xff, 0x49, 0x6c, 0x60, 0x0f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x0e, 0x2c, 0x42, 0x2b, 0x01, 0x00, 0x00,
}
