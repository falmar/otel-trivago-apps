package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/falmar/otel-trivago/pkg/pkg/kithelper"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ reservationpb.ReservationServiceServer = (*grpcServer)(nil)

type grpcServer struct {
	create kitgrpc.Handler
	list   kitgrpc.Handler

	reservationpb.UnimplementedReservationServiceServer
}

func NewGRPCServer(endpoints *endpoint.Endpoints) reservationpb.ReservationServiceServer {
	return &grpcServer{
		create: kitgrpc.NewServer(
			endpoints.CreateEndpoint,
			decodeCreateRequest,
			encodeCreateResponse,
		),
		list: kitgrpc.NewServer(
			endpoints.ListEndpoint,
			decodeListRequest,
			encodeListResponse,
		),
	}
}

func (g *grpcServer) CreateReservation(ctx context.Context, reservation *reservationpb.CreateReservationRequest) (*reservationpb.CreateReservationResponse, error) {
	ctx, resp, err := g.create.ServeGRPC(ctx, reservation)
	if err != nil {
		return nil, kithelper.EncodeError(ctx, err)
	}

	return resp.(*reservationpb.CreateReservationResponse), nil
}

func (g *grpcServer) ListReservations(ctx context.Context, request *reservationpb.ListReservationRequest) (*reservationpb.ListReservationResponse, error) {
	ctx, resp, err := g.list.ServeGRPC(ctx, request)
	if err != nil {
		return nil, kithelper.EncodeError(ctx, err)
	}

	return resp.(*reservationpb.ListReservationResponse), nil
}

func (g *grpcServer) mustEmbedUnimplementedReservationServiceServer() {}

func decodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.ListReservationRequest)

	req := &endpoint.ListRequest{
		RoomID: pbreq.RoomId,
		Limit:  pbreq.Limit,
		Offset: pbreq.Offset,
	}

	if pbreq.StartDate != nil {
		req.Start = pbreq.StartDate.AsTime().UTC()
	}
	if pbreq.EndDate != nil {
		req.End = pbreq.EndDate.AsTime().UTC()
	}

	return req, nil
}

func encodeListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.ListResponse)
	var resv []*reservationpb.Reservation

	for _, r := range resp.Reservations {
		rpb := &reservationpb.Reservation{}
		mapReservation(r, rpb)

		resv = append(resv, rpb)
	}

	return &reservationpb.ListReservationResponse{
		Reservations: resv,
		Total:        resp.Total,
	}, nil
}

func decodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.CreateReservationRequest)
	req := &endpoint.CreateRequest{
		RoomID: pbreq.RoomId,
	}

	if pbreq.StartDate != nil {
		req.Start = pbreq.StartDate.AsTime().UTC()
	}
	if pbreq.EndDate != nil {
		req.End = pbreq.EndDate.AsTime().UTC()
	}

	return req, nil
}

func encodeCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	rpb := &reservationpb.Reservation{}
	mapReservation(response.(*endpoint.CreateResponse).Reservation, rpb)

	return &reservationpb.CreateReservationResponse{
		Reservation: rpb,
	}, nil
}

func mapReservation(r *types.Reservation, rpb *reservationpb.Reservation) {
	rpb.Id = r.ID.String()
	rpb.RoomId = r.RoomID.String()
	rpb.Status = reservationpb.ReservationStatus(r.Status)
	rpb.StartDate = timestamppb.New(r.Start)
	rpb.EndDate = timestamppb.New(r.End)
	rpb.CreatedAt = timestamppb.New(r.CreatedAt)
	rpb.UpdatedAt = timestamppb.New(r.UpdatedAt)
}
