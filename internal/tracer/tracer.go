package tracer

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"os"
)

func NewProvider(svcName string) (*sdktrace.TracerProvider, error) {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(svcName),
			semconv.ServiceVersion("0.0.1"),
			attribute.String("environment", "dev"),
		),
	)

	tex, err := newExporter(os.Stdout)
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
func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	if os.Getenv("JAEGER_ENDPOINT") == "" {
		return stdouttrace.New(
			stdouttrace.WithWriter(w),
			// Use human-readable output.
			stdouttrace.WithPrettyPrint(),
		)
	}

	return jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(os.Getenv("JAEGER_ENDPOINT"))),
	)
}
