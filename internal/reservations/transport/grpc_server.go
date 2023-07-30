package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/pkg/kithelper"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ reservationpb.ReservationServiceServer = (*grpcTransport)(nil)

type grpcTransport struct {
	create kitgrpc.Handler
	list   kitgrpc.Handler

	reservationpb.UnimplementedReservationServiceServer
}

func NewGRPCServer(tr trace.Tracer, endpoints *endpoint.Endpoints) reservationpb.ReservationServiceServer {
	return &grpcTransport{
		create: kitgrpc.NewServer(
			endpoints.CreateEndpoint,
			decodeCreateRequest,
			encodeCreateResponse,
			kitgrpc.ServerBefore(kithelper.SpanTraceBefore(tr, "reservations.grpc.CreateReservation")),
			kitgrpc.ServerAfter(kithelper.SpanTraceAfter),
		),
		list: kitgrpc.NewServer(
			endpoints.ListEndpoint,
			decodeListRequest,
			encodeListResponse,
			kitgrpc.ServerBefore(kithelper.SpanTraceBefore(tr, "reservations.grpc.ListReservations")),
			kitgrpc.ServerAfter(kithelper.SpanTraceAfter),
		),
	}
}

func (g *grpcTransport) CreateReservation(ctx context.Context, reservation *reservationpb.Reservation) (*reservationpb.ReservationResponse, error) {
	ctx, resp, err := g.create.ServeGRPC(ctx, reservation)
	if err != nil {
		return nil, kithelper.EncodeError(ctx, err)
	}

	return resp.(*reservationpb.ReservationResponse), nil
}

func (g *grpcTransport) ListReservations(ctx context.Context, request *reservationpb.ReservationListRequest) (*reservationpb.ReservationListResponse, error) {
	ctx, resp, err := g.list.ServeGRPC(ctx, request)
	if err != nil {
		return nil, kithelper.EncodeError(ctx, err)
	}

	return resp.(*reservationpb.ReservationListResponse), nil
}

func (g *grpcTransport) mustEmbedUnimplementedReservationServiceServer() {}

func decodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.ReservationListRequest)

	req := &endpoint.ListRequest{}

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
		resv = append(resv, &reservationpb.Reservation{
			Id:     r.ID.String(),
			RoomId: r.RoomID.String(),
			Status: reservationpb.ReservationStatus(r.Status),
			StartDate: &timestamppb.Timestamp{
				Seconds: r.Start.Unix(),
			},
			EndDate: &timestamppb.Timestamp{
				Seconds: r.End.Unix(),
			},
			CreatedAt: &timestamppb.Timestamp{
				Seconds: r.CreatedAt.Unix(),
			},
		})
	}

	return &reservationpb.ReservationListResponse{
		Reservations: resv,
		Total:        resp.Total,
	}, nil
}

func decodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.Reservation)
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
	resv := response.(*endpoint.CreateResponse).Reservation

	return &reservationpb.ReservationResponse{
		Reservation: &reservationpb.Reservation{
			Id:     resv.ID.String(),
			RoomId: resv.RoomID.String(),
			Status: reservationpb.ReservationStatus(resv.Status),
			StartDate: &timestamppb.Timestamp{
				Seconds: resv.Start.Unix(),
			},
			EndDate: &timestamppb.Timestamp{
				Seconds: resv.End.Unix(),
			},
			CreatedAt: &timestamppb.Timestamp{
				Seconds: resv.CreatedAt.Unix(),
			},
		},
	}, nil
}
