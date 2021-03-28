package user

import (
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo iface.UserRepo
}

type Params struct {
	fx.In

	Repo iface.UserRepo
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
