package cli

import (
	"os"

	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/config"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/getter"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var errWrongCommandArguments = errors.New("wrong command arguments")

func Run() error {
	defer func() {
		if rvr := recover(); rvr != nil {
			logan.New().WithRecover(rvr).Error("app panicked")
		}
	}()

	gttr := getter.NewGetter()
	cfg := config.New(gttr)
	log := cfg.Log()

	if len(os.Args[1:]) != 1 {
		return errWrongCommandArguments
	}

	switch os.Args[1] {
	case "migrates-up":
		if err := MigrateUp(cfg); err != nil {
			log.Error("failed to migrates-up")
			return errors.Wrap(err, "failed to migrates-up")
		}
	case "migrates-down":
		if err := MigrateDown(cfg); err != nil {
			log.Error("failed to migrates-down")
			return errors.Wrap(err, "failed to migrates-down")
		}
	case "run-server":
		api.Run(cfg)
	default:
		return errors.New("unexpected command argument")
	}

	return nil
}
