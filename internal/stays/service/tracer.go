package service

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ Service = (*svcTracer)(nil)

func NewTracer(svc Service, tr trace.Tracer) Service {
	return &svcTracer{
		svc: svc,
		tr:  tr,
	}
}

type svcTracer struct {
	svc Service
	tr  trace.Tracer
}

func (s *svcTracer) ListStays(ctx context.Context, input *ListStaysInput) (*ListStaysOutput, error) {
	ctx, span := s.tr.Start(ctx, "stays.svc.ListStays")
	defer span.End()

	out, err := s.svc.ListStays(ctx, input)

	defer func() {
		attr := []attribute.KeyValue{
			attribute.Int64("out.total", out.Total),
		}

		if input.RoomID != uuid.Nil {
			attr = append(attr, attribute.String("input.room_id", input.RoomID.String()))
		}
		if input.ReservationID != uuid.Nil {
			attr = append(attr, attribute.String("input.reservation_id", input.ReservationID.String()))
		}

		span.SetAttributes(attr...)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to list stays")
		}
	}()

	return out, err
}

func (s *svcTracer) CreateStay(ctx context.Context, input *CreateStayInput) (*CreateStayOutput, error) {
	ctx, span := s.tr.Start(ctx, "stays.svc.CreateStay")
	defer span.End()
	out, err := s.svc.CreateStay(ctx, input)

	defer func() {
		span.SetAttributes(
			attribute.String("input.room_id", input.RoomID.String()),
			attribute.String("input.reservation_id", input.ReservationID.String()),
			attribute.Int64("input.check_in", input.CheckIn.Unix()),
			attribute.String("output.stay_id", out.Stay.ID.String()),
		)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to create stay")
		}
	}()

	return out, err
}

func (s *svcTracer) UpdateStay(ctx context.Context, input *UpdateStayInput) (*UpdateStayOutput, error) {
	ctx, span := s.tr.Start(ctx, "stays.svc.UpdateStay")
	defer span.End()

	out, err := s.svc.UpdateStay(ctx, input)

	defer func() {
		attr := []attribute.KeyValue{
			attribute.String("input.stay_id", input.ID.String()),
			attribute.String("input.room_id", input.RoomID.String()),
		}

		if !input.CheckOut.IsZero() {
			attr = append(attr, attribute.Int64("input.check_out", input.CheckOut.Unix()))
		}

		span.SetAttributes(attr...)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "failed to update stay")
		}
	}()

	return out, err
}
