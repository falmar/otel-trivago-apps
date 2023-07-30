package main

import (
	"context"
	"github.com/falmar/otel-trivago/internal/otelsvc"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/internal/reservations/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/transport"
	roomtransport "github.com/falmar/otel-trivago/internal/rooms/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"net"
)

const svcName = "reservation-svc"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// tracer/meter setup
	re, err := otelsvc.NewResource(svcName)
	if err != nil {
		log.Fatalln(err)
	}

	tp, err := otelsvc.NewTracerProvider(ctx, re)
	if err != nil {
		log.Fatalln(err)
	}
	tr := otelsvc.InitTracer(svcName, tp)

	mr, err := otelsvc.NewMeterReader()
	if err != nil {
		log.Fatalln(err)
	}

	mp, err := otelsvc.NewMeterProvider(mr, re)
	if err != nil {
		log.Fatalln(err)
	}

	mt := otelsvc.InitMeter(svcName, mp)
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
	svc, err = service.NewMeter(svc, mt)
	if err != nil {
		log.Fatalln(err)
	}

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
	defer func() {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := mr.Shutdown(ctx); err != nil {
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

	go func() {
		log.Println("serving metrics at :9090/metrics")
		http.Handle("/metrics", promhttp.Handler())

		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("starting server on port :" + port)
	if err := server.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
