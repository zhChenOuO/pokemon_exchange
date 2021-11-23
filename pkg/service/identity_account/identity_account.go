package identity_account

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"

	"gitlab.com/howmay/gopher/errors"
)

// GetIdentityAccount 取得IdentityAccount的資訊
func (s *service) GetIdentityAccount(ctx context.Context, opt *option.IdentityAccountWhereOption) (model.IdentityAccount, error) {
	var (
		result model.IdentityAccount
	)
	if err := s.repo.Get(ctx, nil, &result, opt); err != nil {
		return result, err
	}
	return result, nil
}

// CreateIdentityAccount 建立IdentityAccount
func (s *service) CreateIdentityAccount(ctx context.Context, data *model.IdentityAccount) error {
	return s.repo.Create(ctx, nil, data)
}

// ListIdentityAccounts 列出IdentityAccount
func (s *service) ListIdentityAccounts(ctx context.Context, opt *option.IdentityAccountWhereOption) ([]model.IdentityAccount, int64, error) {
	var (
		results []model.IdentityAccount
	)
	total, err := s.repo.List(ctx, nil, &results, opt)
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

// UpdateIdentityAccount 更新IdentityAccount
func (s *service) UpdateIdentityAccount(ctx context.Context, opt *option.IdentityAccountWhereOption, col *option.IdentityAccountUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteIdentityAccount 刪除IdentityAccount
func (s *service) DeleteIdentityAccount(ctx context.Context, opt *option.IdentityAccountWhereOption) error {
	return s.repo.Delete(ctx, nil, &model.IdentityAccount{}, opt)
}

func (s *service) VerifyIdentityAccount(ctx context.Context, data model.IdentityAccount) (model.IdentityAccount, error) {
	var (
		acc model.IdentityAccount
		err error
	)
	acc, err = s.GetIdentityAccount(ctx, &option.IdentityAccountWhereOption{
		IdentityAccount: model.IdentityAccount{
			Name: data.Name,
		},
	})

	switch {
	case errors.Is(err, errors.ErrResourceNotFound):
		return acc, errors.ErrUsernameOrPasswordUnavailable
	case err != nil:
		return acc, err
	case acc.Password.String() != data.Password.String():
		return acc, errors.ErrUsernameOrPasswordUnavailable
	}

	return acc, nil
}
