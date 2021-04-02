package iface

import (
	"context"
	"pokemon/internal/pkg/model"
	"pokemon/internal/pkg/model/option"

	"gorm.io/gorm"
)

type IRepository interface {
	CardRepo
	UserRepo
}

type CardRepo interface {
	GetCard(ctx context.Context, tx *gorm.DB, opt option.CardWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.Card, error)
	CreateCard(ctx context.Context, tx *gorm.DB, data *model.Card, scopes ...func(*gorm.DB) *gorm.DB) error
	ListCards(ctx context.Context, tx *gorm.DB, opt option.CardWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.Card, int64, error)
	UpdateCard(ctx context.Context, tx *gorm.DB, opt option.CardUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error
	DeleteCard(ctx context.Context, tx *gorm.DB, opt option.CardWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
}

type UserRepo interface {
	GetUser(ctx context.Context, tx *gorm.DB, opt option.UserWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.User, error)
	CreateUser(ctx context.Context, tx *gorm.DB, data *model.User, scopes ...func(*gorm.DB) *gorm.DB) error
	ListUsers(ctx context.Context, tx *gorm.DB, opt option.UserWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.User, int64, error)
	UpdateUser(ctx context.Context, tx *gorm.DB, opt option.UserUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error
	DeleteUser(ctx context.Context, tx *gorm.DB, opt option.UserWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
}

type SpotOrderRepo interface {
	GetSpotOrder(ctx context.Context, tx *gorm.DB, opt option.SpotOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.SpotOrder, error)
	CreateSpotOrder(ctx context.Context, tx *gorm.DB, data *model.SpotOrder, scopes ...func(*gorm.DB) *gorm.DB) error
	ListSpotOrders(ctx context.Context, tx *gorm.DB, opt option.SpotOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.SpotOrder, int64, error)
	UpdateSpotOrder(ctx context.Context, tx *gorm.DB, opt option.SpotOrderUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error
	DeleteSpotOrder(ctx context.Context, tx *gorm.DB, opt option.SpotOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
}

// IdentityAccountRepository repository介面層
type IdentityAccountRepository interface {
	GetIdentityAccount(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.IdentityAccount, error)
	CreateIdentityAccount(ctx context.Context, tx *gorm.DB, data *model.IdentityAccount, scopes ...func(*gorm.DB) *gorm.DB) error
	ListIdentityAccounts(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.IdentityAccount, int64, error)
	UpdateIdentityAccount(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error
	DeleteIdentityAccount(ctx context.Context, tx *gorm.DB, opt option.IdentityAccountWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
}

type TradeOrderRepo interface {
	GetTradeOrder(ctx context.Context, tx *gorm.DB, opt option.TradeOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) (model.TradeOrder, error)
	CreateTradeOrder(ctx context.Context, tx *gorm.DB, data *model.TradeOrder, scopes ...func(*gorm.DB) *gorm.DB) error
	ListTradeOrders(ctx context.Context, tx *gorm.DB, opt option.TradeOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) ([]model.TradeOrder, int64, error)
	UpdateTradeOrder(ctx context.Context, tx *gorm.DB, opt option.TradeOrderUpdateOption, scopes ...func(*gorm.DB) *gorm.DB) error
	DeleteTradeOrder(ctx context.Context, tx *gorm.DB, opt option.TradeOrderWhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
}
