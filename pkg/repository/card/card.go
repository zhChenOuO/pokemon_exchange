package card

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
	"reflect"

	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
)

// GetCard 取得Card的資訊
func (repo *repository) GetCard(ctx context.Context, tx *gorm.DB, opt option.CardWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.Card, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.Scopes(scopes...)
	var card model.Card
	err := tx.Table(model.Card{}.TableName()).Scopes(opt.Where).Take(&card).Error
	if err != nil {
		return card, errors.ConvertPostgresError(err)
	}
	return card, nil
}

// CreateCard 建立Card
func (repo *repository) CreateCard(ctx context.Context, tx *gorm.DB, data *model.Card, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.Scopes(scopes...)
	err := tx.Create(data).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}
	return nil
}

// ListCards 列出Card
func (repo *repository) ListCards(ctx context.Context, tx *gorm.DB, opt option.CardWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.Card, int64, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.Scopes(scopes...)
	var cards []model.Card
	var total int64
	db := tx.Table(model.Card{}.TableName()).Scopes(opt.Where)
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, errors.ConvertPostgresError(err)
	}
	err = db.Scopes(opt.Pagination.LimitAndOffset).Find(&cards).Error
	if err != nil {
		return nil, total, errors.ConvertPostgresError(err)
	}
	return cards, total, nil
}

// UpdateCard 更新Card
func (repo *repository) UpdateCard(ctx context.Context, tx *gorm.DB, opt option.CardUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.Scopes(scopes...)
	if reflect.DeepEqual(opt.WhereOpts, option.CardWhereOption{}) {
		return errors.Wrap(errors.ErrInternalError, "database: UpdateCard err: where condition can't empty")
	}
	err := tx.Table(model.Card{}.TableName()).Scopes(opt.WhereOpts.Where).Updates(opt.UpdateCol).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}
	return nil
}

// DeleteCard 刪除Card
func (repo *repository) DeleteCard(ctx context.Context, tx *gorm.DB, opt option.CardWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.Scopes(scopes...)
	if reflect.DeepEqual(opt.Card, model.Card{}) {
		return errors.Wrap(errors.ErrInvalidInput, "database: DeleteCard err: WhereCardCondition is empty")
	}
	err := tx.Scopes(opt.Where).Delete(&model.Card{}).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}
	return nil
}
