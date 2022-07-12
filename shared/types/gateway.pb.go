// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gateway.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetRecentBlocksRequest struct {
	FromBlock            uint64   `protobuf:"varint,1,opt,name=from_block,json=fromBlock,proto3" json:"from_block,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecentBlocksRequest) Reset()         { *m = GetRecentBlocksRequest{} }
func (m *GetRecentBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecentBlocksRequest) ProtoMessage()    {}
func (*GetRecentBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}
func (m *GetRecentBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentBlocksRequest.Unmarshal(m, b)
}
func (m *GetRecentBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentBlocksRequest.Marshal(b, m, deterministic)
}
func (m *GetRecentBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentBlocksRequest.Merge(m, src)
}
func (m *GetRecentBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecentBlocksRequest.Size(m)
}
func (m *GetRecentBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentBlocksRequest proto.InternalMessageInfo

func (m *GetRecentBlocksRequest) GetFromBlock() uint64 {
	if m != nil {
		return m.FromBlock
	}
	return 0
}

func (m *GetRecentBlocksRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetRecentBlocksReply struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecentBlocksReply) Reset()         { *m = GetRecentBlocksReply{} }
func (m *GetRecentBlocksReply) String() string { return proto.CompactTextString(m) }
func (*GetRecentBlocksReply) ProtoMessage()    {}
func (*GetRecentBlocksReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}
func (m *GetRecentBlocksReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentBlocksReply.Unmarshal(m, b)
}
func (m *GetRecentBlocksReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentBlocksReply.Marshal(b, m, deterministic)
}
func (m *GetRecentBlocksReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentBlocksReply.Merge(m, src)
}
func (m *GetRecentBlocksReply) XXX_Size() int {
	return xxx_messageInfo_GetRecentBlocksReply.Size(m)
}
func (m *GetRecentBlocksReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentBlocksReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentBlocksReply proto.InternalMessageInfo

func (m *GetRecentBlocksReply) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type GetRecentTxsRequest struct {
	FromTs               uint64   `protobuf:"varint,1,opt,name=from_ts,json=fromTs,proto3" json:"from_ts,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecentTxsRequest) Reset()         { *m = GetRecentTxsRequest{} }
