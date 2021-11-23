package spot_order

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// GetSpotOrder 取得SpotOrder的資訊
func (s *service) GetSpotOrder(ctx context.Context, opt *option.SpotOrderWhereOption) (model.SpotOrder, error) {
	var (
		result model.SpotOrder
	)
	if err := s.repo.Get(ctx, nil, &result, opt); err != nil {
		return result, err
	}
	return result, nil
}

// CreateSpotOrder 建立SpotOrder
func (s *service) CreateSpotOrder(ctx context.Context, data *model.SpotOrder) error {
	return s.repo.Create(ctx, nil, data)
}

// ListSpotOrders 列出SpotOrder
func (s *service) ListSpotOrders(ctx context.Context, opt *option.SpotOrderWhereOption) ([]model.SpotOrder, int64, error) {
	var (
		results []model.SpotOrder
	)
	total, err := s.repo.List(ctx, nil, &results, opt)
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

// UpdateSpotOrder 更新SpotOrder
func (s *service) UpdateSpotOrder(ctx context.Context, opt *option.SpotOrderWhereOption, col *option.SpotOrderUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpotOrder 刪除SpotOrder
func (s *service) DeleteSpotOrder(ctx context.Context, opt *option.SpotOrderWhereOption) error {
	return s.repo.Delete(ctx, nil, &model.SpotOrder{}, opt)
}
