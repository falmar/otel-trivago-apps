package endpoint

import (
	"github.com/falmar/otel-trivago/internal/reservation/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListEndpoint               kitendpoint.Endpoint
	CreateEndpoint             kitendpoint.Endpoint
	ListAvailableRoomsEndpoint kitendpoint.Endpoint
}

func MakeEndpoints(s service.Service) *Endpoints {
	return &Endpoints{
		ListEndpoint: makeListEndpoint(s),

		CreateEndpoint:             makeCreateEndpoint(s),
		ListAvailableRoomsEndpoint: makeListAvailableRoomsEndpoint(s),
	}
}
