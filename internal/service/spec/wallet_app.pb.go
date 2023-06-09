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

type CreateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateUserReply) Reset() {
	*x = CreateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReply) ProtoMessage() {}

func (x *CreateUserReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateUserReply.ProtoReflect.Descriptor instead.
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserReply) GetSuccess() bool {
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

type UserBalanceAccrualReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserBalanceAccrualReply) Reset() {
	*x = UserBalanceAccrualReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBalanceAccrualReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBalanceAccrualReply) ProtoMessage() {}

func (x *UserBalanceAccrualReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserBalanceAccrualReply.ProtoReflect.Descriptor instead.
func (*UserBalanceAccrualReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{3}
}

func (x *UserBalanceAccrualReply) GetSuccess() bool {
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
	IdService uint32  `protobuf:"varint,2,opt,name=idService,proto3" json:"idService,omitempty"`
	IdOrder   uint32  `protobuf:"varint,3,opt,name=idOrder,proto3" json:"idOrder,omitempty"`
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

func (x *UserOrderReservationRequest) GetIdService() uint32 {
	if x != nil {
		return x.IdService
	}
	return 0
}

func (x *UserOrderReservationRequest) GetIdOrder() uint32 {
	if x != nil {
		return x.IdOrder
	}
	return 0
}

func (x *UserOrderReservationRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UserOrderReservationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserOrderReservationReply) Reset() {
	*x = UserOrderReservationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderReservationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderReservationReply) ProtoMessage() {}

func (x *UserOrderReservationReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserOrderReservationReply.ProtoReflect.Descriptor instead.
func (*UserOrderReservationReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{5}
}

func (x *UserOrderReservationReply) GetSuccess() bool {
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
	IdService uint32  `protobuf:"varint,2,opt,name=idService,proto3" json:"idService,omitempty"`
	IdOrder   uint32  `protobuf:"varint,3,opt,name=idOrder,proto3" json:"idOrder,omitempty"`
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

func (x *UserOrderRevenueRequest) GetIdService() uint32 {
	if x != nil {
		return x.IdService
	}
	return 0
}

func (x *UserOrderRevenueRequest) GetIdOrder() uint32 {
	if x != nil {
		return x.IdOrder
	}
	return 0
}

func (x *UserOrderRevenueRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UserOrderRevenueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserOrderRevenueReply) Reset() {
	*x = UserOrderRevenueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderRevenueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderRevenueReply) ProtoMessage() {}

func (x *UserOrderRevenueReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserOrderRevenueReply.ProtoReflect.Descriptor instead.
func (*UserOrderRevenueReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{7}
}

func (x *UserOrderRevenueReply) GetSuccess() bool {
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
	IdService uint32  `protobuf:"varint,2,opt,name=idService,proto3" json:"idService,omitempty"`
	IdOrder   uint32  `protobuf:"varint,3,opt,name=idOrder,proto3" json:"idOrder,omitempty"`
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

func (x *UserOrderDeleteReservationRequest) GetIdService() uint32 {
	if x != nil {
		return x.IdService
	}
	return 0
}

func (x *UserOrderDeleteReservationRequest) GetIdOrder() uint32 {
	if x != nil {
		return x.IdOrder
	}
	return 0
}

func (x *UserOrderDeleteReservationRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type UserOrderDeleteReservationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserOrderDeleteReservationReply) Reset() {
	*x = UserOrderDeleteReservationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderDeleteReservationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderDeleteReservationReply) ProtoMessage() {}

func (x *UserOrderDeleteReservationReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserOrderDeleteReservationReply.ProtoReflect.Descriptor instead.
func (*UserOrderDeleteReservationReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{9}
}

func (x *UserOrderDeleteReservationReply) GetSuccess() bool {
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

type CreateMonthReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateMonthReportReply) Reset() {
	*x = CreateMonthReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMonthReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMonthReportReply) ProtoMessage() {}

func (x *CreateMonthReportReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateMonthReportReply.ProtoReflect.Descriptor instead.
func (*CreateMonthReportReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{11}
}

func (x *CreateMonthReportReply) GetSuccess() bool {
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

type GetBalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance float32 `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Success bool    `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetBalanceReply) Reset() {
	*x = GetBalanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_app_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceReply) ProtoMessage() {}

func (x *GetBalanceReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetBalanceReply.ProtoReflect.Descriptor instead.
func (*GetBalanceReply) Descriptor() ([]byte, []int) {
	return file_wallet_app_proto_rawDescGZIP(), []int{13}
}

func (x *GetBalanceReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBalanceReply) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *GetBalanceReply) GetSuccess() bool {
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
	0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x45, 0x0a, 0x19, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x63, 0x63, 0x72, 0x75, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x33, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x72, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x79, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x22, 0x35, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x75, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x31,
	0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x7f, 0x0a, 0x21, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x22, 0x3b, 0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x44, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x70, 0x65, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*CreateUserRequest)(nil),                 // 0: grpcAvito.CreateUserRequest
	(*CreateUserReply)(nil),                   // 1: grpcAvito.CreateUserReply
	(*UserBalanceAccrualRequest)(nil),         // 2: grpcAvito.UserBalanceAccrualRequest
	(*UserBalanceAccrualReply)(nil),           // 3: grpcAvito.UserBalanceAccrualReply
	(*UserOrderReservationRequest)(nil),       // 4: grpcAvito.UserOrderReservationRequest
	(*UserOrderReservationReply)(nil),         // 5: grpcAvito.UserOrderReservationReply
	(*UserOrderRevenueRequest)(nil),           // 6: grpcAvito.UserOrderRevenueRequest
	(*UserOrderRevenueReply)(nil),             // 7: grpcAvito.UserOrderRevenueReply
	(*UserOrderDeleteReservationRequest)(nil), // 8: grpcAvito.UserOrderDeleteReservationRequest
	(*UserOrderDeleteReservationReply)(nil),   // 9: grpcAvito.UserOrderDeleteReservationReply
	(*CreateMonthReportRequest)(nil),          // 10: grpcAvito.CreateMonthReportRequest
	(*CreateMonthReportReply)(nil),            // 11: grpcAvito.CreateMonthReportReply
	(*GetBalanceRequest)(nil),                 // 12: grpcAvito.GetBalanceRequest
	(*GetBalanceReply)(nil),                   // 13: grpcAvito.GetBalanceReply
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
			switch v := v.(*CreateUserReply); i {
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
			switch v := v.(*UserBalanceAccrualReply); i {
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
			switch v := v.(*UserOrderReservationReply); i {
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
			switch v := v.(*UserOrderRevenueReply); i {
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
			switch v := v.(*UserOrderDeleteReservationReply); i {
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
			switch v := v.(*CreateMonthReportReply); i {
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
			switch v := v.(*GetBalanceReply); i {
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
