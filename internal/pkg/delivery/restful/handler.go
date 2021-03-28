package restful

import (
	"pokemon/configuration"
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

type handler struct {
	appCfg       *configuration.App
	authSvc      iface.IdentityAccountService
	userSvc      iface.UserService
	cardSvc      iface.CardService
	spotOrderSvc iface.SpotOrderService
}

type Params struct {
	fx.In

	AppCfg       *configuration.App
	AuthSvc      iface.IdentityAccountService
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
		appCfg:       p.AppCfg,
		authSvc:      p.AuthSvc,
		userSvc:      p.UserSvc,
		spotOrderSvc: p.SpotOrderSvc,
		cardSvc:      p.CardSvc,
	}
}
