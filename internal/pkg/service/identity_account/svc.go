package identity_account

import (
	"pokemon/internal/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo iface.IdentityAccountRepository
}

type Params struct {
	fx.In

	Repo iface.IdentityAccountRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.IdentityAccountService {
	return &service{
		repo: p.Repo,
	}
}
