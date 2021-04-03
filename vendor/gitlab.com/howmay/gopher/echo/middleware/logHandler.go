package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// RouteName output http request dump
func RouteName(routeName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()

			requestID := RequestIDFromContext(ctx)
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Response().Writer}
			c.Response().Writer = blw

			// 避免 Request Body 被讀過以後變成空的
			reqBody, _ := ioutil.ReadAll(c.Request().Body)
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			reqDump, _ := httputil.DumpRequest(c.Request(), false)

			// 繼續填充需要往下傳遞的 fields
			logger := log.With().Fields(
				map[string]interface{}{
					"endpoint":   routeName,
					"method":     c.Request().Method,
					"path":       c.Request().URL.Path,
					"request_id": requestID,
				},
			).Logger()
			ctx = logger.WithContext(ctx)
			c.SetRequest(c.Request().WithContext(ctx))

			c.Response().After(func() {
				logger.Info().Fields(
					map[string]interface{}{
						"input":     string(reqBody),
						"req_dump":  string(reqDump),
						"status":    c.Response().Status,
						"resp_dump": blw.body.String(),
					}).
					Msgf("%s access log", routeName)
			})

			return next(c)
		}
	}
}

type bodyLogWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	_, _ = w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyLogWriter) WriteString(s string) (int, error) {
	_, _ = w.body.WriteString(s)
	return w.ResponseWriter.Write([]byte(s))
}
