// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: frontdeskpb/frontdesk.proto

package frontdeskpb

import (
	roompb "github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomType  string                 `protobuf:"bytes,1,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	StartDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *CheckAvailabilityRequest) Reset() {
	*x = CheckAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontdeskpb_frontdesk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAvailabilityRequest) ProtoMessage() {}

func (x *CheckAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frontdeskpb_frontdesk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*CheckAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_frontdeskpb_frontdesk_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAvailabilityRequest) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *CheckAvailabilityRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *CheckAvailabilityRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type CheckAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*roompb.Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *CheckAvailabilityResponse) Reset() {
	*x = CheckAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontdeskpb_frontdesk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAvailabilityResponse) ProtoMessage() {}

func (x *CheckAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_frontdeskpb_frontdesk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*CheckAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_frontdeskpb_frontdesk_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAvailabilityResponse) GetRooms() []*roompb.Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_frontdeskpb_frontdesk_proto protoreflect.FileDescriptor

var file_frontdeskpb_frontdesk_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b, 0x70, 0x62, 0x2f, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72, 0x6f, 0x6f,
	0x6d, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9,
	0x01, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3f, 0x0a, 0x19, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x32, 0x78, 0x0a, 0x10, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x64, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x2f, 0x6f, 0x74, 0x65, 0x6c, 0x2d,
	0x74, 0x72, 0x69, 0x76, 0x61, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x64, 0x65, 0x73, 0x6b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frontdeskpb_frontdesk_proto_rawDescOnce sync.Once
	file_frontdeskpb_frontdesk_proto_rawDescData = file_frontdeskpb_frontdesk_proto_rawDesc
)

func file_frontdeskpb_frontdesk_proto_rawDescGZIP() []byte {
	file_frontdeskpb_frontdesk_proto_rawDescOnce.Do(func() {
		file_frontdeskpb_frontdesk_proto_rawDescData = protoimpl.X.CompressGZIP(file_frontdeskpb_frontdesk_proto_rawDescData)
	})
	return file_frontdeskpb_frontdesk_proto_rawDescData
}

var file_frontdeskpb_frontdesk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_frontdeskpb_frontdesk_proto_goTypes = []interface{}{
	(*CheckAvailabilityRequest)(nil),  // 0: frontdeskpb.CheckAvailabilityRequest
	(*CheckAvailabilityResponse)(nil), // 1: frontdeskpb.CheckAvailabilityResponse
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
	(*roompb.Room)(nil),               // 3: roompb.Room
}
var file_frontdeskpb_frontdesk_proto_depIdxs = []int32{
	2, // 0: frontdeskpb.CheckAvailabilityRequest.start_date:type_name -> google.protobuf.Timestamp
	2, // 1: frontdeskpb.CheckAvailabilityRequest.end_date:type_name -> google.protobuf.Timestamp
	3, // 2: frontdeskpb.CheckAvailabilityResponse.rooms:type_name -> roompb.Room
	0, // 3: frontdeskpb.FrontdeskService.CheckAvailability:input_type -> frontdeskpb.CheckAvailabilityRequest
	1, // 4: frontdeskpb.FrontdeskService.CheckAvailability:output_type -> frontdeskpb.CheckAvailabilityResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_frontdeskpb_frontdesk_proto_init() }
func file_frontdeskpb_frontdesk_proto_init() {
	if File_frontdeskpb_frontdesk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frontdeskpb_frontdesk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAvailabilityRequest); i {
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
		file_frontdeskpb_frontdesk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAvailabilityResponse); i {
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
			RawDescriptor: file_frontdeskpb_frontdesk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_frontdeskpb_frontdesk_proto_goTypes,
		DependencyIndexes: file_frontdeskpb_frontdesk_proto_depIdxs,
		MessageInfos:      file_frontdeskpb_frontdesk_proto_msgTypes,
	}.Build()
	File_frontdeskpb_frontdesk_proto = out.File
	file_frontdeskpb_frontdesk_proto_rawDesc = nil
	file_frontdeskpb_frontdesk_proto_goTypes = nil
	file_frontdeskpb_frontdesk_proto_depIdxs = nil
}
