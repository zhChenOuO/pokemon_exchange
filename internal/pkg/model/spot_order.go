package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// SpotOrder ...
type SpotOrder struct {
	ID          int64           `json:"id"`
	UUID        string          `json:"uuid"`
	Type        OrderType       `json:"type"`
	TradeSide   OrderTradeSide  `json:"trade_side"`
	CardID      int64           `json:"card_id"`
	ExpectedUSD decimal.Decimal `json:"expected_usd"`
	ActualUSD   decimal.Decimal `json:"actual_usd"`
	CreatedAt   time.Time
}

func (SpotOrder) TableName() string {
	return "spot_orders"
}
