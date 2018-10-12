package resolver

import (
	"log"

	"github.com/boromisa/dcdr/cli/api/stores"
	"github.com/boromisa/dcdr/cli/api/stores/consul"
	"github.com/boromisa/dcdr/cli/api/stores/etcd"
	"github.com/boromisa/dcdr/cli/api/stores/redis"
	"github.com/boromisa/dcdr/config"
)

func LoadStore(cfg *config.Config) stores.IFace {
	switch cfg.Storage {
	case "etcd":
		return etcd.New(cfg)
	case "redis":
		r, err := redis.New(cfg)

		if err != nil {
			log.Fatalf("could not load redis: %v", err)
		}

		return r
	default:
		c, err := consul.NewDefault(cfg)

		if err != nil {
			log.Fatalf("could not load consul: %v", err)
		}

		return c
	}
}
