package service

import (
	"context"
	"github.com/falmar/otel-trivago/internal/reservations/reservationrepo"
	"github.com/falmar/otel-trivago/internal/rooms/roomrepo"
)

var _ Service = (*service)(nil)

type Service interface {
	List(ctx context.Context, input *ListInput) (*ListOutput, error)
	Create(ctx context.Context, input *CreateInput) (*CreateOutput, error)
}

type service struct {
	resvRepo reservationrepo.Repository
	roomRepo roomrepo.Repository
}

type Config struct {
	ResvRepo reservationrepo.Repository
	RoomRepo roomrepo.Repository
}

func NewService(cfg *Config) Service {
	return &service{
		resvRepo: cfg.ResvRepo,
		roomRepo: cfg.RoomRepo,
	}
}
