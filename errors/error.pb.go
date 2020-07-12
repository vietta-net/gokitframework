// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: agokit/pb/error.proto

package errors

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code         uint32            `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	NestedErrors map[string]string `protobuf:"bytes,3,rep,name=nested_errors,json=nestedErrors,proto3" json:"nested_errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Details      []string          `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
	Stack        []byte            `protobuf:"bytes,5,opt,name=stack,proto3" json:"stack,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agokit_pb_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_agokit_pb_error_proto_msgTypes[0]
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
	return file_agokit_pb_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetNestedErrors() map[string]string {
	if x != nil {
		return x.NestedErrors
	}
	return nil
}

func (x *Error) GetDetails() []string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Error) GetStack() []byte {
	if x != nil {
		return x.Stack
	}
	return nil
}

var File_agokit_pb_error_proto protoreflect.FileDescriptor

var file_agokit_pb_error_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22,
	0xec, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x1a, 0x3f, 0x0a,
	0x11, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0f,
	0x5a, 0x0d, 0x61, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agokit_pb_error_proto_rawDescOnce sync.Once
	file_agokit_pb_error_proto_rawDescData = file_agokit_pb_error_proto_rawDesc
)

func file_agokit_pb_error_proto_rawDescGZIP() []byte {
	file_agokit_pb_error_proto_rawDescOnce.Do(func() {
		file_agokit_pb_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_agokit_pb_error_proto_rawDescData)
	})
	return file_agokit_pb_error_proto_rawDescData
}

var file_agokit_pb_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agokit_pb_error_proto_goTypes = []interface{}{
	(*Error)(nil), // 0: errors.Error
	nil,           // 1: errors.Error.NestedErrorsEntry
}
var file_agokit_pb_error_proto_depIdxs = []int32{
	1, // 0: errors.Error.nested_errors:type_name -> errors.Error.NestedErrorsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_agokit_pb_error_proto_init() }
func file_agokit_pb_error_proto_init() {
	if File_agokit_pb_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agokit_pb_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agokit_pb_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agokit_pb_error_proto_goTypes,
		DependencyIndexes: file_agokit_pb_error_proto_depIdxs,
		MessageInfos:      file_agokit_pb_error_proto_msgTypes,
	}.Build()
	File_agokit_pb_error_proto = out.File
	file_agokit_pb_error_proto_rawDesc = nil
	file_agokit_pb_error_proto_goTypes = nil
	file_agokit_pb_error_proto_depIdxs = nil
}
