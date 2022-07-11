// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token_transfer.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TokenTransfer struct {
	TxHash               []byte   `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	AddressFrom          []byte   `protobuf:"bytes,2,opt,name=address_from,json=addressFrom,proto3" json:"address_from,omitempty"`
	AddressTo            []byte   `protobuf:"bytes,3,opt,name=address_to,json=addressTo,proto3" json:"address_to,omitempty"`
	TokenContract        []byte   `protobuf:"bytes,4,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
	Amount               uint64   `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	BlockNumber          uint64   `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Timestamp            uint64   `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenTransfer) Reset()         { *m = TokenTransfer{} }
func (m *TokenTransfer) String() string { return proto.CompactTextString(m) }
func (*TokenTransfer) ProtoMessage()    {}
func (*TokenTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dee5df8e2b2416f, []int{0}
}
func (m *TokenTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTransfer.Merge(m, src)
}
func (m *TokenTransfer) XXX_Size() int {
	return m.Size()
}
func (m *TokenTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTransfer proto.InternalMessageInfo

func (m *TokenTransfer) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *TokenTransfer) GetAddressFrom() []byte {
	if m != nil {
		return m.AddressFrom
	}
	return nil
}

func (m *TokenTransfer) GetAddressTo() []byte {
	if m != nil {
		return m.AddressTo
	}
	return nil
}

func (m *TokenTransfer) GetTokenContract() []byte {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *TokenTransfer) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TokenTransfer) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *TokenTransfer) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenTransfer)(nil), "com.ankr.TokenTransfer")
}

func init() { proto.RegisterFile("token_transfer.proto", fileDescriptor_4dee5df8e2b2416f) }

var fileDescriptor_4dee5df8e2b2416f = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xf9, 0xac, 0xa9, 0x9d, 0xfe, 0x59, 0x0c, 0x45, 0x67, 0xa1, 0xa1, 0x0a, 0x42,
	0x57, 0xe9, 0xc2, 0xa5, 0xbb, 0x0a, 0x22, 0x2e, 0x44, 0x4a, 0x56, 0x6e, 0xc2, 0x24, 0x99, 0x12,
	0x89, 0x33, 0x37, 0xcc, 0xdc, 0x42, 0x7d, 0x13, 0x1f, 0xc9, 0xa5, 0x8f, 0x20, 0x71, 0xe7, 0x53,
	0x48, 0x6e, 0xa7, 0x88, 0x9b, 0x81, 0xf9, 0x9d, 0xdf, 0x85, 0xc3, 0xe1, 0x53, 0x84, 0x5a, 0xdb,
	0x0c, 0x9d, 0xb2, 0x7e, 0xad, 0x5d, 0xd2, 0x38, 0x40, 0x10, 0x47, 0x05, 0x98, 0x44, 0xd9, 0xda,
	0x5d, 0x7c, 0x33, 0x3e, 0x4e, 0x3b, 0x25, 0x0d, 0x86, 0x38, 0xe1, 0x7d, 0xdc, 0x66, 0x95, 0xf2,
	0x95, 0x64, 0x33, 0x36, 0x1f, 0xad, 0x22, 0xdc, 0xde, 0x29, 0x5f, 0x89, 0x73, 0x3e, 0x52, 0x65,
	0xe9, 0xb4, 0xf7, 0xd9, 0xda, 0x81, 0x91, 0xff, 0x29, 0x1d, 0x06, 0x76, 0xeb, 0xc0, 0x88, 0x33,
	0xce, 0xf7, 0x0a, 0x82, 0x3c, 0x20, 0x61, 0x10, 0x48, 0x0a, 0xe2, 0x92, 0x4f, 0x76, 0x75, 0x0a,
	0xb0, 0xe8, 0x54, 0x81, 0xb2, 0x47, 0xca, 0x98, 0xe8, 0x4d, 0x80, 0xe2, 0x98, 0x47, 0xca, 0xc0,
	0xc6, 0xa2, 0x3c, 0x9c, 0xb1, 0x79, 0x6f, 0x15, 0x7e, 0x5d, 0x81, 0xfc, 0x05, 0x8a, 0x3a, 0xb3,
	0x1b, 0x93, 0x6b, 0x27, 0x23, 0x4a, 0x87, 0xc4, 0x1e, 0x08, 0x89, 0x53, 0x3e, 0xc0, 0x67, 0xa3,
	0x3d, 0x2a, 0xd3, 0xc8, 0x3e, 0xe5, 0xbf, 0x60, 0x79, 0xff, 0xde, 0xc6, 0xec, 0xa3, 0x8d, 0xd9,
	0x67, 0x1b, 0xb3, 0xb7, 0xaf, 0xf8, 0x1f, 0x9f, 0xec, 0x87, 0xd8, 0x0d, 0xb3, 0x14, 0x7f, 0xb6,
	0x78, 0xec, 0xd8, 0xd3, 0x34, 0x59, 0xf8, 0x4a, 0x39, 0x5d, 0x2e, 0xf0, 0xb5, 0xd1, 0xfe, 0x9a,
	0xde, 0x3c, 0xa2, 0x83, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x80, 0x08, 0x66, 0x61,
	0x01, 0x00, 0x00,
}

func (m *TokenTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != 0 {
		i = encodeVarintTokenTransfer(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x38
	}
	if m.BlockNumber != 0 {
		i = encodeVarintTokenTransfer(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x30
	}
	if m.Amount != 0 {
		i = encodeVarintTokenTransfer(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TokenContract) > 0 {
		i -= len(m.TokenContract)
		copy(dAtA[i:], m.TokenContract)
		i = encodeVarintTokenTransfer(dAtA, i, uint64(len(m.TokenContract)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AddressTo) > 0 {
		i -= len(m.AddressTo)
		copy(dAtA[i:], m.AddressTo)
		i = encodeVarintTokenTransfer(dAtA, i, uint64(len(m.AddressTo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AddressFrom) > 0 {
		i -= len(m.AddressFrom)
		copy(dAtA[i:], m.AddressFrom)
		i = encodeVarintTokenTransfer(dAtA, i, uint64(len(m.AddressFrom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintTokenTransfer(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovTokenTransfer(uint64(l))
	}
	l = len(m.AddressFrom)
	if l > 0 {
		n += 1 + l + sovTokenTransfer(uint64(l))
	}
	l = len(m.AddressTo)
	if l > 0 {
		n += 1 + l + sovTokenTransfer(uint64(l))
	}
	l = len(m.TokenContract)
	if l > 0 {
		n += 1 + l + sovTokenTransfer(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTokenTransfer(uint64(m.Amount))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovTokenTransfer(uint64(m.BlockNumber))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTokenTransfer(uint64(m.Timestamp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTokenTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenTransfer(x uint64) (n int) {
	return sovTokenTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenTransfer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressFrom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressFrom = append(m.AddressFrom[:0], dAtA[iNdEx:postIndex]...)
			if m.AddressFrom == nil {
				m.AddressFrom = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressTo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressTo = append(m.AddressTo[:0], dAtA[iNdEx:postIndex]...)
			if m.AddressTo == nil {
				m.AddressTo = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContract", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContract = append(m.TokenContract[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenContract == nil {
				m.TokenContract = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokenTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTokenTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenTransfer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTokenTransfer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTokenTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenTransfer = fmt.Errorf("proto: unexpected end of group")
)