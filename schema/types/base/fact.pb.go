// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/fact.proto

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

type Fact struct {
	HashId     github_com_persistenceOne_persistenceSDK_schema_types.ID         `protobuf:"bytes,1,opt,name=hash_id,json=hashId,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"hash_id"`
	TypeId     github_com_persistenceOne_persistenceSDK_schema_types.ID         `protobuf:"bytes,2,opt,name=type_id,json=typeId,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"type_id"`
	Signatures github_com_persistenceOne_persistenceSDK_schema_types.Signatures `protobuf:"bytes,3,opt,name=signatures,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Signatures" json:"signatures"`
}

func (m *Fact) Reset()         { *m = Fact{} }
func (m *Fact) String() string { return proto.CompactTextString(m) }
func (*Fact) ProtoMessage()    {}
func (*Fact) Descriptor() ([]byte, []int) {
	return fileDescriptor_559a77c85918f977, []int{0}
}
func (m *Fact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fact.Unmarshal(m, b)
}
func (m *Fact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fact.Marshal(b, m, deterministic)
}
func (m *Fact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fact.Merge(m, src)
}
func (m *Fact) XXX_Size() int {
	return xxx_messageInfo_Fact.Size(m)
}
func (m *Fact) XXX_DiscardUnknown() {
	xxx_messageInfo_Fact.DiscardUnknown(m)
}

var xxx_messageInfo_Fact proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Fact)(nil), "schema.types.base.Fact")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/fact.proto", fileDescriptor_559a77c85918f977)
}

var fileDescriptor_559a77c85918f977 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0x4b,
	0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xc8, 0xea, 0x81, 0x65, 0xf5,
	0x40, 0xb2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0x69,
	0x0b, 0x13, 0x17, 0x8b, 0x5b, 0x62, 0x72, 0x89, 0x50, 0x24, 0x17, 0x7b, 0x46, 0x62, 0x71, 0x46,
	0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xc3, 0x89, 0x7b, 0xf2, 0x0c, 0xb7,
	0xee, 0xc9, 0x5b, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0xb9,
	0xc1, 0x3f, 0x2f, 0x15, 0x99, 0x1b, 0xec, 0xe2, 0x8d, 0xe2, 0x22, 0x3d, 0x4f, 0x97, 0x20, 0x36,
	0x90, 0x81, 0x9e, 0x29, 0x20, 0xa3, 0x41, 0x62, 0x20, 0xa3, 0x99, 0xa8, 0x65, 0x34, 0x88, 0xe5,
	0x99, 0x22, 0x94, 0xc1, 0xc5, 0x55, 0x9c, 0x99, 0x9e, 0x97, 0x58, 0x52, 0x5a, 0x94, 0x5a, 0x2c,
	0xc1, 0x0c, 0x36, 0xdd, 0x03, 0x6a, 0xba, 0x03, 0x79, 0xa6, 0x07, 0xc3, 0xcd, 0x0b, 0x42, 0x32,
	0xdb, 0x8a, 0xa7, 0x63, 0xa1, 0x3c, 0xc3, 0x84, 0x85, 0xf2, 0x0c, 0x0b, 0x16, 0xca, 0x33, 0x38,
	0x85, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x15, 0x59, 0xb6, 0x82,
	0x23, 0x30, 0x89, 0x0d, 0x1c, 0x27, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0xae, 0x78,
	0x3c, 0xec, 0x01, 0x00, 0x00,
}
