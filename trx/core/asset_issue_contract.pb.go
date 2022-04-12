// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.3
// source: core/contract/asset_issue_contract.proto

package core

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssetIssueContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string                             `protobuf:"bytes,41,opt,name=id,proto3" json:"id,omitempty"`
	OwnerAddress            []byte                             `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Name                    []byte                             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbr                    []byte                             `protobuf:"bytes,3,opt,name=abbr,proto3" json:"abbr,omitempty"`
	TotalSupply             int64                              `protobuf:"varint,4,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	FrozenSupply            []*AssetIssueContract_FrozenSupply `protobuf:"bytes,5,rep,name=frozen_supply,json=frozenSupply,proto3" json:"frozen_supply,omitempty"`
	TrxNum                  int32                              `protobuf:"varint,6,opt,name=trx_num,json=trxNum,proto3" json:"trx_num,omitempty"`
	Precision               int32                              `protobuf:"varint,7,opt,name=precision,proto3" json:"precision,omitempty"`
	Num                     int32                              `protobuf:"varint,8,opt,name=num,proto3" json:"num,omitempty"`
	StartTime               int64                              `protobuf:"varint,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime                 int64                              `protobuf:"varint,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Order                   int64                              `protobuf:"varint,11,opt,name=order,proto3" json:"order,omitempty"` // useless
	VoteScore               int32                              `protobuf:"varint,16,opt,name=vote_score,json=voteScore,proto3" json:"vote_score,omitempty"`
	Description             []byte                             `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	Url                     []byte                             `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`
	FreeAssetNetLimit       int64                              `protobuf:"varint,22,opt,name=free_asset_net_limit,json=freeAssetNetLimit,proto3" json:"free_asset_net_limit,omitempty"`
	PublicFreeAssetNetLimit int64                              `protobuf:"varint,23,opt,name=public_free_asset_net_limit,json=publicFreeAssetNetLimit,proto3" json:"public_free_asset_net_limit,omitempty"`
	PublicFreeAssetNetUsage int64                              `protobuf:"varint,24,opt,name=public_free_asset_net_usage,json=publicFreeAssetNetUsage,proto3" json:"public_free_asset_net_usage,omitempty"`
	PublicLatestFreeNetTime int64                              `protobuf:"varint,25,opt,name=public_latest_free_net_time,json=publicLatestFreeNetTime,proto3" json:"public_latest_free_net_time,omitempty"`
}

func (x *AssetIssueContract) Reset() {
	*x = AssetIssueContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_asset_issue_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetIssueContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetIssueContract) ProtoMessage() {}

func (x *AssetIssueContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetIssueContract.ProtoReflect.Descriptor instead.
func (*AssetIssueContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{0}
}

func (x *AssetIssueContract) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssetIssueContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *AssetIssueContract) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *AssetIssueContract) GetAbbr() []byte {
	if x != nil {
		return x.Abbr
	}
	return nil
}

func (x *AssetIssueContract) GetTotalSupply() int64 {
	if x != nil {
		return x.TotalSupply
	}
	return 0
}

func (x *AssetIssueContract) GetFrozenSupply() []*AssetIssueContract_FrozenSupply {
	if x != nil {
		return x.FrozenSupply
	}
	return nil
}

func (x *AssetIssueContract) GetTrxNum() int32 {
	if x != nil {
		return x.TrxNum
	}
	return 0
}

func (x *AssetIssueContract) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *AssetIssueContract) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *AssetIssueContract) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *AssetIssueContract) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *AssetIssueContract) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *AssetIssueContract) GetVoteScore() int32 {
	if x != nil {
		return x.VoteScore
	}
	return 0
}

func (x *AssetIssueContract) GetDescription() []byte {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *AssetIssueContract) GetUrl() []byte {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *AssetIssueContract) GetFreeAssetNetLimit() int64 {
	if x != nil {
		return x.FreeAssetNetLimit
	}
	return 0
}

func (x *AssetIssueContract) GetPublicFreeAssetNetLimit() int64 {
	if x != nil {
		return x.PublicFreeAssetNetLimit
	}
	return 0
}

func (x *AssetIssueContract) GetPublicFreeAssetNetUsage() int64 {
	if x != nil {
		return x.PublicFreeAssetNetUsage
	}
	return 0
}

func (x *AssetIssueContract) GetPublicLatestFreeNetTime() int64 {
	if x != nil {
		return x.PublicLatestFreeNetTime
	}
	return 0
}

type TransferAssetContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetName    []byte `protobuf:"bytes,1,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"` // this field is token name before the proposal ALLOW_SAME_TOKEN_NAME is active, otherwise it is token id and token is should be in string format.
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferAssetContract) Reset() {
	*x = TransferAssetContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_asset_issue_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferAssetContract) ProtoMessage() {}

func (x *TransferAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferAssetContract.ProtoReflect.Descriptor instead.
func (*TransferAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{1}
}

func (x *TransferAssetContract) GetAssetName() []byte {
	if x != nil {
		return x.AssetName
	}
	return nil
}

func (x *TransferAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *TransferAssetContract) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *TransferAssetContract) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type UnfreezeAssetContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (x *UnfreezeAssetContract) Reset() {
	*x = UnfreezeAssetContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_asset_issue_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfreezeAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfreezeAssetContract) ProtoMessage() {}

func (x *UnfreezeAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfreezeAssetContract.ProtoReflect.Descriptor instead.
func (*UnfreezeAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{2}
}

func (x *UnfreezeAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

type UpdateAssetContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress   []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Description    []byte `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url            []byte `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	NewLimit       int64  `protobuf:"varint,4,opt,name=new_limit,json=newLimit,proto3" json:"new_limit,omitempty"`
	NewPublicLimit int64  `protobuf:"varint,5,opt,name=new_public_limit,json=newPublicLimit,proto3" json:"new_public_limit,omitempty"`
}

func (x *UpdateAssetContract) Reset() {
	*x = UpdateAssetContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_asset_issue_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssetContract) ProtoMessage() {}

func (x *UpdateAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssetContract.ProtoReflect.Descriptor instead.
func (*UpdateAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *UpdateAssetContract) GetDescription() []byte {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdateAssetContract) GetUrl() []byte {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *UpdateAssetContract) GetNewLimit() int64 {
	if x != nil {
		return x.NewLimit
	}
	return 0
}

func (x *UpdateAssetContract) GetNewPublicLimit() int64 {
	if x != nil {
		return x.NewPublicLimit
	}
	return 0
}

type ParticipateAssetIssueContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	AssetName    []byte `protobuf:"bytes,3,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"` // this field is token name before the proposal ALLOW_SAME_TOKEN_NAME is active, otherwise it is token id and token is should be in string format.
	Amount       int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`                       // the amount of drops
}

