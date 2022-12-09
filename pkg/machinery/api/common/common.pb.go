// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: common/common.proto

package common

import (
	reflect "reflect"
	sync "sync"

	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_FATAL    Code = 0
	Code_LOCKED   Code = 1
	Code_CANCELED Code = 2
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "FATAL",
		1: "LOCKED",
		2: "CANCELED",
	}
	Code_value = map[string]int32{
		"FATAL":    0,
		"LOCKED":   1,
		"CANCELED": 2,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type ContainerDriver int32

const (
	ContainerDriver_CONTAINERD ContainerDriver = 0
	ContainerDriver_CRI        ContainerDriver = 1
)

// Enum value maps for ContainerDriver.
var (
	ContainerDriver_name = map[int32]string{
		0: "CONTAINERD",
		1: "CRI",
	}
	ContainerDriver_value = map[string]int32{
		"CONTAINERD": 0,
		"CRI":        1,
	}
)

func (x ContainerDriver) Enum() *ContainerDriver {
	p := new(ContainerDriver)
	*p = x
	return p
}

func (x ContainerDriver) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerDriver) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[1].Descriptor()
}

func (ContainerDriver) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[1]
}

func (x ContainerDriver) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerDriver.Descriptor instead.
func (ContainerDriver) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code         `protobuf:"varint,1,opt,name=code,proto3,enum=common.Code" json:"code,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_FATAL
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

// Common metadata message nested in all reply message types
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hostname of the server response comes from (injected by proxy)
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// error is set if request failed to the upstream (rest of response is
	// undefined)
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// error as gRPC Status
	Status *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Metadata) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Metadata) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Bytes    []byte    `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Data) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Data `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *DataResponse) GetMessages() []*Data {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *Empty) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Empty `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *EmptyResponse) GetMessages() []*Empty {
	if x != nil {
		return x.Messages
	}
	return nil
}

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullPath string `protobuf:"bytes,1,opt,name=full_path,json=fullPath,proto3" json:"full_path,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *URL) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

type PEMEncodedCertificateAndKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crt []byte `protobuf:"bytes,1,opt,name=crt,proto3" json:"crt,omitempty"`
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PEMEncodedCertificateAndKey) Reset() {
	*x = PEMEncodedCertificateAndKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PEMEncodedCertificateAndKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PEMEncodedCertificateAndKey) ProtoMessage() {}

func (x *PEMEncodedCertificateAndKey) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PEMEncodedCertificateAndKey.ProtoReflect.Descriptor instead.
func (*PEMEncodedCertificateAndKey) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *PEMEncodedCertificateAndKey) GetCrt() []byte {
	if x != nil {
		return x.Crt
	}
	return nil
}

func (x *PEMEncodedCertificateAndKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type PEMEncodedKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PEMEncodedKey) Reset() {
	*x = PEMEncodedKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PEMEncodedKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PEMEncodedKey) ProtoMessage() {}

func (x *PEMEncodedKey) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PEMEncodedKey.ProtoReflect.Descriptor instead.
func (*PEMEncodedKey) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{8}
}

func (x *PEMEncodedKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type NetIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *NetIP) Reset() {
	*x = NetIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetIP) ProtoMessage() {}

func (x *NetIP) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetIP.ProtoReflect.Descriptor instead.
func (*NetIP) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{9}
}

func (x *NetIP) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

type NetIPPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *NetIPPort) Reset() {
	*x = NetIPPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetIPPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetIPPort) ProtoMessage() {}

func (x *NetIPPort) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetIPPort.ProtoReflect.Descriptor instead.
func (*NetIPPort) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{10}
}

func (x *NetIPPort) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *NetIPPort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type NetIPPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip           []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	PrefixLength int32  `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
}

func (x *NetIPPrefix) Reset() {
	*x = NetIPPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetIPPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetIPPrefix) ProtoMessage() {}

func (x *NetIPPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetIPPrefix.ProtoReflect.Descriptor instead.
func (*NetIPPrefix) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{11}
}

func (x *NetIPPrefix) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *NetIPPrefix) GetPrefixLength() int32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

