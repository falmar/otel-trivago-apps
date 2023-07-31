package rooms

import (
	"context"
	"errors"
	"fmt"
	"github.com/falmar/otel-trivago/internal/bootstrap"
	"github.com/falmar/otel-trivago/internal/rooms/endpoint"
	"github.com/falmar/otel-trivago/internal/rooms/roomrepo"
	"github.com/falmar/otel-trivago/internal/rooms/service"
	"github.com/falmar/otel-trivago/internal/rooms/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/roompb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

const svcName = "rooms-svc"

var roomsCmd = &cobra.Command{
	Use:   "rooms",
	Short: "Starts the rooms service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		// tracer/meter setup
		var otpl *bootstrap.OTPL
		{
			var err error = nil
			otpl, err = bootstrap.NewOTPL(ctx, svcName)
			if err != nil {
				return fmt.Errorf("failed to bootstrap otel: %w", err)
			}

			// shutdown otpl
			defer func() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()

				if err = otpl.Shutdown(ctx); err != nil {
					log.Println(err)
				}
			}()
		}
		// --

		// prometheus setup
		if viper.GetBool("prometheus.metrics.enable") {
			promPort := viper.GetString("prometheus.metrics.port")
			promPath := viper.GetString("prometheus.metrics.path")
			httpServer := &http.Server{Addr: ":" + promPort}

			mux := http.NewServeMux()
			mux.Handle(promPath, promhttp.Handler())

			httpServer.Handler = mux

			defer func() {
				if err := httpServer.Shutdown(ctx); err != nil {
					log.Println(err)
				}
			}()

			go func() {
				log.Printf("Starting prometheus server :%s%s", promPort, promPath)
				if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Println(err)
				}
			}()
		}
		// --

		// service setup
		var grpcService roompb.RoomServiceServer
		{
			svc := service.New(&service.Config{
				RoomRepo: roomrepo.NewMem(),
			})
			svc = service.NewTracer(svc, otpl.Tracer)

			endpoints := endpoint.New(svc, otpl.Tracer)
			grpcService = transport.NewGRPCServer(endpoints, otpl.Tracer)
		}
		// --

		// grpc server setup
		server := grpc.NewServer(
			grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		)
		roompb.RegisterRoomServiceServer(server, grpcService)

		port := viper.GetString("service.port")
		listener, err := net.Listen("tcp", ":"+port)
		if err != nil {
			return err
		}

		go func() {
			<-ctx.Done()
			server.GracefulStop()
		}()

		log.Println("starting server on port :" + port)
		return server.Serve(listener)
	},
}

func Cmd() *cobra.Command {
	return roomsCmd
}