func (x *ParticipateAssetIssueContract) Reset() {
	*x = ParticipateAssetIssueContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_asset_issue_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipateAssetIssueContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateAssetIssueContract) ProtoMessage() {}

func (x *ParticipateAssetIssueContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateAssetIssueContract.ProtoReflect.Descriptor instead.
func (*ParticipateAssetIssueContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{4}
}

func (x *ParticipateAssetIssueContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ParticipateAssetIssueContract) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *ParticipateAssetIssueContract) GetAssetName() []byte {
	if x != nil {
		return x.AssetName
	}
	return nil
}

func (x *ParticipateAssetIssueContract) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AssetIssueContract_FrozenSupply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrozenAmount int64 `protobuf:"varint,1,opt,name=frozen_amount,json=frozenAmount,proto3" json:"frozen_amount,omitempty"`
	FrozenDays   int64 `protobuf:"varint,2,opt,name=frozen_days,json=frozenDays,proto3" json:"frozen_days,omitempty"`
}

func (x *AssetIssueContract_FrozenSupply) Reset() {
	*x = AssetIssueContract_FrozenSupply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_asset_issue_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetIssueContract_FrozenSupply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetIssueContract_FrozenSupply) ProtoMessage() {}

func (x *AssetIssueContract_FrozenSupply) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetIssueContract_FrozenSupply.ProtoReflect.Descriptor instead.
func (*AssetIssueContract_FrozenSupply) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AssetIssueContract_FrozenSupply) GetFrozenAmount() int64 {
	if x != nil {
		return x.FrozenAmount
	}
	return 0
}

