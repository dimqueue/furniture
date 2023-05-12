package types

type CreateOrder struct {
	Description string `db:"description" json:"description"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Delivery    string `db:"delivery" json:"delivery"`
	OrderEmail  string `db:"order_email" json:"order_email"`
	ProductId   int64  `db:"product_id" json:"product_id"`
}
type Order struct {
	UUID      string  `db:"uuid" json:"uuid"`
	ManagerId *string `db:"manager_id" json:"manager_id"`
	CreateOrder
}
