package service

import (
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo iface.IRepository
}

type Params struct {
	fx.In

	IRepo iface.IRepository
}

var Model = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.IServices {
	return &service{
		repo: p.IRepo,
	}
}
