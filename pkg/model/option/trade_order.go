package option

import (
	"pokemon/pkg/model"
	"reflect"

	"gitlab.com/howmay/gopher/common"
	"gorm.io/gorm"
)

type TradeOrderWhereOption struct {
	TradeOrder model.TradeOrder  `json:"trade_order"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`

	UserID uint64 `json:"user_id"`
}

func (where *TradeOrderWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.TradeOrder)

	if where.UserID != 0 {
		db = db.Where("")
	}
	return db
}

func (where *TradeOrderWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *TradeOrderWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *TradeOrderWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.TradeOrder, model.TradeOrder{})
}

func (where *TradeOrderWhereOption) TableName() string {
	return where.TradeOrder.TableName()
}

func (where *TradeOrderWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *TradeOrderWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type TradeOrderUpdateColumn struct{}

func (cols *TradeOrderUpdateColumn) Columns() interface{} {
	return cols
}
