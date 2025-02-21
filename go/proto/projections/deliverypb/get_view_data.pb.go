// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: projections/delivery/get_view_data.proto

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

type GetViewDataRequest struct {
	state                           protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_View                 string                 `protobuf:"bytes,1,opt,name=view,proto3"`
	xxx_hidden_Auth                 *AuthData              `protobuf:"bytes,2,opt,name=auth,proto3"`
	xxx_hidden_Filter               *DataFilter            `protobuf:"bytes,3,opt,name=filter,proto3"`
	xxx_hidden_UseStrongConsistency bool                   `protobuf:"varint,4,opt,name=use_strong_consistency,json=useStrongConsistency,proto3"`
	xxx_hidden_Target               DeploymentTarget       `protobuf:"varint,5,opt,name=target,proto3,enum=freym.projections.delivery.DeploymentTarget"`
	unknownFields                   protoimpl.UnknownFields
	sizeCache                       protoimpl.SizeCache
}

func (x *GetViewDataRequest) Reset() {
	*x = GetViewDataRequest{}
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewDataRequest) ProtoMessage() {}

func (x *GetViewDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetViewDataRequest) GetView() string {
	if x != nil {
		return x.xxx_hidden_View
	}
	return ""
}

func (x *GetViewDataRequest) GetAuth() *AuthData {
	if x != nil {
		return x.xxx_hidden_Auth
	}
	return nil
}

func (x *GetViewDataRequest) GetFilter() *DataFilter {
	if x != nil {
		return x.xxx_hidden_Filter
	}
	return nil
}

func (x *GetViewDataRequest) GetUseStrongConsistency() bool {
	if x != nil {
		return x.xxx_hidden_UseStrongConsistency
	}
	return false
}

func (x *GetViewDataRequest) GetTarget() DeploymentTarget {
	if x != nil {
		return x.xxx_hidden_Target
	}
	return DeploymentTarget_DEPLOYMENT_TARGET_BLUE
}

func (x *GetViewDataRequest) SetView(v string) {
	x.xxx_hidden_View = v
}

func (x *GetViewDataRequest) SetAuth(v *AuthData) {
	x.xxx_hidden_Auth = v
}

func (x *GetViewDataRequest) SetFilter(v *DataFilter) {
	x.xxx_hidden_Filter = v
}

func (x *GetViewDataRequest) SetUseStrongConsistency(v bool) {
	x.xxx_hidden_UseStrongConsistency = v
}

func (x *GetViewDataRequest) SetTarget(v DeploymentTarget) {
	x.xxx_hidden_Target = v
}

func (x *GetViewDataRequest) HasAuth() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Auth != nil
}

func (x *GetViewDataRequest) HasFilter() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Filter != nil
}

func (x *GetViewDataRequest) ClearAuth() {
	x.xxx_hidden_Auth = nil
}

func (x *GetViewDataRequest) ClearFilter() {
	x.xxx_hidden_Filter = nil
}

type GetViewDataRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	View                 string
	Auth                 *AuthData
	Filter               *DataFilter
	UseStrongConsistency bool
	Target               DeploymentTarget
}

func (b0 GetViewDataRequest_builder) Build() *GetViewDataRequest {
	m0 := &GetViewDataRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_View = b.View
	x.xxx_hidden_Auth = b.Auth
	x.xxx_hidden_Filter = b.Filter
	x.xxx_hidden_UseStrongConsistency = b.UseStrongConsistency
	x.xxx_hidden_Target = b.Target
	return m0
}

