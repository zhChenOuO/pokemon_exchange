package spot_order

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"

	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetSpotOrder 取得SpotOrder的資訊
func (s *service) GetSpotOrder(ctx context.Context, opt option.SpotOrderWhereOption) (model.SpotOrder, error) {
	return s.repo.GetSpotOrder(ctx, nil, opt)
}

// CreateSpotOrder 建立SpotOrder
func (s *service) CreateSpotOrder(ctx context.Context, data *model.SpotOrder) error {
	var (
		opt   option.SpotOrderWhereOption
		trade model.TradeOrder
	)

	switch data.TradeSide {
	case model.SellSide:
		opt.SpotOrder.TradeSide = model.BuySide
		opt.ExpectedAmountMoreThan = data.ExpectedAmount
	case model.BuySide:
		opt.SpotOrder.TradeSide = model.SellSide
		opt.ExpectedAmountLessThan = data.ExpectedAmount
	default:
		return errors.WithMessagef(errors.ErrInvalidInput, "trade side no support %d", data.TradeSide)
	}

	opt.SpotOrder.CardID = data.CardID
	opt.SpotOrder.Status = model.OrderWaitingForMatchmaking
	opt.Sorting.SortField = "id"
	opt.Sorting.SortOrder = "ASC"

	// lock, err := s.locker.Obtain(ctx, data.RedisLockKey(), 200*time.Millisecond, &redislock.Options{
	// 	RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(time.Second), 3),
	// })
	// if err == redislock.ErrNotObtained {
	// 	// fmt.Println(errors.New("Could not obtain lock!"))
	// 	log.Error().Msgf("could not obtain lock %+v", time.Since(t1).Milliseconds())
	// 	return nil
	// } else if err != nil {
	// 	return errors.Wrapf(errors.ErrInternalError, "fail to get lock err: %+v", err)
	// }
	// defer lock.Release(ctx)

	txErr := s.GetDB().Transaction(func(tx *gorm.DB) error {

		if _, err := s.cardRepo.GetCard(ctx, tx, option.CardWhereOption{
			Card: model.Card{
				ID: data.CardID,
			},
		}, func(db *gorm.DB) *gorm.DB {
			return db.Clauses(clause.Locking{Strength: "UPDATE"})
		}); err != nil {
			return err
		}

		makerSO, err := s.repo.GetSpotOrder(ctx, tx, opt)
		if err != nil && !errors.Is(err, errors.ErrResourceNotFound) {
			log.Error().Msgf("fail to get spot order")
			return err
		}
		if makerSO == (model.SpotOrder{}) {
			if err := s.repo.CreateSpotOrder(ctx, tx, data); err != nil {
				log.Error().Msgf("fail to create spot order")
				return err
			}
			return nil
		}

		data.SetSuccess(model.OrderTypeTaker)
		if err := s.repo.CreateSpotOrder(ctx, tx, data); err != nil {
			log.Error().Msgf("fail to create spot order")
			return err
		}

		trade.InitTradeOrder(&makerSO, data)
		if err := s.tradeRepo.CreateTradeOrder(ctx, tx, &trade); err != nil {
			log.Error().Msgf("fail to create trade order")
			return err
		}

		if err := s.repo.UpdateSpotOrder(ctx, tx, option.SpotOrderUpdateOption{
			WhereOpts: option.SpotOrderWhereOption{
				SpotOrder: model.SpotOrder{
					ID: makerSO.ID,
				},
			},
			UpdateCol: option.SpotOrderUpdateColumn{
				Status: model.OrderSuccess,
				Type:   makerSO.Type,
			},
		}); err != nil {
			log.Error().Msgf("fail to update spot order err:%+v", err.Error())
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}
	return nil
}

// ListSpotOrders 列出SpotOrder
func (s *service) ListSpotOrders(ctx context.Context, opt option.SpotOrderWhereOption) ([]model.SpotOrder, int64, error) {
	so, total, err := s.repo.ListSpotOrders(ctx, nil, opt)
	if err != nil {
		return nil, 0, err
	}
	return so, total, nil
}

// UpdateSpotOrder 更新SpotOrder
func (s *service) UpdateSpotOrder(ctx context.Context, opt option.SpotOrderUpdateOption) error {
	err := s.repo.UpdateSpotOrder(ctx, nil, opt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpotOrder 刪除SpotOrder
func (s *service) DeleteSpotOrder(ctx context.Context, opt option.SpotOrderWhereOption) error {
	return s.repo.DeleteSpotOrder(ctx, nil, opt)
}
