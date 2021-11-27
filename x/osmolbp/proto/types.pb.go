// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/osmolbp/v1/types.proto

package proto

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
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

type LBP struct {
	// token_out a token denom to be bootstraped. May be referred as base currency.
	TokenOut string `protobuf:"bytes,1,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
	// token_in a denom used to buy LB tokens (`token_out`). May be referred as quote_currency.
	TokenIn string `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// start time when the token emission starts.
	Start time.Time `protobuf:"bytes,3,opt,name=start,proto3,stdtime" json:"start"`
	// end  time when the token emission ends.
	End time.Time `protobuf:"bytes,4,opt,name=end,proto3,stdtime" json:"end"`
	// number of `tokens_out` to be emitted per second. 1 second = 1 round.
	Rate github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"rate"`
	// Running sum of sold coins (token_out)
	Accumulator github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=accumulator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"accumulator"`
	// Round number when the accumulator was updated.
	AccumulatorR uint64 `protobuf:"varint,7,opt,name=accumulator_r,json=accumulatorR,proto3" json:"accumulator_r,omitempty"`
	// total number of currently staked coins (token_in)
	Staked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked"`
	// total amount of earned coins (token_in)
	Income github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=income,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"income"`
}

func (m *LBP) Reset()         { *m = LBP{} }
func (m *LBP) String() string { return proto.CompactTextString(m) }
func (*LBP) ProtoMessage()    {}
func (*LBP) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4260a358f41a501, []int{0}
}
func (m *LBP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LBP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LBP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LBP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LBP.Merge(m, src)
}
func (m *LBP) XXX_Size() int {
	return m.Size()
}
func (m *LBP) XXX_DiscardUnknown() {
	xxx_messageInfo_LBP.DiscardUnknown(m)
}

var xxx_messageInfo_LBP proto.InternalMessageInfo

// UserPosition represents user account in a pool
type UserPosition struct {
	// lbp.accumulator value when the last ping was called and rewards calculated
	Accumulator github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=accumulator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"accumulator"`
	// total number of currently staked tokens
	Staked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked"`
	// Amount of accumulated, not withdrawn purchased tokens (token_out)
	Purchased github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=purchased,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"purchased"`
	// total amount of spent coins (token_in)
	Spent github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=spent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"spent"`
}

func (m *UserPosition) Reset()         { *m = UserPosition{} }
func (m *UserPosition) String() string { return proto.CompactTextString(m) }
func (*UserPosition) ProtoMessage()    {}
func (*UserPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4260a358f41a501, []int{1}
}
func (m *UserPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPosition.Merge(m, src)
}
func (m *UserPosition) XXX_Size() int {
	return m.Size()
}
func (m *UserPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPosition.DiscardUnknown(m)
}

var xxx_messageInfo_UserPosition proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LBP)(nil), "osmosis.osmolbp.v1.LBP")
	proto.RegisterType((*UserPosition)(nil), "osmosis.osmolbp.v1.UserPosition")
}

func init() { proto.RegisterFile("osmosis/osmolbp/v1/types.proto", fileDescriptor_c4260a358f41a501) }

