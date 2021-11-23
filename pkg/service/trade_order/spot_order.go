package trade_order

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// GetTradeOrder 取得TradeOrder的資訊
func (s *service) GetTradeOrder(ctx context.Context, opt option.TradeOrderWhereOption) (model.TradeOrder, error) {
	return s.repo.GetTradeOrder(ctx, nil, opt)
}

// CreateTradeOrder 建立TradeOrder
func (s *service) CreateTradeOrder(ctx context.Context, data *model.TradeOrder) error {
	return s.repo.CreateTradeOrder(ctx, nil, data)
}

// ListTradeOrders 列出TradeOrder
func (s *service) ListTradeOrders(ctx context.Context, opt option.TradeOrderWhereOption) ([]model.TradeOrder, int64, error) {
	so, total, err := s.repo.ListTradeOrders(ctx, nil, opt)
	if err != nil {
		return nil, 0, err
	}
	return so, total, nil
}

// UpdateTradeOrder 更新TradeOrder
func (s *service) UpdateTradeOrder(ctx context.Context, opt option.TradeOrderUpdateOption) error {
	err := s.repo.UpdateTradeOrder(ctx, nil, opt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTradeOrder 刪除TradeOrder
func (s *service) DeleteTradeOrder(ctx context.Context, opt option.TradeOrderWhereOption) error {
	return s.repo.DeleteTradeOrder(ctx, nil, opt)
}
