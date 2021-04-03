package spot_order

import (
	"context"
	"math/rand"
	"pokemon/pkg/model"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func Benchmark_service_CreateSpotOrder(b *testing.B) {
	ctx := log.Logger.WithContext(context.Background())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		getCardID := func() uint64 {
			return uint64(rand.Int63n(3) + 1)
		}

		getTradeSide := func() model.OrderTradeSide {
			return model.OrderTradeSide(rand.Int63n(1) + 1)
		}

		getAmount := func() decimal.Decimal {
			return decimal.NewFromFloat((rand.Float64() * (100 - 0.01)) + 0.01).Round(2)
		}
		so := model.SpotOrder{
			UserID:         1,
			CardID:         getCardID(),
			TradeSide:      getTradeSide(),
			CardQuantity:   decimal.NewFromInt(1),
			ExpectedAmount: getAmount(),
		}
		b.StartTimer()
		if err := suite.svc.CreateSpotOrder(ctx, &so); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
