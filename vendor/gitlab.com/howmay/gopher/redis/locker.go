package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// Locker 提供鎖的相關功能
type Locker interface {
	// Lock 上鎖，並取得對應之解鎖函式，該函式會關閉Locker的Redis連線
	Lock(ctx context.Context, namespace string, targetID uint) (ok bool, unlock func())
	// DoWithLock 直到上鎖才執行 do 的內容
	DoWithLock(ctx context.Context, namespace string, targetID uint, taskInterval time.Duration, doTask func())
}

type locker struct {
	ctx         context.Context
	RedisClient Redis
	LockerID    string
}

// NewLocker new a locker with usedID. (It means who lock the key.)
func NewLocker(ctx context.Context, redisClient Redis, lockerID string) Locker {
	return &locker{
		ctx:         ctx,
		RedisClient: redisClient,
		LockerID:    lockerID,
	}
}

// Lock 上鎖，並取得對應之解鎖函式，該函式會關閉Locker的Redis連線
func (l *locker) Lock(ctx context.Context, namespace string, targetID uint) (bool, func()) {
	logger := log.Ctx(l.ctx)
	unlock := func() {}

	key := l.getKey(namespace, targetID)
	ok, err := l.RedisClient.RedisLock(ctx, key, l.LockerID, 90) // 這邊timeout 時間要比http的timeout還久,先寫死90秒
	if !ok {
		if err != nil {
			logger.Warn().Msgf(err.Error())
		}
		return false, unlock
	}

	unlock = func() {
		err := l.RedisClient.RedisUnlock(ctx, key, l.LockerID)
		if err != nil {
			logger.Warn().Msgf(err.Error())
		}
		return
	}
	return true, unlock
}

// DoWithLock 直到上鎖才執行 do 的內容
func (l *locker) DoWithLock(ctx context.Context, namespace string, targetID uint, taskInterval time.Duration, doTask func()) {
	key := l.getKey(namespace, targetID)

	for {
		result, err := l.RedisClient.RedisLock(ctx, key, l.LockerID, 90) // 這邊timeout 時間要比http的timeout還久,先寫死90秒
		if result == true && err == nil {
			defer l.RedisClient.RedisUnlock(ctx, key, l.LockerID)
			doTask()
			return
		}
		time.Sleep(taskInterval)
	}
}

func (l *locker) getKey(namespace string, targetID uint) string {
	return fmt.Sprintf("%s%s", namespace, strconv.FormatUint(uint64(targetID), 10))
}
