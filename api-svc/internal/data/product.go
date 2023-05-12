package data

import "github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"

type IProductRepository interface {
	CreateProduct(product types.Product) (types.Product, error)
	GetAllProduct() ([]types.Product, error)
	GetProductsByStatusId(types.StatusID) ([]types.Product, error)
	GetProductById(int64) (types.Product, error)
	UpdateStatus(types.Product) (types.Product, error)
	UpdateProduct(manager types.UpdateProduct) (types.Product, error)
}
