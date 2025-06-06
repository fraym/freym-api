// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: streams/management/snapshot.proto

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

type CreateStreamSnapshotRequest struct {
	state                             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TenantId               string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3"`
	xxx_hidden_Topic                  string                 `protobuf:"bytes,2,opt,name=topic,proto3"`
	xxx_hidden_Stream                 string                 `protobuf:"bytes,3,opt,name=stream,proto3"`
	xxx_hidden_LastSnapshottedEventId string                 `protobuf:"bytes,4,opt,name=last_snapshotted_event_id,json=lastSnapshottedEventId,proto3"`
	xxx_hidden_SnapshotEvent          *PublishEvent          `protobuf:"bytes,5,opt,name=snapshot_event,json=snapshotEvent,proto3"`
	unknownFields                     protoimpl.UnknownFields
	sizeCache                         protoimpl.SizeCache
}

func (x *CreateStreamSnapshotRequest) Reset() {
	*x = CreateStreamSnapshotRequest{}
	mi := &file_streams_management_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStreamSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamSnapshotRequest) ProtoMessage() {}

func (x *CreateStreamSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streams_management_snapshot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CreateStreamSnapshotRequest) GetTenantId() string {
	if x != nil {
		return x.xxx_hidden_TenantId
	}
	return ""
}

func (x *CreateStreamSnapshotRequest) GetTopic() string {
	if x != nil {
		return x.xxx_hidden_Topic
	}
	return ""
}

func (x *CreateStreamSnapshotRequest) GetStream() string {
	if x != nil {
		return x.xxx_hidden_Stream
	}
	return ""
}

func (x *CreateStreamSnapshotRequest) GetLastSnapshottedEventId() string {
	if x != nil {
		return x.xxx_hidden_LastSnapshottedEventId
	}
	return ""
}

func (x *CreateStreamSnapshotRequest) GetSnapshotEvent() *PublishEvent {
	if x != nil {
		return x.xxx_hidden_SnapshotEvent
	}
	return nil
}

func (x *CreateStreamSnapshotRequest) SetTenantId(v string) {
	x.xxx_hidden_TenantId = v
}

func (x *CreateStreamSnapshotRequest) SetTopic(v string) {
	x.xxx_hidden_Topic = v
}

func (x *CreateStreamSnapshotRequest) SetStream(v string) {
	x.xxx_hidden_Stream = v
}

func (x *CreateStreamSnapshotRequest) SetLastSnapshottedEventId(v string) {
	x.xxx_hidden_LastSnapshottedEventId = v
}

func (x *CreateStreamSnapshotRequest) SetSnapshotEvent(v *PublishEvent) {
	x.xxx_hidden_SnapshotEvent = v
}

func (x *CreateStreamSnapshotRequest) HasSnapshotEvent() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_SnapshotEvent != nil
}

func (x *CreateStreamSnapshotRequest) ClearSnapshotEvent() {
	x.xxx_hidden_SnapshotEvent = nil
}

type CreateStreamSnapshotRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	TenantId               string
	Topic                  string
	Stream                 string
	LastSnapshottedEventId string
	SnapshotEvent          *PublishEvent
}

func (b0 CreateStreamSnapshotRequest_builder) Build() *CreateStreamSnapshotRequest {
	m0 := &CreateStreamSnapshotRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_TenantId = b.TenantId
	x.xxx_hidden_Topic = b.Topic
	x.xxx_hidden_Stream = b.Stream
	x.xxx_hidden_LastSnapshottedEventId = b.LastSnapshottedEventId
	x.xxx_hidden_SnapshotEvent = b.SnapshotEvent
	return m0
}

type CreateStreamSnapshotResponse struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStreamSnapshotResponse) Reset() {
	*x = CreateStreamSnapshotResponse{}
	mi := &file_streams_management_snapshot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStreamSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamSnapshotResponse) ProtoMessage() {}

func (x *CreateStreamSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streams_management_snapshot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type CreateStreamSnapshotResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 CreateStreamSnapshotResponse_builder) Build() *CreateStreamSnapshotResponse {
	m0 := &CreateStreamSnapshotResponse{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_streams_management_snapshot_proto protoreflect.FileDescriptor

const file_streams_management_snapshot_proto_rawDesc = "" +
	"\n" +
	"!streams/management/snapshot.proto\x12\x18freym.streams.management\x1a streams/management/publish.proto\"\xf2\x01\n" +
	"\x1bCreateStreamSnapshotRequest\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\x12\x14\n" +
	"\x05topic\x18\x02 \x01(\tR\x05topic\x12\x16\n" +
	"\x06stream\x18\x03 \x01(\tR\x06stream\x129\n" +
	"\x19last_snapshotted_event_id\x18\x04 \x01(\tR\x16lastSnapshottedEventId\x12M\n" +
	"\x0esnapshot_event\x18\x05 \x01(\v2&.freym.streams.management.PublishEventR\rsnapshotEvent\"\x1e\n" +
	"\x1cCreateStreamSnapshotResponseb\x06proto3"

var file_streams_management_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_streams_management_snapshot_proto_goTypes = []any{
	(*CreateStreamSnapshotRequest)(nil),  // 0: freym.streams.management.CreateStreamSnapshotRequest
	(*CreateStreamSnapshotResponse)(nil), // 1: freym.streams.management.CreateStreamSnapshotResponse
	(*PublishEvent)(nil),                 // 2: freym.streams.management.PublishEvent
}
var file_streams_management_snapshot_proto_depIdxs = []int32{
	2, // 0: freym.streams.management.CreateStreamSnapshotRequest.snapshot_event:type_name -> freym.streams.management.PublishEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_streams_management_snapshot_proto_init() }
func file_streams_management_snapshot_proto_init() {
	if File_streams_management_snapshot_proto != nil {
		return
	}
	file_streams_management_publish_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_streams_management_snapshot_proto_rawDesc), len(file_streams_management_snapshot_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_streams_management_snapshot_proto_goTypes,
		DependencyIndexes: file_streams_management_snapshot_proto_depIdxs,
		MessageInfos:      file_streams_management_snapshot_proto_msgTypes,
	}.Build()
	File_streams_management_snapshot_proto = out.File
	file_streams_management_snapshot_proto_goTypes = nil
	file_streams_management_snapshot_proto_depIdxs = nil
}
