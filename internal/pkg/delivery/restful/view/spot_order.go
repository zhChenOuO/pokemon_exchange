package view

import (
	"pokemon/internal/pkg/model"
	"pokemon/pkg/claims"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gitlab.com/howmay/gopher/errors"
)

type SpotOrderReq struct {
	CardID         uint64               `json:"card_id"`
	TradeSide      model.OrderTradeSide `json:"trade_side"`
	ExpectedAmount decimal.Decimal      `json:"expected_amount"`
	CardQuantity   decimal.Decimal      `json:"card_quantity"`
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
	if req.CardQuantity.IsZero() || req.CardQuantity.IsNegative() {
		return errors.NewWithMessage(errors.ErrInvalidInput, "need fill in card quantity")
	}

	return nil
}

func (req *SpotOrderReq) ConvertToSpotOrder(claims *claims.Claims) model.SpotOrder {
	return model.SpotOrder{
		CardID:         req.CardID,
		UserID:         claims.GetID(),
		TradeSide:      req.TradeSide,
		ExpectedAmount: req.ExpectedAmount,
		CardQuantity:   req.CardQuantity,
	}
}
