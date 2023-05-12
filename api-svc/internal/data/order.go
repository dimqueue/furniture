package data

import "github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"

type IOrderRepository interface {
	OrderExists(uuid string) (bool, error)
	GetAllOrders() ([]types.Order, error)
	CreateOrder(order types.Order) (types.Order, error)
	GetOrderByUUID(uuid string) (types.Order, error)
	GetOrderByProductId(id int64) (types.Order, error)
	DeleteOrderById(string) error
}
