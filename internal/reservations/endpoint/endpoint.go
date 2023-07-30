package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/pkg/kithelper"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"time"
)

var _ service.Service = (*Endpoints)(nil)

type Endpoints struct {
	ListEndpoint   kitendpoint.Endpoint
	CreateEndpoint kitendpoint.Endpoint
}

func New(tr trace.Tracer, svc service.Service) *Endpoints {
	return &Endpoints{
		ListEndpoint: MakeTracerEndpointMiddleware(
			"reservations.endpoint.ListRooms", tr,
			makeListEndpoint(svc),
		),

		CreateEndpoint: MakeTracerEndpointMiddleware(
			"reservations.endpoint.Create", tr,
			makeCreateEndpoint(svc),
		),
	}
}

func (e *Endpoints) List(ctx context.Context, input *service.ListInput) (*service.ListOutput, error) {
	response, err := e.ListEndpoint(ctx, &ListRequest{
		Start:  input.Start,
		End:    input.End,
		Offset: input.Offset,
		Limit:  input.Limit,
	})
	if err != nil {
		return nil, err
	}

	resp := response.(*ListResponse)

	return &service.ListOutput{
		Reservations: resp.Reservations,
		Total:        resp.Total,
	}, nil
}

func (e *Endpoints) Create(ctx context.Context, input *service.CreateInput) (*service.CreateOutput, error) {
	response, err := e.CreateEndpoint(ctx, &CreateRequest{
		RoomID: input.RoomID.String(),

		Start: input.Start,
		End:   input.End,
	})
	if err != nil {
		return nil, err
	}

	resp := response.(*CreateResponse)

	return &service.CreateOutput{
		Reservation: resp.Reservation,
	}, nil
}

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
			return "", &kithelper.ErrInvalidArgument{Message: "room id must be a valid uuid"}
		}
		if req.Start.IsZero() {
			return "", &kithelper.ErrInvalidArgument{Message: "start date must be a valid date"}
		}
		if req.End.IsZero() {
			return "", &kithelper.ErrInvalidArgument{Message: "end date must be a valid date"}
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
