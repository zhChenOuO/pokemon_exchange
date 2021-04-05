package matching

import (
	"context"
	"pokemon/pkg/iface"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo      iface.SpotOrderRepo
	tradeRepo iface.TradeOrderRepo
	cardRepo  iface.CardRepo
	db        *gorm.DB
	// locker    *redislock.Client
	buyOrder  map[uint64]*rbt.Tree
	sellOrder map[uint64]*rbt.Tree
}

type Params struct {
	fx.In

	Repo      iface.SpotOrderRepo
	TradeRepo iface.TradeOrderRepo
	CardRepo  iface.CardRepo
	Conns     *db.Connections
	// RedisConns redis.Redis
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) (iface.MatchingUsecase, error) {
	s := service{
		repo: p.Repo,
		db:   p.Conns.WriteDB,
		// locker:    redislock.New(p.RedisConns),
		tradeRepo: p.TradeRepo,
		cardRepo:  p.CardRepo,
		sellOrder: make(map[uint64]*rbt.Tree),
		buyOrder:  make(map[uint64]*rbt.Tree),
	}
	ctx := log.Logger.WithContext(context.Background())
	cards, _, err := s.cardRepo.ListCards(ctx, nil, option.CardWhereOption{})
	if err != nil {
		return nil, err
	}
	for i := range cards {
		s.sellOrder[cards[i].ID] = rbt.NewWith(model.DecimalComparator)
		s.buyOrder[cards[i].ID] = rbt.NewWith(model.DecimalComparator)
	}

	return &s, nil
}

func (s *service) GetDB() *gorm.DB {
	return s.db
}

func (s *service) PubOrder(o *model.SpotOrder) {
	var tree *rbt.Tree
	switch o.TradeSide {
	case model.BuySide:
		tree = s.buyOrder[o.CardID]
	case model.SellSide:
		tree = s.sellOrder[o.CardID]
	}

	if v, found := tree.Get(o.ExpectedAmount); found {
		subTree := v.(*rbt.Tree)
		subTree.Put(o.ID, o)
		return
	}

	subTree := rbt.NewWith(utils.UInt64Comparator)
	subTree.Put(o.ID, o)
	tree.Put(o.ExpectedAmount, subTree)
}

func (s *service) GetMatchOrder(o *model.SpotOrder) []*model.SpotOrder {
	// TODO 尚未補上 方向性
	// TODO 買單 由小到大的查看掛單簿
	// TODO 賣單 由大到小的查看掛單簿
	var (
		tree   *rbt.Tree
		result = make([]*model.SpotOrder, 0)
	)
	switch o.TradeSide {
	case model.BuySide:
		tree = s.sellOrder[o.CardID]
	case model.SellSide:
		tree = s.buyOrder[o.CardID]
	}
	if v, found := tree.Get(o.ExpectedAmount); found {
		subTree := v.(*rbt.Tree)

		var (
			needQuantity = o.CardQuantity
		)
		for _, iKey := range subTree.Values() {
			_order := iKey.(*model.SpotOrder)
			result = append(result, _order)
			if needQuantity.LessThanOrEqual(_order.CardQuantity) {
				break
			} else {
				needQuantity.Sub(_order.CardQuantity)
			}
		}
	}
	return result
}
