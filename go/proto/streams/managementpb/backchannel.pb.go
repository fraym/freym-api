// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: streams/management/backchannel.proto

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

type BackchannelEventRequest struct {
	state              protoimpl.MessageState            `protogen:"opaque.v1"`
	xxx_hidden_Payload isBackchannelEventRequest_Payload `protobuf_oneof:"payload"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *BackchannelEventRequest) Reset() {
	*x = BackchannelEventRequest{}
	mi := &file_streams_management_backchannel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackchannelEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackchannelEventRequest) ProtoMessage() {}

func (x *BackchannelEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streams_management_backchannel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BackchannelEventRequest) GetBroadcast() *Event {
	if x != nil {
		if x, ok := x.xxx_hidden_Payload.(*backchannelEventRequest_Broadcast); ok {
			return x.Broadcast
		}
	}
	return nil
}

func (x *BackchannelEventRequest) GetNotice() *Event {
	if x != nil {
		if x, ok := x.xxx_hidden_Payload.(*backchannelEventRequest_Notice); ok {
			return x.Notice
		}
	}
	return nil
}

func (x *BackchannelEventRequest) SetBroadcast(v *Event) {
	if v == nil {
		x.xxx_hidden_Payload = nil
		return
	}
	x.xxx_hidden_Payload = &backchannelEventRequest_Broadcast{v}
}

func (x *BackchannelEventRequest) SetNotice(v *Event) {
	if v == nil {
		x.xxx_hidden_Payload = nil
		return
	}
	x.xxx_hidden_Payload = &backchannelEventRequest_Notice{v}
}

func (x *BackchannelEventRequest) HasPayload() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Payload != nil
}

func (x *BackchannelEventRequest) HasBroadcast() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Payload.(*backchannelEventRequest_Broadcast)
	return ok
}

func (x *BackchannelEventRequest) HasNotice() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Payload.(*backchannelEventRequest_Notice)
	return ok
}

func (x *BackchannelEventRequest) ClearPayload() {
	x.xxx_hidden_Payload = nil
}

func (x *BackchannelEventRequest) ClearBroadcast() {
	if _, ok := x.xxx_hidden_Payload.(*backchannelEventRequest_Broadcast); ok {
		x.xxx_hidden_Payload = nil
	}
}

func (x *BackchannelEventRequest) ClearNotice() {
	if _, ok := x.xxx_hidden_Payload.(*backchannelEventRequest_Notice); ok {
		x.xxx_hidden_Payload = nil
	}
}

const BackchannelEventRequest_Payload_not_set_case case_BackchannelEventRequest_Payload = 0
const BackchannelEventRequest_Broadcast_case case_BackchannelEventRequest_Payload = 1
const BackchannelEventRequest_Notice_case case_BackchannelEventRequest_Payload = 2

func (x *BackchannelEventRequest) WhichPayload() case_BackchannelEventRequest_Payload {
	if x == nil {
		return BackchannelEventRequest_Payload_not_set_case
	}
	switch x.xxx_hidden_Payload.(type) {
	case *backchannelEventRequest_Broadcast:
		return BackchannelEventRequest_Broadcast_case
	case *backchannelEventRequest_Notice:
		return BackchannelEventRequest_Notice_case
	default:
		return BackchannelEventRequest_Payload_not_set_case
	}
}

type BackchannelEventRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_Payload:
	Broadcast *Event
	Notice    *Event
	// -- end of xxx_hidden_Payload
}

func (b0 BackchannelEventRequest_builder) Build() *BackchannelEventRequest {
	m0 := &BackchannelEventRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Broadcast != nil {
		x.xxx_hidden_Payload = &backchannelEventRequest_Broadcast{b.Broadcast}
	}
	if b.Notice != nil {
		x.xxx_hidden_Payload = &backchannelEventRequest_Notice{b.Notice}
	}
	return m0
}

type case_BackchannelEventRequest_Payload protoreflect.FieldNumber

func (x case_BackchannelEventRequest_Payload) String() string {
	md := file_streams_management_backchannel_proto_msgTypes[0].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isBackchannelEventRequest_Payload interface {
	isBackchannelEventRequest_Payload()
}

type backchannelEventRequest_Broadcast struct {
	Broadcast *Event `protobuf:"bytes,1,opt,name=broadcast,proto3,oneof"`
}

type backchannelEventRequest_Notice struct {
	Notice *Event `protobuf:"bytes,2,opt,name=notice,proto3,oneof"`
}

func (*backchannelEventRequest_Broadcast) isBackchannelEventRequest_Payload() {}

func (*backchannelEventRequest_Notice) isBackchannelEventRequest_Payload() {}

type BackchannelEventResponse struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackchannelEventResponse) Reset() {
	*x = BackchannelEventResponse{}
	mi := &file_streams_management_backchannel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackchannelEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackchannelEventResponse) ProtoMessage() {}

func (x *BackchannelEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streams_management_backchannel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type BackchannelEventResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 BackchannelEventResponse_builder) Build() *BackchannelEventResponse {
	m0 := &BackchannelEventResponse{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_streams_management_backchannel_proto protoreflect.FileDescriptor

var file_streams_management_backchannel_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x1e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x09,
	0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x42, 0x61, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_streams_management_backchannel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_streams_management_backchannel_proto_goTypes = []any{
	(*BackchannelEventRequest)(nil),  // 0: freym.streams.management.BackchannelEventRequest
	(*BackchannelEventResponse)(nil), // 1: freym.streams.management.BackchannelEventResponse
	(*Event)(nil),                    // 2: freym.streams.management.Event
}
var file_streams_management_backchannel_proto_depIdxs = []int32{
	2, // 0: freym.streams.management.BackchannelEventRequest.broadcast:type_name -> freym.streams.management.Event
	2, // 1: freym.streams.management.BackchannelEventRequest.notice:type_name -> freym.streams.management.Event
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_streams_management_backchannel_proto_init() }
func file_streams_management_backchannel_proto_init() {
	if File_streams_management_backchannel_proto != nil {
		return
	}
	file_streams_management_event_proto_init()
	file_streams_management_backchannel_proto_msgTypes[0].OneofWrappers = []any{
		(*backchannelEventRequest_Broadcast)(nil),
		(*backchannelEventRequest_Notice)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_streams_management_backchannel_proto_rawDesc), len(file_streams_management_backchannel_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_streams_management_backchannel_proto_goTypes,
		DependencyIndexes: file_streams_management_backchannel_proto_depIdxs,
		MessageInfos:      file_streams_management_backchannel_proto_msgTypes,
	}.Build()
	File_streams_management_backchannel_proto = out.File
	file_streams_management_backchannel_proto_goTypes = nil
	file_streams_management_backchannel_proto_depIdxs = nil
}
