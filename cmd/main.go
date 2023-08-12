package main

import (
	"context"
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

var rootCmd = &cobra.Command{}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	rootCmd.AddCommand(reservationsvc.Cmd())
	rootCmd.AddCommand(rooms.Cmd())
	rootCmd.AddCommand(stays.Cmd())

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
}
