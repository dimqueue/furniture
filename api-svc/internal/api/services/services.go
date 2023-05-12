package services

import (
	"github.com/labstack/echo/v4"
)

type UserService interface {
	GetAllProductsByStatus(ctx echo.Context) error
	CreateOrder(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
}

type ManagerService interface {
	GetOrdersByStatus(ctx echo.Context) error
	CreateProduct(ctx echo.Context) error
	GetAllOrders(ctx echo.Context) error
}

type AdminService interface {
	UpdateProduct(ctx echo.Context) error
	DeleteOrderById(ctx echo.Context) error
}
