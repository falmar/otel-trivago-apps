package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stay/types"
	"sync"
)

var _ Repository = (*memRepository)(nil)

func NewMemRepo() Repository {
	return &memRepository{
		mu: sync.RWMutex{},
	}
}

type memRepository struct {
	mu sync.RWMutex
}

func (m *memRepository) List(ctx context.Context, options *ListOptions) ([]*types.Stay, error) {
	panic("implement me")
}

func (m *memRepository) Create(ctx context.Context, stay *types.Stay) error {
	panic("implement me")
}
