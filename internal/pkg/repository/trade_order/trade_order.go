package trade_order

import (
	"context"
	"pokemon/internal/pkg/model"
	"pokemon/internal/pkg/model/option"
	"reflect"

	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
)

// GetTradeOrder 取得TradeOrder的資訊
func (repo *repository) GetTradeOrder(ctx context.Context, tx *gorm.DB, opt option.TradeOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.TradeOrder, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallet model.TradeOrder
	err := tx.Table(wallet.TableName()).Scopes(opt.Where).Take(&wallet).Error
	if err != nil {
		return wallet, errors.ConvertPostgresError(err)
	}
	return wallet, nil
}

// CreateTradeOrder 建立TradeOrder
func (repo *repository) CreateTradeOrder(ctx context.Context, tx *gorm.DB, data *model.TradeOrder, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// CreateTradeOrder 建立TradeOrder
func (repo *repository) CreateTradeOrders(ctx context.Context, tx *gorm.DB, data *[]model.TradeOrder, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// ListTradeOrders 列出TradeOrder
func (repo *repository) ListTradeOrders(ctx context.Context, tx *gorm.DB, opt option.TradeOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.TradeOrder, int64, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallets []model.TradeOrder
	var total int64
	db := tx.Table(model.TradeOrder{}.TableName()).Scopes(opt.Where)
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, errors.Wrapf(errors.ErrInternalError, "database: ListTradeOrder err: %s", err.Error())
	}
	err = db.Scopes(opt.Pagination.LimitAndOffset, opt.Sorting.Sort).Find(&wallets).Error
	if err != nil {
		return nil, total, errors.Wrapf(errors.ErrInternalError, "database: ListTradeOrder err: %s", err.Error())
	}
	return wallets, total, nil
}

// UpdateTradeOrder 更新TradeOrder
func (repo *repository) UpdateTradeOrder(ctx context.Context, tx *gorm.DB, opt option.TradeOrderUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.WhereOpts, option.TradeOrderWhereOption{}) {
		return errors.Wrap(errors.ErrInternalError, "database: UpdateTradeOrder err: where condition can't empty")
	}
	err := tx.Table(model.TradeOrder{}.TableName()).Scopes(opt.Update).Error
	if err != nil {
		return errors.Wrapf(errors.ErrInternalError, "database: UpdateTradeOrder err: %s", err.Error())
	}

	return nil
}

// DeleteTradeOrder 刪除TradeOrder
func (repo *repository) DeleteTradeOrder(ctx context.Context, tx *gorm.DB, opt option.TradeOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.TradeOrder, model.TradeOrder{}) {
		return errors.Wrap(errors.ErrInvalidInput, "database: DeleteTradeOrder err: TradeOrderWhereOption is empty")
	}
	err := tx.Scopes(opt.Where).Delete(&model.TradeOrder{}).Error
	if err != nil {
		return errors.Wrapf(errors.ErrInternalError, "database: DeleteTradeOrder err: %s", err.Error())
	}
	return nil
}
