// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: xds.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnvoyConfigurationEvent_EventKind int32

const (
	EnvoyConfigurationEvent_EVENT_KIND_UNDEFINED EnvoyConfigurationEvent_EventKind = 0
	// envoy_service_discovery_v3.DeltaDiscoveryRequest
	EnvoyConfigurationEvent_EVENT_DISCOVERY_REQUEST EnvoyConfigurationEvent_EventKind = 1
	// envoy_service_discovery_v3.DeltaDiscoveryResponse
	EnvoyConfigurationEvent_EVENT_DISCOVERY_RESPONSE EnvoyConfigurationEvent_EventKind = 2
)

// Enum value maps for EnvoyConfigurationEvent_EventKind.
var (
	EnvoyConfigurationEvent_EventKind_name = map[int32]string{
		0: "EVENT_KIND_UNDEFINED",
		1: "EVENT_DISCOVERY_REQUEST",
		2: "EVENT_DISCOVERY_RESPONSE",
	}
	EnvoyConfigurationEvent_EventKind_value = map[string]int32{
		"EVENT_KIND_UNDEFINED":     0,
		"EVENT_DISCOVERY_REQUEST":  1,
		"EVENT_DISCOVERY_RESPONSE": 2,
	}
)

func (x EnvoyConfigurationEvent_EventKind) Enum() *EnvoyConfigurationEvent_EventKind {
	p := new(EnvoyConfigurationEvent_EventKind)
	*p = x
	return p
}

func (x EnvoyConfigurationEvent_EventKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnvoyConfigurationEvent_EventKind) Descriptor() protoreflect.EnumDescriptor {
	return file_xds_proto_enumTypes[0].Descriptor()
}

func (EnvoyConfigurationEvent_EventKind) Type() protoreflect.EnumType {
	return &file_xds_proto_enumTypes[0]
}

func (x EnvoyConfigurationEvent_EventKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnvoyConfigurationEvent_EventKind.Descriptor instead.
func (EnvoyConfigurationEvent_EventKind) EnumDescriptor() ([]byte, []int) {
	return file_xds_proto_rawDescGZIP(), []int{1, 0}
}

// EnvoyConfigurationEvents is a list of envoy configuration events.
type EnvoyConfigurationEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*EnvoyConfigurationEvent `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *EnvoyConfigurationEvents) Reset() {
	*x = EnvoyConfigurationEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvoyConfigurationEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvoyConfigurationEvents) ProtoMessage() {}

func (x *EnvoyConfigurationEvents) ProtoReflect() protoreflect.Message {
	mi := &file_xds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvoyConfigurationEvents.ProtoReflect.Descriptor instead.
func (*EnvoyConfigurationEvents) Descriptor() ([]byte, []int) {
	return file_xds_proto_rawDescGZIP(), []int{0}
}

func (x *EnvoyConfigurationEvents) GetValues() []*EnvoyConfigurationEvent {
	if x != nil {
		return x.Values
	}
	return nil
}

// EnvoyConfigurationEvent is an envoy configuration event.
type EnvoyConfigurationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32                  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Details []*anypb.Any           `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
	// databroker config version
	ConfigVersion uint64 `protobuf:"varint,5,opt,name=config_version,json=configVersion,proto3" json:"config_version,omitempty"`
	// envoy resource type (i.e. listener, cluster)
	TypeUrl              string                            `protobuf:"bytes,6,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Kind                 EnvoyConfigurationEvent_EventKind `protobuf:"varint,7,opt,name=kind,proto3,enum=pomerium.events.EnvoyConfigurationEvent_EventKind" json:"kind,omitempty"`
	ResourceSubscribed   []string                          `protobuf:"bytes,8,rep,name=resource_subscribed,json=resourceSubscribed,proto3" json:"resource_subscribed,omitempty"`
	ResourceUnsubscribed []string                          `protobuf:"bytes,9,rep,name=resource_unsubscribed,json=resourceUnsubscribed,proto3" json:"resource_unsubscribed,omitempty"`
	// instance this event originated from
	Instance string `protobuf:"bytes,10,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *EnvoyConfigurationEvent) Reset() {
	*x = EnvoyConfigurationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvoyConfigurationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvoyConfigurationEvent) ProtoMessage() {}

