package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"go.opentelemetry.io/otel/trace"
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
			kitgrpc.ServerBefore(spanBefore(tr, "reservation.grpc.Create")),
			kitgrpc.ServerAfter(spanAfter),
		),
		list: kitgrpc.NewServer(
			endpoints.ListEndpoint,
			decodeListRequest,
			encodeListResponse,
			kitgrpc.ServerBefore(spanBefore(tr, "reservation.grpc.List")),
			kitgrpc.ServerAfter(spanAfter),
		),
	}
}

func (g *grpcTransport) CreateReservation(ctx context.Context, reservation *reservationpb.Reservation) (*reservationpb.ReservationResponse, error) {
	ctx, resp, err := g.create.ServeGRPC(ctx, reservation)
	if err != nil {
		return nil, encodeError(ctx, err)
	}

	return resp.(*reservationpb.ReservationResponse), nil
}

func (g *grpcTransport) ListReservations(ctx context.Context, request *reservationpb.ReservationListRequest) (*reservationpb.ReservationListResponse, error) {
	ctx, resp, err := g.list.ServeGRPC(ctx, request)
	if err != nil {
		return nil, encodeError(ctx, err)
	}

	return resp.(*reservationpb.ReservationListResponse), nil
}

func (g *grpcTransport) mustEmbedUnimplementedReservationServiceServer() {}
