package test_fixture

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"pokemon/configuration"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cenk/backoff"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pressly/goose"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gitlab.com/howmay/gopher/db"
	"gitlab.com/howmay/gopher/redis"
	"gitlab.com/howmay/gopher/zlog"
	"go.uber.org/fx"
)

// Suite ...
type Suite struct {
	app            *fx.App
	pgContainer    *dockertest.Resource
	redisContainer *dockertest.Resource
	pool           *dockertest.Pool
}

var suite Suite

// Initialize 初始化 suite
func Initialize(fxOption ...fx.Option) error {
	viper.AutomaticEnv()
	if os.Getenv("CONFIG_NAME") == "" {
		_ = os.Setenv("CONFIG_NAME", "app-test")
	}

	var cfg *configuration.Configuration = &configuration.Configuration{
		Log: &zlog.Config{
			Local: true,
		},
		Database: &db.Config{},
	}
	var err error
	isLocalTest := os.Getenv("LOCAL_TEST")
	if isLocalTest == "1" {
		suite.pool, err = dockertest.NewPool("")
		if err != nil {
			log.Error().Msgf("Could not connect to docker: %s", err)
			return err
		}
		databaseName := "postgres"
		suite.pgContainer, err = suite.pool.RunWithOptions(
			&dockertest.RunOptions{
				Repository: "postgres",
				Tag:        "12",
				Env: []string{
					"POSTGRES_PASSWORD=postgres",
					"POSTGRES_DB =" + databaseName,
				},
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{
					Name: "no",
				}
			})
		if err != nil {
			log.Error().Msgf("Could not start resource: %s", err)
			return err
		}
		pgPort, _ := strconv.Atoi(suite.pgContainer.GetPort("5432/tcp"))

		dbCfg := &db.Database{
			Host:         "localhost",
			Port:         pgPort,
			Password:     "postgres",
			Username:     "postgres",
			Type:         db.Postgres,
			DBName:       databaseName,
			Debug:        false,
			SearchPath:   "pokemon",
			MaxOpenConns: 400,
			MaxIdleConns: 200,
		}
		cfg.Database.Read = dbCfg
		cfg.Database.Write = dbCfg

		suite.redisContainer, err = suite.pool.RunWithOptions(&dockertest.RunOptions{
			Repository: "redis",
			Tag:        "alpine",
		}, func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		})
		if err != nil {
			log.Error().Msgf("Could not start resource: %s", err)
			return err
		}
		// redisPort, _ := strconv.Atoi(suite.redisContainer.GetPort("6379/tcp"))
		// cfg.Redis = &redis.Config{
		// 	Addresses: []string{fmt.Sprintf("localhost:%d", redisPort)},
		// }
	} else {
		cfg, err = configuration.New()
		if err != nil {
			return err
		}
	}
	// configuration.Apis = &client.APIConfigs{}
	t := &testing.T{}
	base := []fx.Option{
		fx.Supply(*cfg),
		fx.Supply(t),
		fx.Provide(
			db.InitDatabases,
			redis.InitRedisClient,
		),
		fx.Invoke(zlog.Init),
		fx.Invoke(Migration),
	}

	base = append(base, fxOption...)

	app := fx.New(
		base...,
	)

	suite.app = app
	return app.Start(context.Background())
}

// Close 停止 container
func Close() {
	log.Info().Msg("close app")
	isLocalTest := viper.GetString("LOCAL_TEST")
	if isLocalTest == "1" {
		if err := suite.pool.Purge(suite.pgContainer); err != nil {
			log.Error().Msgf("Could not purge resource: %s", err)
		}
		if err := suite.pool.Purge(suite.redisContainer); err != nil {
			log.Error().Msgf("Could not purge resource: %s", err)
		}
	}
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

	fmt.Println(viper.GetString("PROJ_DIR") + "/deployment/database")
	if err := goose.Run("up", db, viper.GetString("PROJ_DIR")+"/deployment/database"); err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}
