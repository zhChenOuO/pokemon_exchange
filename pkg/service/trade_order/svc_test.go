package trade_order

import (
	"context"
	"os"
	"pokemon/internal/test_fixture"
	"pokemon/pkg/iface"
	"pokemon/pkg/repository"
	"testing"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type Suite struct {
	ctx context.Context
	svc iface.SpotOrderService
}

var suite Suite

func TestMain(m *testing.M) {
	test_fixture.Initialize(
		fx.Provide(New),
		repository.Module,
		fx.Populate(&suite.svc),
	)

	ctx := log.Logger.WithContext(context.Background())
	suite.ctx = ctx
	e := m.Run()
	test_fixture.Close()
	os.Exit(e)
}
