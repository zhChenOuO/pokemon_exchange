package service

import (
	"pokemon/pkg/service/card"
	"pokemon/pkg/service/identity_account"
	"pokemon/pkg/service/spot_order"
	"pokemon/pkg/service/trade_order"
	"pokemon/pkg/service/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	identity_account.Module,
	card.Module,
	spot_order.Module,
	user.Module,
	trade_order.Module,
)
