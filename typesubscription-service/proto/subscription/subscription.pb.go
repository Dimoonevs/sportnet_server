// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: subscription.proto

package subscription

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

type TypeSub int32

const (
	TypeSub_FIXED_COUNT  TypeSub = 0
	TypeSub_DATE_LIMITED TypeSub = 1
)

// Enum value maps for TypeSub.
var (
	TypeSub_name = map[int32]string{
		0: "FIXED_COUNT",
		1: "DATE_LIMITED",
	}
	TypeSub_value = map[string]int32{
		"FIXED_COUNT":  0,
		"DATE_LIMITED": 1,
	}
)

func (x TypeSub) Enum() *TypeSub {
	p := new(TypeSub)
	*p = x
	return p
}

func (x TypeSub) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeSub) Descriptor() protoreflect.EnumDescriptor {
	return file_subscription_proto_enumTypes[0].Descriptor()
}

func (TypeSub) Type() protoreflect.EnumType {
	return &file_subscription_proto_enumTypes[0]
}

func (x TypeSub) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeSub.Descriptor instead.
func (TypeSub) EnumDescriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

type TimeLimited int32

const (
	TimeLimited_WEEK   TimeLimited = 0
	TimeLimited_MONTH  TimeLimited = 1
	TimeLimited_YEAR   TimeLimited = 2
	TimeLimited_CUSTOM TimeLimited = 3
)

// Enum value maps for TimeLimited.
var (
	TimeLimited_name = map[int32]string{
		0: "WEEK",
		1: "MONTH",
		2: "YEAR",
		3: "CUSTOM",
	}
	TimeLimited_value = map[string]int32{
		"WEEK":   0,
		"MONTH":  1,
		"YEAR":   2,
		"CUSTOM": 3,
	}
)

func (x TimeLimited) Enum() *TimeLimited {
	p := new(TimeLimited)
	*p = x
	return p
}

func (x TimeLimited) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeLimited) Descriptor() protoreflect.EnumDescriptor {
	return file_subscription_proto_enumTypes[1].Descriptor()
}

func (TimeLimited) Type() protoreflect.EnumType {
	return &file_subscription_proto_enumTypes[1]
}

func (x TimeLimited) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeLimited.Descriptor instead.
func (TimeLimited) EnumDescriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                    string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description             string              `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	StatusSubscription      *StatusSubscription `protobuf:"bytes,3,opt,name=statusSubscription,proto3" json:"statusSubscription,omitempty"`
	Price                   *PriceSubscription  `protobuf:"bytes,4,opt,name=Price,proto3" json:"Price,omitempty"`
	CoachId                 int32               `protobuf:"varint,5,opt,name=coachId,proto3" json:"coachId,omitempty"`
	DaysOfWeek              []string            `protobuf:"bytes,6,rep,name=daysOfWeek,proto3" json:"daysOfWeek,omitempty"`
	AutomaticallyManagement bool                `protobuf:"varint,7,opt,name=automaticallyManagement,proto3" json:"automaticallyManagement,omitempty"`
	Time                    []string            `protobuf:"bytes,8,rep,name=time,proto3" json:"time,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriptionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubscriptionRequest) GetStatusSubscription() *StatusSubscription {
	if x != nil {
		return x.StatusSubscription
	}
	return nil
}

func (x *SubscriptionRequest) GetPrice() *PriceSubscription {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *SubscriptionRequest) GetCoachId() int32 {
	if x != nil {
		return x.CoachId
	}
	return 0
}

func (x *SubscriptionRequest) GetDaysOfWeek() []string {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

func (x *SubscriptionRequest) GetAutomaticallyManagement() bool {
	if x != nil {
		return x.AutomaticallyManagement
	}
	return false
}

func (x *SubscriptionRequest) GetTime() []string {
	if x != nil {
		return x.Time
	}
	return nil
}

type StatusSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeSub           TypeSub     `protobuf:"varint,1,opt,name=typeSub,proto3,enum=subscription.TypeSub" json:"typeSub,omitempty"`
	TimeLimited       TimeLimited `protobuf:"varint,2,opt,name=timeLimited,proto3,enum=subscription.TimeLimited" json:"timeLimited,omitempty"`
	CustomTimeLimited int32       `protobuf:"varint,3,opt,name=customTimeLimited,proto3" json:"customTimeLimited,omitempty"`
}

func (x *StatusSubscription) Reset() {
	*x = StatusSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSubscription) ProtoMessage() {}

func (x *StatusSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSubscription.ProtoReflect.Descriptor instead.
func (*StatusSubscription) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *StatusSubscription) GetTypeSub() TypeSub {
	if x != nil {
		return x.TypeSub
	}
	return TypeSub_FIXED_COUNT
}

func (x *StatusSubscription) GetTimeLimited() TimeLimited {
	if x != nil {
		return x.TimeLimited
	}
	return TimeLimited_WEEK
}

func (x *StatusSubscription) GetCustomTimeLimited() int32 {
	if x != nil {
		return x.CustomTimeLimited
	}
	return 0
}

type PriceSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price    int32  `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *PriceSubscription) Reset() {
	*x = PriceSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceSubscription) ProtoMessage() {}

func (x *PriceSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceSubscription.ProtoReflect.Descriptor instead.
func (*PriceSubscription) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *PriceSubscription) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PriceSubscription) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type SubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id      int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubscriptionResponse) Reset() {
	*x = SubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionResponse) ProtoMessage() {}

func (x *SubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SubscriptionResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*SubscriptionData `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *GetSubscriptionResponse) Reset() {
	*x = GetSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionResponse) ProtoMessage() {}

func (x *GetSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *GetSubscriptionResponse) GetSubscriptions() []*SubscriptionData {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachId int32 `protobuf:"varint,1,opt,name=coachId,proto3" json:"coachId,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *GetSubscriptionRequest) GetCoachId() int32 {
	if x != nil {
		return x.CoachId
	}
	return 0
}

type SubscriptionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description             string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TypeSub                 string   `protobuf:"bytes,4,opt,name=typeSub,proto3" json:"typeSub,omitempty"`
	TimeLimited             string   `protobuf:"bytes,5,opt,name=timeLimited,proto3" json:"timeLimited,omitempty"`
	CustomTimeLimited       int32    `protobuf:"varint,6,opt,name=customTimeLimited,proto3" json:"customTimeLimited,omitempty"`
	Price                   int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Currency                string   `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	CoachId                 int32    `protobuf:"varint,9,opt,name=coachId,proto3" json:"coachId,omitempty"`
	DaysOfWeek              []string `protobuf:"bytes,10,rep,name=daysOfWeek,proto3" json:"daysOfWeek,omitempty"`
	AutomaticallyManagement bool     `protobuf:"varint,11,opt,name=automaticallyManagement,proto3" json:"automaticallyManagement,omitempty"`
	Time                    []string `protobuf:"bytes,12,rep,name=time,proto3" json:"time,omitempty"`
	CronId                  int32    `protobuf:"varint,13,opt,name=cronId,proto3" json:"cronId,omitempty"`
	IdScheduler             int32    `protobuf:"varint,14,opt,name=idScheduler,proto3" json:"idScheduler,omitempty"`
}

func (x *SubscriptionData) Reset() {
	*x = SubscriptionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionData) ProtoMessage() {}

func (x *SubscriptionData) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionData.ProtoReflect.Descriptor instead.
func (*SubscriptionData) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *SubscriptionData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SubscriptionData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriptionData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubscriptionData) GetTypeSub() string {
	if x != nil {
		return x.TypeSub
	}
	return ""
}

func (x *SubscriptionData) GetTimeLimited() string {
	if x != nil {
		return x.TimeLimited
	}
	return ""
}

func (x *SubscriptionData) GetCustomTimeLimited() int32 {
	if x != nil {
		return x.CustomTimeLimited
	}
	return 0
}

func (x *SubscriptionData) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SubscriptionData) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SubscriptionData) GetCoachId() int32 {
	if x != nil {
		return x.CoachId
	}
	return 0
}

func (x *SubscriptionData) GetDaysOfWeek() []string {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

func (x *SubscriptionData) GetAutomaticallyManagement() bool {
	if x != nil {
		return x.AutomaticallyManagement
	}
	return false
}

func (x *SubscriptionData) GetTime() []string {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *SubscriptionData) GetCronId() int32 {
	if x != nil {
		return x.CronId
	}
	return 0
}

func (x *SubscriptionData) GetIdScheduler() int32 {
	if x != nil {
		return x.IdScheduler
	}
	return 0
}

type SubscriptionEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                    string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description             string              `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	StatusSubscription      *StatusSubscription `protobuf:"bytes,3,opt,name=statusSubscription,proto3" json:"statusSubscription,omitempty"`
	Price                   *PriceSubscription  `protobuf:"bytes,4,opt,name=Price,proto3" json:"Price,omitempty"`
	CoachId                 int32               `protobuf:"varint,5,opt,name=coachId,proto3" json:"coachId,omitempty"`
	DaysOfWeek              []string            `protobuf:"bytes,6,rep,name=daysOfWeek,proto3" json:"daysOfWeek,omitempty"`
	AutomaticallyManagement bool                `protobuf:"varint,7,opt,name=automaticallyManagement,proto3" json:"automaticallyManagement,omitempty"`
	Time                    []string            `protobuf:"bytes,8,rep,name=time,proto3" json:"time,omitempty"`
	CronId                  int32               `protobuf:"varint,9,opt,name=cronId,proto3" json:"cronId,omitempty"`
	IdScheduler             int32               `protobuf:"varint,10,opt,name=idScheduler,proto3" json:"idScheduler,omitempty"`
	Id                      int32               `protobuf:"varint,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubscriptionEditRequest) Reset() {
	*x = SubscriptionEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionEditRequest) ProtoMessage() {}

