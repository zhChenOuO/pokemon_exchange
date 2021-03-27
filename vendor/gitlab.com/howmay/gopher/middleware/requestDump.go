package middleware

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// RequestDump outpt http request dump
func RequestDump() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			logger := log.Ctx(c.Request().Context())
			req := c.Request()
			b, err := httputil.DumpRequest(req, true)

			if err == nil {
				logger.Info().Str("path", req.URL.Path).Str("req_body", string(b)).Msg("request dump")
			}
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer
			resp := c.Response()
			resp.After(func() {
				logger.Info().Str("path", req.URL.Path).Str("resp_body", resBody.String()).Msg("response dump")
			})
			return next(c)
		}
	}
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
