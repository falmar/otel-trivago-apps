package service

import (
	"context"
	"github.com/falmar/krun"
	reservationsvc "github.com/falmar/otel-trivago/internal/reservations/service"
	roomsvc "github.com/falmar/otel-trivago/internal/rooms/service"
	roomtypes "github.com/falmar/otel-trivago/internal/rooms/types"
	staysvc "github.com/falmar/otel-trivago/internal/stays/service"
	"time"
)

var _ Service = (*service)(nil)

type roomCtxKey struct{}

var roomKey = roomCtxKey{}

type Config struct {
	RoomService         roomsvc.Service
	ReservationsService reservationsvc.Service
	StaysService        staysvc.Service

	KQueue krun.Krun
}

func New(cfg *Config) Service {
	return &service{
		roomService:         cfg.RoomService,
		reservationsService: cfg.ReservationsService,
		staysService:        cfg.StaysService,
		kQueue:              cfg.KQueue,
	}
}

type Service interface {
	CheckAvailability(ctx context.Context, input *CheckAvailabilityInput) (*CheckAvailabilityOutput, error)
	CheckIn(ctx context.Context, input *CheckInInput) (*CheckInOutput, error)
	CheckOut(ctx context.Context, input *CheckOutInput) (*CheckOutOutput, error)
}

type service struct {
	roomService         roomsvc.Service
	reservationsService reservationsvc.Service
	staysService        staysvc.Service

	kQueue krun.Krun
}

type CheckAvailabilityInput struct {
	Capacity int64
	Start    time.Time
	End      time.Time
}
type CheckAvailabilityOutput struct {
	Rooms []*roomtypes.Room
}

func (s *service) CheckAvailability(ctx context.Context, input *CheckAvailabilityInput) (*CheckAvailabilityOutput, error) {
	// call room service to list rooms by a given capacity
	// then check against reservation service to discard rooms that are already reserved
	roomOut, err := s.roomService.ListRooms(ctx, &roomsvc.ListRoomsInput{
		Capacity: input.Capacity,
		Limit:    10,
		Offset:   0,
	})
	if err != nil {
		return nil, err
	}

	resInput := &reservationsvc.ListReservationsInput{
		Start: input.Start,
		End:   input.End,
		Limit: 1,
	}

	rooms := make([]*roomtypes.Room, 0)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	resultChan := make(chan *krun.Result)

	for _, r := range roomOut.Rooms {
		ctx := context.WithValue(ctx, roomKey, r)
		result := s.kQueue.Run(ctx, func(ctx context.Context) (interface{}, error) {
			room := ctx.Value(roomKey).(*roomtypes.Room)

			resInput.RoomID = room.ID
			resvOut, err := s.reservationsService.ListReservations(ctx, resInput)
			if err != nil {
				return nil, err
			}

			if resvOut.Total > 0 {
				return nil, nil
			}

			return room, nil
		})

		go func(result <-chan *krun.Result) { resultChan <- <-result }(result)
	}

	n := len(roomOut.Rooms)

loop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case result := <-resultChan:
			if result.Error != nil {
				cancel()
				return nil, result.Error
			}

			if room := result.Data.(*roomtypes.Room); room != nil {
				rooms = append(rooms, room)
			}

			if n--; n == 0 {
				break loop
			}
		}
	}

	return &CheckAvailabilityOutput{
		Rooms: rooms,
	}, nil
}

type CheckInInput struct{}
type CheckInOutput struct{}

func (s *service) CheckIn(ctx context.Context, input *CheckInInput) (*CheckInOutput, error) {
	// call reservations service to check if reservation exists for the given id
	// call stays service to create a new stay
	panic("implement me")
}

type CheckOutInput struct{}
type CheckOutOutput struct{}

func (s *service) CheckOut(ctx context.Context, input *CheckOutInput) (*CheckOutOutput, error) {
	// call stays service to check if stay exists for the given id
	// call reservations service to check if reservation exists for the given id
	// call stays service to update stay with check out date
	panic("implement me")
}
