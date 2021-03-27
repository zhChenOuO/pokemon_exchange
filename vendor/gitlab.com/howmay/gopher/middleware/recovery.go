package middleware

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/errors"
)

// NewErrorHandlingMiddleware handles panic error
func NewErrorHandlingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() error {
				if r := recover(); r != nil {
					trace := make([]byte, 4096)
					runtime.Stack(trace, true)
					requestID := c.Request().Header.Get(echo.HeaderXRequestID)
					customFields := map[string]interface{}{
						"url":         c.Request().RequestURI,
						"stack_error": string(trace),
						"request_id":  requestID,
					}
					var msg string
					for i := 2; ; i++ {
						_, file, line, ok := runtime.Caller(i)
						if !ok {
							break
						}
						msg += fmt.Sprintf("%s:%d\n", file, line)
					}
					logger := log.With().Fields(customFields).Logger()
					logger.Error().Msgf("%s\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥", r, msg)

					return c.JSON(500, errors.WithMessage(errors.ErrInternalError, msg))
				}
				return nil
			}()
			return next(c)
		}
	}
}
