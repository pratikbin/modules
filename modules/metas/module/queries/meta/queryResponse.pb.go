// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: modules/metas/module/queries/meta/queryResponse.proto

package meta

import (
	base "github.com/AssetMantle/modules/schema/data/base"
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

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	List    []*base.DataI `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_metas_module_queries_meta_queryResponse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_metas_module_queries_meta_queryResponse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_modules_metas_module_queries_meta_queryResponse_proto_rawDescGZIP(), []int{0}
}

func (x *QueryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *QueryResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *QueryResponse) GetList() []*base.DataI {
	if x != nil {
		return x.List
	}
	return nil
}

var File_modules_metas_module_queries_meta_queryResponse_proto protoreflect.FileDescriptor

var file_modules_metas_module_queries_meta_queryResponse_proto_rawDesc = []byte{
	0x0a, 0x35, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0d, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x90, 0x01,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x42, 0x12, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x4d, 0x65, 0x74, 0x61, 0xca,
	0x02, 0x04, 0x4d, 0x65, 0x74, 0x61, 0xe2, 0x02, 0x10, 0x4d, 0x65, 0x74, 0x61, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_metas_module_queries_meta_queryResponse_proto_rawDescOnce sync.Once
	file_modules_metas_module_queries_meta_queryResponse_proto_rawDescData = file_modules_metas_module_queries_meta_queryResponse_proto_rawDesc
)

func file_modules_metas_module_queries_meta_queryResponse_proto_rawDescGZIP() []byte {
	file_modules_metas_module_queries_meta_queryResponse_proto_rawDescOnce.Do(func() {
		file_modules_metas_module_queries_meta_queryResponse_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_metas_module_queries_meta_queryResponse_proto_rawDescData)
	})
	return file_modules_metas_module_queries_meta_queryResponse_proto_rawDescData
}

var file_modules_metas_module_queries_meta_queryResponse_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_modules_metas_module_queries_meta_queryResponse_proto_goTypes = []interface{}{
	(*QueryResponse)(nil), // 0: meta.QueryResponse
	(*base.DataI)(nil),    // 1: base.DataI
}
var file_modules_metas_module_queries_meta_queryResponse_proto_depIdxs = []int32{
	1, // 0: meta.QueryResponse.list:type_name -> base.DataI
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_modules_metas_module_queries_meta_queryResponse_proto_init() }
func file_modules_metas_module_queries_meta_queryResponse_proto_init() {
	if File_modules_metas_module_queries_meta_queryResponse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_metas_module_queries_meta_queryResponse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
			RawDescriptor: file_modules_metas_module_queries_meta_queryResponse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modules_metas_module_queries_meta_queryResponse_proto_goTypes,
		DependencyIndexes: file_modules_metas_module_queries_meta_queryResponse_proto_depIdxs,
		MessageInfos:      file_modules_metas_module_queries_meta_queryResponse_proto_msgTypes,
	}.Build()
	File_modules_metas_module_queries_meta_queryResponse_proto = out.File
	file_modules_metas_module_queries_meta_queryResponse_proto_rawDesc = nil
	file_modules_metas_module_queries_meta_queryResponse_proto_goTypes = nil
	file_modules_metas_module_queries_meta_queryResponse_proto_depIdxs = nil
}
