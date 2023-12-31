package main

import (
	"context"
	"github.com/falmar/otel-trivago/cmd/frontdesk"
	"github.com/falmar/otel-trivago/cmd/reservations"
	"github.com/falmar/otel-trivago/cmd/rooms"
	"github.com/falmar/otel-trivago/cmd/stays"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//if dport := viper.GetString("debug_port"); dport != "" {
		//	go func() {
		//		log.Printf("starting pprof server on port :%s\n", dport)
		//		_ = http.ListenAndServe("localhost:"+dport, nil)
		//	}()
		//}
	},
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	rootCmd.AddCommand(reservationsvc.Cmd())
	rootCmd.AddCommand(rooms.Cmd())
	rootCmd.AddCommand(stays.Cmd())
	rootCmd.AddCommand(frontdesk.Cmd())

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigFile(viper.GetString("config"))
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringP("config", "c", "./config.yaml", "config file path")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().String("prometheus-port", "9090", "port to serve prometheus metrics")
	rootCmd.PersistentFlags().String("prometheus-path", "/metrics", "path to serve prometheus metrics")
	_ = viper.BindPFlag("prometheus.metrics.port", rootCmd.PersistentFlags().Lookup("prometheus-port"))
	_ = viper.BindPFlag("prometheus.metrics.path", rootCmd.PersistentFlags().Lookup("prometheus-path"))

	rootCmd.PersistentFlags().String("service-version", "0.0.1", "service version")
	_ = viper.BindPFlag("service.version", rootCmd.PersistentFlags().Lookup("service-version"))

	rootCmd.PersistentFlags().String("otlp-endpoint", "", "otel grpc exporter endpoint eg: jaeger localhost:4317")
	_ = viper.BindPFlag("otpl_endpoint", rootCmd.PersistentFlags().Lookup("otlp-endpoint"))
	_ = viper.BindEnv("otpl_endpoint", "OTEL_EXPORTER_OTLP_ENDPOINT")

	rootCmd.PersistentFlags().String("debug-port", "", "debug port eg: 6060")
	_ = viper.BindPFlag("debug_port", rootCmd.PersistentFlags().Lookup("debug-port"))
}
