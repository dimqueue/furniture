package api

import (
	_ "github.com/dmytroserhiienko02/furniture/api-svc/docs"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *apiGateway) router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAuthorization,
			echo.HeaderAccept,
			echo.HeaderXRequestedWith,
		},
		AllowMethods: []string{
			echo.POST,
			echo.DELETE,
			echo.GET,
			echo.PATCH,
			echo.OPTIONS,
		},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middlewares.CtxLog(a.log))
	e.Use(middlewares.CtxRepository(a.repository))

	user := e.Group("/users")
	{
		user.POST("", a.userService.CreateOrder)
		user.GET("/:status", a.userService.GetAllProductsByStatus)
		user.GET("", a.userService.GetAllProducts)
	}

	manager := e.Group("/managers", middlewares.WithUuidKey())
	{
		manager.POST("/products", a.managerService.CreateProduct)
		manager.GET("/orders", a.managerService.GetAllOrders)
		manager.GET("/orders/:status", a.managerService.GetOrdersByStatus)
	}
	admin := e.Group("/admins", middlewares.WithApiKey(a.listenerConfig.ApiKey))
	{
		admin.PATCH("/products/:id", a.adminService.UpdateProduct)
		admin.DELETE("/orders/:uuid", a.adminService.DeleteOrderById)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
