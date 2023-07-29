package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/internal/rooms/types"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	grpc "google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) service.Service {
	listEndpoint := kitgrpc.NewClient(
		conn,
		"roompb.RoomService",
		"ListRooms",
		encodeListRoomsRequest,
		decodeListRoomsResponse,
		&roompb.ListRoomsResponse{},
	).Endpoint()

	return &endpoint.Endpoints{
		ListEndpoint: listEndpoint,
	}
}

func decodeListRoomsResponse(_ context.Context, response interface{}) (request interface{}, err error) {
	respb := response.(*roompb.ListRoomsResponse)

	var rooms []*types.Room

	for _, r := range respb.Rooms {
		rooms = append(rooms, &types.Room{
			ID:          uuid.MustParse(r.Id),
			Capacity:    r.Capacity,
			CleanStatus: types.CleanStatus(r.CleanStatus),
		})
	}

	resp := &endpoint.ListRoomsResponse{
		Rooms: rooms,
	}

	return resp, nil
}

func encodeListRoomsRequest(_ context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*endpoint.ListRoomsRequest)

	reqpb := &roompb.ListRoomsRequest{
		Capacity: req.Capacity,
		Limit:    req.Limit,
		Offset:   req.Offset,
	}

	return reqpb, nil
}
