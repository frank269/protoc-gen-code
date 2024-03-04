// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: enviroment.proto

package extension

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnviromentOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DefaultValue string `protobuf:"bytes,2,opt,name=defaultValue,proto3" json:"defaultValue,omitempty"`
}

func (x *EnviromentOptions) Reset() {
	*x = EnviromentOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enviroment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnviromentOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnviromentOptions) ProtoMessage() {}

func (x *EnviromentOptions) ProtoReflect() protoreflect.Message {
	mi := &file_enviroment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnviromentOptions.ProtoReflect.Descriptor instead.
func (*EnviromentOptions) Descriptor() ([]byte, []int) {
	return file_enviroment_proto_rawDescGZIP(), []int{0}
}

func (x *EnviromentOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnviromentOptions) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

var file_enviroment_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         40001,
		Name:          "extension.enabled",
		Tag:           "varint,40001,opt,name=enabled",
		Filename:      "enviroment.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*EnviromentOptions)(nil),
		Field:         50000,
		Name:          "extension.env",
		Tag:           "bytes,50000,opt,name=env",
		Filename:      "enviroment.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool enabled = 40001;
	E_Enabled = &file_enviroment_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional extension.EnviromentOptions env = 50000;
	E_Env = &file_enviroment_proto_extTypes[1]
)

var File_enviroment_proto protoreflect.FileDescriptor

var file_enviroment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4b, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x3e, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc1, 0xb8, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x4f, 0x0a, 0x03,
	0x65, 0x6e, 0x76, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x61, 0x6e,
	0x6b, 0x32, 0x36, 0x39, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enviroment_proto_rawDescOnce sync.Once
	file_enviroment_proto_rawDescData = file_enviroment_proto_rawDesc
)

func file_enviroment_proto_rawDescGZIP() []byte {
	file_enviroment_proto_rawDescOnce.Do(func() {
		file_enviroment_proto_rawDescData = protoimpl.X.CompressGZIP(file_enviroment_proto_rawDescData)
	})
	return file_enviroment_proto_rawDescData
}

var file_enviroment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enviroment_proto_goTypes = []interface{}{
	(*EnviromentOptions)(nil),           // 0: extension.EnviromentOptions
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
}
var file_enviroment_proto_depIdxs = []int32{
	1, // 0: extension.enabled:extendee -> google.protobuf.MessageOptions
	2, // 1: extension.env:extendee -> google.protobuf.FieldOptions
	0, // 2: extension.env:type_name -> extension.EnviromentOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enviroment_proto_init() }
func file_enviroment_proto_init() {
	if File_enviroment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enviroment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnviromentOptions); i {
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
			RawDescriptor: file_enviroment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_enviroment_proto_goTypes,
		DependencyIndexes: file_enviroment_proto_depIdxs,
		MessageInfos:      file_enviroment_proto_msgTypes,
		ExtensionInfos:    file_enviroment_proto_extTypes,
	}.Build()
	File_enviroment_proto = out.File
	file_enviroment_proto_rawDesc = nil
	file_enviroment_proto_goTypes = nil
	file_enviroment_proto_depIdxs = nil
}
