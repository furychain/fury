// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlin/strategicreserve/stats.proto

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

// OrderBookStats holds statistics on the order book.
type OrderBookStats struct {
	// resolved_unsettled is the list of order book universal unique identifiers
	// that needs to be settled.
	ResolvedUnsettled []string `protobuf:"bytes,1,rep,name=resolved_unsettled,json=resolvedUnsettled,proto3" json:"resolved_unsettled,omitempty"`
}

func (m *OrderBookStats) Reset()         { *m = OrderBookStats{} }
func (m *OrderBookStats) String() string { return proto.CompactTextString(m) }
func (*OrderBookStats) ProtoMessage()    {}
func (*OrderBookStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_392a01451260450d, []int{0}
}
func (m *OrderBookStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBookStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBookStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBookStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBookStats.Merge(m, src)
}
func (m *OrderBookStats) XXX_Size() int {
	return m.Size()
}
func (m *OrderBookStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBookStats.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBookStats proto.InternalMessageInfo

func (m *OrderBookStats) GetResolvedUnsettled() []string {
	if m != nil {
		return m.ResolvedUnsettled
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderBookStats)(nil), "merlinnetwork.merlin.strategicreserve.OrderBookStats")
}

func init() { proto.RegisterFile("merlin/strategicreserve/stats.proto", fileDescriptor_392a01451260450d) }

var fileDescriptor_392a01451260450d = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0x4f, 0xd5,
	0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xcf, 0x4c, 0x2e, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a, 0x03,
	0x09, 0x24, 0x96, 0x14, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0x17, 0xa7, 0xa7, 0xe6,
	0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0xeb, 0x15, 0xa7, 0xa7, 0xea, 0xa1, 0x2b, 0x56, 0xb2, 0xe7,
	0xe2, 0xf3, 0x2f, 0x4a, 0x49, 0x2d, 0x72, 0xca, 0xcf, 0xcf, 0x0e, 0x06, 0x69, 0x14, 0xd2, 0xe5,
	0x12, 0x2a, 0x4a, 0x2d, 0xce, 0xcf, 0x29, 0x4b, 0x4d, 0x89, 0x2f, 0xcd, 0x2b, 0x4e, 0x2d, 0x29,
	0xc9, 0x49, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x12, 0x84, 0xc9, 0x84, 0xc2, 0x24,
	0x9c, 0x7c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x38, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0x3d, 0x55, 0x17, 0xea, 0x0e, 0x10,
	0x5b, 0xbf, 0x02, 0xd3, 0xd9, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x77, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x53, 0x81, 0x9e, 0x43, 0xdb, 0x00, 0x00, 0x00,
}

func (m *OrderBookStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBookStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBookStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResolvedUnsettled) > 0 {
		for iNdEx := len(m.ResolvedUnsettled) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResolvedUnsettled[iNdEx])
			copy(dAtA[i:], m.ResolvedUnsettled[iNdEx])
			i = encodeVarintStats(dAtA, i, uint64(len(m.ResolvedUnsettled[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderBookStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ResolvedUnsettled) > 0 {
		for _, s := range m.ResolvedUnsettled {
			l = len(s)
			n += 1 + l + sovStats(uint64(l))
		}
	}
	return n
}

func sovStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBookStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: OrderBookStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBookStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolvedUnsettled", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResolvedUnsettled = append(m.ResolvedUnsettled, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
				return 0, ErrInvalidLengthStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStats = fmt.Errorf("proto: unexpected end of group")
)
