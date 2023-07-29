package types

type ErrReservedRoom struct {
	RoomID string
}

func (e ErrReservedRoom) Error() string {
	return "room " + e.RoomID + " is already reserved"
}
