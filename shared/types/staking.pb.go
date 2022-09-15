// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staking.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type ChainNativeCurrency struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimals             uint32   `protobuf:"varint,3,opt,name=decimals,proto3" json:"decimals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainNativeCurrency) Reset()         { *m = ChainNativeCurrency{} }
func (m *ChainNativeCurrency) String() string { return proto.CompactTextString(m) }
func (*ChainNativeCurrency) ProtoMessage()    {}
func (*ChainNativeCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{0}
}
func (m *ChainNativeCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainNativeCurrency.Unmarshal(m, b)
}
func (m *ChainNativeCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainNativeCurrency.Marshal(b, m, deterministic)
}
func (m *ChainNativeCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainNativeCurrency.Merge(m, src)
}
func (m *ChainNativeCurrency) XXX_Size() int {
	return xxx_messageInfo_ChainNativeCurrency.Size(m)
}
func (m *ChainNativeCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainNativeCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_ChainNativeCurrency proto.InternalMessageInfo

func (m *ChainNativeCurrency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChainNativeCurrency) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ChainNativeCurrency) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

type ChainExplorer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Standard             string   `protobuf:"bytes,3,opt,name=standard,proto3" json:"standard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainExplorer) Reset()         { *m = ChainExplorer{} }
func (m *ChainExplorer) String() string { return proto.CompactTextString(m) }
func (*ChainExplorer) ProtoMessage()    {}
func (*ChainExplorer) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{1}
}
func (m *ChainExplorer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainExplorer.Unmarshal(m, b)
}
func (m *ChainExplorer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainExplorer.Marshal(b, m, deterministic)
}
func (m *ChainExplorer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainExplorer.Merge(m, src)
}
func (m *ChainExplorer) XXX_Size() int {
	return xxx_messageInfo_ChainExplorer.Size(m)
}
func (m *ChainExplorer) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainExplorer.DiscardUnknown(m)
}

var xxx_messageInfo_ChainExplorer proto.InternalMessageInfo

func (m *ChainExplorer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChainExplorer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ChainExplorer) GetStandard() string {
	if m != nil {
		return m.Standard
	}
	return ""
}

type Chain struct {
	ChainId              uint64               `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ProjectUrl           string               `protobuf:"bytes,2,opt,name=project_url,json=projectUrl,proto3" json:"project_url,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ShortName            string               `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	Rpc                  []string             `protobuf:"bytes,5,rep,name=rpc,proto3" json:"rpc,omitempty"`
	Faucet               []string             `protobuf:"bytes,6,rep,name=faucet,proto3" json:"faucet,omitempty"`
	NativeCurrency       *ChainNativeCurrency `protobuf:"bytes,7,opt,name=native_currency,json=nativeCurrency,proto3" json:"native_currency,omitempty"`
	Explorer             []*ChainExplorer     `protobuf:"bytes,8,rep,name=explorer,proto3" json:"explorer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{2}
}
func (m *Chain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chain.Unmarshal(m, b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
}
func (m *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(m, src)
}
func (m *Chain) XXX_Size() int {
	return xxx_messageInfo_Chain.Size(m)
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Chain) GetProjectUrl() string {
	if m != nil {
		return m.ProjectUrl
	}
	return ""
}

func (m *Chain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chain) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *Chain) GetRpc() []string {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Chain) GetFaucet() []string {
	if m != nil {
		return m.Faucet
	}
	return nil
}

func (m *Chain) GetNativeCurrency() *ChainNativeCurrency {
	if m != nil {
		return m.NativeCurrency
	}
	return nil
}

func (m *Chain) GetExplorer() []*ChainExplorer {
	if m != nil {
		return m.Explorer
	}
	return nil
}

type Delegator struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Delegator) Reset()         { *m = Delegator{} }
func (m *Delegator) String() string { return proto.CompactTextString(m) }
func (*Delegator) ProtoMessage()    {}
func (*Delegator) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{3}
}
func (m *Delegator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delegator.Unmarshal(m, b)
}
func (m *Delegator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delegator.Marshal(b, m, deterministic)
}
func (m *Delegator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegator.Merge(m, src)
}
func (m *Delegator) XXX_Size() int {
	return xxx_messageInfo_Delegator.Size(m)
}
func (m *Delegator) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegator.DiscardUnknown(m)
}

var xxx_messageInfo_Delegator proto.InternalMessageInfo

func (m *Delegator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Delegator) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type GetChainsRequest struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Size_                uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChainsRequest) Reset()         { *m = GetChainsRequest{} }
func (m *GetChainsRequest) String() string { return proto.CompactTextString(m) }
func (*GetChainsRequest) ProtoMessage()    {}
func (*GetChainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{4}
}
func (m *GetChainsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChainsRequest.Unmarshal(m, b)
}
func (m *GetChainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChainsRequest.Marshal(b, m, deterministic)
}
func (m *GetChainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChainsRequest.Merge(m, src)
}
func (m *GetChainsRequest) XXX_Size() int {
	return xxx_messageInfo_GetChainsRequest.Size(m)
}
func (m *GetChainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChainsRequest proto.InternalMessageInfo

func (m *GetChainsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetChainsRequest) GetSize_() uint32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

type GetChainsReply struct {
	Chains               []*Chain `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains,omitempty"`
	HasMore              bool     `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChainsReply) Reset()         { *m = GetChainsReply{} }
func (m *GetChainsReply) String() string { return proto.CompactTextString(m) }
func (*GetChainsReply) ProtoMessage()    {}
func (*GetChainsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{5}
}
func (m *GetChainsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChainsReply.Unmarshal(m, b)
}
func (m *GetChainsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChainsReply.Marshal(b, m, deterministic)
}
func (m *GetChainsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChainsReply.Merge(m, src)
}
func (m *GetChainsReply) XXX_Size() int {
	return xxx_messageInfo_GetChainsReply.Size(m)
}
func (m *GetChainsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChainsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetChainsReply proto.InternalMessageInfo

func (m *GetChainsReply) GetChains() []*Chain {
	if m != nil {
		return m.Chains
	}
	return nil
}

func (m *GetChainsReply) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

type GetDelegatorsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDelegatorsRequest) Reset()         { *m = GetDelegatorsRequest{} }
func (m *GetDelegatorsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDelegatorsRequest) ProtoMessage()    {}
func (*GetDelegatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{6}
}
func (m *GetDelegatorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDelegatorsRequest.Unmarshal(m, b)
}
func (m *GetDelegatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDelegatorsRequest.Marshal(b, m, deterministic)
}
func (m *GetDelegatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDelegatorsRequest.Merge(m, src)
}
func (m *GetDelegatorsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDelegatorsRequest.Size(m)
}
func (m *GetDelegatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDelegatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDelegatorsRequest proto.InternalMessageInfo

type GetDelegatorsReply struct {
	Delegators           []*Delegator `protobuf:"bytes,1,rep,name=delegators,proto3" json:"delegators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetDelegatorsReply) Reset()         { *m = GetDelegatorsReply{} }