func (x *SubscriptionEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionEditRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionEditRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *SubscriptionEditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriptionEditRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubscriptionEditRequest) GetStatusSubscription() *StatusSubscription {
	if x != nil {
		return x.StatusSubscription
	}
	return nil
}

func (x *SubscriptionEditRequest) GetPrice() *PriceSubscription {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *SubscriptionEditRequest) GetCoachId() int32 {
	if x != nil {
		return x.CoachId
	}
	return 0
}

func (x *SubscriptionEditRequest) GetDaysOfWeek() []string {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

func (x *SubscriptionEditRequest) GetAutomaticallyManagement() bool {
	if x != nil {
		return x.AutomaticallyManagement
	}
	return false
}

func (x *SubscriptionEditRequest) GetTime() []string {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *SubscriptionEditRequest) GetCronId() int32 {
	if x != nil {
		return x.CronId
	}
	return 0
}

func (x *SubscriptionEditRequest) GetIdScheduler() int32 {
	if x != nil {
		return x.IdScheduler
	}
	return 0
}

func (x *SubscriptionEditRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xdc, 0x02, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x50, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x61,
	0x63, 0x68, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x38, 0x0a, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x6c, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x6c, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x53, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x62,
	0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x53, 0x75, 0x62, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x40, 0x0a, 0x14, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x32,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x49, 0x64, 0x22, 0xb6, 0x03, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x79, 0x70, 0x65, 0x53, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x79, 0x70, 0x65, 0x53, 0x75, 0x62, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69,
	0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x61,
	0x63, 0x68, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x38, 0x0a, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x6c, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x6c, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x63, 0x72, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x69, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0xaa, 0x03, 0x0a, 0x17,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a,
	0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x35, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b,
	0x12, 0x38, 0x0a, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x6c,
	0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x72, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x63, 0x72, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x2c, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x75, 0x62, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x58, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d,
	0x49, 0x54, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x38, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x45, 0x4b, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x45,
	0x41, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x03,
	0x32, 0xbf, 0x02, 0x0a, 0x17, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscription_proto_rawDescOnce sync.Once
	file_subscription_proto_rawDescData = file_subscription_proto_rawDesc
)

func file_subscription_proto_rawDescGZIP() []byte {
	file_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscription_proto_rawDescData)
	})
	return file_subscription_proto_rawDescData
}

var file_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_subscription_proto_goTypes = []interface{}{
	(TypeSub)(0),                    // 0: subscription.TypeSub
	(TimeLimited)(0),                // 1: subscription.TimeLimited
	(*SubscriptionRequest)(nil),     // 2: subscription.SubscriptionRequest
	(*StatusSubscription)(nil),      // 3: subscription.StatusSubscription
	(*PriceSubscription)(nil),       // 4: subscription.PriceSubscription
	(*SubscriptionResponse)(nil),    // 5: subscription.SubscriptionResponse
	(*GetSubscriptionResponse)(nil), // 6: subscription.GetSubscriptionResponse
	(*GetSubscriptionRequest)(nil),  // 7: subscription.GetSubscriptionRequest
	(*SubscriptionData)(nil),        // 8: subscription.SubscriptionData
	(*SubscriptionEditRequest)(nil), // 9: subscription.SubscriptionEditRequest
}
var file_subscription_proto_depIdxs = []int32{
	3,  // 0: subscription.SubscriptionRequest.statusSubscription:type_name -> subscription.StatusSubscription
	4,  // 1: subscription.SubscriptionRequest.Price:type_name -> subscription.PriceSubscription
	0,  // 2: subscription.StatusSubscription.typeSub:type_name -> subscription.TypeSub
	1,  // 3: subscription.StatusSubscription.timeLimited:type_name -> subscription.TimeLimited
	8,  // 4: subscription.GetSubscriptionResponse.subscriptions:type_name -> subscription.SubscriptionData
	3,  // 5: subscription.SubscriptionEditRequest.statusSubscription:type_name -> subscription.StatusSubscription
	4,  // 6: subscription.SubscriptionEditRequest.Price:type_name -> subscription.PriceSubscription
	2,  // 7: subscription.TypeSubscriptionService.CreateSubscription:input_type -> subscription.SubscriptionRequest
	7,  // 8: subscription.TypeSubscriptionService.GetAllSubscriptions:input_type -> subscription.GetSubscriptionRequest
	9,  // 9: subscription.TypeSubscriptionService.EditSubscription:input_type -> subscription.SubscriptionEditRequest
	5,  // 10: subscription.TypeSubscriptionService.CreateSubscription:output_type -> subscription.SubscriptionResponse
	6,  // 11: subscription.TypeSubscriptionService.GetAllSubscriptions:output_type -> subscription.GetSubscriptionResponse
	5,  // 12: subscription.TypeSubscriptionService.EditSubscription:output_type -> subscription.SubscriptionResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusSubscription); i {
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
		file_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceSubscription); i {
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
		file_subscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionData); i {
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
		file_subscription_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionEditRequest); i {
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
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
		EnumInfos:         file_subscription_proto_enumTypes,
		MessageInfos:      file_subscription_proto_msgTypes,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}
