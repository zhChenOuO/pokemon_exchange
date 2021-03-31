package spot_order

import (
	"pokemon/internal/pkg/iface"

	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo iface.SpotOrderRepo
	db   *gorm.DB
}

type Params struct {
	fx.In

	Repo  iface.SpotOrderRepo
	conns *db.Connections
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.SpotOrderService {
	return &service{
		repo: p.Repo,
		db:   p.conns.WriteDB,
	}
}

func (s *service) GetDB() *gorm.DB {
	return s.db
}
