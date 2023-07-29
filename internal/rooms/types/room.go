package types

import "github.com/google/uuid"

type CleanStatus int

const (
	CleanStatusClean = iota
	CleanStatusDirty
)

type Room struct {
	ID       uuid.UUID
	Capacity int64

	CleanStatus CleanStatus
}
