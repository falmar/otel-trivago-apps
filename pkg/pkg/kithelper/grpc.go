package kithelper

import (
	"context"
	"errors"
	"github.com/falmar/otel-trivago/internal/reservations/types"
	"github.com/go-kit/kit/transport/grpc"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var trackerSpanCtxKey = struct{}{}

func EncodeError(_ context.Context, err error) error {
	if err == nil {
		return nil
	}

	code := codes.Internal
	msg := err.Error()

	var eInvalidArgument *ErrInvalidArgument
	if errors.As(err, &eInvalidArgument) {
		code = codes.InvalidArgument
		msg = eInvalidArgument.Error()
	}
	var eReserved *types.ErrReservedRoom
	if errors.As(err, &eReserved) {
		code = codes.AlreadyExists
		msg = eReserved.Error()
	}

	//span := ctx.Value(trackerSpanCtxKey).(trace.Span)
	//span.RecordError(err)
	//span.SetStatus(otelcodes.Error, "transport error")
	//defer span.End()

	return status.Error(code, msg)
}

func SpanTraceBefore(tracer trace.Tracer, name string) grpc.ServerRequestFunc {
	return func(ctx context.Context, _ metadata.MD) context.Context {
		ctx, span := tracer.Start(ctx, name)
		return context.WithValue(ctx, trackerSpanCtxKey, span)
	}
}

func SpanTraceAfter(ctx context.Context, _ *metadata.MD, _ *metadata.MD) context.Context {
	span := ctx.Value(trackerSpanCtxKey).(trace.Span)
	span.SetStatus(otelcodes.Ok, "")
	span.End()

	return ctx
}
