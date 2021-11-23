package iface

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
)

// type IServices interface {
// 	UserService
// 	CardService
// }

type UserService interface {
	GetUser(ctx context.Context, opt *option.UserWhereOption) (model.User, error)
	CreateUser(ctx context.Context, data *model.User) error
	ListUsers(ctx context.Context, opt *option.UserWhereOption) ([]model.User, int64, error)
	UpdateUser(ctx context.Context, opt *option.UserWhereOption, col *option.UserUpdateColumn) error
	DeleteUser(ctx context.Context, opt *option.UserWhereOption) error
}

type CardService interface {
	GetCard(ctx context.Context, opt *option.CardWhereOption) (model.Card, error)
	CreateCard(ctx context.Context, data *model.Card) error
	ListCards(ctx context.Context, opt *option.CardWhereOption) ([]model.Card, int64, error)
	UpdateCard(ctx context.Context, opt *option.CardWhereOption, col *option.CardUpdateColumn) error
	DeleteCard(ctx context.Context, opt *option.CardWhereOption) error
}

type SpotOrderService interface {
	GetSpotOrder(ctx context.Context, opt *option.SpotOrderWhereOption) (model.SpotOrder, error)
	CreateSpotOrder(ctx context.Context, data *model.SpotOrder) error
	ListSpotOrders(ctx context.Context, opt *option.SpotOrderWhereOption) ([]model.SpotOrder, int64, error)
	UpdateSpotOrder(ctx context.Context, opt *option.SpotOrderWhereOption, col *option.SpotOrderUpdateColumn) error
	DeleteSpotOrder(ctx context.Context, opt *option.SpotOrderWhereOption) error
}

type TradeOrderService interface {
	GetTradeOrder(ctx context.Context, opt *option.TradeOrderWhereOption) (model.TradeOrder, error)
	CreateTradeOrder(ctx context.Context, data *model.TradeOrder) error
	ListTradeOrders(ctx context.Context, opt *option.TradeOrderWhereOption) ([]model.TradeOrder, int64, error)
	UpdateTradeOrder(ctx context.Context, opt *option.TradeOrderWhereOption, col *option.TradeOrderUpdateColumn) error
	DeleteTradeOrder(ctx context.Context, opt *option.TradeOrderWhereOption) error
}

// IdentityAccountService service介面層
type IdentityAccountService interface {
	GetIdentityAccount(ctx context.Context, opt *option.IdentityAccountWhereOption) (model.IdentityAccount, error)
	CreateIdentityAccount(ctx context.Context, data *model.IdentityAccount) error
	ListIdentityAccounts(ctx context.Context, opt *option.IdentityAccountWhereOption) ([]model.IdentityAccount, int64, error)
	UpdateIdentityAccount(ctx context.Context, opt *option.IdentityAccountWhereOption, col *option.IdentityAccountUpdateColumn) error
	DeleteIdentityAccount(ctx context.Context, opt *option.IdentityAccountWhereOption) error

	VerifyIdentityAccount(ctx context.Context, data model.IdentityAccount) (model.IdentityAccount, error)
}

type IUsecase interface {
	MatchingUsecase
}

type MatchingUsecase interface {
	MatchingSpotOrder(ctx context.Context, data *model.SpotOrder) error
}
