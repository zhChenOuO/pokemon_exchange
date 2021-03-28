package option

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/common"

	"gorm.io/gorm"
)

type SpotOrderWhereOption struct {
	SpotOrder  model.SpotOrder   `json:"spot_order"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *SpotOrderWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.SpotOrder)

	return db
}

type SpotOrderUpdateOption struct {
	WhereOpts SpotOrderWhereOption
	UpdateCol SpotOrderUpdateColumn
}

type SpotOrderUpdateColumn struct{}

func (opts *SpotOrderUpdateOption) Update(db *gorm.DB) *gorm.DB {
	db = db.Scopes(opts.WhereOpts.Where)
	db = db.Updates(opts.UpdateCol)

	return db
}
