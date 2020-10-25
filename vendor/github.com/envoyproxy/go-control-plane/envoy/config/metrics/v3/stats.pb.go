// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v3/stats.proto

package envoy_config_metrics_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type StatsSink struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*StatsSink_HiddenEnvoyDeprecatedConfig
	//	*StatsSink_TypedConfig
	ConfigType           isStatsSink_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StatsSink) Reset()         { *m = StatsSink{} }
func (m *StatsSink) String() string { return proto.CompactTextString(m) }
func (*StatsSink) ProtoMessage()    {}
func (*StatsSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{0}
}

func (m *StatsSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsSink.Unmarshal(m, b)
}
func (m *StatsSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsSink.Marshal(b, m, deterministic)
}
func (m *StatsSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsSink.Merge(m, src)
}
func (m *StatsSink) XXX_Size() int {
	return xxx_messageInfo_StatsSink.Size(m)
}
func (m *StatsSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsSink proto.InternalMessageInfo

func (m *StatsSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isStatsSink_ConfigType interface {
	isStatsSink_ConfigType()
}

type StatsSink_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type StatsSink_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*StatsSink_HiddenEnvoyDeprecatedConfig) isStatsSink_ConfigType() {}

func (*StatsSink_TypedConfig) isStatsSink_ConfigType() {}

func (m *StatsSink) GetConfigType() isStatsSink_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *StatsSink) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*StatsSink_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *StatsSink) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*StatsSink_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsSink_HiddenEnvoyDeprecatedConfig)(nil),
		(*StatsSink_TypedConfig)(nil),
	}
}

type StatsConfig struct {
	StatsTags            []*TagSpecifier     `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags,proto3" json:"stats_tags,omitempty"`
	UseAllDefaultTags    *wrappers.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags,proto3" json:"use_all_default_tags,omitempty"`
	StatsMatcher         *StatsMatcher       `protobuf:"bytes,3,opt,name=stats_matcher,json=statsMatcher,proto3" json:"stats_matcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StatsConfig) Reset()         { *m = StatsConfig{} }
func (m *StatsConfig) String() string { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()    {}
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{1}
}

func (m *StatsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsConfig.Unmarshal(m, b)
}
func (m *StatsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsConfig.Marshal(b, m, deterministic)
}
func (m *StatsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsConfig.Merge(m, src)
}
func (m *StatsConfig) XXX_Size() int {
	return xxx_messageInfo_StatsConfig.Size(m)
}
func (m *StatsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StatsConfig proto.InternalMessageInfo

func (m *StatsConfig) GetStatsTags() []*TagSpecifier {
	if m != nil {
		return m.StatsTags
	}
	return nil
}

func (m *StatsConfig) GetUseAllDefaultTags() *wrappers.BoolValue {
	if m != nil {
		return m.UseAllDefaultTags
	}
	return nil
}

func (m *StatsConfig) GetStatsMatcher() *StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

type StatsMatcher struct {
	// Types that are valid to be assigned to StatsMatcher:
	//	*StatsMatcher_RejectAll
	//	*StatsMatcher_ExclusionList
	//	*StatsMatcher_InclusionList
	StatsMatcher         isStatsMatcher_StatsMatcher `protobuf_oneof:"stats_matcher"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StatsMatcher) Reset()         { *m = StatsMatcher{} }
func (m *StatsMatcher) String() string { return proto.CompactTextString(m) }
func (*StatsMatcher) ProtoMessage()    {}
func (*StatsMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{2}
}

