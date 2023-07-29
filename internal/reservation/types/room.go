package types

import "github.com/google/uuid"

type Room struct {
	ID       uuid.UUID
	Capacity int64
}
