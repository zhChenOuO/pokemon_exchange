package restful

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCard(c echo.Context) error {
	return nil
}

func (h *handler) CreateCard(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (h *handler) ListCards(c echo.Context) error {
	return nil
}

func (h *handler) UpdateCard(c echo.Context) error {
	return nil
}
