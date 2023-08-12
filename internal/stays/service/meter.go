package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var _ Service = (*svcMeter)(nil)

func NewMeter(svc Service, mt metric.Meter) (Service, error) {
	var err error = nil
	metered := &svcMeter{
		svc: svc,
	}

	metered.countListStays, err = mt.Int64Counter("stays.svc.list_stays")
	if err != nil {
		return nil, err
	}

	metered.countCreateStay, err = mt.Int64Counter("stays.svc.create_stay")
	if err != nil {
		return nil, err
	}

	metered.countUpdateStay, err = mt.Int64Counter("stays.svc.update_stay")
	if err != nil {
		return nil, err
	}

	return metered, nil
}

type svcMeter struct {
	svc Service

	countListStays  metric.Int64Counter
	countCreateStay metric.Int64Counter
	countUpdateStay metric.Int64Counter
}

func (s *svcMeter) ListStays(ctx context.Context, input *ListStaysInput) (*ListStaysOutput, error) {
	out, err := s.svc.ListStays(ctx, input)

	defer func() {
		s.countListStays.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}

func (s *svcMeter) CreateStay(ctx context.Context, input *CreateStayInput) (*CreateStayOutput, error) {
	out, err := s.svc.CreateStay(ctx, input)

	defer func() {
		s.countCreateStay.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}

func (s *svcMeter) UpdateStay(ctx context.Context, input *UpdateStayInput) (*UpdateStayOutput, error) {
	out, err := s.svc.UpdateStay(ctx, input)

	defer func() {
		s.countUpdateStay.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}
