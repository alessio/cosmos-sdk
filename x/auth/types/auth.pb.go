// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/auth/v1beta1/auth.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
	Address       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	PubKey        []byte                                        `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"public_key,omitempty" yaml:"public_key"`
	AccountNumber uint64                                        `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty" yaml:"account_number"`
	Sequence      uint64                                        `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *BaseAccount) Reset()      { *m = BaseAccount{} }
func (*BaseAccount) ProtoMessage() {}
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{0}
}
func (m *BaseAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAccount.Merge(m, src)
}
func (m *BaseAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAccount proto.InternalMessageInfo

// ModuleAccount defines an account for modules that holds coins on a pool
type ModuleAccount struct {
	*BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions  []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *ModuleAccount) Reset()      { *m = ModuleAccount{} }
func (*ModuleAccount) ProtoMessage() {}
func (*ModuleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{1}
}
func (m *ModuleAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleAccount.Merge(m, src)
}
func (m *ModuleAccount) XXX_Size() int {
	return m.Size()
}
func (m *ModuleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleAccount proto.InternalMessageInfo

// Params defines the parameters for the auth module.
type Params struct {
	MaxMemoCharacters      *types.UInt64Value `protobuf:"bytes,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty" yaml:"max_memo_characters"`
	TxSigLimit             *types.UInt64Value `protobuf:"bytes,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty" yaml:"tx_sig_limit"`
	TxSizeCostPerByte      *types.UInt64Value `protobuf:"bytes,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty" yaml:"tx_size_cost_per_byte"`
	SigVerifyCostED25519   *types.UInt64Value `protobuf:"bytes,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty" yaml:"sig_verify_cost_ed25519"`
	SigVerifyCostSecp256k1 *types.UInt64Value `protobuf:"bytes,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty" yaml:"sig_verify_cost_secp256k1"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxMemoCharacters() *types.UInt64Value {
	if m != nil {
		return m.MaxMemoCharacters
	}
	return nil
}

func (m *Params) GetTxSigLimit() *types.UInt64Value {
	if m != nil {
		return m.TxSigLimit
	}
	return nil
}

func (m *Params) GetTxSizeCostPerByte() *types.UInt64Value {
	if m != nil {
		return m.TxSizeCostPerByte
	}
	return nil
}

func (m *Params) GetSigVerifyCostED25519() *types.UInt64Value {
	if m != nil {
		return m.SigVerifyCostED25519
	}
	return nil
}

func (m *Params) GetSigVerifyCostSecp256k1() *types.UInt64Value {
	if m != nil {
		return m.SigVerifyCostSecp256k1
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseAccount)(nil), "cosmos.auth.v1beta1.BaseAccount")
	proto.RegisterType((*ModuleAccount)(nil), "cosmos.auth.v1beta1.ModuleAccount")
	proto.RegisterType((*Params)(nil), "cosmos.auth.v1beta1.Params")
}

func init() { proto.RegisterFile("cosmos/auth/v1beta1/auth.proto", fileDescriptor_7e1f7e915d020d2d) }