func (m *GetRecentTxsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecentTxsRequest) ProtoMessage()    {}
func (*GetRecentTxsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}
func (m *GetRecentTxsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentTxsRequest.Unmarshal(m, b)
}
func (m *GetRecentTxsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentTxsRequest.Marshal(b, m, deterministic)
}
func (m *GetRecentTxsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentTxsRequest.Merge(m, src)
}
func (m *GetRecentTxsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecentTxsRequest.Size(m)
}
func (m *GetRecentTxsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentTxsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentTxsRequest proto.InternalMessageInfo

func (m *GetRecentTxsRequest) GetFromTs() uint64 {
	if m != nil {
		return m.FromTs
	}
	return 0
}

func (m *GetRecentTxsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetRecentTxsReply struct {
	Txs                  []*Transaction `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRecentTxsReply) Reset()         { *m = GetRecentTxsReply{} }
func (m *GetRecentTxsReply) String() string { return proto.CompactTextString(m) }
func (*GetRecentTxsReply) ProtoMessage()    {}
func (*GetRecentTxsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{3}
}
func (m *GetRecentTxsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentTxsReply.Unmarshal(m, b)
}
func (m *GetRecentTxsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentTxsReply.Marshal(b, m, deterministic)
}
func (m *GetRecentTxsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentTxsReply.Merge(m, src)
}
func (m *GetRecentTxsReply) XXX_Size() int {
	return xxx_messageInfo_GetRecentTxsReply.Size(m)
}
func (m *GetRecentTxsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentTxsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentTxsReply proto.InternalMessageInfo

func (m *GetRecentTxsReply) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type GetBlockByHashOrNumberRequest struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockByHashOrNumberRequest) Reset()         { *m = GetBlockByHashOrNumberRequest{} }
func (m *GetBlockByHashOrNumberRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashOrNumberRequest) ProtoMessage()    {}
func (*GetBlockByHashOrNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{4}
}
func (m *GetBlockByHashOrNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashOrNumberRequest.Unmarshal(m, b)
}
func (m *GetBlockByHashOrNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashOrNumberRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockByHashOrNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashOrNumberRequest.Merge(m, src)
}
func (m *GetBlockByHashOrNumberRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashOrNumberRequest.Size(m)
}
func (m *GetBlockByHashOrNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashOrNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashOrNumberRequest proto.InternalMessageInfo

func (m *GetBlockByHashOrNumberRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GetBlockByHashOrNumberRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockByHashOrNumberReply struct {
	Block                *BlockDetails `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetBlockByHashOrNumberReply) Reset()         { *m = GetBlockByHashOrNumberReply{} }
func (m *GetBlockByHashOrNumberReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashOrNumberReply) ProtoMessage()    {}
func (*GetBlockByHashOrNumberReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{5}
}
func (m *GetBlockByHashOrNumberReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashOrNumberReply.Unmarshal(m, b)
}
func (m *GetBlockByHashOrNumberReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashOrNumberReply.Marshal(b, m, deterministic)
}
func (m *GetBlockByHashOrNumberReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashOrNumberReply.Merge(m, src)
}
func (m *GetBlockByHashOrNumberReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashOrNumberReply.Size(m)
}
func (m *GetBlockByHashOrNumberReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashOrNumberReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashOrNumberReply proto.InternalMessageInfo

func (m *GetBlockByHashOrNumberReply) GetBlock() *BlockDetails {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetTransactionByHashRequest struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionByHashRequest) Reset()         { *m = GetTransactionByHashRequest{} }
func (m *GetTransactionByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionByHashRequest) ProtoMessage()    {}
func (*GetTransactionByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{6}
}
func (m *GetTransactionByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionByHashRequest.Unmarshal(m, b)
}
func (m *GetTransactionByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionByHashRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionByHashRequest.Merge(m, src)
}
func (m *GetTransactionByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionByHashRequest.Size(m)
}
func (m *GetTransactionByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionByHashRequest proto.InternalMessageInfo

func (m *GetTransactionByHashRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type GetTransactionByHashReply struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetTransactionByHashReply) Reset()         { *m = GetTransactionByHashReply{} }
func (m *GetTransactionByHashReply) String() string { return proto.CompactTextString(m) }
func (*GetTransactionByHashReply) ProtoMessage()    {}
func (*GetTransactionByHashReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{7}
}
func (m *GetTransactionByHashReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionByHashReply.Unmarshal(m, b)
}
func (m *GetTransactionByHashReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionByHashReply.Marshal(b, m, deterministic)
}
func (m *GetTransactionByHashReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionByHashReply.Merge(m, src)
}
func (m *GetTransactionByHashReply) XXX_Size() int {
	return xxx_messageInfo_GetTransactionByHashReply.Size(m)
}
func (m *GetTransactionByHashReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionByHashReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionByHashReply proto.InternalMessageInfo

func (m *GetTransactionByHashReply) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type GetTokenTransfersRequest struct {
	TokenContract        []byte   `protobuf:"bytes,1,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
	FromBlock            uint64   `protobuf:"varint,2,opt,name=from_block,json=fromBlock,proto3" json:"from_block,omitempty"`
	Limit                uint64   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenTransfersRequest) Reset()         { *m = GetTokenTransfersRequest{} }
func (m *GetTokenTransfersRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenTransfersRequest) ProtoMessage()    {}
func (*GetTokenTransfersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{8}
}
func (m *GetTokenTransfersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenTransfersRequest.Unmarshal(m, b)
}
func (m *GetTokenTransfersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenTransfersRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenTransfersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenTransfersRequest.Merge(m, src)
}
func (m *GetTokenTransfersRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenTransfersRequest.Size(m)
}
func (m *GetTokenTransfersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenTransfersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenTransfersRequest proto.InternalMessageInfo

func (m *GetTokenTransfersRequest) GetTokenContract() []byte {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *GetTokenTransfersRequest) GetFromBlock() uint64 {
	if m != nil {
		return m.FromBlock
	}
	return 0
}

func (m *GetTokenTransfersRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetTokenTransfersReply struct {
	Transfers            []*TokenTransfer `protobuf:"bytes,1,rep,name=transfers,proto3" json:"transfers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetTokenTransfersReply) Reset()         { *m = GetTokenTransfersReply{} }
func (m *GetTokenTransfersReply) String() string { return proto.CompactTextString(m) }
func (*GetTokenTransfersReply) ProtoMessage()    {}
func (*GetTokenTransfersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{9}
}
func (m *GetTokenTransfersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenTransfersReply.Unmarshal(m, b)
}
func (m *GetTokenTransfersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenTransfersReply.Marshal(b, m, deterministic)
}
func (m *GetTokenTransfersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenTransfersReply.Merge(m, src)
}
func (m *GetTokenTransfersReply) XXX_Size() int {
	return xxx_messageInfo_GetTokenTransfersReply.Size(m)
}
func (m *GetTokenTransfersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenTransfersReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenTransfersReply proto.InternalMessageInfo

func (m *GetTokenTransfersReply) GetTransfers() []*TokenTransfer {
	if m != nil {
		return m.Transfers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRecentBlocksRequest)(nil), "com.ankr.GetRecentBlocksRequest")
	proto.RegisterType((*GetRecentBlocksReply)(nil), "com.ankr.GetRecentBlocksReply")
	proto.RegisterType((*GetRecentTxsRequest)(nil), "com.ankr.GetRecentTxsRequest")
	proto.RegisterType((*GetRecentTxsReply)(nil), "com.ankr.GetRecentTxsReply")
	proto.RegisterType((*GetBlockByHashOrNumberRequest)(nil), "com.ankr.GetBlockByHashOrNumberRequest")
	proto.RegisterType((*GetBlockByHashOrNumberReply)(nil), "com.ankr.GetBlockByHashOrNumberReply")
	proto.RegisterType((*GetTransactionByHashRequest)(nil), "com.ankr.GetTransactionByHashRequest")
	proto.RegisterType((*GetTransactionByHashReply)(nil), "com.ankr.GetTransactionByHashReply")
	proto.RegisterType((*GetTokenTransfersRequest)(nil), "com.ankr.GetTokenTransfersRequest")
	proto.RegisterType((*GetTokenTransfersReply)(nil), "com.ankr.GetTokenTransfersReply")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0x55, 0xfa, 0x93, 0xef, 0xeb, 0x4d, 0x7f, 0xa8, 0x1b, 0xd2, 0x74, 0x4a, 0x20, 0x1a, 0x54,
	0xb5, 0x0b, 0x94, 0x51, 0x83, 0x10, 0x0b, 0x10, 0x48, 0xa1, 0x52, 0x90, 0x2a, 0x28, 0x1a, 0x65,
	0xc5, 0x26, 0x72, 0x06, 0x37, 0x89, 0x32, 0x19, 0x0f, 0x63, 0xa7, 0x24, 0x8a, 0xca, 0x82, 0x1d,
	0x6b, 0x1e, 0x8d, 0x57, 0x60, 0xc5, 0x53, 0x20, 0x5f, 0xcf, 0x6f, 0x92, 0x69, 0x37, 0x51, 0xe6,
	0xfa, 0xf8, 0x9c, 0xe3, 0xeb, 0x73, 0x0d, 0x3b, 0x7d, 0x2a, 0xd9, 0x37, 0x3a, 0x6b, 0xf8, 0x01,
	0x97, 0x9c, 0xfc, 0xef, 0xf0, 0x71, 0x83, 0x7a, 0xa3, 0xc0, 0x28, 0xf5, 0x5c, 0xee, 0x8c, 0x74,
	0xd9, 0xd8, 0x97, 0x01, 0xf5, 0x04, 0x75, 0xe4, 0x90, 0x7b, 0x61, 0xa9, 0x2c, 0xf9, 0x88, 0x79,
	0x5d, 0x5c, 0xb8, 0x66, 0x41, 0x58, 0x7d, 0xd4, 0xe7, 0xbc, 0xef, 0x32, 0x8b, 0xfa, 0x43, 0x8b,
	0x7a, 0x1e, 0x97, 0x54, 0x6d, 0x11, 0x7a, 0xd5, 0xfc, 0x00, 0x95, 0x36, 0x93, 0x36, 0x73, 0x98,
	0x27, 0x5b, 0x8a, 0x5e, 0xd8, 0xec, 0xeb, 0x84, 0x09, 0x49, 0x6a, 0x00, 0xd7, 0x01, 0x1f, 0x77,
	0x51, 0xb4, 0x5a, 0xa8, 0x17, 0xce, 0x36, 0xec, 0x2d, 0x55, 0x41, 0x18, 0x29, 0xc3, 0xa6, 0x3b,
	0x1c, 0x0f, 0x65, 0x75, 0x0d, 0x57, 0xf4, 0x87, 0xf9, 0x16, 0xca, 0x4b, 0x74, 0xbe, 0x3b, 0x23,
	0xa7, 0x50, 0x44, 0x1e, 0x51, 0x2d, 0xd4, 0xd7, 0xcf, 0x4a, 0xcd, 0xbd, 0x46, 0x74, 0xaa, 0x06,
	0xc2, 0xec, 0x70, 0xd9, 0xbc, 0x80, 0x83, 0x98, 0xa0, 0x33, 0x8d, 0xcd, 0x1c, 0xc2, 0x7f, 0x68,
	0x46, 0x8a, 0xd0, 0x49, 0x51, 0x7d, 0x76, 0x44, 0x8e, 0x8d, 0xd7, 0xb0, 0x9f, 0x65, 0xd1, 0x1e,
	0xd6, 0xe5, 0x34, 0x32, 0xf0, 0x30, 0x31, 0xd0, 0x49, 0x1a, 0x69, 0x2b, 0x84, 0x79, 0x09, 0xb5,
	0x36, 0xd3, 0xf6, 0x5b, 0xb3, 0xf7, 0x54, 0x0c, 0xae, 0x82, 0x8f, 0x93, 0x71, 0x8f, 0x05, 0x91,
	0x9b, 0x0a, 0x14, 0x3d, 0x2c, 0x44, 0x66, 0xf4, 0x17, 0x21, 0xb0, 0x31, 0xa0, 0x62, 0x80, 0x5e,
	0xb6, 0x6c, 0xfc, 0x6f, 0x5e, 0xc2, 0x71, 0x1e, 0x99, 0x32, 0xf5, 0x0c, 0x36, 0x93, 0x06, 0x97,
	0x9a, 0x95, 0x85, 0xbe, 0x5c, 0x30, 0x49, 0x87, 0xae, 0xb0, 0x35, 0xc8, 0x3c, 0x47, 0xb2, 0x94,
	0x61, 0x4d, 0x19, 0xf9, 0x8a, 0xf4, 0x15, 0xd7, 0x76, 0xa8, 0xdf, 0x81, 0xa3, 0xd5, 0x5b, 0x94,
	0xfa, 0x4b, 0x28, 0xa5, 0x62, 0x14, 0x7a, 0xc8, 0x69, 0x4d, 0x1a, 0x69, 0xde, 0x40, 0x55, 0xb1,
	0xaa, 0xbc, 0x75, 0xc2, 0xb8, 0xc5, 0x77, 0x75, 0x02, 0xbb, 0x3a, 0x88, 0x0e, 0xf7, 0x64, 0x40,
	0x1d, 0x19, 0xfa, 0xd9, 0xc1, 0xea, 0xbb, 0xb0, 0xb8, 0x90, 0xaf, 0xb5, 0xdc, 0x7c, 0xad, 0xa7,
	0x2f, 0xf6, 0x0a, 0xe3, 0xba, 0xa8, 0xab, 0x8e, 0xf2, 0x02, 0xb6, 0xa2, 0xe0, 0x47, 0x77, 0x7c,
	0x98, 0x3a, 0x48, 0x7a, 0x87, 0x9d, 0x20, 0x9b, 0x7f, 0x37, 0x60, 0x5f, 0x07, 0xd5, 0xe1, 0x13,
	0xd9, 0xd6, 0x93, 0x47, 0x46, 0xb0, 0xb7, 0x10, 0x63, 0x52, 0x4f, 0xc8, 0x56, 0x0f, 0x8c, 0xf1,
	0xf8, 0x0e, 0x84, 0xef, 0xce, 0xcc, 0xca, 0x8f, 0xdf, 0x7f, 0x7e, 0xad, 0x3d, 0x20, 0xbb, 0xd6,
	0xcd, 0x39, 0x75, 0xfd, 0x01, 0xb5, 0xf0, 0xe8, 0xa4, 0x0b, 0xdb, 0xe9, 0xb0, 0x92, 0xda, 0x0a,
	0x9e, 0x64, 0x14, 0x8c, 0xe3, 0xbc, 0x65, 0xa5, 0x71, 0x80, 0x1a, 0x3b, 0xa4, 0x14, 0x6b, 0xc8,
	0x29, 0xf9, 0x59, 0xc0, 0xae, 0xad, 0xc8, 0x20, 0x39, 0xcd, 0x90, 0xe5, 0x47, 0xde, 0x38, 0xb9,
	0x1f, 0xa8, 0xf4, 0x9f, 0xa0, 0xfe, 0x11, 0x39, 0xcc, 0x9e, 0xd1, 0x9a, 0xeb, 0x09, 0xb9, 0x25,
	0xdf, 0xf1, 0x81, 0x58, 0x8a, 0x23, 0xc9, 0xf2, 0xe7, 0x25, 0xdc, 0x78, 0x7a, 0x1f, 0x4c, 0x99,
	0x30, 0xd0, 0x44, 0x99, 0x90, 0x54, 0x13, 0xac, 0xb9, 0x9a, 0x86, 0x5b, 0x32, 0xc7, 0x97, 0x21,
	0x1b, 0x20, 0x62, 0x66, 0x59, 0x57, 0xa5, 0xda, 0xa8, 0xdf, 0x89, 0x51, 0xb2, 0x75, 0x94, 0x35,
	0x48, 0x35, 0x91, 0x55, 0x28, 0x2b, 0x0e, 0x5b, 0xeb, 0x0d, 0xec, 0xc6, 0x24, 0xf8, 0xfc, 0xb6,
	0x2a, 0x4b, 0xd9, 0xfb, 0xa4, 0xea, 0x9f, 0xcb, 0x0d, 0x4b, 0x0c, 0x68, 0xc0, 0xbe, 0x58, 0x72,
	0xe6, 0x33, 0xf1, 0x0a, 0x7f, 0x7b, 0x45, 0xdc, 0xf4, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x28, 0xe4, 0x16, 0x53, 0x22, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockscoutGatewayClient is the client API for BlockscoutGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockscoutGatewayClient interface {
	GetRecentBlocks(ctx context.Context, in *GetRecentBlocksRequest, opts ...grpc.CallOption) (*GetRecentBlocksReply, error)
	GetRecentTxs(ctx context.Context, in *GetRecentTxsRequest, opts ...grpc.CallOption) (*GetRecentTxsReply, error)
	GetBlockByHashOrNumber(ctx context.Context, in *GetBlockByHashOrNumberRequest, opts ...grpc.CallOption) (*GetBlockByHashOrNumberReply, error)
	GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*GetTransactionByHashReply, error)
	GetTokenTransfers(ctx context.Context, in *GetTokenTransfersRequest, opts ...grpc.CallOption) (*GetTokenTransfersReply, error)
}

type blockscoutGatewayClient struct {
	cc *grpc.ClientConn
}

func NewBlockscoutGatewayClient(cc *grpc.ClientConn) BlockscoutGatewayClient {
	return &blockscoutGatewayClient{cc}
}

func (c *blockscoutGatewayClient) GetRecentBlocks(ctx context.Context, in *GetRecentBlocksRequest, opts ...grpc.CallOption) (*GetRecentBlocksReply, error) {
	out := new(GetRecentBlocksReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetRecentBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetRecentTxs(ctx context.Context, in *GetRecentTxsRequest, opts ...grpc.CallOption) (*GetRecentTxsReply, error) {
	out := new(GetRecentTxsReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetRecentTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetBlockByHashOrNumber(ctx context.Context, in *GetBlockByHashOrNumberRequest, opts ...grpc.CallOption) (*GetBlockByHashOrNumberReply, error) {
	out := new(GetBlockByHashOrNumberReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetBlockByHashOrNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*GetTransactionByHashReply, error) {
	out := new(GetTransactionByHashReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetTransactionByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetTokenTransfers(ctx context.Context, in *GetTokenTransfersRequest, opts ...grpc.CallOption) (*GetTokenTransfersReply, error) {
	out := new(GetTokenTransfersReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetTokenTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockscoutGatewayServer is the server API for BlockscoutGateway service.
type BlockscoutGatewayServer interface {
	GetRecentBlocks(context.Context, *GetRecentBlocksRequest) (*GetRecentBlocksReply, error)
	GetRecentTxs(context.Context, *GetRecentTxsRequest) (*GetRecentTxsReply, error)
	GetBlockByHashOrNumber(context.Context, *GetBlockByHashOrNumberRequest) (*GetBlockByHashOrNumberReply, error)
	GetTransactionByHash(context.Context, *GetTransactionByHashRequest) (*GetTransactionByHashReply, error)
	GetTokenTransfers(context.Context, *GetTokenTransfersRequest) (*GetTokenTransfersReply, error)
}

// UnimplementedBlockscoutGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedBlockscoutGatewayServer struct {
}

func (*UnimplementedBlockscoutGatewayServer) GetRecentBlocks(ctx context.Context, req *GetRecentBlocksRequest) (*GetRecentBlocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentBlocks not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetRecentTxs(ctx context.Context, req *GetRecentTxsRequest) (*GetRecentTxsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentTxs not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetBlockByHashOrNumber(ctx context.Context, req *GetBlockByHashOrNumberRequest) (*GetBlockByHashOrNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHashOrNumber not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetTransactionByHash(ctx context.Context, req *GetTransactionByHashRequest) (*GetTransactionByHashReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByHash not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetTokenTransfers(ctx context.Context, req *GetTokenTransfersRequest) (*GetTokenTransfersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenTransfers not implemented")
}

func RegisterBlockscoutGatewayServer(s *grpc.Server, srv BlockscoutGatewayServer) {
	s.RegisterService(&_BlockscoutGateway_serviceDesc, srv)
}

func _BlockscoutGateway_GetRecentBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecentBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetRecentBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetRecentBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetRecentBlocks(ctx, req.(*GetRecentBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetRecentTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecentTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetRecentTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetRecentTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetRecentTxs(ctx, req.(*GetRecentTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetBlockByHashOrNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHashOrNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetBlockByHashOrNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetBlockByHashOrNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetBlockByHashOrNumber(ctx, req.(*GetBlockByHashOrNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetTransactionByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetTransactionByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetTransactionByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetTransactionByHash(ctx, req.(*GetTransactionByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetTokenTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetTokenTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetTokenTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetTokenTransfers(ctx, req.(*GetTokenTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockscoutGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.ankr.BlockscoutGateway",
	HandlerType: (*BlockscoutGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecentBlocks",
			Handler:    _BlockscoutGateway_GetRecentBlocks_Handler,
		},
		{
			MethodName: "GetRecentTxs",
			Handler:    _BlockscoutGateway_GetRecentTxs_Handler,
		},
		{
			MethodName: "GetBlockByHashOrNumber",
			Handler:    _BlockscoutGateway_GetBlockByHashOrNumber_Handler,
		},
		{
			MethodName: "GetTransactionByHash",
			Handler:    _BlockscoutGateway_GetTransactionByHash_Handler,
		},
		{
			MethodName: "GetTokenTransfers",
			Handler:    _BlockscoutGateway_GetTokenTransfers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
