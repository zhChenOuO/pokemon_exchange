package spot_order

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
	"reflect"

	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetSpotOrder 取得SpotOrder的資訊
func (repo *repository) GetSpotOrder(ctx context.Context, tx *gorm.DB, opt option.SpotOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.SpotOrder, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallet model.SpotOrder
	err := tx.Table(wallet.TableName()).Scopes(opt.Where).Take(&wallet).Error
	if err != nil {
		return wallet, errors.ConvertPostgresError(err)
	}
	return wallet, nil
}

// CreateSpotOrder 建立SpotOrder
func (repo *repository) CreateSpotOrder(ctx context.Context, tx *gorm.DB, data *model.SpotOrder, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}

	if err := data.VerifyCreateSpotOrder(); err != nil {
		return err
	}

	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// CreateSpotOrder 建立SpotOrder
func (repo *repository) CreateSpotOrders(ctx context.Context, tx *gorm.DB, data *[]model.SpotOrder, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// ListSpotOrders 列出SpotOrder
func (repo *repository) ListSpotOrders(ctx context.Context, tx *gorm.DB, opt option.SpotOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.SpotOrder, int64, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallets []model.SpotOrder
	var total int64
	db := tx.Table(model.SpotOrder{}.TableName()).Scopes(opt.Where)
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, errors.ConvertPostgresError(err)
	}
	err = db.Scopes(opt.Pagination.LimitAndOffset, opt.Sorting.Sort).Find(&wallets).Error
	if err != nil {
		return nil, total, errors.ConvertPostgresError(err)
	}
	return wallets, total, nil
}

// UpdateSpotOrder 更新SpotOrder
func (repo *repository) UpdateSpotOrder(ctx context.Context, tx *gorm.DB, opt option.SpotOrderUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.WhereOpts, option.SpotOrderWhereOption{}) {
		return errors.Wrap(errors.ErrInternalError, "database: UpdateSpotOrder err: where condition can't empty")
	}
	err := tx.Table(model.SpotOrder{}.TableName()).Scopes(opt.WhereOpts.Where).Updates(opt.UpdateCol).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}

	return nil
}

// DeleteSpotOrder 刪除SpotOrder
func (repo *repository) DeleteSpotOrder(ctx context.Context, tx *gorm.DB, opt option.SpotOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.SpotOrder, model.SpotOrder{}) {
		return errors.Wrap(errors.ErrInvalidInput, "database: DeleteSpotOrder err: SpotOrderWhereOption is empty")
	}
	err := tx.Scopes(opt.Where).Delete(&model.SpotOrder{}).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}
	return nil
}

func (repo *repository) ListSpotOrdersWithLock(ctx context.Context) ([]model.SpotOrder, error) {
	var (
		sos   []model.SpotOrder = make([]model.SpotOrder, 0)
		txErr error
	)
	txErr = repo.writeDB.Transaction(func(tx *gorm.DB) error {
		var cards []model.Card
		cardTx := tx.Clauses(clause.Locking{Strength: "UPDATE"})
		err := cardTx.Table("cards").Find(&cards).Error
		if err != nil {
			return err
		}
		sos, _, err = repo.ListSpotOrders(ctx, tx, option.SpotOrderWhereOption{})
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return sos, txErr
	}
	return sos, nil
}
