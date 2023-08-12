package endpoint

import (
	"context"
	"github.com/falmar/otel-trivago/internal/stays/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListStaysEndpoint  kitendpoint.Endpoint
	CreateStayEndpoint kitendpoint.Endpoint
	UpdateStayEndpoint kitendpoint.Endpoint
}

func New(svc service.Service) Endpoints {
	return Endpoints{
		ListStaysEndpoint:  MakeListStaysEndpoint(svc),
		CreateStayEndpoint: MakeCreateStayEndpoint(svc),
		UpdateStayEndpoint: MakeUpdateStayEndpoint(svc),
	}
}

func MakeListStaysEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		panic("implement me")
	}
}

func MakeCreateStayEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		panic("implement me")
	}
}

func MakeUpdateStayEndpoint(svc service.Service) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		panic("implement me")
	}
}
