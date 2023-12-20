package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var _ Service = (*svcMeter)(nil)

func NewMeter(svc Service, mt metric.Meter) (Service, error) {
	var err error
	metered := &svcMeter{svc: svc}

	metered.listRoomsCounter, err = mt.Int64Counter("list_rooms")
	if err != nil {
		return nil, err
	}

	return metered, nil
}

type svcMeter struct {
	svc Service

	listRoomsCounter metric.Int64Counter
}

func (s *svcMeter) ListRooms(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error) {

	out, err := s.svc.ListRooms(ctx, input)

	defer func() {
		s.listRoomsCounter.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}
