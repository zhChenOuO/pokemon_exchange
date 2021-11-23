package option

import (
	"pokemon/pkg/model"
	"reflect"

	"gitlab.com/howmay/gopher/common"
	"gorm.io/gorm"
)

// IdentityAccountWhereOption ORM查詢條件
type IdentityAccountWhereOption struct {
	IdentityAccount model.IdentityAccount `json:"identity_account"`
	Pagination      common.Pagination     `json:"pagination"`
	BaseWhere       common.BaseWhere      `json:"base_where"`
	Sorting         common.Sorting        `json:"sorting"`
}

// Where 基礎的查詢條件
func (where *IdentityAccountWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.IdentityAccount)
	db = db.Scopes(where.BaseWhere.Where)
	return db
}

func (where *IdentityAccountWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *IdentityAccountWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *IdentityAccountWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.IdentityAccount, model.IdentityAccount{})
}

func (where *IdentityAccountWhereOption) TableName() string {
	return where.IdentityAccount.TableName()
}

func (where *IdentityAccountWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *IdentityAccountWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

// UpdateIdentityAccountColumn ORM更新欄位
type IdentityAccountUpdateColumn struct {
}

func (cols *IdentityAccountUpdateColumn) Columns() interface{} {
	return cols
}
