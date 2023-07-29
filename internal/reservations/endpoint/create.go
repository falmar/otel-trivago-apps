package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"time"
)

type CreateRequest struct {
	RoomID string

	Start time.Time
	End   time.Time
}

type CreateResponse struct {
	Reservation *types.Reservation
}

func makeCreateEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*CreateRequest)

		id, err := uuid.Parse(req.RoomID)
		if err != nil {
			return "", &ErrInvalidArgument{Message: "room id must be a valid uuid"}
		}
		if req.Start.IsZero() {
			return "", &ErrInvalidArgument{Message: "start date must be a valid date"}
		}
		if req.End.IsZero() {
			return "", &ErrInvalidArgument{Message: "end date must be a valid date"}
		}

		out, err := svc.Create(ctx, &service.CreateInput{
			RoomID: id,

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
