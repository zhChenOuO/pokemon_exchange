package model

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/shopspring/decimal"
	"gitlab.com/howmay/gopher/errors"
)

type OrderStatus int8

const (
	OrderInitialization        OrderStatus = 1
	OrderWaitingForMatchmaking OrderStatus = 2
	OrderCancel                OrderStatus = 3
	OrderSuccess               OrderStatus = 4
)

// SpotOrder ...
type SpotOrder struct {
	ID             uint64          `json:"id"`
	UUID           string          `json:"uuid"`
	CardID         uint64          `json:"card_id"`
	UserID         uint64          `json:"user_id"`
	Status         OrderStatus     `json:"status"`
	Type           OrderType       `json:"type"`
	TradeSide      OrderTradeSide  `json:"trade_side"`
	CardQuantity   decimal.Decimal `json:"card_quantity"`
	ExpectedAmount decimal.Decimal `json:"expected_amount"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func (SpotOrder) TableName() string {
	return "spot_orders"
}

func (so *SpotOrder) VerifyCreateSpotOrder() error {
	if so.CardID == 0 {
		return errors.WithMessagef(errors.ErrInvalidInput, "card id is empty")
	}
	if so.UserID == 0 {
		return errors.WithMessagef(errors.ErrInvalidInput, "user id is empty")
	}
	if so.TradeSide == 0 {
		return errors.WithMessagef(errors.ErrInvalidInput, "trade side is empty")
	}
	if so.ExpectedAmount.IsZero() {
		return errors.WithMessagef(errors.ErrInvalidInput, "expected amount is empty")
	}
	if so.CardQuantity.IsZero() {
		return errors.WithMessagef(errors.ErrInvalidInput, "card quantity is empty")
	}

	if !so.CardQuantity.Equal(decimal.NewFromInt(1)) {
		return errors.WithMessagef(errors.ErrInvalidInput, "card quantity allow 1")
	}

	so.UUID = xid.New().String()
	so.Status = OrderWaitingForMatchmaking
	return nil
}

func (so *SpotOrder) RedisLockKey() string {
	return fmt.Sprintf("%s:%d", "trade", so.CardID)
}

func (so *SpotOrder) SetSuccess(t OrderType) {
	so.Status = OrderSuccess
	so.Type = t
}
