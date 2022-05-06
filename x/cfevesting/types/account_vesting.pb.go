// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfevesting/account_vesting.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccountVestingsList struct {
	Vestings []*AccountVestings `protobuf:"bytes,1,rep,name=vestings,proto3" json:"vestings,omitempty"`
}

func (m *AccountVestingsList) Reset()         { *m = AccountVestingsList{} }
func (m *AccountVestingsList) String() string { return proto.CompactTextString(m) }
func (*AccountVestingsList) ProtoMessage()    {}
func (*AccountVestingsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_587a03f29f224c3f, []int{0}
}
func (m *AccountVestingsList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountVestingsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountVestingsList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountVestingsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountVestingsList.Merge(m, src)
}
func (m *AccountVestingsList) XXX_Size() int {
	return m.Size()
}
func (m *AccountVestingsList) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountVestingsList.DiscardUnknown(m)
}

var xxx_messageInfo_AccountVestingsList proto.InternalMessageInfo

func (m *AccountVestingsList) GetVestings() []*AccountVestings {
	if m != nil {
		return m.Vestings
	}
	return nil
}

type AccountVestings struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// string delegable_address = 2;
	VestingPools []*VestingPool `protobuf:"bytes,3,rep,name=vesting_pools,json=vestingPools,proto3" json:"vesting_pools,omitempty"`
}

func (m *AccountVestings) Reset()         { *m = AccountVestings{} }
func (m *AccountVestings) String() string { return proto.CompactTextString(m) }
func (*AccountVestings) ProtoMessage()    {}
func (*AccountVestings) Descriptor() ([]byte, []int) {
	return fileDescriptor_587a03f29f224c3f, []int{1}
}
func (m *AccountVestings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountVestings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountVestings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountVestings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountVestings.Merge(m, src)
}
func (m *AccountVestings) XXX_Size() int {
	return m.Size()
}
func (m *AccountVestings) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountVestings.DiscardUnknown(m)
}

var xxx_messageInfo_AccountVestings proto.InternalMessageInfo

func (m *AccountVestings) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountVestings) GetVestingPools() []*VestingPool {
	if m != nil {
		return m.VestingPools
	}
	return nil
}

type VestingPool struct {
	Id          int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VestingType string    `protobuf:"bytes,3,opt,name=vesting_type,json=vestingType,proto3" json:"vesting_type,omitempty"`
	LockStart   time.Time `protobuf:"bytes,4,opt,name=lock_start,json=lockStart,proto3,stdtime" json:"lock_start"`
	LockEnd     time.Time `protobuf:"bytes,5,opt,name=lock_end,json=lockEnd,proto3,stdtime" json:"lock_end"`
	// google.protobuf.Timestamp vesting_end = 5 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
	Vested github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=vested,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"vested"`
	// google.protobuf.Duration release_period = 7 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
	// bool delegation_allowed = 8;
	Withdrawn                 github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=withdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"withdrawn"`
	Sent                      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=sent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sent"`
	LastModification          time.Time                              `protobuf:"bytes,9,opt,name=last_modification,json=lastModification,proto3,stdtime" json:"last_modification"`
	LastModificationVested    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=last_modification_vested,json=lastModificationVested,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_modification_vested"`
	LastModificationWithdrawn github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=last_modification_withdrawn,json=lastModificationWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_modification_withdrawn"`
	TransferAllowed           bool                                   `protobuf:"varint,12,opt,name=transfer_allowed,json=transferAllowed,proto3" json:"transfer_allowed,omitempty"`
}

func (m *VestingPool) Reset()         { *m = VestingPool{} }
func (m *VestingPool) String() string { return proto.CompactTextString(m) }
func (*VestingPool) ProtoMessage()    {}
func (*VestingPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_587a03f29f224c3f, []int{2}
}
func (m *VestingPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingPool.Merge(m, src)
}
func (m *VestingPool) XXX_Size() int {
	return m.Size()
}
func (m *VestingPool) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingPool.DiscardUnknown(m)
}

var xxx_messageInfo_VestingPool proto.InternalMessageInfo

