package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/internal/rooms/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"go.opentelemetry.io/otel/trace"
)

var _ service.Service = (*Endpoints)(nil)

type Endpoints struct {
	ListEndpoint kitendpoint.Endpoint
}

func New(svc service.Service, tr trace.Tracer) *Endpoints {
	return &Endpoints{
		ListEndpoint: MakeTracerEndpointMiddleware(
			"rooms.endpoint.ListRooms", tr,
			MakeListEndpoint(svc),
		),
	}
}

func (e *Endpoints) ListRooms(ctx context.Context, input *service.ListRoomsInput) (*service.ListRoomsOutput, error) {
	response, err := e.ListEndpoint(ctx, &ListRoomsRequest{
		Capacity: input.Capacity,
		Limit:    input.Limit,
		Offset:   input.Offset,
	})
	if err != nil {
		return nil, err
	}

	resp := response.(*ListRoomsResponse)

	return &service.ListRoomsOutput{
		Rooms: resp.Rooms,
	}, nil
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
