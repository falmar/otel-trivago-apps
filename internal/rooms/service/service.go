package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/repo"
	"github.com/falmar/otel-trivago/internal/rooms/types"
)

var _ Service = (*service)(nil)

type Config struct {
	RoomRepo repo.Repository
}

func New(config *Config) Service {
	return &service{
		roomRepo: config.RoomRepo,
	}
}

type Service interface {
	ListRooms(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error)
}

type service struct {
	roomRepo repo.Repository
}

type ListRoomsInput struct {
	Capacity int64
	Limit    int64
	Offset   int64
}
type ListRoomsOutput struct {
	Rooms []*types.Room
	Total int64
}

func (s *service) ListRooms(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error) {
	rooms, err := s.roomRepo.List(ctx, input.Capacity)
	if err != nil {
		return nil, err
	}

	if len(rooms) > int(input.Limit) && input.Limit > 0 {
		rooms = rooms[:input.Limit]
	}

	return &ListRoomsOutput{
		Rooms: rooms,
		Total: int64(len(rooms)),
	}, nil
}
