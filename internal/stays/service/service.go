package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stays/repo"
	"github.com/falmar/otel-trivago/internal/stays/types"
	"github.com/google/uuid"
	"time"
)

var _ Service = (*service)(nil)

type Config struct {
	Repo repo.Repository
}

func New(cfg *Config) Service {
	return &service{
		repo: cfg.Repo,
	}
}

type Service interface {
	ListStays(ctx context.Context, input *ListStaysInput) (*ListStaysOutput, error)
	CreateStay(ctx context.Context, input *CreateStayInput) (*CreateStayOutput, error)
	UpdateStay(ctx context.Context, input *UpdateStayInput) (*UpdateStayOutput, error)
}

type service struct {
	repo repo.Repository
}

type ListStaysInput struct {
	RoomID        uuid.UUID
	ReservationID uuid.UUID

	Limit  int64
	Offset int64
}
type ListStaysOutput struct {
	Stays []*types.Stay
	Total int64
}

func (s *service) ListStays(ctx context.Context, input *ListStaysInput) (*ListStaysOutput, error) {
	stays, err := s.repo.List(ctx, &repo.ListOptions{
		RoomID:        input.RoomID,
		ReservationID: input.ReservationID,

		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return nil, err
	}

	return &ListStaysOutput{
		Stays: stays,
		Total: int64(len(stays)),
	}, nil
}

type CreateStayInput struct {
	RoomID        uuid.UUID
	ReservationID uuid.UUID

	CheckIn time.Time
	Notes   string
}
type CreateStayOutput struct {
	Stay *types.Stay
}

func (s *service) CreateStay(ctx context.Context, input *CreateStayInput) (*CreateStayOutput, error) {
	stay := &types.Stay{
		RoomID:        input.RoomID,
		ReservationID: input.ReservationID,
		CheckIn:       input.CheckIn,
		Notes:         input.Notes,
	}

	err := s.repo.Create(ctx, stay)
	if err != nil {
		return nil, err
	}

	return &CreateStayOutput{
		Stay: stay,
	}, nil
}

type UpdateStayInput struct {
	ID     uuid.UUID
	RoomID uuid.UUID

	CheckOut time.Time
	Note     string
}
type UpdateStayOutput struct {
	Stay *types.Stay
}

func (s *service) UpdateStay(ctx context.Context, input *UpdateStayInput) (*UpdateStayOutput, error) {
	stay, err := s.repo.GetById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.RoomID != uuid.Nil {
		stay.RoomID = input.RoomID
	}
	if !input.CheckOut.IsZero() {
		stay.CheckOut = input.CheckOut
	}
	if input.Note != "" {
		stay.Notes = input.Note
	}

	err = s.repo.Update(ctx, stay)
	if err != nil {
		return nil, err
	}

	return &UpdateStayOutput{
		Stay: stay,
	}, nil
}
