// Code generated by protoc-gen-gogo.
// source: test_schema_xxx.proto
// DO NOT EDIT!

package protein

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OtherTestSchemaXXX struct {
	Ts *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=ts" json:"ts,omitempty"`
}

func (m *OtherTestSchemaXXX) Reset()                    { *m = OtherTestSchemaXXX{} }
func (m *OtherTestSchemaXXX) String() string            { return proto.CompactTextString(m) }
func (*OtherTestSchemaXXX) ProtoMessage()               {}
func (*OtherTestSchemaXXX) Descriptor() ([]byte, []int) { return fileDescriptorTestSchemaXxx, []int{0} }

func (m *OtherTestSchemaXXX) GetTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

type TestSchemaXXX struct {
	Uid    string                                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FqName []string                               `protobuf:"bytes,2,rep,name=fq_name,json=fqName" json:"fq_name,omitempty"`
	Deps   map[string]*TestSchemaXXX_NestedSchema `protobuf:"bytes,4,rep,name=deps" json:"deps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Ids    map[int32]string                       `protobuf:"bytes,12,rep,name=ids" json:"ids,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to Descr:
	//	*TestSchemaXXX_Str
	//	*TestSchemaXXX_Buf
	Descr isTestSchemaXXX_Descr         `protobuf_oneof:"descr"`
	Ts    *google_protobuf1.Timestamp   `protobuf:"bytes,7,opt,name=ts" json:"ts,omitempty"`
	Ots   *OtherTestSchemaXXX           `protobuf:"bytes,9,opt,name=ots" json:"ots,omitempty"`
	Nss   []*TestSchemaXXX_NestedSchema `protobuf:"bytes,8,rep,name=nss" json:"nss,omitempty"`
}

func (m *TestSchemaXXX) Reset()                    { *m = TestSchemaXXX{} }
func (m *TestSchemaXXX) String() string            { return proto.CompactTextString(m) }
func (*TestSchemaXXX) ProtoMessage()               {}
func (*TestSchemaXXX) Descriptor() ([]byte, []int) { return fileDescriptorTestSchemaXxx, []int{1} }

type isTestSchemaXXX_Descr interface {
	isTestSchemaXXX_Descr()
}

type TestSchemaXXX_Str struct {
	Str string `protobuf:"bytes,30,opt,name=str,proto3,oneof"`
}
type TestSchemaXXX_Buf struct {
	Buf []byte `protobuf:"bytes,31,opt,name=buf,proto3,oneof"`
}

func (*TestSchemaXXX_Str) isTestSchemaXXX_Descr() {}
func (*TestSchemaXXX_Buf) isTestSchemaXXX_Descr() {}

func (m *TestSchemaXXX) GetDescr() isTestSchemaXXX_Descr {
	if m != nil {
		return m.Descr
	}
	return nil
}

func (m *TestSchemaXXX) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TestSchemaXXX) GetFqName() []string {
	if m != nil {
		return m.FqName
	}
	return nil
}

func (m *TestSchemaXXX) GetDeps() map[string]*TestSchemaXXX_NestedSchema {
	if m != nil {
		return m.Deps
	}
	return nil
}

func (m *TestSchemaXXX) GetIds() map[int32]string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *TestSchemaXXX) GetStr() string {
	if x, ok := m.GetDescr().(*TestSchemaXXX_Str); ok {
		return x.Str
	}
	return ""
}

func (m *TestSchemaXXX) GetBuf() []byte {
	if x, ok := m.GetDescr().(*TestSchemaXXX_Buf); ok {
		return x.Buf
	}
	return nil
}

func (m *TestSchemaXXX) GetTs() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *TestSchemaXXX) GetOts() *OtherTestSchemaXXX {
	if m != nil {
		return m.Ots
	}
	return nil
}

func (m *TestSchemaXXX) GetNss() []*TestSchemaXXX_NestedSchema {
	if m != nil {
		return m.Nss
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestSchemaXXX) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TestSchemaXXX_OneofMarshaler, _TestSchemaXXX_OneofUnmarshaler, _TestSchemaXXX_OneofSizer, []interface{}{
		(*TestSchemaXXX_Str)(nil),
		(*TestSchemaXXX_Buf)(nil),
	}
}

func _TestSchemaXXX_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestSchemaXXX)
	// descr
	switch x := m.Descr.(type) {
	case *TestSchemaXXX_Str:
		_ = b.EncodeVarint(30<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Str)
	case *TestSchemaXXX_Buf:
		_ = b.EncodeVarint(31<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Buf)
	case nil:
	default:
		return fmt.Errorf("TestSchemaXXX.Descr has unexpected type %T", x)
	}
	return nil
}

func _TestSchemaXXX_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestSchemaXXX)
	switch tag {
	case 30: // descr.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Descr = &TestSchemaXXX_Str{x}
		return true, err
	case 31: // descr.buf
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Descr = &TestSchemaXXX_Buf{x}
		return true, err
	default:
		return false, nil
	}
}

