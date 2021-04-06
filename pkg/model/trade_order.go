package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type TradeOrder struct {
	ID           uint64          `json:"id"`
	Turnover     decimal.Decimal `json:"turnover"` // 成交金額
	Quantity     decimal.Decimal `json:"quantity"` // 成交數量
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

	makerOrder.Status = OrderSuccess
	takerOrder.Status = OrderSuccess

	if makerOrder.CardQuantity.GreaterThan(takerOrder.CardQuantity) {
		t.Quantity = takerOrder.CardQuantity
		makerOrder.CardQuantity = makerOrder.CardQuantity.Sub(takerOrder.CardQuantity)
		takerOrder.CardQuantity = decimal.Zero
	} else if makerOrder.CardQuantity.LessThan(takerOrder.CardQuantity) {
		t.Quantity = makerOrder.CardQuantity
		takerOrder.CardQuantity = takerOrder.CardQuantity.Sub(makerOrder.CardQuantity)
		makerOrder.CardQuantity = decimal.Zero
	} else {
		t.Quantity = takerOrder.CardQuantity
		takerOrder.CardQuantity = decimal.Zero
		makerOrder.CardQuantity = decimal.Zero
	}
}
