package middleware

// import (
// 	"context"

// 	"github.com/rs/zerolog/log"
// )

// type loggerFieldsKey struct{}

// // InitLoggerFields 透過這個來 init Logger Fileds
// func InitLoggerFields(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, loggerFieldsKey{}, &log.Fields{})
// }

// // Logger 拿出 WithFieldsTags 的 logger
// func Logger(ctx context.Context) log.Entry {
// 	logger := log.FromContext(ctx)
// 	ctxFields := ctx.Value(loggerFieldsKey{}).(log.Fields)
// 	return logger.WithFields(ctxFields)
// }

// // SetLoggerFields 透過這個來塞 loggerFields
// func SetLoggerFields(ctx context.Context, fields log.Fields) {
// 	ctxFields := ctx.Value(loggerFieldsKey{}).(log.Fields)
// 	for i := range fields {
// 		(ctxFields)[i] = fields[i]
// 	}
// }
