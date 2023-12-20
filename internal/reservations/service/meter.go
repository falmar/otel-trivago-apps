package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type serviceMeter struct {
	svc Service

	listReservationsCounter  metric.Int64Counter
	createReservationCounter metric.Int64Counter
}

func NewMeter(svc Service, mt metric.Meter) (Service, error) {
	var err error
	svcMeter := &serviceMeter{svc: svc}

	svcMeter.listReservationsCounter, err = mt.Int64Counter("list_reservations")
	if err != nil {
		return nil, err
	}

	svcMeter.createReservationCounter, err = mt.Int64Counter("create_reservation")
	if err != nil {
		return nil, err
	}

	return svcMeter, err
}

func (s *serviceMeter) ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error) {
	out, err := s.svc.ListReservations(ctx, input)

	defer func() {
		s.listReservationsCounter.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}

func (s *serviceMeter) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	out, err := s.svc.CreateReservation(ctx, input)

	defer func() {
		s.createReservationCounter.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}