func (m *GetDelegatorsReply) String() string { return proto.CompactTextString(m) }
func (*GetDelegatorsReply) ProtoMessage()    {}
func (*GetDelegatorsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{7}
}
func (m *GetDelegatorsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDelegatorsReply.Unmarshal(m, b)
}
func (m *GetDelegatorsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDelegatorsReply.Marshal(b, m, deterministic)
}
func (m *GetDelegatorsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDelegatorsReply.Merge(m, src)
}
func (m *GetDelegatorsReply) XXX_Size() int {
	return xxx_messageInfo_GetDelegatorsReply.Size(m)
}
func (m *GetDelegatorsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDelegatorsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetDelegatorsReply proto.InternalMessageInfo

func (m *GetDelegatorsReply) GetDelegators() []*Delegator {
	if m != nil {
		return m.Delegators
	}
	return nil
}

func init() {
	proto.RegisterType((*ChainNativeCurrency)(nil), "com.ankr.ChainNativeCurrency")
	proto.RegisterType((*ChainExplorer)(nil), "com.ankr.ChainExplorer")
	proto.RegisterType((*Chain)(nil), "com.ankr.Chain")
	proto.RegisterType((*Delegator)(nil), "com.ankr.Delegator")
	proto.RegisterType((*GetChainsRequest)(nil), "com.ankr.GetChainsRequest")
	proto.RegisterType((*GetChainsReply)(nil), "com.ankr.GetChainsReply")
	proto.RegisterType((*GetDelegatorsRequest)(nil), "com.ankr.GetDelegatorsRequest")
	proto.RegisterType((*GetDelegatorsReply)(nil), "com.ankr.GetDelegatorsReply")
}

func init() { proto.RegisterFile("staking.proto", fileDescriptor_289e7c8aea278311) }

var fileDescriptor_289e7c8aea278311 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x9a, 0x34, 0x71, 0x26, 0x72, 0x5a, 0x6d, 0xaa, 0x62, 0x22, 0x0a, 0x91, 0x2f, 0xe4,
	0x94, 0x4a, 0xc9, 0xb1, 0x02, 0x89, 0x16, 0x88, 0x7a, 0x68, 0x05, 0x0b, 0x5c, 0x90, 0x50, 0xb4,
	0xb5, 0x27, 0x8d, 0xa9, 0xed, 0x35, 0xbb, 0x1b, 0xc0, 0xfc, 0x17, 0x2e, 0xfc, 0x52, 0xb4, 0x13,
	0xdb, 0xf9, 0x20, 0x5c, 0xac, 0x7d, 0x6f, 0xc6, 0x6f, 0xdf, 0xbc, 0xd1, 0x82, 0xab, 0x8d, 0x78,
	0x88, 0xd2, 0xfb, 0x51, 0xa6, 0xa4, 0x91, 0xcc, 0x09, 0x64, 0x32, 0x12, 0xe9, 0x83, 0xf2, 0xbf,
	0x40, 0xef, 0x6a, 0x21, 0xa2, 0xf4, 0x56, 0x98, 0xe8, 0x3b, 0x5e, 0x2d, 0x95, 0xc2, 0x34, 0xc8,
	0x19, 0x83, 0x46, 0x2a, 0x12, 0xf4, 0x6a, 0x83, 0xda, 0xb0, 0xcd, 0xe9, 0xcc, 0x4e, 0xa1, 0xa9,
	0xf3, 0xe4, 0x4e, 0xc6, 0xde, 0x01, 0xb1, 0x05, 0x62, 0x7d, 0x70, 0x42, 0x0c, 0xa2, 0x44, 0xc4,
	0xda, 0xab, 0x0f, 0x6a, 0x43, 0x97, 0x57, 0xd8, 0x7f, 0x0f, 0x2e, 0xc9, 0xbf, 0xf9, 0x99, 0xc5,
	0x52, 0xa1, 0xda, 0x2b, 0x7c, 0x0c, 0xf5, 0xa5, 0x2a, 0x55, 0xed, 0xd1, 0x4a, 0x6a, 0x23, 0xd2,
	0x50, 0xa8, 0x90, 0x24, 0xdb, 0xbc, 0xc2, 0xfe, 0xef, 0x03, 0x38, 0x24, 0x4d, 0xf6, 0x18, 0x9c,
	0xc0, 0x1e, 0x66, 0x51, 0x48, 0x7a, 0x0d, 0xde, 0x22, 0x7c, 0x1d, 0xb2, 0x67, 0xd0, 0xc9, 0x94,
	0xfc, 0x8a, 0x81, 0x99, 0xad, 0xa5, 0xa1, 0xa0, 0x3e, 0xa9, 0xb8, 0xf2, 0x51, 0xdf, 0xf0, 0x71,
	0x06, 0xa0, 0x17, 0x52, 0x99, 0x19, 0x55, 0x1a, 0x54, 0x69, 0x13, 0x73, 0x5b, 0xd8, 0x54, 0x59,
	0xe0, 0x1d, 0x0e, 0xea, 0xd6, 0xa6, 0xca, 0x02, 0x9b, 0xc8, 0x5c, 0x2c, 0x03, 0x34, 0x5e, 0x93,
	0xc8, 0x02, 0xb1, 0xb7, 0x70, 0x94, 0x52, 0x9e, 0xb3, 0xa0, 0x08, 0xd4, 0x6b, 0x0d, 0x6a, 0xc3,
	0xce, 0xf8, 0x6c, 0x54, 0x06, 0x3f, 0xda, 0x93, 0x3a, 0xef, 0xa6, 0xdb, 0x5b, 0x98, 0x80, 0x83,
	0x45, 0x70, 0x9e, 0x33, 0xa8, 0x0f, 0x3b, 0xe3, 0x47, 0x3b, 0x02, 0x65, 0xae, 0xbc, 0x6a, 0xf4,
	0x5f, 0x40, 0xfb, 0x35, 0xc6, 0x78, 0x2f, 0x8c, 0x54, 0xcc, 0x83, 0x96, 0x08, 0x43, 0x85, 0x5a,
	0x17, 0x89, 0x97, 0xd0, 0x7a, 0x17, 0x89, 0x5c, 0xa6, 0xa6, 0xdc, 0xe6, 0x0a, 0xf9, 0x2f, 0xe1,
	0x78, 0x8a, 0x86, 0xc4, 0x35, 0xc7, 0x6f, 0x4b, 0xd4, 0xc6, 0xf6, 0xca, 0xf9, 0x5c, 0xa3, 0x21,
	0x11, 0x97, 0x17, 0xc8, 0x86, 0xa8, 0xa3, 0x5f, 0x48, 0x0a, 0x2e, 0xa7, 0xb3, 0xff, 0x11, 0xba,
	0x1b, 0xff, 0x67, 0x71, 0xce, 0x9e, 0x43, 0x93, 0xd6, 0x62, 0x2d, 0xd8, 0x19, 0x8e, 0x76, 0x66,
	0xe0, 0x45, 0xd9, 0xee, 0x73, 0x21, 0xf4, 0x2c, 0x91, 0x6a, 0x25, 0xe9, 0xf0, 0xd6, 0x42, 0xe8,
	0x1b, 0xa9, 0xd0, 0x3f, 0x85, 0x93, 0x29, 0x9a, 0x6a, 0xae, 0xd2, 0x99, 0x7f, 0x0d, 0x6c, 0x87,
	0xb7, 0x37, 0x4e, 0x00, 0xc2, 0x8a, 0x2a, 0x6e, 0xed, 0xad, 0x6f, 0xad, 0xda, 0xf9, 0x46, 0xdb,
	0xf8, 0x4f, 0x0d, 0xba, 0x1f, 0x56, 0xaf, 0x64, 0x2a, 0x0c, 0xfe, 0x10, 0x39, 0x7b, 0x05, 0xed,
	0x6a, 0x16, 0xd6, 0x5f, 0x0b, 0xec, 0x06, 0xd4, 0xf7, 0xf6, 0xd6, 0xac, 0x95, 0x1b, 0x70, 0xb7,
	0x0c, 0xb2, 0xa7, 0x5b, 0xad, 0xff, 0x4c, 0xd4, 0x7f, 0xf2, 0xdf, 0x7a, 0x16, 0xe7, 0x97, 0x17,
	0xd0, 0xad, 0xca, 0xf4, 0x94, 0x2f, 0x7b, 0xdb, 0x9e, 0xdf, 0x59, 0xf2, 0xf3, 0xc9, 0xe8, 0x5c,
	0x2f, 0x84, 0xc2, 0xf0, 0xdc, 0xe4, 0x19, 0xea, 0x0b, 0xfa, 0xde, 0x35, 0xe9, 0x8f, 0xc9, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xcb, 0xc3, 0xed, 0x0d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StakingGatewayClient is the client API for StakingGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StakingGatewayClient interface {
	GetChains(ctx context.Context, in *GetChainsRequest, opts ...grpc.CallOption) (*GetChainsReply, error)
	GetDelegators(ctx context.Context, in *GetDelegatorsRequest, opts ...grpc.CallOption) (*GetDelegatorsReply, error)
}

type stakingGatewayClient struct {
	cc *grpc.ClientConn
}

func NewStakingGatewayClient(cc *grpc.ClientConn) StakingGatewayClient {
	return &stakingGatewayClient{cc}
}

func (c *stakingGatewayClient) GetChains(ctx context.Context, in *GetChainsRequest, opts ...grpc.CallOption) (*GetChainsReply, error) {
	out := new(GetChainsReply)
	err := c.cc.Invoke(ctx, "/com.ankr.StakingGateway/GetChains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakingGatewayClient) GetDelegators(ctx context.Context, in *GetDelegatorsRequest, opts ...grpc.CallOption) (*GetDelegatorsReply, error) {
	out := new(GetDelegatorsReply)
	err := c.cc.Invoke(ctx, "/com.ankr.StakingGateway/GetDelegators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakingGatewayServer is the server API for StakingGateway service.
type StakingGatewayServer interface {
	GetChains(context.Context, *GetChainsRequest) (*GetChainsReply, error)
	GetDelegators(context.Context, *GetDelegatorsRequest) (*GetDelegatorsReply, error)
}

// UnimplementedStakingGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedStakingGatewayServer struct {
}

func (*UnimplementedStakingGatewayServer) GetChains(ctx context.Context, req *GetChainsRequest) (*GetChainsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChains not implemented")
}
func (*UnimplementedStakingGatewayServer) GetDelegators(ctx context.Context, req *GetDelegatorsRequest) (*GetDelegatorsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegators not implemented")
}

func RegisterStakingGatewayServer(s *grpc.Server, srv StakingGatewayServer) {
	s.RegisterService(&_StakingGateway_serviceDesc, srv)
}

func _StakingGateway_GetChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingGatewayServer).GetChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.StakingGateway/GetChains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingGatewayServer).GetChains(ctx, req.(*GetChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakingGateway_GetDelegators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingGatewayServer).GetDelegators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.StakingGateway/GetDelegators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingGatewayServer).GetDelegators(ctx, req.(*GetDelegatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StakingGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.ankr.StakingGateway",
	HandlerType: (*StakingGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChains",
			Handler:    _StakingGateway_GetChains_Handler,
		},
		{
			MethodName: "GetDelegators",
			Handler:    _StakingGateway_GetDelegators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staking.proto",
}
