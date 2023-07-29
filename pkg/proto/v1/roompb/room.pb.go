// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: roompb/room.proto

package roompb

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

type CleanStatus int32

const (
	CleanStatus_CLEAN CleanStatus = 0
	CleanStatus_DIRTY CleanStatus = 1
)

// Enum value maps for CleanStatus.
var (
	CleanStatus_name = map[int32]string{
		0: "CLEAN",
		1: "DIRTY",
	}
	CleanStatus_value = map[string]int32{
		"CLEAN": 0,
		"DIRTY": 1,
	}
)

func (x CleanStatus) Enum() *CleanStatus {
	p := new(CleanStatus)
	*p = x
	return p
}

func (x CleanStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CleanStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_roompb_room_proto_enumTypes[0].Descriptor()
}

func (CleanStatus) Type() protoreflect.EnumType {
	return &file_roompb_room_proto_enumTypes[0]
}

func (x CleanStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CleanStatus.Descriptor instead.
func (CleanStatus) EnumDescriptor() ([]byte, []int) {
	return file_roompb_room_proto_rawDescGZIP(), []int{0}
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity    int64       `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	CleanStatus CleanStatus `protobuf:"varint,3,opt,name=clean_status,json=cleanStatus,proto3,enum=roompb.CleanStatus" json:"clean_status,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roompb_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_roompb_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_roompb_room_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Room) GetCleanStatus() CleanStatus {
	if x != nil {
		return x.CleanStatus
	}
	return CleanStatus_CLEAN
}

type ListRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capacity int64 `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Limit    int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRoomsRequest) Reset() {
	*x = ListRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roompb_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsRequest) ProtoMessage() {}

func (x *ListRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_roompb_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsRequest.ProtoReflect.Descriptor instead.
func (*ListRoomsRequest) Descriptor() ([]byte, []int) {
	return file_roompb_room_proto_rawDescGZIP(), []int{1}
}

func (x *ListRoomsRequest) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *ListRoomsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRoomsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListRoomsResponse) Reset() {
	*x = ListRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roompb_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsResponse) ProtoMessage() {}

func (x *ListRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roompb_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsResponse.ProtoReflect.Descriptor instead.
func (*ListRoomsResponse) Descriptor() ([]byte, []int) {
	return file_roompb_room_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *ListRoomsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_roompb_room_proto protoreflect.FileDescriptor

var file_roompb_room_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x22, 0x6a, 0x0a, 0x04, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x36, 0x0a, 0x0c, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x43,
	0x6c, 0x65, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x63, 0x6c, 0x65, 0x61,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x23, 0x0a, 0x0b, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x49, 0x52, 0x54, 0x59, 0x10, 0x01, 0x32, 0x4f, 0x0a, 0x0b, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x2f,
	0x6f, 0x74, 0x65, 0x6c, 0x2d, 0x74, 0x72, 0x69, 0x76, 0x61, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_roompb_room_proto_rawDescOnce sync.Once
	file_roompb_room_proto_rawDescData = file_roompb_room_proto_rawDesc
)

func file_roompb_room_proto_rawDescGZIP() []byte {
	file_roompb_room_proto_rawDescOnce.Do(func() {
		file_roompb_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_roompb_room_proto_rawDescData)
	})
	return file_roompb_room_proto_rawDescData
}

var file_roompb_room_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_roompb_room_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_roompb_room_proto_goTypes = []interface{}{
	(CleanStatus)(0),          // 0: roompb.CleanStatus
	(*Room)(nil),              // 1: roompb.Room
	(*ListRoomsRequest)(nil),  // 2: roompb.ListRoomsRequest
	(*ListRoomsResponse)(nil), // 3: roompb.ListRoomsResponse
}
var file_roompb_room_proto_depIdxs = []int32{
	0, // 0: roompb.Room.clean_status:type_name -> roompb.CleanStatus
	1, // 1: roompb.ListRoomsResponse.rooms:type_name -> roompb.Room
	2, // 2: roompb.RoomService.ListRooms:input_type -> roompb.ListRoomsRequest
	3, // 3: roompb.RoomService.ListRooms:output_type -> roompb.ListRoomsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_roompb_room_proto_init() }
func file_roompb_room_proto_init() {
	if File_roompb_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roompb_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_roompb_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsRequest); i {
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
		file_roompb_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsResponse); i {
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
			RawDescriptor: file_roompb_room_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roompb_room_proto_goTypes,
		DependencyIndexes: file_roompb_room_proto_depIdxs,
		EnumInfos:         file_roompb_room_proto_enumTypes,
		MessageInfos:      file_roompb_room_proto_msgTypes,
	}.Build()
	File_roompb_room_proto = out.File
	file_roompb_room_proto_rawDesc = nil
	file_roompb_room_proto_goTypes = nil
	file_roompb_room_proto_depIdxs = nil
}
