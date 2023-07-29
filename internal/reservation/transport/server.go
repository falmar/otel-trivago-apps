package transport

import (
	"context"
	"errors"
	"github.com/falmar/otel-trivago/internal/reservation/endpoint"
	"github.com/falmar/otel-trivago/internal/reservation/service"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ reservationpb.ReservationServiceServer = (*grpcTransport)(nil)

type grpcTransport struct {
	create         kitgrpc.Handler
	get            kitgrpc.Handler
	update         kitgrpc.Handler
	delete         kitgrpc.Handler
	list           kitgrpc.Handler
	availableRooms kitgrpc.Handler

	reservationpb.UnimplementedReservationServiceServer
}

func NewGRPCServer(endpoints *endpoint.Endpoints) reservationpb.ReservationServiceServer {
	return &grpcTransport{
		create:         kitgrpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse),
		list:           kitgrpc.NewServer(endpoints.ListEndpoint, decodeListRequest, encodeListResponse),
		availableRooms: kitgrpc.NewServer(endpoints.ListAvailableRoomsEndpoint, decodeListAvailableRoomsRequest, encodeListAvailableRoomsResponse),
	}
}

func (g *grpcTransport) CreateReservation(ctx context.Context, reservation *reservationpb.Reservation) (*reservationpb.ReservationResponse, error) {
	_, resp, err := g.create.ServeGRPC(ctx, reservation)
	if err != nil {
		return nil, encodeError(err)
	}

	return resp.(*reservationpb.ReservationResponse), nil
}

func (g *grpcTransport) ListReservations(ctx context.Context, request *reservationpb.ReservationListRequest) (*reservationpb.ReservationListResponse, error) {
	_, resp, err := g.list.ServeGRPC(ctx, request)
	if err != nil {
		return nil, encodeError(err)
	}

	return resp.(*reservationpb.ReservationListResponse), nil
}

func (g *grpcTransport) ListAvailableRooms(ctx context.Context, request *reservationpb.RoomAvailabilityRequest) (*reservationpb.RoomAvailabilityResponse, error) {
	_, resp, err := g.availableRooms.ServeGRPC(ctx, request)
	if err != nil {
		return nil, encodeError(err)
	}

	return resp.(*reservationpb.RoomAvailabilityResponse), nil
}

func (g *grpcTransport) mustEmbedUnimplementedReservationServiceServer() {}

func encodeError(err error) error {
	if err == nil {
		return nil
	}

	var eInvalidArgument *endpoint.ErrInvalidArgument
	if errors.As(err, &eInvalidArgument) {
		return status.Error(codes.InvalidArgument, eInvalidArgument.Error())
	}
	var eReserved *service.ErrRoomReserved
	if errors.As(err, &eReserved) {
		return status.Error(codes.AlreadyExists, eReserved.Error())
	}

	return status.Error(codes.Unknown, err.Error())
}
