package cmd

import (
	"context"
	"os"
	"os/signal"
	"pokemon/configuration"
	"pokemon/internal/pkg/delivery/restful"
	"pokemon/internal/pkg/repository"
	"pokemon/internal/pkg/service"
	"syscall"
	"time"

	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	"gitlab.com/howmay/gopher/db"
	"gitlab.com/howmay/gopher/delivery/http"
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
		http.StartEcho,
	),
)

func run(_ *cobra.Command, _ []string) {
	defer cmdRecover()

	config, err := configuration.New()
	if err != nil {
		os.Exit(0)
		return
	}

	zlog.InitV2(config.Log)

	app := fx.New(
		fx.Supply(*config),
		Module,
		repository.Model,
		service.Model,
		restful.Model,
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
