package restful

import (
	"pokemon/internal/pkg/delivery/restful/view"
	"pokemon/internal/pkg/iface"
	"pokemon/pkg/claims"

	"github.com/labstack/echo/v4"
)

// setSpotOrderRoutes ...
func setSpotOrderRoutes(group *echo.Group, e echo.MiddlewareFunc, h iface.IRestfulHandler) {
	spotOrder := group.Group("/spotOrder", e)
	spotOrder.GET("", h.ListSpotOrders)
	spotOrder.GET("/:id", h.GetSpotOrder)
	spotOrder.POST("", h.CreateSpotOrder)
	spotOrder.PUT("/:id", h.UpdateSpotOrder)
}

func (h *handler) CreateSpotOrder(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req view.SpotOrderReq
	)

	claims, err := claims.GetClaims(c)
	if err != nil {
		return err
	}

	if err := req.BindAndVerify(c); err != nil {
		return err
	}

	so := req.ConvertToSpotOrder(claims)
	if err := h.spotOrderSvc.CreateSpotOrder(ctx, &so); err != nil {
		return err
	}
	return nil
}

func (h *handler) GetSpotOrder(c echo.Context) error {
	return nil
}

func (h *handler) ListSpotOrders(c echo.Context) error {
	return nil
}

func (h *handler) UpdateSpotOrder(c echo.Context) error {
	return nil
}
