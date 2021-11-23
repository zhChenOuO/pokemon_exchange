package card

import (
	"pokemon/pkg/iface"

	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo iface.IRepository

	tx *gorm.DB
}

type Params struct {
	fx.In

	iface.IRepository
	Conns *db.Connections
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.CardService {
	return &service{
		repo: p.IRepository,
		tx:   p.Conns.WriteDB,
	}
}

func (s *service) GetTX() *gorm.DB {
	return s.tx
}
