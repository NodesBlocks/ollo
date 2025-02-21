// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/prices.proto

package types

import (
	fmt "fmt"
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

type PricesCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *PricesCallData) Reset()         { *m = PricesCallData{} }
func (m *PricesCallData) String() string { return proto.CompactTextString(m) }
func (*PricesCallData) ProtoMessage()    {}
func (*PricesCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e5ddc2c96b666f0, []int{0}
}
func (m *PricesCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PricesCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PricesCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PricesCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricesCallData.Merge(m, src)
}
func (m *PricesCallData) XXX_Size() int {
	return m.Size()
}
func (m *PricesCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_PricesCallData.DiscardUnknown(m)
}

var xxx_messageInfo_PricesCallData proto.InternalMessageInfo

func (m *PricesCallData) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *PricesCallData) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type PricesResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty"`
}

func (m *PricesResult) Reset()         { *m = PricesResult{} }
func (m *PricesResult) String() string { return proto.CompactTextString(m) }
func (*PricesResult) ProtoMessage()    {}
func (*PricesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e5ddc2c96b666f0, []int{1}
}
func (m *PricesResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PricesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PricesResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PricesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricesResult.Merge(m, src)
}
func (m *PricesResult) XXX_Size() int {
	return m.Size()
}
func (m *PricesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PricesResult.DiscardUnknown(m)
}

var xxx_messageInfo_PricesResult proto.InternalMessageInfo

func (m *PricesResult) GetRates() []uint64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*PricesCallData)(nil), "ollo.oracle.PricesCallData")
	proto.RegisterType((*PricesResult)(nil), "ollo.oracle.PricesResult")
}

func init() { proto.RegisterFile("oracle/prices.proto", fileDescriptor_5e5ddc2c96b666f0) }

var fileDescriptor_5e5ddc2c96b666f0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0xce, 0xcf, 0xc9, 0xc9, 0xd7, 0x83, 0xc8, 0x28, 0x79, 0x71, 0xf1, 0x05, 0x80, 0x25, 0x9d,
	0x13, 0x73, 0x72, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0x2b, 0x73, 0x93, 0xf2,
	0x73, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x39, 0x2e, 0xae, 0xdc,
	0xd2, 0x9c, 0x92, 0xcc, 0x82, 0x9c, 0xcc, 0xd4, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20,
	0x24, 0x11, 0x25, 0x15, 0x2e, 0x1e, 0x88, 0x59, 0x41, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x42, 0x22,
	0x5c, 0xac, 0x45, 0x89, 0x25, 0xa9, 0x10, 0x73, 0x58, 0x82, 0x20, 0x1c, 0x27, 0xdd, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x06, 0x39, 0x4c, 0xbf, 0x42, 0x1f, 0xea,
	0xe8, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xa3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x51, 0x3f, 0x66, 0xdb, 0xcb, 0x00, 0x00, 0x00,
}

func (m *PricesCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PricesCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PricesCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintPrices(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintPrices(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PricesResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PricesResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PricesResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rates) > 0 {
		dAtA2 := make([]byte, len(m.Rates)*10)
		var j1 int
		for _, num := range m.Rates {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintPrices(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrices(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrices(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PricesCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovPrices(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovPrices(uint64(m.Multiplier))
	}
	return n
}

func (m *PricesResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovPrices(uint64(e))
		}
		n += 1 + sovPrices(uint64(l)) + l
	}
	return n
}

func sovPrices(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrices(x uint64) (n int) {
	return sovPrices(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PricesCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: PricesCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PricesCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
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
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func (m *PricesResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: PricesResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PricesResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPrices
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rates = append(m.Rates, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPrices
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPrices
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPrices
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Rates) == 0 {
					m.Rates = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPrices
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rates = append(m.Rates, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func skipPrices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrices
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
					return 0, ErrIntOverflowPrices
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
					return 0, ErrIntOverflowPrices
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
				return 0, ErrInvalidLengthPrices
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrices
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrices
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrices        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrices          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrices = fmt.Errorf("proto: unexpected end of group")
)
