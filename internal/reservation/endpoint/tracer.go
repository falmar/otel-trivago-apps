package endpoint

import (
	"context"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"go.opentelemetry.io/otel/trace"
)

func MakeTrackerEndpoint(name string, tracer trace.Tracer, next kitendpoint.Endpoint) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctx, span := tracer.Start(ctx, name)
		defer span.End()

		return next(ctx, request)
	}
}
