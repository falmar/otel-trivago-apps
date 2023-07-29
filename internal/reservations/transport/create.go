package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func decodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.Reservation)
	req := &endpoint.CreateRequest{
		RoomID: pbreq.RoomId,
	}

	if pbreq.StartDate != nil {
		req.Start = pbreq.StartDate.AsTime().UTC()
	}
	if pbreq.EndDate != nil {
		req.End = pbreq.EndDate.AsTime().UTC()
	}

	return req, nil
}

func encodeCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resv := response.(*endpoint.CreateResponse).Reservation

	return &reservationpb.ReservationResponse{
		Reservation: &reservationpb.Reservation{
			Id:     resv.ID.String(),
			RoomId: resv.RoomID.String(),
			Status: reservationpb.ReservationStatus(resv.Status),
			StartDate: &timestamppb.Timestamp{
				Seconds: resv.Start.Unix(),
			},
			EndDate: &timestamppb.Timestamp{
				Seconds: resv.End.Unix(),
			},
			CreatedAt: &timestamppb.Timestamp{
				Seconds: resv.CreatedAt.Unix(),
			},
		},
	}, nil
}
