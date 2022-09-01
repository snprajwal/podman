// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: ns.proto

package images

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

type NsFileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	NsId    *uint32 `protobuf:"varint,2,req,name=ns_id,json=nsId" json:"ns_id,omitempty"`
	NsCflag *uint32 `protobuf:"varint,3,req,name=ns_cflag,json=nsCflag" json:"ns_cflag,omitempty"`
	Flags   *uint32 `protobuf:"varint,4,req,name=flags" json:"flags,omitempty"`
}

func (x *NsFileEntry) Reset() {
	*x = NsFileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsFileEntry) ProtoMessage() {}

func (x *NsFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_ns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsFileEntry.ProtoReflect.Descriptor instead.
func (*NsFileEntry) Descriptor() ([]byte, []int) {
	return file_ns_proto_rawDescGZIP(), []int{0}
}

func (x *NsFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *NsFileEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

func (x *NsFileEntry) GetNsCflag() uint32 {
	if x != nil && x.NsCflag != nil {
		return *x.NsCflag
	}
	return 0
}

func (x *NsFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

var File_ns_proto protoreflect.FileDescriptor

var file_ns_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x65,
	0x0a, 0x0d, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x13, 0x0a, 0x05, 0x6e, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04,
	0x6e, 0x73, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x73, 0x5f, 0x63, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x73, 0x43, 0x66, 0x6c, 0x61, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73,
}

var (
	file_ns_proto_rawDescOnce sync.Once
	file_ns_proto_rawDescData = file_ns_proto_rawDesc
)

func file_ns_proto_rawDescGZIP() []byte {
	file_ns_proto_rawDescOnce.Do(func() {
		file_ns_proto_rawDescData = protoimpl.X.CompressGZIP(file_ns_proto_rawDescData)
	})
	return file_ns_proto_rawDescData
}

var file_ns_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ns_proto_goTypes = []interface{}{
	(*NsFileEntry)(nil), // 0: checkpoint.restore.ns_file_entry
}
var file_ns_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ns_proto_init() }
func file_ns_proto_init() {
	if File_ns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsFileEntry); i {
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
			RawDescriptor: file_ns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ns_proto_goTypes,
		DependencyIndexes: file_ns_proto_depIdxs,
		MessageInfos:      file_ns_proto_msgTypes,
	}.Build()
	File_ns_proto = out.File
	file_ns_proto_rawDesc = nil
	file_ns_proto_goTypes = nil
	file_ns_proto_depIdxs = nil
}
