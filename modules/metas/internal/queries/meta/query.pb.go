// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: modules/metas/internal/queries/meta/query.proto

package meta

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_modules_metas_internal_queries_meta_query_proto protoreflect.FileDescriptor

var file_modules_metas_internal_queries_meta_query_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x34, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x65,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x5c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x12, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x12, 0x23, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x64, 0x61,
	0x74, 0x61, 0x49, 0x44, 0x7d, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_modules_metas_internal_queries_meta_query_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),  // 0: meta.QueryRequest
	(*QueryResponse)(nil), // 1: meta.QueryResponse
}
var file_modules_metas_internal_queries_meta_query_proto_depIdxs = []int32{
	0, // 0: meta.Query.Meta:input_type -> meta.QueryRequest
	1, // 1: meta.Query.Meta:output_type -> meta.QueryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_metas_internal_queries_meta_query_proto_init() }
func file_modules_metas_internal_queries_meta_query_proto_init() {
	if File_modules_metas_internal_queries_meta_query_proto != nil {
		return
	}
	file_modules_metas_internal_queries_meta_queryRequest_proto_init()
	file_modules_metas_internal_queries_meta_queryResponse_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_metas_internal_queries_meta_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_metas_internal_queries_meta_query_proto_goTypes,
		DependencyIndexes: file_modules_metas_internal_queries_meta_query_proto_depIdxs,
	}.Build()
	File_modules_metas_internal_queries_meta_query_proto = out.File
	file_modules_metas_internal_queries_meta_query_proto_rawDesc = nil
	file_modules_metas_internal_queries_meta_query_proto_goTypes = nil
	file_modules_metas_internal_queries_meta_query_proto_depIdxs = nil
}