var file_common_common_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93117,
		Name:          "common.remove_deprecated_message",
		Tag:           "bytes,93117,opt,name=remove_deprecated_message",
		Filename:      "common/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93117,
		Name:          "common.remove_deprecated_field",
		Tag:           "bytes,93117,opt,name=remove_deprecated_field",
		Filename:      "common/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93117,
		Name:          "common.remove_deprecated_enum",
		Tag:           "bytes,93117,opt,name=remove_deprecated_enum",
		Filename:      "common/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93117,
		Name:          "common.remove_deprecated_enum_value",
		Tag:           "bytes,93117,opt,name=remove_deprecated_enum_value",
		Filename:      "common/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93117,
		Name:          "common.remove_deprecated_method",
		Tag:           "bytes,93117,opt,name=remove_deprecated_method",
		Filename:      "common/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93117,
		Name:          "common.remove_deprecated_service",
		Tag:           "bytes,93117,opt,name=remove_deprecated_service",
		Filename:      "common/common.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Indicates the Talos version when this deprecated message will be removed from API.
	//
	// optional string remove_deprecated_message = 93117;
	E_RemoveDeprecatedMessage = &file_common_common_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// Indicates the Talos version when this deprecated filed will be removed from API.
	//
	// optional string remove_deprecated_field = 93117;
	E_RemoveDeprecatedField = &file_common_common_proto_extTypes[1]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// Indicates the Talos version when this deprecated enum will be removed from API.
	//
	// optional string remove_deprecated_enum = 93117;
	E_RemoveDeprecatedEnum = &file_common_common_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// Indicates the Talos version when this deprecated enum value will be removed from API.
	//
	// optional string remove_deprecated_enum_value = 93117;
	E_RemoveDeprecatedEnumValue = &file_common_common_proto_extTypes[3]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// Indicates the Talos version when this deprecated method will be removed from API.
	//
	// optional string remove_deprecated_method = 93117;
	E_RemoveDeprecatedMethod = &file_common_common_proto_extTypes[4]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// Indicates the Talos version when this deprecated service will be removed from API.
	//
	// optional string remove_deprecated_service = 93117;
	E_RemoveDeprecatedService = &file_common_common_proto_extTypes[5]
)

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x68, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x4a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x38,
	0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x3a, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x03, 0x55,
	0x52, 0x4c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x41, 0x0a, 0x1b, 0x50, 0x45, 0x4d, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x21, 0x0a, 0x0d, 0x50, 0x45, 0x4d, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x17, 0x0a, 0x05, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x22, 0x2f,
	0x0a, 0x09, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x42, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x2a, 0x2b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x46,
	0x41, 0x54, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x2a, 0x2a, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52,
	0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x52, 0x49, 0x10, 0x01, 0x3a, 0x5d, 0x0a, 0x19,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xd7, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x57, 0x0a, 0x17, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xd7, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x3a, 0x54, 0x0a, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xd7, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x3a, 0x64, 0x0a, 0x1c, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xd7,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x5a, 0x0a, 0x18, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xd7, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x5d, 0x0a, 0x19,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xd7, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_common_common_proto_goTypes = []interface{}{
	(Code)(0),                             // 0: common.Code
	(ContainerDriver)(0),                  // 1: common.ContainerDriver
	(*Error)(nil),                         // 2: common.Error
	(*Metadata)(nil),                      // 3: common.Metadata
	(*Data)(nil),                          // 4: common.Data
	(*DataResponse)(nil),                  // 5: common.DataResponse
	(*Empty)(nil),                         // 6: common.Empty
	(*EmptyResponse)(nil),                 // 7: common.EmptyResponse
	(*URL)(nil),                           // 8: common.URL
	(*PEMEncodedCertificateAndKey)(nil),   // 9: common.PEMEncodedCertificateAndKey
	(*PEMEncodedKey)(nil),                 // 10: common.PEMEncodedKey
	(*NetIP)(nil),                         // 11: common.NetIP
	(*NetIPPort)(nil),                     // 12: common.NetIPPort
	(*NetIPPrefix)(nil),                   // 13: common.NetIPPrefix
	(*anypb.Any)(nil),                     // 14: google.protobuf.Any
	(*status.Status)(nil),                 // 15: google.rpc.Status
	(*descriptorpb.MessageOptions)(nil),   // 16: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),     // 17: google.protobuf.FieldOptions
	(*descriptorpb.EnumOptions)(nil),      // 18: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 19: google.protobuf.EnumValueOptions
	(*descriptorpb.MethodOptions)(nil),    // 20: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil),   // 21: google.protobuf.ServiceOptions
}
var file_common_common_proto_depIdxs = []int32{
	0,  // 0: common.Error.code:type_name -> common.Code
	14, // 1: common.Error.details:type_name -> google.protobuf.Any
	15, // 2: common.Metadata.status:type_name -> google.rpc.Status
	3,  // 3: common.Data.metadata:type_name -> common.Metadata
	4,  // 4: common.DataResponse.messages:type_name -> common.Data
	3,  // 5: common.Empty.metadata:type_name -> common.Metadata
	6,  // 6: common.EmptyResponse.messages:type_name -> common.Empty
	16, // 7: common.remove_deprecated_message:extendee -> google.protobuf.MessageOptions
	17, // 8: common.remove_deprecated_field:extendee -> google.protobuf.FieldOptions
	18, // 9: common.remove_deprecated_enum:extendee -> google.protobuf.EnumOptions
	19, // 10: common.remove_deprecated_enum_value:extendee -> google.protobuf.EnumValueOptions
	20, // 11: common.remove_deprecated_method:extendee -> google.protobuf.MethodOptions
	21, // 12: common.remove_deprecated_service:extendee -> google.protobuf.ServiceOptions
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	7,  // [7:13] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
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
		file_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_common_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_common_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PEMEncodedCertificateAndKey); i {
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
		file_common_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PEMEncodedKey); i {
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
		file_common_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetIP); i {
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
		file_common_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetIPPort); i {
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
		file_common_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetIPPrefix); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
		ExtensionInfos:    file_common_common_proto_extTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
