package option

import (
	"pokemon/internal/common"
	"pokemon/pkg/model"

	"gorm.io/gorm"
)

type TradeOrderWhereOption struct {
	TradeOrder model.TradeOrder  `json:"trade_order"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`

	MakerOrderIDs []uint64 `gorm:"-"`
	TakerOrderIDs []uint64 `gorm:"-"`
}

func (where *TradeOrderWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.TradeOrder)

	if len(where.MakerOrderIDs) != 0 {
		db = db.Where("maker_order_id IN (?)", where.MakerOrderIDs)
	}

	if len(where.TakerOrderIDs) != 0 {
		db = db.Where("taker_order_id IN (?)", where.TakerOrderIDs)
	}

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
