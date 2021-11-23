package option

import (
	"pokemon/pkg/model"
	"reflect"

	"gitlab.com/howmay/gopher/common"
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

func (where *UserWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *UserWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *UserWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.User, model.User{})
}

func (where *UserWhereOption) TableName() string {
	return where.User.TableName()
}

func (where *UserWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *UserWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type UserUpdateColumn struct{}

func (cols *UserUpdateColumn) Columns() interface{} {
	return cols
}
