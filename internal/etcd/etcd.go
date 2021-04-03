package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"
)

type Config struct {
	Endpoints []string `yaml:"endpoints" mapstructure:"endpoints"`
}

func New(lc fx.Lifecycle, cfg *Config) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{Endpoints: cfg.Endpoints})
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			defer cli.Close()
			return nil
		},
	})
	return cli, nil
}
