package transport

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/endpoint"
	"github.com/falmar/otel-trivago/pkg/proto/v1/reservationpb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func decodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*reservationpb.Reservation)

	return endpoint.CreateRequest{
		RoomID: uuid.MustParse(req.RoomId),
		Start:  time.Unix(req.StartDate.Seconds, 0),
		End:    time.Unix(req.EndDate.Seconds, 0),
	}, nil
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
