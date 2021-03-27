package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type ctxKeyAfterMiddlewareFunc struct{}

// AddFuncAfterMiddleware 加入當 middleware 結束時去執行的 function
func AddFuncAfterMiddleware(ctx context.Context, fn func()) {
	funcList := ctx.Value(ctxKeyAfterMiddlewareFunc{}).(*[]func())
	*funcList = append(*funcList, fn)
}

// ScheduleCallAddFuncAfterMiddleware 給 schedule 用的 function
func ScheduleCallAddFuncAfterMiddleware(ctx context.Context) {
	funcList := ctx.Value(ctxKeyAfterMiddlewareFunc{}).(*[]func())
	for i := range *funcList {
		(*funcList)[i]()
	}
}

// ScheduleInitAddFuncAfterMiddleware 給 schedule 用的 init function
func ScheduleInitAddFuncAfterMiddleware(ctx context.Context) context.Context {
	funcList := []func(){}
	return context.WithValue(ctx, ctxKeyAfterMiddlewareFunc{}, &funcList)
}

// AfterMiddleware 設定預設 after middleware 要執行的 function
func AfterMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()
			ctx = ScheduleInitAddFuncAfterMiddleware(ctx)
			c.SetRequest(c.Request().WithContext(ctx))
			defer func() {
				ctx := c.Request().Context()
				ScheduleCallAddFuncAfterMiddleware(ctx)
			}()

			return next(c)
		}
	}
}

// GinAfterMiddleware 設定預設 after middleware 要執行的 function
func GinAfterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = ScheduleInitAddFuncAfterMiddleware(ctx)
		c.Request = c.Request.WithContext(ctx)
		defer func() {
			ctx := c.Request.Context()
			ScheduleCallAddFuncAfterMiddleware(ctx)
		}()
		c.Next()
	}
}
