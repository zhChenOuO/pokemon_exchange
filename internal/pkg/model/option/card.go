package option

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/common"

	"gorm.io/gorm"
)

type CardWhereOption struct {
	Card       model.Card        `json:"card"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *CardWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.Card)

	return db
}

type CardUpdateOption struct {
	WhereOpts CardWhereOption
	UpdateCol CardUpdateColumn
}

type CardUpdateColumn struct{}
