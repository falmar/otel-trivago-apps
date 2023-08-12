package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ Service = (*svcTracer)(nil)

func NewTracer(svc Service, tr trace.Tracer) Service {
	return &svcTracer{
		svc:    svc,
		tracer: tr,
	}
}

type svcTracer struct {
	tracer trace.Tracer
	svc    Service
}

func (t *svcTracer) ListReservations(ctx context.Context, input *ListReservationsInput) (*ListReservationsOutput, error) {
	ctx, span := t.tracer.Start(ctx, "svc.ListRooms")
	defer span.End()

	out, err := t.svc.ListReservations(ctx, input)

	defer func() {
		attr := []attribute.KeyValue{
			attribute.Int("output.count", len(out.Reservations)),
		}

		if !input.Start.IsZero() {
			attr = append(attr, attribute.Int64("input.start", input.Start.Unix()))
		}
		if !input.End.IsZero() {
			attr = append(attr, attribute.Int64("input.end", input.End.Unix()))
		}

		span.SetAttributes(attr...)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to list reservations")
		}
	}()

	return out, err
}

func (t *svcTracer) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	ctx, span := t.tracer.Start(ctx, "svc.CreateReservation")
	defer span.End()

	out, err := t.svc.CreateReservation(ctx, input)

	defer func() {
		span.SetAttributes(
			attribute.Int64("input.start", input.Start.Unix()),
			attribute.Int64("input.end", input.End.Unix()),
			attribute.Float64("input.days", input.End.Sub(input.Start).Hours()/24),
			attribute.String("input.room_id", input.RoomID.String()),
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to create reservation")
		}
	}()

	return out, err
}
