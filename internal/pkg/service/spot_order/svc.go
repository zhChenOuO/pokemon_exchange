package spot_order

import (
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo iface.SpotOrderRepo
}

type Params struct {
	fx.In

	Repo iface.SpotOrderRepo
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.SpotOrderService {
	return &service{
		repo: p.Repo,
	}
}
