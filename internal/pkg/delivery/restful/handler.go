package restful

import (
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

type handler struct {
	svc iface.IServices
}

type Params struct {
	fx.In

	ISvc iface.IServices
}

var Model = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(
		setRoutes,
	),
)

func New(p Params) iface.IRestfulHandler {
	return &handler{
		svc: p.ISvc,
	}
}
