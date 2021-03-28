package model

type OrderTradeSide int8

const (
	BuySide OrderTradeSide = iota + 1
	SellSide
)

type OrderType int8

const (
	OrderTypeMaker OrderType = iota + 1
	OrderTypeTaker
)
