package service

import (
	"context"
	"fmt"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	"github.com/google/uuid"
	"time"
)

type CreateInput struct {
	RoomID uuid.UUID

	Start time.Time
	End   time.Time
}

type CreateOutput struct {
	Reservation *types.Reservation
}

func (s *service) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	current, err := s.resvRepo.ByRoomID(ctx, input.RoomID)
	if err != nil {
		return nil, err
	}

	// if overlap, return error
	for _, resv := range current {
		if input.Start.After(resv.Start) && input.Start.Before(resv.End) {
			return nil, &ErrRoomReserved{RoomID: resv.RoomID.String()}
		} else if input.End.After(resv.Start) && input.End.Before(resv.End) {
			return nil, &ErrRoomReserved{RoomID: resv.RoomID.String()}
		} else if input.Start.Equal(resv.Start) || input.End.Equal(resv.End) {
			return nil, &ErrRoomReserved{RoomID: resv.RoomID.String()}
		}
	}

	fmt.Println(current)

	resv := &types.Reservation{
		RoomID: input.RoomID,
		Start:  input.Start,
		End:    input.End,
	}

	err = s.resvRepo.Create(ctx, resv)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Reservation: resv,
	}, nil
}