func (m *StatsMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsMatcher.Unmarshal(m, b)
}
func (m *StatsMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsMatcher.Marshal(b, m, deterministic)
}
func (m *StatsMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsMatcher.Merge(m, src)
}
func (m *StatsMatcher) XXX_Size() int {
	return xxx_messageInfo_StatsMatcher.Size(m)
}
func (m *StatsMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StatsMatcher proto.InternalMessageInfo

type isStatsMatcher_StatsMatcher interface {
	isStatsMatcher_StatsMatcher()
}

type StatsMatcher_RejectAll struct {
	RejectAll bool `protobuf:"varint,1,opt,name=reject_all,json=rejectAll,proto3,oneof"`
}

type StatsMatcher_ExclusionList struct {
	ExclusionList *v3.ListStringMatcher `protobuf:"bytes,2,opt,name=exclusion_list,json=exclusionList,proto3,oneof"`
}

type StatsMatcher_InclusionList struct {
	InclusionList *v3.ListStringMatcher `protobuf:"bytes,3,opt,name=inclusion_list,json=inclusionList,proto3,oneof"`
}

func (*StatsMatcher_RejectAll) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_ExclusionList) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_InclusionList) isStatsMatcher_StatsMatcher() {}

func (m *StatsMatcher) GetStatsMatcher() isStatsMatcher_StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

func (m *StatsMatcher) GetRejectAll() bool {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_RejectAll); ok {
		return x.RejectAll
	}
	return false
}

func (m *StatsMatcher) GetExclusionList() *v3.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_ExclusionList); ok {
		return x.ExclusionList
	}
	return nil
}

func (m *StatsMatcher) GetInclusionList() *v3.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_InclusionList); ok {
		return x.InclusionList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsMatcher_RejectAll)(nil),
		(*StatsMatcher_ExclusionList)(nil),
		(*StatsMatcher_InclusionList)(nil),
	}
}

type TagSpecifier struct {
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	// Types that are valid to be assigned to TagValue:
	//	*TagSpecifier_Regex
	//	*TagSpecifier_FixedValue
	TagValue             isTagSpecifier_TagValue `protobuf_oneof:"tag_value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TagSpecifier) Reset()         { *m = TagSpecifier{} }
func (m *TagSpecifier) String() string { return proto.CompactTextString(m) }
func (*TagSpecifier) ProtoMessage()    {}
func (*TagSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{3}
}

func (m *TagSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagSpecifier.Unmarshal(m, b)
}
func (m *TagSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagSpecifier.Marshal(b, m, deterministic)
}
func (m *TagSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagSpecifier.Merge(m, src)
}
func (m *TagSpecifier) XXX_Size() int {
	return xxx_messageInfo_TagSpecifier.Size(m)
}
func (m *TagSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TagSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_TagSpecifier proto.InternalMessageInfo

func (m *TagSpecifier) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

type isTagSpecifier_TagValue interface {
	isTagSpecifier_TagValue()
}

type TagSpecifier_Regex struct {
	Regex string `protobuf:"bytes,2,opt,name=regex,proto3,oneof"`
}

type TagSpecifier_FixedValue struct {
	FixedValue string `protobuf:"bytes,3,opt,name=fixed_value,json=fixedValue,proto3,oneof"`
}

func (*TagSpecifier_Regex) isTagSpecifier_TagValue() {}

func (*TagSpecifier_FixedValue) isTagSpecifier_TagValue() {}

func (m *TagSpecifier) GetTagValue() isTagSpecifier_TagValue {
	if m != nil {
		return m.TagValue
	}
	return nil
}

func (m *TagSpecifier) GetRegex() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *TagSpecifier) GetFixedValue() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_FixedValue); ok {
		return x.FixedValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagSpecifier_Regex)(nil),
		(*TagSpecifier_FixedValue)(nil),
	}
}

type StatsdSink struct {
	// Types that are valid to be assigned to StatsdSpecifier:
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier      isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
	Prefix               string                       `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StatsdSink) Reset()         { *m = StatsdSink{} }
func (m *StatsdSink) String() string { return proto.CompactTextString(m) }
func (*StatsdSink) ProtoMessage()    {}
func (*StatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{4}
}

