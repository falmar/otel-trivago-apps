package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/internal/rooms/types"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	kitendpoint "github.com/go-kit/kit/endpoint"
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

	return &grpcClient{
		listRooms: listEndpoint,
	}
}

type grpcClient struct {
	listRooms kitendpoint.Endpoint
}

func (g *grpcClient) ListRooms(ctx context.Context, input *service.ListRoomsInput) (*service.ListRoomsOutput, error) {
	response, err := g.listRooms(ctx, &endpoint.ListRoomsRequest{
		Capacity: input.Capacity,
		Limit:    input.Limit,
		Offset:   input.Offset,
	})
	if err != nil {
		return nil, err
	}

	resp := response.(*endpoint.ListRoomsResponse)

	return &service.ListRoomsOutput{
		Rooms: resp.Rooms,
	}, nil
}

func decodeListRoomsResponse(_ context.Context, response interface{}) (request interface{}, err error) {
	respb := response.(*roompb.ListRoomsResponse)
	resp := &endpoint.ListRoomsResponse{
		Rooms: make([]*types.Room, respb.Total),
		Total: respb.Total,
	}

	var rooms []*types.Room

	for i, rpb := range respb.Rooms {
		r := &types.Room{}
		mapRoomPB(rpb, r)

		rooms[i] = r
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

func mapRoomPB(rpb *roompb.Room, r *types.Room) {
	r.ID = uuid.MustParse(rpb.Id)
	r.Capacity = rpb.Capacity
	r.CleanStatus = types.CleanStatus(rpb.CleanStatus)
}
