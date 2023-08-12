package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/internal/rooms/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListEndpoint kitendpoint.Endpoint
}

func New(svc service.Service) *Endpoints {
	return &Endpoints{
		ListEndpoint: MakeListEndpoint(svc),
	}
}

type ListRoomsRequest struct {
	Capacity int64
	Limit    int64
	Offset   int64
}

type ListRoomsResponse struct {
	Rooms []*types.Room
}

func MakeListEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*ListRoomsRequest)

		output, err := svc.ListRooms(ctx, &service.ListRoomsInput{
			Capacity: req.Capacity,
			Limit:    req.Limit,
			Offset:   req.Offset,
		})
		if err != nil {
			return nil, err
		}

		return &ListRoomsResponse{
			Rooms: output.Rooms,
		}, nil
	}
}
