package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/google/uuid"
	"time"
)

var _ Service = (*service)(nil)

type Service interface {
	ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error)
	CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error)
}

type service struct {
	resvRepo reservationrepo.Repository
}

type Config struct {
	ResvRepo reservationrepo.Repository
}

func NewService(cfg *Config) Service {
	return &service{
		resvRepo: cfg.ResvRepo,
	}
}

type ListReservationsInput struct {
	Start time.Time
	End   time.Time

	Limit  int64
	Offset int64
}

type ListReservationsOutput struct {
	Reservations []*types.Reservation
	Total        int64
}

func (s *service) ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error) {
	resv, err := s.resvRepo.List(ctx, input.Start, input.End)
	if err != nil {
		return nil, err
	}

	return &ListReservationsOutput{
		Reservations: resv,
		Total:        int64(len(resv)),
	}, nil
}

type CreateReservationInput struct {
	RoomID uuid.UUID

	Start time.Time
	End   time.Time
}

type CreateReservationOutput struct {
	Reservation *types.Reservation
}

func (s *service) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	current, err := s.resvRepo.ByRoomID(ctx, input.RoomID)
	if err != nil {
		return nil, err
	}

	// if overlap, return error
	for _, resv := range current {
		if input.Start.After(resv.Start) && input.Start.Before(resv.End) {
			return nil, &types.ErrReservedRoom{RoomID: resv.RoomID.String()}
		} else if input.End.After(resv.Start) && input.End.Before(resv.End) {
			return nil, &types.ErrReservedRoom{RoomID: resv.RoomID.String()}
		} else if input.Start.Equal(resv.Start) || input.End.Equal(resv.End) {
			return nil, &types.ErrReservedRoom{RoomID: resv.RoomID.String()}
		}
	}

	resv := &types.Reservation{
		RoomID: input.RoomID,
		Start:  input.Start,
		End:    input.End,
	}

	err = s.resvRepo.Create(ctx, resv)
	if err != nil {
		return nil, err
	}

	return &CreateReservationOutput{
		Reservation: resv,
	}, nil
}
