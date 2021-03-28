package restful

import (
	"pokemon/internal/pkg/iface"

	"github.com/labstack/echo/v4"
)

// setRoutes ...
func setRoutes(e *echo.Echo, h iface.IRestfulHandler) {
	rootV1 := e.Group("/apis/v1")

	auth := rootV1.Group("auth")
	auth.POST("/register", h.Register)

	card := rootV1.Group("cards")
	card.GET("/", func(c echo.Context) error {
		return nil
	})
	card.POST("/", func(c echo.Context) error {
		return nil
	})
	card.PUT("/:id", func(c echo.Context) error {
		return nil
	})
	card.DELETE("/:id", func(c echo.Context) error {
		return nil
	})

}
