// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: schema/ids/base/hashID.proto

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

type HashID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdBytes []byte `protobuf:"bytes,1,opt,name=id_bytes,json=idBytes,proto3" json:"id_bytes,omitempty"`
}

func (x *HashID) Reset() {
	*x = HashID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_ids_base_hashID_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashID) ProtoMessage() {}

func (x *HashID) ProtoReflect() protoreflect.Message {
	mi := &file_schema_ids_base_hashID_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashID.ProtoReflect.Descriptor instead.
func (*HashID) Descriptor() ([]byte, []int) {
	return file_schema_ids_base_hashID_proto_rawDescGZIP(), []int{0}
}

func (x *HashID) GetIdBytes() []byte {
	if x != nil {
		return x.IdBytes
	}
	return nil
}

var File_schema_ids_base_hashID_proto protoreflect.FileDescriptor

var file_schema_ids_base_hashID_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x06, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x69, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x77, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x42, 0x0b, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x64, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x42, 0x61, 0x73,
	0x65, 0xca, 0x02, 0x04, 0x42, 0x61, 0x73, 0x65, 0xe2, 0x02, 0x10, 0x42, 0x61, 0x73, 0x65, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x42, 0x61,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_ids_base_hashID_proto_rawDescOnce sync.Once
	file_schema_ids_base_hashID_proto_rawDescData = file_schema_ids_base_hashID_proto_rawDesc
)

func file_schema_ids_base_hashID_proto_rawDescGZIP() []byte {
	file_schema_ids_base_hashID_proto_rawDescOnce.Do(func() {
		file_schema_ids_base_hashID_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_ids_base_hashID_proto_rawDescData)
	})
	return file_schema_ids_base_hashID_proto_rawDescData
}

var file_schema_ids_base_hashID_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_ids_base_hashID_proto_goTypes = []interface{}{
	(*HashID)(nil), // 0: base.HashID
}
var file_schema_ids_base_hashID_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_ids_base_hashID_proto_init() }
func file_schema_ids_base_hashID_proto_init() {
	if File_schema_ids_base_hashID_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_ids_base_hashID_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashID); i {
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
			RawDescriptor: file_schema_ids_base_hashID_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_ids_base_hashID_proto_goTypes,
		DependencyIndexes: file_schema_ids_base_hashID_proto_depIdxs,
		MessageInfos:      file_schema_ids_base_hashID_proto_msgTypes,
	}.Build()
	File_schema_ids_base_hashID_proto = out.File
	file_schema_ids_base_hashID_proto_rawDesc = nil
	file_schema_ids_base_hashID_proto_goTypes = nil
	file_schema_ids_base_hashID_proto_depIdxs = nil
}
