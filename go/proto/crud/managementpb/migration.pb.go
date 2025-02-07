// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: crud/management/migration.proto

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
	return file_crud_management_migration_proto_enumTypes[0].Descriptor()
}

func (DeploymentTarget) Type() protoreflect.EnumType {
	return &file_crud_management_migration_proto_enumTypes[0]
}

func (x DeploymentTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type DeploySchemaRequest struct {
	state                      protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Namespace       string                 `protobuf:"bytes,1,opt,name=namespace,proto3"`
	xxx_hidden_ProjectionTypes *[]*ObjectType         `protobuf:"bytes,2,rep,name=projection_types,json=projectionTypes,proto3"`
	xxx_hidden_CrudTypes       *[]*ObjectType         `protobuf:"bytes,3,rep,name=crud_types,json=crudTypes,proto3"`
	xxx_hidden_NestedTypes     *[]*ObjectType         `protobuf:"bytes,4,rep,name=nested_types,json=nestedTypes,proto3"`
	xxx_hidden_EnumTypes       *[]*EnumType           `protobuf:"bytes,5,rep,name=enum_types,json=enumTypes,proto3"`
	xxx_hidden_Options         *DeploymentOptions     `protobuf:"bytes,6,opt,name=options,proto3"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *DeploySchemaRequest) Reset() {
	*x = DeploySchemaRequest{}
	mi := &file_crud_management_migration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploySchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploySchemaRequest) ProtoMessage() {}

func (x *DeploySchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DeploySchemaRequest) GetNamespace() string {
	if x != nil {
		return x.xxx_hidden_Namespace
	}
	return ""
}

func (x *DeploySchemaRequest) GetProjectionTypes() []*ObjectType {
	if x != nil {
		if x.xxx_hidden_ProjectionTypes != nil {
			return *x.xxx_hidden_ProjectionTypes
		}
	}
	return nil
}

func (x *DeploySchemaRequest) GetCrudTypes() []*ObjectType {
	if x != nil {
		if x.xxx_hidden_CrudTypes != nil {
			return *x.xxx_hidden_CrudTypes
		}
	}
	return nil
}

func (x *DeploySchemaRequest) GetNestedTypes() []*ObjectType {
	if x != nil {
		if x.xxx_hidden_NestedTypes != nil {
			return *x.xxx_hidden_NestedTypes
		}
	}
	return nil
}

func (x *DeploySchemaRequest) GetEnumTypes() []*EnumType {
	if x != nil {
		if x.xxx_hidden_EnumTypes != nil {
			return *x.xxx_hidden_EnumTypes
		}
	}
	return nil
}

func (x *DeploySchemaRequest) GetOptions() *DeploymentOptions {
	if x != nil {
		return x.xxx_hidden_Options
	}
	return nil
}

func (x *DeploySchemaRequest) SetNamespace(v string) {
	x.xxx_hidden_Namespace = v
}

func (x *DeploySchemaRequest) SetProjectionTypes(v []*ObjectType) {
	x.xxx_hidden_ProjectionTypes = &v
}

func (x *DeploySchemaRequest) SetCrudTypes(v []*ObjectType) {
	x.xxx_hidden_CrudTypes = &v
}

func (x *DeploySchemaRequest) SetNestedTypes(v []*ObjectType) {
	x.xxx_hidden_NestedTypes = &v
}

func (x *DeploySchemaRequest) SetEnumTypes(v []*EnumType) {
	x.xxx_hidden_EnumTypes = &v
}

func (x *DeploySchemaRequest) SetOptions(v *DeploymentOptions) {
	x.xxx_hidden_Options = v
}

func (x *DeploySchemaRequest) HasOptions() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Options != nil
}

func (x *DeploySchemaRequest) ClearOptions() {
	x.xxx_hidden_Options = nil
}

type DeploySchemaRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Namespace       string
	ProjectionTypes []*ObjectType
	CrudTypes       []*ObjectType
	NestedTypes     []*ObjectType
	EnumTypes       []*EnumType
	Options         *DeploymentOptions
}

func (b0 DeploySchemaRequest_builder) Build() *DeploySchemaRequest {
	m0 := &DeploySchemaRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Namespace = b.Namespace
	x.xxx_hidden_ProjectionTypes = &b.ProjectionTypes
	x.xxx_hidden_CrudTypes = &b.CrudTypes
	x.xxx_hidden_NestedTypes = &b.NestedTypes
	x.xxx_hidden_EnumTypes = &b.EnumTypes
	x.xxx_hidden_Options = b.Options
	return m0
}

type DeploySchemaResponse struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploySchemaResponse) Reset() {
	*x = DeploySchemaResponse{}
	mi := &file_crud_management_migration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploySchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploySchemaResponse) ProtoMessage() {}

func (x *DeploySchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type DeploySchemaResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 DeploySchemaResponse_builder) Build() *DeploySchemaResponse {
	m0 := &DeploySchemaResponse{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type ConfirmSchemaRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DeploymentId int64                  `protobuf:"varint,1,opt,name=deployment_id,json=deploymentId,proto3"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *ConfirmSchemaRequest) Reset() {
	*x = ConfirmSchemaRequest{}
	mi := &file_crud_management_migration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmSchemaRequest) ProtoMessage() {}

func (x *ConfirmSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ConfirmSchemaRequest) GetDeploymentId() int64 {
	if x != nil {
		return x.xxx_hidden_DeploymentId
	}
	return 0
}

func (x *ConfirmSchemaRequest) SetDeploymentId(v int64) {
	x.xxx_hidden_DeploymentId = v
}

type ConfirmSchemaRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DeploymentId int64
}

