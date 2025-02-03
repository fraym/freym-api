// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: auth/management/roles_get.proto

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

type GetRolesRequest struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TenantId string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *GetRolesRequest) Reset() {
	*x = GetRolesRequest{}
	mi := &file_auth_management_roles_get_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesRequest) ProtoMessage() {}

func (x *GetRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_management_roles_get_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetRolesRequest) GetTenantId() string {
	if x != nil {
		return x.xxx_hidden_TenantId
	}
	return ""
}

func (x *GetRolesRequest) SetTenantId(v string) {
	x.xxx_hidden_TenantId = v
}

type GetRolesRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	TenantId string
}

func (b0 GetRolesRequest_builder) Build() *GetRolesRequest {
	m0 := &GetRolesRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_TenantId = b.TenantId
	return m0
}

type GetRolesResponse struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Roles *[]*Role               `protobuf:"bytes,1,rep,name=roles,proto3"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetRolesResponse) Reset() {
	*x = GetRolesResponse{}
	mi := &file_auth_management_roles_get_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesResponse) ProtoMessage() {}

func (x *GetRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_management_roles_get_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetRolesResponse) GetRoles() []*Role {
	if x != nil {
		if x.xxx_hidden_Roles != nil {
			return *x.xxx_hidden_Roles
		}
	}
	return nil
}

func (x *GetRolesResponse) SetRoles(v []*Role) {
	x.xxx_hidden_Roles = &v
}

type GetRolesResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Roles []*Role
}

func (b0 GetRolesResponse_builder) Build() *GetRolesResponse {
	m0 := &GetRolesResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Roles = &b.Roles
	return m0
}

type Role struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id            string                 `protobuf:"bytes,1,opt,name=id,proto3"`
	xxx_hidden_AllowedScopes *[]*RoleScope          `protobuf:"bytes,2,rep,name=allowed_scopes,json=allowedScopes,proto3"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_auth_management_roles_get_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_auth_management_roles_get_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Role) GetId() string {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return ""
}

func (x *Role) GetAllowedScopes() []*RoleScope {
	if x != nil {
		if x.xxx_hidden_AllowedScopes != nil {
			return *x.xxx_hidden_AllowedScopes
		}
	}
	return nil
}

func (x *Role) SetId(v string) {
	x.xxx_hidden_Id = v
}

func (x *Role) SetAllowedScopes(v []*RoleScope) {
	x.xxx_hidden_AllowedScopes = &v
}

type Role_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id            string
	AllowedScopes []*RoleScope
}

func (b0 Role_builder) Build() *Role {
	m0 := &Role{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	x.xxx_hidden_AllowedScopes = &b.AllowedScopes
	return m0
}

var File_auth_management_roles_get_proto protoreflect.FileDescriptor

var file_auth_management_roles_get_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x72, 0x65, 0x79,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x5f, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_auth_management_roles_get_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auth_management_roles_get_proto_goTypes = []any{
	(*GetRolesRequest)(nil),  // 0: freym.auth.management.GetRolesRequest
	(*GetRolesResponse)(nil), // 1: freym.auth.management.GetRolesResponse
	(*Role)(nil),             // 2: freym.auth.management.Role
	(*RoleScope)(nil),        // 3: freym.auth.management.RoleScope
}
var file_auth_management_roles_get_proto_depIdxs = []int32{
	2, // 0: freym.auth.management.GetRolesResponse.roles:type_name -> freym.auth.management.Role
	3, // 1: freym.auth.management.Role.allowed_scopes:type_name -> freym.auth.management.RoleScope
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_management_roles_get_proto_init() }
func file_auth_management_roles_get_proto_init() {
	if File_auth_management_roles_get_proto != nil {
		return
	}
	file_auth_management_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_management_roles_get_proto_rawDesc), len(file_auth_management_roles_get_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_management_roles_get_proto_goTypes,
		DependencyIndexes: file_auth_management_roles_get_proto_depIdxs,
		MessageInfos:      file_auth_management_roles_get_proto_msgTypes,
	}.Build()
	File_auth_management_roles_get_proto = out.File
	file_auth_management_roles_get_proto_goTypes = nil
	file_auth_management_roles_get_proto_depIdxs = nil
}
