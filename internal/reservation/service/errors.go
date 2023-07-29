package service

type ErrRoomReserved struct {
	RoomID string
}

func (e ErrRoomReserved) Error() string {
	return "room " + e.RoomID + " is already reserved"
}
