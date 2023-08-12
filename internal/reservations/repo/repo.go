package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	List(ctx context.Context, start time.Time, end time.Time) ([]*types.Reservation, error)
	ListByRoomID(ctx context.Context, roomID uuid.UUID) ([]*types.Reservation, error)
	Create(ctx context.Context, res *types.Reservation) error
}
