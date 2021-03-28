package card

import (
	"pokemon/internal/pkg/iface"

	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type repository struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
}

type Params struct {
	fx.In

	DBConns *db.Connections
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.CardRepo {
	return &repository{
		readDB:  p.DBConns.ReadDB,
		writeDB: p.DBConns.WriteDB,
	}
}
