// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlin/bet/odds_type.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// OddsType is the representation of the type of the odds.
type OddsType int32

const (
	// invalid odds type
	OddsType_ODDS_TYPE_UNSPECIFIED OddsType = 0
	// decimal odds type (european)
	OddsType_ODDS_TYPE_DECIMAL OddsType = 1
	// fractional odds type (british)
	OddsType_ODDS_TYPE_FRACTIONAL OddsType = 2
	// moneyline odds type (american)
	OddsType_ODDS_TYPE_MONEYLINE OddsType = 3
)

var OddsType_name = map[int32]string{
	0: "ODDS_TYPE_UNSPECIFIED",
	1: "ODDS_TYPE_DECIMAL",
	2: "ODDS_TYPE_FRACTIONAL",
	3: "ODDS_TYPE_MONEYLINE",
}

var OddsType_value = map[string]int32{
	"ODDS_TYPE_UNSPECIFIED": 0,
	"ODDS_TYPE_DECIMAL":     1,
	"ODDS_TYPE_FRACTIONAL":  2,
	"ODDS_TYPE_MONEYLINE":   3,
}

func (x OddsType) String() string {
	return proto.EnumName(OddsType_name, int32(x))
}

func (OddsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9054da535c81741c, []int{0}
}

func init() {
	proto.RegisterEnum("merlinnetwork.merlin.bet.OddsType", OddsType_name, OddsType_value)
}

func init() { proto.RegisterFile("merlin/bet/odds_type.proto", fileDescriptor_9054da535c81741c) }

var fileDescriptor_9054da535c81741c = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x4f, 0xd5,
	0x4f, 0x4a, 0x2d, 0xd1, 0xcf, 0x4f, 0x49, 0x29, 0x8e, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6,
	0x2b, 0x4e, 0x4f, 0xd5, 0x4b, 0x4a, 0x2d, 0xd1, 0xca, 0xe7, 0xe2, 0xf0, 0x4f, 0x49, 0x29, 0x0e,
	0xa9, 0x2c, 0x48, 0x15, 0x92, 0xe4, 0x12, 0xf5, 0x77, 0x71, 0x09, 0x8e, 0x0f, 0x89, 0x0c, 0x70,
	0x8d, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0x12,
	0xe5, 0x12, 0x44, 0x48, 0xb9, 0xb8, 0x3a, 0x7b, 0xfa, 0x3a, 0xfa, 0x08, 0x30, 0x0a, 0x49, 0x70,
	0x89, 0x20, 0x84, 0xdd, 0x82, 0x1c, 0x9d, 0x43, 0x3c, 0xfd, 0xfd, 0x1c, 0x7d, 0x04, 0x98, 0x84,
	0xc4, 0xb9, 0x84, 0x11, 0x32, 0xbe, 0xfe, 0x7e, 0xae, 0x91, 0x3e, 0x9e, 0x7e, 0xae, 0x02, 0xcc,
	0x4e, 0x0e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x96, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x9c, 0x9e, 0xaa, 0x0b, 0x75, 0x2a, 0x88,
	0xad, 0x5f, 0x01, 0xf6, 0x10, 0xc8, 0x2f, 0xc5, 0x49, 0x6c, 0x60, 0xdf, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x6e, 0x4e, 0x56, 0xe8, 0x00, 0x00, 0x00,
}
