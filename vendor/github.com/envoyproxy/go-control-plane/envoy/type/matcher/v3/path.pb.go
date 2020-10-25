// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/path.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type PathMatcher struct {
	// Types that are valid to be assigned to Rule:
	//	*PathMatcher_Path
	Rule                 isPathMatcher_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PathMatcher) Reset()         { *m = PathMatcher{} }
func (m *PathMatcher) String() string { return proto.CompactTextString(m) }
func (*PathMatcher) ProtoMessage()    {}
func (*PathMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad26ce2b727dd011, []int{0}
}

func (m *PathMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathMatcher.Unmarshal(m, b)
}
func (m *PathMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathMatcher.Marshal(b, m, deterministic)
}
func (m *PathMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathMatcher.Merge(m, src)
}
func (m *PathMatcher) XXX_Size() int {
	return xxx_messageInfo_PathMatcher.Size(m)
}
func (m *PathMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_PathMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_PathMatcher proto.InternalMessageInfo

type isPathMatcher_Rule interface {
	isPathMatcher_Rule()
}

type PathMatcher_Path struct {
	Path *StringMatcher `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

func (*PathMatcher_Path) isPathMatcher_Rule() {}

func (m *PathMatcher) GetRule() isPathMatcher_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *PathMatcher) GetPath() *StringMatcher {
	if x, ok := m.GetRule().(*PathMatcher_Path); ok {
		return x.Path
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PathMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PathMatcher_Path)(nil),
	}
}

func init() {
	proto.RegisterType((*PathMatcher)(nil), "envoy.type.matcher.v3.PathMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/path.proto", fileDescriptor_ad26ce2b727dd011) }

var fileDescriptor_ad26ce2b727dd011 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x33, 0xd6, 0x2f, 0x48, 0x2c, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x05, 0xab,
	0xd0, 0x03, 0xa9, 0xd0, 0x83, 0xaa, 0xd0, 0x2b, 0x33, 0x96, 0x52, 0xc2, 0xae, 0xb1, 0xb8, 0xa4,
	0x28, 0x33, 0x2f, 0x1d, 0xa2, 0x55, 0x4a, 0xb1, 0x34, 0xa5, 0x20, 0x51, 0x3f, 0x31, 0x2f, 0x2f,
	0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xbf, 0x2c, 0xb5, 0xa8, 0x38, 0x33, 0x3f, 0x0f,
	0xa1, 0x44, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80, 0x48, 0x28,
	0xb5, 0x33, 0x72, 0x71, 0x07, 0x24, 0x96, 0x64, 0xf8, 0x42, 0xcc, 0x16, 0x72, 0xe1, 0x62, 0x01,
	0x39, 0x4a, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x45, 0x0f, 0xab, 0xab, 0xf4, 0x82, 0xc1,
	0xd6, 0x43, 0xf5, 0x38, 0x71, 0xfc, 0x72, 0x62, 0xed, 0x62, 0x64, 0x12, 0x60, 0xf4, 0x60, 0x08,
	0x02, 0xeb, 0xb6, 0x52, 0x9d, 0x75, 0xb4, 0x43, 0x4e, 0x81, 0x4b, 0x0e, 0x8b, 0x6e, 0x24, 0xcb,
	0x9c, 0xb8, 0xb9, 0x58, 0x8a, 0x4a, 0x73, 0x52, 0x85, 0x98, 0x7f, 0x38, 0x31, 0x3a, 0x19, 0x71,
	0x29, 0x67, 0xe6, 0x43, 0xec, 0x2b, 0x28, 0xca, 0xaf, 0xa8, 0xc4, 0x6e, 0xb5, 0x13, 0x27, 0xc8,
	0x80, 0x00, 0x90, 0xdb, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0x9e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0x66, 0x7d, 0x37, 0x5f, 0x01, 0x00, 0x00,
}
