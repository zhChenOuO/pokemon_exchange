package option

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/common"

	"gorm.io/gorm"
)

type TradeOrderWhereOption struct {
	TradeOrder model.TradeOrder  `json:"trade_order"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *TradeOrderWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.TradeOrder)
	
	return db
}

type TradeOrderUpdateOption struct {
	WhereOpts TradeOrderWhereOption
	UpdateCol TradeOrderUpdateColumn
}

type TradeOrderUpdateColumn struct{}

func (opts *TradeOrderUpdateOption) Update(db *gorm.DB) *gorm.DB {
	db = db.Scopes(opts.WhereOpts.Where)
	db = db.Updates(opts.UpdateCol)

	return db
}
