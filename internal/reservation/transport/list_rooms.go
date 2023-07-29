package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
)

func decodeListAvailableRoomsRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.RoomAvailabilityRequest)
	req := &endpoint.ListAvailableRoomsRequest{}

	req.Capacity = pbreq.Capacity
	if pbreq.StartDate != nil {
		req.Start = pbreq.StartDate.AsTime().UTC()
	}
	if pbreq.EndDate != nil {
		req.End = pbreq.EndDate.AsTime().UTC()
	}

	return req, nil
}

func encodeListAvailableRoomsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.ListAvailableRoomsResponse)

	var rooms []*reservationpb.Room
	for _, room := range resp.Rooms {
		rooms = append(rooms, &reservationpb.Room{
			Id:       room.ID.String(),
			Capacity: room.Capacity,
		})
	}

	return &reservationpb.RoomAvailabilityResponse{
		Rooms: rooms,
		Total: resp.Total,
	}, nil
}
