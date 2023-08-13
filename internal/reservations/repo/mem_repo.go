package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/google/uuid"
	"sync"
	"time"
)

var _ Repository = (*memRepository)(nil)

type memRepository struct {
	data map[uuid.UUID]*types.Reservation
	mu   sync.RWMutex
}

func NewMem() Repository {
	return &memRepository{
		data: map[uuid.UUID]*types.Reservation{},
	}
}

func (r *memRepository) List(_ context.Context, opts *ListOptions) ([]*types.Reservation, error) {
	r.mu.RLock()

	var reservations []*types.Reservation

	var count int64
	for _, reservation := range r.data {
		if opts.RoomID != uuid.Nil && reservation.RoomID != opts.RoomID {
			continue
		} else if !opts.Start.IsZero() && reservation.End.Before(opts.Start) {
			continue
		} else if !opts.End.IsZero() && reservation.Start.After(opts.End) {
			continue
		}

		reservations = append(reservations, reservation)
		count++

		if count >= opts.Limit {
			break
		}
	}

	r.mu.RUnlock()

	return reservations, nil
}

func (r *memRepository) ListByRoomID(_ context.Context, roomID uuid.UUID) ([]*types.Reservation, error) {
	r.mu.RLock()
	var reservations []*types.Reservation

	for _, reservation := range r.data {
		if reservation.RoomID == roomID {
			reservations = append(reservations, reservation)
		}
	}

	r.mu.RUnlock()

	return reservations, nil
}

func (r *memRepository) Create(_ context.Context, res *types.Reservation) error {
	res.ID = uuid.New()
	res.Status = types.ReservationStatusReserved
	res.CreatedAt = time.Now()

	r.mu.Lock()
	r.data[res.ID] = res
	r.mu.Unlock()

	return nil
}
