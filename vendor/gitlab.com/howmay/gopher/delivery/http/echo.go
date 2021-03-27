package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/errors"
	"gitlab.com/howmay/gopher/middleware"
	"go.uber.org/fx"
)

// Config setting http config
type Config struct {
	Debug         bool   `json:"mode"`
	Address       string `json:"address"`
	AppID         string `yaml:"app_id" mapstructure:"app_id"`
	IsRequestDump bool   `yaml:"is_request_dump" mapstructure:"is_request_dump"`
}

// NewInjection 注入 config
func (hc *Config) NewInjection() *Config {
	return hc
}

// NewEcho create new engine for handler to register
func NewEcho(cfg *Config) *echo.Echo {
	echo.NotFoundHandler = errors.NotFoundHandlerForEcho
	echo.MethodNotAllowedHandler = errors.NotFoundHandlerForEcho

	e := echo.New()

	if cfg.Debug {
		e.Debug = true
		e.HideBanner = false
		e.HidePort = false
	} else {
		e.Debug = false
		e.HideBanner = true
		e.HidePort = true
	}
	e.HTTPErrorHandler = errors.HTTPErrorHandlerForEcho

	// 所有 API 皆經過讓 Header 帶著 request id 的 middleware
	// Context 會帶著有 request id 的 logger
	e.Pre(middleware.NewRequestIDMiddleware())

	// Request dump
	if cfg.IsRequestDump {
		e.Use(middleware.RequestDump())
	}

	// 處理所有 API 的異常 panic
	e.Use(middleware.NewErrorHandlingMiddleware())

	// 所有 API 皆經過 CORS middeware
	e.Use(middleware.CorsConfig)

	// 紀錄所有 API 的錯誤
	e.Use(errors.ErrMiddleware)

	RegisterDefaultRoute(e)
	return e
}

// StartEcho create new engine for handler to register
func StartEcho(s *Config, lc fx.Lifecycle) *echo.Echo {
	e := NewEcho(s)
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info().Msgf("Starting echo server, listen on %s", s.Address)
			go func() {
				err := e.Start(s.Address)
				if err != nil {
					log.Error().Msgf("Error echo server, err: %s", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Stopping echo HTTP server.")
			return e.Shutdown(ctx)
		},
	})
	return e
}

// RegisterDefaultRoute provide default handler
func RegisterDefaultRoute(e *echo.Echo) {
	// TODO: 後續補上普羅米修斯相關的 middleware
	// app.Use(middleware.NewPrometheusExporterMiddleware(config.C.Agent.Prometheus))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!!!")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!!")
	})
}

// ContextWithXRequestID returns a context.Context with given X-Request-Id value.
func ContextWithXRequestID(ctx context.Context, xrid string) context.Context {
	return context.WithValue(ctx, echo.HeaderXRequestID, xrid)
}
