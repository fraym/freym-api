// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: crud/delivery/shared.proto

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

type DeploymentTarget int32

const (
	DeploymentTarget_DEPLOYMENT_TARGET_BLUE  DeploymentTarget = 0
	DeploymentTarget_DEPLOYMENT_TARGET_GREEN DeploymentTarget = 1
)

// Enum value maps for DeploymentTarget.
var (
	DeploymentTarget_name = map[int32]string{
		0: "DEPLOYMENT_TARGET_BLUE",
		1: "DEPLOYMENT_TARGET_GREEN",
	}
	DeploymentTarget_value = map[string]int32{
		"DEPLOYMENT_TARGET_BLUE":  0,
		"DEPLOYMENT_TARGET_GREEN": 1,
	}
)

func (x DeploymentTarget) Enum() *DeploymentTarget {
	p := new(DeploymentTarget)
	*p = x
	return p
}

func (x DeploymentTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_crud_delivery_shared_proto_enumTypes[0].Descriptor()
}

func (DeploymentTarget) Type() protoreflect.EnumType {
	return &file_crud_delivery_shared_proto_enumTypes[0]
}

func (x DeploymentTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type AuthData struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TenantId string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3"`
	xxx_hidden_Scopes   []string               `protobuf:"bytes,2,rep,name=scopes,proto3"`
	xxx_hidden_Data     map[string]string      `protobuf:"bytes,3,rep,name=data,proto3" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_UserId   string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	mi := &file_crud_delivery_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_crud_delivery_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *AuthData) GetTenantId() string {
	if x != nil {
		return x.xxx_hidden_TenantId
	}
	return ""
}

func (x *AuthData) GetScopes() []string {
	if x != nil {
		return x.xxx_hidden_Scopes
	}
	return nil
}

func (x *AuthData) GetData() map[string]string {
	if x != nil {
		return x.xxx_hidden_Data
	}
	return nil
}

func (x *AuthData) GetUserId() string {
	if x != nil {
		return x.xxx_hidden_UserId
	}
	return ""
}

func (x *AuthData) SetTenantId(v string) {
	x.xxx_hidden_TenantId = v
}

func (x *AuthData) SetScopes(v []string) {
	x.xxx_hidden_Scopes = v
}

func (x *AuthData) SetData(v map[string]string) {
	x.xxx_hidden_Data = v
}

func (x *AuthData) SetUserId(v string) {
	x.xxx_hidden_UserId = v
}

type AuthData_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	TenantId string
	Scopes   []string
	Data     map[string]string
	UserId   string
}

func (b0 AuthData_builder) Build() *AuthData {
	m0 := &AuthData{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_TenantId = b.TenantId
	x.xxx_hidden_Scopes = b.Scopes
	x.xxx_hidden_Data = b.Data
	x.xxx_hidden_UserId = b.UserId
	return m0
}

type DataFilter struct {
	state             protoimpl.MessageState      `protogen:"opaque.v1"`
	xxx_hidden_Fields map[string]*DataFieldFilter `protobuf:"bytes,1,rep,name=fields,proto3" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_And    *[]*DataFilter              `protobuf:"bytes,2,rep,name=and,proto3"`
	xxx_hidden_Or     *[]*DataFilter              `protobuf:"bytes,3,rep,name=or,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DataFilter) Reset() {
	*x = DataFilter{}
	mi := &file_crud_delivery_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFilter) ProtoMessage() {}

func (x *DataFilter) ProtoReflect() protoreflect.Message {
	mi := &file_crud_delivery_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DataFilter) GetFields() map[string]*DataFieldFilter {
	if x != nil {
		return x.xxx_hidden_Fields
	}
	return nil
}

func (x *DataFilter) GetAnd() []*DataFilter {
	if x != nil {
		if x.xxx_hidden_And != nil {
			return *x.xxx_hidden_And
		}
	}
	return nil
}

func (x *DataFilter) GetOr() []*DataFilter {
	if x != nil {
		if x.xxx_hidden_Or != nil {
			return *x.xxx_hidden_Or
		}
	}
	return nil
}

func (x *DataFilter) SetFields(v map[string]*DataFieldFilter) {
	x.xxx_hidden_Fields = v
}

func (x *DataFilter) SetAnd(v []*DataFilter) {
	x.xxx_hidden_And = &v
}

func (x *DataFilter) SetOr(v []*DataFilter) {
	x.xxx_hidden_Or = &v
}

type DataFilter_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Fields map[string]*DataFieldFilter
	And    []*DataFilter
	Or     []*DataFilter
}

func (b0 DataFilter_builder) Build() *DataFilter {
	m0 := &DataFilter{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Fields = b.Fields
	x.xxx_hidden_And = &b.And
	x.xxx_hidden_Or = &b.Or
	return m0
}

type DataFieldFilter struct {
	state                protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Type      string                 `protobuf:"bytes,1,opt,name=type,proto3"`
	xxx_hidden_Operation string                 `protobuf:"bytes,2,opt,name=operation,proto3"`
	xxx_hidden_Value     string                 `protobuf:"bytes,3,opt,name=value,proto3"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *DataFieldFilter) Reset() {
	*x = DataFieldFilter{}
	mi := &file_crud_delivery_shared_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataFieldFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFieldFilter) ProtoMessage() {}

func (x *DataFieldFilter) ProtoReflect() protoreflect.Message {
	mi := &file_crud_delivery_shared_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DataFieldFilter) GetType() string {
	if x != nil {
		return x.xxx_hidden_Type
	}
	return ""
}

func (x *DataFieldFilter) GetOperation() string {
	if x != nil {
		return x.xxx_hidden_Operation
	}
	return ""
}

