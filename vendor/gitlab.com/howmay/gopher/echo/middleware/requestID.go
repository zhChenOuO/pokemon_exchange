package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/ctxutil"
)

// CtxKey 用來代表 context.Context 的 key
type CtxKey string

func (ck CtxKey) String() string {
	return string(ck)
}

// 一些神奇的 key 只剩gam在用 盡量用 echo.HeaderXRequestID 用了會被 Harvey 罵 瑞翔說的
const (
	CtxKeyRequestID CtxKey = "X-Request-Id"
	CtxKeyEndpoint  CtxKey = "endPoint"
)

// RequestIDFromContext 從 ctx 中取得 request id, 如果沒有即時產生一個
func RequestIDFromContext(ctx context.Context) string {
	rid, ok := ctx.Value(echo.HeaderXRequestID).(string)
	if !ok {
		// 產生 requestID 並傳下去
		rid = xid.New().String()
		return rid
	}
	return rid
}

// NewRequestIDMiddleware Default returns the location middleware with default configuration.
func NewRequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			requestID := c.Request().Header.Get(ctxutil.XRequestID.String())
			if requestID == "" {
				requestID = xid.New().String()
			}
			ctx := context.WithValue(c.Request().Context(), ctxutil.XRequestID, requestID)
			logger := log.With().Str("request_id", requestID).Logger()

			ctx = logger.WithContext(ctx)

			c.SetRequest(c.Request().WithContext(ctx))
			// Set X-Request-Id header
			c.Response().Writer.Header().Set(ctxutil.XRequestID.String(), requestID)
			return next(c)
		}
	}
}
