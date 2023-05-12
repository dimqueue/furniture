package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const managerTableName = "managers"

func newManagerRepository(db *pgdb.DB) data.IManagerRepository {
	return &managerRepository{
		db: db,
	}
}

type managerRepository struct {
	db *pgdb.DB
}

func (m managerRepository) ManagerExists(uuid string) (bool, error) {
	var count int64

	if err := m.db.Get(&count, sq.Select("count(*)").
		From(managerTableName).
		Where(sq.Eq{"uuid": uuid})); err != nil {
		return false, errors.Wrap(err, "failed to check existence")
	}

	return count > 0, nil
}

func (m managerRepository) CreateManager(manager types.Manager) (types.Manager, error) {
	var returning types.Manager

	err := m.db.Get(&returning, sq.
		Insert(managerTableName).
		Columns("uuid", "login", "manager_email").
		Values(manager.UUID, manager.Login, manager.ManagerEmail).
		Suffix("returning *"))

	return returning, errors.Wrap(err, "failed to create manager")
}

func (m managerRepository) DeleteManager(uuid string) error {
	return m.db.Exec(sq.Delete(managerTableName).Where(sq.Eq{"uuid": uuid}))
}

func (m managerRepository) GetManagerByUUID(uuid string) (types.Manager, error) {
	var manager types.Manager

	err := m.db.Get(&manager, sq.
		Select("uuid", "username", "manager_email").
		From(managerTableName).
		Where(sq.Eq{"uuid": uuid}))

	return manager, errors.Wrap(err, "failed to get manager by uuid")
}

func (m managerRepository) GetAllManagers() ([]types.Manager, error) {
	var managers []types.Manager

	err := m.db.Select(&managers, sq.Select("*").From(managerTableName))

	return managers, errors.Wrap(err, "failed to get all managers")
}
