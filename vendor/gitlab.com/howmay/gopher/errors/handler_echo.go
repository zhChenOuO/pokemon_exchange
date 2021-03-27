package errors

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryClientErrorInterceptor ...
func UnaryClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(parentCtx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		lastErr := invoker(ContextWithMeta(parentCtx, XRequestIDFromContext(parentCtx)), method, req, reply, cc, opts...)
		return ConvertHttpErr(lastErr)
	}
}

//UnaryErrorInterceptor ...
func UnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {

		resp, err := handler(ctx, req)
		if err != nil {
			logFields := map[string]interface{}{}
			requestID := MetaFromContext(ctx)
			logFields["requestID"] = requestID
			logFields["input"] = req
			logFields["output"] = resp
			causeErr := errors.Cause(err)
			_err, ok := causeErr.(*_error)
			if !ok || _err == nil {
				return resp, status.Error(ErrInternalError.GRPCCode, err.Error())
			}
			// 根據狀態碼用不同等級來紀錄
			logger := log.With().Fields(logFields).Logger()
			if _err.Status >= http.StatusInternalServerError {
				logger.Error().Msgf("%+v", err)
			} else {
				logger.Debug().Msgf("%+v", err)
			}
			return resp, ConvertProtoErr(err)
		}
		return resp, err
	}
}

// HTTPErrorHandlerForEcho responds error response according to given error.
func HTTPErrorHandlerForEcho(err error, c echo.Context) {
	if err == nil {
		return
	}

	echoErr, ok := err.(*echo.HTTPError)
	if ok {
		_ = c.JSON(echoErr.Code, echoErr)
		return
	}

	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok || _err == nil {
		_ = c.JSON(http.StatusInternalServerError, ErrInternalError)
		return
	}
	if len(_err.Code) < 3 {
		_ = c.JSON(http.StatusInternalServerError, ErrInternalError)
		return
	}
	_ = c.JSON(_err.Status, GetHTTPError(_err))
}

//ErrMiddleware provide error middleware
func ErrMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			logFields := map[string]interface{}{}

			// 紀錄 Request 資料
			req := c.Request()
			{
				logFields["requestMethod"] = req.Method
				logFields["requestURL"] = req.URL.String()
			}

			// 紀錄 Response 資料
			resp := c.Response()
			resp.After(func() {
				logFields["responseStatus"] = resp.Status
				// 根據狀態碼用不同等級來紀錄
				logger := log.Ctx(req.Context()).With().Fields(logFields).Logger()
				if resp.Status >= http.StatusInternalServerError {
					logger.Error().Msgf("%+v", err)
				} else if resp.Status >= http.StatusBadRequest {
					logger.Debug().Msgf("%+v", err)
				} else {
					logger.Debug().Msgf("%+v", err)
				}
			})
		}
		return err
	}

}

// NotFoundHandlerForEcho responds not found response.
func NotFoundHandlerForEcho(c echo.Context) error {
	return c.JSON(http.StatusNotFound, ErrPageNotFound)
}

//UnaryAccessInterceptor ...
func UnaryAccessInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		go func() {
			reqB, _ := json.Marshal(req)
			reqStr := string(reqB)
			requestID := MetaFromContext(ctx)
			logger := log.With().Str("requestID", requestID).Str("input", reqStr).Str("endpoint", info.FullMethod).Logger()
			logger.Debug().Msg("Access log")
		}()
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func GQLErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	var (
		logger    = log.Ctx(ctx)
		unwrapErr error
		path      = graphql.GetPath(ctx)
	)

	gqlErr, ok := err.(*gqlerror.Error)
	if ok {
		unwrapErr = gqlErr.Unwrap()
	} else {
		if path != nil {
			gqlErr = gqlerror.WrapPath(path, err)
		} else {
			gqlErr = gqlerror.Errorf(err.Error())
		}
	}

	if unwrapErr == nil {
		unwrapErr = err
	}

	if gqlErr.Extensions == nil {
		gqlErr.Extensions = make(map[string]interface{})
	}

	causeErr := Cause(unwrapErr)
	_err, ok := causeErr.(*_error)
	if !ok || _err == nil {
		gqlErr.Extensions["raw_err_msg"] = unwrapErr.Error()
		_err = ErrInternalError
	}

	if _err.Details != nil {
		for k, v := range _err.Details {
			gqlErr.Extensions[k] = v
		}
	}

	gqlErr.Message = _err.Code
	gqlErr.Extensions["err_msg"] = _err.Message

	logger.Info().Msgf("error handle: %+v", unwrapErr)
	return gqlErr
}
