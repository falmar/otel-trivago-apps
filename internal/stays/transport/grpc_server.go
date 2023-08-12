package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stays/endpoint"
	"github.com/falmar/otel-trivago/internal/stays/types"
	"github.com/falmar/otel-trivago/pkg/proto/v1/staypb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ staypb.StayServiceServer = (*grpcServer)(nil)

func NewGRPCServer(endpoints endpoint.Endpoints) staypb.StayServiceServer {
	return &grpcServer{
		list: kitgrpc.NewServer(
			endpoints.ListStaysEndpoint,
			decodeListRequest,
			encodeListResponse,
		),
		create: kitgrpc.NewServer(
			endpoints.CreateStayEndpoint,
			decodeCreateRequest,
			encodeCreateResponse,
		),
		update: kitgrpc.NewServer(
			endpoints.UpdateStayEndpoint,
			decodeUpdateRequest,
			encodeUpdateResponse,
		),
	}
}

type grpcServer struct {
	list   kitgrpc.Handler
	create kitgrpc.Handler
	update kitgrpc.Handler

	staypb.StayServiceServer
}

func (g *grpcServer) ListStays(ctx context.Context, request *staypb.ListStaysRequest) (*staypb.ListStaysResponse, error) {
	ctx, resp, err := g.list.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*staypb.ListStaysResponse), nil
}

func (g *grpcServer) CreateStay(ctx context.Context, request *staypb.CreateStayRequest) (*staypb.CreateStayResponse, error) {
	ctx, resp, err := g.create.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*staypb.CreateStayResponse), nil
}

func (g *grpcServer) UpdateStay(ctx context.Context, request *staypb.UpdateStayRequest) (*staypb.UpdateStayResponse, error) {
	ctx, resp, err := g.update.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*staypb.UpdateStayResponse), nil
}

func (g *grpcServer) mustEmbedUnimplementedStayServiceServer() {}

func decodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*staypb.ListStaysRequest)
	req := &endpoint.ListStaysRequest{
		RoomID:        pbreq.RoomId,
		ReservationID: pbreq.ReservationId,
		Limit:         pbreq.Limit,
		Offset:        pbreq.Offset,
	}

	return req, nil
}

func encodeListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.ListStaysResponse)
	pbresp := &staypb.ListStaysResponse{
		Stays: make([]*staypb.Stay, len(resp.Stays)),
		Total: resp.Total,
	}

	for i, s := range resp.Stays {
		pbs := &staypb.Stay{}
		mapStay(s, pbs)

		pbresp.Stays[i] = pbs
	}

	return pbresp, nil
}

func decodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	respb := request.(*staypb.CreateStayRequest)
	req := &endpoint.CreateStayRequest{
		RoomID:        respb.RoomId,
		ReservationID: respb.ReservationId,
		Notes:         respb.Note,
	}

	if respb.CheckIn != nil {
		req.CheckIn = respb.CheckIn.AsTime()
	}

	return req, nil
}

func encodeCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.CreateStayResponse)
	pbresp := &staypb.CreateStayResponse{
		Stay: &staypb.Stay{},
	}

	mapStay(resp.Stay, pbresp.Stay)

	return pbresp, nil
}

func decodeUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	respb := request.(*staypb.UpdateStayRequest)
	req := &endpoint.UpdateStayRequest{
		ID:     respb.Id,
		RoomID: respb.RoomId,
		Notes:  respb.Note,
	}

	if respb.CheckOut != nil {
		req.CheckOut = respb.CheckOut.AsTime()
	}

	return req, nil
}

func encodeUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.UpdateStayResponse)
	pbresp := &staypb.UpdateStayResponse{
		Stay: &staypb.Stay{},
	}

	mapStay(resp.Stay, pbresp.Stay)

	return pbresp, nil
}

func mapStay(s *types.Stay, rs *staypb.Stay) {
	rs.Id = s.ID.String()
	rs.RoomId = s.RoomID.String()
	rs.ReservationId = s.ReservationID.String()

	if !s.CheckIn.IsZero() {
		rs.CheckIn = timestamppb.New(s.CheckIn)
	}
	if !s.CheckOut.IsZero() {
		rs.CheckOut = timestamppb.New(s.CheckOut)
	}

	rs.Note = s.Note

	rs.CreatedAt = timestamppb.New(s.CreatedAt)
	rs.UpdatedAt = timestamppb.New(s.UpdatedAt)
}
