package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/service"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"time"
)

type CreateRequest struct {
	RoomID uuid.UUID

	Start time.Time
	End   time.Time
}

type CreateResponse struct {
	Reservation *types.Reservation
}

func makeCreateEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)

		out, err := svc.Create(ctx, &service.CreateInput{
			RoomID: req.RoomID,

			Start: req.Start,
			End:   req.End,
		})
		if err != nil {
			return nil, err
		}

		return &CreateResponse{
			Reservation: out.Reservation,
		}, nil
	}
}
