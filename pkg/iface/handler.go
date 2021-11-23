package iface

import "github.com/labstack/echo/v4"

type IRestfulHandler interface {
	AuthHandler
	CardHandler
	SpotOrderHandler
	MeHandler
}

type AuthHandler interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type CardHandler interface {
	GetCard(ctx echo.Context) error
	ListCards(ctx echo.Context) error
	CreateCard(ctx echo.Context) error
	UpdateCard(ctx echo.Context) error
}

type SpotOrderHandler interface {
	GetSpotOrder(ctx echo.Context) error
	ListSpotOrders(ctx echo.Context) error
	CreateSpotOrder(ctx echo.Context) error
	UpdateSpotOrder(ctx echo.Context) error
}

type MeHandler interface {
	ListMySpotOrder(ctx echo.Context) error
	ListMyTradeOrder(ctx echo.Context) error
}
