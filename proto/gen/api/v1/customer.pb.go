// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: api/v1/customer.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Customer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender        string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	IsActive      bool                   `protobuf:"varint,6,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     int64                  `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_api_v1_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Customer) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Customer) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Customer) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Customer) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerRequest) Reset() {
	*x = CustomerRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerRequest) ProtoMessage() {}

func (x *CustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerRequest.ProtoReflect.Descriptor instead.
func (*CustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CustomerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Customer              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CustomerResponse) GetData() *Customer {
	if x != nil {
		return x.Data
	}
	return nil
}

type CustomerListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []int32                `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Search        string                 `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerListRequest) Reset() {
	*x = CustomerListRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerListRequest) ProtoMessage() {}

func (x *CustomerListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerListRequest.ProtoReflect.Descriptor instead.
func (*CustomerListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{3}
}

func (x *CustomerListRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CustomerListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type CustomerListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*Customer            `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerListResponse) Reset() {
	*x = CustomerListResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerListResponse) ProtoMessage() {}

func (x *CustomerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerListResponse.ProtoReflect.Descriptor instead.
func (*CustomerListResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CustomerListResponse) GetData() []*Customer {
	if x != nil {
		return x.Data
	}
	return nil
}

type CustomerCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	IsActive      bool                   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerCreateRequest) Reset() {
	*x = CustomerCreateRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerCreateRequest) ProtoMessage() {}

func (x *CustomerCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerCreateRequest.ProtoReflect.Descriptor instead.
func (*CustomerCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{5}
}

func (x *CustomerCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerCreateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CustomerCreateRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CustomerCreateRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type CustomerCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Customer              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerCreateResponse) Reset() {
	*x = CustomerCreateResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerCreateResponse) ProtoMessage() {}

func (x *CustomerCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerCreateResponse.ProtoReflect.Descriptor instead.
func (*CustomerCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{6}
}

func (x *CustomerCreateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerCreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CustomerCreateResponse) GetData() *Customer {
	if x != nil {
		return x.Data
	}
	return nil
}

type CustomerUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	IsActive      bool                   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerUpdateRequest) Reset() {
	*x = CustomerUpdateRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUpdateRequest) ProtoMessage() {}

func (x *CustomerUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUpdateRequest.ProtoReflect.Descriptor instead.
func (*CustomerUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{7}
}

func (x *CustomerUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerUpdateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CustomerUpdateRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CustomerUpdateRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type CustomerUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Customer              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerUpdateResponse) Reset() {
	*x = CustomerUpdateResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUpdateResponse) ProtoMessage() {}

func (x *CustomerUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUpdateResponse.ProtoReflect.Descriptor instead.
func (*CustomerUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{8}
}

func (x *CustomerUpdateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerUpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CustomerUpdateResponse) GetData() *Customer {
	if x != nil {
		return x.Data
	}
	return nil
}

type CustomerDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerDeleteRequest) Reset() {
	*x = CustomerDeleteRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDeleteRequest) ProtoMessage() {}

func (x *CustomerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDeleteRequest.ProtoReflect.Descriptor instead.
func (*CustomerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{9}
}

func (x *CustomerDeleteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CustomerDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerDeleteResponse) Reset() {
	*x = CustomerDeleteResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDeleteResponse) ProtoMessage() {}

func (x *CustomerDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDeleteResponse.ProtoReflect.Descriptor instead.
func (*CustomerDeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{10}
}

func (x *CustomerDeleteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CustomerActivateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerActivateRequest) Reset() {
	*x = CustomerActivateRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerActivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerActivateRequest) ProtoMessage() {}

func (x *CustomerActivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerActivateRequest.ProtoReflect.Descriptor instead.
func (*CustomerActivateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{11}
}

func (x *CustomerActivateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CustomerActivateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerActivateResponse) Reset() {
	*x = CustomerActivateResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerActivateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerActivateResponse) ProtoMessage() {}

func (x *CustomerActivateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerActivateResponse.ProtoReflect.Descriptor instead.
func (*CustomerActivateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{12}
}

func (x *CustomerActivateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerActivateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CustomerDeactivateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerDeactivateRequest) Reset() {
	*x = CustomerDeactivateRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerDeactivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDeactivateRequest) ProtoMessage() {}

func (x *CustomerDeactivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDeactivateRequest.ProtoReflect.Descriptor instead.
func (*CustomerDeactivateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{13}
}

func (x *CustomerDeactivateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CustomerDeactivateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerDeactivateResponse) Reset() {
	*x = CustomerDeactivateResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerDeactivateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDeactivateResponse) ProtoMessage() {}

func (x *CustomerDeactivateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDeactivateResponse.ProtoReflect.Descriptor instead.
func (*CustomerDeactivateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{14}
}

func (x *CustomerDeactivateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerDeactivateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CustomerRegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerRegisterRequest) Reset() {
	*x = CustomerRegisterRequest{}
	mi := &file_api_v1_customer_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerRegisterRequest) ProtoMessage() {}

func (x *CustomerRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerRegisterRequest.ProtoReflect.Descriptor instead.
func (*CustomerRegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{15}
}

func (x *CustomerRegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerRegisterRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CustomerRegisterRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type CustomerRegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *Customer              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerRegisterResponse) Reset() {
	*x = CustomerRegisterResponse{}
	mi := &file_api_v1_customer_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerRegisterResponse) ProtoMessage() {}

func (x *CustomerRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_customer_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerRegisterResponse.ProtoReflect.Descriptor instead.
func (*CustomerRegisterResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_customer_proto_rawDescGZIP(), []int{16}
}

func (x *CustomerRegisterResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CustomerRegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CustomerRegisterResponse) GetData() *Customer {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_v1_customer_proto protoreflect.FileDescriptor

const file_api_v1_customer_proto_rawDesc = "" +
	"\n" +
	"\x15api/v1/customer.proto\x12\x06api.v1\"\xca\x01\n" +
	"\bCustomer\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x05 \x01(\tR\x06gender\x12\x1a\n" +
	"\bisActive\x18\x06 \x01(\bR\bisActive\x12\x1c\n" +
	"\tcreatedAt\x18\a \x01(\x03R\tcreatedAt\x12\x1c\n" +
	"\tupdatedAt\x18\b \x01(\x03R\tupdatedAt\"!\n" +
	"\x0fCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"r\n" +
	"\x10CustomerResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12$\n" +
	"\x04data\x18\x03 \x01(\v2\x10.api.v1.CustomerR\x04data\"?\n" +
	"\x13CustomerListRequest\x12\x10\n" +
	"\x03ids\x18\x01 \x03(\x05R\x03ids\x12\x16\n" +
	"\x06search\x18\x02 \x01(\tR\x06search\"v\n" +
	"\x14CustomerListResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12$\n" +
	"\x04data\x18\x03 \x03(\v2\x10.api.v1.CustomerR\x04data\"\x8c\x01\n" +
	"\x15CustomerCreateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x04 \x01(\tR\x06gender\x12\x1b\n" +
	"\tis_active\x18\x05 \x01(\bR\bisActive\"x\n" +
	"\x16CustomerCreateResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12$\n" +
	"\x04data\x18\x03 \x01(\v2\x10.api.v1.CustomerR\x04data\"\x8c\x01\n" +
	"\x15CustomerUpdateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x04 \x01(\tR\x06gender\x12\x1b\n" +
	"\tis_active\x18\x05 \x01(\bR\bisActive\"x\n" +
	"\x16CustomerUpdateResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12$\n" +
	"\x04data\x18\x03 \x01(\v2\x10.api.v1.CustomerR\x04data\"'\n" +
	"\x15CustomerDeleteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"R\n" +
	"\x16CustomerDeleteResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\")\n" +
	"\x17CustomerActivateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"T\n" +
	"\x18CustomerActivateResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"+\n" +
	"\x19CustomerDeactivateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"V\n" +
	"\x1aCustomerDeactivateResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"q\n" +
	"\x17CustomerRegisterRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x04 \x01(\tR\x06gender\"z\n" +
	"\x18CustomerRegisterResponse\x12\x1e\n" +
	"\n" +
	"statusCode\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12$\n" +
	"\x04data\x18\x03 \x01(\v2\x10.api.v1.CustomerR\x04data2\xfe\x04\n" +
	"\x0eCustomerMethod\x12F\n" +
	"\aGetList\x12\x1b.api.v1.CustomerListRequest\x1a\x1c.api.v1.CustomerListResponse\"\x00\x12>\n" +
	"\aGetById\x12\x17.api.v1.CustomerRequest\x1a\x18.api.v1.CustomerResponse\"\x00\x12I\n" +
	"\x06Create\x12\x1d.api.v1.CustomerCreateRequest\x1a\x1e.api.v1.CustomerCreateResponse\"\x00\x12I\n" +
	"\x06Update\x12\x1d.api.v1.CustomerUpdateRequest\x1a\x1e.api.v1.CustomerUpdateResponse\"\x00\x12M\n" +
	"\n" +
	"DeleteById\x12\x1d.api.v1.CustomerDeleteRequest\x1a\x1e.api.v1.CustomerDeleteResponse\"\x00\x12S\n" +
	"\fActivateById\x12\x1f.api.v1.CustomerActivateRequest\x1a .api.v1.CustomerActivateResponse\"\x00\x12Y\n" +
	"\x0eDeactivateById\x12!.api.v1.CustomerDeactivateRequest\x1a\".api.v1.CustomerDeactivateResponse\"\x00\x12O\n" +
	"\bRegister\x12\x1f.api.v1.CustomerRegisterRequest\x1a .api.v1.CustomerRegisterResponse\"\x00B$Z\"hexagonal-go-grpc/proto/gen/api/v1b\x06proto3"

var (
	file_api_v1_customer_proto_rawDescOnce sync.Once
	file_api_v1_customer_proto_rawDescData []byte
)

func file_api_v1_customer_proto_rawDescGZIP() []byte {
	file_api_v1_customer_proto_rawDescOnce.Do(func() {
		file_api_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_customer_proto_rawDesc), len(file_api_v1_customer_proto_rawDesc)))
	})
	return file_api_v1_customer_proto_rawDescData
}

var file_api_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_api_v1_customer_proto_goTypes = []any{
	(*Customer)(nil),                   // 0: api.v1.Customer
	(*CustomerRequest)(nil),            // 1: api.v1.CustomerRequest
	(*CustomerResponse)(nil),           // 2: api.v1.CustomerResponse
	(*CustomerListRequest)(nil),        // 3: api.v1.CustomerListRequest
	(*CustomerListResponse)(nil),       // 4: api.v1.CustomerListResponse
	(*CustomerCreateRequest)(nil),      // 5: api.v1.CustomerCreateRequest
	(*CustomerCreateResponse)(nil),     // 6: api.v1.CustomerCreateResponse
	(*CustomerUpdateRequest)(nil),      // 7: api.v1.CustomerUpdateRequest
	(*CustomerUpdateResponse)(nil),     // 8: api.v1.CustomerUpdateResponse
	(*CustomerDeleteRequest)(nil),      // 9: api.v1.CustomerDeleteRequest
	(*CustomerDeleteResponse)(nil),     // 10: api.v1.CustomerDeleteResponse
	(*CustomerActivateRequest)(nil),    // 11: api.v1.CustomerActivateRequest
	(*CustomerActivateResponse)(nil),   // 12: api.v1.CustomerActivateResponse
	(*CustomerDeactivateRequest)(nil),  // 13: api.v1.CustomerDeactivateRequest
	(*CustomerDeactivateResponse)(nil), // 14: api.v1.CustomerDeactivateResponse
	(*CustomerRegisterRequest)(nil),    // 15: api.v1.CustomerRegisterRequest
	(*CustomerRegisterResponse)(nil),   // 16: api.v1.CustomerRegisterResponse
}
var file_api_v1_customer_proto_depIdxs = []int32{
	0,  // 0: api.v1.CustomerResponse.data:type_name -> api.v1.Customer
	0,  // 1: api.v1.CustomerListResponse.data:type_name -> api.v1.Customer
	0,  // 2: api.v1.CustomerCreateResponse.data:type_name -> api.v1.Customer
	0,  // 3: api.v1.CustomerUpdateResponse.data:type_name -> api.v1.Customer
	0,  // 4: api.v1.CustomerRegisterResponse.data:type_name -> api.v1.Customer
	3,  // 5: api.v1.CustomerMethod.GetList:input_type -> api.v1.CustomerListRequest
	1,  // 6: api.v1.CustomerMethod.GetById:input_type -> api.v1.CustomerRequest
	5,  // 7: api.v1.CustomerMethod.Create:input_type -> api.v1.CustomerCreateRequest
	7,  // 8: api.v1.CustomerMethod.Update:input_type -> api.v1.CustomerUpdateRequest
	9,  // 9: api.v1.CustomerMethod.DeleteById:input_type -> api.v1.CustomerDeleteRequest
	11, // 10: api.v1.CustomerMethod.ActivateById:input_type -> api.v1.CustomerActivateRequest
	13, // 11: api.v1.CustomerMethod.DeactivateById:input_type -> api.v1.CustomerDeactivateRequest
	15, // 12: api.v1.CustomerMethod.Register:input_type -> api.v1.CustomerRegisterRequest
	4,  // 13: api.v1.CustomerMethod.GetList:output_type -> api.v1.CustomerListResponse
	2,  // 14: api.v1.CustomerMethod.GetById:output_type -> api.v1.CustomerResponse
	6,  // 15: api.v1.CustomerMethod.Create:output_type -> api.v1.CustomerCreateResponse
	8,  // 16: api.v1.CustomerMethod.Update:output_type -> api.v1.CustomerUpdateResponse
	10, // 17: api.v1.CustomerMethod.DeleteById:output_type -> api.v1.CustomerDeleteResponse
	12, // 18: api.v1.CustomerMethod.ActivateById:output_type -> api.v1.CustomerActivateResponse
	14, // 19: api.v1.CustomerMethod.DeactivateById:output_type -> api.v1.CustomerDeactivateResponse
	16, // 20: api.v1.CustomerMethod.Register:output_type -> api.v1.CustomerRegisterResponse
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_customer_proto_init() }
func file_api_v1_customer_proto_init() {
	if File_api_v1_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_customer_proto_rawDesc), len(file_api_v1_customer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_customer_proto_goTypes,
		DependencyIndexes: file_api_v1_customer_proto_depIdxs,
		MessageInfos:      file_api_v1_customer_proto_msgTypes,
	}.Build()
	File_api_v1_customer_proto = out.File
	file_api_v1_customer_proto_goTypes = nil
	file_api_v1_customer_proto_depIdxs = nil
}
