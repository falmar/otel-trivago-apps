package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/falmar/otel-trivago/pkg/pkg/kithelper"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"time"
)

type Endpoints struct {
	ListEndpoint   kitendpoint.Endpoint
	CreateEndpoint kitendpoint.Endpoint
}

func New(svc service.Service) *Endpoints {
	return &Endpoints{
		ListEndpoint:   makeListReservationsEndpoint(svc),
		CreateEndpoint: makeCreateReservationEndpoint(svc),
	}
}

type ListRequest struct {
	RoomID string
	Offset int64
	Limit  int64

	Start time.Time
	End   time.Time
}
type ListResponse struct {
	Reservations []*types.Reservation
	Total        int64
}

func makeListReservationsEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error = nil
		req := request.(*ListRequest)
		in := &service.ListReservationsInput{
			Start: req.Start,
			End:   req.End,

			Limit:  req.Limit,
			Offset: req.Offset,
		}

		if in.Limit == 0 {
			in.Limit = 10
		}
		if in.Offset < 0 {
			in.Offset = 0
		}

		if req.RoomID != "" {
			in.RoomID, err = uuid.Parse(req.RoomID)
			if err != nil {
				return nil, &kithelper.ErrInvalidArgument{Message: "room id must be a valid uuid"}
			}
		}

		out, err := svc.ListReservations(ctx, in)
		if err != nil {
			return nil, err
		}

		return &ListResponse{
			Reservations: out.Reservations,
			Total:        out.Total,
		}, nil
	}
}

type CreateRequest struct {
	RoomID string

	Start time.Time
	End   time.Time
}
type CreateResponse struct {
	Reservation *types.Reservation
}

func makeCreateReservationEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*CreateRequest)

		id, err := uuid.Parse(req.RoomID)
		if err != nil {
			return "", &kithelper.ErrInvalidArgument{Message: "room id must be a valid uuid"}
		}
		if req.Start.IsZero() {
			return "", &kithelper.ErrInvalidArgument{Message: "start date must be a valid date"}
		}
		if req.End.IsZero() {
			return "", &kithelper.ErrInvalidArgument{Message: "end date must be a valid date"}
		}

		out, err := svc.CreateReservation(ctx, &service.CreateReservationInput{
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
