package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Database ...
type Database struct {
	Debug          bool
	Type           DatabaseType
	Host           string
	Port           int
	Username       string
	Password       string
	DBName         string
	MaxIdleConns   int `yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns   int `yaml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxLifetimeSec int
	ReadTimeout    string `yaml:"read_timeout"`
	WriteTimeout   string `yaml:"write_timeout"`
	SearchPath     string `yaml:"search_path" mapstructure:"search_path"` // pg should setting this value. It will restrict access to db schema.
	SSLEnable      bool   `yaml:"ssl_enable" mapstructure:"ssl_enable"`   // pg ssl mode
	WithColor      bool   `yaml:"with_color" mapstructure:"with_color"`
}

// DatabaseType 類型
type DatabaseType string

const (
	// MySQL ...
	MySQL DatabaseType = "mysql"
	// Postgres ...
	Postgres DatabaseType = "postgres"
)

// GetConnectionStr ...
func GetConnectionStr(database *Database) (string, error) {
	var connectionString string
	switch database.Type {
	case MySQL:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&multiStatements=true&readTimeout=%s&writeTimeout=%s", database.Username, database.Password, database.Host+":"+strconv.Itoa(database.Port), database.DBName, database.ReadTimeout, database.WriteTimeout)
	case Postgres:
		connectionString = fmt.Sprintf(`user=%s password=%s host=%s port=%d dbname=%s`, database.Username, database.Password, database.Host, database.Port, database.DBName)

		if database.SSLEnable {
			connectionString += " sslmode=require"
		} else {
			connectionString += " sslmode=disable"
		}

		if strings.TrimSpace(database.SearchPath) != "" {
			connectionString = fmt.Sprintf("%s search_path=%s", connectionString, database.SearchPath)
		}
	default:
		return "", errors.New("not support driver")
	}
	return connectionString, nil
}
