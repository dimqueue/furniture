package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const orderTableName = "orders"

func newOrderRepository(db *pgdb.DB) data.IOrderRepository {
	return &orderRepository{
		db: db,
	}
}

type orderRepository struct {
	db *pgdb.DB
}

func (o orderRepository) OrderExists(uuid string) (bool, error) {
	var count int64

	err := o.db.Get(&count, sq.Select("count(*)").
		From(orderTableName).
		Where(sq.Eq{"uuid": uuid}))

	return count > 0, errors.Wrap(err, "failed to check existence order")
}

func (o orderRepository) CreateOrder(order types.Order) (types.Order, error) {
	var returning types.Order

	err := o.db.Get(&returning, sq.
		Insert(orderTableName).
		Columns("uuid", "description", "first_name", "last_name", "delivery", "order_email", "product_id").
		Values(order.UUID, order.Description, order.FirstName, order.LastName, order.Delivery, order.OrderEmail, order.ProductId).
		Suffix("returning *"))

	return returning, errors.Wrap(err, "failed to create order")
}

func (o orderRepository) GetOrderByUUID(uuid string) (types.Order, error) {
	var order types.Order

	err := o.db.Get(&order, sq.Select("*").From(orderTableName).Where(sq.Eq{"uuid": uuid}))

	return order, errors.Wrap(err, "failed to get order by uuid")
}

func (o orderRepository) GetAllOrders() ([]types.Order, error) {
	var orders []types.Order

	err := o.db.Select(&orders, sq.Select("*").From(orderTableName))

	return orders, errors.Wrap(err, "failed to get all orders")
}

func (o orderRepository) GetOrderByProductId(id int64) (types.Order, error) {
	var order types.Order
	err := o.db.Get(&order, sq.Select("*").From(orderTableName).Where(sq.Eq{"product_id": id}))
	return order, errors.Wrap(err, "failed to get order by product id")
}
func (o orderRepository) DeleteOrderById(uuid string) error {
	return o.db.Exec(sq.Delete(orderTableName).Where(sq.Eq{"uuid": uuid}))
}
