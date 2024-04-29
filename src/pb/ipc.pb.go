// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.26.1
// source: ipc.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_STATUS_UNSPECIFIED   Status = 0
	Status_STATUS_STARTING      Status = 1
	Status_STATUS_RUNNING       Status = 2
	Status_STATUS_SCANNING      Status = 3
	Status_STATUS_SHUTTING_DOWN Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_STARTING",
		2: "STATUS_RUNNING",
		3: "STATUS_SCANNING",
		4: "STATUS_SHUTTING_DOWN",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":   0,
		"STATUS_STARTING":      1,
		"STATUS_RUNNING":       2,
		"STATUS_SCANNING":      3,
		"STATUS_SHUTTING_DOWN": 4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_ipc_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_ipc_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{0}
}

type UTXOState int32

const (
	UTXOState_UNKNOWN     UTXOState = 0
	UTXOState_UNCONFIRMED UTXOState = 1
	UTXOState_UNSPENT     UTXOState = 2
	UTXOState_SPENT       UTXOState = 3
)

// Enum value maps for UTXOState.
var (
	UTXOState_name = map[int32]string{
		0: "UNKNOWN",
		1: "UNCONFIRMED",
		2: "UNSPENT",
		3: "SPENT",
	}
	UTXOState_value = map[string]int32{
		"UNKNOWN":     0,
		"UNCONFIRMED": 1,
		"UNSPENT":     2,
		"SPENT":       3,
	}
)

func (x UTXOState) Enum() *UTXOState {
	p := new(UTXOState)
	*p = x
	return p
}

func (x UTXOState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UTXOState) Descriptor() protoreflect.EnumDescriptor {
	return file_ipc_proto_enumTypes[1].Descriptor()
}

func (UTXOState) Type() protoreflect.EnumType {
	return &file_ipc_proto_enumTypes[1]
}

func (x UTXOState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UTXOState.Descriptor instead.
func (UTXOState) EnumDescriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{1}
}

// This is an empty message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{0}
}

// The response message containing the greetings
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=ipc.Status" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

type UTXOCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Utxos []*OwnedUTXO `protobuf:"bytes,1,rep,name=utxos,proto3" json:"utxos,omitempty"` // Array of OwnedUTXO
}

func (x *UTXOCollection) Reset() {
	*x = UTXOCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UTXOCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UTXOCollection) ProtoMessage() {}

func (x *UTXOCollection) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UTXOCollection.ProtoReflect.Descriptor instead.
func (*UTXOCollection) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{2}
}

func (x *UTXOCollection) GetUtxos() []*OwnedUTXO {
	if x != nil {
		return x.Utxos
	}
	return nil
}

type PasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PasswordRequest) Reset() {
	*x = PasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRequest) ProtoMessage() {}

func (x *PasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRequest.ProtoReflect.Descriptor instead.
func (*PasswordRequest) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{3}
}

func (x *PasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"` // only filled if success false
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{4}
}

func (x *BoolResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BoolResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type OwnedUTXO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid               []byte                 `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Vout               uint32                 `protobuf:"varint,2,opt,name=vout,proto3" json:"vout,omitempty"`
	Amount             uint64                 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	PubKey             []byte                 `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	TimestampConfirmed *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp_confirmed,json=timestampConfirmed,proto3" json:"timestamp_confirmed,omitempty"`
	UtxoState          UTXOState              `protobuf:"varint,6,opt,name=utxo_state,json=utxoState,proto3,enum=ipc.UTXOState" json:"utxo_state,omitempty"`
	Label              []byte                 `protobuf:"bytes,7,opt,name=label,proto3,oneof" json:"label,omitempty"`
}

func (x *OwnedUTXO) Reset() {
	*x = OwnedUTXO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnedUTXO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnedUTXO) ProtoMessage() {}

func (x *OwnedUTXO) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnedUTXO.ProtoReflect.Descriptor instead.
func (*OwnedUTXO) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{5}
}

func (x *OwnedUTXO) GetTxid() []byte {
	if x != nil {
		return x.Txid
	}
	return nil
}

func (x *OwnedUTXO) GetVout() uint32 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *OwnedUTXO) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OwnedUTXO) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *OwnedUTXO) GetTimestampConfirmed() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampConfirmed
	}
	return nil
}

func (x *OwnedUTXO) GetUtxoState() UTXOState {
	if x != nil {
		return x.UtxoState
	}
	return UTXOState_UNKNOWN
}

func (x *OwnedUTXO) GetLabel() []byte {
	if x != nil {
		return x.Label
	}
	return nil
}

type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipients []*TransactionRecipient `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	FeeRate    int64                   `protobuf:"varint,2,opt,name=feeRate,proto3" json:"feeRate,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{6}
}

func (x *CreateTransactionRequest) GetRecipients() []*TransactionRecipient {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *CreateTransactionRequest) GetFeeRate() int64 {
	if x != nil {
		return x.FeeRate
	}
	return 0
}

type TransactionRecipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount     uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Annotation string `protobuf:"bytes,3,opt,name=annotation,proto3" json:"annotation,omitempty"`
}

func (x *TransactionRecipient) Reset() {
	*x = TransactionRecipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRecipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRecipient) ProtoMessage() {}

func (x *TransactionRecipient) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRecipient.ProtoReflect.Descriptor instead.
func (*TransactionRecipient) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionRecipient) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TransactionRecipient) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionRecipient) GetAnnotation() string {
	if x != nil {
		return x.Annotation
	}
	return ""
}

type RawTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawTx []byte `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
}

func (x *RawTransaction) Reset() {
	*x = RawTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawTransaction) ProtoMessage() {}

func (x *RawTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawTransaction.ProtoReflect.Descriptor instead.
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return file_ipc_proto_rawDescGZIP(), []int{8}
}

func (x *RawTransaction) GetRawTx() []byte {
	if x != nil {
		return x.RawTx
	}
	return nil
}

var File_ipc_proto protoreflect.FileDescriptor

var file_ipc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x70, 0x63,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x69,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x36, 0x0a, 0x0e, 0x55, 0x54, 0x58, 0x4f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x75, 0x74, 0x78, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x55, 0x54,
	0x58, 0x4f, 0x52, 0x05, 0x75, 0x74, 0x78, 0x6f, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x85, 0x02, 0x0a, 0x09, 0x4f, 0x77, 0x6e,
	0x65, 0x64, 0x55, 0x54, 0x58, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x4b, 0x0a, 0x13, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x0a,
	0x75, 0x74, 0x78, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x55, 0x54, 0x58, 0x4f, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x09, 0x75, 0x74, 0x78, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x22, 0x6f, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x22, 0x68, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x0e, 0x52,
	0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a,
	0x06, 0x72, 0x61, 0x77, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72,
	0x61, 0x77, 0x54, 0x78, 0x2a, 0x78, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x48, 0x55, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x2a, 0x41,
	0x0a, 0x09, 0x55, 0x54, 0x58, 0x4f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x45, 0x4e, 0x54, 0x10,
	0x03, 0x32, 0xc4, 0x02, 0x0a, 0x0a, 0x49, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x2e, 0x69, 0x70, 0x63,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x69, 0x70,
	0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x0a, 0x2e, 0x69, 0x70, 0x63,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x53, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x69, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x54, 0x58, 0x4f, 0x73, 0x12, 0x0a,
	0x2e, 0x69, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x69, 0x70, 0x63,
	0x2e, 0x55, 0x54, 0x58, 0x4f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x47, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ipc_proto_rawDescOnce sync.Once
	file_ipc_proto_rawDescData = file_ipc_proto_rawDesc
)

func file_ipc_proto_rawDescGZIP() []byte {
	file_ipc_proto_rawDescOnce.Do(func() {
		file_ipc_proto_rawDescData = protoimpl.X.CompressGZIP(file_ipc_proto_rawDescData)
	})
	return file_ipc_proto_rawDescData
}

var file_ipc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ipc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ipc_proto_goTypes = []interface{}{
	(Status)(0),                      // 0: ipc.Status
	(UTXOState)(0),                   // 1: ipc.UTXOState
	(*Empty)(nil),                    // 2: ipc.Empty
	(*StatusResponse)(nil),           // 3: ipc.StatusResponse
	(*UTXOCollection)(nil),           // 4: ipc.UTXOCollection
	(*PasswordRequest)(nil),          // 5: ipc.PasswordRequest
	(*BoolResponse)(nil),             // 6: ipc.BoolResponse
	(*OwnedUTXO)(nil),                // 7: ipc.OwnedUTXO
	(*CreateTransactionRequest)(nil), // 8: ipc.CreateTransactionRequest
	(*TransactionRecipient)(nil),     // 9: ipc.TransactionRecipient
	(*RawTransaction)(nil),           // 10: ipc.RawTransaction
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_ipc_proto_depIdxs = []int32{
	0,  // 0: ipc.StatusResponse.status:type_name -> ipc.Status
	7,  // 1: ipc.UTXOCollection.utxos:type_name -> ipc.OwnedUTXO
	11, // 2: ipc.OwnedUTXO.timestamp_confirmed:type_name -> google.protobuf.Timestamp
	1,  // 3: ipc.OwnedUTXO.utxo_state:type_name -> ipc.UTXOState
	9,  // 4: ipc.CreateTransactionRequest.recipients:type_name -> ipc.TransactionRecipient
	2,  // 5: ipc.IpcService.Status:input_type -> ipc.Empty
	5,  // 6: ipc.IpcService.Unlock:input_type -> ipc.PasswordRequest
	2,  // 7: ipc.IpcService.Shutdown:input_type -> ipc.Empty
	5,  // 8: ipc.IpcService.SetPassword:input_type -> ipc.PasswordRequest
	2,  // 9: ipc.IpcService.ListUTXOs:input_type -> ipc.Empty
	8,  // 10: ipc.IpcService.CreateTransaction:input_type -> ipc.CreateTransactionRequest
	3,  // 11: ipc.IpcService.Status:output_type -> ipc.StatusResponse
	6,  // 12: ipc.IpcService.Unlock:output_type -> ipc.BoolResponse
	6,  // 13: ipc.IpcService.Shutdown:output_type -> ipc.BoolResponse
	6,  // 14: ipc.IpcService.SetPassword:output_type -> ipc.BoolResponse
	4,  // 15: ipc.IpcService.ListUTXOs:output_type -> ipc.UTXOCollection
	10, // 16: ipc.IpcService.CreateTransaction:output_type -> ipc.RawTransaction
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ipc_proto_init() }
func file_ipc_proto_init() {
	if File_ipc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ipc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_ipc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_ipc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UTXOCollection); i {
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
		file_ipc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordRequest); i {
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
		file_ipc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
		file_ipc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnedUTXO); i {
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
		file_ipc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionRequest); i {
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
		file_ipc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRecipient); i {
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
		file_ipc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawTransaction); i {
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
	file_ipc_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ipc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ipc_proto_goTypes,
		DependencyIndexes: file_ipc_proto_depIdxs,
		EnumInfos:         file_ipc_proto_enumTypes,
		MessageInfos:      file_ipc_proto_msgTypes,
	}.Build()
	File_ipc_proto = out.File
	file_ipc_proto_rawDesc = nil
	file_ipc_proto_goTypes = nil
	file_ipc_proto_depIdxs = nil
}
