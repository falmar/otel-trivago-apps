package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	"time"
)

type ListInput struct {
	Start time.Time
	End   time.Time

	Limit  int64
	Offset int64
}

type ListOutput struct {
	Reservations []*types.Reservation
	Total        int64
}

func (s *service) List(ctx context.Context, input *ListInput) (*ListOutput, error) {
	resv, err := s.resvRepo.List(ctx, input.Start, input.End)
	if err != nil {
		return nil, err
	}

	return &ListOutput{
		Reservations: resv,
		Total:        int64(len(resv)),
	}, nil
}
