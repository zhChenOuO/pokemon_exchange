package option

import (
	"pokemon/pkg/model"
	"reflect"

	"github.com/shopspring/decimal"
	"gitlab.com/howmay/gopher/common"
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

func (where *SpotOrderWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *SpotOrderWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *SpotOrderWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.SpotOrder, model.SpotOrder{})
}

func (where *SpotOrderWhereOption) TableName() string {
	return where.SpotOrder.TableName()
}

func (where *SpotOrderWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *SpotOrderWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type SpotOrderUpdateColumn struct {
	Status         model.OrderStatus `json:"status" gorm:"status"`
	Type           model.OrderType   `json:"type" gorm:"type"`
	ActuallyAmount decimal.Decimal   `json:"actually_amount" gorm:"actually_amount"`
}

func (cols *SpotOrderUpdateColumn) Columns() interface{} {
	return cols
}