func (m *VestingPool) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VestingPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VestingPool) GetVestingType() string {
	if m != nil {
		return m.VestingType
	}
	return ""
}

func (m *VestingPool) GetLockStart() time.Time {
	if m != nil {
		return m.LockStart
	}
	return time.Time{}
}

func (m *VestingPool) GetLockEnd() time.Time {
	if m != nil {
		return m.LockEnd
	}
	return time.Time{}
}

func (m *VestingPool) GetLastModification() time.Time {
	if m != nil {
		return m.LastModification
	}
	return time.Time{}
}

func (m *VestingPool) GetTransferAllowed() bool {
	if m != nil {
		return m.TransferAllowed
	}
	return false
}

func init() {
	proto.RegisterType((*AccountVestingsList)(nil), "chain4energy.c4echain.cfevesting.AccountVestingsList")
	proto.RegisterType((*AccountVestings)(nil), "chain4energy.c4echain.cfevesting.AccountVestings")
	proto.RegisterType((*VestingPool)(nil), "chain4energy.c4echain.cfevesting.VestingPool")
}

func init() { proto.RegisterFile("cfevesting/account_vesting.proto", fileDescriptor_587a03f29f224c3f) }

var fileDescriptor_587a03f29f224c3f = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x30,
	0x1c, 0x6f, 0xda, 0xae, 0x1f, 0xee, 0x60, 0xc3, 0x20, 0x64, 0x8a, 0x94, 0x96, 0x1e, 0x50, 0x39,
	0xd4, 0x11, 0xa3, 0x77, 0xb4, 0x22, 0x90, 0x90, 0x36, 0x04, 0x61, 0x1a, 0x12, 0x97, 0xc8, 0x8d,
	0xdd, 0xd4, 0x5a, 0x62, 0x57, 0xb1, 0xbb, 0xd2, 0x13, 0x2f, 0xc0, 0x61, 0x8f, 0xb5, 0xe3, 0x8e,
	0x88, 0xc3, 0x40, 0xed, 0x8b, 0xa0, 0x38, 0xe9, 0x1a, 0x75, 0x87, 0x69, 0x3d, 0xc5, 0xff, 0x8f,
	0xdf, 0x47, 0x7e, 0x51, 0x0c, 0xda, 0xfe, 0x88, 0x9d, 0x33, 0xa5, 0xb9, 0x08, 0x1c, 0xe2, 0xfb,
	0x72, 0x2a, 0xb4, 0x97, 0xd5, 0x78, 0x12, 0x4b, 0x2d, 0x61, 0xdb, 0x1f, 0x13, 0x2e, 0xfa, 0x4c,
	0xb0, 0x38, 0x98, 0x63, 0xbf, 0xcf, 0x4c, 0x8d, 0xd7, 0xb8, 0xe6, 0x93, 0x40, 0x06, 0xd2, 0x2c,
	0x3b, 0xc9, 0x29, 0xc5, 0x35, 0x5b, 0x81, 0x94, 0x41, 0xc8, 0x1c, 0x53, 0x0d, 0xa7, 0x23, 0x47,
	0xf3, 0x88, 0x29, 0x4d, 0xa2, 0x49, 0xba, 0xd0, 0xa1, 0xe0, 0xf1, 0x61, 0xaa, 0x78, 0x9a, 0x12,
	0xa9, 0x23, 0xae, 0x34, 0x3c, 0x06, 0xb5, 0x8c, 0x58, 0x21, 0xab, 0x5d, 0xea, 0x36, 0x0e, 0x5e,
	0xe3, 0xbb, 0x2c, 0xe0, 0x0d, 0x22, 0xf7, 0x86, 0xa2, 0xf3, 0x13, 0xec, 0x6d, 0x0c, 0x21, 0x02,
	0x55, 0x42, 0x69, 0xcc, 0x54, 0x22, 0x60, 0x75, 0xeb, 0xee, 0xaa, 0x84, 0x2e, 0x78, 0x90, 0x01,
	0xbd, 0x89, 0x94, 0xa1, 0x42, 0x25, 0x63, 0xa0, 0x77, 0xb7, 0x81, 0x8c, 0xfc, 0xb3, 0x94, 0xa1,
	0xbb, 0x7b, 0xbe, 0x2e, 0x54, 0xe7, 0x57, 0x05, 0x34, 0x72, 0x53, 0xf8, 0x10, 0x14, 0x39, 0x35,
	0xc2, 0x3b, 0x6e, 0x91, 0x53, 0x08, 0x41, 0x59, 0x90, 0x88, 0xa1, 0xa2, 0xb1, 0x62, 0xce, 0xf0,
	0x05, 0x58, 0x71, 0x78, 0x7a, 0x3e, 0x61, 0xa8, 0x64, 0x66, 0x8d, 0xac, 0x77, 0x32, 0x9f, 0x30,
	0xf8, 0x0e, 0x80, 0x50, 0xfa, 0x67, 0x9e, 0xd2, 0x24, 0xd6, 0xa8, 0xdc, 0xb6, 0xba, 0x8d, 0x83,
	0x26, 0x4e, 0x33, 0xc7, 0xab, 0xcc, 0xf1, 0xc9, 0x2a, 0xf3, 0x41, 0xed, 0xf2, 0xba, 0x55, 0xb8,
	0xf8, 0xdb, 0xb2, 0xdc, 0x7a, 0x82, 0xfb, 0x9a, 0xc0, 0xe0, 0x5b, 0x50, 0x33, 0x24, 0x4c, 0x50,
	0xb4, 0x73, 0x0f, 0x8a, 0x6a, 0x82, 0x7a, 0x2f, 0x28, 0xfc, 0x00, 0x2a, 0x89, 0x29, 0x46, 0x51,
	0x25, 0xb1, 0x38, 0xc0, 0xc9, 0xca, 0x9f, 0xeb, 0xd6, 0xcb, 0x80, 0xeb, 0xf1, 0x74, 0x88, 0x7d,
	0x19, 0x39, 0xbe, 0x54, 0x91, 0x54, 0xd9, 0xa3, 0xa7, 0xe8, 0x99, 0x93, 0xbc, 0x93, 0xc2, 0x1f,
	0x85, 0x76, 0x33, 0x34, 0x3c, 0x02, 0xf5, 0x19, 0xd7, 0x63, 0x1a, 0x93, 0x99, 0x40, 0xd5, 0xad,
	0xa8, 0xd6, 0x04, 0x70, 0x00, 0xca, 0x8a, 0x09, 0x8d, 0x6a, 0x5b, 0x11, 0x19, 0x2c, 0xfc, 0x02,
	0x1e, 0x85, 0x44, 0x69, 0x2f, 0x92, 0x94, 0x8f, 0xb8, 0x4f, 0x34, 0x97, 0x02, 0xd5, 0xef, 0x91,
	0xd1, 0x7e, 0x02, 0x3f, 0xce, 0xa1, 0xe1, 0x18, 0xa0, 0x5b, 0x94, 0x5e, 0x16, 0x1f, 0xd8, 0xca,
	0xea, 0xd3, 0x4d, 0x8d, 0xd3, 0x34, 0x4e, 0x01, 0x9e, 0xdf, 0x56, 0x5a, 0x07, 0xdc, 0xd8, 0x4a,
	0xec, 0xd9, 0xa6, 0xd8, 0xb7, 0x9b, 0xc0, 0x5f, 0x81, 0x7d, 0x1d, 0x13, 0xa1, 0x46, 0x2c, 0xf6,
	0x48, 0x18, 0xca, 0x19, 0xa3, 0x68, 0xb7, 0x6d, 0x75, 0x6b, 0xee, 0xde, 0xaa, 0x7f, 0x98, 0xb6,
	0x07, 0x9f, 0x2e, 0x17, 0xb6, 0x75, 0xb5, 0xb0, 0xad, 0x7f, 0x0b, 0xdb, 0xba, 0x58, 0xda, 0x85,
	0xab, 0xa5, 0x5d, 0xf8, 0xbd, 0xb4, 0x0b, 0xdf, 0xfb, 0x79, 0x1f, 0xb9, 0xff, 0xcd, 0xf1, 0xfb,
	0xac, 0x67, 0x1a, 0xce, 0x0f, 0x27, 0x77, 0x5d, 0x19, 0x67, 0xc3, 0x8a, 0xf9, 0x08, 0x6f, 0xfe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x78, 0x36, 0x74, 0x53, 0xc9, 0x04, 0x00, 0x00,
}

