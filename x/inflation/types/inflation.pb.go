// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: echelon/inflation/v1/inflation.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// InflationDistribution defines the distribution in which inflation is
// allocated through minting on each epoch (staking, incentives, community). It
// excludes the team vesting distribution, as this is minted once at genesis.
// The initial InflationDistribution can be calculated from the Evmvos Token
// Model like this:
// mintDistribution1 = distribution1 / (1 - teamVestingDistribution)
// 0.5333333         = 40%           / (1 - 25%)
type InflationDistribution struct {
	// staking_rewards defines the proportion of the minted minted_denom that is
	// to be allocated as staking rewards
	StakingRewards github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=staking_rewards,json=stakingRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_rewards"`
	// usage_incentives defines the proportion of the minted minted_denom that is
	// to be allocated to the incentives module address
	UsageIncentives github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=usage_incentives,json=usageIncentives,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"usage_incentives"`
	// community_pool defines the proportion of the minted minted_denom that is to
	// be allocated to the community pool
	CommunityPool github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=community_pool,json=communityPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_pool"`
}

func (m *InflationDistribution) Reset()         { *m = InflationDistribution{} }
func (m *InflationDistribution) String() string { return proto.CompactTextString(m) }
func (*InflationDistribution) ProtoMessage()    {}
func (*InflationDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_d064cb35c3ff7df8, []int{0}
}
func (m *InflationDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationDistribution.Merge(m, src)
}
func (m *InflationDistribution) XXX_Size() int {
	return m.Size()
}
func (m *InflationDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_InflationDistribution proto.InternalMessageInfo

// ExponentialCalculation holds factors to calculate exponential inflation on
// each period. Calculation reference:
// periodProvision = exponentialDecay       *  bondingIncentive
// f(x)            = (a * (1 - r) ^ x + c)  *  (1 + max_variance - bondedRatio *
// (max_variance / bonding_target))
type ExponentialCalculation struct {
	// initial value
	A github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=a,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"a"`
	// reduction factor
	R github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=r,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"r"`
	// long term inflation
	C github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=c,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"c"`
	// bonding target
	BondingTarget github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=bonding_target,json=bondingTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bonding_target"`
	// max variance
	MaxVariance github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=max_variance,json=maxVariance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_variance"`
}

func (m *ExponentialCalculation) Reset()         { *m = ExponentialCalculation{} }
func (m *ExponentialCalculation) String() string { return proto.CompactTextString(m) }
func (*ExponentialCalculation) ProtoMessage()    {}
func (*ExponentialCalculation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d064cb35c3ff7df8, []int{1}
}
func (m *ExponentialCalculation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExponentialCalculation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExponentialCalculation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExponentialCalculation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExponentialCalculation.Merge(m, src)
}
func (m *ExponentialCalculation) XXX_Size() int {
	return m.Size()
}
func (m *ExponentialCalculation) XXX_DiscardUnknown() {
	xxx_messageInfo_ExponentialCalculation.DiscardUnknown(m)
}

var xxx_messageInfo_ExponentialCalculation proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InflationDistribution)(nil), "echelon.inflation.v1.InflationDistribution")
	proto.RegisterType((*ExponentialCalculation)(nil), "echelon.inflation.v1.ExponentialCalculation")
}

func init() {
	proto.RegisterFile("echelon/inflation/v1/inflation.proto", fileDescriptor_d064cb35c3ff7df8)
}

