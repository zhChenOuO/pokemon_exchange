package card

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// GetCard 取得Card的資訊
func (s *service) GetCard(ctx context.Context, opt *option.CardWhereOption) (model.Card, error) {
	var (
		result model.Card
	)
	if err := s.repo.Get(ctx, nil, &result, opt); err != nil {
		return result, err
	}
	return result, nil
}

// CreateCard 建立Card
func (s *service) CreateCard(ctx context.Context, data *model.Card) error {
	return s.repo.Create(ctx, nil, data)
}

// ListCards 列出Card
func (s *service) ListCards(ctx context.Context, opt *option.CardWhereOption) ([]model.Card, int64, error) {
	var (
		results []model.Card
	)
	total, err := s.repo.List(ctx, nil, &results, opt)
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

// UpdateCard 更新Card
func (s *service) UpdateCard(ctx context.Context, opt *option.CardWhereOption, col *option.CardUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCard 刪除Card
func (s *service) DeleteCard(ctx context.Context, opt *option.CardWhereOption) error {
	return s.repo.Delete(ctx, nil, &model.Card{}, opt)
}
