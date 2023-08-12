package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stay/types"
	"github.com/google/uuid"
)

type Repository interface {
	List(ctx context.Context, options *ListOptions) ([]*types.Stay, error)
	GetById(ctx context.Context, id uuid.UUID) (*types.Stay, error)
	Create(ctx context.Context, stay *types.Stay) error
	Update(ctx context.Context, stay *types.Stay) error
}

type ListOptions struct {
	RoomID        uuid.UUID
	ReservationID uuid.UUID

	Limit  int64
	Offset int64
}
