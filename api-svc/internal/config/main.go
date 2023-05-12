package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	Listener
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	Listener
}

func New(getter kv.Getter) Config {
	return &config{
		Logger:    comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Databaser: pgdb.NewDatabaser(getter),
		Listener:  NewListener(getter),
	}
}
