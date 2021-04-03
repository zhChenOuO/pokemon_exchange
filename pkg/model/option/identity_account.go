package option

import (
	"pokemon/internal/common"
	"pokemon/pkg/model"

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

// IdentityAccountUpdateOption ORM更新條件
type IdentityAccountUpdateOption struct {
	WhereOpts    IdentityAccountWhereOption
	UpdateColumn IdentityAccountUpdateColumn
}

// UpdateIdentityAccountColumn ORM更新欄位
type IdentityAccountUpdateColumn struct {
}

func (opts *IdentityAccountUpdateOption) Update(db *gorm.DB) *gorm.DB {
	return db
}
