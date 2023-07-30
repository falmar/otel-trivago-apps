package otelsvc

import (
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewMeterProvider(r sdkmetric.Reader, re *resource.Resource) (metric.MeterProvider, error) {
	return sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(r),
		sdkmetric.WithResource(re),
	), nil
}

func NewMeterReader() (*prometheus.Exporter, error) {
	return prometheus.New()
}

func InitMeter(name string, mp metric.MeterProvider) metric.Meter {
	return mp.Meter(name)
}