var fileDescriptor_d064cb35c3ff7df8 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe3, 0x72, 0x91, 0x30, 0xd0, 0xa2, 0x08, 0x50, 0xc4, 0x90, 0xa2, 0x0e, 0x88, 0xa5,
	0x89, 0xaa, 0xae, 0x4c, 0xa5, 0x08, 0x75, 0x83, 0x8a, 0x8b, 0x60, 0x89, 0x1c, 0xd7, 0xa4, 0x56,
	0x13, 0x3b, 0xb2, 0x9d, 0x90, 0xbe, 0x05, 0xcf, 0xc0, 0xd3, 0x74, 0xec, 0x88, 0x18, 0x2a, 0xd4,
	0xbe, 0x06, 0x03, 0x72, 0x52, 0xda, 0xce, 0x99, 0x7c, 0x6c, 0xff, 0xe7, 0xb3, 0xcf, 0x6f, 0x1f,
	0xd8, 0x20, 0x69, 0xc4, 0xa5, 0x4b, 0xd9, 0x5b, 0x88, 0x14, 0xe5, 0xcc, 0x4d, 0x5b, 0xeb, 0x89,
	0x13, 0x0b, 0xae, 0xb8, 0x69, 0xe6, 0x1a, 0x67, 0xbd, 0x9c, 0xb6, 0xce, 0x8e, 0x03, 0x1e, 0xf0,
	0x7c, 0xdb, 0xd5, 0x51, 0xa1, 0x6c, 0x7c, 0x56, 0xe0, 0x49, 0xef, 0x5f, 0xd6, 0xa5, 0x52, 0x09,
	0xea, 0x27, 0x3a, 0x36, 0x9f, 0x61, 0x4d, 0x2a, 0x34, 0xa2, 0x2c, 0xf0, 0x04, 0x79, 0x47, 0x62,
	0x20, 0x2d, 0x70, 0x0e, 0x2e, 0xf7, 0x3a, 0xce, 0x64, 0x56, 0x37, 0xbe, 0x67, 0xf5, 0x8b, 0x80,
	0xaa, 0x61, 0xe2, 0x3b, 0x98, 0x47, 0x2e, 0xe6, 0x52, 0x5f, 0xaa, 0x18, 0x9a, 0x72, 0x30, 0x72,
	0xd5, 0x38, 0x26, 0xd2, 0xe9, 0x12, 0xdc, 0xaf, 0x2e, 0x31, 0xfd, 0x82, 0x62, 0xbe, 0xc0, 0xa3,
	0x44, 0xa2, 0x80, 0x78, 0x94, 0x61, 0xc2, 0x14, 0x4d, 0x89, 0xb4, 0x2a, 0xa5, 0xc8, 0xb5, 0x9c,
	0xd3, 0x5b, 0x61, 0xcc, 0x47, 0x58, 0xc5, 0x3c, 0x8a, 0x12, 0x46, 0xd5, 0xd8, 0x8b, 0x39, 0x0f,
	0xad, 0xad, 0x52, 0xe0, 0xc3, 0x15, 0xe5, 0x8e, 0xf3, 0xb0, 0xf1, 0x5b, 0x81, 0xa7, 0x37, 0x59,
	0xcc, 0x99, 0x3e, 0x07, 0x85, 0xd7, 0x28, 0xc4, 0x49, 0xe1, 0x98, 0x79, 0x05, 0x01, 0x2a, 0xe9,
	0x0b, 0x40, 0x3a, 0x5b, 0x94, 0xac, 0x1d, 0x08, 0x9d, 0x8d, 0x4b, 0x16, 0x08, 0xb0, 0xf6, 0xca,
	0xe7, 0x6c, 0xa0, 0xdf, 0x57, 0x21, 0x11, 0x10, 0x65, 0x6d, 0x97, 0xf3, 0x6a, 0x49, 0x79, 0xc8,
	0x21, 0xe6, 0x3d, 0x3c, 0x88, 0x50, 0xe6, 0xa5, 0x48, 0x50, 0xc4, 0x30, 0xb1, 0x76, 0x4a, 0x41,
	0xf7, 0x23, 0x94, 0x3d, 0x2d, 0x11, 0x9d, 0xdb, 0xc9, 0xdc, 0x06, 0xd3, 0xb9, 0x0d, 0x7e, 0xe6,
	0x36, 0xf8, 0x58, 0xd8, 0xc6, 0x74, 0x61, 0x1b, 0x5f, 0x0b, 0xdb, 0x78, 0x6d, 0x6e, 0xe0, 0xd4,
	0x10, 0x09, 0x49, 0xa5, 0x5b, 0xb4, 0x47, 0xda, 0x76, 0xb3, 0x8d, 0x1e, 0xc9, 0xc9, 0xfe, 0x6e,
	0xfe, 0xe7, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x91, 0x17, 0x7d, 0x43, 0x03, 0x00,
	0x00,
}

func (m *InflationDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UsageIncentives.Size()
		i -= size
		if _, err := m.UsageIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.StakingRewards.Size()
		i -= size
		if _, err := m.StakingRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ExponentialCalculation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExponentialCalculation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExponentialCalculation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxVariance.Size()
		i -= size
		if _, err := m.MaxVariance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.BondingTarget.Size()
		i -= size
		if _, err := m.BondingTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.C.Size()
		i -= size
		if _, err := m.C.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.R.Size()
		i -= size
		if _, err := m.R.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.A.Size()
		i -= size
		if _, err := m.A.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintInflation(dAtA []byte, offset int, v uint64) int {
	offset -= sovInflation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InflationDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StakingRewards.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.UsageIncentives.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func (m *ExponentialCalculation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.A.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.R.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.C.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.BondingTarget.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.MaxVariance.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func sovInflation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInflation(x uint64) (n int) {
	return sovInflation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InflationDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: InflationDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsageIncentives", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UsageIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func (m *ExponentialCalculation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: ExponentialCalculation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExponentialCalculation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.R.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.C.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondingTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondingTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVariance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxVariance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func skipInflation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
				return 0, ErrInvalidLengthInflation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInflation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInflation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInflation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInflation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInflation = fmt.Errorf("proto: unexpected end of group")
)