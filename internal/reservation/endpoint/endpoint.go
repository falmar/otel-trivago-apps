package endpoint

import (
	"github.com/falmar/otel-trivago/internal/reservation/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"go.opentelemetry.io/otel/trace"
)

type Endpoints struct {
	ListEndpoint               kitendpoint.Endpoint
	CreateEndpoint             kitendpoint.Endpoint
	ListAvailableRoomsEndpoint kitendpoint.Endpoint
}

func MakeEndpoints(tr trace.Tracer, s service.Service) *Endpoints {
	return &Endpoints{
		ListEndpoint: MakeTrackerEndpoint(
			"reservation.endpoint.List",
			tr,
			makeListEndpoint(s),
		),

		ListAvailableRoomsEndpoint: MakeTrackerEndpoint(
			"reservation.endpoint.ListAvailableRooms",
			tr,
			makeListAvailableRoomsEndpoint(s),
		),

		CreateEndpoint: MakeTrackerEndpoint(
			"reservation.endpoint.Create",
			tr,
			makeCreateEndpoint(s),
		),
	}
}
