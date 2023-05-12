package api

import (
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services/admins"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services/managers"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services/users"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/config"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/pg"
	"gitlab.com/distributed_lab/logan/v3"
)

type apiGateway struct {
	log            *logan.Entry
	listenerConfig config.ListenerConfig
	repository     data.IRepository
	userService    services.UserService
	managerService services.ManagerService
	adminService   services.AdminService
}

func (a *apiGateway) run() error {
	return a.router().Start(a.listenerConfig.Addr)
}

func newService(cfg config.Config) *apiGateway {
	return &apiGateway{
		log:            cfg.Log(),
		listenerConfig: cfg.ListenerConfig(),
		repository:     pg.NewRepository(cfg.DB()),
		userService:    users.NewService(),
		managerService: managers.NewService(),
		adminService:   admins.NewService(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
