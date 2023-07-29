package main

import (
	"github.com/falmar/otel-trivago/internal/reservation/endpoint"
	"github.com/falmar/otel-trivago/internal/reservation/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservation/roomrepo"
	"github.com/falmar/otel-trivago/internal/reservation/service"
	"github.com/falmar/otel-trivago/internal/reservation/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"net"
)

const tracerName = "reservation-svc"

func initTracer() trace.Tracer {
	tex, err := newExporter(os.Stdout)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(newResource()),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(tex))

	return tp.Tracer(tracerName)
}

func main() {
	tr := initTracer()

	svc := service.NewService(&service.Config{
		ResvRepo: reservationrepo.NewMem(),
		RoomRepo: roomrepo.NewMem(),
	})
	svc = service.NewTracer(tr, svc)

	endpoints := endpoint.MakeEndpoints(tr, svc)
	grpcServer := transport.NewGRPCServer(tr, endpoints)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	reservationpb.RegisterReservationServiceServer(server, grpcServer)

	go func() {
		sigChan := make(chan os.Signal)

		signal.Notify(sigChan, syscall.SIGINT)
		signal.Notify(sigChan, syscall.SIGTERM)

		<-sigChan
		log.Println("stop signal received")
		server.GracefulStop()
	}()

	log.Println("Starting server on port :8080")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
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
