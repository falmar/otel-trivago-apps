package service

import (
	"context"
	"errors"
	"github.com/falmar/otel-trivago/internal/reservations/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	roomsvc "github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/google/uuid"
	"time"
)

var _ Service = (*service)(nil)

type Service interface {
	List(ctx context.Context, input *ListInput) (*ListOutput, error)
	Create(ctx context.Context, input *CreateInput) (*CreateOutput, error)
}

type service struct {
	resvRepo reservationrepo.Repository
	roomSvc  roomsvc.Service
}

type Config struct {
	ResvRepo reservationrepo.Repository
	RoomSvc  roomsvc.Service
}

func NewService(cfg *Config) Service {
	return &service{
		resvRepo: cfg.ResvRepo,
		roomSvc:  cfg.RoomSvc,
	}
}

type ListInput struct {
	Start time.Time
	End   time.Time

	Limit  int64
	Offset int64
}

type ListOutput struct {
	Reservations []*types.Reservation
	Total        int64
}

func (s *service) List(ctx context.Context, input *ListInput) (*ListOutput, error) {
	resv, err := s.resvRepo.List(ctx, input.Start, input.End)
	if err != nil {
		return nil, err
	}

	return &ListOutput{
		Reservations: resv,
		Total:        int64(len(resv)),
	}, nil
}

type CreateInput struct {
	RoomID uuid.UUID

	Start time.Time
	End   time.Time
}

type CreateOutput struct {
	Reservation *types.Reservation
}

func (s *service) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	rooms, err := s.roomSvc.ListRooms(ctx, &roomsvc.ListRoomsInput{})
	if err != nil {
		return nil, err
	}

	// mimic checking if room exists
	var exists bool
	for _, room := range rooms.Rooms {
		if room.ID == input.RoomID {
			exists = true
			break
		}
	}
	if !exists {
		return nil, errors.New("room does not exist")
	}

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

	return &CreateOutput{
		Reservation: resv,
	}, nil
}
