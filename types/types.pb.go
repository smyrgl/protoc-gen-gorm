// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UUIDValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDValue) Reset()         { *m = UUIDValue{} }
func (m *UUIDValue) String() string { return proto.CompactTextString(m) }
func (*UUIDValue) ProtoMessage()    {}
func (*UUIDValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0}
}

func (m *UUIDValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDValue.Unmarshal(m, b)
}
func (m *UUIDValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDValue.Marshal(b, m, deterministic)
}
func (m *UUIDValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDValue.Merge(m, src)
}
func (m *UUIDValue) XXX_Size() int {
	return xxx_messageInfo_UUIDValue.Size(m)
}
func (m *UUIDValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDValue.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDValue proto.InternalMessageInfo

func (m *UUIDValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type JSONValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONValue) Reset()         { *m = JSONValue{} }
func (m *JSONValue) String() string { return proto.CompactTextString(m) }
func (*JSONValue) ProtoMessage()    {}
func (*JSONValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{1}
}

func (m *JSONValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONValue.Unmarshal(m, b)
}
func (m *JSONValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONValue.Marshal(b, m, deterministic)
}
func (m *JSONValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONValue.Merge(m, src)
}
func (m *JSONValue) XXX_Size() int {
	return xxx_messageInfo_JSONValue.Size(m)
}
func (m *JSONValue) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONValue.DiscardUnknown(m)
}

var xxx_messageInfo_JSONValue proto.InternalMessageInfo

func (m *JSONValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UUID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{2}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type InetValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InetValue) Reset()         { *m = InetValue{} }
func (m *InetValue) String() string { return proto.CompactTextString(m) }
func (*InetValue) ProtoMessage()    {}
func (*InetValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{3}
}

func (m *InetValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InetValue.Unmarshal(m, b)
}
func (m *InetValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InetValue.Marshal(b, m, deterministic)
}
func (m *InetValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InetValue.Merge(m, src)
}
func (m *InetValue) XXX_Size() int {
	return xxx_messageInfo_InetValue.Size(m)
}
func (m *InetValue) XXX_DiscardUnknown() {
	xxx_messageInfo_InetValue.DiscardUnknown(m)
}

var xxx_messageInfo_InetValue proto.InternalMessageInfo

func (m *InetValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TimeOnly struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeOnly) Reset()         { *m = TimeOnly{} }
func (m *TimeOnly) String() string { return proto.CompactTextString(m) }
func (*TimeOnly) ProtoMessage()    {}
func (*TimeOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{4}
}

func (m *TimeOnly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeOnly.Unmarshal(m, b)
}
func (m *TimeOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeOnly.Marshal(b, m, deterministic)
}
func (m *TimeOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeOnly.Merge(m, src)
}
func (m *TimeOnly) XXX_Size() int {
	return xxx_messageInfo_TimeOnly.Size(m)
}
func (m *TimeOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeOnly.DiscardUnknown(m)
}

var xxx_messageInfo_TimeOnly proto.InternalMessageInfo

func (m *TimeOnly) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*UUIDValue)(nil), "gorm.types.UUIDValue")
	proto.RegisterType((*JSONValue)(nil), "gorm.types.JSONValue")
	proto.RegisterType((*UUID)(nil), "gorm.types.UUID")
	proto.RegisterType((*InetValue)(nil), "gorm.types.InetValue")
	proto.RegisterType((*TimeOnly)(nil), "gorm.types.TimeOnly")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptor_2c0f90c600ad7e2e) }

var fileDescriptor_2c0f90c600ad7e2e = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c, 0xe9, 0xf9, 0x45, 0xb9,
	0x7a, 0x60, 0x11, 0x25, 0x45, 0x2e, 0xce, 0xd0, 0x50, 0x4f, 0x97, 0xb0, 0xc4, 0x9c, 0xd2, 0x54,
	0x21, 0x11, 0x2e, 0xd6, 0x32, 0x10, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x01,
	0x29, 0xf1, 0x0a, 0xf6, 0xf7, 0xc3, 0xa7, 0x44, 0x86, 0x8b, 0x05, 0x64, 0x0a, 0x6e, 0x03, 0x3c,
	0xf3, 0x52, 0x4b, 0xf0, 0x19, 0xa0, 0xc0, 0xc5, 0x11, 0x92, 0x99, 0x9b, 0xea, 0x9f, 0x97, 0x53,
	0x89, 0xaa, 0x82, 0x17, 0xaa, 0xc2, 0x49, 0x3f, 0x4a, 0x37, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0xb7, 0xb2, 0x28, 0x3d, 0x47, 0x1f, 0xec, 0x9f, 0x64, 0xdd,
	0xf4, 0xd4, 0x3c, 0x5d, 0x90, 0x9f, 0x20, 0xbe, 0xb4, 0x06, 0x93, 0x49, 0x6c, 0x60, 0x49, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xcc, 0x56, 0x1e, 0x01, 0x01, 0x00, 0x00,
}
