package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stay/types"
	"github.com/google/uuid"
	"sync"
	"time"
)

var _ Repository = (*memRepository)(nil)

func NewMemRepo() Repository {
	return &memRepository{
		mu:   sync.RWMutex{},
		data: map[string]*types.Stay{},
	}
}

type memRepository struct {
	mu   sync.RWMutex
	data map[string]*types.Stay
}

func (m *memRepository) List(_ context.Context, options *ListOptions) ([]*types.Stay, error) {
	var stays []*types.Stay
	m.mu.RLock()

	for _, stay := range m.data {
		if options.RoomID != "" && stay.RoomID.String() != options.RoomID {
			continue
		}
		if options.ReservationID != "" && stay.ReservationID.String() != options.ReservationID {
			continue
		}

		stays = append(stays, stay)
	}

	m.mu.RUnlock()

	return stays, nil
}

func (m *memRepository) Create(_ context.Context, stay *types.Stay) error {
	m.mu.Lock()

	stay.ID = uuid.New()
	stay.CreatedAt = time.Now()
	stay.UpdatedAt = stay.CreatedAt

	m.data[stay.ID.String()] = stay

	m.mu.Unlock()

	return nil
}

func (m *memRepository) Update(_ context.Context, stay *types.Stay) error {
	m.mu.Lock()

	stay.UpdatedAt = time.Now()

	m.data[stay.ID.String()] = stay

	m.mu.Unlock()

	return nil
}