type GetViewDataResponse struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Result *Data                  `protobuf:"bytes,1,opt,name=result,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetViewDataResponse) Reset() {
	*x = GetViewDataResponse{}
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewDataResponse) ProtoMessage() {}

func (x *GetViewDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetViewDataResponse) GetResult() *Data {
	if x != nil {
		return x.xxx_hidden_Result
	}
	return nil
}

func (x *GetViewDataResponse) SetResult(v *Data) {
	x.xxx_hidden_Result = v
}

func (x *GetViewDataResponse) HasResult() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Result != nil
}

func (x *GetViewDataResponse) ClearResult() {
	x.xxx_hidden_Result = nil
}

type GetViewDataResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Result *Data
}

func (b0 GetViewDataResponse_builder) Build() *GetViewDataResponse {
	m0 := &GetViewDataResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Result = b.Result
	return m0
}

type GetViewDataListRequest struct {
	state                           protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_View                 string                 `protobuf:"bytes,1,opt,name=view,proto3"`
	xxx_hidden_Auth                 *AuthData              `protobuf:"bytes,2,opt,name=auth,proto3"`
	xxx_hidden_Limit                int64                  `protobuf:"varint,3,opt,name=limit,proto3"`
	xxx_hidden_Page                 int64                  `protobuf:"varint,4,opt,name=page,proto3"`
	xxx_hidden_Filter               *DataFilter            `protobuf:"bytes,5,opt,name=filter,proto3"`
	xxx_hidden_Order                *[]*DataOrder          `protobuf:"bytes,6,rep,name=order,proto3"`
	xxx_hidden_UseStrongConsistency bool                   `protobuf:"varint,7,opt,name=use_strong_consistency,json=useStrongConsistency,proto3"`
	xxx_hidden_Target               DeploymentTarget       `protobuf:"varint,8,opt,name=target,proto3,enum=freym.projections.delivery.DeploymentTarget"`
	unknownFields                   protoimpl.UnknownFields
	sizeCache                       protoimpl.SizeCache
}

func (x *GetViewDataListRequest) Reset() {
	*x = GetViewDataListRequest{}
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewDataListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewDataListRequest) ProtoMessage() {}

func (x *GetViewDataListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetViewDataListRequest) GetView() string {
	if x != nil {
		return x.xxx_hidden_View
	}
	return ""
}

func (x *GetViewDataListRequest) GetAuth() *AuthData {
	if x != nil {
		return x.xxx_hidden_Auth
	}
	return nil
}

func (x *GetViewDataListRequest) GetLimit() int64 {
	if x != nil {
		return x.xxx_hidden_Limit
	}
	return 0
}

func (x *GetViewDataListRequest) GetPage() int64 {
	if x != nil {
		return x.xxx_hidden_Page
	}
	return 0
}

func (x *GetViewDataListRequest) GetFilter() *DataFilter {
	if x != nil {
		return x.xxx_hidden_Filter
	}
	return nil
}

func (x *GetViewDataListRequest) GetOrder() []*DataOrder {
	if x != nil {
		if x.xxx_hidden_Order != nil {
			return *x.xxx_hidden_Order
		}
	}
	return nil
}

func (x *GetViewDataListRequest) GetUseStrongConsistency() bool {
	if x != nil {
		return x.xxx_hidden_UseStrongConsistency
	}
	return false
}

func (x *GetViewDataListRequest) GetTarget() DeploymentTarget {
	if x != nil {
		return x.xxx_hidden_Target
	}
	return DeploymentTarget_DEPLOYMENT_TARGET_BLUE
}

func (x *GetViewDataListRequest) SetView(v string) {
	x.xxx_hidden_View = v
}

func (x *GetViewDataListRequest) SetAuth(v *AuthData) {
	x.xxx_hidden_Auth = v
}

func (x *GetViewDataListRequest) SetLimit(v int64) {
	x.xxx_hidden_Limit = v
}

func (x *GetViewDataListRequest) SetPage(v int64) {
	x.xxx_hidden_Page = v
}

func (x *GetViewDataListRequest) SetFilter(v *DataFilter) {
	x.xxx_hidden_Filter = v
}

func (x *GetViewDataListRequest) SetOrder(v []*DataOrder) {
	x.xxx_hidden_Order = &v
}

func (x *GetViewDataListRequest) SetUseStrongConsistency(v bool) {
	x.xxx_hidden_UseStrongConsistency = v
}

func (x *GetViewDataListRequest) SetTarget(v DeploymentTarget) {
	x.xxx_hidden_Target = v
}

func (x *GetViewDataListRequest) HasAuth() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Auth != nil
}

func (x *GetViewDataListRequest) HasFilter() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Filter != nil
}

func (x *GetViewDataListRequest) ClearAuth() {
	x.xxx_hidden_Auth = nil
}

func (x *GetViewDataListRequest) ClearFilter() {
	x.xxx_hidden_Filter = nil
}

type GetViewDataListRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	View                 string
	Auth                 *AuthData
	Limit                int64
	Page                 int64
	Filter               *DataFilter
	Order                []*DataOrder
	UseStrongConsistency bool
	Target               DeploymentTarget
}

func (b0 GetViewDataListRequest_builder) Build() *GetViewDataListRequest {
	m0 := &GetViewDataListRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_View = b.View
	x.xxx_hidden_Auth = b.Auth
	x.xxx_hidden_Limit = b.Limit
	x.xxx_hidden_Page = b.Page
	x.xxx_hidden_Filter = b.Filter
	x.xxx_hidden_Order = &b.Order
	x.xxx_hidden_UseStrongConsistency = b.UseStrongConsistency
	x.xxx_hidden_Target = b.Target
	return m0
}

type GetViewDataListResponse struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Result *[]*Data               `protobuf:"bytes,1,rep,name=result,proto3"`
	xxx_hidden_Limit  int64                  `protobuf:"varint,2,opt,name=limit,proto3"`
	xxx_hidden_Page   int64                  `protobuf:"varint,3,opt,name=page,proto3"`
	xxx_hidden_Total  int64                  `protobuf:"varint,4,opt,name=total,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetViewDataListResponse) Reset() {
	*x = GetViewDataListResponse{}
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewDataListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewDataListResponse) ProtoMessage() {}

