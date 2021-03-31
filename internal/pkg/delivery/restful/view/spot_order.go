package view

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/claims"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/shopspring/decimal"
	"gitlab.com/howmay/gopher/errors"
)

type SpotOrderReq struct {
	CardID         int64                `json:"card_id"`
	TradeSide      model.OrderTradeSide `json:"trade_side"`
	ExpectedAmount decimal.Decimal      `json:"expected_amount"`
}

func (req *SpotOrderReq) BindAndVerify(c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return errors.WithStack(errors.ErrInvalidInput)
	}

	if req.CardID == 0 {
		return errors.NewWithMessage(errors.ErrInvalidInput, "need choose card")
	}
	if req.TradeSide == 0 {
		return errors.NewWithMessage(errors.ErrInvalidInput, "need choose buy or sell")
	}
	if req.ExpectedAmount.IsZero() || req.ExpectedAmount.IsNegative() {
		return errors.NewWithMessage(errors.ErrInvalidInput, "need fill in money")
	}

	return nil
}

func (req *SpotOrderReq) ConvertToSpotOrder(claims *claims.Claims) model.SpotOrder {
	return model.SpotOrder{
		UUID:           xid.New().String(),
		CardID:         req.CardID,
		UserID:         claims.GetID(),
		TradeSide:      req.TradeSide,
		ExpectedAmount: req.ExpectedAmount,
	}
}
