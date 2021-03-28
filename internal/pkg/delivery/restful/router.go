package restful

import (
	"pokemon/internal/pkg/iface"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// setRoutes ...
func setRoutes(e *echo.Echo, h iface.IRestfulHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	rootV1 := e.Group("/apis/v1")

	auth := rootV1.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	card := rootV1.Group("/cards")
	card.GET("/", h.ListCards)
	card.GET("/:id", h.GetCard)
	card.POST("/", h.CreateCard)
	card.PUT("/:id", h.UpdateCard)

}
