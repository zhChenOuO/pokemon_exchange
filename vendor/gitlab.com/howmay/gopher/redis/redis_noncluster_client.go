package redis

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/cenk/backoff"
	"github.com/go-redis/redis/v8"
	goredis "github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"gitlab.com/howmay/gopher/errors"
)

// NonClusterClient ...
type NonClusterClient struct {
	*goredis.Client
}

func newClient(redisCfg *Config) (Redis, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	if len(redisCfg.Addresses) == 0 {
		return nil, fmt.Errorf("redis config address is empty")
	}

	var client *redis.Client
	err := backoff.Retry(func() error {
		client = redis.NewClient(&redis.Options{
			Addr:       redisCfg.Addresses[0],
			Password:   redisCfg.Password,
			MaxRetries: redisCfg.MaxRetries,
			PoolSize:   redisCfg.PoolSizePerNode,
			DB:         redisCfg.DB,
		})
		err := client.Ping(context.Background()).Err()
		if err != nil {
			return fmt.Errorf("ping occurs error after connecting to redis: %s", err)
		}
		return nil
	}, bo)

	if err != nil {
		return nil, err
	}

	return &NonClusterClient{Client: client}, nil
}

//RegisterCasbinPubSub 註冊redis pubsub功能, 用來同步casbin policy
func (c *NonClusterClient) RegisterCasbinPubSub(ctx context.Context, externalFunc func() error) error {
	pubsub := c.Subscribe(ctx, "Casbin")

	_, err := pubsub.Receive(ctx)
	if err != nil {
		return err
	}
	go func(externalFunc func() error) {
		defer pubsub.Close()
		for {
			_, err := pubsub.Receive(ctx)
			if err != nil {
				return
			}
			delay := rand.Intn(5)
			time.Sleep(time.Millisecond * 100 * time.Duration(delay))
			err = externalFunc()
			if err != nil {
				return
			}
			log.Info().Msg("reload policy: ")
		}
	}(externalFunc)

	return nil
}

// RedisLock ...
func (c *NonClusterClient) RedisLock(ctx context.Context, key, lockerID string, expireSeconds int) (bool, error) {
	ok, err := c.SetNX(ctx, key, lockerID, time.Duration(expireSeconds)*time.Second).Result() // 這邊timeout 時間要比http的timeout還久
	if err != nil {
		return false, errors.Wrapf(errors.ErrInternalError, "[RedisLock] SetNX Error: %v", err.Error())
	}
	if !ok { // 其他人佔用中
		return false, errors.Wrapf(errors.ErrResourceNotFound, "[RedisLock] SetNX with key %v already exists", key)
	}
	return true, nil
}

// RedisUnlock ...
func (c *NonClusterClient) RedisUnlock(ctx context.Context, key, lockerID string) error {
	result, err := c.Get(ctx, key).Result()
	if err != nil {
		return errors.Wrapf(errors.ErrInternalError, "[RedisUnlock] Get %v Error: %v", key, err.Error())
	}
	if result == lockerID { // 如果使用者還是自己，將鎖解開
		err := c.Del(ctx, key).Err()
		if err != nil {
			return errors.Wrapf(errors.ErrInternalError, "[RedisUnlock] Del %v Error: %v", key, err.Error())
		}
	}

	return nil
}
