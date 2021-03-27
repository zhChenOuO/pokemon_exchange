package errors

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"google.golang.org/grpc/metadata"
)

type ctxKey string

const (
	ctxKeyXRequestID ctxKey = echo.HeaderXRequestID
)

var key = "traceid"

// ContextWithXRequestID returns a context.Context with given X-Request-Id value.
func ContextWithMeta(ctx context.Context, xrid string) context.Context {
	if xrid == "" {
		xrid = uuid.New().String()
	}
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}
	md[key] = []string{xrid}
	return metadata.NewOutgoingContext(ctx, md)
}

// XRequestIDFromContext returns the X-Request-Id value from a given context.Context ctx.
func MetaFromContext(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)

	// Set request ID for context.
	requestIDs := md[key]
	if len(requestIDs) >= 1 {
		return requestIDs[0]
	}

	return ""
}

// ContextWithXRequestID returns a context.Context with given X-Request-Id value.
func ContextWithXRequestID(ctx context.Context, xrid string) context.Context {
	return context.WithValue(ctx, ctxKeyXRequestID, xrid)
}

// XRequestIDFromContext returns the X-Request-Id value from a given context.Context ctx.
func XRequestIDFromContext(ctx context.Context) string {
	v := ctx.Value(string(ctxKeyXRequestID))
	if v == nil {
		v = ctx.Value(ctxKeyXRequestID)
		if v == nil {
			return ""
		}
	}
	xrid, ok := v.(string)
	if !ok {
		return ""
	}
	return xrid
}
