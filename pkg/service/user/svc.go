package user

import (
	"pokemon/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo iface.IRepository
}

type Params struct {
	fx.In

	Repo iface.IRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.UserService {
	return &service{
		repo: p.Repo,
	}
}
