package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/keypair/figurekeypair"
)

const listenerMapId = "listener"

type Listener interface {
	ListenerConfig() ListenerConfig
}

type ListenerConfig struct {
	Addr   string `fig:"addr,required"`
	ApiKey string `fig:"api_key,required"`
}

type listener struct {
	once   comfig.Once
	getter kv.Getter
}

func (l *listener) ListenerConfig() ListenerConfig {
	return l.once.Do(func() interface{} {
		var cfg ListenerConfig
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(l.getter, listenerMapId)).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out listener"))
		}

		return cfg
	}).(ListenerConfig)
}

func NewListener(getter kv.Getter) Listener {
	return &listener{
		getter: getter,
	}
}
