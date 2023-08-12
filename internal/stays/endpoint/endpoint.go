package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stays/service"
	"github.com/falmar/otel-trivago/internal/stays/types"
	"github.com/falmar/otel-trivago/pkg/pkg/kithelper"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"time"
)

type Endpoints struct {
	ListStaysEndpoint  kitendpoint.Endpoint
	CreateStayEndpoint kitendpoint.Endpoint
	UpdateStayEndpoint kitendpoint.Endpoint
}

func New(svc service.Service) Endpoints {
	return Endpoints{
		ListStaysEndpoint:  MakeListStaysEndpoint(svc),
		CreateStayEndpoint: MakeCreateStayEndpoint(svc),
		UpdateStayEndpoint: MakeUpdateStayEndpoint(svc),
	}
}

type ListStaysRequest struct {
	RoomID        string
	ReservationID string

	Limit  int64
	Offset int64
}

type ListStaysResponse struct {
	Stays []*types.Stay
	Total int64
}

func MakeListStaysEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error = nil
		req := request.(*ListStaysRequest)
		in := &service.ListStaysInput{
			Limit:  req.Limit,
			Offset: req.Offset,
		}

		if in.Limit == 0 {
			in.Limit = 10
		}

		if req.RoomID != "" {
			in.RoomID, err = uuid.Parse(req.RoomID)
			if err != nil {
				return nil, err
			}
		}
		if req.ReservationID != "" {
			in.ReservationID, err = uuid.Parse(req.ReservationID)
			if err != nil {
				return nil, err
			}
		}

		out, err := svc.ListStays(ctx, in)
		if err != nil {
			return nil, err
		}

		return &ListStaysResponse{
			Stays: out.Stays,
			Total: out.Total,
		}, nil
	}
}

type CreateStayRequest struct {
	RoomID        string
	ReservationID string
	CheckIn       time.Time
	Notes         string
}
type CreateStayResponse struct {
	Stay *types.Stay
}

func MakeCreateStayEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error = nil
		req := request.(*CreateStayRequest)
		in := &service.CreateStayInput{
			CheckIn: req.CheckIn,
			Notes:   req.Notes,
		}

		if req.RoomID == "" {
			return nil, &kithelper.ErrInvalidArgument{Message: "room id must be a valid uuid"}
		} else if req.RoomID != "" {
			in.RoomID, err = uuid.Parse(req.RoomID)
			if err != nil {
				return nil, err
			}
		}

		if req.ReservationID == "" {
			return nil, &kithelper.ErrInvalidArgument{Message: "reservation id must be a valid uuid"}
		} else if req.ReservationID != "" {
			in.ReservationID, err = uuid.Parse(req.ReservationID)
			if err != nil {
				return nil, err
			}
		}

		if in.CheckIn.IsZero() {
			return nil, &kithelper.ErrInvalidArgument{Message: "check in date must be a valid date"}
		}

		out, err := svc.CreateStay(ctx, in)
		if err != nil {
			return nil, err
		}

		return &CreateStayResponse{
			Stay: out.Stay,
		}, nil
	}
}

type UpdateStayRequest struct {
	ID       string
	RoomID   string
	CheckOut time.Time
	Notes    string
}

func MakeUpdateStayEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error = nil
		req := request.(*UpdateStayRequest)
		in := &service.UpdateStayInput{
			Notes: req.Notes,
		}

		if req.ID == "" {
			return nil, &kithelper.ErrInvalidArgument{Message: "id must be a valid uuid"}
		} else if req.ID != "" {
			in.ID, err = uuid.Parse(req.ID)
			if err != nil {
				return nil, err
			}
		}

		if req.RoomID != "" {
			in.RoomID, err = uuid.Parse(req.RoomID)
			if err != nil {
				return nil, err
			}
		}

		if !req.CheckOut.IsZero() {
			in.CheckOut = req.CheckOut
		}

		out, err := svc.UpdateStay(ctx, in)
		if err != nil {
			return nil, err
		}

		return &CreateStayResponse{
			Stay: out.Stay,
		}, nil
	}
}
