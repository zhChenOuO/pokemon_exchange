package spot_order

import (
	"pokemon/internal/pkg/iface"

	"github.com/bsm/redislock"
	"gitlab.com/howmay/gopher/db"
	"gitlab.com/howmay/gopher/redis"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo      iface.SpotOrderRepo
	tradeRepo iface.TradeOrderRepo
	cardRepo  iface.CardRepo
	db        *gorm.DB
	locker    *redislock.Client
}

type Params struct {
	fx.In

	Repo       iface.SpotOrderRepo
	TradeRepo  iface.TradeOrderRepo
	CardRepo   iface.CardRepo
	Conns      *db.Connections
	RedisConns redis.Redis
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.SpotOrderService {

	return &service{
		repo:      p.Repo,
		db:        p.Conns.WriteDB,
		locker:    redislock.New(p.RedisConns),
		tradeRepo: p.TradeRepo,
		cardRepo:  p.CardRepo,
	}
}

func (s *service) GetDB() *gorm.DB {
	return s.db
}
