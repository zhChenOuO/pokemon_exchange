package db

import (
	"time"

	"github.com/cenk/backoff"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connections for gorm v2 ...
type Connections struct {
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

// InitDatabases init and return write and read DB objects
func InitDatabases(cfg *Config) (*Connections, error) {
	var (
		conn Connections
		err  error
	)
	conn.ReadDB, err = SetupDatabase(cfg.Read)
	if err != nil {
		return nil, err
	}

	conn.WriteDB, err = SetupDatabase(cfg.Write)
	if err != nil {
		return nil, err
	}
	return &conn, err
}

// SetupDatabase ...
func SetupDatabase(database *Database) (*gorm.DB, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	if database.WriteTimeout == "" {
		database.WriteTimeout = "10s"
	}
	if database.ReadTimeout == "" {
		database.ReadTimeout = "10s"
	}

	var dialector gorm.Dialector

	dsn, err := GetConnectionStr(database)
	if err != nil {
		return nil, err
	}

	switch database.Type {
	case MySQL:
		dialector = mysql.Open(dsn)
	case Postgres:
		dialector = postgres.Open(dsn)
	}

	log.Debug().Msgf("main: database connection string: %s", dsn)

	colorful := false
	logLevel := logger.Silent
	if database.Debug {
		logLevel = logger.Info
		colorful = true
	}

	newLogger := NewLogger(logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logLevel,    // Log level
		Colorful:      colorful,    // Disable color
	})

	var conn *gorm.DB
	err = backoff.Retry(func() error {
		db, err := gorm.Open(dialector, &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			return err
		}
		conn = db

		sqlDB, err := conn.DB()
		if err != nil {
			return err
		}

		err = sqlDB.Ping()
		return err
	}, bo)

	if err != nil {
		log.Error().Msgf("main: database connect err: %s", err.Error())
		return nil, err
	}
	log.Info().Msgf("database ping success")

	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}

	if database.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	} else {
		sqlDB.SetMaxIdleConns(2)
	}

	if database.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(database.MaxOpenConns)
	} else {
		sqlDB.SetMaxOpenConns(5)
	}

	if database.MaxLifetimeSec != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(database.MaxLifetimeSec) * time.Second)
	} else {
		sqlDB.SetConnMaxLifetime(14400 * time.Second)
	}

	return conn, nil
}
