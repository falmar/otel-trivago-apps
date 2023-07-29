package main

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/internal/rooms/roomrepo"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/internal/rooms/transport"
	"github.com/falmar/otel-trivago/internal/tracer"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"net"
)

const svcName = "room-svc"
const tracerName = "room-svc"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// tracer setup
	tp, err := tracer.NewProvider(svcName)
	if err != nil {
		log.Fatalln(err)
	}
	tr := tracer.InitTracer(tracerName, tp)
	// --

	// service setup
	svc := service.New(&service.Config{
		RoomRepo: roomrepo.NewMem(),
	})
	//svc = service.NewTracer(tr, svc)

	endpoints := endpoint.New(tr, svc)
	grpcServer := transport.NewGRPCServer(endpoints, tr)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)

	roompb.RegisterRoomServiceServer(server, grpcServer)
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
