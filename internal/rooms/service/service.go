package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/rooms/roomrepo"
	"github.com/falmar/otel-trivago/internal/rooms/types"
)

var _ Service = (*service)(nil)

type Service interface {
	ListRooms(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error)
}

type service struct {
	roomRepo roomrepo.Repository
}

type Config struct {
	RoomRepo roomrepo.Repository
}

func New(config *Config) Service {
	return &service{
		roomRepo: config.RoomRepo,
	}
}

type ListRoomsInput struct {
	Capacity int64
	Limit    int64
	Offset   int64
}

type ListRoomsOutput struct {
	Rooms []*types.Room
}

func (s *service) ListRooms(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error) {
	rooms, err := s.roomRepo.List(ctx, input.Capacity)
	if err != nil {
		return nil, err
	}

	if len(rooms) > int(input.Limit) {
		rooms = rooms[:input.Limit]
	}

	return &ListRoomsOutput{
		Rooms: rooms,
	}, nil
}
