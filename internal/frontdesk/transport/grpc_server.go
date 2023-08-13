package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/frontdesk/endpoint"
	roomtransport "github.com/falmar/otel-trivago/internal/rooms/transport"
	"github.com/falmar/otel-trivago/pkg/pkg/kithelper"
	"github.com/falmar/otel-trivago/pkg/proto/v1/frontdeskpb"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

var _ frontdeskpb.FrontDeskServiceServer = (*grpcServer)(nil)

func NewGRPCServer(e endpoint.Endpoints) frontdeskpb.FrontDeskServiceServer {
	return &grpcServer{
		checkAvailabilityEndpoint: kitgrpc.NewServer(
			e.CheckAvailabilityEndpoint,
			decodeCheckAvailabilityRequest,
			encodeCheckAvailabilityResponse,
		),
	}
}

type grpcServer struct {
	checkAvailabilityEndpoint kitgrpc.Handler

	frontdeskpb.UnimplementedFrontDeskServiceServer
}

func (s *grpcServer) CheckAvailability(ctx context.Context, request *frontdeskpb.CheckAvailabilityRequest) (*frontdeskpb.CheckAvailabilityResponse, error) {
	_, resp, err := s.checkAvailabilityEndpoint.ServeGRPC(ctx, request)
	if err != nil {
		return nil, kithelper.EncodeError(ctx, err)
	}

	return resp.(*frontdeskpb.CheckAvailabilityResponse), nil
}

func decodeCheckAvailabilityRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*frontdeskpb.CheckAvailabilityRequest)
	req := &endpoint.CheckAvailabilityRequest{
		Capacity: pbreq.Capacity,
	}

	if pbreq.StartDate != nil {
		req.Start = pbreq.StartDate.AsTime()
	}
	if pbreq.EndDate != nil {
		req.End = pbreq.EndDate.AsTime()
	}

	return req, nil
}

func encodeCheckAvailabilityResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*endpoint.CheckAvailabilityResponse)
	pbresp := &frontdeskpb.CheckAvailabilityResponse{
		Rooms: make([]*roompb.Room, len(res.Rooms)),
	}

	for i, room := range res.Rooms {
		pbresp.Rooms[i] = &roompb.Room{}

		roomtransport.MapRoom(room, pbresp.Rooms[i])
	}

	return pbresp, nil
}
