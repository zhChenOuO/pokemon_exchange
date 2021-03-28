package card

import (
	"pokemon/internal/pkg/iface"

	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo iface.CardRepo

	tx *gorm.DB
}

type Params struct {
	fx.In

	CardRepo iface.CardRepo
	Conns    *db.Connections
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.CardService {
	return &service{
		repo: p.CardRepo,
		tx:   p.Conns.WriteDB,
	}
}

func (s *service) GetTX() *gorm.DB {
	return s.tx
}
