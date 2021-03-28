package repository

import (
	"pokemon/internal/pkg/repository/card"
	"pokemon/internal/pkg/repository/spot_order"
	"pokemon/internal/pkg/repository/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	card.Module,
	spot_order.Module,
	user.Module,
)
