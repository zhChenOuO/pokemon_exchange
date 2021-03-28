package option

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/common"

	"gorm.io/gorm"
)

type UserWhereOption struct {
	User       model.User        `json:"user"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *UserWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.User)

	return db
}

type UserUpdateOption struct {
	WhereOpts UserWhereOption
	UpdateCol UserUpdateColumn
}

type UserUpdateColumn struct{}

func (opts *UserUpdateOption) Update(db *gorm.DB) *gorm.DB {
	db = db.Scopes(opts.WhereOpts.Where)
	db = db.Updates(opts.UpdateCol)
	return db
}
