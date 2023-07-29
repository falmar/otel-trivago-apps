package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	"github.com/google/uuid"
	"time"
)

var _ Service = (*service)(nil)

type Service interface {
	List(ctx context.Context, input *ListInput) (*ListOutput, error)
	Create(ctx context.Context, input *CreateInput) (*CreateOutput, error)
	ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
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
	return &ListOutput{}, nil
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
	return &CreateOutput{}, nil
}

type ListAvailableRoomsInput struct {
	Start time.Time
	End   time.Time
}

type ListAvailableRoomsOutput struct {
	Rooms []*types.Room
}

func (s *service) ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error) {
	return &ListAvailableRoomsOutput{}, nil
}
