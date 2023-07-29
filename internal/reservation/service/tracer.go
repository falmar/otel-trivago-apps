package service

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ Service = (*svcTracer)(nil)

func NewTracer(tr trace.Tracer, svc Service) Service {
	return &svcTracer{
		svc:    svc,
		tracer: tr,
	}
}

type svcTracer struct {
	tracer trace.Tracer
	svc    Service
}

func (t *svcTracer) List(ctx context.Context, input *ListInput) (*ListOutput, error) {
	ctx, span := t.tracer.Start(ctx, "reservation.service.List")
	defer span.End()

	out, err := t.svc.List(ctx, input)

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

func (t *svcTracer) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	ctx, span := t.tracer.Start(ctx, "reservation.service.Create")
	defer span.End()

	out, err := t.svc.Create(ctx, input)

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

func (t *svcTracer) ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error) {
	ctx, span := t.tracer.Start(ctx, "reservation.service.ListAvailableRooms")
	defer span.End()

	out, err := t.svc.ListAvailableRooms(ctx, input)

	defer func() {
		span.SetAttributes(
			attribute.Int("output.count", len(out.Rooms)),
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to list available rooms")
		}
	}()

	return out, err
}
