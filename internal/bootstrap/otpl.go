package bootstrap

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type OTPL struct {
	Tracer trace.Tracer
	Meter  metric.Meter

	mr sdkmetric.Reader
	tp *sdktrace.TracerProvider
}

func (o *OTPL) Shutdown(ctx context.Context) error {
	if err := o.mr.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown meter reader: %w", err)
	}

	if err := o.tp.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown tracer provider: %w", err)
	}

	return nil
}

func NewOTPL(ctx context.Context, svcName string) (*OTPL, error) {
	re, err := NewResource(svcName)
	if err != nil {
		return nil, err
	}

	tp, err := NewTracerProvider(ctx, re)
	if err != nil {
		return nil, err
	}
	tr := InitTracer(svcName, tp)

	mr, err := NewMeterReader()
	if err != nil {
		return nil, err
	}

	mp, err := NewMeterProvider(mr, re)
	if err != nil {
		return nil, err
	}

	mt := InitMeter(svcName, mp)

	return &OTPL{
		Tracer: tr,
		Meter:  mt,

		mr: mr,
		tp: tp,
	}, nil
}
