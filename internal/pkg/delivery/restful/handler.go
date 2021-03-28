package restful

import (
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

type handler struct {
	userSvc      iface.UserService
	cardSvc      iface.CardService
	spotOrderSvc iface.SpotOrderService
}

type Params struct {
	fx.In

	UserSvc      iface.UserService
	CardSvc      iface.CardService
	SpotOrderSvc iface.SpotOrderService
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(
		setRoutes,
	),
)

func New(p Params) iface.IRestfulHandler {
	return &handler{
		userSvc:      p.UserSvc,
		spotOrderSvc: p.SpotOrderSvc,
		cardSvc:      p.CardSvc,
	}
}
