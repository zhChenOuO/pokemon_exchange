// Package zlog 負責初始化 zerolog 的格式和等級
package zlog

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	// Teal ...
	Teal = Color("\033[1;36m%s\033[0m")
	// Yellow ...
	Yellow = Color("\033[35m%s\033[0m")
)

// Color ...
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// Graylog 的錯誤等級
const (
	levelEmerg   = int8(0)
	levelAlert   = int8(1)
	levelCrit    = int8(2)
	levelErr     = int8(3)
	levelWarning = int8(4)
	levelNotice  = int8(5)
	levelInfo    = int8(6)
	levelDebug   = int8(7)
)

type severityHook struct{}

// Run ...
func (h severityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	lvl := int8(0)
	switch level {
	case zerolog.DebugLevel:
		lvl = levelDebug
	case zerolog.InfoLevel:
		lvl = levelInfo
	case zerolog.WarnLevel:
		lvl = levelWarning
	case zerolog.ErrorLevel:
		lvl = levelErr
	case zerolog.FatalLevel:
		lvl = levelCrit
	}
	e.Int8("graylog_level", lvl).
		Float64("timestamp", float64(time.Now().UnixNano()/int64(time.Millisecond))/1000)
	if msg == "" {
		e.Str("message", "no message")
	}
}

// Init 初始化 zerolog
func Init(debug bool) {
	zerolog.DisableSampling(true)
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	hostname, _ := os.Hostname()
	lvl := zerolog.InfoLevel
	if debug {
		lvl = zerolog.DebugLevel
	}
	log.Logger = zerolog.New(os.Stdout).Hook(severityHook{}).With().
		Str("host", hostname).
		Logger().
		Level(lvl)
}

// Config ...
type Config struct {
	Debug bool
	Local bool
	AppID string `yaml:"app_id" mapstructure:"app_id"`
	Env   string
}

// NewInjection ...
func (c Config) NewInjection() *Config {
	return &c
}

// InitV2 ...
func InitV2(c *Config) {
	zerolog.DisableSampling(true)
	zerolog.TimestampFieldName = "local_timestamp"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	hostname, _ := os.Hostname()
	lvl := zerolog.InfoLevel
	if c.Debug {
		lvl = zerolog.DebugLevel
	}

	var z zerolog.Logger

	if c.Local {
		output := zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("[ %s ]", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", Teal(i))
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatTimestamp = func(i interface{}) string {
			t := fmt.Sprintf("%s", i)
			millisecond, err := strconv.ParseInt(fmt.Sprintf("%s", i), 10, 64)
			if err == nil {
				t = time.Unix(int64(millisecond/1000), 0).Local().Format("2006/01/02 15:04:05")
			}
			return Yellow(t)
		}
		z = zerolog.New(output)
	} else {
		z = zerolog.New(os.Stdout)
	}

	log.Logger = z.Hook(severityHook{}).With().
		Fields(map[string]interface{}{
			"app_id": c.AppID,
			"env":    c.Env,
		}).
		Str("host", hostname).
		Timestamp().
		Caller().
		Logger().
		Level(lvl)
}

// Ctx wrap zerolog Ctx func, if ctx not setting logger, return a default prevent for panic
func Ctx(ctx context.Context) *zerolog.Logger {
	defaultLogger := log.Logger
	if ctx == nil {
		defaultLogger.Warn().Msg("zlog func Ctx() not set context.Context in right way.")
		return &defaultLogger
	}

	return log.Ctx(ctx) // if ctx is not null and not set logger yet. A disabled logger is returned.
}
