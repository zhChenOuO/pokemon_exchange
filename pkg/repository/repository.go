package repository

import (
	"pokemon/pkg/iface"

	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type repository struct {
	*db.Connections
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type Params struct {
	fx.In

	DBConns   *db.Connections
}

func New(p Params) (iface.IRepository, error) {
	repo := &repository{
		Connections: p.DBConns,
	}
	return repo, nil
}

func (repo *repository) GetDB() *gorm.DB {
	return repo.Connections.WriteDB
}