func (m *AccountVestingsList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountVestingsList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountVestingsList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vestings) > 0 {
		for iNdEx := len(m.Vestings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vestings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AccountVestings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountVestings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountVestings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPools) > 0 {
		for iNdEx := len(m.VestingPools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccountVesting(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VestingPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TransferAllowed {
		i--
		if m.TransferAllowed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.LastModificationWithdrawn.Size()
		i -= size
		if _, err := m.LastModificationWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVesting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.LastModificationVested.Size()
		i -= size
		if _, err := m.LastModificationVested.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVesting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastModification, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastModification):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAccountVesting(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	{
		size := m.Sent.Size()
		i -= size
		if _, err := m.Sent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVesting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.Withdrawn.Size()
		i -= size
		if _, err := m.Withdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVesting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Vested.Size()
		i -= size
		if _, err := m.Vested.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVesting(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LockEnd, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LockEnd):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAccountVesting(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LockStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LockStart):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintAccountVesting(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if len(m.VestingType) > 0 {
		i -= len(m.VestingType)
		copy(dAtA[i:], m.VestingType)
		i = encodeVarintAccountVesting(dAtA, i, uint64(len(m.VestingType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAccountVesting(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAccountVesting(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountVestingsList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Vestings) > 0 {
		for _, e := range m.Vestings {
			l = e.Size()
			n += 1 + l + sovAccountVesting(uint64(l))
		}
	}
	return n
}

func (m *AccountVestings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccountVesting(uint64(l))
	}
	if len(m.VestingPools) > 0 {
		for _, e := range m.VestingPools {
			l = e.Size()
			n += 1 + l + sovAccountVesting(uint64(l))
		}
	}
	return n
}

func (m *VestingPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAccountVesting(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAccountVesting(uint64(l))
	}
	l = len(m.VestingType)
	if l > 0 {
		n += 1 + l + sovAccountVesting(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LockStart)
	n += 1 + l + sovAccountVesting(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LockEnd)
	n += 1 + l + sovAccountVesting(uint64(l))
	l = m.Vested.Size()
	n += 1 + l + sovAccountVesting(uint64(l))
	l = m.Withdrawn.Size()
	n += 1 + l + sovAccountVesting(uint64(l))
	l = m.Sent.Size()
	n += 1 + l + sovAccountVesting(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastModification)
	n += 1 + l + sovAccountVesting(uint64(l))
	l = m.LastModificationVested.Size()
	n += 1 + l + sovAccountVesting(uint64(l))
	l = m.LastModificationWithdrawn.Size()
	n += 1 + l + sovAccountVesting(uint64(l))
	if m.TransferAllowed {
		n += 2
	}
	return n
}

func sovAccountVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountVesting(x uint64) (n int) {
	return sovAccountVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountVestingsList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountVesting
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
			return fmt.Errorf("proto: AccountVestingsList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountVestingsList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vestings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vestings = append(m.Vestings, &AccountVestings{})
			if err := m.Vestings[len(m.Vestings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountVesting
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
func (m *AccountVestings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountVesting
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
			return fmt.Errorf("proto: AccountVestings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountVestings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPools = append(m.VestingPools, &VestingPool{})
			if err := m.VestingPools[len(m.VestingPools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountVesting
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
func (m *VestingPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountVesting
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
			return fmt.Errorf("proto: VestingPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LockStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LockEnd, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vested", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Vested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Withdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModification", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastModification, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModificationVested", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastModificationVested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModificationWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
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
				return ErrInvalidLengthAccountVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastModificationWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferAllowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TransferAllowed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAccountVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountVesting
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
func skipAccountVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountVesting
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
					return 0, ErrIntOverflowAccountVesting
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
					return 0, ErrIntOverflowAccountVesting
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
				return 0, ErrInvalidLengthAccountVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountVesting = fmt.Errorf("proto: unexpected end of group")
)
