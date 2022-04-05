package cmd

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"pokemon/configuration"
	"pokemon/pkg/delivery/restful"
	"pokemon/pkg/repository"
	"pokemon/pkg/service"
	"pokemon/pkg/usecase"
	"syscall"
	"time"

	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	"gitlab.com/howmay/gopher/db"
	"gitlab.com/howmay/gopher/echo"
	"gitlab.com/howmay/gopher/helper"
	"gitlab.com/howmay/gopher/redis"
	"gitlab.com/howmay/gopher/zlog"
	fx "go.uber.org/fx"
)

// ServerCmd 是此程式的Service入口點
var ServerCmd = &cobra.Command{
	Run: run,
	Use: "server",
}

var Module = fx.Options(
	fx.Provide(
		db.InitDatabases,
		echo.StartEcho,
		redis.InitRedisClient,
	),
)

func run(command *cobra.Command, _ []string) {
	defer helper.Recover(command.Context())

	config, err := configuration.New()
	if err != nil {
		os.Exit(0)
		return
	}

	zlog.New(config.Log)

	app := fx.New(
		fx.Supply(*config),
		Module,
		service.Module,
		repository.Module,
		usecase.Module,
		restful.Module,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopChan
	log.Info().Msg("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Error().Msg(err.Error())
	}

	os.Exit(exitCode)
}
