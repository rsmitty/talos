// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: resource/definitions/siderolink/siderolink.proto

package siderolink

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	common "github.com/siderolabs/talos/pkg/machinery/api/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConfigSpec describes Siderolink configuration.
type ConfigSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiEndpoint   string                 `protobuf:"bytes,1,opt,name=api_endpoint,json=apiEndpoint,proto3" json:"api_endpoint,omitempty"`
	Host          string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	JoinToken     string                 `protobuf:"bytes,3,opt,name=join_token,json=joinToken,proto3" json:"join_token,omitempty"`
	Insecure      bool                   `protobuf:"varint,4,opt,name=insecure,proto3" json:"insecure,omitempty"`
	Tunnel        bool                   `protobuf:"varint,5,opt,name=tunnel,proto3" json:"tunnel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigSpec) Reset() {
	*x = ConfigSpec{}
	mi := &file_resource_definitions_siderolink_siderolink_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSpec) ProtoMessage() {}

func (x *ConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_siderolink_siderolink_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSpec.ProtoReflect.Descriptor instead.
func (*ConfigSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_siderolink_siderolink_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSpec) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

func (x *ConfigSpec) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ConfigSpec) GetJoinToken() string {
	if x != nil {
		return x.JoinToken
	}
	return ""
}

func (x *ConfigSpec) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *ConfigSpec) GetTunnel() bool {
	if x != nil {
		return x.Tunnel
	}
	return false
}

// StatusSpec describes Siderolink status.
type StatusSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Host          string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Connected     bool                   `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusSpec) Reset() {
	*x = StatusSpec{}
	mi := &file_resource_definitions_siderolink_siderolink_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSpec) ProtoMessage() {}

func (x *StatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_siderolink_siderolink_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSpec.ProtoReflect.Descriptor instead.
func (*StatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_siderolink_siderolink_proto_rawDescGZIP(), []int{1}
}

func (x *StatusSpec) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *StatusSpec) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

// TunnelSpec describes Siderolink GRPC Tunnel configuration.
type TunnelSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiEndpoint   string                 `protobuf:"bytes,1,opt,name=api_endpoint,json=apiEndpoint,proto3" json:"api_endpoint,omitempty"`
	LinkName      string                 `protobuf:"bytes,2,opt,name=link_name,json=linkName,proto3" json:"link_name,omitempty"`
	Mtu           int64                  `protobuf:"varint,3,opt,name=mtu,proto3" json:"mtu,omitempty"`
	NodeAddress   *common.NetIPPort      `protobuf:"bytes,4,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelSpec) Reset() {
	*x = TunnelSpec{}
	mi := &file_resource_definitions_siderolink_siderolink_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelSpec) ProtoMessage() {}

func (x *TunnelSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_siderolink_siderolink_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelSpec.ProtoReflect.Descriptor instead.
func (*TunnelSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_siderolink_siderolink_proto_rawDescGZIP(), []int{2}
}

func (x *TunnelSpec) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

func (x *TunnelSpec) GetLinkName() string {
	if x != nil {
		return x.LinkName
	}
	return ""
}

func (x *TunnelSpec) GetMtu() int64 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *TunnelSpec) GetNodeAddress() *common.NetIPPort {
	if x != nil {
		return x.NodeAddress
	}
	return nil
}

var File_resource_definitions_siderolink_siderolink_proto protoreflect.FileDescriptor

var file_resource_definitions_siderolink_siderolink_proto_rawDesc = []byte{
	0x0a, 0x30, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x25, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96,
	0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x3e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70,
	0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69,
	0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x34, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x6f, 0x72,
	0x74, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x7e,
	0x0a, 0x2d, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x5a,
	0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65,
	0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_siderolink_siderolink_proto_rawDescOnce sync.Once
	file_resource_definitions_siderolink_siderolink_proto_rawDescData = file_resource_definitions_siderolink_siderolink_proto_rawDesc
)

func file_resource_definitions_siderolink_siderolink_proto_rawDescGZIP() []byte {
	file_resource_definitions_siderolink_siderolink_proto_rawDescOnce.Do(func() {
		file_resource_definitions_siderolink_siderolink_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_siderolink_siderolink_proto_rawDescData)
	})
	return file_resource_definitions_siderolink_siderolink_proto_rawDescData
}

var file_resource_definitions_siderolink_siderolink_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resource_definitions_siderolink_siderolink_proto_goTypes = []any{
	(*ConfigSpec)(nil),       // 0: talos.resource.definitions.siderolink.ConfigSpec
	(*StatusSpec)(nil),       // 1: talos.resource.definitions.siderolink.StatusSpec
	(*TunnelSpec)(nil),       // 2: talos.resource.definitions.siderolink.TunnelSpec
	(*common.NetIPPort)(nil), // 3: common.NetIPPort
}
var file_resource_definitions_siderolink_siderolink_proto_depIdxs = []int32{
	3, // 0: talos.resource.definitions.siderolink.TunnelSpec.node_address:type_name -> common.NetIPPort
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resource_definitions_siderolink_siderolink_proto_init() }
func file_resource_definitions_siderolink_siderolink_proto_init() {
	if File_resource_definitions_siderolink_siderolink_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_definitions_siderolink_siderolink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_siderolink_siderolink_proto_goTypes,
		DependencyIndexes: file_resource_definitions_siderolink_siderolink_proto_depIdxs,
		MessageInfos:      file_resource_definitions_siderolink_siderolink_proto_msgTypes,
	}.Build()
	File_resource_definitions_siderolink_siderolink_proto = out.File
	file_resource_definitions_siderolink_siderolink_proto_rawDesc = nil
	file_resource_definitions_siderolink_siderolink_proto_goTypes = nil
	file_resource_definitions_siderolink_siderolink_proto_depIdxs = nil
}