func (x *EnvoyConfigurationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_xds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvoyConfigurationEvent.ProtoReflect.Descriptor instead.
func (*EnvoyConfigurationEvent) Descriptor() ([]byte, []int) {
	return file_xds_proto_rawDescGZIP(), []int{1}
}

func (x *EnvoyConfigurationEvent) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EnvoyConfigurationEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EnvoyConfigurationEvent) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EnvoyConfigurationEvent) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *EnvoyConfigurationEvent) GetConfigVersion() uint64 {
	if x != nil {
		return x.ConfigVersion
	}
	return 0
}

func (x *EnvoyConfigurationEvent) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *EnvoyConfigurationEvent) GetKind() EnvoyConfigurationEvent_EventKind {
	if x != nil {
		return x.Kind
	}
	return EnvoyConfigurationEvent_EVENT_KIND_UNDEFINED
}

func (x *EnvoyConfigurationEvent) GetResourceSubscribed() []string {
	if x != nil {
		return x.ResourceSubscribed
	}
	return nil
}

func (x *EnvoyConfigurationEvent) GetResourceUnsubscribed() []string {
	if x != nil {
		return x.ResourceUnsubscribed
	}
	return nil
}

func (x *EnvoyConfigurationEvent) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

var File_xds_proto protoreflect.FileDescriptor

var file_xds_proto_rawDesc = []byte{
	0x0a, 0x09, 0x78, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x6f, 0x6d,
	0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x18, 0x45, 0x6e, 0x76, 0x6f,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x95, 0x04, 0x0a, 0x17, 0x45, 0x6e, 0x76, 0x6f, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x46, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x15, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x60, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6d,
	0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_proto_rawDescOnce sync.Once
	file_xds_proto_rawDescData = file_xds_proto_rawDesc
)

func file_xds_proto_rawDescGZIP() []byte {
	file_xds_proto_rawDescOnce.Do(func() {
		file_xds_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_proto_rawDescData)
	})
	return file_xds_proto_rawDescData
}

var file_xds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_xds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xds_proto_goTypes = []interface{}{
	(EnvoyConfigurationEvent_EventKind)(0), // 0: pomerium.events.EnvoyConfigurationEvent.EventKind
	(*EnvoyConfigurationEvents)(nil),       // 1: pomerium.events.EnvoyConfigurationEvents
	(*EnvoyConfigurationEvent)(nil),        // 2: pomerium.events.EnvoyConfigurationEvent
	(*timestamppb.Timestamp)(nil),          // 3: google.protobuf.Timestamp
	(*anypb.Any)(nil),                      // 4: google.protobuf.Any
}
var file_xds_proto_depIdxs = []int32{
	2, // 0: pomerium.events.EnvoyConfigurationEvents.values:type_name -> pomerium.events.EnvoyConfigurationEvent
	3, // 1: pomerium.events.EnvoyConfigurationEvent.time:type_name -> google.protobuf.Timestamp
	4, // 2: pomerium.events.EnvoyConfigurationEvent.details:type_name -> google.protobuf.Any
	0, // 3: pomerium.events.EnvoyConfigurationEvent.kind:type_name -> pomerium.events.EnvoyConfigurationEvent.EventKind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_xds_proto_init() }
func file_xds_proto_init() {
	if File_xds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvoyConfigurationEvents); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_xds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvoyConfigurationEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_proto_goTypes,
		DependencyIndexes: file_xds_proto_depIdxs,
		EnumInfos:         file_xds_proto_enumTypes,
		MessageInfos:      file_xds_proto_msgTypes,
	}.Build()
	File_xds_proto = out.File
	file_xds_proto_rawDesc = nil
	file_xds_proto_goTypes = nil
	file_xds_proto_depIdxs = nil
}