func (x *GetViewDataListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_projections_delivery_get_view_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetViewDataListResponse) GetResult() []*Data {
	if x != nil {
		if x.xxx_hidden_Result != nil {
			return *x.xxx_hidden_Result
		}
	}
	return nil
}

func (x *GetViewDataListResponse) GetLimit() int64 {
	if x != nil {
		return x.xxx_hidden_Limit
	}
	return 0
}

func (x *GetViewDataListResponse) GetPage() int64 {
	if x != nil {
		return x.xxx_hidden_Page
	}
	return 0
}

func (x *GetViewDataListResponse) GetTotal() int64 {
	if x != nil {
		return x.xxx_hidden_Total
	}
	return 0
}

func (x *GetViewDataListResponse) SetResult(v []*Data) {
	x.xxx_hidden_Result = &v
}

func (x *GetViewDataListResponse) SetLimit(v int64) {
	x.xxx_hidden_Limit = v
}

func (x *GetViewDataListResponse) SetPage(v int64) {
	x.xxx_hidden_Page = v
}

func (x *GetViewDataListResponse) SetTotal(v int64) {
	x.xxx_hidden_Total = v
}

type GetViewDataListResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Result []*Data
	Limit  int64
	Page   int64
	Total  int64
}

func (b0 GetViewDataListResponse_builder) Build() *GetViewDataListResponse {
	m0 := &GetViewDataListResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Result = &b.Result
	x.xxx_hidden_Limit = b.Limit
	x.xxx_hidden_Page = b.Page
	x.xxx_hidden_Total = b.Total
	return m0
}

var File_projections_delivery_get_view_data_proto protoreflect.FileDescriptor

var file_projections_delivery_get_view_data_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x66, 0x72, 0x65, 0x79,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x38, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x3e,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x34,
	0x0a, 0x16, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x75, 0x73, 0x65, 0x53, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x44, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x89, 0x03, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12, 0x38, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3e,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x75,
	0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x75, 0x73, 0x65,
	0x53, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x44, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x72, 0x65, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_projections_delivery_get_view_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_projections_delivery_get_view_data_proto_goTypes = []any{
	(*GetViewDataRequest)(nil),      // 0: freym.projections.delivery.GetViewDataRequest
	(*GetViewDataResponse)(nil),     // 1: freym.projections.delivery.GetViewDataResponse
	(*GetViewDataListRequest)(nil),  // 2: freym.projections.delivery.GetViewDataListRequest
	(*GetViewDataListResponse)(nil), // 3: freym.projections.delivery.GetViewDataListResponse
	(*AuthData)(nil),                // 4: freym.projections.delivery.AuthData
	(*DataFilter)(nil),              // 5: freym.projections.delivery.DataFilter
	(DeploymentTarget)(0),           // 6: freym.projections.delivery.DeploymentTarget
	(*Data)(nil),                    // 7: freym.projections.delivery.Data
	(*DataOrder)(nil),               // 8: freym.projections.delivery.DataOrder
}
var file_projections_delivery_get_view_data_proto_depIdxs = []int32{
	4, // 0: freym.projections.delivery.GetViewDataRequest.auth:type_name -> freym.projections.delivery.AuthData
	5, // 1: freym.projections.delivery.GetViewDataRequest.filter:type_name -> freym.projections.delivery.DataFilter
	6, // 2: freym.projections.delivery.GetViewDataRequest.target:type_name -> freym.projections.delivery.DeploymentTarget
	7, // 3: freym.projections.delivery.GetViewDataResponse.result:type_name -> freym.projections.delivery.Data
	4, // 4: freym.projections.delivery.GetViewDataListRequest.auth:type_name -> freym.projections.delivery.AuthData
	5, // 5: freym.projections.delivery.GetViewDataListRequest.filter:type_name -> freym.projections.delivery.DataFilter
	8, // 6: freym.projections.delivery.GetViewDataListRequest.order:type_name -> freym.projections.delivery.DataOrder
	6, // 7: freym.projections.delivery.GetViewDataListRequest.target:type_name -> freym.projections.delivery.DeploymentTarget
	7, // 8: freym.projections.delivery.GetViewDataListResponse.result:type_name -> freym.projections.delivery.Data
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_projections_delivery_get_view_data_proto_init() }
func file_projections_delivery_get_view_data_proto_init() {
	if File_projections_delivery_get_view_data_proto != nil {
		return
	}
	file_projections_delivery_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_projections_delivery_get_view_data_proto_rawDesc), len(file_projections_delivery_get_view_data_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_projections_delivery_get_view_data_proto_goTypes,
		DependencyIndexes: file_projections_delivery_get_view_data_proto_depIdxs,
		MessageInfos:      file_projections_delivery_get_view_data_proto_msgTypes,
	}.Build()
	File_projections_delivery_get_view_data_proto = out.File
	file_projections_delivery_get_view_data_proto_goTypes = nil
	file_projections_delivery_get_view_data_proto_depIdxs = nil
}
