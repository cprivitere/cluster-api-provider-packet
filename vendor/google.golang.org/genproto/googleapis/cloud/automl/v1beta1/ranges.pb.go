// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/ranges.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A range between two double numbers.
type DoubleRange struct {
	// Start of the range, inclusive.
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// End of the range, exclusive.
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a166e9960f6d1167, []int{0}
}

func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*DoubleRange)(nil), "google.cloud.automl.v1beta1.DoubleRange")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/ranges.proto", fileDescriptor_a166e9960f6d1167)
}

var fileDescriptor_a166e9960f6d1167 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xc9, 0x8a, 0x16, 0xb9, 0x46, 0x16, 0x8b, 0xc3, 0x13, 0x15, 0xab, 0xab, 0x12, 0x0e,
	0xb1, 0x89, 0xd5, 0x9d, 0x82, 0x8d, 0xc2, 0x71, 0x85, 0x85, 0x6c, 0x33, 0x7b, 0x17, 0xc2, 0x42,
	0x2e, 0xb3, 0x6c, 0x26, 0x3e, 0x98, 0x8f, 0xe1, 0xa3, 0xf8, 0x14, 0x92, 0x49, 0x4a, 0xb1, 0xcb,
	0xe4, 0xfb, 0xf2, 0xff, 0x19, 0xb9, 0x74, 0x88, 0xce, 0x5b, 0xbd, 0xf7, 0x98, 0x0e, 0x1a, 0x12,
	0xe1, 0xd1, 0xeb, 0xcf, 0x55, 0x6f, 0x09, 0x56, 0x7a, 0x82, 0xe0, 0x6c, 0x54, 0xe3, 0x84, 0x84,
	0xed, 0xa2, 0x98, 0x8a, 0x4d, 0x55, 0x4c, 0x55, 0xcd, 0xcb, 0xab, 0x1a, 0x03, 0xe3, 0xa0, 0x21,
	0x04, 0x24, 0xa0, 0x01, 0x43, 0x7d, 0x7a, 0xf7, 0x20, 0x67, 0xcf, 0x98, 0x7a, 0x6f, 0x77, 0x39,
	0xb0, 0xbd, 0x90, 0xa7, 0x91, 0x60, 0xa2, 0xb9, 0xb8, 0x15, 0x4b, 0xb1, 0x2b, 0x43, 0x7b, 0x2e,
	0x4f, 0x6c, 0x38, 0xcc, 0x1b, 0xbe, 0xcb, 0xc7, 0xcd, 0x97, 0x90, 0x37, 0x7b, 0x3c, 0xaa, 0x7f,
	0x8a, 0x37, 0x33, 0x8e, 0x8c, 0xdb, 0xdc, 0xb3, 0x15, 0x1f, 0xeb, 0xea, 0x3a, 0xf4, 0x10, 0x9c,
	0xc2, 0xc9, 0x69, 0x67, 0x03, 0xff, 0x42, 0x17, 0x04, 0xe3, 0x10, 0xff, 0xdc, 0xf6, 0xb1, 0x8c,
	0xdf, 0xcd, 0xe2, 0x85, 0xc5, 0xee, 0x29, 0x4b, 0xdd, 0x3a, 0x11, 0xbe, 0xf9, 0xee, 0xbd, 0x48,
	0x3f, 0xcd, 0x75, 0xa1, 0xc6, 0x30, 0x36, 0x86, 0xf9, 0xab, 0x31, 0x55, 0xe8, 0xcf, 0xb8, 0xec,
	0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x05, 0xcb, 0x11, 0xcf, 0x59, 0x01, 0x00, 0x00,
}