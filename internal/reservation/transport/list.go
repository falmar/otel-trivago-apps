package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func decodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*reservationpb.ReservationListRequest)

	return endpoint.ListRequest{
		Offset: req.Offset,
		Limit:  req.Limit,
	}, nil
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
