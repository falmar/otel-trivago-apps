package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/internal/rooms/types"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

var _ roompb.RoomServiceServer = (*grpcServer)(nil)

type grpcServer struct {
	listRooms kitgrpc.Handler

	roompb.UnimplementedRoomServiceServer
}

func NewGRPCServer(endpoints *endpoint.Endpoints) roompb.RoomServiceServer {
	return &grpcServer{
		listRooms: kitgrpc.NewServer(
			endpoints.ListEndpoint,
			decodeListRoomsRequest,
			encodeListRoomsResponse,
		),
	}
}

func (g *grpcServer) ListRooms(ctx context.Context, request *roompb.ListRoomsRequest) (*roompb.ListRoomsResponse, error) {
	ctx, resp, err := g.listRooms.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*roompb.ListRoomsResponse), nil
}

func (g *grpcServer) mustEmbedUnimplementedRoomServiceServer() {}

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
	respb := &roompb.ListRoomsResponse{
		Rooms: make([]*roompb.Room, resp.Total),
		Total: resp.Total,
	}

	for i, r := range resp.Rooms {
		rpb := &roompb.Room{}
		mapRoom(r, rpb)

		respb.Rooms[i] = rpb
	}

	return respb, nil
}

func mapRoom(r *types.Room, rpb *roompb.Room) {
	rpb.Id = r.ID.String()
	rpb.Capacity = r.Capacity
	rpb.CleanStatus = roompb.CleanStatus(r.CleanStatus)
}
