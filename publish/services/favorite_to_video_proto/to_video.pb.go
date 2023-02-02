// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: to_video.proto

package proto

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

type GetFavoriteStatus_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *GetFavoriteStatus_Request) Reset() {
	*x = GetFavoriteStatus_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavoriteStatus_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavoriteStatus_Request) ProtoMessage() {}

func (x *GetFavoriteStatus_Request) ProtoReflect() protoreflect.Message {
	mi := &file_to_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavoriteStatus_Request.ProtoReflect.Descriptor instead.
func (*GetFavoriteStatus_Request) Descriptor() ([]byte, []int) {
	return file_to_video_proto_rawDescGZIP(), []int{0}
}

func (x *GetFavoriteStatus_Request) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFavoriteStatus_Request) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type GetFavoriteStatus_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFavorite bool `protobuf:"varint,1,opt,name=isFavorite,proto3" json:"isFavorite,omitempty"`
}

func (x *GetFavoriteStatus_Response) Reset() {
	*x = GetFavoriteStatus_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavoriteStatus_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavoriteStatus_Response) ProtoMessage() {}

func (x *GetFavoriteStatus_Response) ProtoReflect() protoreflect.Message {
	mi := &file_to_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavoriteStatus_Response.ProtoReflect.Descriptor instead.
func (*GetFavoriteStatus_Response) Descriptor() ([]byte, []int) {
	return file_to_video_proto_rawDescGZIP(), []int{1}
}

func (x *GetFavoriteStatus_Response) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

var File_to_video_proto protoreflect.FileDescriptor

var file_to_video_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x6f, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x32, 0x70, 0x0a, 0x0e, 0x54, 0x6f, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_to_video_proto_rawDescOnce sync.Once
	file_to_video_proto_rawDescData = file_to_video_proto_rawDesc
)

func file_to_video_proto_rawDescGZIP() []byte {
	file_to_video_proto_rawDescOnce.Do(func() {
		file_to_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_to_video_proto_rawDescData)
	})
	return file_to_video_proto_rawDescData
}

var file_to_video_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_to_video_proto_goTypes = []interface{}{
	(*GetFavoriteStatus_Request)(nil),  // 0: services.GetFavoriteStatus_Request
	(*GetFavoriteStatus_Response)(nil), // 1: services.GetFavoriteStatus_Response
}
var file_to_video_proto_depIdxs = []int32{
	0, // 0: services.ToVideoService.GetFavoriteStatus:input_type -> services.GetFavoriteStatus_Request
	1, // 1: services.ToVideoService.GetFavoriteStatus:output_type -> services.GetFavoriteStatus_Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_to_video_proto_init() }
func file_to_video_proto_init() {
	if File_to_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_to_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavoriteStatus_Request); i {
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
		file_to_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavoriteStatus_Response); i {
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
			RawDescriptor: file_to_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_to_video_proto_goTypes,
		DependencyIndexes: file_to_video_proto_depIdxs,
		MessageInfos:      file_to_video_proto_msgTypes,
	}.Build()
	File_to_video_proto = out.File
	file_to_video_proto_rawDesc = nil
	file_to_video_proto_goTypes = nil
	file_to_video_proto_depIdxs = nil
}
