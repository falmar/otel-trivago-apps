package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/service"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"time"
)

type ListAvailableRoomsRequest struct {
	Start time.Time
	End   time.Time

	Capacity int64
}

type ListAvailableRoomsResponse struct {
	Rooms []*types.Room
	Total int64
}

func makeListAvailableRoomsEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*ListAvailableRoomsRequest)

		out, err := svc.ListAvailableRooms(ctx, &service.ListAvailableRoomsInput{
			Start:    req.Start,
			End:      req.End,
			Capacity: req.Capacity,
		})
		if err != nil {
			return nil, err
		}

		return &ListAvailableRoomsResponse{
			Rooms: out.Rooms,
			Total: out.Total,
		}, nil
	}
}
