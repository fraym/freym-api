// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: projections/delivery/service.proto

package deliverypb

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

var File_projections_delivery_service_proto protoreflect.FileDescriptor

const file_projections_delivery_service_proto_rawDesc = "" +
	"\n" +
	"\"projections/delivery/service.proto\x12\x1afreym.projections.delivery\x1a#projections/delivery/get_data.proto\x1a(projections/delivery/get_view_data.proto\x1a!projections/delivery/upsert.proto\x1a!projections/delivery/delete.proto2\x8b\x05\n" +
	"\aService\x12b\n" +
	"\aGetData\x12*.freym.projections.delivery.GetDataRequest\x1a+.freym.projections.delivery.GetDataResponse\x12n\n" +
	"\vGetViewData\x12..freym.projections.delivery.GetViewDataRequest\x1a/.freym.projections.delivery.GetViewDataResponse\x12n\n" +
	"\vGetDataList\x12..freym.projections.delivery.GetDataListRequest\x1a/.freym.projections.delivery.GetDataListResponse\x12z\n" +
	"\x0fGetViewDataList\x122.freym.projections.delivery.GetViewDataListRequest\x1a3.freym.projections.delivery.GetViewDataListResponse\x12_\n" +
	"\x06Upsert\x12).freym.projections.delivery.UpsertRequest\x1a*.freym.projections.delivery.UpsertResponse\x12_\n" +
	"\x06Delete\x12).freym.projections.delivery.DeleteRequest\x1a*.freym.projections.delivery.DeleteResponseb\x06proto3"

var file_projections_delivery_service_proto_goTypes = []any{
	(*GetDataRequest)(nil),          // 0: freym.projections.delivery.GetDataRequest
	(*GetViewDataRequest)(nil),      // 1: freym.projections.delivery.GetViewDataRequest
	(*GetDataListRequest)(nil),      // 2: freym.projections.delivery.GetDataListRequest
	(*GetViewDataListRequest)(nil),  // 3: freym.projections.delivery.GetViewDataListRequest
	(*UpsertRequest)(nil),           // 4: freym.projections.delivery.UpsertRequest
	(*DeleteRequest)(nil),           // 5: freym.projections.delivery.DeleteRequest
	(*GetDataResponse)(nil),         // 6: freym.projections.delivery.GetDataResponse
	(*GetViewDataResponse)(nil),     // 7: freym.projections.delivery.GetViewDataResponse
	(*GetDataListResponse)(nil),     // 8: freym.projections.delivery.GetDataListResponse
	(*GetViewDataListResponse)(nil), // 9: freym.projections.delivery.GetViewDataListResponse
	(*UpsertResponse)(nil),          // 10: freym.projections.delivery.UpsertResponse
	(*DeleteResponse)(nil),          // 11: freym.projections.delivery.DeleteResponse
}
var file_projections_delivery_service_proto_depIdxs = []int32{
	0,  // 0: freym.projections.delivery.Service.GetData:input_type -> freym.projections.delivery.GetDataRequest
	1,  // 1: freym.projections.delivery.Service.GetViewData:input_type -> freym.projections.delivery.GetViewDataRequest
	2,  // 2: freym.projections.delivery.Service.GetDataList:input_type -> freym.projections.delivery.GetDataListRequest
	3,  // 3: freym.projections.delivery.Service.GetViewDataList:input_type -> freym.projections.delivery.GetViewDataListRequest
	4,  // 4: freym.projections.delivery.Service.Upsert:input_type -> freym.projections.delivery.UpsertRequest
	5,  // 5: freym.projections.delivery.Service.Delete:input_type -> freym.projections.delivery.DeleteRequest
	6,  // 6: freym.projections.delivery.Service.GetData:output_type -> freym.projections.delivery.GetDataResponse
	7,  // 7: freym.projections.delivery.Service.GetViewData:output_type -> freym.projections.delivery.GetViewDataResponse
	8,  // 8: freym.projections.delivery.Service.GetDataList:output_type -> freym.projections.delivery.GetDataListResponse
	9,  // 9: freym.projections.delivery.Service.GetViewDataList:output_type -> freym.projections.delivery.GetViewDataListResponse
	10, // 10: freym.projections.delivery.Service.Upsert:output_type -> freym.projections.delivery.UpsertResponse
	11, // 11: freym.projections.delivery.Service.Delete:output_type -> freym.projections.delivery.DeleteResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_projections_delivery_service_proto_init() }
func file_projections_delivery_service_proto_init() {
	if File_projections_delivery_service_proto != nil {
		return
	}
	file_projections_delivery_get_data_proto_init()
	file_projections_delivery_get_view_data_proto_init()
	file_projections_delivery_upsert_proto_init()
	file_projections_delivery_delete_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_projections_delivery_service_proto_rawDesc), len(file_projections_delivery_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_projections_delivery_service_proto_goTypes,
		DependencyIndexes: file_projections_delivery_service_proto_depIdxs,
	}.Build()
	File_projections_delivery_service_proto = out.File
	file_projections_delivery_service_proto_goTypes = nil
	file_projections_delivery_service_proto_depIdxs = nil
}
