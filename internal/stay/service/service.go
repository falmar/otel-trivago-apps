package service

import "context"

var _ Service = (*service)(nil)

type Config struct{}

func New() Service {
	return &service{}
}

type Service interface {
	ListStays(ctx context.Context, input *ListStaysInput) (*ListStaysOutput, error)
	CreateStay(ctx context.Context, input *CreateStayInput) (*CreateStayOutput, error)
	UpdateStay(ctx context.Context, input *UpdateStayInput) (*UpdateStayOutput, error)
}

type service struct{}

type ListStaysInput struct{}
type ListStaysOutput struct{}

func (s *service) ListStays(ctx context.Context, input *ListStaysInput) (*ListStaysOutput, error) {
	panic("implement me")
}

type CreateStayInput struct{}
type CreateStayOutput struct{}

func (s *service) CreateStay(ctx context.Context, input *CreateStayInput) (*CreateStayOutput, error) {
	panic("implement me")
}

type UpdateStayInput struct{}
type UpdateStayOutput struct{}

func (s *service) UpdateStay(ctx context.Context, input *UpdateStayInput) (*UpdateStayOutput, error) {
	panic("implement me")
}
