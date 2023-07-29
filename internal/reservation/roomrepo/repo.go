package roomrepo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/types"
)

type Repository interface {
	List(ctx context.Context, capacity int64) ([]*types.Room, error)
}