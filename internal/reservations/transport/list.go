package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func decodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
	pbreq := request.(*reservationpb.ReservationListRequest)

	req := &endpoint.ListRequest{}

	if pbreq.StartDate != nil {
		req.Start = pbreq.StartDate.AsTime().UTC()
	}
	if pbreq.EndDate != nil {
		req.End = pbreq.EndDate.AsTime().UTC()
	}

	return req, nil
}

func encodeListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*endpoint.ListResponse)
	var resv []*reservationpb.Reservation

	for _, r := range resp.Reservations {
		resv = append(resv, &reservationpb.Reservation{
			Id:     r.ID.String(),
			RoomId: r.RoomID.String(),
			Status: reservationpb.ReservationStatus(r.Status),
			StartDate: &timestamppb.Timestamp{
				Seconds: r.Start.Unix(),
			},
			EndDate: &timestamppb.Timestamp{
				Seconds: r.End.Unix(),
			},
			CreatedAt: &timestamppb.Timestamp{
				Seconds: r.CreatedAt.Unix(),
			},
		})
	}

	return &reservationpb.ReservationListResponse{
		Reservations: resv,
		Total:        resp.Total,
	}, nil
}
