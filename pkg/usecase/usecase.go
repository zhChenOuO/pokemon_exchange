package usecase

import (
	"pokemon/pkg/usecase/matching"

	"go.uber.org/fx"
)

var Module = fx.Options(
	matching.Module,
)
