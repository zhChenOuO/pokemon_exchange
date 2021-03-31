package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// SpotOrder ...
type SpotOrder struct {
	ID             int64           `json:"id"`
	UUID           string          `json:"uuid"`
	CardID         int64           `json:"card_id"`
	UserID         int64           `json:"user_id"`
	Type           OrderType       `json:"type"`
	TradeSide      OrderTradeSide  `json:"trade_side"`
	ExpectedAmount decimal.Decimal `json:"expected_amount"`
	ActuallyAmount decimal.Decimal `json:"actually_amount"`
	CreatedAt      time.Time
}

func (SpotOrder) TableName() string {
	return "spot_orders"
}
