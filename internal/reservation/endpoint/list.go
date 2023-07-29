package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/service"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"time"
)

type ListRequest struct {
	Start time.Time
	End   time.Time

	Offset int64
	Limit  int64
}

type ListResponse struct {
	Reservations []*types.Reservation
	Total        int64
}

func makeListEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*ListRequest)

		out, err := svc.List(ctx, &service.ListInput{
			Start: req.Start,
			End:   req.End,

			Offset: req.Offset,
			Limit:  req.Limit,
		})
		if err != nil {
			return nil, err
		}

		return &ListResponse{
			Reservations: out.Reservations,
			Total:        out.Total,
		}, nil
	}
}
