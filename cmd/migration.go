package cmd

import (
	"os"
	"pokemon/configuration"
	"pokemon/internal/testsuite"

	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	"gitlab.com/howmay/gopher/zlog"
)

// MigrationCmd 是此程式的Service入口點
var MigrationCmd = &cobra.Command{
	Run: migrationRun,
	Use: "migration",
}

func migrationRun(_ *cobra.Command, _ []string) {
	defer cmdRecover()

	config, err := configuration.New()
	if err != nil {
		os.Exit(0)
		return
	}

	zlog.InitV2(config.Log)

	testsuite.Migration(config.Database)

	log.Info().Msgf("finish")
}
