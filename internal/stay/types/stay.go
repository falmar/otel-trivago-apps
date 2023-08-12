package types

import "github.com/google/uuid"

type Stay struct {
	ID            uuid.UUID
	RoomID        uuid.UUID
	ReservationID uuid.UUID
}
