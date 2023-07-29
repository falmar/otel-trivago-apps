package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
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

func decodeListRoomsResponse(context.Context, interface{}) (request interface{}, err error) {
	return nil, nil
}

func encodeListRoomsRequest(context.Context, interface{}) (response interface{}, err error) {
	return nil, nil
}