func (b0 ConfirmSchemaRequest_builder) Build() *ConfirmSchemaRequest {
	m0 := &ConfirmSchemaRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_DeploymentId = b.DeploymentId
	return m0
}

type ConfirmSchemaResponse struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmSchemaResponse) Reset() {
	*x = ConfirmSchemaResponse{}
	mi := &file_crud_management_migration_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmSchemaResponse) ProtoMessage() {}

func (x *ConfirmSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type ConfirmSchemaResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 ConfirmSchemaResponse_builder) Build() *ConfirmSchemaResponse {
	m0 := &ConfirmSchemaResponse{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type RollbackSchemaRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DeploymentId int64                  `protobuf:"varint,1,opt,name=deployment_id,json=deploymentId,proto3"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *RollbackSchemaRequest) Reset() {
	*x = RollbackSchemaRequest{}
	mi := &file_crud_management_migration_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackSchemaRequest) ProtoMessage() {}

func (x *RollbackSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RollbackSchemaRequest) GetDeploymentId() int64 {
	if x != nil {
		return x.xxx_hidden_DeploymentId
	}
	return 0
}

func (x *RollbackSchemaRequest) SetDeploymentId(v int64) {
	x.xxx_hidden_DeploymentId = v
}

type RollbackSchemaRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DeploymentId int64
}

func (b0 RollbackSchemaRequest_builder) Build() *RollbackSchemaRequest {
	m0 := &RollbackSchemaRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_DeploymentId = b.DeploymentId
	return m0
}

type RollbackSchemaResponse struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RollbackSchemaResponse) Reset() {
	*x = RollbackSchemaResponse{}
	mi := &file_crud_management_migration_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RollbackSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackSchemaResponse) ProtoMessage() {}

func (x *RollbackSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type RollbackSchemaResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 RollbackSchemaResponse_builder) Build() *RollbackSchemaResponse {
	m0 := &RollbackSchemaResponse{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type GetSchemaDeploymentRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DeploymentId int64                  `protobuf:"varint,1,opt,name=deployment_id,json=deploymentId,proto3"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *GetSchemaDeploymentRequest) Reset() {
	*x = GetSchemaDeploymentRequest{}
	mi := &file_crud_management_migration_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchemaDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaDeploymentRequest) ProtoMessage() {}

func (x *GetSchemaDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetSchemaDeploymentRequest) GetDeploymentId() int64 {
	if x != nil {
		return x.xxx_hidden_DeploymentId
	}
	return 0
}

func (x *GetSchemaDeploymentRequest) SetDeploymentId(v int64) {
	x.xxx_hidden_DeploymentId = v
}

type GetSchemaDeploymentRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DeploymentId int64
}

