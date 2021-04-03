package user

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
	"reflect"

	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
)

// GetUser 取得User的資訊
func (repo *repository) GetUser(ctx context.Context, tx *gorm.DB, opt option.UserWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.User, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallet model.User
	err := tx.Table(wallet.TableName()).Scopes(opt.Where).Take(&wallet).Error
	if err != nil {
		return wallet, errors.ConvertPostgresError(err)
	}
	return wallet, nil
}

// CreateUser 建立User
func (repo *repository) CreateUser(ctx context.Context, tx *gorm.DB, data *model.User, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// CreateUser 建立User
func (repo *repository) CreateUsers(ctx context.Context, tx *gorm.DB, data *[]model.User, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertPostgresError(err)
}

// ListUsers 列出User
func (repo *repository) ListUsers(ctx context.Context, tx *gorm.DB, opt option.UserWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.User, int64, error) {
	if tx == nil {
		tx = repo.readDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	var wallets []model.User
	var total int64
	db := tx.Table(model.User{}.TableName()).Scopes(opt.Where)
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, errors.Wrapf(errors.ErrInternalError, "database: ListUser err: %s", err.Error())
	}
	err = db.Scopes(opt.Pagination.LimitAndOffset, opt.Sorting.Sort).Find(&wallets).Error
	if err != nil {
		return nil, total, errors.Wrapf(errors.ErrInternalError, "database: ListUser err: %s", err.Error())
	}
	return wallets, total, nil
}

// UpdateUser 更新User
func (repo *repository) UpdateUser(ctx context.Context, tx *gorm.DB, opt option.UserUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.WhereOpts, option.UserWhereOption{}) {
		return errors.Wrap(errors.ErrInternalError, "database: UpdateUser err: where condition can't empty")
	}
	err := tx.Table(model.User{}.TableName()).Scopes(opt.Update).Error
	if err != nil {
		return errors.Wrapf(errors.ErrInternalError, "database: UpdateUser err: %s", err.Error())
	}

	return nil
}

// DeleteUser 刪除User
func (repo *repository) DeleteUser(ctx context.Context, tx *gorm.DB, opt option.UserWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB
	}
	tx = tx.WithContext(ctx).Scopes(scopes...)
	if reflect.DeepEqual(opt.User, model.User{}) {
		return errors.Wrap(errors.ErrInvalidInput, "database: DeleteUser err: UserWhereOption is empty")
	}
	err := tx.Scopes(opt.Where).Delete(&model.User{}).Error
	if err != nil {
		return errors.Wrapf(errors.ErrInternalError, "database: DeleteUser err: %s", err.Error())
	}
	return nil
}
