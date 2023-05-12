package data

import "github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"

type IManagerRepository interface {
	ManagerExists(uuid string) (bool, error)
	CreateManager(manager types.Manager) (types.Manager, error)
	DeleteManager(uuid string) error
	GetManagerByUUID(uuid string) (types.Manager, error)
	GetAllManagers() ([]types.Manager, error)
}