func (b0 GetSchemaDeploymentRequest_builder) Build() *GetSchemaDeploymentRequest {
	m0 := &GetSchemaDeploymentRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_DeploymentId = b.DeploymentId
	return m0
}

type GetSchemaDeploymentResponse struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Progress uint32                 `protobuf:"varint,1,opt,name=progress,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *GetSchemaDeploymentResponse) Reset() {
	*x = GetSchemaDeploymentResponse{}
	mi := &file_crud_management_migration_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchemaDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaDeploymentResponse) ProtoMessage() {}

func (x *GetSchemaDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetSchemaDeploymentResponse) GetProgress() uint32 {
	if x != nil {
		return x.xxx_hidden_Progress
	}
	return 0
}

func (x *GetSchemaDeploymentResponse) SetProgress(v uint32) {
	x.xxx_hidden_Progress = v
}

type GetSchemaDeploymentResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Progress uint32
}

func (b0 GetSchemaDeploymentResponse_builder) Build() *GetSchemaDeploymentResponse {
	m0 := &GetSchemaDeploymentResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Progress = b.Progress
	return m0
}

type DeploymentOptions struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Target DeploymentTarget       `protobuf:"varint,1,opt,name=target,proto3,enum=freym.crud.management.DeploymentTarget"`
	xxx_hidden_Force  bool                   `protobuf:"varint,2,opt,name=force,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DeploymentOptions) Reset() {
	*x = DeploymentOptions{}
	mi := &file_crud_management_migration_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentOptions) ProtoMessage() {}

func (x *DeploymentOptions) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DeploymentOptions) GetTarget() DeploymentTarget {
	if x != nil {
		return x.xxx_hidden_Target
	}
	return DeploymentTarget_DEPLOYMENT_TARGET_BLUE
}

func (x *DeploymentOptions) GetForce() bool {
	if x != nil {
		return x.xxx_hidden_Force
	}
	return false
}

func (x *DeploymentOptions) SetTarget(v DeploymentTarget) {
	x.xxx_hidden_Target = v
}

func (x *DeploymentOptions) SetForce(v bool) {
	x.xxx_hidden_Force = v
}

type DeploymentOptions_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Target DeploymentTarget
	Force  bool
}

func (b0 DeploymentOptions_builder) Build() *DeploymentOptions {
	m0 := &DeploymentOptions{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Target = b.Target
	x.xxx_hidden_Force = b.Force
	return m0
}

type ObjectType struct {
	state                 protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name       string                 `protobuf:"bytes,1,opt,name=name,proto3"`
	xxx_hidden_Directives *[]*TypeDirective      `protobuf:"bytes,2,rep,name=directives,proto3"`
	xxx_hidden_Fields     *[]*TypeField          `protobuf:"bytes,3,rep,name=fields,proto3"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ObjectType) Reset() {
	*x = ObjectType{}
	mi := &file_crud_management_migration_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectType) ProtoMessage() {}

func (x *ObjectType) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ObjectType) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *ObjectType) GetDirectives() []*TypeDirective {
	if x != nil {
		if x.xxx_hidden_Directives != nil {
			return *x.xxx_hidden_Directives
		}
	}
	return nil
}

func (x *ObjectType) GetFields() []*TypeField {
	if x != nil {
		if x.xxx_hidden_Fields != nil {
			return *x.xxx_hidden_Fields
		}
	}
	return nil
}

func (x *ObjectType) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *ObjectType) SetDirectives(v []*TypeDirective) {
	x.xxx_hidden_Directives = &v
}

func (x *ObjectType) SetFields(v []*TypeField) {
	x.xxx_hidden_Fields = &v
}

type ObjectType_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Name       string
	Directives []*TypeDirective
	Fields     []*TypeField
}

func (b0 ObjectType_builder) Build() *ObjectType {
	m0 := &ObjectType{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Directives = &b.Directives
	x.xxx_hidden_Fields = &b.Fields
	return m0
}

type TypeDirective struct {
	state                protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name      string                 `protobuf:"bytes,1,opt,name=name,proto3"`
	xxx_hidden_Arguments *[]*TypeArgument       `protobuf:"bytes,2,rep,name=arguments,proto3"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *TypeDirective) Reset() {
	*x = TypeDirective{}
	mi := &file_crud_management_migration_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TypeDirective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeDirective) ProtoMessage() {}

func (x *TypeDirective) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TypeDirective) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *TypeDirective) GetArguments() []*TypeArgument {
	if x != nil {
		if x.xxx_hidden_Arguments != nil {
			return *x.xxx_hidden_Arguments
		}
	}
	return nil
}

