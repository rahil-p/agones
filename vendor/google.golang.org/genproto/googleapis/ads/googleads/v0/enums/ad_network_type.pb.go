// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_network_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enumerates Google Ads network types.
type AdNetworkTypeEnum_AdNetworkType int32

const (
	// Not specified.
	AdNetworkTypeEnum_UNSPECIFIED AdNetworkTypeEnum_AdNetworkType = 0
	// The value is unknown in this version.
	AdNetworkTypeEnum_UNKNOWN AdNetworkTypeEnum_AdNetworkType = 1
	// Google search.
	AdNetworkTypeEnum_SEARCH AdNetworkTypeEnum_AdNetworkType = 2
	// Search partners.
	AdNetworkTypeEnum_SEARCH_PARTNERS AdNetworkTypeEnum_AdNetworkType = 3
	// Display Network.
	AdNetworkTypeEnum_CONTENT AdNetworkTypeEnum_AdNetworkType = 4
	// YouTube Search.
	AdNetworkTypeEnum_YOUTUBE_SEARCH AdNetworkTypeEnum_AdNetworkType = 5
	// YouTube Videos
	AdNetworkTypeEnum_YOUTUBE_WATCH AdNetworkTypeEnum_AdNetworkType = 6
	// Cross-network.
	AdNetworkTypeEnum_MIXED AdNetworkTypeEnum_AdNetworkType = 7
)

var AdNetworkTypeEnum_AdNetworkType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH",
	3: "SEARCH_PARTNERS",
	4: "CONTENT",
	5: "YOUTUBE_SEARCH",
	6: "YOUTUBE_WATCH",
	7: "MIXED",
}
var AdNetworkTypeEnum_AdNetworkType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"SEARCH":          2,
	"SEARCH_PARTNERS": 3,
	"CONTENT":         4,
	"YOUTUBE_SEARCH":  5,
	"YOUTUBE_WATCH":   6,
	"MIXED":           7,
}

func (x AdNetworkTypeEnum_AdNetworkType) String() string {
	return proto.EnumName(AdNetworkTypeEnum_AdNetworkType_name, int32(x))
}
func (AdNetworkTypeEnum_AdNetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_network_type_8e7c5c4b59c79566, []int{0, 0}
}

// Container for enumeration of Google Ads network types.
type AdNetworkTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdNetworkTypeEnum) Reset()         { *m = AdNetworkTypeEnum{} }
func (m *AdNetworkTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdNetworkTypeEnum) ProtoMessage()    {}
func (*AdNetworkTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_network_type_8e7c5c4b59c79566, []int{0}
}
func (m *AdNetworkTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdNetworkTypeEnum.Unmarshal(m, b)
}
func (m *AdNetworkTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdNetworkTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdNetworkTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdNetworkTypeEnum.Merge(dst, src)
}
func (m *AdNetworkTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdNetworkTypeEnum.Size(m)
}
func (m *AdNetworkTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdNetworkTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdNetworkTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdNetworkTypeEnum)(nil), "google.ads.googleads.v0.enums.AdNetworkTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdNetworkTypeEnum_AdNetworkType", AdNetworkTypeEnum_AdNetworkType_name, AdNetworkTypeEnum_AdNetworkType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_network_type.proto", fileDescriptor_ad_network_type_8e7c5c4b59c79566)
}

var fileDescriptor_ad_network_type_8e7c5c4b59c79566 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4e, 0xf2, 0x30,
	0x14, 0xc7, 0xbf, 0xc1, 0x07, 0xc4, 0x43, 0x90, 0x52, 0xaf, 0xb9, 0x90, 0x07, 0xe8, 0x96, 0x70,
	0xe9, 0x55, 0x37, 0x2a, 0x10, 0x63, 0x59, 0x60, 0x03, 0x35, 0x4b, 0x96, 0xe9, 0x96, 0xc6, 0x08,
	0xeb, 0x42, 0x01, 0xc3, 0x43, 0xf8, 0x12, 0x7a, 0xe7, 0x6b, 0x78, 0xe7, 0x53, 0x99, 0xad, 0x8c,
	0x84, 0x0b, 0xbd, 0x69, 0x7e, 0x39, 0xe7, 0xfc, 0xda, 0x9e, 0x3f, 0xf4, 0x85, 0x94, 0x62, 0x99,
	0x98, 0x51, 0xac, 0x4c, 0x8d, 0x39, 0xed, 0x2c, 0x33, 0x49, 0xb7, 0x2b, 0x65, 0x46, 0x71, 0x98,
	0x26, 0x9b, 0x57, 0xb9, 0x7e, 0x09, 0x37, 0xfb, 0x2c, 0x21, 0xd9, 0x5a, 0x6e, 0x24, 0xee, 0xea,
	0x49, 0x12, 0xc5, 0x8a, 0x1c, 0x25, 0xb2, 0xb3, 0x48, 0x21, 0xf5, 0x3e, 0x0c, 0xe8, 0xd0, 0x98,
	0x6b, 0xcf, 0xdb, 0x67, 0x09, 0x4b, 0xb7, 0xab, 0xde, 0x9b, 0x01, 0xad, 0x93, 0x2a, 0x6e, 0x43,
	0xd3, 0xe7, 0x33, 0x97, 0x39, 0xe3, 0xeb, 0x31, 0x1b, 0xa0, 0x7f, 0xb8, 0x09, 0x0d, 0x9f, 0xdf,
	0xf0, 0xc9, 0x82, 0x23, 0x03, 0x03, 0xd4, 0x67, 0x8c, 0x4e, 0x9d, 0x11, 0xaa, 0xe0, 0x0b, 0x68,
	0x6b, 0x0e, 0x5d, 0x3a, 0xf5, 0x38, 0x9b, 0xce, 0x50, 0x35, 0x9f, 0x76, 0x26, 0xdc, 0x63, 0xdc,
	0x43, 0xff, 0x31, 0x86, 0xf3, 0xfb, 0x89, 0xef, 0xf9, 0x36, 0x0b, 0x0f, 0x56, 0x0d, 0x77, 0xa0,
	0x55, 0xd6, 0x16, 0xd4, 0x73, 0x46, 0xa8, 0x8e, 0xcf, 0xa0, 0x76, 0x3b, 0xbe, 0x63, 0x03, 0xd4,
	0xb0, 0xbf, 0x0c, 0xb8, 0x7c, 0x92, 0x2b, 0xf2, 0xe7, 0x2e, 0x36, 0x3e, 0xf9, 0xb2, 0x9b, 0xaf,
	0xef, 0x1a, 0x0f, 0xf6, 0x41, 0x12, 0x72, 0x19, 0xa5, 0x82, 0xc8, 0xb5, 0x30, 0x45, 0x92, 0x16,
	0xe1, 0x94, 0x29, 0x66, 0xcf, 0xea, 0x97, 0x50, 0xaf, 0x8a, 0xf3, 0xbd, 0x52, 0x1d, 0x52, 0xfa,
	0x59, 0xe9, 0x0e, 0xf5, 0x55, 0x34, 0x56, 0x44, 0x63, 0x4e, 0x73, 0x8b, 0xe4, 0xa1, 0xa9, 0xef,
	0xb2, 0x1f, 0xd0, 0x58, 0x05, 0xc7, 0x7e, 0x30, 0xb7, 0x82, 0xa2, 0xff, 0x58, 0x2f, 0x1e, 0xed,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xe6, 0xe0, 0x73, 0xc8, 0x01, 0x00, 0x00,
}