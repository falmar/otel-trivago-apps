package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ Service = (*serviceTracer)(nil)

func NewTracer(svc Service, tr trace.Tracer) Service {
	return &serviceTracer{
		svc: svc,
		tr:  tr,
	}
}

type serviceTracer struct {
	svc Service
	tr  trace.Tracer
}

func (s *serviceTracer) CheckAvailability(ctx context.Context, input *CheckAvailabilityInput) (*CheckAvailabilityOutput, error) {
	ctx, span := s.tr.Start(ctx, "svc.CheckAvailability")
	defer span.End()

	out, err := s.svc.CheckAvailability(ctx, input)

	defer func() {
		span.SetAttributes(
			attribute.Int64("input.start", input.Start.Unix()),
			attribute.Int64("input.end", input.End.Unix()),
			attribute.Int64("input.capacity", int64(input.Capacity)),
			attribute.Int("output.count", len(out.Rooms)),
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to list reservations")
		}
	}()

	return out, err
}

func (s *serviceTracer) CheckIn(ctx context.Context, input *CheckInInput) (*CheckInOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceTracer) CheckOut(ctx context.Context, input *CheckOutInput) (*CheckOutOutput, error) {
	//TODO implement me
	panic("implement me")
}
