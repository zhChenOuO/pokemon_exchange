package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"pokemon/configuration"
	"strconv"
	"strings"
	"time"

	"github.com/cenk/backoff"
	"github.com/pressly/goose"
	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/howmay/gopher/db"
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

	Migration(config.Database)

	log.Info().Msgf("finish")
}

func Migration(dbCfg *db.Config) error {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		dbCfg.Write.Username, dbCfg.Write.Password, dbCfg.Write.Host+":"+strconv.Itoa(dbCfg.Write.Port), dbCfg.Write.DBName)

	if strings.TrimSpace(dbCfg.Write.SearchPath) != "" {
		connectionString = fmt.Sprintf("%s&search_path=%s", connectionString, dbCfg.Write.SearchPath)
	}

	fmt.Println(connectionString)

	err := goose.SetDialect("postgres")
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	var db *sql.DB
	err = backoff.Retry(func() error {
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Error().Msg(err.Error())
			return err
		}
		err = db.Ping()
		if err != nil {
			log.Error().Msgf("main: %s ping error: %v", dbCfg.Write.Type, err)
			return err
		}
		return nil
	}, bo)

	if err != nil {
		log.Error().Msgf("main: %s connect err: %s", dbCfg.Write.Type, err.Error())
		return err
	}

	if err := goose.Run("up", db, viper.GetString("PROJ_DIR")+"/deployment/database"); err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}
