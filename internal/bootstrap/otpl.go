package bootstrap

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"os"
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

type OTPLConfig struct {
	ServiceName          string
	ServiceVersion       string
	GRPCExporterEndpoint string
	InstrumentAttributes []attribute.KeyValue
}

func NewOTPL(ctx context.Context, config *OTPLConfig) (*OTPL, error) {
	re, err := NewResource(config.InstrumentAttributes)
	if err != nil {
		return nil, err
	}

	// tracing
	var ex sdktrace.SpanExporter
	if config.GRPCExporterEndpoint != "" {
		ex, err = NewGRPCExporter(ctx, config.GRPCExporterEndpoint)
	} else {
		ex, err = NewStdoutExporter(os.Stdout)
	}
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(re),
		sdktrace.WithBatcher(ex),
	)
	tr := InitTracer(tp, config.ServiceName, config.ServiceVersion, config.InstrumentAttributes)
	// --

	// metrics
	mr, err := NewMeterReader()
	if err != nil {
		return nil, err
	}

	mp, err := NewMeterProvider(mr, re)
	if err != nil {
		return nil, err
	}

	mt := InitMeter(mp, config.ServiceName, config.ServiceVersion, config.InstrumentAttributes)
	// --

	return &OTPL{
		Tracer: tr,
		Meter:  mt,

		mr: mr,
		tp: tp,
	}, nil
}

func NewResource(attr []attribute.KeyValue) (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			attr...,
		),
	)
}
