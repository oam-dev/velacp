// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/proto/model/definitions.proto

package model

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Definition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc       string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Jsonschema string `protobuf:"bytes,3,opt,name=jsonschema,proto3" json:"jsonschema,omitempty"`
	Namespace  string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Definition) Reset() {
	*x = Definition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_definitions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Definition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Definition) ProtoMessage() {}

func (x *Definition) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_definitions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Definition.ProtoReflect.Descriptor instead.
func (*Definition) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_definitions_proto_rawDescGZIP(), []int{0}
}

func (x *Definition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Definition) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Definition) GetJsonschema() string {
	if x != nil {
		return x.Jsonschema
	}
	return ""
}

func (x *Definition) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type DefinitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definitions []*Definition `protobuf:"bytes,1,rep,name=definitions,proto3" json:"definitions,omitempty"`
}

func (x *DefinitionsResponse) Reset() {
	*x = DefinitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_model_definitions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefinitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefinitionsResponse) ProtoMessage() {}

func (x *DefinitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_model_definitions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefinitionsResponse.ProtoReflect.Descriptor instead.
func (*DefinitionsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_model_definitions_proto_rawDescGZIP(), []int{1}
}

func (x *DefinitionsResponse) GetDefinitions() []*Definition {
	if x != nil {
		return x.Definitions
	}
	return nil
}

var File_pkg_proto_model_definitions_proto protoreflect.FileDescriptor

var file_pkg_proto_model_definitions_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x72, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x73, 0x6f,
	0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a,
	0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x61, 0x6d, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x76, 0x65, 0x6c, 0x61, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_proto_model_definitions_proto_rawDescOnce sync.Once
	file_pkg_proto_model_definitions_proto_rawDescData = file_pkg_proto_model_definitions_proto_rawDesc
)

func file_pkg_proto_model_definitions_proto_rawDescGZIP() []byte {
	file_pkg_proto_model_definitions_proto_rawDescOnce.Do(func() {
		file_pkg_proto_model_definitions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_model_definitions_proto_rawDescData)
	})
	return file_pkg_proto_model_definitions_proto_rawDescData
}

var file_pkg_proto_model_definitions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_model_definitions_proto_goTypes = []interface{}{
	(*Definition)(nil),          // 0: vela.api.model.Definition
	(*DefinitionsResponse)(nil), // 1: vela.api.model.DefinitionsResponse
}
var file_pkg_proto_model_definitions_proto_depIdxs = []int32{
	0, // 0: vela.api.model.DefinitionsResponse.definitions:type_name -> vela.api.model.Definition
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_model_definitions_proto_init() }
func file_pkg_proto_model_definitions_proto_init() {
	if File_pkg_proto_model_definitions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_model_definitions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Definition); i {
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
		file_pkg_proto_model_definitions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefinitionsResponse); i {
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
			RawDescriptor: file_pkg_proto_model_definitions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_model_definitions_proto_goTypes,
		DependencyIndexes: file_pkg_proto_model_definitions_proto_depIdxs,
		MessageInfos:      file_pkg_proto_model_definitions_proto_msgTypes,
	}.Build()
	File_pkg_proto_model_definitions_proto = out.File
	file_pkg_proto_model_definitions_proto_rawDesc = nil
	file_pkg_proto_model_definitions_proto_goTypes = nil
	file_pkg_proto_model_definitions_proto_depIdxs = nil
}
