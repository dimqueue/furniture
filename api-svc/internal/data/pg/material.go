package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const materialTableName = "materials"

func newMaterialRepository(db *pgdb.DB) data.IMaterialRepository {
	return &materialRepository{
		db: db,
	}
}

type materialRepository struct {
	db *pgdb.DB
}

func (m materialRepository) GetMaterialById(id int64) (types.Material, error) {
	var material types.Material
	err := m.db.Get(&material, sq.Select("*").From(materialTableName).Where(sq.Eq{"id": id}))
	return material, errors.Wrap(err, "failed to select material by id")
}
