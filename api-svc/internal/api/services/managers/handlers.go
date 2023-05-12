package managers

import (
	"net/http"

	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/middlewares"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"github.com/labstack/echo/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewService() services.ManagerService {
	return &service{}
}

type service struct {
}

// CreateProduct
// @Summary create product
// @Description creates product
// @Tags Managers
// @Accept  application/json
// @Produce application/json
// @Param X-UUID-Key header string true "managers`s api key"
// @Param manager body types.CreateProduct true "create product body"
// @Success 201 {object} services.Response{data=ProductView}
// @Failure 400 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /managers/products [post]
func (s *service) CreateProduct(ctx echo.Context) error {
	var request types.CreateProduct
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}
	var (
		product types.Product
		err     error
	)
	product.CreateProduct = request
	product.StatusID = 2

	product, err = middlewares.Repository(ctx).CreateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}
	return ctx.JSON(http.StatusCreated, services.NewSuccessResponseWithData(product))
}

// GetAllOrders
// @Summary get all orders
// @Description returns all orders
// @Tags Managers
// @Accept  application/json
// @Produce application/json
// @Param X-UUID-Key header string true "managers`s api key"
// @Success 201 {object} services.Response{data=[]OrderView}
// @Failure 400 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /managers/orders [get]
func (s service) GetAllOrders(ctx echo.Context) error {
	orders, err := middlewares.Repository(ctx).GetAllOrders()
	var response []OrderView
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}

	if err := middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {

		for _, o := range orders {
			var orderView OrderView
			orderView.Order = o
			orderView.ProductView.Product, err = r.GetProductById(o.ProductId)
			if err != nil {
				return errors.Wrap(err, "failed to get product by id")
			}

			orderView.ProductView.Status, err = orderView.ProductView.Product.StatusID.ToStatus()
			if err != nil {
				return errors.Wrap(err, "failed to parse status")
			}

			orderView.ProductView.Material, err = r.GetMaterialById(orderView.ProductView.Product.MaterialId)
			if err != nil {
				return errors.Wrap(err, "failed to get material by id")
			}
			response = append(response, orderView)
		}

		return nil
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, services.NewSuccessResponseWithData(response))
}

// GetOrdersByStatus
// @Summary get orders by product status
// @Description returns orders by  product status
// @Tags Managers
// @Accept  application/json
// @Produce application/json
// @Param X-UUID-Key header string true "managers`s api key"
// @Param status   path string false "product status"
// @Success 200 {object} services.Response{data=[]OrderView}
// @Failure 400 {object} services.Response
// @Failure 404 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /managers/orders/{status} [get]
func (s service) GetOrdersByStatus(ctx echo.Context) error {
	status, err := types.StatusTextCode(ctx.Param("status")).ToStatus()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}
	var response []OrderView
	if err := middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {
		products, err := r.GetProductsByStatusId(status.ID)
		for _, p := range products {

			var orderView OrderView

			orderView.Order, err = r.GetOrderByProductId(p.Id)
			if err != nil {
				return errors.Wrap(err, "failed to get order by product id")
			}

			orderView.ProductView.Product = p

			orderView.ProductView.Status, err = orderView.ProductView.Product.StatusID.ToStatus()
			if err != nil {
				return errors.Wrap(err, "failed to parse status")
			}

			orderView.ProductView.Material, err = r.GetMaterialById(orderView.ProductView.Product.MaterialId)
			if err != nil {
				return errors.Wrap(err, "failed to get material by id")
			}
			response = append(response, orderView)
		}

		return nil
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, services.NewSuccessResponseWithData(response))
}
