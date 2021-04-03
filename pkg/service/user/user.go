package user

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// GetUser 取得User的資訊
func (s *service) GetUser(ctx context.Context, opt option.UserWhereOption) (model.User, error) {
	return s.repo.GetUser(ctx, nil, opt)
}

// CreateUser 建立User
func (s *service) CreateUser(ctx context.Context, data *model.User) error {
	return s.repo.CreateUser(ctx, nil, data)
}

// ListUsers 列出User
func (s *service) ListUsers(ctx context.Context, opt option.UserWhereOption) ([]model.User, int64, error) {
	so, total, err := s.repo.ListUsers(ctx, nil, opt)
	if err != nil {
		return nil, 0, err
	}
	return so, total, nil
}

// UpdateUser 更新User
func (s *service) UpdateUser(ctx context.Context, opt option.UserUpdateOption) error {
	err := s.repo.UpdateUser(ctx, nil, opt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser 刪除User
func (s *service) DeleteUser(ctx context.Context, opt option.UserWhereOption) error {
	return s.repo.DeleteUser(ctx, nil, opt)
}
