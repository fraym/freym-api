// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: sync/management/service.proto

package managementpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_sync_management_service_proto protoreflect.FileDescriptor

var file_sync_management_service_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xc2, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x29, 0x2e, 0x66,
	0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x4c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x27, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x66, 0x72, 0x65, 0x79,
	0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x09, 0x44, 0x72, 0x6f, 0x70, 0x4c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x27, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x66, 0x72, 0x65, 0x79,
	0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x4f,
	0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x66, 0x72, 0x65,
	0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x66, 0x72, 0x65, 0x79,
	0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_sync_management_service_proto_goTypes = []any{
	(*CreateLeaseRequest)(nil),   // 0: freym.sync.management.CreateLeaseRequest
	(*KeepLeaseRequest)(nil),     // 1: freym.sync.management.KeepLeaseRequest
	(*DropLeaseRequest)(nil),     // 2: freym.sync.management.DropLeaseRequest
	(*GetPeerNodesRequest)(nil),  // 3: freym.sync.management.GetPeerNodesRequest
	(*LockRequest)(nil),          // 4: freym.sync.management.LockRequest
	(*UnlockRequest)(nil),        // 5: freym.sync.management.UnlockRequest
	(*CreateLeaseResponse)(nil),  // 6: freym.sync.management.CreateLeaseResponse
	(*KeepLeaseResponse)(nil),    // 7: freym.sync.management.KeepLeaseResponse
	(*DropLeaseResponse)(nil),    // 8: freym.sync.management.DropLeaseResponse
	(*GetPeerNodesResponse)(nil), // 9: freym.sync.management.GetPeerNodesResponse
	(*LockResponse)(nil),         // 10: freym.sync.management.LockResponse
	(*UnlockResponse)(nil),       // 11: freym.sync.management.UnlockResponse
}
var file_sync_management_service_proto_depIdxs = []int32{
	0,  // 0: freym.sync.management.Service.CreateLease:input_type -> freym.sync.management.CreateLeaseRequest
	1,  // 1: freym.sync.management.Service.KeepLease:input_type -> freym.sync.management.KeepLeaseRequest
	2,  // 2: freym.sync.management.Service.DropLease:input_type -> freym.sync.management.DropLeaseRequest
	3,  // 3: freym.sync.management.Service.GetPeerNodes:input_type -> freym.sync.management.GetPeerNodesRequest
	4,  // 4: freym.sync.management.Service.Lock:input_type -> freym.sync.management.LockRequest
	5,  // 5: freym.sync.management.Service.Unlock:input_type -> freym.sync.management.UnlockRequest
	6,  // 6: freym.sync.management.Service.CreateLease:output_type -> freym.sync.management.CreateLeaseResponse
	7,  // 7: freym.sync.management.Service.KeepLease:output_type -> freym.sync.management.KeepLeaseResponse
	8,  // 8: freym.sync.management.Service.DropLease:output_type -> freym.sync.management.DropLeaseResponse
	9,  // 9: freym.sync.management.Service.GetPeerNodes:output_type -> freym.sync.management.GetPeerNodesResponse
	10, // 10: freym.sync.management.Service.Lock:output_type -> freym.sync.management.LockResponse
	11, // 11: freym.sync.management.Service.Unlock:output_type -> freym.sync.management.UnlockResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_sync_management_service_proto_init() }
func file_sync_management_service_proto_init() {
	if File_sync_management_service_proto != nil {
		return
	}
	file_sync_management_lease_proto_init()
	file_sync_management_peer_nodes_proto_init()
	file_sync_management_lock_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sync_management_service_proto_rawDesc), len(file_sync_management_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sync_management_service_proto_goTypes,
		DependencyIndexes: file_sync_management_service_proto_depIdxs,
	}.Build()
	File_sync_management_service_proto = out.File
	file_sync_management_service_proto_goTypes = nil
	file_sync_management_service_proto_depIdxs = nil
}
