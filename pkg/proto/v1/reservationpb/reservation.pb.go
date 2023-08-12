// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: reservationpb/reservation.proto

package reservationpb

import (
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

type ReservationStatus int32

const (
	ReservationStatus_RESERVED    ReservationStatus = 0
	ReservationStatus_CHECKED_IN  ReservationStatus = 1
	ReservationStatus_CHECKED_OUT ReservationStatus = 2
	ReservationStatus_CANCELLED   ReservationStatus = 3
)

// Enum value maps for ReservationStatus.
var (
	ReservationStatus_name = map[int32]string{
		0: "RESERVED",
		1: "CHECKED_IN",
		2: "CHECKED_OUT",
		3: "CANCELLED",
	}
	ReservationStatus_value = map[string]int32{
		"RESERVED":    0,
		"CHECKED_IN":  1,
		"CHECKED_OUT": 2,
		"CANCELLED":   3,
	}
)

func (x ReservationStatus) Enum() *ReservationStatus {
	p := new(ReservationStatus)
	*p = x
	return p
}

func (x ReservationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReservationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_reservationpb_reservation_proto_enumTypes[0].Descriptor()
}

func (ReservationStatus) Type() protoreflect.EnumType {
	return &file_reservationpb_reservation_proto_enumTypes[0]
}

func (x ReservationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReservationStatus.Descriptor instead.
func (ReservationStatus) EnumDescriptor() ([]byte, []int) {
	return file_reservationpb_reservation_proto_rawDescGZIP(), []int{0}
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomId    string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Status    ReservationStatus      `protobuf:"varint,3,opt,name=status,proto3,enum=reservationpb.ReservationStatus" json:"status,omitempty"`
	StartDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservationpb_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservationpb_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservationpb_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Reservation) GetStatus() ReservationStatus {
	if x != nil {
		return x.Status
	}
	return ReservationStatus_RESERVED
}

func (x *Reservation) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Reservation) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Reservation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ListReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset    int64                  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	StartDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *ListReservationRequest) Reset() {
	*x = ListReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservationpb_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReservationRequest) ProtoMessage() {}

func (x *ListReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservationpb_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReservationRequest.ProtoReflect.Descriptor instead.
func (*ListReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservationpb_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *ListReservationRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListReservationRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListReservationRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ListReservationRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type ListReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
	Total        int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListReservationResponse) Reset() {
	*x = ListReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservationpb_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReservationResponse) ProtoMessage() {}

func (x *ListReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservationpb_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReservationResponse.ProtoReflect.Descriptor instead.
func (*ListReservationResponse) Descriptor() ([]byte, []int) {
	return file_reservationpb_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *ListReservationResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

func (x *ListReservationResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId    string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	StartDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *CreateReservationRequest) Reset() {
	*x = CreateReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservationpb_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationRequest) ProtoMessage() {}

func (x *CreateReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservationpb_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationRequest.ProtoReflect.Descriptor instead.
func (*CreateReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservationpb_reservation_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReservationRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CreateReservationRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *CreateReservationRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type CreateReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *CreateReservationResponse) Reset() {
	*x = CreateReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservationpb_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationResponse) ProtoMessage() {}

func (x *CreateReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservationpb_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationResponse.ProtoReflect.Descriptor instead.
func (*CreateReservationResponse) Descriptor() ([]byte, []int) {
	return file_reservationpb_reservation_proto_rawDescGZIP(), []int{4}
}

func (x *CreateReservationResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

var File_reservationpb_reservation_proto protoreflect.FileDescriptor

var file_reservationpb_reservation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xb8, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x6f, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa5, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x59, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2a, 0x51, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x49,
	0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x4f,
	0x55, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x32, 0xdf, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25,
	0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6c, 0x6d, 0x61, 0x72, 0x2f, 0x6f, 0x74, 0x65, 0x6c, 0x2d,
	0x74, 0x72, 0x69, 0x76, 0x61, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservationpb_reservation_proto_rawDescOnce sync.Once
	file_reservationpb_reservation_proto_rawDescData = file_reservationpb_reservation_proto_rawDesc
)

func file_reservationpb_reservation_proto_rawDescGZIP() []byte {
	file_reservationpb_reservation_proto_rawDescOnce.Do(func() {
		file_reservationpb_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservationpb_reservation_proto_rawDescData)
	})
	return file_reservationpb_reservation_proto_rawDescData
}

var file_reservationpb_reservation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_reservationpb_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reservationpb_reservation_proto_goTypes = []interface{}{
	(ReservationStatus)(0),            // 0: reservationpb.ReservationStatus
	(*Reservation)(nil),               // 1: reservationpb.Reservation
	(*ListReservationRequest)(nil),    // 2: reservationpb.ListReservationRequest
	(*ListReservationResponse)(nil),   // 3: reservationpb.ListReservationResponse
	(*CreateReservationRequest)(nil),  // 4: reservationpb.CreateReservationRequest
	(*CreateReservationResponse)(nil), // 5: reservationpb.CreateReservationResponse
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_reservationpb_reservation_proto_depIdxs = []int32{
	0,  // 0: reservationpb.Reservation.status:type_name -> reservationpb.ReservationStatus
	6,  // 1: reservationpb.Reservation.start_date:type_name -> google.protobuf.Timestamp
	6,  // 2: reservationpb.Reservation.end_date:type_name -> google.protobuf.Timestamp
	6,  // 3: reservationpb.Reservation.created_at:type_name -> google.protobuf.Timestamp
	6,  // 4: reservationpb.ListReservationRequest.start_date:type_name -> google.protobuf.Timestamp
	6,  // 5: reservationpb.ListReservationRequest.end_date:type_name -> google.protobuf.Timestamp
	1,  // 6: reservationpb.ListReservationResponse.reservations:type_name -> reservationpb.Reservation
	6,  // 7: reservationpb.CreateReservationRequest.start_date:type_name -> google.protobuf.Timestamp
	6,  // 8: reservationpb.CreateReservationRequest.end_date:type_name -> google.protobuf.Timestamp
	1,  // 9: reservationpb.CreateReservationResponse.reservation:type_name -> reservationpb.Reservation
	2,  // 10: reservationpb.ReservationService.ListReservations:input_type -> reservationpb.ListReservationRequest
	4,  // 11: reservationpb.ReservationService.CreateReservation:input_type -> reservationpb.CreateReservationRequest
	3,  // 12: reservationpb.ReservationService.ListReservations:output_type -> reservationpb.ListReservationResponse
	5,  // 13: reservationpb.ReservationService.CreateReservation:output_type -> reservationpb.CreateReservationResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_reservationpb_reservation_proto_init() }
func file_reservationpb_reservation_proto_init() {
	if File_reservationpb_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservationpb_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_reservationpb_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReservationRequest); i {
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
		file_reservationpb_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReservationResponse); i {
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
		file_reservationpb_reservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationRequest); i {
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
		file_reservationpb_reservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationResponse); i {
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
			RawDescriptor: file_reservationpb_reservation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservationpb_reservation_proto_goTypes,
		DependencyIndexes: file_reservationpb_reservation_proto_depIdxs,
		EnumInfos:         file_reservationpb_reservation_proto_enumTypes,
		MessageInfos:      file_reservationpb_reservation_proto_msgTypes,
	}.Build()
	File_reservationpb_reservation_proto = out.File
	file_reservationpb_reservation_proto_rawDesc = nil
	file_reservationpb_reservation_proto_goTypes = nil
	file_reservationpb_reservation_proto_depIdxs = nil
}
