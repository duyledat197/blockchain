// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.14.0
// source: contract/contract_reader.proto

package contract

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	common "openmyth/blockchain/idl/pb/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetListApprovalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.Approval `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetListApprovalResponse) Reset() {
	*x = GetListApprovalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListApprovalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListApprovalResponse) ProtoMessage() {}

func (x *GetListApprovalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListApprovalResponse.ProtoReflect.Descriptor instead.
func (*GetListApprovalResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{0}
}

func (x *GetListApprovalResponse) GetData() []*common.Approval {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetListTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.Transfer `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetListTransferResponse) Reset() {
	*x = GetListTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListTransferResponse) ProtoMessage() {}

func (x *GetListTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListTransferResponse.ProtoReflect.Descriptor instead.
func (*GetListTransferResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{1}
}

func (x *GetListTransferResponse) GetData() []*common.Transfer {
	if x != nil {
		return x.Data
	}
	return nil
}

type RetrieveLatestBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number    uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Nonce     uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Hash      string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	GasLimit  uint64 `protobuf:"varint,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed   uint64 `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *RetrieveLatestBlockResponse) Reset() {
	*x = RetrieveLatestBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveLatestBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveLatestBlockResponse) ProtoMessage() {}

func (x *RetrieveLatestBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveLatestBlockResponse.ProtoReflect.Descriptor instead.
func (*RetrieveLatestBlockResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveLatestBlockResponse) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *RetrieveLatestBlockResponse) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *RetrieveLatestBlockResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *RetrieveLatestBlockResponse) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *RetrieveLatestBlockResponse) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *RetrieveLatestBlockResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type RetrieveBalanceOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetrieveBalanceOfRequest) Reset() {
	*x = RetrieveBalanceOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBalanceOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBalanceOfRequest) ProtoMessage() {}

func (x *RetrieveBalanceOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBalanceOfRequest.ProtoReflect.Descriptor instead.
func (*RetrieveBalanceOfRequest) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{3}
}

type RetrieveBalanceOfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance       uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	NativeBalance uint64 `protobuf:"varint,2,opt,name=native_balance,json=nativeBalance,proto3" json:"native_balance,omitempty"`
}

func (x *RetrieveBalanceOfResponse) Reset() {
	*x = RetrieveBalanceOfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBalanceOfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBalanceOfResponse) ProtoMessage() {}

