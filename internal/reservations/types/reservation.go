package types

import (
	"github.com/google/uuid"
	"time"
)

type ReservationStatus int

const (
	ReservationStatusReserved ReservationStatus = iota
	ReservationStatusCheckedIn
	ReservationStatusCheckedOut
	ReservationStatusCancelled
)

type Reservation struct {
	ID     uuid.UUID
	RoomID uuid.UUID
	Status ReservationStatus

	Start time.Time
	End   time.Time

	CreatedAt time.Time
	UpdateAt  time.Time
}
