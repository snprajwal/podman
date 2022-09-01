// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: eventpoll.proto

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

type EventpollTfdEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Tfd    *uint32 `protobuf:"varint,2,req,name=tfd" json:"tfd,omitempty"`
	Events *uint32 `protobuf:"varint,3,req,name=events" json:"events,omitempty"`
	Data   *uint64 `protobuf:"varint,4,req,name=data" json:"data,omitempty"`
	// to find dup'ed target files
	Dev   *uint32 `protobuf:"varint,5,opt,name=dev" json:"dev,omitempty"`
	Inode *uint64 `protobuf:"varint,6,opt,name=inode" json:"inode,omitempty"`
	Pos   *uint64 `protobuf:"varint,7,opt,name=pos" json:"pos,omitempty"`
}

func (x *EventpollTfdEntry) Reset() {
	*x = EventpollTfdEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventpoll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventpollTfdEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventpollTfdEntry) ProtoMessage() {}

func (x *EventpollTfdEntry) ProtoReflect() protoreflect.Message {
	mi := &file_eventpoll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventpollTfdEntry.ProtoReflect.Descriptor instead.
func (*EventpollTfdEntry) Descriptor() ([]byte, []int) {
	return file_eventpoll_proto_rawDescGZIP(), []int{0}
}

func (x *EventpollTfdEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *EventpollTfdEntry) GetTfd() uint32 {
	if x != nil && x.Tfd != nil {
		return *x.Tfd
	}
	return 0
}

func (x *EventpollTfdEntry) GetEvents() uint32 {
	if x != nil && x.Events != nil {
		return *x.Events
	}
	return 0
}

func (x *EventpollTfdEntry) GetData() uint64 {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return 0
}

func (x *EventpollTfdEntry) GetDev() uint32 {
	if x != nil && x.Dev != nil {
		return *x.Dev
	}
	return 0
}

func (x *EventpollTfdEntry) GetInode() uint64 {
	if x != nil && x.Inode != nil {
		return *x.Inode
	}
	return 0
}

func (x *EventpollTfdEntry) GetPos() uint64 {
	if x != nil && x.Pos != nil {
		return *x.Pos
	}
	return 0
}

type EventpollFileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *uint32              `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags *uint32              `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Fown  *FownEntry           `protobuf:"bytes,3,req,name=fown" json:"fown,omitempty"`
	Tfd   []*EventpollTfdEntry `protobuf:"bytes,4,rep,name=tfd" json:"tfd,omitempty"`
}

func (x *EventpollFileEntry) Reset() {
	*x = EventpollFileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventpoll_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventpollFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventpollFileEntry) ProtoMessage() {}

func (x *EventpollFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_eventpoll_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventpollFileEntry.ProtoReflect.Descriptor instead.
func (*EventpollFileEntry) Descriptor() ([]byte, []int) {
	return file_eventpoll_proto_rawDescGZIP(), []int{1}
}

func (x *EventpollFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *EventpollFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *EventpollFileEntry) GetFown() *FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *EventpollFileEntry) GetTfd() []*EventpollTfdEntry {
	if x != nil {
		return x.Tfd
	}
	return nil
}

var File_eventpoll_proto protoreflect.FileDescriptor

var file_eventpoll_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6c, 0x6c, 0x5f,
	0x74, 0x66, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x66, 0x64,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x66, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x65, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x6f,
	0x73, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6c, 0x6c, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x32, 0x0a, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x66, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x66, 0x6f, 0x77, 0x6e, 0x12, 0x39, 0x0a, 0x03, 0x74, 0x66, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6c, 0x6c,
	0x5f, 0x74, 0x66, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x74, 0x66, 0x64,
}

var (
	file_eventpoll_proto_rawDescOnce sync.Once
	file_eventpoll_proto_rawDescData = file_eventpoll_proto_rawDesc
)

func file_eventpoll_proto_rawDescGZIP() []byte {
	file_eventpoll_proto_rawDescOnce.Do(func() {
		file_eventpoll_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventpoll_proto_rawDescData)
	})
	return file_eventpoll_proto_rawDescData
}

var file_eventpoll_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eventpoll_proto_goTypes = []interface{}{
	(*EventpollTfdEntry)(nil),  // 0: checkpoint.restore.eventpoll_tfd_entry
	(*EventpollFileEntry)(nil), // 1: checkpoint.restore.eventpoll_file_entry
	(*FownEntry)(nil),          // 2: checkpoint.restore.fown_entry
}
var file_eventpoll_proto_depIdxs = []int32{
	2, // 0: checkpoint.restore.eventpoll_file_entry.fown:type_name -> checkpoint.restore.fown_entry
	0, // 1: checkpoint.restore.eventpoll_file_entry.tfd:type_name -> checkpoint.restore.eventpoll_tfd_entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eventpoll_proto_init() }
func file_eventpoll_proto_init() {
	if File_eventpoll_proto != nil {
		return
	}
	file_fown_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eventpoll_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventpollTfdEntry); i {
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
		file_eventpoll_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventpollFileEntry); i {
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
			RawDescriptor: file_eventpoll_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventpoll_proto_goTypes,
		DependencyIndexes: file_eventpoll_proto_depIdxs,
		MessageInfos:      file_eventpoll_proto_msgTypes,
	}.Build()
	File_eventpoll_proto = out.File
	file_eventpoll_proto_rawDesc = nil
	file_eventpoll_proto_goTypes = nil
	file_eventpoll_proto_depIdxs = nil
}
