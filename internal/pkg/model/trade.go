package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	ID           int64           `json:"id"`
	Turnover     decimal.Decimal `json:"turnover"`
	TakerOrderID int64           `json:"taker_order_id"`
	MakerOrderID int64           `json:"maker_order_id"`
	CreatedAt    time.Time       `json:"created_at"`
}

func (t *Trade) TableName() string {
	return "trades"
}
