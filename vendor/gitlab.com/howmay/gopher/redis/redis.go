package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Config ...
type Config struct {
	ClusterMode     bool     `yaml:"cluster_mode" mapstructure:"cluster_mode"`
	Addresses       []string `yaml:"addresses" mapstructure:"addresses"`
	Password        string   `yaml:"password" mapstructure:"password"`
	MaxRetries      int      `yaml:"max_retries" mapstructure:"max_retries"`
	PoolSizePerNode int      `yaml:"pool_size_per_node" mapstructure:"pool_size_per_node"`
	DB              int      `yaml:"db" mapstructure:"db"`
}

// NewInjection ...
func (c Config) NewInjection() *Config {
	return &c
}

// Nil redis key not found
const Nil = redis.Nil

// Redis 提供操作 redis 的介面
type Redis interface {
	redis.Cmdable
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub

	// 客製化的函數

	// RegisterCasbinPubSub 註冊redis pubsub功能, 用來同步casbin polic
	RegisterCasbinPubSub(ctx context.Context, externalFunc func() error) error

	RedisLock(ctx context.Context, key, lockerID string, expireSeconds int) (bool, error)
	RedisUnlock(ctx context.Context, key, lockerID string) error
}

// InitRedisClient init redis client
func InitRedisClient(redisCfg *Config) (Redis, error) {
	if redisCfg.ClusterMode {
		return newClusterClient(redisCfg)
	}
	return newClient(redisCfg)
}