func (x *TypeDirective) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *TypeDirective) SetArguments(v []*TypeArgument) {
	x.xxx_hidden_Arguments = &v
}

type TypeDirective_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Name      string
	Arguments []*TypeArgument
}

func (b0 TypeDirective_builder) Build() *TypeDirective {
	m0 := &TypeDirective{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Arguments = &b.Arguments
	return m0
}

type TypeField struct {
	state                 protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name       string                 `protobuf:"bytes,1,opt,name=name,proto3"`
	xxx_hidden_Type       []string               `protobuf:"bytes,2,rep,name=type,proto3"`
	xxx_hidden_Directives *[]*TypeDirective      `protobuf:"bytes,3,rep,name=directives,proto3"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *TypeField) Reset() {
	*x = TypeField{}
	mi := &file_crud_management_migration_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TypeField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeField) ProtoMessage() {}

func (x *TypeField) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TypeField) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *TypeField) GetType() []string {
	if x != nil {
		return x.xxx_hidden_Type
	}
	return nil
}

func (x *TypeField) GetDirectives() []*TypeDirective {
	if x != nil {
		if x.xxx_hidden_Directives != nil {
			return *x.xxx_hidden_Directives
		}
	}
	return nil
}

func (x *TypeField) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *TypeField) SetType(v []string) {
	x.xxx_hidden_Type = v
}

func (x *TypeField) SetDirectives(v []*TypeDirective) {
	x.xxx_hidden_Directives = &v
}

type TypeField_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Name       string
	Type       []string
	Directives []*TypeDirective
}

func (b0 TypeField_builder) Build() *TypeField {
	m0 := &TypeField{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Type = b.Type
	x.xxx_hidden_Directives = &b.Directives
	return m0
}

type TypeArgument struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name  string                 `protobuf:"bytes,1,opt,name=name,proto3"`
	xxx_hidden_Value string                 `protobuf:"bytes,2,opt,name=value,proto3"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TypeArgument) Reset() {
	*x = TypeArgument{}
	mi := &file_crud_management_migration_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TypeArgument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeArgument) ProtoMessage() {}

func (x *TypeArgument) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TypeArgument) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *TypeArgument) GetValue() string {
	if x != nil {
		return x.xxx_hidden_Value
	}
	return ""
}

func (x *TypeArgument) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *TypeArgument) SetValue(v string) {
	x.xxx_hidden_Value = v
}

type TypeArgument_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Name  string
	Value string
}

