// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fetchprice/v1/packet.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// FetchPricePacketData defines a struct for the packet payload
// to fetch the oracle prices.
type FetchPricePacketData struct {
	// The currency id is the string with "BASE/QUOTE" format.
	CurrencyIds []string `protobuf:"bytes,1,rep,name=currency_ids,json=currencyIds,proto3" json:"currency_ids,omitempty" yaml:"currency_ids"`
	// optional memo field for future use
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty" yaml:"memo"`
}

func (m *FetchPricePacketData) Reset()         { *m = FetchPricePacketData{} }
func (m *FetchPricePacketData) String() string { return proto.CompactTextString(m) }
func (*FetchPricePacketData) ProtoMessage()    {}
func (*FetchPricePacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1e61fa9a3dc8fa, []int{0}
}
func (m *FetchPricePacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPricePacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPricePacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPricePacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPricePacketData.Merge(m, src)
}
func (m *FetchPricePacketData) XXX_Size() int {
	return m.Size()
}
func (m *FetchPricePacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPricePacketData.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPricePacketData proto.InternalMessageInfo

func (m *FetchPricePacketData) GetCurrencyIds() []string {
	if m != nil {
		return m.CurrencyIds
	}
	return nil
}

func (m *FetchPricePacketData) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// FetchPriceAckData defines a sturct for the ack payload
// of FetchPricePacket.
type FetchPriceAckData struct {
	CurrencyPrices []CurrencyPrice `protobuf:"bytes,1,rep,name=currency_prices,json=currencyPrices,proto3" json:"currency_prices" yaml:"currency_prices"`
}

func (m *FetchPriceAckData) Reset()         { *m = FetchPriceAckData{} }
func (m *FetchPriceAckData) String() string { return proto.CompactTextString(m) }
func (*FetchPriceAckData) ProtoMessage()    {}
func (*FetchPriceAckData) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1e61fa9a3dc8fa, []int{1}
}
func (m *FetchPriceAckData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceAckData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceAckData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceAckData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceAckData.Merge(m, src)
}
func (m *FetchPriceAckData) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceAckData) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceAckData.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceAckData proto.InternalMessageInfo

func (m *FetchPriceAckData) GetCurrencyPrices() []CurrencyPrice {
	if m != nil {
		return m.CurrencyPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchPricePacketData)(nil), "ibc.applications.fetchprice.v1.FetchPricePacketData")
	proto.RegisterType((*FetchPriceAckData)(nil), "ibc.applications.fetchprice.v1.FetchPriceAckData")
}

func init() {
	proto.RegisterFile("ibc/applications/fetchprice/v1/packet.proto", fileDescriptor_cb1e61fa9a3dc8fa)
}

var fileDescriptor_cb1e61fa9a3dc8fa = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xce, 0x4c, 0x4a, 0xd6,
	0x4f, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0x4b, 0x2d,
	0x49, 0xce, 0x28, 0x28, 0xca, 0x4c, 0x4e, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x4c, 0xce, 0x4e,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcb, 0x4c, 0x4a, 0xd6, 0x43, 0x56, 0xac,
	0x87, 0x50, 0xac, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xaa, 0x0f, 0x62,
	0x41, 0x74, 0x49, 0x69, 0x11, 0xb0, 0xa2, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0xa2, 0x56, 0xa9, 0x9c,
	0x4b, 0xc4, 0x0d, 0x24, 0x19, 0x00, 0x92, 0x0c, 0x00, 0xdb, 0xed, 0x92, 0x58, 0x92, 0x28, 0x64,
	0xc5, 0xc5, 0x93, 0x5c, 0x5a, 0x54, 0x94, 0x9a, 0x97, 0x5c, 0x19, 0x9f, 0x99, 0x52, 0x2c, 0xc1,
	0xa8, 0xc0, 0xac, 0xc1, 0xe9, 0x24, 0xfe, 0xe9, 0x9e, 0xbc, 0x70, 0x65, 0x62, 0x6e, 0x8e, 0x95,
	0x12, 0xb2, 0xac, 0x52, 0x10, 0x37, 0x8c, 0xeb, 0x99, 0x52, 0x2c, 0xa4, 0xcc, 0xc5, 0x92, 0x9b,
	0x9a, 0x9b, 0x2f, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xff, 0xe9, 0x9e, 0x3c, 0x37, 0x44,
	0x0f, 0x48, 0x54, 0x29, 0x08, 0x2c, 0xa9, 0xd4, 0xcd, 0xc8, 0x25, 0x88, 0xb0, 0xd9, 0x31, 0x39,
	0x1b, 0x6c, 0x6d, 0x19, 0x17, 0x3f, 0xdc, 0x60, 0xb0, 0x7b, 0x21, 0x36, 0x73, 0x1b, 0xe9, 0xea,
	0xe1, 0x0f, 0x0a, 0x3d, 0x67, 0xa8, 0x36, 0xb0, 0x71, 0x4e, 0x72, 0x27, 0xee, 0xc9, 0x33, 0x7c,
	0xba, 0x27, 0x2f, 0x86, 0xe6, 0x58, 0x88, 0x99, 0x4a, 0x41, 0x7c, 0xc9, 0xc8, 0xca, 0x8b, 0x9d,
	0xfc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x24, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x33, 0x2f, 0xb3, 0x24, 0x33, 0x51, 0x37, 0x27,
	0x31, 0xa9, 0x18, 0xca, 0xd6, 0xaf, 0xd0, 0x07, 0x05, 0x36, 0x52, 0xf8, 0x82, 0x03, 0x37, 0x89,
	0x0d, 0x1c, 0xba, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xe5, 0xf9, 0x02, 0xee, 0x01,
	0x00, 0x00,
}

func (m *FetchPricePacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPricePacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPricePacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CurrencyIds) > 0 {
		for iNdEx := len(m.CurrencyIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CurrencyIds[iNdEx])
			copy(dAtA[i:], m.CurrencyIds[iNdEx])
			i = encodeVarintPacket(dAtA, i, uint64(len(m.CurrencyIds[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FetchPriceAckData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceAckData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceAckData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CurrencyPrices) > 0 {
		for iNdEx := len(m.CurrencyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrencyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FetchPricePacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CurrencyIds) > 0 {
		for _, s := range m.CurrencyIds {
			l = len(s)
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *FetchPriceAckData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CurrencyPrices) > 0 {
		for _, e := range m.CurrencyPrices {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FetchPricePacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: FetchPricePacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPricePacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyIds = append(m.CurrencyIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *FetchPriceAckData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: FetchPriceAckData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceAckData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyPrices = append(m.CurrencyPrices, CurrencyPrice{})
			if err := m.CurrencyPrices[len(m.CurrencyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)