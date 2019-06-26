// Code generated by protoc-gen-go. DO NOT EDIT.
// source: postgres_arrays.proto

package postgres_arrays // import "github.com/smyrgl/protoc-gen-gorm/example/postgres_arrays"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/smyrgl/protoc-gen-gorm/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Example struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description          string    `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	ArrayOfBools         []bool    `protobuf:"varint,20,rep,packed,name=array_of_bools,json=arrayOfBools" json:"array_of_bools,omitempty"`
	ArrayOfFloat64       []float64 `protobuf:"fixed64,30,rep,packed,name=array_of_float64,json=arrayOfFloat64" json:"array_of_float64,omitempty"`
	ArrayOfInt64         []int64   `protobuf:"varint,40,rep,packed,name=array_of_int64,json=arrayOfInt64" json:"array_of_int64,omitempty"`
	ArrayOfString        []string  `protobuf:"bytes,50,rep,name=array_of_string,json=arrayOfString" json:"array_of_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgres_arrays_b477d8058a1bc73b, []int{0}
}
func (m *Example) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example.Unmarshal(m, b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example.Marshal(b, m, deterministic)
}
func (dst *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(dst, src)
}
func (m *Example) XXX_Size() int {
	return xxx_messageInfo_Example.Size(m)
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Example) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Example) GetArrayOfBools() []bool {
	if m != nil {
		return m.ArrayOfBools
	}
	return nil
}

func (m *Example) GetArrayOfFloat64() []float64 {
	if m != nil {
		return m.ArrayOfFloat64
	}
	return nil
}

func (m *Example) GetArrayOfInt64() []int64 {
	if m != nil {
		return m.ArrayOfInt64
	}
	return nil
}

func (m *Example) GetArrayOfString() []string {
	if m != nil {
		return m.ArrayOfString
	}
	return nil
}

func init() {
	proto.RegisterType((*Example)(nil), "postgres.arrays.Example")
}

func init() {
	proto.RegisterFile("postgres_arrays.proto", fileDescriptor_postgres_arrays_b477d8058a1bc73b)
}

var fileDescriptor_postgres_arrays_b477d8058a1bc73b = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0x80, 0x49, 0xfb, 0x63, 0xbf, 0x2d, 0xea, 0x26, 0x41, 0xa1, 0xf3, 0x30, 0x82, 0x88, 0xe4,
	0xb2, 0x16, 0x54, 0x3c, 0xcc, 0x5b, 0x41, 0xc1, 0xd3, 0xb0, 0xde, 0xbc, 0x94, 0xfe, 0x49, 0x6b,
	0xa0, 0xcd, 0x5b, 0x92, 0x14, 0xe6, 0x47, 0xdb, 0xbe, 0x95, 0xdf, 0x40, 0x9a, 0xce, 0xa2, 0x3b,
	0x79, 0x7c, 0x1f, 0x9e, 0x3c, 0x84, 0xf7, 0xc5, 0xe7, 0x0d, 0x68, 0x53, 0x2a, 0xae, 0xe3, 0x44,
	0xa9, 0xe4, 0x43, 0xfb, 0x8d, 0x02, 0x03, 0x64, 0xf6, 0x8d, 0xfd, 0x1e, 0x5f, 0xac, 0x4a, 0x61,
	0xde, 0xdb, 0xd4, 0xcf, 0xa0, 0x0e, 0x84, 0x2c, 0x20, 0xad, 0x60, 0x03, 0x0d, 0x97, 0x81, 0xf5,
	0xb3, 0x65, 0xc9, 0xe5, 0xb2, 0x04, 0x55, 0x07, 0xd0, 0x18, 0x01, 0x52, 0x07, 0xdd, 0xd0, 0xc7,
	0x2e, 0x3f, 0x11, 0xfe, 0xff, 0xb8, 0x49, 0xea, 0xa6, 0xe2, 0x64, 0x81, 0x1d, 0x91, 0x7b, 0x88,
	0x22, 0x36, 0x09, 0xa7, 0xbb, 0xed, 0x1c, 0xe3, 0x31, 0xf9, 0xd7, 0xb6, 0x22, 0x67, 0x28, 0x72,
	0x44, 0x4e, 0x28, 0x3e, 0xca, 0xb9, 0xce, 0x94, 0xb0, 0x19, 0xcf, 0xe9, 0xc4, 0xe8, 0x27, 0x22,
	0x57, 0x78, 0x6a, 0xff, 0x14, 0x43, 0x11, 0xa7, 0x00, 0x95, 0xf6, 0xce, 0xa8, 0xcb, 0xc6, 0xd1,
	0xb1, 0xa5, 0xeb, 0x22, 0xec, 0x18, 0x61, 0xf8, 0x74, 0xb0, 0x8a, 0x0a, 0x12, 0x73, 0x7f, 0xe7,
	0x2d, 0xa8, 0xcb, 0x50, 0x34, 0xdd, 0x7b, 0x4f, 0x3d, 0xfd, 0xd5, 0x13, 0xb2, 0xf3, 0x18, 0x75,
	0x99, 0x3b, 0xf4, 0x9e, 0x3b, 0x46, 0xae, 0xf1, 0x6c, 0xb0, 0xb4, 0x51, 0x42, 0x96, 0xde, 0x0d,
	0x75, 0xd9, 0x24, 0x3a, 0xd9, 0x6b, 0xaf, 0x16, 0xae, 0x46, 0xbb, 0xed, 0xdc, 0x19, 0xa3, 0xf0,
	0xe5, 0x6d, 0xfd, 0xd7, 0x8d, 0xf1, 0x7e, 0x35, 0xc1, 0xc1, 0x25, 0x1e, 0x0e, 0xe6, 0x74, 0x64,
	0x1f, 0xde, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xed, 0x71, 0x7d, 0x45, 0xb3, 0x01, 0x00, 0x00,
}
