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
	ctx, span := t.tracer.Start(ctx, "reservations.svc.ListRooms")
	defer span.End()

	out, err := t.svc.ListReservations(ctx, input)

	defer func() {
		span.SetAttributes(
			attribute.Int("output.count", len(out.Reservations)),
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to list reservations")
		}
	}()

	return out, err
}

func (t *svcTracer) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	ctx, span := t.tracer.Start(ctx, "reservations.svc.CreateReservation")
	defer span.End()

	out, err := t.svc.CreateReservation(ctx, input)

	defer func() {
		span.SetAttributes(
			attribute.String("output.room_id", input.RoomID.String()),
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to create reservation")
		}
	}()

	return out, err
}