func (x *AssetIssueContract_FrozenSupply) GetFrozenDays() int64 {
	if x != nil {
		return x.FrozenDays
	}
	return 0
}

var File_core_contract_asset_issue_contract_proto protoreflect.FileDescriptor

var file_core_contract_asset_issue_contract_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x91, 0x06, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x62, 0x62, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x61, 0x62, 0x62, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0d, 0x66,
	0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x0c, 0x66,
	0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x72, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x72,
	0x78, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x14, 0x66, 0x72, 0x65, 0x65, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x72, 0x65, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x4e, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3c, 0x0a, 0x1b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x65,
	0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x46, 0x72, 0x65, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3c, 0x0a, 0x1b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x46, 0x72, 0x65, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x1b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x46, 0x72, 0x65, 0x65, 0x4e, 0x65, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x1a, 0x54, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72,
	0x6f, 0x7a, 0x65, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a,
	0x15, 0x55, 0x6e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x65, 0x77, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6e, 0x65, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x1d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x45, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x6f, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_contract_asset_issue_contract_proto_rawDescOnce sync.Once
	file_core_contract_asset_issue_contract_proto_rawDescData = file_core_contract_asset_issue_contract_proto_rawDesc
)

func file_core_contract_asset_issue_contract_proto_rawDescGZIP() []byte {
	file_core_contract_asset_issue_contract_proto_rawDescOnce.Do(func() {
		file_core_contract_asset_issue_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_contract_asset_issue_contract_proto_rawDescData)
	})
	return file_core_contract_asset_issue_contract_proto_rawDescData
}

var file_core_contract_asset_issue_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_core_contract_asset_issue_contract_proto_goTypes = []interface{}{
	(*AssetIssueContract)(nil),              // 0: protocol.AssetIssueContract
	(*TransferAssetContract)(nil),           // 1: protocol.TransferAssetContract
	(*UnfreezeAssetContract)(nil),           // 2: protocol.UnfreezeAssetContract
	(*UpdateAssetContract)(nil),             // 3: protocol.UpdateAssetContract
	(*ParticipateAssetIssueContract)(nil),   // 4: protocol.ParticipateAssetIssueContract
	(*AssetIssueContract_FrozenSupply)(nil), // 5: protocol.AssetIssueContract.FrozenSupply
}
var file_core_contract_asset_issue_contract_proto_depIdxs = []int32{
	5, // 0: protocol.AssetIssueContract.frozen_supply:type_name -> protocol.AssetIssueContract.FrozenSupply
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_contract_asset_issue_contract_proto_init() }
func file_core_contract_asset_issue_contract_proto_init() {
	if File_core_contract_asset_issue_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_contract_asset_issue_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetIssueContract); i {
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
		file_core_contract_asset_issue_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferAssetContract); i {
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
		file_core_contract_asset_issue_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfreezeAssetContract); i {
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
		file_core_contract_asset_issue_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAssetContract); i {
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
		file_core_contract_asset_issue_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipateAssetIssueContract); i {
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
		file_core_contract_asset_issue_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetIssueContract_FrozenSupply); i {
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
			RawDescriptor: file_core_contract_asset_issue_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_contract_asset_issue_contract_proto_goTypes,
		DependencyIndexes: file_core_contract_asset_issue_contract_proto_depIdxs,
		MessageInfos:      file_core_contract_asset_issue_contract_proto_msgTypes,
	}.Build()
	File_core_contract_asset_issue_contract_proto = out.File
	file_core_contract_asset_issue_contract_proto_rawDesc = nil
	file_core_contract_asset_issue_contract_proto_goTypes = nil
	file_core_contract_asset_issue_contract_proto_depIdxs = nil
}
