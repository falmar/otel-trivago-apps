package endpoint

import (
	"context"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func MakeTracerEndpointMiddleware(name string, tracer trace.Tracer, next kitendpoint.Endpoint) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctx, span := tracer.Start(ctx, name)
		defer span.End()

		resp, err := next(ctx, request)

		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, "endpoint received error")
			}
		}()

		return resp, err
	}
}
