package roomrepo

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservation/types"
	"github.com/google/uuid"
)

var _ Repository = (*memRepository)(nil)

type memRepository struct {
	rooms map[uuid.UUID]*types.Room
}

func NewMem() Repository {
	var rooms = []*types.Room{
		{ID: uuid.MustParse("7e2d4217-336e-47b2-9247-881f8c843921"), Capacity: 2},
		{ID: uuid.MustParse("7e2d4217-336e-47b2-9247-881f8c843922"), Capacity: 2},
		{ID: uuid.MustParse("7e2d4217-336e-47b2-9247-881f8c843941"), Capacity: 4},
		{ID: uuid.MustParse("7e2d4217-336e-47b2-9247-881f8c843911"), Capacity: 1},
		{ID: uuid.MustParse("7e2d4217-336e-47b2-9247-881f8c843942"), Capacity: 4},
	}
	var mapRooms = map[uuid.UUID]*types.Room{}

	for _, room := range rooms {
		mapRooms[room.ID] = room
	}

	return &memRepository{
		rooms: mapRooms,
	}
}

func (r *memRepository) List(_ context.Context, capacity int64) ([]*types.Room, error) {
	var rooms []*types.Room

	for _, room := range r.rooms {
		if room.Capacity >= capacity {
			rooms = append(rooms, room)
		}
	}

	return rooms, nil
}
