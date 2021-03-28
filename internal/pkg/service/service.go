package service

import (
	"pokemon/internal/pkg/service/card"
	"pokemon/internal/pkg/service/spot_order"
	"pokemon/internal/pkg/service/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	card.Module,
	spot_order.Module,
	user.Module,
)
