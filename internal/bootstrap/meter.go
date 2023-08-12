package bootstrap

import (
	"go.opentelemetry.io/otel/attribute"
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

func InitMeter(mp metric.MeterProvider, name, version string, attr []attribute.KeyValue) metric.Meter {
	return mp.Meter(
		name,
		metric.WithInstrumentationVersion(version),
		metric.WithInstrumentationAttributes(attr...),
	)
}
