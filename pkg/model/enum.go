package model

type OrderTradeSide int8

const (
	BuySide OrderTradeSide = iota + 1
	SellSide
)

type OrderType int8

const (
	// OrderTypeMaker 掛單者
	OrderTypeMaker OrderType = iota + 1
	// OrderTypeTaker 吃單者
	OrderTypeTaker
)
