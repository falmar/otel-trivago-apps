package repo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stay/types"
)

type Repository interface {
	List(ctx context.Context, options *ListOptions) ([]*types.Stay, error)
	Create(ctx context.Context, stay *types.Stay) error
}

type ListOptions struct{}
