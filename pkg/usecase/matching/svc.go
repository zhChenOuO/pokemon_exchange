package matching

import (
	"context"
	"pokemon/pkg/iface"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"
	"sync"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"gitlab.com/howmay/gopher/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// service ...
type service struct {
	repo          iface.SpotOrderRepo
	tradeRepo     iface.TradeOrderRepo
	cardRepo      iface.CardRepo
	db            *gorm.DB
	buyOrderLock  *sync.RWMutex
	buyOrder      map[uint64]*rbt.Tree
	sellOrderLock *sync.RWMutex
	sellOrder     map[uint64]*rbt.Tree
	// locker    *redislock.Client

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
		tradeRepo:     p.TradeRepo,
		cardRepo:      p.CardRepo,
		sellOrder:     make(map[uint64]*rbt.Tree),
		buyOrder:      make(map[uint64]*rbt.Tree),
		buyOrderLock:  &sync.RWMutex{},
		sellOrderLock: &sync.RWMutex{},
	}
	ctx := log.Logger.WithContext(context.Background())
	cards, _, err := s.cardRepo.ListCards(ctx, nil, option.CardWhereOption{})
	if err != nil {
		return nil, err
	}
	for i := range cards {
		s.sellOrder[cards[i].ID] = rbt.NewWith(model.DecimalASCComparator)
		s.buyOrder[cards[i].ID] = rbt.NewWith(model.DecimalDESCComparator)
	}

	sos, err := p.Repo.ListSpotOrdersWithLock(ctx)
	if err != nil {
		return nil, err
	}

	// TODO 若開放數字不為1時 card_quantity 需要再扣除 trade_order 成交的數量
	for i := range sos {
		s.PubOrder(&sos[i])
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
		s.buyOrderLock.Lock()
		defer s.buyOrderLock.Unlock()
		tree = s.buyOrder[o.CardID]
	case model.SellSide:
		s.sellOrderLock.Lock()
		defer s.sellOrderLock.Unlock()
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

func (s *service) RemoveOrder(o *model.SpotOrder) {
	var tree *rbt.Tree
	switch o.TradeSide {
	case model.BuySide:
		s.buyOrderLock.Lock()
		defer s.buyOrderLock.Unlock()
		tree = s.buyOrder[o.CardID]
	case model.SellSide:
		s.sellOrderLock.Lock()
		defer s.sellOrderLock.Unlock()
		tree = s.sellOrder[o.CardID]
	}
	if v, found := tree.Get(o.ExpectedAmount); found {
		subTree := v.(*rbt.Tree)
		subTree.Remove(o.ID)
		return
	}
}

func (s *service) GetMatchOrder(o *model.SpotOrder) []*model.SpotOrder {
	var (
		tree   *rbt.Tree
		result = make([]*model.SpotOrder, 0)
	)
	switch o.TradeSide {
	case model.BuySide:
		s.buyOrderLock.RLock()
		defer s.buyOrderLock.RUnlock()
		tree = s.sellOrder[o.CardID]
	case model.SellSide:
		s.sellOrderLock.RLock()
		defer s.sellOrderLock.RUnlock()
		tree = s.buyOrder[o.CardID]
	default:
		return result
	}

	var (
		needQuantity = o.CardQuantity
	)

	for i := 0; i < tree.Size(); i++ {
		l := tree.Left()
		log.Info().Msg(tree.String())
		log.Info().Msg(l.String())
		orderBookAmount := l.Key.(decimal.Decimal)
		switch o.TradeSide {
		case model.BuySide:
			// 買單 , 賣單簿的最低價 小於下單價則略過
			if o.ExpectedAmount.LessThan(orderBookAmount) {
				return result
			}
		case model.SellSide:
			// 賣單 , 買單簿的最高價 大於下單價則略過
			if o.ExpectedAmount.GreaterThan(orderBookAmount) {
				return result
			}
		}
		subTree := l.Value.(*rbt.Tree)
		log.Info().Msgf("subtree \n%s\n", subTree.String())
		for j := 0; j < subTree.Size(); j++ {
			idLeft := subTree.Left()
			defer subTree.Remove(idLeft.Key)
			_order := idLeft.Value.(*model.SpotOrder)
			result = append(result, _order)
			needQuantity = needQuantity.Sub(_order.CardQuantity)
			if needQuantity.IsNegative() || needQuantity.IsZero() {
				return result
			}
		}
	}

	// for _, v := range tree.Keys() {
	// 	orderBookAmount := v.(decimal.Decimal)
	// 	switch o.TradeSide {
	// 	case model.BuySide:
	// 		// 買單 , 賣單簿的最低價 小於下單價則略過
	// 		if o.ExpectedAmount.LessThan(orderBookAmount) {
	// 			return result
	// 		}
	// 	case model.SellSide:
	// 		// 賣單 , 買單簿的最高價 大於下單價則略過
	// 		if o.ExpectedAmount.GreaterThan(orderBookAmount) {
	// 			return result
	// 		}
	// 	}

	// 	subTree, found := tree.Get(v)
	// 	if !found {
	// 		log.Error().Msgf("not found key in tree")
	// 		continue
	// 	}
	// 	for _, iKey := range subTree.(*rbt.Tree).Values() {
	// 		_order := iKey.(*model.SpotOrder)
	// 		result = append(result, _order)
	// 		needQuantity = needQuantity.Sub(_order.CardQuantity)
	// 		if needQuantity.IsNegative() || needQuantity.IsZero() {
	// 			return result
	// 		}
	// 	}
	// }
	return result
}
