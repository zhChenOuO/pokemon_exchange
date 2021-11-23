package spot_order

import (
	"pokemon/pkg/iface"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo iface.IRepository
	db   *gorm.DB
}

type Params struct {
	fx.In

	Repo iface.IRepository
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

func (s *service) GetDB() *gorm.DB {
	return s.db
}
