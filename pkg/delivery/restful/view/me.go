package view

import (
	"pokemon/pkg/model"

	"github.com/labstack/echo/v4"
	"gitlab.com/howmay/gopher/errors"
)

type ListMySpotOrderReq struct {
	Page    int64 `query:"page"`
	PerPage int64 `query:"perPage"`
}

type ListMySpotOrderResp struct {
	SpotOrders []model.SpotOrder `json:"spot_orders"`
	Meta       Meta              `json:"meta"`
}

func (req *ListMySpotOrderReq) BindAndVerify(c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return errors.WithStack(errors.ErrInvalidInput)
	}

	return nil
}

type ListMyTradeOrderReq struct {
	Page    int64 `query:"page"`
	PerPage int64 `query:"perPage"`
}

type ListMyTradeOrderResp struct {
	TradeOrders []model.TradeOrder `json:"trade_orders"`
	Meta        Meta               `json:"meta"`
}

func (req *ListMyTradeOrderReq) BindAndVerify(c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return errors.WithStack(errors.ErrInvalidInput)
	}

	return nil
}
