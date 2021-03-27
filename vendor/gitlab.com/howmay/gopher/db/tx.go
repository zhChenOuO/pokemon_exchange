package db

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
)

// CtxKey 用來代表 context.Context 的 key
type CtxKey string

// AfterCommitCallback ...
type AfterCommitCallback func()

type contextKey struct {
	name string
}

const (
	// CtxKeyAfterCommitCallback 註冊 transaction commit 後要執行的動作的 context key
	CtxKeyAfterCommitCallback CtxKey = "afterCommit"
)

// ExecuteTx 執行一個 Database 交易，如果 `fn` 執行過程中遇到失敗，會自動執行 rollback, 如果成功則會自動 commit,
// 另外需要注意的地方是 db 的 connection 需要用 write 的 connection 來執行 tx
func ExecuteTx(ctx context.Context, db *gorm.DB, fn func(*gorm.DB) error) error {
	f := GetFrame(1)
	logger := log.Ctx(ctx).With().
		Str("caller", fmt.Sprintf("%s:%v1", f.File, f.Line)).
		Str("caller_func", fmt.Sprintf("%s", f.Function)).Logger()
	defer func(now time.Time) {
		costMilliSeconds := time.Since(now).Milliseconds()
		logger = logger.With().Int64("cost_milliseconds", costMilliSeconds).Logger()
		if costMilliSeconds > 2000 {
			logger.Warn().Msgf("db: transaction time too long: %v1 millsSeconds", costMilliSeconds)
		}
	}(time.Now())
	// Start a transactions.
	return executeInTx(ctx, db, fn)
}

func executeInTx(ctx context.Context, db *gorm.DB, fn func(*gorm.DB) error) (err error) {
	panicked := true
	tx := db.Begin()
	if tx.Error != nil {
		return errors.Wrapf(errors.ErrInternalError, "executeInTx Begin error %+v", tx.Error.Error())
	}

	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked {
			err = errors.Wrap(errors.ErrInternalError, "executeInTx occurs panic, start Rollback transaction")
		} else if err != nil {
			err = errors.Wrapf(err, "executeInTx occurs error, start Rollback transaction.")
		}

		if err != nil {
			log.Ctx(ctx).Error().Msgf("%s", err)
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				log.Ctx(ctx).Error().Msgf("executeInTx Rollback transaction then error again! error: %+v ", rollbackErr)
				return
			}
			return
		}
	}()

	err = fn(tx)

	if err == nil {
		if commitErr := tx.Commit().Error; commitErr != nil {
			log.Ctx(ctx).Error().Msgf("executeInTx failed to commit, start Rollback transaction. error: %+v", err)
			return commitErr
		}
		runAfterCommitCallback(ctx)
	}
	panicked = false
	return err
}

// InitAfterTxCommitCallback init after tx commit callback callback list
func InitAfterTxCommitCallback(ctx context.Context) context.Context {
	afterCommitFunc := []AfterCommitCallback{}
	return context.WithValue(ctx, CtxKeyAfterCommitCallback, &afterCommitFunc)
}

// AddAfterCommitCallback ...
func AddAfterCommitCallback(ctx context.Context, fn AfterCommitCallback) context.Context {
	logger := log.Ctx(ctx).With().Logger()
	if fn == nil {
		logger.Warn().Msg("AddAfterCommitCallback fn can not be nil")
		return ctx
	}
	callbacks := ctx.Value(CtxKeyAfterCommitCallback)
	fns, ok := callbacks.(*[]AfterCommitCallback)
	if !ok {
		logger.Warn().Msg("AddAfterCommitCallback convert failed")
		return ctx
	}
	*fns = append(*fns, fn)
	return context.WithValue(ctx, CtxKeyAfterCommitCallback, fns)
}

func runAfterCommitCallback(ctx context.Context) {
	logger := log.Ctx(ctx).With().Logger()
	callbacks := ctx.Value(CtxKeyAfterCommitCallback)
	if callbacks == nil {
		return
	}

	fns, ok := callbacks.(*[]AfterCommitCallback)
	if !ok {
		logger.Warn().Msg("runAfterCommitCallback convert failed")
		return
	}
	for i := range *fns {
		fn := (*fns)[i]
		fn()
	}
	*fns = []AfterCommitCallback{}
}

// GetFrame get caller frame
// example:
// f := stack.GetFrame(0) // 0: currentFrame
// fmt.Println(f.File, f.Function, f.Line)
func GetFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
