package service

import (
	"context"
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

	return t.svc.List(ctx, input)
}

func (t *svcTracer) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	ctx, span := t.tracer.Start(ctx, "reservation.service.Create")
	defer span.End()

	return t.svc.Create(ctx, input)
}

func (t *svcTracer) ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error) {
	ctx, span := t.tracer.Start(ctx, "reservation.service.ListAvailableRooms")
	defer span.End()

	return t.svc.ListAvailableRooms(ctx, input)
}
