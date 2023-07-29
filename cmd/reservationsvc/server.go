package main

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/internal/reservations/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/transport"
	"github.com/falmar/otel-trivago/internal/rooms/roomrepo"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"net"
)

const tracerName = "reservation-svc"

func newProvider() (*sdktrace.TracerProvider, error) {
	tex, err := newExporter(os.Stdout)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(newResource()),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(tex))

	return tp, nil
}

func initTracer(tp trace.TracerProvider) trace.Tracer {
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp.Tracer(tracerName)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// tracer setup
	tp, err := newProvider()
	if err != nil {
		log.Fatalln(err)
	}
	tr := initTracer(tp)

	svc := service.NewService(&service.Config{
		ResvRepo: reservationrepo.NewMem(),
		RoomRepo: roomrepo.NewMem(),
	})
	svc = service.NewTracer(tr, svc)
	// --

	// service setup
	endpoints := endpoint.MakeEndpoints(tr, svc)
	grpcServer := transport.NewGRPCServer(tr, endpoints)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)

	reservationpb.RegisterReservationServiceServer(server, grpcServer)
	// --

	defer func() {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := tp.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		sigChan := make(chan os.Signal)

		signal.Notify(sigChan, syscall.SIGINT)
		signal.Notify(sigChan, syscall.SIGTERM)

		<-sigChan
		log.Println("stop signal received")
		server.GracefulStop()
	}()

	log.Println("Starting server on port :" + port)
	if err := server.Serve(listener); err != nil {
		log.Fatalln(err)
	}
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

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("reservation-svc"),
			semconv.ServiceVersion("0.0.1"),
			attribute.String("environment", "dev"),
		),
	)
	return r
}
