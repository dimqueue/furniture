package data

type IRepository interface {
	IOrderRepository
	IManagerRepository
	IProductRepository
	IMaterialRepository

	Transaction(func(r IRepository) error) error
}
