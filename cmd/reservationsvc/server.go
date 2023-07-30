package main

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/internal/reservations/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/transport"
	roomtransport "github.com/falmar/otel-trivago/internal/rooms/transport"
	"github.com/falmar/otel-trivago/internal/tracer"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"net"
)

const svcName = "reservation-svc"
const tracerName = "reservation-svc"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// tracer setup
	tp, err := tracer.NewProvider(ctx, svcName)
	if err != nil {
		log.Fatalln(err)
	}
	tr := tracer.InitTracer(tracerName, tp)
	// --

	// service setup
	roomHost := os.Getenv("ROOM_HOST")
	if roomHost == "" {
		roomHost = "localhost:8081"
	}

	roomConn, err := grpc.DialContext(ctx, roomHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalln(err)
	}

	svc := service.NewService(&service.Config{
		ResvRepo: reservationrepo.NewMem(),
		RoomSvc:  roomtransport.NewGRPCClient(roomConn),
	})
	svc = service.NewTracer(svc, tr)

	endpoints := endpoint.New(tr, svc)
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
