package stays

import (
	"context"
	"errors"
	"fmt"
	"github.com/falmar/otel-trivago/internal/bootstrap"
	"github.com/falmar/otel-trivago/internal/stays/endpoint"
	"github.com/falmar/otel-trivago/internal/stays/repo"
	"github.com/falmar/otel-trivago/internal/stays/service"
	"github.com/falmar/otel-trivago/internal/stays/transport"
	"github.com/falmar/otel-trivago/pkg/proto/v1/staypb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

const svcName = "stays-svc"

var staysCmd = &cobra.Command{
	Use:   "stays",
	Short: "Starts the stays service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		version := viper.GetString("version")
		if version == "" {
			version = "0.0.0"
		}

		// tracer/meter setup
		var otlp *bootstrap.OTLP
		{
			var err error = nil
			otlp, err = bootstrap.NewOTLP(ctx, &bootstrap.OTLPConfig{
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

			// shutdown otlp
			defer func() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()

				if err = otlp.Shutdown(ctx); err != nil {
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
		var grpcService staypb.StayServiceServer
		{
			svc := service.New(&service.Config{
				Repo: repo.NewMemRepo(),
			})
			svc = service.NewTracer(svc, otlp.Tracer)
			svc, err := service.NewMeter(svc, otlp.Meter)
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
		staypb.RegisterStayServiceServer(server, grpcService)

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
	return staysCmd
}
