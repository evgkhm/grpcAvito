// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: wallet_app.proto

package spec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance float32 `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateUserRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UserBalanceAccrualRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance float32 `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *UserBalanceAccrualRequest) Reset() {
	*x = UserBalanceAccrualRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBalanceAccrualRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBalanceAccrualRequest) ProtoMessage() {}

func (x *UserBalanceAccrualRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBalanceAccrualRequest.ProtoReflect.Descriptor instead.
func (*UserBalanceAccrualRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{2}
}

func (x *UserBalanceAccrualRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserBalanceAccrualRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type UserBalanceAccrualResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserBalanceAccrualResponse) Reset() {
	*x = UserBalanceAccrualResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBalanceAccrualResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBalanceAccrualResponse) ProtoMessage() {}

func (x *UserBalanceAccrualResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBalanceAccrualResponse.ProtoReflect.Descriptor instead.
func (*UserBalanceAccrualResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{3}
}

func (x *UserBalanceAccrualResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UserOrderReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceId uint32  `protobuf:"varint,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	OrderId   uint32  `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Cost      float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *UserOrderReservationRequest) Reset() {
	*x = UserOrderReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderReservationRequest) ProtoMessage() {}

func (x *UserOrderReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderReservationRequest.ProtoReflect.Descriptor instead.
func (*UserOrderReservationRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{4}
}

func (x *UserOrderReservationRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserOrderReservationRequest) GetServiceId() uint32 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *UserOrderReservationRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UserOrderReservationRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UserOrderReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserOrderReservationResponse) Reset() {
	*x = UserOrderReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderReservationResponse) ProtoMessage() {}

func (x *UserOrderReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderReservationResponse.ProtoReflect.Descriptor instead.
func (*UserOrderReservationResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{5}
}

func (x *UserOrderReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UserOrderRevenueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceId uint32  `protobuf:"varint,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	OrderId   uint32  `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Cost      float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *UserOrderRevenueRequest) Reset() {
	*x = UserOrderRevenueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderRevenueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderRevenueRequest) ProtoMessage() {}

func (x *UserOrderRevenueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderRevenueRequest.ProtoReflect.Descriptor instead.
func (*UserOrderRevenueRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{6}
}

func (x *UserOrderRevenueRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserOrderRevenueRequest) GetServiceId() uint32 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *UserOrderRevenueRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UserOrderRevenueRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UserOrderRevenueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserOrderRevenueResponse) Reset() {
	*x = UserOrderRevenueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderRevenueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderRevenueResponse) ProtoMessage() {}

func (x *UserOrderRevenueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderRevenueResponse.ProtoReflect.Descriptor instead.
func (*UserOrderRevenueResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{7}
}

func (x *UserOrderRevenueResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UserOrderDeleteReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceId uint32  `protobuf:"varint,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	OrderId   uint32  `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Cost      float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *UserOrderDeleteReservationRequest) Reset() {
	*x = UserOrderDeleteReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderDeleteReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderDeleteReservationRequest) ProtoMessage() {}

func (x *UserOrderDeleteReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderDeleteReservationRequest.ProtoReflect.Descriptor instead.
func (*UserOrderDeleteReservationRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{8}
}

func (x *UserOrderDeleteReservationRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserOrderDeleteReservationRequest) GetServiceId() uint32 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *UserOrderDeleteReservationRequest) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UserOrderDeleteReservationRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UserOrderDeleteReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserOrderDeleteReservationResponse) Reset() {
	*x = UserOrderDeleteReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderDeleteReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderDeleteReservationResponse) ProtoMessage() {}

func (x *UserOrderDeleteReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderDeleteReservationResponse.ProtoReflect.Descriptor instead.
func (*UserOrderDeleteReservationResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{9}
}

func (x *UserOrderDeleteReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateMonthReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  uint32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month uint32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *CreateMonthReportRequest) Reset() {
	*x = CreateMonthReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMonthReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMonthReportRequest) ProtoMessage() {}

func (x *CreateMonthReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMonthReportRequest.ProtoReflect.Descriptor instead.
func (*CreateMonthReportRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{10}
}

func (x *CreateMonthReportRequest) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CreateMonthReportRequest) GetMonth() uint32 {
	if x != nil {
		return x.Month
	}
	return 0
}

type CreateMonthReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateMonthReportResponse) Reset() {
	*x = CreateMonthReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMonthReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMonthReportResponse) ProtoMessage() {}

func (x *CreateMonthReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMonthReportResponse.ProtoReflect.Descriptor instead.
func (*CreateMonthReportResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{11}
}

func (x *CreateMonthReportResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{12}
}

func (x *GetBalanceRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance float32 `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Success bool    `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_app_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{13}
}

func (x *GetBalanceResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBalanceResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *GetBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_wallet_app_proto protoreflect.FileDescriptor

var file_wallet_app_proto_rawDesc = []byte{
	0x0a, 0x10, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x41, 0x76, 0x69, 0x74, 0x6f, 0x22, 0x3d, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x45, 0x0a, 0x19,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x63, 0x63, 0x72, 0x75,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x72, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x79, 0x0a, 0x1b, 0x55,
	0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x75, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x76,
	0x65, 0x6e, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x7f, 0x0a,
	0x21, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x3e,
	0x0a, 0x22, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x44,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x22, 0x35, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x58, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x2e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_app_proto_rawDescOnce sync.Once
	file_wallet_app_proto_rawDescData = file_wallet_app_proto_rawDesc
)

func file_wallet_app_proto_rawDescGZIP() []byte {
	file_wallet_app_proto_rawDescOnce.Do(func() {
		file_wallet_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_app_proto_rawDescData)
	})
	return file_wallet_app_proto_rawDescData
}

var file_wallet_app_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_wallet_app_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),                  // 0: grpcAvito.CreateUserRequest
	(*CreateUserResponse)(nil),                 // 1: grpcAvito.CreateUserResponse
	(*UserBalanceAccrualRequest)(nil),          // 2: grpcAvito.UserBalanceAccrualRequest
	(*UserBalanceAccrualResponse)(nil),         // 3: grpcAvito.UserBalanceAccrualResponse
	(*UserOrderReservationRequest)(nil),        // 4: grpcAvito.UserOrderReservationRequest
	(*UserOrderReservationResponse)(nil),       // 5: grpcAvito.UserOrderReservationResponse
	(*UserOrderRevenueRequest)(nil),            // 6: grpcAvito.UserOrderRevenueRequest
	(*UserOrderRevenueResponse)(nil),           // 7: grpcAvito.UserOrderRevenueResponse
	(*UserOrderDeleteReservationRequest)(nil),  // 8: grpcAvito.UserOrderDeleteReservationRequest
	(*UserOrderDeleteReservationResponse)(nil), // 9: grpcAvito.UserOrderDeleteReservationResponse
	(*CreateMonthReportRequest)(nil),           // 10: grpcAvito.CreateMonthReportRequest
	(*CreateMonthReportResponse)(nil),          // 11: grpcAvito.CreateMonthReportResponse
	(*GetBalanceRequest)(nil),                  // 12: grpcAvito.GetBalanceRequest
	(*GetBalanceResponse)(nil),                 // 13: grpcAvito.GetBalanceResponse
}
var file_wallet_app_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wallet_app_proto_init() }
func file_wallet_app_proto_init() {
	if File_wallet_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBalanceAccrualRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBalanceAccrualResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderReservationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderReservationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderRevenueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderRevenueResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderDeleteReservationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderDeleteReservationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMonthReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMonthReportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_app_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wallet_app_proto_goTypes,
		DependencyIndexes: file_wallet_app_proto_depIdxs,
		MessageInfos:      file_wallet_app_proto_msgTypes,
	}.Build()
	File_wallet_app_proto = out.File
	file_wallet_app_proto_rawDesc = nil
	file_wallet_app_proto_goTypes = nil
	file_wallet_app_proto_depIdxs = nil
}