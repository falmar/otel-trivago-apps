package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ Service = (*serviceTracer)(nil)

type serviceTracer struct {
	svc Service

	tr trace.Tracer
}

func NewTracer(svc Service, tr trace.Tracer) Service {
	return &serviceTracer{svc: svc, tr: tr}
}

func (s *serviceTracer) ListRooms(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error) {
	ctx, span := s.tr.Start(ctx, "svc.ListRooms")
	defer span.End()

	out, err := s.svc.ListRooms(ctx, input)

	defer func() {
		attr := []attribute.KeyValue{
			attribute.Int64("output.count", out.Total),
		}

		if input.Capacity != 0 {
			attr = append(attr, attribute.Int64("input.capacity", input.Capacity))
		}

		span.SetAttributes(attr...)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to list rooms")
		}
	}()

	return out, err
}
