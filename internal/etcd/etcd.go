package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/fx"
)

type Config struct {
	Endpoints []string `yaml:"endpoints" mapstructure:"endpoints"`
}

func New(lc fx.Lifecycle, cfg *Config) (*concurrency.Session, error) {
	cli, err := clientv3.New(clientv3.Config{Endpoints: cfg.Endpoints})
	if err != nil {
		return nil, err
	}
	s, err := concurrency.NewSession(cli)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			defer cli.Close()
			defer s.Close()

			return nil
		},
	})
	return s, nil
}
