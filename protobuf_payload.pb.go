// Code generated by protoc-gen-gogo.
// source: protobuf_payload.proto
// DO NOT EDIT!

/*
Package protein is a generated protocol buffer package.

It is generated from these files:
	protobuf_payload.proto
	protobuf_schema.proto

It has these top-level messages:
	ProtobufPayload
	ProtobufSchema
*/
package protein

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ProtobufPayload is a versioned protobuf payload.
type ProtobufPayload struct {
	// UID is the unique, deterministic & versioned identifier for this schema.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Payload is the actual, marshaled protobuf payload.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *ProtobufPayload) Reset()                    { *m = ProtobufPayload{} }
func (m *ProtobufPayload) String() string            { return proto.CompactTextString(m) }
func (*ProtobufPayload) ProtoMessage()               {}
func (*ProtobufPayload) Descriptor() ([]byte, []int) { return fileDescriptorProtobufPayload, []int{0} }

func (m *ProtobufPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *ProtobufPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtobufPayload)(nil), "protein.ProtobufPayload")
}
func (this *ProtobufPayload) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protein.ProtobufPayload{")
	s = append(s, "UID: "+fmt.Sprintf("%#v", this.UID)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtobufPayload(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func init() { proto.RegisterFile("protobuf_payload.proto", fileDescriptorProtobufPayload) }

var fileDescriptorProtobufPayload = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0x8b, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0xd1, 0x03, 0x0b, 0x08,
	0xb1, 0x83, 0xa8, 0xd4, 0xcc, 0x3c, 0x29, 0x59, 0x98, 0x02, 0xfd, 0xf4, 0xfc, 0xf4, 0x7c, 0x30,
	0x07, 0xcc, 0x82, 0xa8, 0x53, 0x0a, 0xe6, 0xe2, 0x0f, 0x80, 0x2a, 0x08, 0x80, 0x18, 0x20, 0x24,
	0xc9, 0xc5, 0x5c, 0x9a, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xfe, 0xe8, 0x9e,
	0x3c, 0x73, 0xa8, 0xa7, 0x4b, 0x10, 0x48, 0x4c, 0x48, 0x95, 0x8b, 0x1d, 0x6a, 0x8d, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x8f, 0x13, 0xf7, 0xa3, 0x7b, 0xf2, 0xec, 0x50, 0x8d, 0x41, 0x30, 0x39, 0x27,
	0x83, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0x74, 0x3c, 0x92, 0x63, 0x38, 0xf1, 0x48,
	0x8e, 0xe1, 0xc2, 0x23, 0x39, 0x86, 0x1b, 0x8f, 0xe4, 0x18, 0x1e, 0x3c, 0x92, 0x63, 0xf8, 0xf1,
	0x48, 0x8e, 0xa1, 0xe1, 0xb1, 0x1c, 0xc3, 0x84, 0xc7, 0x72, 0x0c, 0x0b, 0x1e, 0xcb, 0x31, 0x6c,
	0x78, 0x2c, 0xc7, 0x98, 0xc4, 0x06, 0x76, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x20, 0xe2,
	0x7a, 0xfb, 0xcf, 0x00, 0x00, 0x00,
}
