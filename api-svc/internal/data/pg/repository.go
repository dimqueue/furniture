package pg

import (
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func NewRepository(db *pgdb.DB) data.IRepository {
	return &repository{
		IOrderRepository:    newOrderRepository(db),
		IManagerRepository:  newManagerRepository(db),
		IProductRepository:  newProductRepository(db),
		IMaterialRepository: newMaterialRepository(db),
		db:                  db,
	}
}

type repository struct {
	data.IOrderRepository
	data.IManagerRepository
	data.IProductRepository
	data.IMaterialRepository
	db *pgdb.DB
}

func (r repository) Transaction(fn func(r data.IRepository) error) error {
	return r.db.Transaction(func() error {
		return fn(r)
	})
}
