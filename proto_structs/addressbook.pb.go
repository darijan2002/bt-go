// See README.md for information and build instructions.
//
// Note: START and END tags are used in comments to define sections used in
// tutorials.  They are not part of the syntax for Protocol Buffers.
//
// To get an in-depth walkthrough of this file and the related examples, see:
// https://developers.google.com/protocol-buffers/docs/tutorials

// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.6
// source: addressbook.proto

package proto_structs

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

// [START messages]
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length uint64 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addressbook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_addressbook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_addressbook_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FileInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*FileInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *FileInfos) Reset() {
	*x = FileInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addressbook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfos) ProtoMessage() {}

func (x *FileInfos) ProtoReflect() protoreflect.Message {
	mi := &file_addressbook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfos.ProtoReflect.Descriptor instead.
func (*FileInfos) Descriptor() ([]byte, []int) {
	return file_addressbook_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfos) GetInfos() []*FileInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type MetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Announce string         `protobuf:"bytes,1,opt,name=announce,proto3" json:"announce,omitempty"`
	Info     *MetaInfo_Info `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *MetaInfo) Reset() {
	*x = MetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addressbook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaInfo) ProtoMessage() {}

func (x *MetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_addressbook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaInfo.ProtoReflect.Descriptor instead.
func (*MetaInfo) Descriptor() ([]byte, []int) {
	return file_addressbook_proto_rawDescGZIP(), []int{2}
}

func (x *MetaInfo) GetAnnounce() string {
	if x != nil {
		return x.Announce
	}
	return ""
}

func (x *MetaInfo) GetInfo() *MetaInfo_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type MetaInfo_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PieceLength int32    `protobuf:"varint,2,opt,name=piece_length,json=pieceLength,proto3" json:"piece_length,omitempty"`
	Pieces      []string `protobuf:"bytes,3,rep,name=pieces,proto3" json:"pieces,omitempty"`
	// Types that are assignable to Data:
	//
	//	*MetaInfo_Info_Length
	//	*MetaInfo_Info_Files
	Data isMetaInfo_Info_Data `protobuf_oneof:"Data"`
}

func (x *MetaInfo_Info) Reset() {
	*x = MetaInfo_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addressbook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaInfo_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaInfo_Info) ProtoMessage() {}

func (x *MetaInfo_Info) ProtoReflect() protoreflect.Message {
	mi := &file_addressbook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaInfo_Info.ProtoReflect.Descriptor instead.
func (*MetaInfo_Info) Descriptor() ([]byte, []int) {
	return file_addressbook_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MetaInfo_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaInfo_Info) GetPieceLength() int32 {
	if x != nil {
		return x.PieceLength
	}
	return 0
}

func (x *MetaInfo_Info) GetPieces() []string {
	if x != nil {
		return x.Pieces
	}
	return nil
}

func (m *MetaInfo_Info) GetData() isMetaInfo_Info_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *MetaInfo_Info) GetLength() uint64 {
	if x, ok := x.GetData().(*MetaInfo_Info_Length); ok {
		return x.Length
	}
	return 0
}

func (x *MetaInfo_Info) GetFiles() *FileInfos {
	if x, ok := x.GetData().(*MetaInfo_Info_Files); ok {
		return x.Files
	}
	return nil
}

type isMetaInfo_Info_Data interface {
	isMetaInfo_Info_Data()
}

type MetaInfo_Info_Length struct {
	Length uint64 `protobuf:"varint,4,opt,name=length,proto3,oneof"`
}

type MetaInfo_Info_Files struct {
	Files *FileInfos `protobuf:"bytes,5,opt,name=files,proto3,oneof"`
}

func (*MetaInfo_Info_Length) isMetaInfo_Info_Data() {}

func (*MetaInfo_Info_Files) isMetaInfo_Info_Data() {}

var File_addressbook_proto protoreflect.FileDescriptor

var file_addressbook_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x3a, 0x0a, 0x09, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0xa9, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x69, 0x65, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x69, 0x65, 0x63, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x48, 0x00, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_addressbook_proto_rawDescOnce sync.Once
	file_addressbook_proto_rawDescData = file_addressbook_proto_rawDesc
)

func file_addressbook_proto_rawDescGZIP() []byte {
	file_addressbook_proto_rawDescOnce.Do(func() {
		file_addressbook_proto_rawDescData = protoimpl.X.CompressGZIP(file_addressbook_proto_rawDescData)
	})
	return file_addressbook_proto_rawDescData
}

var file_addressbook_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_addressbook_proto_goTypes = []interface{}{
	(*FileInfo)(nil),      // 0: proto_structs.FileInfo
	(*FileInfos)(nil),     // 1: proto_structs.FileInfos
	(*MetaInfo)(nil),      // 2: proto_structs.MetaInfo
	(*MetaInfo_Info)(nil), // 3: proto_structs.MetaInfo.Info
}
var file_addressbook_proto_depIdxs = []int32{
	0, // 0: proto_structs.FileInfos.infos:type_name -> proto_structs.FileInfo
	3, // 1: proto_structs.MetaInfo.info:type_name -> proto_structs.MetaInfo.Info
	1, // 2: proto_structs.MetaInfo.Info.files:type_name -> proto_structs.FileInfos
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_addressbook_proto_init() }
func file_addressbook_proto_init() {
	if File_addressbook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_addressbook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_addressbook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfos); i {
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
		file_addressbook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaInfo); i {
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
		file_addressbook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaInfo_Info); i {
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
	file_addressbook_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MetaInfo_Info_Length)(nil),
		(*MetaInfo_Info_Files)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_addressbook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_addressbook_proto_goTypes,
		DependencyIndexes: file_addressbook_proto_depIdxs,
		MessageInfos:      file_addressbook_proto_msgTypes,
	}.Build()
	File_addressbook_proto = out.File
	file_addressbook_proto_rawDesc = nil
	file_addressbook_proto_goTypes = nil
	file_addressbook_proto_depIdxs = nil
}
