package trade_order

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// GetTradeOrder 取得TradeOrder的資訊
func (s *service) GetTradeOrder(ctx context.Context, opt *option.TradeOrderWhereOption) (model.TradeOrder, error) {
	var (
		result model.TradeOrder
	)
	if err := s.repo.Get(ctx, nil, &result, opt); err != nil {
		return result, err
	}
	return result, nil
}

// CreateTradeOrder 建立TradeOrder
func (s *service) CreateTradeOrder(ctx context.Context, data *model.TradeOrder) error {
	return s.repo.Create(ctx, nil, data)
}

// ListTradeOrders 列出TradeOrder
func (s *service) ListTradeOrders(ctx context.Context, opt *option.TradeOrderWhereOption) ([]model.TradeOrder, int64, error) {
	var (
		results []model.TradeOrder
	)
	total, err := s.repo.List(ctx, nil, &results, opt)
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}

// UpdateTradeOrder 更新TradeOrder
func (s *service) UpdateTradeOrder(ctx context.Context, opt *option.TradeOrderWhereOption, col *option.TradeOrderUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTradeOrder 刪除TradeOrder
func (s *service) DeleteTradeOrder(ctx context.Context, opt *option.TradeOrderWhereOption) error {
	return s.repo.Delete(ctx, nil, &model.TradeOrder{}, opt)
}
