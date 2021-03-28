package iface

import "github.com/labstack/echo/v4"

type IRestfulHandler interface {
	AuthHandler
	CardHandler
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
