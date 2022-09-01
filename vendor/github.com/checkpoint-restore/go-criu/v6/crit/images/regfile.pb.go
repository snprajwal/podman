// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: regfile.proto

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

type RegFileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *uint32    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags *uint32    `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Pos   *uint64    `protobuf:"varint,3,req,name=pos" json:"pos,omitempty"`
	Fown  *FownEntry `protobuf:"bytes,5,req,name=fown" json:"fown,omitempty"`
	Name  *string    `protobuf:"bytes,6,req,name=name" json:"name,omitempty"`
	MntId *int32     `protobuf:"zigzag32,7,opt,name=mnt_id,json=mntId,def=-1" json:"mnt_id,omitempty"`
	Size  *uint64    `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
	Ext   *bool      `protobuf:"varint,9,opt,name=ext" json:"ext,omitempty"`
	Mode  *uint32    `protobuf:"varint,10,opt,name=mode" json:"mode,omitempty"`
	// This field stores the build-ID of the file if it could be obtained.
	BuildId []uint32 `protobuf:"varint,11,rep,name=build_id,json=buildId" json:"build_id,omitempty"`
	// This field stores the CRC32C checksum of the file if it could be obtained.
	Checksum *uint32 `protobuf:"varint,12,opt,name=checksum" json:"checksum,omitempty"`
	// This field stores the configuration of bytes which were used in the
	// calculation of the checksum, if it could be obtained.
	ChecksumConfig *uint32 `protobuf:"varint,13,opt,name=checksum_config,json=checksumConfig" json:"checksum_config,omitempty"`
	// This field stores the checksum parameter if it was used in the calculation
	// of the checksum, if it could be obtained.
	ChecksumParameter *uint32 `protobuf:"varint,14,opt,name=checksum_parameter,json=checksumParameter" json:"checksum_parameter,omitempty"`
}

// Default values for RegFileEntry fields.
const (
	Default_RegFileEntry_MntId = int32(-1)
)

func (x *RegFileEntry) Reset() {
	*x = RegFileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regfile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegFileEntry) ProtoMessage() {}

func (x *RegFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_regfile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegFileEntry.ProtoReflect.Descriptor instead.
func (*RegFileEntry) Descriptor() ([]byte, []int) {
	return file_regfile_proto_rawDescGZIP(), []int{0}
}

func (x *RegFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *RegFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *RegFileEntry) GetPos() uint64 {
	if x != nil && x.Pos != nil {
		return *x.Pos
	}
	return 0
}

func (x *RegFileEntry) GetFown() *FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *RegFileEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RegFileEntry) GetMntId() int32 {
	if x != nil && x.MntId != nil {
		return *x.MntId
	}
	return Default_RegFileEntry_MntId
}

func (x *RegFileEntry) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *RegFileEntry) GetExt() bool {
	if x != nil && x.Ext != nil {
		return *x.Ext
	}
	return false
}

func (x *RegFileEntry) GetMode() uint32 {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return 0
}

func (x *RegFileEntry) GetBuildId() []uint32 {
	if x != nil {
		return x.BuildId
	}
	return nil
}

func (x *RegFileEntry) GetChecksum() uint32 {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return 0
}

func (x *RegFileEntry) GetChecksumConfig() uint32 {
	if x != nil && x.ChecksumConfig != nil {
		return *x.ChecksumConfig
	}
	return 0
}

func (x *RegFileEntry) GetChecksumParameter() uint32 {
	if x != nil && x.ChecksumParameter != nil {
		return *x.ChecksumParameter
	}
	return 0
}

var File_regfile_proto protoreflect.FileDescriptor

var file_regfile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x0e,
	0x72, 0x65, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x10, 0xd2,
	0x3f, 0x0d, 0x1a, 0x0b, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x52,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x66, 0x6f, 0x77, 0x6e,
	0x18, 0x05, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x6f, 0x77, 0x6e,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x06, 0x6d, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11,
	0x3a, 0x02, 0x2d, 0x31, 0x52, 0x05, 0x6d, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x78,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72,
}

var (
	file_regfile_proto_rawDescOnce sync.Once
	file_regfile_proto_rawDescData = file_regfile_proto_rawDesc
)

func file_regfile_proto_rawDescGZIP() []byte {
	file_regfile_proto_rawDescOnce.Do(func() {
		file_regfile_proto_rawDescData = protoimpl.X.CompressGZIP(file_regfile_proto_rawDescData)
	})
	return file_regfile_proto_rawDescData
}

var file_regfile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_regfile_proto_goTypes = []interface{}{
	(*RegFileEntry)(nil), // 0: checkpoint.restore.reg_file_entry
	(*FownEntry)(nil),    // 1: checkpoint.restore.fown_entry
}
var file_regfile_proto_depIdxs = []int32{
	1, // 0: checkpoint.restore.reg_file_entry.fown:type_name -> checkpoint.restore.fown_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_regfile_proto_init() }
func file_regfile_proto_init() {
	if File_regfile_proto != nil {
		return
	}
	file_opts_proto_init()
	file_fown_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regfile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegFileEntry); i {
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
			RawDescriptor: file_regfile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regfile_proto_goTypes,
		DependencyIndexes: file_regfile_proto_depIdxs,
		MessageInfos:      file_regfile_proto_msgTypes,
	}.Build()
	File_regfile_proto = out.File
	file_regfile_proto_rawDesc = nil
	file_regfile_proto_goTypes = nil
	file_regfile_proto_depIdxs = nil
}
