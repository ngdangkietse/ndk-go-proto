// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: common/shared.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PCode int32

const (
	PCode_SUCCESS           PCode = 0
	PCode_NOT_FOUND         PCode = -1
	PCode_INVALID_DATA      PCode = -2
	PCode_ALREADY_EXISTS    PCode = -3
	PCode_UNKNOWN_EXCEPTION PCode = -4
)

// Enum value maps for PCode.
var (
	PCode_name = map[int32]string{
		0:  "SUCCESS",
		-1: "NOT_FOUND",
		-2: "INVALID_DATA",
		-3: "ALREADY_EXISTS",
		-4: "UNKNOWN_EXCEPTION",
	}
	PCode_value = map[string]int32{
		"SUCCESS":           0,
		"NOT_FOUND":         -1,
		"INVALID_DATA":      -2,
		"ALREADY_EXISTS":    -3,
		"UNKNOWN_EXCEPTION": -4,
	}
)

func (x PCode) Enum() *PCode {
	p := new(PCode)
	*p = x
	return p
}

func (x PCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_shared_proto_enumTypes[0].Descriptor()
}

func (PCode) Type() protoreflect.EnumType {
	return &file_common_shared_proto_enumTypes[0]
}

func (x PCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PCode.Descriptor instead.
func (PCode) EnumDescriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{0}
}

type PError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors map[string]string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PError) Reset() {
	*x = PError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PError) ProtoMessage() {}

func (x *PError) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PError.ProtoReflect.Descriptor instead.
func (*PError) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{0}
}

func (x *PError) GetErrors() map[string]string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_common_shared_proto protoreflect.FileDescriptor

var file_common_shared_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x77, 0x0a,
	0x06, 0x50, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x84, 0x01, 0x0a, 0x05, 0x50, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x19, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x12, 0x1b, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x53, 0x10, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1e, 0x0a,
	0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x88, 0x01,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0b, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67, 0x6b, 0x69,
	0x65, 0x74, 0x73, 0x65, 0x2f, 0x6e, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xca, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xe2, 0x02, 0x12, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_shared_proto_rawDescOnce sync.Once
	file_common_shared_proto_rawDescData = file_common_shared_proto_rawDesc
)

func file_common_shared_proto_rawDescGZIP() []byte {
	file_common_shared_proto_rawDescOnce.Do(func() {
		file_common_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_shared_proto_rawDescData)
	})
	return file_common_shared_proto_rawDescData
}

var file_common_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_shared_proto_goTypes = []interface{}{
	(PCode)(0),     // 0: common.PCode
	(*PError)(nil), // 1: common.PError
	nil,            // 2: common.PError.ErrorsEntry
}
var file_common_shared_proto_depIdxs = []int32{
	2, // 0: common.PError.errors:type_name -> common.PError.ErrorsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_shared_proto_init() }
func file_common_shared_proto_init() {
	if File_common_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PError); i {
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
			RawDescriptor: file_common_shared_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_shared_proto_goTypes,
		DependencyIndexes: file_common_shared_proto_depIdxs,
		EnumInfos:         file_common_shared_proto_enumTypes,
		MessageInfos:      file_common_shared_proto_msgTypes,
	}.Build()
	File_common_shared_proto = out.File
	file_common_shared_proto_rawDesc = nil
	file_common_shared_proto_goTypes = nil
	file_common_shared_proto_depIdxs = nil
}