func (m *StatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsdSink.Unmarshal(m, b)
}
func (m *StatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsdSink.Marshal(b, m, deterministic)
}
func (m *StatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsdSink.Merge(m, src)
}
func (m *StatsdSink) XXX_Size() int {
	return xxx_messageInfo_StatsdSink.Size(m)
}
func (m *StatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsdSink proto.InternalMessageInfo

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
}

type StatsdSink_Address struct {
	Address *v31.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type StatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier() {}

func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (m *StatsdSink) GetAddress() *v31.Address {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *StatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

func (m *StatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
}

type DogStatsdSink struct {
	// Types that are valid to be assigned to DogStatsdSpecifier:
	//	*DogStatsdSink_Address
	DogStatsdSpecifier   isDogStatsdSink_DogStatsdSpecifier `protobuf_oneof:"dog_statsd_specifier"`
	Prefix               string                             `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *DogStatsdSink) Reset()         { *m = DogStatsdSink{} }
func (m *DogStatsdSink) String() string { return proto.CompactTextString(m) }
func (*DogStatsdSink) ProtoMessage()    {}
func (*DogStatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{5}
}

func (m *DogStatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogStatsdSink.Unmarshal(m, b)
}
func (m *DogStatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogStatsdSink.Marshal(b, m, deterministic)
}
func (m *DogStatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogStatsdSink.Merge(m, src)
}
func (m *DogStatsdSink) XXX_Size() int {
	return xxx_messageInfo_DogStatsdSink.Size(m)
}
func (m *DogStatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_DogStatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_DogStatsdSink proto.InternalMessageInfo

type isDogStatsdSink_DogStatsdSpecifier interface {
	isDogStatsdSink_DogStatsdSpecifier()
}

type DogStatsdSink_Address struct {
	Address *v31.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

func (*DogStatsdSink_Address) isDogStatsdSink_DogStatsdSpecifier() {}

func (m *DogStatsdSink) GetDogStatsdSpecifier() isDogStatsdSink_DogStatsdSpecifier {
	if m != nil {
		return m.DogStatsdSpecifier
	}
	return nil
}

func (m *DogStatsdSink) GetAddress() *v31.Address {
	if x, ok := m.GetDogStatsdSpecifier().(*DogStatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *DogStatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DogStatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DogStatsdSink_Address)(nil),
	}
}

type HystrixSink struct {
	NumBuckets           int64    `protobuf:"varint,1,opt,name=num_buckets,json=numBuckets,proto3" json:"num_buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HystrixSink) Reset()         { *m = HystrixSink{} }
func (m *HystrixSink) String() string { return proto.CompactTextString(m) }
func (*HystrixSink) ProtoMessage()    {}
func (*HystrixSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_38edc339428e9971, []int{6}
}

func (m *HystrixSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HystrixSink.Unmarshal(m, b)
}
func (m *HystrixSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HystrixSink.Marshal(b, m, deterministic)
}
func (m *HystrixSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HystrixSink.Merge(m, src)
}
func (m *HystrixSink) XXX_Size() int {
	return xxx_messageInfo_HystrixSink.Size(m)
}
func (m *HystrixSink) XXX_DiscardUnknown() {
	xxx_messageInfo_HystrixSink.DiscardUnknown(m)
}

var xxx_messageInfo_HystrixSink proto.InternalMessageInfo

func (m *HystrixSink) GetNumBuckets() int64 {
	if m != nil {
		return m.NumBuckets
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsSink)(nil), "envoy.config.metrics.v3.StatsSink")
	proto.RegisterType((*StatsConfig)(nil), "envoy.config.metrics.v3.StatsConfig")
	proto.RegisterType((*StatsMatcher)(nil), "envoy.config.metrics.v3.StatsMatcher")
	proto.RegisterType((*TagSpecifier)(nil), "envoy.config.metrics.v3.TagSpecifier")
	proto.RegisterType((*StatsdSink)(nil), "envoy.config.metrics.v3.StatsdSink")
	proto.RegisterType((*DogStatsdSink)(nil), "envoy.config.metrics.v3.DogStatsdSink")
	proto.RegisterType((*HystrixSink)(nil), "envoy.config.metrics.v3.HystrixSink")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v3/stats.proto", fileDescriptor_38edc339428e9971)
}

var fileDescriptor_38edc339428e9971 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xf2, 0x46,
	0x14, 0xc5, 0xf0, 0xfd, 0xc0, 0x35, 0x7c, 0x4a, 0x2d, 0x54, 0xf8, 0x92, 0x36, 0x21, 0x24, 0xb4,
	0x34, 0xad, 0x6c, 0x89, 0xa8, 0x8b, 0xb0, 0xc3, 0xa1, 0x12, 0x4a, 0x7f, 0x94, 0x9a, 0xa8, 0x8b,
	0x6e, 0xac, 0xc1, 0x1e, 0x9c, 0x69, 0x8c, 0xc7, 0x1a, 0x8f, 0x29, 0xec, 0xba, 0xec, 0x23, 0x54,
	0x7d, 0x82, 0xee, 0xfb, 0x0c, 0x5d, 0xf7, 0x71, 0x5a, 0x75, 0x55, 0xcd, 0x0f, 0x14, 0x92, 0x90,
	0xa8, 0xd2, 0xb7, 0xc3, 0x3e, 0xf7, 0x9e, 0x39, 0xe7, 0xf8, 0xde, 0x01, 0x4e, 0x70, 0x32, 0xa7,
	0x4b, 0x27, 0xa0, 0xc9, 0x94, 0x44, 0xce, 0x0c, 0x73, 0x46, 0x82, 0xcc, 0x99, 0x9f, 0x3b, 0x19,
	0x47, 0x3c, 0xb3, 0x53, 0x46, 0x39, 0xb5, 0x1a, 0xb2, 0xc8, 0x56, 0x45, 0xb6, 0x2e, 0xb2, 0xe7,
	0xe7, 0xfb, 0xed, 0xad, 0xee, 0x80, 0x32, 0x2c, 0x5a, 0x51, 0x18, 0x32, 0x9c, 0xe9, 0xe6, 0x55,
	0x0d, 0x5f, 0xa6, 0xd8, 0x99, 0x21, 0x1e, 0xdc, 0x62, 0xa6, 0xf8, 0x19, 0x49, 0x22, 0x5d, 0xf3,
	0x36, 0xa2, 0x34, 0x8a, 0xb1, 0x23, 0x9f, 0x26, 0xf9, 0xd4, 0x41, 0xc9, 0x52, 0x43, 0x1f, 0xdc,
	0x87, 0x32, 0xce, 0xf2, 0x80, 0x6b, 0xf4, 0xf0, 0x3e, 0xfa, 0x23, 0x43, 0x69, 0x8a, 0xd9, 0xea,
	0xf0, 0xe3, 0x3c, 0x4c, 0x91, 0x83, 0x92, 0x84, 0x72, 0xc4, 0x09, 0x4d, 0x32, 0x67, 0x8e, 0x59,
	0x46, 0x68, 0xf2, 0xdf, 0xd9, 0x8d, 0x39, 0x8a, 0x49, 0x88, 0x38, 0x76, 0x56, 0x3f, 0x14, 0xd0,
	0xfe, 0xcb, 0x80, 0xca, 0x58, 0xa4, 0x30, 0x26, 0xc9, 0x9d, 0x65, 0xc1, 0x8b, 0x04, 0xcd, 0x70,
	0xd3, 0x68, 0x19, 0xdd, 0x8a, 0x27, 0x7f, 0x5b, 0x13, 0x38, 0xbc, 0x25, 0x61, 0x88, 0x13, 0x5f,
	0x7a, 0xf4, 0x43, 0x9c, 0x32, 0x1c, 0x20, 0x8e, 0x43, 0x5f, 0x45, 0xd2, 0x2c, 0xb6, 0x8c, 0xae,
	0xd9, 0x6b, 0xd8, 0x4a, 0xa6, 0xbd, 0x92, 0x69, 0x8f, 0xa5, 0x09, 0xb7, 0xd8, 0x34, 0x46, 0x05,
	0xef, 0x40, 0x91, 0x7c, 0x21, 0x38, 0x86, 0x6b, 0x8a, 0x4b, 0xc9, 0x60, 0x5d, 0x40, 0x55, 0x44,
	0xb7, 0x66, 0x2c, 0x49, 0xc6, 0xfa, 0x03, 0xc6, 0x41, 0xb2, 0x1c, 0x15, 0x3c, 0x53, 0xd6, 0xaa,
	0xd6, 0x7e, 0xf7, 0xd7, 0x3f, 0x7e, 0x3e, 0x3c, 0x81, 0xe3, 0xc7, 0xbf, 0x5e, 0xcf, 0x5e, 0x9b,
	0x73, 0x6b, 0x60, 0x2a, 0xd8, 0x17, 0xfd, 0xed, 0x5f, 0x8a, 0x60, 0x4a, 0x50, 0x6b, 0x18, 0x02,
	0xc8, 0x71, 0xf0, 0x39, 0x8a, 0xb2, 0xa6, 0xd1, 0x2a, 0x75, 0xcd, 0x5e, 0xc7, 0xde, 0x31, 0x14,
	0xf6, 0x0d, 0x8a, 0xc6, 0x29, 0x0e, 0xc8, 0x94, 0x60, 0xe6, 0x55, 0x64, 0xe3, 0x0d, 0x8a, 0x32,
	0xeb, 0x4b, 0xa8, 0xe7, 0x19, 0xf6, 0x51, 0x1c, 0xfb, 0x21, 0x9e, 0xa2, 0x3c, 0xe6, 0x8a, 0x4f,
	0x65, 0xb4, 0xff, 0xc0, 0x91, 0x4b, 0x69, 0xfc, 0x1d, 0x8a, 0x73, 0xec, 0xbd, 0x97, 0x67, 0x78,
	0x10, 0xc7, 0x43, 0xd5, 0x25, 0xc9, 0xae, 0xa0, 0xa6, 0x24, 0xe9, 0x91, 0xd2, 0xb9, 0xec, 0x56,
	0x25, 0xfd, 0x7c, 0xad, 0x8a, 0xbd, 0x6a, 0xb6, 0xf1, 0xd4, 0x3f, 0x13, 0x39, 0x75, 0xf4, 0x2a,
	0xec, 0xca, 0x49, 0x45, 0x21, 0xa2, 0xa9, 0x6e, 0x52, 0x59, 0x47, 0x00, 0x0c, 0xff, 0x80, 0x03,
	0x2e, 0x8c, 0xc9, 0xe9, 0x28, 0x8f, 0x0a, 0x5e, 0x45, 0xbd, 0x1b, 0xc4, 0xb1, 0xf5, 0x2d, 0xbc,
	0xc1, 0x8b, 0x20, 0xce, 0xc5, 0xd4, 0xf9, 0x31, 0xc9, 0xb8, 0x36, 0xdc, 0xd5, 0x52, 0x45, 0xe2,
	0xb6, 0x76, 0x21, 0x84, 0x7e, 0x45, 0x32, 0x3e, 0x96, 0xcb, 0xa1, 0x8f, 0x18, 0x15, 0xbc, 0xda,
	0x9a, 0x41, 0xa0, 0x82, 0x92, 0x24, 0x5b, 0x94, 0xa5, 0xff, 0x4f, 0xb9, 0x66, 0x10, 0x68, 0xff,
	0x53, 0x91, 0xc1, 0x47, 0x70, 0xfa, 0x64, 0x06, 0xba, 0xdb, 0xad, 0xdf, 0x0b, 0xdf, 0x2a, 0xfd,
	0xed, 0x1a, 0xed, 0xdf, 0x0c, 0xa8, 0x6e, 0x7e, 0x7b, 0xeb, 0x2d, 0x94, 0x39, 0x8a, 0xfc, 0x8d,
	0xb5, 0x79, 0xcd, 0x51, 0xf4, 0x8d, 0xd8, 0x9c, 0x16, 0xbc, 0x64, 0x38, 0xc2, 0x0b, 0x99, 0x45,
	0xc5, 0x2d, 0xff, 0xe3, 0xbe, 0x64, 0xa5, 0xee, 0x4f, 0x22, 0x3a, 0x05, 0x58, 0xc7, 0x60, 0x4e,
	0xc9, 0x02, 0x87, 0xfe, 0x5c, 0x8c, 0x80, 0x34, 0x58, 0x19, 0x15, 0x3c, 0x90, 0x2f, 0xe5, 0x58,
	0x3c, 0xab, 0x79, 0x53, 0x8c, 0x6b, 0x42, 0x45, 0x88, 0x91, 0x6c, 0xed, 0x3f, 0x0d, 0x00, 0xe9,
	0x28, 0x94, 0xbb, 0x7d, 0x01, 0xaf, 0xf5, 0x9d, 0x25, 0x75, 0x9a, 0xbd, 0x0f, 0xb7, 0xc7, 0x48,
	0x5c, 0x6c, 0x22, 0xc7, 0x81, 0x2a, 0x1a, 0x15, 0xbc, 0x55, 0xbd, 0x75, 0x06, 0x7b, 0x3c, 0x48,
	0x7d, 0x11, 0x25, 0xc7, 0x4c, 0x79, 0x2d, 0x6a, 0xad, 0x6f, 0x78, 0x90, 0x5e, 0x2a, 0x40, 0x9a,
	0x7e, 0x1f, 0x5e, 0xa5, 0x0c, 0x4f, 0xc9, 0x42, 0xb9, 0xf1, 0xf4, 0x53, 0xff, 0x13, 0xe1, 0xe3,
	0x14, 0xda, 0x4f, 0x66, 0x2f, 0x95, 0xba, 0x0d, 0xd8, 0x93, 0xc9, 0x87, 0x7e, 0xb6, 0x8e, 0x59,
	0x86, 0xff, 0xbb, 0x01, 0xb5, 0x21, 0x8d, 0xde, 0x8d, 0xa9, 0x5d, 0x42, 0x3f, 0x13, 0x42, 0x3f,
	0x86, 0xce, 0x2e, 0xa1, 0x5b, 0x02, 0xdc, 0x03, 0xa8, 0x87, 0x34, 0xf2, 0x1f, 0xd5, 0x7b, 0xf5,
	0xa2, 0x5c, 0xdc, 0x2b, 0xb5, 0xbf, 0x07, 0x73, 0xb4, 0x14, 0xff, 0x04, 0x0b, 0x29, 0xf9, 0x08,
	0xcc, 0x24, 0x9f, 0xf9, 0x93, 0x3c, 0xb8, 0xc3, 0x5c, 0xc9, 0x2e, 0x79, 0x90, 0xe4, 0x33, 0x57,
	0xbd, 0x79, 0x76, 0x53, 0x37, 0xc8, 0xdc, 0xcf, 0xa1, 0x43, 0xa8, 0xb2, 0x9c, 0x32, 0xba, 0x58,
	0xee, 0xba, 0x19, 0x5c, 0x35, 0x09, 0xd7, 0xe2, 0xda, 0xb9, 0x36, 0x26, 0xaf, 0xe4, 0xfd, 0x73,
	0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x84, 0xc2, 0x83, 0x19, 0x07, 0x00, 0x00,
}
