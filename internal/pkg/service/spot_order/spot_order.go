package spot_order

import (
	"context"
	"pokemon/internal/pkg/model"
	"pokemon/internal/pkg/model/option"
)

// GetSpotOrder 取得SpotOrder的資訊
func (s *service) GetSpotOrder(ctx context.Context, opt option.SpotOrderWhereOption) (model.SpotOrder, error) {
	return s.repo.GetSpotOrder(ctx, nil, opt)
}

// CreateSpotOrder 建立SpotOrder
func (s *service) CreateSpotOrder(ctx context.Context, data *model.SpotOrder) error {
	return s.repo.CreateSpotOrder(ctx, nil, data)
}

// ListSpotOrders 列出SpotOrder
func (s *service) ListSpotOrders(ctx context.Context, opt option.SpotOrderWhereOption) ([]model.SpotOrder, int64, error) {
	so, total, err := s.repo.ListSpotOrders(ctx, nil, opt)
	if err != nil {
		return nil, 0, err
	}
	return so, total, nil
}

// UpdateSpotOrder 更新SpotOrder
func (s *service) UpdateSpotOrder(ctx context.Context, opt option.SpotOrderUpdateOption) error {
	err := s.repo.UpdateSpotOrder(ctx, nil, opt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpotOrder 刪除SpotOrder
func (s *service) DeleteSpotOrder(ctx context.Context, opt option.SpotOrderWhereOption) error {
	return s.repo.DeleteSpotOrder(ctx, nil, opt)
}
