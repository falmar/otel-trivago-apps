package reservationrepo

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

func (r *memRepository) List(_ context.Context, start time.Time, end time.Time) ([]*types.Reservation, error) {
	r.mu.RLock()

	var reservations []*types.Reservation

	if !start.IsZero() && !end.IsZero() {
		for _, reservation := range r.data {
			if (reservation.Start.Equal(start) || reservation.Start.Equal(start)) &&
				(reservation.End.Equal(end) || reservation.End.Equal(end)) {
				reservations = append(reservations, reservation)
			}
		}

		r.mu.RUnlock()

		return reservations, nil
	}

	for _, reservation := range r.data {
		reservations = append(reservations, reservation)
	}

	r.mu.RUnlock()

	return reservations, nil
}

func (r *memRepository) ByRoomID(_ context.Context, roomID uuid.UUID) ([]*types.Reservation, error) {
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
