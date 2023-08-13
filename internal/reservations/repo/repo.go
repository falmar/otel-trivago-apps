package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	List(ctx context.Context, opts *ListOptions) ([]*types.Reservation, error)
	ListByRoomID(ctx context.Context, roomID uuid.UUID) ([]*types.Reservation, error)
	Create(ctx context.Context, res *types.Reservation) error
}

type ListOptions struct {
	RoomID uuid.UUID
	Start  time.Time
	End    time.Time

	Limit  int64
	Offset int64
}
