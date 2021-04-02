package option

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/common"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SpotOrderWhereOption struct {
	SpotOrder              model.SpotOrder   `json:"spot_order"`
	Pagination             common.Pagination `json:"pagination"`
	BaseWhere              common.BaseWhere  `json:"base_where"`
	Sorting                common.Sorting    `json:"sorting"`
	ExpectedAmountMoreThan decimal.Decimal   `gorm:"-"`
	ExpectedAmountLessThan decimal.Decimal   `gorm:"-"`
}

func (where *SpotOrderWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.SpotOrder)
	db = db.Scopes(where.Sorting.Sort)

	if !where.ExpectedAmountMoreThan.IsZero() {
		db = db.Where("expected_amount >= ?", where.ExpectedAmountMoreThan)
	}

	if !where.ExpectedAmountLessThan.IsZero() {
		db = db.Where("expected_amount <= ?", where.ExpectedAmountLessThan)
	}

	return db
}

type SpotOrderUpdateOption struct {
	WhereOpts SpotOrderWhereOption
	UpdateCol SpotOrderUpdateColumn
}

type SpotOrderUpdateColumn struct {
	Status         model.OrderStatus `json:"status" gorm:"status"`
	Type           model.OrderType   `json:"type" gorm:"type"`
	ActuallyAmount decimal.Decimal   `json:"actually_amount" gorm:"actually_amount"`
}

func (opts *SpotOrderUpdateOption) Update(db *gorm.DB) *gorm.DB {
	db = db.Scopes(opts.WhereOpts.Where)
	db = db.Updates(opts.UpdateCol)
	return db
}