func (x *DataFieldFilter) GetValue() string {
	if x != nil {
		return x.xxx_hidden_Value
	}
	return ""
}

func (x *DataFieldFilter) SetType(v string) {
	x.xxx_hidden_Type = v
}

func (x *DataFieldFilter) SetOperation(v string) {
	x.xxx_hidden_Operation = v
}

func (x *DataFieldFilter) SetValue(v string) {
	x.xxx_hidden_Value = v
}

type DataFieldFilter_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Type      string
	Operation string
	Value     string
}

func (b0 DataFieldFilter_builder) Build() *DataFieldFilter {
	m0 := &DataFieldFilter{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Type = b.Type
	x.xxx_hidden_Operation = b.Operation
	x.xxx_hidden_Value = b.Value
	return m0
}

type EventMetadata struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_CausationId   string                 `protobuf:"bytes,1,opt,name=causation_id,json=causationId,proto3"`
	xxx_hidden_CorrelationId string                 `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3"`
	xxx_hidden_UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3"`
	xxx_hidden_Target        DeploymentTarget       `protobuf:"varint,4,opt,name=target,proto3,enum=freym.crud.delivery.DeploymentTarget"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *EventMetadata) Reset() {
	*x = EventMetadata{}
	mi := &file_crud_delivery_shared_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetadata) ProtoMessage() {}

func (x *EventMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_crud_delivery_shared_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *EventMetadata) GetCausationId() string {
	if x != nil {
		return x.xxx_hidden_CausationId
	}
	return ""
}

func (x *EventMetadata) GetCorrelationId() string {
	if x != nil {
		return x.xxx_hidden_CorrelationId
	}
	return ""
}

func (x *EventMetadata) GetUserId() string {
	if x != nil {
		return x.xxx_hidden_UserId
	}
	return ""
}

func (x *EventMetadata) GetTarget() DeploymentTarget {
	if x != nil {
		return x.xxx_hidden_Target
	}
	return DeploymentTarget_DEPLOYMENT_TARGET_BLUE
}

func (x *EventMetadata) SetCausationId(v string) {
	x.xxx_hidden_CausationId = v
}

func (x *EventMetadata) SetCorrelationId(v string) {
	x.xxx_hidden_CorrelationId = v
}

func (x *EventMetadata) SetUserId(v string) {
	x.xxx_hidden_UserId = v
}

func (x *EventMetadata) SetTarget(v DeploymentTarget) {
	x.xxx_hidden_Target = v
}

type EventMetadata_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	CausationId   string
	CorrelationId string
	UserId        string
	Target        DeploymentTarget
}

func (b0 EventMetadata_builder) Build() *EventMetadata {
	m0 := &EventMetadata{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_CausationId = b.CausationId
	x.xxx_hidden_CorrelationId = b.CorrelationId
	x.xxx_hidden_UserId = b.UserId
	x.xxx_hidden_Target = b.Target
	return m0
}

var File_crud_delivery_shared_proto protoreflect.FileDescriptor

var file_crud_delivery_shared_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x66, 0x72,
	0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x96, 0x02, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x43, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x03, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x03, 0x61, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x02, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x02, 0x6f, 0x72, 0x1a, 0x5f, 0x0a, 0x0b, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x72, 0x65,
	0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a, 0x0f, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x75, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x61, 0x75, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x66, 0x72,
	0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2a, 0x4b, 0x0a, 0x10, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1a,
	0x0a, 0x16, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x5f, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f,
	0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_crud_delivery_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crud_delivery_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_crud_delivery_shared_proto_goTypes = []any{
	(DeploymentTarget)(0),   // 0: freym.crud.delivery.DeploymentTarget
	(*AuthData)(nil),        // 1: freym.crud.delivery.AuthData
	(*DataFilter)(nil),      // 2: freym.crud.delivery.DataFilter
	(*DataFieldFilter)(nil), // 3: freym.crud.delivery.DataFieldFilter
	(*EventMetadata)(nil),   // 4: freym.crud.delivery.EventMetadata
	nil,                     // 5: freym.crud.delivery.AuthData.DataEntry
	nil,                     // 6: freym.crud.delivery.DataFilter.FieldsEntry
}
var file_crud_delivery_shared_proto_depIdxs = []int32{
	5, // 0: freym.crud.delivery.AuthData.data:type_name -> freym.crud.delivery.AuthData.DataEntry
	6, // 1: freym.crud.delivery.DataFilter.fields:type_name -> freym.crud.delivery.DataFilter.FieldsEntry
	2, // 2: freym.crud.delivery.DataFilter.and:type_name -> freym.crud.delivery.DataFilter
	2, // 3: freym.crud.delivery.DataFilter.or:type_name -> freym.crud.delivery.DataFilter
	0, // 4: freym.crud.delivery.EventMetadata.target:type_name -> freym.crud.delivery.DeploymentTarget
	3, // 5: freym.crud.delivery.DataFilter.FieldsEntry.value:type_name -> freym.crud.delivery.DataFieldFilter
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_crud_delivery_shared_proto_init() }
func file_crud_delivery_shared_proto_init() {
	if File_crud_delivery_shared_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_crud_delivery_shared_proto_rawDesc), len(file_crud_delivery_shared_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crud_delivery_shared_proto_goTypes,
		DependencyIndexes: file_crud_delivery_shared_proto_depIdxs,
		EnumInfos:         file_crud_delivery_shared_proto_enumTypes,
		MessageInfos:      file_crud_delivery_shared_proto_msgTypes,
	}.Build()
	File_crud_delivery_shared_proto = out.File
	file_crud_delivery_shared_proto_goTypes = nil
	file_crud_delivery_shared_proto_depIdxs = nil
}
