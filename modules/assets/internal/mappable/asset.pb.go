// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/assets/internal/mappable/asset.proto

package mappable

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/traits/base"
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

type Asset struct {
	ID            github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"i_d"`
	HasImmutables base.HasImmutables                                       `protobuf:"bytes,2,opt,name=has_immutables,json=hasImmutables,proto3" json:"has_immutables"`
	HasMutables   base.HasMutables                                         `protobuf:"bytes,3,opt,name=has_mutables,json=hasMutables,proto3" json:"has_mutables"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2afd895f76376e2, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetHasImmutables() base.HasImmutables {
	if m != nil {
		return m.HasImmutables
	}
	return base.HasImmutables{}
}

func (m *Asset) GetHasMutables() base.HasMutables {
	if m != nil {
		return m.HasMutables
	}
	return base.HasMutables{}
}

func init() {
	proto.RegisterType((*Asset)(nil), "assets.internal.mappable.Asset")
}

func init() {
	proto.RegisterFile("persistence_sdk/assets/internal/mappable/asset.proto", fileDescriptor_c2afd895f76376e2)
}

var fileDescriptor_c2afd895f76376e2 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e,
	0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4d, 0x2c, 0x28,
	0x48, 0x4c, 0xca, 0x49, 0x85, 0x48, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x40, 0x54,
	0xe9, 0xc1, 0x54, 0xe9, 0xc1, 0x54, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x15, 0xe9, 0x83,
	0x58, 0x10, 0xf5, 0x52, 0x66, 0xe8, 0xb6, 0x14, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0xea, 0x97, 0x14,
	0x25, 0x66, 0x96, 0x14, 0xeb, 0x27, 0x25, 0x16, 0xa7, 0xea, 0x67, 0x24, 0x16, 0x7b, 0xe6, 0xe6,
	0x96, 0x96, 0x80, 0x8c, 0x29, 0x86, 0xea, 0x33, 0x21, 0x4e, 0x9f, 0x2f, 0x8a, 0x2e, 0xa5, 0x26,
	0x26, 0x2e, 0x56, 0x47, 0x90, 0x03, 0x85, 0x02, 0xb9, 0x98, 0x33, 0xe3, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x9d, 0x1c, 0x4e, 0xdc, 0x93, 0x67, 0xb8, 0x75, 0x4f, 0xde, 0x22, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xc9, 0x7c, 0xff, 0xbc, 0x54, 0x64, 0x6e,
	0xb0, 0x8b, 0x37, 0xdc, 0xb6, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0x4f, 0x97, 0x20, 0xa6, 0x4c, 0x17,
	0x21, 0x3f, 0x2e, 0xbe, 0x8c, 0xc4, 0xe2, 0xf8, 0x4c, 0xb8, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35,
	0xb8, 0x8d, 0x14, 0xf5, 0x20, 0xaa, 0xf5, 0x20, 0x6e, 0xd3, 0x03, 0xb9, 0x4d, 0xcf, 0x03, 0xd9,
	0x4f, 0x4e, 0x2c, 0x20, 0x07, 0x04, 0xf1, 0xa2, 0x78, 0x54, 0xc8, 0x83, 0x8b, 0x07, 0x64, 0x1e,
	0xdc, 0x34, 0x66, 0xb0, 0x69, 0xf2, 0x38, 0x4c, 0xf3, 0x45, 0x35, 0x8b, 0x1b, 0xc9, 0xf3, 0x56,
	0x3c, 0x1d, 0x0b, 0xe5, 0x19, 0x26, 0x2c, 0x94, 0x67, 0x58, 0xb0, 0x50, 0x9e, 0xc1, 0x29, 0xf9,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x3c, 0x89, 0xf6, 0x7f, 0x6e, 0x7e,
	0x4a, 0x69, 0x4e, 0x6a, 0x31, 0xce, 0x34, 0x91, 0xc4, 0x06, 0x0e, 0x70, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x21, 0x54, 0x85, 0xb6, 0x46, 0x02, 0x00, 0x00,
}
