package service

import "context"

var _ Service = (*service)(nil)

type Config struct{}

func NewService(cfg *Config) Service {
	return &service{}
}

type Service interface {
	CheckAvailability(ctx context.Context, input *CheckAvailabilityInput) (*CheckAvailabilityOutput, error)
	CheckIn(ctx context.Context, input *CheckInInput) (*CheckInOutput, error)
	CheckOut(ctx context.Context, input *CheckOutInput) (*CheckOutOutput, error)
}

type service struct{}

type CheckAvailabilityInput struct{}
type CheckAvailabilityOutput struct{}

func (s *service) CheckAvailability(ctx context.Context, input *CheckAvailabilityInput) (*CheckAvailabilityOutput, error) {
	panic("implement me")
}

type CheckInInput struct{}
type CheckInOutput struct{}

func (s *service) CheckIn(ctx context.Context, input *CheckInInput) (*CheckInOutput, error) {
	panic("implement me")
}

type CheckOutInput struct{}
type CheckOutOutput struct{}

func (s *service) CheckOut(ctx context.Context, input *CheckOutInput) (*CheckOutOutput, error) {
	panic("implement me")
}
