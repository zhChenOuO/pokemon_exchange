package user

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// GetUser 取得User的資訊
func (s *service) GetUser(ctx context.Context, opt *option.UserWhereOption) (model.User, error) {
	var (
		result model.User
	)
	if err := s.repo.Get(ctx, nil, &result, opt); err != nil {
		return result, err
	}
	return result, nil
}

// CreateUser 建立User
func (s *service) CreateUser(ctx context.Context, data *model.User) error {
	return s.repo.Create(ctx, nil, data)
}

// ListUsers 列出User
func (s *service) ListUsers(ctx context.Context, opt *option.UserWhereOption) ([]model.User, int64, error) {
	var (
		results []model.User
	)
	total, err := s.repo.List(ctx, nil, &results, opt)
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

// UpdateUser 更新User
func (s *service) UpdateUser(ctx context.Context, opt *option.UserWhereOption, col *option.UserUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser 刪除User
func (s *service) DeleteUser(ctx context.Context, opt *option.UserWhereOption) error {
	return s.repo.Delete(ctx, nil, &model.User{}, opt)
}
