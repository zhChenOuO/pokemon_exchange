package db

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/ctxutil"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

type gormLogger struct {
	LogLevel                            logger.LogLevel
	Config                              logger.Config
	SlowThreshold                       time.Duration
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// NewLogger for gorm log use zerolog print
func NewLogger(config logger.Config) logger.Interface {
	var (
		infoStr      = "%s\n" + "[info] "
		warnStr      = "%s\n" + "[warn] "
		errStr       = "%s\n" + "[error] "
		traceStr     = "%s\n" + "[%.3fms]" + "[rows:%v]" + "\n%s\n"
		traceWarnStr = "%s" + ", %s\n" + "[%.3fms]" + "[rows:%v]" + "\n%s" + "\n"
		traceErrStr  = "%s" + ", %s\n" + "[%.3fms]" + "[rows:%v]" + "\n%s\n"
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr = BlueBold + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Cyan + " \n%s\n"
		traceWarnStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " \n%s" + Reset + "\n"
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + "\n%s\n"
	}

	l := &gormLogger{
		LogLevel:      config.LogLevel,
		SlowThreshold: config.SlowThreshold,
		Config:        config,
		infoStr:       infoStr,
		warnStr:       warnStr,
		errStr:        errStr,
		traceStr:      traceStr,
		traceWarnStr:  traceWarnStr,
		traceErrStr:   traceErrStr,
	}

	return l
}

//LogMode ...
func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *g
	newlogger.LogLevel = level
	return &newlogger
}

//Info ...
func (g gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if g.LogLevel >= logger.Info {
		requestID := ctxutil.FromContext(ctx)
		currentLogger := log.With().Fields(map[string]interface{}{
			"request_id": requestID,
			"db_log":     true,
		}).Logger()

		currentLogger.Info().Msgf(
			g.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

//Warn ....
func (g gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if g.LogLevel >= logger.Warn {
		requestID := ctxutil.FromContext(ctx)
		currentLogger := log.With().Fields(map[string]interface{}{
			"request_id": requestID,
			"db_log":     true,
		}).Logger()

		currentLogger.Warn().Msgf(
			g.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

//Error ...
func (g gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if g.LogLevel >= logger.Error {
		requestID := ctxutil.FromContext(ctx)
		currentLogger := log.With().Fields(map[string]interface{}{
			"request_id": requestID,
			"db_log":     true,
		}).Logger()

		currentLogger.Error().Msgf(g.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

//Trace ...
func (g gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.LogLevel > 0 {
		requestID := ctxutil.FromContext(ctx)
		currentLogger := log.With().Fields(map[string]interface{}{
			"request_id": requestID,
			"db_log":     true,
		}).Logger()

		elapsed := time.Since(begin)
		switch {
		case err != nil && g.LogLevel >= logger.Error:
			sql, rows := fc()
			currentLogger.Error().Msgf(g.traceErrStr, utils.FileWithLineNum(), err, sql, rows, float64(elapsed.Nanoseconds())/1e6)
		case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= logger.Warn:
			sql, rows := fc()
			currentLogger.Warn().Msgf(g.traceWarnStr, utils.FileWithLineNum(), sql, rows, float64(elapsed.Nanoseconds())/1e6)
		case g.LogLevel >= logger.Info:
			sql, rows := fc()
			currentLogger.Info().Msgf(g.traceStr, utils.FileWithLineNum(), sql, rows, float64(elapsed.Nanoseconds())/1e6)
		}
	}
}
