package restful

import (
	"net/http"
	"pokemon/internal/claims"
	"pokemon/pkg/delivery/restful/view"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"

	"github.com/labstack/echo/v4"
	"gitlab.com/howmay/gopher/common"
)

func (h *handler) ListMySpotOrder(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req view.ListMySpotOrderReq
	)

	claims, err := claims.GetClaims(c)
	if err != nil {
		return err
	}

	if err := req.BindAndVerify(c); err != nil {
		return err
	}

	page := common.Pagination{
		Page:    req.Page,
		PerPage: req.PerPage,
	}
	page.CheckOrSetDefault(50)
	spotOrders, total, err := h.spotOrderSvc.ListSpotOrders(ctx, &option.SpotOrderWhereOption{
		SpotOrder: model.SpotOrder{
			UserID: claims.ID,
		},
		Pagination: page,
		Sorting: common.Sorting{
			SortField: "id",
			Type:      common.SortingOrderType_DESC,
		},
	})
	if err != nil {
		return err
	}

	var resp view.ListMySpotOrderResp
	resp.Meta.Total = total
	resp.SpotOrders = spotOrders

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) ListMyTradeOrder(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req view.ListMyTradeOrderReq
	)

	claims, err := claims.GetClaims(c)
	if err != nil {
		return err
	}

	if err := req.BindAndVerify(c); err != nil {
		return err
	}

	page := common.Pagination{
		Page:    req.Page,
		PerPage: req.PerPage,
	}
	page.CheckOrSetDefault(50)
	tradeOrders, total, err := h.tradeOrderSvc.ListTradeOrders(ctx, &option.TradeOrderWhereOption{
		UserID:     claims.ID,
		Pagination: page,
		Sorting: common.Sorting{
			SortField: "id",
			Type:      common.SortingOrderType_DESC,
		},
	})
	if err != nil {
		return err
	}

	var resp view.ListMyTradeOrderResp
	resp.Meta.Total = total
	resp.TradeOrders = tradeOrders

	return c.JSON(http.StatusOK, resp)
}
