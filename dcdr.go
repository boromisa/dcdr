package main

import (
	"os"

	"github.com/PagerDuty/godspeed"
	"github.com/boromisa/dcdr/cli"
	"github.com/boromisa/dcdr/cli/api"
	"github.com/boromisa/dcdr/cli/api/resolver"
	"github.com/boromisa/dcdr/cli/controller"
	"github.com/boromisa/dcdr/cli/printer"
	"github.com/boromisa/dcdr/cli/repo"
	"github.com/boromisa/dcdr/config"
)

func main() {
	cfg := config.LoadConfig()
	store := resolver.LoadStore(cfg)

	rp := repo.New(cfg)

	var gs *godspeed.Godspeed
	var err error
	if cfg.StatsEnabled() {
		gs, err = godspeed.New(cfg.Stats.Host, cfg.Stats.Port, false)

		if err != nil {
			printer.SayErr("%v", err)
			os.Exit(1)
		}
	}

	kv := api.New(store, rp, cfg, gs)
	ctrl := controller.New(cfg, kv)

	dcdr := cli.New(ctrl)
	dcdr.Run()
}
