package bootstrap

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"io"
)

func InitTracer(tp trace.TracerProvider, name, version string, attr []attribute.KeyValue) trace.Tracer {
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp.Tracer(
		name,
		trace.WithInstrumentationVersion(version),
		trace.WithInstrumentationAttributes(attr...),
	)
}

func NewGRPCExporter(ctx context.Context, endpoint string) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithInsecure(),
	)
}

func NewStdoutExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		// Use human-readable output.
		stdouttrace.WithWriter(w),
		stdouttrace.WithPrettyPrint(),
	)
}
