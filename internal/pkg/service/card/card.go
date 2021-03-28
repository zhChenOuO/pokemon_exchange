package card

import (
	"context"
	"pokemon/internal/pkg/model"
	"pokemon/internal/pkg/model/option"
)

// GetCard 取得Card的資訊
func (s *service) GetCard(ctx context.Context, opt option.CardWhereOption) (model.Card, error) {
	return s.repo.GetCard(ctx, nil, opt)
}

// CreateCard 建立Card
func (s *service) CreateCard(ctx context.Context, data *model.Card) error {
	return s.repo.CreateCard(ctx, nil, data)
}

// ListCards 列出Card
func (s *service) ListCards(ctx context.Context, opt option.CardWhereOption) ([]model.Card, int64, error) {
	sw, total, err := s.repo.ListCards(ctx, nil, opt)
	if err != nil {
		return nil, 0, err
	}

	return sw, total, nil
}

// UpdateCard 更新Card
func (s *service) UpdateCard(ctx context.Context, opt option.CardUpdateOption) error {
	err := s.repo.UpdateCard(ctx, nil, opt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCard 刪除Card
func (s *service) DeleteCard(ctx context.Context, opt option.CardWhereOption) error {
	return s.repo.DeleteCard(ctx, nil, opt)
}
