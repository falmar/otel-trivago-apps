package reservationsvc

import (
	"context"
	"errors"
	"fmt"
	"github.com/falmar/otel-trivago/internal/bootstrap"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/internal/reservations/repo"
	"github.com/falmar/otel-trivago/internal/reservations/service"
	"github.com/falmar/otel-trivago/internal/reservations/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
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

const svcName = "reservations-svc"

var reservationCmd = &cobra.Command{
	Use:   "reservations",
	Short: "Starts the reservations service",
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
		var grpcService reservationpb.ReservationServiceServer
		{
			svc := service.NewService(&service.Config{
				ResvRepo: repo.NewMem(),
			})
			svc = service.NewTracer(svc, otpl.Tracer)
			svc, err := service.NewMeter(svc, otpl.Meter)
			if err != nil {
				return err
			}

			endpoints := endpoint.New(svc)
			grpcService = transport.NewGRPCServer(endpoints)
		}
		// --

		// grpc server setup
		server := grpc.NewServer(
			grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		)
		reservationpb.RegisterReservationServiceServer(server, grpcService)

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
	reservationCmd.Flags().String("rooms-endpoint", "8081", "rooms service endpoint")
	_ = viper.BindPFlag("rooms.endpoint", reservationCmd.Flags().Lookup("rooms-endpoint"))

	return reservationCmd
}
