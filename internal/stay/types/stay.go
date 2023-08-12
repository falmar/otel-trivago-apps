package types

import (
	"github.com/google/uuid"
	"time"
)

type Stay struct {
	ID            uuid.UUID
	RoomID        uuid.UUID
	ReservationID uuid.UUID

	CheckedInAt  time.Time
	CheckedOutAt time.Time

	Notes string
}
