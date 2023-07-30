package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
)

func NewProvider(ctx context.Context, svcName string) (*sdktrace.TracerProvider, error) {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(svcName),
			semconv.ServiceVersion("0.0.1"),
			attribute.String("environment", "dev"),
		),
	)

	tex, err := newExporter(ctx, os.Stdout)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(r),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(tex))

	return tp, nil
}

func InitTracer(name string, tp trace.TracerProvider) trace.Tracer {
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp.Tracer(name)
}

// newExporter returns a console exporter.
func newExporter(ctx context.Context, w io.Writer) (sdktrace.SpanExporter, error) {
	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") == "" {
		return stdouttrace.New(
			stdouttrace.WithWriter(w),
			// Use human-readable output.
			stdouttrace.WithPrettyPrint(),
		)
	}

	return otlptracegrpc.New(ctx)
}
