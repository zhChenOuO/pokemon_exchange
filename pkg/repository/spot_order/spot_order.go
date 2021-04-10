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
	var so []model.SpotOrder
	var total int64
	db := tx.Table(model.SpotOrder{}.TableName()).Scopes(opt.Where)
	if !opt.WithoutCount {
		err := db.Count(&total).Error
		if err != nil {
			return nil, total, errors.ConvertPostgresError(err)
		}
	}
	err := db.Scopes(opt.Pagination.LimitAndOffset, opt.Sorting.Sort).Find(&so).Error
	if err != nil {
		return nil, total, errors.ConvertPostgresError(err)
	}
	return so, total, nil
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
		err = repo.readDB.Raw(`
		SELECT 
			spot_orders.id AS id,
			spot_orders.uuid AS uuid,
			spot_orders.card_id AS card_id,
			spot_orders.user_id AS user_id,
			spot_orders.status AS status,
			spot_orders.trade_side AS trade_side,
			spot_orders.card_quantity - COALESCE(maker_trade.quantity,0) - COALESCE(taker_trade.quantity,0) AS card_quantity,
			spot_orders.expected_amount AS expected_amount,
			spot_orders.created_at AS created_at,
			spot_orders.updated_at AS updated_at
		FROM "spot_orders" 
			LEFT JOIN trade_orders AS maker_trade ON maker_trade.maker_order_id = spot_orders.id 
			LEFT JOIN trade_orders AS taker_trade ON taker_trade.taker_order_id = spot_orders.id 
		WHERE "spot_orders"."status" = '2'
		`).Scan(&sos).Error
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
