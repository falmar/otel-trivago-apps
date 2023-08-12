package types

import (
	"github.com/google/uuid"
	"time"
)

type Stay struct {
	ID            uuid.UUID
	RoomID        uuid.UUID
	ReservationID uuid.UUID

	CheckIn  time.Time
	CheckOut time.Time

	Notes string
}
