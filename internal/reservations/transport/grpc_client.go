package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ service.Service = (*grpcClient)(nil)

func NewGRPCClient(conn *grpc.ClientConn) service.Service {
	listEndpoint := kitgrpc.NewClient(
		conn,
		"reservationpb.ReservationService",
		"ListReservations",
		encodeListReservationsRequest,
		decodeListReservationsResponse,
		&reservationpb.ListReservationResponse{},
	).Endpoint()
	createEndpoint := kitgrpc.NewClient(
		conn,
		"reservationpb.ReservationService",
		"CreateReservation",
		encodeCreateReservationRequest,
		decodeCreateReservationResponse,
		&reservationpb.CreateReservationResponse{},
	).Endpoint()

	return &grpcClient{
		listReservations:  listEndpoint,
		createReservation: createEndpoint,
	}
}

type grpcClient struct {
	listReservations  kitendpoint.Endpoint
	createReservation kitendpoint.Endpoint
}

func (g *grpcClient) ListReservations(ctx context.Context, input *service.ListReservationsInput) (*service.ListReservationsOutput, error) {
	response, err := g.listReservations(ctx, &endpoint.ListRequest{
		RoomID: input.RoomID.String(),
		Start:  input.Start,
		End:    input.End,
		Offset: input.Offset,
		Limit:  input.Limit,
	})
	if err != nil {
		return nil, err
	}

	resp := response.(*endpoint.ListResponse)

	return &service.ListReservationsOutput{
		Reservations: resp.Reservations,
		Total:        resp.Total,
	}, nil
}

func (g *grpcClient) CreateReservation(ctx context.Context, input *service.CreateReservationInput) (*service.CreateReservationOutput, error) {
	response, err := g.createReservation(ctx, &endpoint.CreateRequest{
		RoomID: input.RoomID.String(),

		Start: input.Start,
		End:   input.End,
	})
	if err != nil {
		return nil, err
	}

	resp := response.(*endpoint.CreateResponse)

	return &service.CreateReservationOutput{
		Reservation: resp.Reservation,
	}, nil
}

func encodeListReservationsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*endpoint.ListRequest)

	return &reservationpb.ListReservationRequest{
		RoomId:    req.RoomID,
		StartDate: timestamppb.New(req.Start),
		EndDate:   timestamppb.New(req.End),

		Offset: req.Offset,
		Limit:  req.Limit,
	}, nil
}

func decodeListReservationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*reservationpb.ListReservationResponse)
	res := &endpoint.ListResponse{
		Total:        respb.Total,
		Reservations: make([]*types.Reservation, len(respb.Reservations)),
	}

	for i, rpb := range respb.Reservations {
		r := &types.Reservation{}
		mapReservationPB(rpb, r)

		res.Reservations[i] = r
	}

	return res, nil
}

func encodeCreateReservationRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*endpoint.CreateRequest)

	return &reservationpb.CreateReservationRequest{
		RoomId: req.RoomID,

		StartDate: timestamppb.New(req.Start),
		EndDate:   timestamppb.New(req.End),
	}, nil
}

func decodeCreateReservationResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*reservationpb.CreateReservationResponse)
	res := &endpoint.CreateResponse{
		Reservation: &types.Reservation{},
	}

	mapReservationPB(respb.Reservation, res.Reservation)

	return res, nil
}

func mapReservationPB(rpb *reservationpb.Reservation, r *types.Reservation) {
	r.ID = uuid.MustParse(rpb.Id)
	r.RoomID = uuid.MustParse(rpb.RoomId)
	r.Status = types.ReservationStatus(rpb.Status)
	r.Start = rpb.StartDate.AsTime()
	r.End = rpb.EndDate.AsTime()
	r.CreatedAt = rpb.CreatedAt.AsTime()
	r.UpdatedAt = rpb.UpdatedAt.AsTime()
}