func (x *RetrieveBalanceOfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBalanceOfResponse.ProtoReflect.Descriptor instead.
func (*RetrieveBalanceOfResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveBalanceOfResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *RetrieveBalanceOfResponse) GetNativeBalance() uint64 {
	if x != nil {
		return x.NativeBalance
	}
	return 0
}

type SendTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivKey string `protobuf:"bytes,1,opt,name=priv_key,json=privKey,proto3" json:"priv_key,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount  string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SendTransactionRequest) Reset() {
	*x = SendTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionRequest) ProtoMessage() {}

func (x *SendTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionRequest.ProtoReflect.Descriptor instead.
func (*SendTransactionRequest) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{5}
}

func (x *SendTransactionRequest) GetPrivKey() string {
	if x != nil {
		return x.PrivKey
	}
	return ""
}

func (x *SendTransactionRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendTransactionRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type SendTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendTransactionResponse) Reset() {
	*x = SendTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionResponse) ProtoMessage() {}

func (x *SendTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionResponse.ProtoReflect.Descriptor instead.
func (*SendTransactionResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{6}
}

type SendTransactionV2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SendTransactionV2Request) Reset() {
	*x = SendTransactionV2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_reader_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionV2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionV2Request) ProtoMessage() {}

func (x *SendTransactionV2Request) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_reader_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionV2Request.ProtoReflect.Descriptor instead.
func (*SendTransactionV2Request) Descriptor() ([]byte, []int) {
	return file_contract_contract_reader_proto_rawDescGZIP(), []int{7}
}

func (x *SendTransactionV2Request) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SendTransactionV2Request) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendTransactionV2Request) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_contract_contract_reader_proto protoreflect.FileDescriptor

var file_contract_contract_reader_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb5, 0x01, 0x0a,
	0x1b, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67,
	0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x5c, 0x0a, 0x19, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5b,
	0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x76,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76,
	0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xae, 0x05, 0x0a, 0x15, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x63, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x63, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x13,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x72, 0x0a,
	0x11, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x4f, 0x66, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x70, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x73, 0x12, 0x74, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x32, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x6f, 0x70, 0x65,
	0x6e, 0x6d, 0x79, 0x74, 0x68, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_contract_reader_proto_rawDescOnce sync.Once
	file_contract_contract_reader_proto_rawDescData = file_contract_contract_reader_proto_rawDesc
)

func file_contract_contract_reader_proto_rawDescGZIP() []byte {
	file_contract_contract_reader_proto_rawDescOnce.Do(func() {
		file_contract_contract_reader_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_contract_reader_proto_rawDescData)
	})
	return file_contract_contract_reader_proto_rawDescData
}

var file_contract_contract_reader_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_contract_contract_reader_proto_goTypes = []interface{}{
	(*GetListApprovalResponse)(nil),     // 0: contract.GetListApprovalResponse
	(*GetListTransferResponse)(nil),     // 1: contract.GetListTransferResponse
	(*RetrieveLatestBlockResponse)(nil), // 2: contract.RetrieveLatestBlockResponse
	(*RetrieveBalanceOfRequest)(nil),    // 3: contract.RetrieveBalanceOfRequest
	(*RetrieveBalanceOfResponse)(nil),   // 4: contract.RetrieveBalanceOfResponse
	(*SendTransactionRequest)(nil),      // 5: contract.SendTransactionRequest
	(*SendTransactionResponse)(nil),     // 6: contract.SendTransactionResponse
	(*SendTransactionV2Request)(nil),    // 7: contract.SendTransactionV2Request
	(*common.Approval)(nil),             // 8: common.Approval
	(*common.Transfer)(nil),             // 9: common.Transfer
	(*emptypb.Empty)(nil),               // 10: google.protobuf.Empty
}
var file_contract_contract_reader_proto_depIdxs = []int32{
	8,  // 0: contract.GetListApprovalResponse.data:type_name -> common.Approval
	9,  // 1: contract.GetListTransferResponse.data:type_name -> common.Transfer
	10, // 2: contract.ContractReaderService.GetListApproval:input_type -> google.protobuf.Empty
	10, // 3: contract.ContractReaderService.GetListTransfer:input_type -> google.protobuf.Empty
	10, // 4: contract.ContractReaderService.RetrieveLatestBlock:input_type -> google.protobuf.Empty
	3,  // 5: contract.ContractReaderService.RetrieveBalanceOf:input_type -> contract.RetrieveBalanceOfRequest
	5,  // 6: contract.ContractReaderService.SendTransaction:input_type -> contract.SendTransactionRequest
	7,  // 7: contract.ContractReaderService.SendTransactionV2:input_type -> contract.SendTransactionV2Request
	0,  // 8: contract.ContractReaderService.GetListApproval:output_type -> contract.GetListApprovalResponse
	1,  // 9: contract.ContractReaderService.GetListTransfer:output_type -> contract.GetListTransferResponse
	2,  // 10: contract.ContractReaderService.RetrieveLatestBlock:output_type -> contract.RetrieveLatestBlockResponse
	4,  // 11: contract.ContractReaderService.RetrieveBalanceOf:output_type -> contract.RetrieveBalanceOfResponse
	6,  // 12: contract.ContractReaderService.SendTransaction:output_type -> contract.SendTransactionResponse
	6,  // 13: contract.ContractReaderService.SendTransactionV2:output_type -> contract.SendTransactionResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_contract_contract_reader_proto_init() }
func file_contract_contract_reader_proto_init() {
	if File_contract_contract_reader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_contract_reader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListApprovalResponse); i {
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
		file_contract_contract_reader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListTransferResponse); i {
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
		file_contract_contract_reader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveLatestBlockResponse); i {
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
		file_contract_contract_reader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveBalanceOfRequest); i {
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
		file_contract_contract_reader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveBalanceOfResponse); i {
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
		file_contract_contract_reader_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransactionRequest); i {
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
		file_contract_contract_reader_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransactionResponse); i {
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
		file_contract_contract_reader_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransactionV2Request); i {
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
			RawDescriptor: file_contract_contract_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_contract_reader_proto_goTypes,
		DependencyIndexes: file_contract_contract_reader_proto_depIdxs,
		MessageInfos:      file_contract_contract_reader_proto_msgTypes,
	}.Build()
	File_contract_contract_reader_proto = out.File
	file_contract_contract_reader_proto_rawDesc = nil
	file_contract_contract_reader_proto_goTypes = nil
	file_contract_contract_reader_proto_depIdxs = nil
}
