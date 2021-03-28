package identity_account

import (
	"context"
	"pokemon/internal/pkg/model"
	"pokemon/internal/pkg/model/option"
)

// GetIdentityAccount 取得IdentityAccount的資訊
func (s *service) GetIdentityAccount(ctx context.Context, opt option.IdentityAccountWhereOption) (model.IdentityAccount, error) {
	return s.repo.GetIdentityAccount(ctx, nil, opt)
}

// CreateIdentityAccount 建立IdentityAccount
func (s *service) CreateIdentityAccount(ctx context.Context, data *model.IdentityAccount) error {
	return s.repo.CreateIdentityAccount(ctx, nil, data)
}

// ListIdentityAccounts 列出IdentityAccount
func (s *service) ListIdentityAccounts(ctx context.Context, opt option.IdentityAccountWhereOption) ([]model.IdentityAccount, int64, error) {
	so, total, err := s.repo.ListIdentityAccounts(ctx, nil, opt)
	if err != nil {
		return nil, 0, err
	}
	return so, total, nil
}

// UpdateIdentityAccount 更新IdentityAccount
func (s *service) UpdateIdentityAccount(ctx context.Context, opt option.IdentityAccountUpdateOption) error {
	err := s.repo.UpdateIdentityAccount(ctx, nil, opt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteIdentityAccount 刪除IdentityAccount
func (s *service) DeleteIdentityAccount(ctx context.Context, opt option.IdentityAccountWhereOption) error {
	return s.repo.DeleteIdentityAccount(ctx, nil, opt)
}
