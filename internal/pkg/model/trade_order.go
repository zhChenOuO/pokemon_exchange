package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type TradeOrder struct {
	ID           uint64          `json:"id"`
	Turnover     decimal.Decimal `json:"turnover"`
	TakerOrderID uint64          `json:"taker_order_id"`
	MakerOrderID uint64          `json:"maker_order_id"`
	CreatedAt    time.Time       `json:"created_at"`
}

func (TradeOrder) TableName() string {
	return "trade_orders"
}

func (t *TradeOrder) InitTradeOrder(makerOrder, takerOrder *SpotOrder) {
	t.MakerOrderID = makerOrder.ID
	t.TakerOrderID = takerOrder.ID
	t.Turnover = makerOrder.ExpectedAmount

	makerOrder.Type = OrderTypeMaker
	makerOrder.Status = OrderSuccess

	takerOrder.Type = OrderTypeTaker
	takerOrder.Status = OrderSuccess
}
