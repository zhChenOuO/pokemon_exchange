package matching

import (
	"context"
	"pokemon/pkg/model"
	"pokemon/pkg/model/option"

	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/common"
	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *service) MatchingSpotOrder(ctx context.Context, data *model.SpotOrder) error {
	var (
		opt   option.SpotOrderWhereOption
		trade model.TradeOrder
	)

	switch data.TradeSide {
	case model.BuySide:
		opt.SpotOrder.TradeSide = model.SellSide
		opt.ExpectedAmountLessThan = data.ExpectedAmount
	case model.SellSide:
		opt.SpotOrder.TradeSide = model.BuySide
		opt.ExpectedAmountMoreThan = data.ExpectedAmount
	default:
		return errors.WithMessagef(errors.ErrInvalidInput, "trade side no support %d", data.TradeSide)
	}

	opt.SpotOrder.CardID = data.CardID
	opt.SpotOrder.Status = model.OrderWaitingForMatchmaking
	opt.Sorting.SortField = "id"
	opt.Sorting.Type = common.SortingOrderType_ASC

	txErr := s.GetDB().Transaction(func(tx *gorm.DB) error {
		var card model.Card
		if err := s.repo.Get(ctx, tx, &card, &option.CardWhereOption{
			Card: model.Card{
				ID: data.CardID,
			},
		}, func(db *gorm.DB) *gorm.DB {
			return db.Clauses(clause.Locking{Strength: "UPDATE"})
		}); err != nil {
			return err
		}

		makerSOs := s.GetMatchOrder(data)
		if len(makerSOs) == 0 {
			if err := s.repo.Create(ctx, tx, data); err != nil {
				log.Error().Msgf("fail to create spot order %+v", err)
				return err
			}
			s.PubOrder(data)
			return nil
		}

		data.SetSuccess(model.OrderTypeTaker)
		if err := s.repo.Create(ctx, tx, data); err != nil {
			log.Error().Msgf("fail to create spot order")
			return err
		}

		for _, makerSO := range makerSOs {
			trade.InitTradeOrder(makerSO, data)
			if err := s.repo.Create(ctx, tx, &trade); err != nil {
				log.Error().Msgf("fail to create trade order")
				return err
			}

			if makerSO.CardQuantity.IsZero() {
				if err := s.repo.Update(ctx, tx, &option.SpotOrderWhereOption{
					SpotOrder: model.SpotOrder{
						ID: makerSO.ID,
					}},
					&option.SpotOrderUpdateColumn{
						Status: model.OrderSuccess,
					}); err != nil {
					log.Error().Msgf("fail to update spot order err:%+v", err.Error())
					return err
				}
				s.RemoveOrder(makerSO)
			} else {
				s.PubOrder(makerSO)
			}
		}

		if data.CardQuantity.IsZero() {
			if err := s.repo.Update(ctx, tx, &option.SpotOrderWhereOption{
				SpotOrder: model.SpotOrder{
					ID: data.ID,
				}},
				&option.SpotOrderUpdateColumn{
					Status: model.OrderSuccess,
				}); err != nil {
				return err
			}
		} else {
			s.PubOrder(data)
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
