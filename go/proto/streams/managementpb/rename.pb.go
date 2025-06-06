// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: streams/management/rename.proto

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

type RenameEventTypeRequest struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Topic   string                 `protobuf:"bytes,1,opt,name=topic,proto3"`
	xxx_hidden_OldType string                 `protobuf:"bytes,2,opt,name=old_type,json=oldType,proto3"`
	xxx_hidden_NewType string                 `protobuf:"bytes,3,opt,name=new_type,json=newType,proto3"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *RenameEventTypeRequest) Reset() {
	*x = RenameEventTypeRequest{}
	mi := &file_streams_management_rename_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenameEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameEventTypeRequest) ProtoMessage() {}

func (x *RenameEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streams_management_rename_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RenameEventTypeRequest) GetTopic() string {
	if x != nil {
		return x.xxx_hidden_Topic
	}
	return ""
}

func (x *RenameEventTypeRequest) GetOldType() string {
	if x != nil {
		return x.xxx_hidden_OldType
	}
	return ""
}

func (x *RenameEventTypeRequest) GetNewType() string {
	if x != nil {
		return x.xxx_hidden_NewType
	}
	return ""
}

func (x *RenameEventTypeRequest) SetTopic(v string) {
	x.xxx_hidden_Topic = v
}

func (x *RenameEventTypeRequest) SetOldType(v string) {
	x.xxx_hidden_OldType = v
}

func (x *RenameEventTypeRequest) SetNewType(v string) {
	x.xxx_hidden_NewType = v
}

type RenameEventTypeRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Topic   string
	OldType string
	NewType string
}

func (b0 RenameEventTypeRequest_builder) Build() *RenameEventTypeRequest {
	m0 := &RenameEventTypeRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Topic = b.Topic
	x.xxx_hidden_OldType = b.OldType
	x.xxx_hidden_NewType = b.NewType
	return m0
}

type RenameEventTypeResponse struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RenameEventTypeResponse) Reset() {
	*x = RenameEventTypeResponse{}
	mi := &file_streams_management_rename_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenameEventTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameEventTypeResponse) ProtoMessage() {}

func (x *RenameEventTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streams_management_rename_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type RenameEventTypeResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 RenameEventTypeResponse_builder) Build() *RenameEventTypeResponse {
	m0 := &RenameEventTypeResponse{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_streams_management_rename_proto protoreflect.FileDescriptor

const file_streams_management_rename_proto_rawDesc = "" +
	"\n" +
	"\x1fstreams/management/rename.proto\x12\x18freym.streams.management\"d\n" +
	"\x16RenameEventTypeRequest\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12\x19\n" +
	"\bold_type\x18\x02 \x01(\tR\aoldType\x12\x19\n" +
	"\bnew_type\x18\x03 \x01(\tR\anewType\"\x19\n" +
	"\x17RenameEventTypeResponseb\x06proto3"

var file_streams_management_rename_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_streams_management_rename_proto_goTypes = []any{
	(*RenameEventTypeRequest)(nil),  // 0: freym.streams.management.RenameEventTypeRequest
	(*RenameEventTypeResponse)(nil), // 1: freym.streams.management.RenameEventTypeResponse
}
var file_streams_management_rename_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_streams_management_rename_proto_init() }
func file_streams_management_rename_proto_init() {
	if File_streams_management_rename_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_streams_management_rename_proto_rawDesc), len(file_streams_management_rename_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_streams_management_rename_proto_goTypes,
		DependencyIndexes: file_streams_management_rename_proto_depIdxs,
		MessageInfos:      file_streams_management_rename_proto_msgTypes,
	}.Build()
	File_streams_management_rename_proto = out.File
	file_streams_management_rename_proto_goTypes = nil
	file_streams_management_rename_proto_depIdxs = nil
}
