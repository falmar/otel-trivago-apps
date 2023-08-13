package frontdesk

import (
	"context"
	"errors"
	"fmt"
	"github.com/falmar/krun"
	"github.com/falmar/otel-trivago/internal/bootstrap"
	"github.com/falmar/otel-trivago/internal/frontdesk/endpoint"
	"github.com/falmar/otel-trivago/internal/frontdesk/service"
	"github.com/falmar/otel-trivago/internal/frontdesk/transport"
	reservationtransport "github.com/falmar/otel-trivago/internal/reservations/transport"
	roomtransport "github.com/falmar/otel-trivago/internal/rooms/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/frontdeskpb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"time"
)

const svcName = "frontdesk-svc"

var frontdeskCmd = &cobra.Command{
	Use:   "frontdesk",
	Short: "Starts the frontdesk service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		version := viper.GetString("version")
		if version == "" {
			version = "0.0.0"
		}

		// tracer/meter setup
		var otpl *bootstrap.OTPL
		{
			var err error = nil
			otpl, err = bootstrap.NewOTPL(ctx, &bootstrap.OTPLConfig{
				ServiceName:          svcName,
				ServiceVersion:       version,
				GRPCExporterEndpoint: viper.GetString("otpl_endpoint"),
				InstrumentAttributes: []attribute.KeyValue{
					semconv.ServiceName(svcName),
					semconv.ServiceVersion(version),
					semconv.DeploymentEnvironment("dev"),
				},
			})
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
				log.Printf("starting prometheus server :%s%s", promPort, promPath)
				if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Println(err)
				}
			}()
		}
		// --

		// service setup
		var grpcService frontdeskpb.FrontDeskServiceServer
		kQueue := krun.New(&krun.Config{
			Size:      3,
			WaitSleep: time.Millisecond * 20,
		})
		{
			roomConn, err := grpc.DialContext(ctx, viper.GetString("rooms.endpoint"),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			)
			if err != nil {
				return err
			}
			defer roomConn.Close()

			reservationConn, err := grpc.DialContext(ctx, viper.GetString("reservations.endpoint"),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			)
			if err != nil {
				return err
			}
			defer reservationConn.Close()

			roomSvc := roomtransport.NewGRPCClient(roomConn)
			reservationSvc := reservationtransport.NewGRPCClient(reservationConn)

			svc := service.New(&service.Config{
				RoomService:         roomSvc,
				ReservationsService: reservationSvc,
				KQueue:              kQueue,
			})
			svc = service.NewTracer(svc, otpl.Tracer)
			svc, err = service.NewMeter(svc, otpl.Meter)
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
		frontdeskpb.RegisterFrontDeskServiceServer(server, grpcService)

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
		err = server.Serve(listener)

		// wait for all pending requests to finish
		ctx, cancel2 := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel2()
		kQueue.Wait(ctx)

		return err
	},
}

func Cmd() *cobra.Command {
	return frontdeskCmd
}
