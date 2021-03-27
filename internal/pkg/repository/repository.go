package repository

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

var Model = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.IRepository {
	return &repository{
		readDB:  p.DBConns.ReadDB,
		writeDB: p.DBConns.WriteDB,
	}
}
