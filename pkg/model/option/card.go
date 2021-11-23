package option

import (
	"pokemon/pkg/model"
	"reflect"

	"gitlab.com/howmay/gopher/common"
	"gorm.io/gorm"
)

type CardWhereOption struct {
	Card       model.Card        `json:"card"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *CardWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.Card)

	return db
}

func (where *CardWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *CardWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *CardWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.Card, model.Card{})
}

func (where *CardWhereOption) TableName() string {
	return where.Card.TableName()
}

func (where *CardWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *CardWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type CardUpdateColumn struct{}

func (cols *CardUpdateColumn) Columns() interface{} {
	return cols
}
