package main

import (
	"github.com/falmar/otel-trivago/internal/reservation/endpoint"
	"github.com/falmar/otel-trivago/internal/reservation/reservationrepo"
	"github.com/falmar/otel-trivago/internal/reservation/roomrepo"
	"github.com/falmar/otel-trivago/internal/reservation/service"
	"github.com/falmar/otel-trivago/internal/reservation/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"log"

	"google.golang.org/grpc"
	"net"
)

func main() {
	svc := service.NewService(&service.Config{
		ResvRepo: reservationrepo.NewMem(),
		RoomRepo: roomrepo.NewMem(),
	})
	endpoints := endpoint.MakeEndpoints(svc)
	grpcServer := transport.NewGRPCServer(endpoints)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	reservationpb.RegisterReservationServiceServer(server, grpcServer)

	log.Println("Starting server on port :8080")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