var fileDescriptor_7e1f7e915d020d2d = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x4f, 0x13, 0x41,
	0x14, 0xee, 0x42, 0xe5, 0xc7, 0x14, 0x48, 0x58, 0x0a, 0x94, 0x4a, 0x76, 0x9b, 0x3d, 0x61, 0x62,
	0xb7, 0x29, 0x0a, 0x89, 0x3d, 0x18, 0x29, 0x6a, 0x42, 0x10, 0x43, 0x96, 0x88, 0x89, 0x97, 0x75,
	0x76, 0xfb, 0x58, 0x36, 0x74, 0x3a, 0xeb, 0xce, 0x2c, 0x76, 0xf9, 0x0b, 0xbc, 0xe9, 0x49, 0x3d,
	0xf2, 0x47, 0xf8, 0x1f, 0x78, 0xf1, 0x48, 0x3c, 0x79, 0xda, 0x98, 0x72, 0x31, 0x1e, 0x7b, 0x31,
	0xf1, 0x64, 0x76, 0x76, 0x29, 0x2d, 0xa9, 0xe2, 0xa9, 0xfb, 0xbe, 0x37, 0xdf, 0xf7, 0xbe, 0xf9,
	0xfa, 0x32, 0x48, 0xb1, 0x29, 0x23, 0x94, 0x55, 0x70, 0xc0, 0x0f, 0x2b, 0xc7, 0x55, 0x0b, 0x38,
	0xae, 0x8a, 0x42, 0xf7, 0x7c, 0xca, 0xa9, 0x3c, 0x97, 0xf4, 0x75, 0x01, 0xa5, 0xfd, 0xe2, 0x52,
	0x02, 0x9a, 0xe2, 0x48, 0x25, 0x3d, 0x21, 0x8a, 0x62, 0xde, 0xa1, 0x0e, 0x4d, 0xf0, 0xf8, 0x2b,
	0x45, 0x15, 0x87, 0x52, 0xa7, 0x09, 0x15, 0x51, 0x59, 0xc1, 0x41, 0xe5, 0xb5, 0x8f, 0x3d, 0x0f,
	0xfc, 0x94, 0xa5, 0xbd, 0x1f, 0x41, 0xb9, 0x3a, 0x66, 0xb0, 0x61, 0xdb, 0x34, 0x68, 0x71, 0x79,
	0x1b, 0x8d, 0xe3, 0x46, 0xc3, 0x07, 0xc6, 0x0a, 0x52, 0x49, 0x5a, 0x99, 0xaa, 0x57, 0x7f, 0x47,
	0x6a, 0xd9, 0x71, 0xf9, 0x61, 0x60, 0xe9, 0x36, 0x25, 0xe9, 0xcc, 0xf4, 0xa7, 0xcc, 0x1a, 0x47,
	0x15, 0x1e, 0x7a, 0xc0, 0xf4, 0x0d, 0xdb, 0xde, 0x48, 0x88, 0xc6, 0x85, 0x82, 0xfc, 0x18, 0x8d,
	0x7b, 0x81, 0x65, 0x1e, 0x41, 0x58, 0x18, 0x11, 0x62, 0xe5, 0x9f, 0x91, 0x9a, 0xf7, 0x02, 0xab,
	0xe9, 0xda, 0x31, 0x7a, 0x9b, 0x12, 0x97, 0x03, 0xf1, 0x78, 0xd8, 0x8d, 0xd4, 0xd9, 0x10, 0x93,
	0x66, 0x4d, 0xbb, 0xec, 0x6a, 0xc6, 0x98, 0x17, 0x58, 0xdb, 0x10, 0xca, 0x0f, 0xd0, 0x0c, 0x4e,
	0xfc, 0x99, 0xad, 0x80, 0x58, 0xe0, 0x17, 0x46, 0x4b, 0xd2, 0x4a, 0xb6, 0xbe, 0xd4, 0x8d, 0xd4,
	0xf9, 0x84, 0x36, 0xd8, 0xd7, 0x8c, 0xe9, 0x14, 0x78, 0x2a, 0x6a, 0xb9, 0x88, 0x26, 0x18, 0xbc,
	0x0a, 0xa0, 0x65, 0x43, 0x21, 0x1b, 0x73, 0x8d, 0x5e, 0x5d, 0xcb, 0xbf, 0x39, 0x55, 0x33, 0x1f,
	0x4f, 0xd5, 0xcc, 0xd7, 0x4f, 0xe5, 0x89, 0x34, 0x87, 0x2d, 0xed, 0xb3, 0x84, 0xa6, 0x77, 0x68,
	0x23, 0x68, 0xf6, 0xa2, 0x79, 0x89, 0xa6, 0x2c, 0xcc, 0xc0, 0x4c, 0x95, 0x45, 0x3e, 0xb9, 0xd5,
	0x92, 0x3e, 0xe4, 0x7f, 0xd2, 0xfb, 0x22, 0xad, 0xdf, 0x3c, 0x8b, 0x54, 0xa9, 0x1b, 0xa9, 0x73,
	0x89, 0xd3, 0x7e, 0x0d, 0xcd, 0xc8, 0x59, 0x7d, 0xe1, 0xcb, 0x28, 0xdb, 0xc2, 0x04, 0x44, 0x58,
	0x93, 0x86, 0xf8, 0x96, 0x4b, 0x28, 0xe7, 0x81, 0x4f, 0x5c, 0xc6, 0x5c, 0xda, 0x62, 0x85, 0xd1,
	0xd2, 0xe8, 0xca, 0xa4, 0xd1, 0x0f, 0xd5, 0x8a, 0x7d, 0xfe, 0x67, 0x06, 0x2c, 0x6f, 0x69, 0xbf,
	0xb2, 0x68, 0x6c, 0x17, 0xfb, 0x98, 0x30, 0xb9, 0x89, 0xe6, 0x08, 0x6e, 0x9b, 0x04, 0x08, 0x35,
	0xed, 0x43, 0xec, 0x63, 0x9b, 0x83, 0xcf, 0xd2, 0x5b, 0x2c, 0xeb, 0xc9, 0x9e, 0xe8, 0x17, 0x7b,
	0xa2, 0x3f, 0xdb, 0x6a, 0xf1, 0xf5, 0xbb, 0xfb, 0xb8, 0x19, 0x40, 0x5d, 0xe9, 0x46, 0x6a, 0x31,
	0x71, 0x3f, 0x44, 0x42, 0x33, 0x66, 0x09, 0x6e, 0xef, 0x00, 0xa1, 0x9b, 0x3d, 0x4c, 0x7e, 0x8e,
	0xa6, 0x78, 0xdb, 0x64, 0xae, 0x63, 0x36, 0x5d, 0xe2, 0x72, 0x71, 0xa5, 0xeb, 0xc6, 0x2c, 0x5e,
	0x86, 0xd4, 0xcf, 0xd5, 0x0c, 0xc4, 0xdb, 0x7b, 0xae, 0xf3, 0x24, 0x2e, 0x64, 0x0f, 0xcd, 0x8b,
	0xe6, 0x09, 0x98, 0x36, 0x65, 0xdc, 0xf4, 0xc0, 0x37, 0xad, 0x90, 0x83, 0x58, 0x89, 0xeb, 0x26,
	0x94, 0xba, 0x91, 0xba, 0xdc, 0x37, 0xe1, 0xaa, 0x88, 0x66, 0xcc, 0xc6, 0xa3, 0x4e, 0x60, 0x93,
	0x32, 0xbe, 0x0b, 0x7e, 0x3d, 0xe4, 0x20, 0xbf, 0x95, 0xd0, 0x62, 0x6c, 0xe6, 0x18, 0x7c, 0xf7,
	0x20, 0x4c, 0x08, 0xd0, 0x58, 0x5d, 0x5b, 0xab, 0xde, 0x13, 0xbb, 0x74, 0xdd, 0xd0, 0x5a, 0x27,
	0x52, 0xf3, 0x7b, 0xae, 0xb3, 0x2f, 0xf8, 0xb1, 0xf2, 0xa3, 0x87, 0x82, 0xdd, 0x8d, 0x54, 0x25,
	0x31, 0xf3, 0x17, 0x79, 0xcd, 0xc8, 0xb3, 0x01, 0x5e, 0x02, 0xcb, 0x1f, 0x24, 0xb4, 0x74, 0x95,
	0xc2, 0xc0, 0xf6, 0x56, 0xd7, 0xd6, 0x8f, 0xaa, 0x85, 0x1b, 0xff, 0xe1, 0xe9, 0x7e, 0x27, 0x52,
	0x17, 0x06, 0x3c, 0xed, 0x5d, 0xf0, 0xbb, 0x91, 0x5a, 0x1a, 0xee, 0xaa, 0x37, 0x42, 0x33, 0x16,
	0xd8, 0x50, 0x6e, 0x6d, 0x22, 0xde, 0xc3, 0x1f, 0xa7, 0xaa, 0x54, 0xdf, 0xfc, 0xd2, 0x51, 0xa4,
	0xb3, 0x8e, 0x22, 0x7d, 0xef, 0x28, 0xd2, 0xbb, 0x73, 0x25, 0x73, 0x76, 0xae, 0x64, 0xbe, 0x9d,
	0x2b, 0x99, 0x17, 0xb7, 0xfe, 0xf9, 0x9a, 0xb4, 0x93, 0x07, 0x51, 0x3c, 0x2a, 0xd6, 0x98, 0x30,
	0x7f, 0xe7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc2, 0xb8, 0x4b, 0x2c, 0x05, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MaxMemoCharacters.Equal(that1.MaxMemoCharacters) {
		return false
	}
	if !this.TxSigLimit.Equal(that1.TxSigLimit) {
		return false
	}
	if !this.TxSizeCostPerByte.Equal(that1.TxSizeCostPerByte) {
		return false
	}
	if !this.SigVerifyCostED25519.Equal(that1.SigVerifyCostED25519) {
		return false
	}
	if !this.SigVerifyCostSecp256k1.Equal(that1.SigVerifyCostSecp256k1) {
		return false
	}
	return true
}
func (m *BaseAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if m.AccountNumber != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModuleAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Permissions[iNdEx])
			copy(dAtA[i:], m.Permissions[iNdEx])
			i = encodeVarintAuth(dAtA, i, uint64(len(m.Permissions[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SigVerifyCostSecp256k1 != nil {
		{
			size, err := m.SigVerifyCostSecp256k1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.SigVerifyCostED25519 != nil {
		{
			size, err := m.SigVerifyCostED25519.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TxSizeCostPerByte != nil {
		{
			size, err := m.TxSizeCostPerByte.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TxSigLimit != nil {
		{
			size, err := m.TxSigLimit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxMemoCharacters != nil {
		{
			size, err := m.MaxMemoCharacters.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovAuth(uint64(m.AccountNumber))
	}
	if m.Sequence != 0 {
		n += 1 + sovAuth(uint64(m.Sequence))
	}
	return n
}

func (m *ModuleAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, s := range m.Permissions {
			l = len(s)
			n += 1 + l + sovAuth(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxMemoCharacters != nil {
		l = m.MaxMemoCharacters.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.TxSigLimit != nil {
		l = m.TxSigLimit.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.TxSizeCostPerByte != nil {
		l = m.TxSizeCostPerByte.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.SigVerifyCostED25519 != nil {
		l = m.SigVerifyCostED25519.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.SigVerifyCostSecp256k1 != nil {
		l = m.SigVerifyCostSecp256k1.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func sovAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: BaseAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ModuleAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: ModuleAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMemoCharacters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxMemoCharacters == nil {
				m.MaxMemoCharacters = &types.UInt64Value{}
			}
			if err := m.MaxMemoCharacters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSigLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxSigLimit == nil {
				m.TxSigLimit = &types.UInt64Value{}
			}
			if err := m.TxSigLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSizeCostPerByte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxSizeCostPerByte == nil {
				m.TxSizeCostPerByte = &types.UInt64Value{}
			}
			if err := m.TxSizeCostPerByte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostED25519", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SigVerifyCostED25519 == nil {
				m.SigVerifyCostED25519 = &types.UInt64Value{}
			}
			if err := m.SigVerifyCostED25519.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostSecp256k1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SigVerifyCostSecp256k1 == nil {
				m.SigVerifyCostSecp256k1 = &types.UInt64Value{}
			}
			if err := m.SigVerifyCostSecp256k1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuth = fmt.Errorf("proto: unexpected end of group")
)