func (b0 TypeArgument_builder) Build() *TypeArgument {
	m0 := &TypeArgument{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Value = b.Value
	return m0
}

type EnumType struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name   string                 `protobuf:"bytes,1,opt,name=name,proto3"`
	xxx_hidden_Values []string               `protobuf:"bytes,2,rep,name=values,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *EnumType) Reset() {
	*x = EnumType{}
	mi := &file_crud_management_migration_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnumType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumType) ProtoMessage() {}

func (x *EnumType) ProtoReflect() protoreflect.Message {
	mi := &file_crud_management_migration_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *EnumType) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *EnumType) GetValues() []string {
	if x != nil {
		return x.xxx_hidden_Values
	}
	return nil
}

func (x *EnumType) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *EnumType) SetValues(v []string) {
	x.xxx_hidden_Values = v
}

type EnumType_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Name   string
	Values []string
}

func (b0 EnumType_builder) Build() *EnumType {
	m0 := &EnumType{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Values = b.Values
	return m0
}

var File_crud_management_migration_proto protoreflect.FileDescriptor

var file_crud_management_migration_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x8d, 0x03, 0x0a, 0x13, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d,
	0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x72, 0x75, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x44,
	0x0a, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d,
	0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3b, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x15, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x39, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6a, 0x0a, 0x11,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3f, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x66, 0x0a, 0x0d, 0x54,
	0x79, 0x70, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x41, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x79, 0x0a, 0x09, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66,
	0x72, 0x65, 0x79, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x22, 0x38,
	0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x2a, 0x4b, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_crud_management_migration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crud_management_migration_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_crud_management_migration_proto_goTypes = []any{
	(DeploymentTarget)(0),               // 0: freym.crud.management.DeploymentTarget
	(*DeploySchemaRequest)(nil),         // 1: freym.crud.management.DeploySchemaRequest
	(*DeploySchemaResponse)(nil),        // 2: freym.crud.management.DeploySchemaResponse
	(*ConfirmSchemaRequest)(nil),        // 3: freym.crud.management.ConfirmSchemaRequest
	(*ConfirmSchemaResponse)(nil),       // 4: freym.crud.management.ConfirmSchemaResponse
	(*RollbackSchemaRequest)(nil),       // 5: freym.crud.management.RollbackSchemaRequest
	(*RollbackSchemaResponse)(nil),      // 6: freym.crud.management.RollbackSchemaResponse
	(*GetSchemaDeploymentRequest)(nil),  // 7: freym.crud.management.GetSchemaDeploymentRequest
	(*GetSchemaDeploymentResponse)(nil), // 8: freym.crud.management.GetSchemaDeploymentResponse
	(*DeploymentOptions)(nil),           // 9: freym.crud.management.DeploymentOptions
	(*ObjectType)(nil),                  // 10: freym.crud.management.ObjectType
	(*TypeDirective)(nil),               // 11: freym.crud.management.TypeDirective
	(*TypeField)(nil),                   // 12: freym.crud.management.TypeField
	(*TypeArgument)(nil),                // 13: freym.crud.management.TypeArgument
	(*EnumType)(nil),                    // 14: freym.crud.management.EnumType
}
var file_crud_management_migration_proto_depIdxs = []int32{
	10, // 0: freym.crud.management.DeploySchemaRequest.projection_types:type_name -> freym.crud.management.ObjectType
	10, // 1: freym.crud.management.DeploySchemaRequest.crud_types:type_name -> freym.crud.management.ObjectType
	10, // 2: freym.crud.management.DeploySchemaRequest.nested_types:type_name -> freym.crud.management.ObjectType
	14, // 3: freym.crud.management.DeploySchemaRequest.enum_types:type_name -> freym.crud.management.EnumType
	9,  // 4: freym.crud.management.DeploySchemaRequest.options:type_name -> freym.crud.management.DeploymentOptions
	0,  // 5: freym.crud.management.DeploymentOptions.target:type_name -> freym.crud.management.DeploymentTarget
	11, // 6: freym.crud.management.ObjectType.directives:type_name -> freym.crud.management.TypeDirective
	12, // 7: freym.crud.management.ObjectType.fields:type_name -> freym.crud.management.TypeField
	13, // 8: freym.crud.management.TypeDirective.arguments:type_name -> freym.crud.management.TypeArgument
	11, // 9: freym.crud.management.TypeField.directives:type_name -> freym.crud.management.TypeDirective
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_crud_management_migration_proto_init() }
func file_crud_management_migration_proto_init() {
	if File_crud_management_migration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_crud_management_migration_proto_rawDesc), len(file_crud_management_migration_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crud_management_migration_proto_goTypes,
		DependencyIndexes: file_crud_management_migration_proto_depIdxs,
		EnumInfos:         file_crud_management_migration_proto_enumTypes,
		MessageInfos:      file_crud_management_migration_proto_msgTypes,
	}.Build()
	File_crud_management_migration_proto = out.File
	file_crud_management_migration_proto_goTypes = nil
	file_crud_management_migration_proto_depIdxs = nil
}
