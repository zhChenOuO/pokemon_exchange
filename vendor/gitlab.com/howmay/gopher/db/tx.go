package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/errors"
	"gorm.io/gorm"
)

// CtxKey 用來代表 context.Context 的 key
type CtxKey string

// AfterCommitCallback ...
type AfterCommitCallback func() error

type contextKey struct {
	name string
}

const (
	// CtxKeyAfterCommitCallback 註冊 transaction commit 後要執行的動作的 context key
	CtxKeyAfterCommitCallback CtxKey = "afterCommit"
	// DefaultCostMilliSeconds 預設的db預期的執行時間，超過會出警示
	DefaultCostMilliSeconds int64 = 2000

	// DefaultAfterCommitFuncsCostMilliSeconds 預設的db after commit func預期的執行時間，超過會出警示
	DefaultAfterCommitFuncsCostMilliSeconds int64 = 2000
)

// ExecuteTx 執行一個 Database 交易，如果 `fn` 執行過程中遇到失敗，會自動執行 rollback, 如果成功則會自動 commit,
// 另外需要注意的地方是 db 的 connection 需要用 write 的 connection 來執行 tx
func ExecuteTx(ctx context.Context, db *gorm.DB, exceptCostMilliSeconds int64, fn func(*gorm.DB) error) error {
	var err error
	if exceptCostMilliSeconds == 0 {
		exceptCostMilliSeconds = DefaultCostMilliSeconds
	}
	f := GetFrame(1)
	logger := log.Ctx(ctx).With().
		Str("caller", fmt.Sprintf("%s:%v1", f.File, f.Line)).
		Str("caller_func", fmt.Sprintf("%s", f.Function)).Logger()
	defer func(now time.Time) {
		costMilliSeconds := time.Since(now).Milliseconds()
		logger = logger.With().Int64("cost_milliseconds", costMilliSeconds).Logger()
		if costMilliSeconds > exceptCostMilliSeconds {
			logger.Warn().Msgf("db: transaction time too long: %v1 millsSeconds", costMilliSeconds)
		}
		if err == nil { // trigger after commit
			runAfterCommitCallback(logger.WithContext(ctx))
		}
	}(time.Now())
	// Start a transactions.
	err = executeInTx(ctx, db, fn)
	if err != nil {
		return err
	}

	return nil
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
	}
	panicked = false
	return err
}

// ExecuteTxWithSavePoint 支援SavePoint的tx func
func ExecuteTxWithSavePoint(ctx context.Context, db *gorm.DB, exceptCostMilliSeconds int64, fn func(*gorm.DB) (rollbackToSavepoint string, err error)) error {
	var err error
	if exceptCostMilliSeconds == 0 {
		exceptCostMilliSeconds = DefaultCostMilliSeconds
	}
	f := GetFrame(1)
	logger := log.Ctx(ctx).With().
		Str("caller", fmt.Sprintf("%s:%v1", f.File, f.Line)).
		Str("caller_func", fmt.Sprintf("%s", f.Function)).Logger()
	defer func(now time.Time) {
		costMilliSeconds := time.Since(now).Milliseconds()
		logger = logger.With().Int64("cost_milliseconds", costMilliSeconds).Logger()
		if costMilliSeconds > exceptCostMilliSeconds {
			logger.Warn().Msgf("db: transaction time too long: %v1 millsSeconds", costMilliSeconds)
		}
		if err == nil { // trigger after commit
			runAfterCommitCallback(logger.WithContext(ctx))
		}
	}(time.Now())
	// Start a transactions which driver support Savepoint.
	err = executeInTxWithSavePoint(ctx, db, fn)
	if err != nil {
		return err
	}
	return nil
}

func executeInTxWithSavePoint(ctx context.Context, db *gorm.DB, fn func(*gorm.DB) (rollbackToSavepoint string, txErr error)) (err error) {
	panicked := true
	rollbackToSavepoint := ""
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
			if rollbackToSavepoint != "" {
				rollbackErr := tx.RollbackTo(rollbackToSavepoint).Error
				if rollbackErr == nil {
					tx.Commit()
					return
				} else if rollbackErr != gorm.ErrUnsupportedDriver {
					log.Ctx(ctx).Error().Msgf("executeInTx RollbackTo %v transaction then error again! error: %+v ", rollbackToSavepoint, rollbackErr)
					return
				}
				log.Ctx(ctx).Error().Msgf("this driver not support rollback with savepoint. please refine your code. err: %v", rollbackErr)
			}
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				log.Ctx(ctx).Error().Msgf("executeInTx Rollback transaction then error again! error: %+v ", rollbackErr)
				return
			}

			return
		}
	}()

	rollbackToSavepoint, err = fn(tx)

	if err == nil {
		if commitErr := tx.Commit().Error; commitErr != nil {
			log.Ctx(ctx).Error().Msgf("executeInTx failed to commit, start Rollback transaction. error: %+v", err)
			return commitErr
		}
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
func AddAfterCommitCallback(ctx context.Context, fn AfterCommitCallback) (context.Context, error) {
	logger := log.Ctx(ctx).With().Logger()
	if fn == nil {
		logger.Warn().Msg("AddAfterCommitCallback fn can not be nil")
		return ctx, nil
	}
	callbacks := ctx.Value(CtxKeyAfterCommitCallback)
	if callbacks == nil {
		return ctx, errors.NewWithMessagef(errors.ErrInternalServerError, "AddAfterCommitCallback failed, not set CtxKey")
	}
	fns, ok := callbacks.(*[]AfterCommitCallback)
	if !ok {
		logger.Warn().Msg("AddAfterCommitCallback convert failed")
		return ctx, errors.NewWithMessagef(errors.ErrInternalServerError, "AddAfterCommitCallback convert failed, the callbacks funcs unexpected")
	}
	*fns = append(*fns, fn)
	return context.WithValue(ctx, CtxKeyAfterCommitCallback, fns), nil
}

func runAfterCommitCallback(ctx context.Context) {
	var err error
	logger := log.Ctx(ctx).With().Logger()
	defer func(now time.Time) {
		costMilliSeconds := time.Since(now).Milliseconds()
		logger = logger.With().Int64("db: runAfterCommitCallback cost_milliseconds", costMilliSeconds).Logger()
		if costMilliSeconds > DefaultAfterCommitFuncsCostMilliSeconds {
			logger.Warn().Msgf("db: runAfterCommitCallback too long: %v1 millsSeconds", costMilliSeconds)
		}
	}(time.Now())
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
		if err = fn(); err != nil {
			logger.Error().Msgf("fail to run after commit callback, err: %s", err.Error())
		}
	}
	*fns = []AfterCommitCallback{}
}
