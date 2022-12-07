// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: schema/ids/base/propertyID.proto

package base

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PropertyID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyID  *StringID `protobuf:"bytes,1,opt,name=key_i_d,json=keyID,proto3" json:"key_i_d,omitempty"`
	TypeID *StringID `protobuf:"bytes,2,opt,name=type_i_d,json=typeID,proto3" json:"type_i_d,omitempty"`
}

func (x *PropertyID) Reset() {
	*x = PropertyID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_ids_base_propertyID_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyID) ProtoMessage() {}

func (x *PropertyID) ProtoReflect() protoreflect.Message {
	mi := &file_schema_ids_base_propertyID_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyID.ProtoReflect.Descriptor instead.
func (*PropertyID) Descriptor() ([]byte, []int) {
	return file_schema_ids_base_propertyID_proto_rawDescGZIP(), []int{0}
}

func (x *PropertyID) GetKeyID() *StringID {
	if x != nil {
		return x.KeyID
	}
	return nil
}

func (x *PropertyID) GetTypeID() *StringID {
	if x != nil {
		return x.TypeID
	}
	return nil
}

var File_schema_ids_base_propertyID_proto protoreflect.FileDescriptor

var file_schema_ids_base_propertyID_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x5f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x44, 0x12, 0x28,
	0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x5f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x44,
	0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x7b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x42, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x44,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69,
	0x64, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x04,
	0x42, 0x61, 0x73, 0x65, 0xca, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0xe2, 0x02, 0x10, 0x42, 0x61,
	0x73, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x04, 0x42, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_ids_base_propertyID_proto_rawDescOnce sync.Once
	file_schema_ids_base_propertyID_proto_rawDescData = file_schema_ids_base_propertyID_proto_rawDesc
)

func file_schema_ids_base_propertyID_proto_rawDescGZIP() []byte {
	file_schema_ids_base_propertyID_proto_rawDescOnce.Do(func() {
		file_schema_ids_base_propertyID_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_ids_base_propertyID_proto_rawDescData)
	})
	return file_schema_ids_base_propertyID_proto_rawDescData
}

var file_schema_ids_base_propertyID_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_ids_base_propertyID_proto_goTypes = []interface{}{
	(*PropertyID)(nil), // 0: base.PropertyID
	(*StringID)(nil),   // 1: base.StringID
}
var file_schema_ids_base_propertyID_proto_depIdxs = []int32{
	1, // 0: base.PropertyID.key_i_d:type_name -> base.StringID
	1, // 1: base.PropertyID.type_i_d:type_name -> base.StringID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schema_ids_base_propertyID_proto_init() }
func file_schema_ids_base_propertyID_proto_init() {
	if File_schema_ids_base_propertyID_proto != nil {
		return
	}
	file_schema_ids_base_stringID_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_ids_base_propertyID_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_ids_base_propertyID_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_ids_base_propertyID_proto_goTypes,
		DependencyIndexes: file_schema_ids_base_propertyID_proto_depIdxs,
		MessageInfos:      file_schema_ids_base_propertyID_proto_msgTypes,
	}.Build()
	File_schema_ids_base_propertyID_proto = out.File
	file_schema_ids_base_propertyID_proto_rawDesc = nil
	file_schema_ids_base_propertyID_proto_goTypes = nil
	file_schema_ids_base_propertyID_proto_depIdxs = nil
}
