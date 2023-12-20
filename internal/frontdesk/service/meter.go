package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var _ Service = (*serviceMeter)(nil)

func NewMeter(svc Service, mt metric.Meter) (Service, error) {
	var err error
	metered := &serviceMeter{
		svc: svc,
	}

	metered.checkAvailabilityCounter, err = mt.Int64Counter("svc.check_availability")
	if err != nil {
		return nil, err
	}

	return metered, nil
}

type serviceMeter struct {
	svc Service

	checkAvailabilityCounter metric.Int64Counter
}

func (m *serviceMeter) CheckAvailability(ctx context.Context, input *CheckAvailabilityInput) (*CheckAvailabilityOutput, error) {
	out, err := m.svc.CheckAvailability(ctx, input)

	defer func() {
		m.checkAvailabilityCounter.Add(ctx, 1, metric.WithAttributes(
			attribute.Bool("error", err != nil),
		))
	}()

	return out, err
}

func (m *serviceMeter) CheckIn(ctx context.Context, input *CheckInInput) (*CheckInOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (m *serviceMeter) CheckOut(ctx context.Context, input *CheckOutInput) (*CheckOutOutput, error) {
	//TODO implement me
	panic("implement me")
}
