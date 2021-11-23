package trade_order

import (
	"pokemon/pkg/iface"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo iface.TradeOrderRepo
	db   *gorm.DB
}

type Params struct {
	fx.In

	Repo iface.TradeOrderRepo
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.TradeOrderService {
	return &service{
		repo: p.Repo,
	}
}

func (s *service) GetDB() *gorm.DB {
	return s.db
}
