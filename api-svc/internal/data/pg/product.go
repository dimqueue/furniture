package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const productTableName = "products"

func newProductRepository(db *pgdb.DB) data.IProductRepository {
	return &productRepository{
		db: db,
	}
}

type productRepository struct {
	db *pgdb.DB
}

func (p productRepository) CreateProduct(product types.Product) (types.Product, error) {
	var returning types.Product

	err := p.db.Get(&returning, sq.
		Insert(productTableName).
		Columns("title", "price", "material_id", "status_id").
		Values(product.Title, product.Price, product.MaterialId, product.StatusID).
		Suffix("returning *"))

	if err != nil {
		return types.Product{}, errors.Wrap(err, "failed to create product")
	}

	return returning, nil
}

func (p productRepository) GetAllProduct() ([]types.Product, error) {
	var products []types.Product
	err := p.db.Select(&products, sq.Select("*").From(productTableName))
	return products, errors.Wrap(err, "failed to select all products")
}

func (p productRepository) GetProductsByStatusId(statusID types.StatusID) ([]types.Product, error) {
	var products []types.Product

	err := p.db.Select(&products, sq.Select("*").From(productTableName).Where(sq.Eq{"status_id": statusID}))

	return products, errors.Wrap(err, "failed to select products by status")
}

func (p productRepository) GetProductById(id int64) (types.Product, error) {
	var product types.Product
	err := p.db.Get(&product, sq.Select("*").From(productTableName).Where(sq.Eq{"id": id}))
	return product, errors.Wrap(err, "failed to select product by id")
}

func (p productRepository) UpdateStatus(product types.Product) (types.Product, error) {
	var returning types.Product
	err := p.db.Get(&returning, sq.Update(productTableName).Set("status_id", product.StatusID).Where(sq.Eq{"id": product.Id}).Suffix("returning *"))
	return returning, errors.Wrap(err, "failed to update status")
}

func (p productRepository) UpdateProduct(up types.UpdateProduct) (types.Product, error) {
	var returning types.Product
	err := p.db.Get(&returning, sq.Update(productTableName).
		SetMap(up.SetData()).
		Where(sq.Eq{"id": up.Id}).
		Suffix("returning *"))

	return returning, errors.Wrap(err, "failed to update product")
}
