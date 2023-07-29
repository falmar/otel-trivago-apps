package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	"time"
)

type ListAvailableRoomsInput struct {
	Start time.Time
	End   time.Time

	Capacity int64
}

type ListAvailableRoomsOutput struct {
	Rooms []*types.Room
	Total int64
}

func (s *service) ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error) {
	rooms, err := s.roomRepo.List(ctx, input.Capacity)
	if err != nil {
		return nil, err
	}

	reservations, err := s.resvRepo.List(ctx, input.Start, input.End)
	if err != nil {
		return nil, err
	}

	var availableRooms []*types.Room

	for _, room := range rooms {
		available := true

		for _, reservation := range reservations {
			if reservation.RoomID == room.ID {
				available = false
				break
			}
		}

		if available {
			availableRooms = append(availableRooms, room)
		}
	}

	return &ListAvailableRoomsOutput{
		Rooms: availableRooms,
		Total: int64(len(availableRooms)),
	}, nil
}
