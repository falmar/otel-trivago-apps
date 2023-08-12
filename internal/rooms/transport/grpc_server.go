package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

var _ roompb.RoomServiceServer = (*grpcTransport)(nil)

type grpcTransport struct {
	listRooms kitgrpc.Handler

	roompb.UnimplementedRoomServiceServer
}

func NewGRPCServer(endpoints *endpoint.Endpoints) roompb.RoomServiceServer {
	return &grpcTransport{
		listRooms: kitgrpc.NewServer(
			endpoints.ListEndpoint,
			decodeListRoomsRequest,
			encodeListRoomsResponse,
		),
	}
}

func (g *grpcTransport) ListRooms(ctx context.Context, request *roompb.ListRoomsRequest) (*roompb.ListRoomsResponse, error) {
	ctx, resp, err := g.listRooms.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*roompb.ListRoomsResponse), nil
}

func (g *grpcTransport) mustEmbedUnimplementedRoomServiceServer() {}

func decodeListRoomsRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*roompb.ListRoomsRequest)
	req := &endpoint.ListRoomsRequest{
		Capacity: reqpb.Capacity,
		Limit:    reqpb.Limit,
		Offset:   reqpb.Offset,
	}

	return req, nil
}

func encodeListRoomsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.ListRoomsResponse)

	var rooms []*roompb.Room

	for _, r := range resp.Rooms {
		rooms = append(rooms, &roompb.Room{
			Id:          r.ID.String(),
			Capacity:    r.Capacity,
			CleanStatus: roompb.CleanStatus(r.CleanStatus),
		})
	}

	return &roompb.ListRoomsResponse{
		Rooms: rooms,
	}, nil
}