var fileDescriptor_c4260a358f41a501 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x5e, 0xe7, 0xaf, 0x89, 0x5b, 0x2e, 0x16, 0x07, 0x37, 0x48, 0x9b, 0xa8, 0x48, 0x28, 0x97,
	0xda, 0x2a, 0x48, 0x1c, 0x38, 0xae, 0x10, 0x52, 0xa5, 0x22, 0xa2, 0x15, 0x5c, 0xb8, 0x44, 0xce,
	0xc6, 0x6c, 0x57, 0xc9, 0xda, 0xab, 0xf5, 0x6c, 0x05, 0x6f, 0xd1, 0x07, 0xe2, 0x01, 0x72, 0xec,
	0x11, 0x71, 0x28, 0x90, 0xbc, 0x06, 0x07, 0xb4, 0xb6, 0x43, 0xa3, 0xde, 0x92, 0x9c, 0xc6, 0xdf,
	0x7c, 0xfb, 0x7d, 0x3b, 0x33, 0x9a, 0xc1, 0xa1, 0x36, 0xb9, 0x36, 0x99, 0xe1, 0x75, 0x5c, 0x4c,
	0x0b, 0x7e, 0x73, 0xc1, 0xe1, 0x5b, 0x21, 0x0d, 0x2b, 0x4a, 0x0d, 0x9a, 0x10, 0xcf, 0x33, 0xcf,
	0xb3, 0x9b, 0x8b, 0xfe, 0x20, 0xd5, 0x3a, 0x5d, 0x48, 0x6e, 0xbf, 0x98, 0x56, 0x5f, 0x38, 0x64,
	0xb9, 0x34, 0x20, 0xf2, 0xc2, 0x89, 0xfa, 0xa7, 0x89, 0x55, 0x4d, 0x2c, 0xe2, 0x0e, 0x78, 0xea,
	0x69, 0xaa, 0x53, 0xed, 0xf2, 0xf5, 0xcb, 0x65, 0xcf, 0xfe, 0x36, 0x71, 0xf3, 0x2a, 0x1a, 0x93,
	0x67, 0xb8, 0x07, 0x7a, 0x2e, 0xd5, 0x44, 0x57, 0x40, 0xd1, 0x10, 0x8d, 0x7a, 0x71, 0xd7, 0x26,
	0x3e, 0x54, 0x40, 0x4e, 0xb1, 0x7b, 0x4f, 0x32, 0x45, 0x1b, 0x96, 0x3b, 0xb2, 0xf8, 0x52, 0x91,
	0x37, 0xb8, 0x6d, 0x40, 0x94, 0x40, 0x9b, 0x43, 0x34, 0x3a, 0x7e, 0xd9, 0x67, 0xae, 0x42, 0xb6,
	0xa9, 0x90, 0x7d, 0xdc, 0x54, 0x18, 0x75, 0x97, 0xf7, 0x83, 0xe0, 0xf6, 0xd7, 0x00, 0xc5, 0x4e,
	0x42, 0x5e, 0xe3, 0xa6, 0x54, 0x33, 0xda, 0xda, 0x41, 0x59, 0x0b, 0x48, 0x84, 0x5b, 0xa5, 0x00,
	0x49, 0xdb, 0x75, 0x29, 0x11, 0xab, 0xc9, 0x9f, 0xf7, 0x83, 0x17, 0x69, 0x06, 0xd7, 0xd5, 0x94,
	0x25, 0x3a, 0xf7, 0x8d, 0xfb, 0x70, 0x6e, 0x66, 0x73, 0x3f, 0xd9, 0x4b, 0x05, 0xb1, 0xd5, 0x92,
	0x31, 0x3e, 0x16, 0x49, 0x52, 0xe5, 0xd5, 0x42, 0x80, 0x2e, 0x69, 0x67, 0x2f, 0xab, 0x6d, 0x0b,
	0xf2, 0x1c, 0x3f, 0xd9, 0x82, 0x93, 0x92, 0x1e, 0x0d, 0xd1, 0xa8, 0x15, 0x9f, 0x6c, 0x25, 0x63,
	0xf2, 0x0e, 0x77, 0x0c, 0x88, 0xb9, 0x9c, 0xd1, 0xee, 0x5e, 0x7f, 0xf4, 0xea, 0xda, 0x27, 0x53,
	0x89, 0xce, 0x25, 0xed, 0xed, 0xe7, 0xe3, 0xd4, 0x67, 0xdf, 0x1b, 0xf8, 0xe4, 0x93, 0x91, 0xe5,
	0x58, 0x9b, 0x0c, 0x32, 0xad, 0x1e, 0xcf, 0x05, 0x1d, 0x3e, 0x97, 0x87, 0x96, 0x1b, 0x07, 0xb5,
	0x7c, 0x85, 0x7b, 0x45, 0x55, 0x26, 0xd7, 0xc2, 0xc8, 0x99, 0xdd, 0xb6, 0xdd, 0xad, 0x1e, 0x0c,
	0xc8, 0x5b, 0xdc, 0x36, 0x85, 0x54, 0x60, 0xb7, 0x6f, 0x77, 0x27, 0x27, 0x8e, 0xde, 0x2f, 0xff,
	0x84, 0xc1, 0x72, 0x15, 0xa2, 0xbb, 0x55, 0x88, 0x7e, 0xaf, 0x42, 0x74, 0xbb, 0x0e, 0x83, 0xbb,
	0x75, 0x18, 0xfc, 0x58, 0x87, 0xc1, 0x67, 0xbe, 0x65, 0xe6, 0x8f, 0xf9, 0x7c, 0x21, 0xa6, 0x66,
	0x03, 0xf8, 0xd7, 0xff, 0xb7, 0xef, 0x96, 0xbe, 0x63, 0xc3, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x3f, 0x02, 0x6d, 0xba, 0x1b, 0x04, 0x00, 0x00,
}

func (m *LBP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LBP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LBP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Income.Size()
		i -= size
		if _, err := m.Income.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.AccumulatorR != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AccumulatorR))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.Accumulator.Size()
		i -= size
		if _, err := m.Accumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Spent.Size()
		i -= size
		if _, err := m.Spent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Purchased.Size()
		i -= size
		if _, err := m.Purchased.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Accumulator.Size()
		i -= size
		if _, err := m.Accumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LBP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovTypes(uint64(l))
	l = m.Rate.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Accumulator.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.AccumulatorR != 0 {
		n += 1 + sovTypes(uint64(m.AccumulatorR))
	}
	l = m.Staked.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Income.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *UserPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Accumulator.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Staked.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Purchased.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Spent.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LBP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LBP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LBP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Accumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumulatorR", wireType)
			}
			m.AccumulatorR = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccumulatorR |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Income", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Income.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *UserPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: UserPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Accumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purchased", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Purchased.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
