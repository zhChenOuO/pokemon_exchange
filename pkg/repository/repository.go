package repository

import (
	"pokemon/pkg/repository/card"
	"pokemon/pkg/repository/identity_account"
	"pokemon/pkg/repository/spot_order"
	"pokemon/pkg/repository/trade_order"
	"pokemon/pkg/repository/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	identity_account.Module,
	card.Module,
	spot_order.Module,
	user.Module,
	trade_order.Module,
)