func _TestSchemaXXX_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TestSchemaXXX)
	// descr
	switch x := m.Descr.(type) {
	case *TestSchemaXXX_Str:
		n += proto.SizeVarint(30<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *TestSchemaXXX_Buf:
		n += proto.SizeVarint(31<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Buf)))
		n += len(x.Buf)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TestSchemaXXX_NestedSchema struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TestSchemaXXX_NestedSchema) Reset()         { *m = TestSchemaXXX_NestedSchema{} }
func (m *TestSchemaXXX_NestedSchema) String() string { return proto.CompactTextString(m) }
func (*TestSchemaXXX_NestedSchema) ProtoMessage()    {}
func (*TestSchemaXXX_NestedSchema) Descriptor() ([]byte, []int) {
	return fileDescriptorTestSchemaXxx, []int{1, 0}
}

func (m *TestSchemaXXX_NestedSchema) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TestSchemaXXX_NestedSchema) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*OtherTestSchemaXXX)(nil), "protein.OtherTestSchemaXXX")
	proto.RegisterType((*TestSchemaXXX)(nil), "protein.TestSchemaXXX")
	proto.RegisterType((*TestSchemaXXX_NestedSchema)(nil), "protein.TestSchemaXXX.NestedSchema")
}
func (this *OtherTestSchemaXXX) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protein.OtherTestSchemaXXX{")
	if this.Ts != nil {
		s = append(s, "Ts: "+fmt.Sprintf("%#v", this.Ts)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestSchemaXXX) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&protein.TestSchemaXXX{")
	s = append(s, "Uid: "+fmt.Sprintf("%#v", this.Uid)+",\n")
	s = append(s, "FqName: "+fmt.Sprintf("%#v", this.FqName)+",\n")
	keysForDeps := make([]string, 0, len(this.Deps))
	for k, _ := range this.Deps {
		keysForDeps = append(keysForDeps, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDeps)
	mapStringForDeps := "map[string]*TestSchemaXXX_NestedSchema{"
	for _, k := range keysForDeps {
		mapStringForDeps += fmt.Sprintf("%#v: %#v,", k, this.Deps[k])
	}
	mapStringForDeps += "}"
	if this.Deps != nil {
		s = append(s, "Deps: "+mapStringForDeps+",\n")
	}
	keysForIds := make([]int32, 0, len(this.Ids))
	for k, _ := range this.Ids {
		keysForIds = append(keysForIds, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForIds)
	mapStringForIds := "map[int32]string{"
	for _, k := range keysForIds {
		mapStringForIds += fmt.Sprintf("%#v: %#v,", k, this.Ids[k])
	}
	mapStringForIds += "}"
	if this.Ids != nil {
		s = append(s, "Ids: "+mapStringForIds+",\n")
	}
	if this.Descr != nil {
		s = append(s, "Descr: "+fmt.Sprintf("%#v", this.Descr)+",\n")
	}
	if this.Ts != nil {
		s = append(s, "Ts: "+fmt.Sprintf("%#v", this.Ts)+",\n")
	}
	if this.Ots != nil {
		s = append(s, "Ots: "+fmt.Sprintf("%#v", this.Ots)+",\n")
	}
	if this.Nss != nil {
		s = append(s, "Nss: "+fmt.Sprintf("%#v", this.Nss)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestSchemaXXX_Str) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&protein.TestSchemaXXX_Str{` +
		`Str:` + fmt.Sprintf("%#v", this.Str) + `}`}, ", ")
	return s
}
func (this *TestSchemaXXX_Buf) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&protein.TestSchemaXXX_Buf{` +
		`Buf:` + fmt.Sprintf("%#v", this.Buf) + `}`}, ", ")
	return s
}
func (this *TestSchemaXXX_NestedSchema) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protein.TestSchemaXXX_NestedSchema{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTestSchemaXxx(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func init() { proto.RegisterFile("test_schema_xxx.proto", fileDescriptorTestSchemaXxx) }

var fileDescriptorTestSchemaXxx = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x5d, 0xc7, 0x49, 0xd3, 0x4c, 0x83, 0x84, 0x56, 0x20, 0x56, 0x46, 0x6c, 0xac, 0x72, 0x89,
	0x90, 0x70, 0xa1, 0x40, 0x05, 0x9c, 0x50, 0x05, 0x12, 0x5c, 0x8a, 0x64, 0x7a, 0xc8, 0x01, 0x29,
	0xb2, 0xeb, 0xb1, 0x6b, 0xb5, 0xfe, 0xa8, 0x77, 0x8d, 0x92, 0x5b, 0x8f, 0x1c, 0xf9, 0x09, 0xfc,
	0x04, 0x0e, 0xfc, 0x08, 0x8e, 0x1c, 0x39, 0x82, 0xfd, 0x07, 0x38, 0x72, 0x44, 0xbb, 0x8e, 0xa3,
	0x54, 0xd0, 0x2a, 0x27, 0xef, 0xbc, 0x79, 0x6f, 0xe6, 0xcd, 0x33, 0xdc, 0x94, 0x28, 0xe4, 0x54,
	0x1c, 0x1d, 0x63, 0xe2, 0x4d, 0x67, 0xb3, 0x99, 0x93, 0x17, 0x99, 0xcc, 0x68, 0x5f, 0x7d, 0x30,
	0x4e, 0xad, 0x51, 0x94, 0x65, 0xd1, 0x29, 0xee, 0x68, 0xd8, 0x2f, 0xc3, 0x1d, 0x19, 0x27, 0x28,
	0xa4, 0x97, 0xe4, 0x0d, 0xd3, 0xba, 0xb3, 0xec, 0x44, 0x59, 0x94, 0xe9, 0x42, 0xbf, 0x9a, 0xf6,
	0xf6, 0x0b, 0xa0, 0x6f, 0xe5, 0x31, 0x16, 0x87, 0x28, 0xe4, 0x3b, 0xbd, 0x65, 0x32, 0x99, 0xd0,
	0x7b, 0xd0, 0x91, 0x82, 0x19, 0xb6, 0x31, 0xde, 0xda, 0xb5, 0x9c, 0x66, 0x85, 0xd3, 0x0e, 0x72,
	0x0e, 0xdb, 0x15, 0x6e, 0x47, 0x8a, 0xed, 0xaf, 0x5d, 0xb8, 0x76, 0x51, 0x7d, 0x1d, 0xcc, 0x32,
	0x0e, 0xb4, 0x7c, 0xe0, 0xaa, 0x27, 0xbd, 0x05, 0xfd, 0xf0, 0x6c, 0x9a, 0x7a, 0x09, 0xb2, 0x8e,
	0x6d, 0x8e, 0x07, 0xee, 0x46, 0x78, 0x76, 0xe0, 0x25, 0x48, 0x1f, 0x43, 0x37, 0xc0, 0x5c, 0xb0,
	0xae, 0x6d, 0x8e, 0xb7, 0x76, 0x6d, 0x67, 0x71, 0x96, 0x73, 0x61, 0xa0, 0xf3, 0x12, 0x73, 0xf1,
	0x2a, 0x95, 0xc5, 0xdc, 0xd5, 0x6c, 0xfa, 0x10, 0xcc, 0x38, 0x10, 0x6c, 0xa8, 0x45, 0xa3, 0x4b,
	0x44, 0x6f, 0x82, 0x85, 0x46, 0x71, 0x29, 0x05, 0x53, 0xc8, 0x82, 0x71, 0xe5, 0xe9, 0x35, 0x71,
	0x55, 0xa1, 0x30, 0xbf, 0x0c, 0xd9, 0xc8, 0x36, 0xc6, 0x43, 0x85, 0xf9, 0x65, 0xb8, 0xb8, 0xbc,
	0xbf, 0xce, 0xe5, 0xf4, 0x3e, 0x98, 0x99, 0x14, 0x6c, 0xa0, 0xc9, 0xb7, 0x97, 0x36, 0xfe, 0xcd,
	0xd3, 0x55, 0x3c, 0xfa, 0x04, 0xcc, 0x54, 0x08, 0xb6, 0xa9, 0x5d, 0xdf, 0xbd, 0xc4, 0xf5, 0x01,
	0x0a, 0x89, 0x41, 0x53, 0xbb, 0x8a, 0x6f, 0xed, 0xc1, 0x70, 0x15, 0x54, 0xe9, 0x9e, 0xe0, 0xbc,
	0x4d, 0xf7, 0x04, 0xe7, 0xf4, 0x06, 0xf4, 0x3e, 0x78, 0xa7, 0xa5, 0xca, 0x56, 0x61, 0x4d, 0x61,
	0xbd, 0x87, 0xc1, 0x32, 0xb7, 0xff, 0x88, 0x9e, 0xad, 0x8a, 0xd6, 0xf4, 0xd3, 0x28, 0x9e, 0x77,
	0x9e, 0x1a, 0xd6, 0x1e, 0x6c, 0xb6, 0x01, 0xaf, 0x0e, 0xef, 0x5d, 0xe1, 0x48, 0xe9, 0xf6, 0xfb,
	0xd0, 0x0b, 0x50, 0x1c, 0x15, 0xfb, 0x0f, 0x7e, 0xff, 0xe2, 0xc6, 0x79, 0xc5, 0xc9, 0xc7, 0x8a,
	0x93, 0x6f, 0x15, 0x27, 0xdf, 0x2b, 0x4e, 0x7e, 0x54, 0x9c, 0xfc, 0xac, 0x38, 0xf9, 0x53, 0x71,
	0x72, 0x5e, 0x73, 0xf2, 0xa9, 0xe6, 0xe4, 0x73, 0xcd, 0xc9, 0x97, 0x9a, 0x1b, 0xfe, 0x86, 0xfe,
	0x0d, 0x8f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x8a, 0x7d, 0x64, 0x13, 0x03, 0x00, 0x00,
}
