// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.8.0
// source: auht_service_simple_bank.proto

package pb

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

var File_service_simple_bank_proto protoreflect.FileDescriptor

var file_service_simple_bank_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70,
	0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xda, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x6e,
	0x6b, 0x12, 0x67, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x63, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42,
	0x1d, 0x5a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x63, 0x6c, 0x65, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_simple_bank_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),  // 0: proto.auth.CreateUserRequest
	(*LoginUserRequest)(nil),   // 1: proto.auth.LoginUserRequest
	(*CreateUserResponse)(nil), // 2: proto.auth.CreateUserResponse
	(*LoginUserResponse)(nil),  // 3: proto.auth.LoginUserResponse
}
var file_service_simple_bank_proto_depIdxs = []int32{
	0, // 0: proto.auth.SimpleBank.CreateUser:input_type -> proto.auth.CreateUserRequest
	1, // 1: proto.auth.SimpleBank.LoginUser:input_type -> proto.auth.LoginUserRequest
	2, // 2: proto.auth.SimpleBank.CreateUser:output_type -> proto.auth.CreateUserResponse
	3, // 3: proto.auth.SimpleBank.LoginUser:output_type -> proto.auth.LoginUserResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_simple_bank_proto_init() }
func file_service_simple_bank_proto_init() {
	if File_service_simple_bank_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_login_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_simple_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_simple_bank_proto_goTypes,
		DependencyIndexes: file_service_simple_bank_proto_depIdxs,
	}.Build()
	File_service_simple_bank_proto = out.File
	file_service_simple_bank_proto_rawDesc = nil
	file_service_simple_bank_proto_goTypes = nil
	file_service_simple_bank_proto_depIdxs = nil
}