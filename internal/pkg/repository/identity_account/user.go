package identity_account

import (
	"context"
	"pokemon/internal/pkg/model"
	"pokemon/internal/pkg/model/option"
	"reflect"

	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
)

// GetIdentityAccount 取得IdentityAccount的資訊
func (repo *repository) GetIdentityAccount(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.IdentityAccount, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallet model.IdentityAccount
	err := tx.Table(wallet.TableName()).Scopes(opt.Where).Take(&wallet).Error
	if err != nil {
		return wallet, errors.ConvertPostgresError(err)
	}
	return wallet, nil
}

// CreateIdentityAccount 建立IdentityAccount
func (repo *repository) CreateIdentityAccount(ctx context.Context, tx *gorm.DB, data *model.IdentityAccount, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// CreateIdentityAccount 建立IdentityAccount
func (repo *repository) CreateIdentityAccounts(ctx context.Context, tx *gorm.DB, data *[]model.IdentityAccount, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// ListIdentityAccounts 列出IdentityAccount
func (repo *repository) ListIdentityAccounts(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.IdentityAccount, int64, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallets []model.IdentityAccount
	var total int64
	db := tx.Table(model.IdentityAccount{}.TableName()).Scopes(opt.Where)
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, errors.Wrapf(errors.ErrInternalError, "database: ListIdentityAccount err: %s", err.Error())
	}
	err = db.Scopes(opt.Pagination.LimitAndOffset, opt.Sorting.Sort).Find(&wallets).Error
	if err != nil {
		return nil, total, errors.ConvertPostgresError(err)
	}
	return wallets, total, nil
}

// UpdateIdentityAccount 更新IdentityAccount
func (repo *repository) UpdateIdentityAccount(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.WhereOpts, option.IdentityAccountWhereOption{}) {
		return errors.Wrap(errors.ErrInternalError, "database: UpdateIdentityAccount err: where condition can't empty")
	}
	err := tx.Table(model.IdentityAccount{}.TableName()).Scopes(opt.Update).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}

	return nil
}

// DeleteIdentityAccount 刪除IdentityAccount
func (repo *repository) DeleteIdentityAccount(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.IdentityAccount, model.IdentityAccount{}) {
		return errors.Wrap(errors.ErrInvalidInput, "database: DeleteIdentityAccount err: IdentityAccountWhereOption is empty")
	}
	err := tx.Scopes(opt.Where).Delete(&model.IdentityAccount{}).Error
	if err != nil {
		return errors.ConvertPostgresError(err)
	}
	return nil
}
