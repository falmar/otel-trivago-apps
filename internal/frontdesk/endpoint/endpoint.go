package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/frontdesk/service"
	roomtypes "github.com/falmar/otel-trivago/internal/rooms/types"
	"github.com/falmar/otel-trivago/pkg/pkg/kithelper"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"time"
)

func New(svc service.Service) Endpoints {
	return Endpoints{
		CheckAvailabilityEndpoint: MakeCheckAvailabilityEndpoint(svc),
	}
}

type Endpoints struct {
	CheckAvailabilityEndpoint kitendpoint.Endpoint
}

type CheckAvailabilityRequest struct {
	Capacity int64
	Start    time.Time
	End      time.Time
}
type CheckAvailabilityResponse struct {
	Rooms []*roomtypes.Room
}

func MakeCheckAvailabilityEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		req := request.(*CheckAvailabilityRequest)

		if req.Start.IsZero() {
			return nil, &kithelper.ErrInvalidArgument{Message: "start date is required"}
		} else if req.End.IsZero() {
			return nil, &kithelper.ErrInvalidArgument{Message: "end date is required"}
		} else if req.Start.After(req.End) {
			return nil, &kithelper.ErrInvalidArgument{Message: "start date must be before end date"}
		}

		out, err := svc.CheckAvailability(ctx, &service.CheckAvailabilityInput{
			Capacity: req.Capacity,
			Start:    req.Start,
			End:      req.End,
		})
		if err != nil {
			return nil, err
		}

		return &CheckAvailabilityResponse{
			Rooms: out.Rooms,
		}, nil
	}
}
