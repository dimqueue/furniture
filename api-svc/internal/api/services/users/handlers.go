package users

import (
	"net/http"

	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/middlewares"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services/managers"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewService() services.UserService {
	return &service{}
}

type service struct {
}

// GetAllProductsByStatus
// @Summary get products by status
// @Description returns products by status
// @Tags Users
// @Accept  application/json
// @Produce application/json
// @Param status   path string false "product status"
// @Success 200 {object} services.Response{data=[]managers.ProductView}
// @Failure 400 {object} services.Response
// @Failure 404 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /users/{status} [get]
func (s *service) GetAllProductsByStatus(ctx echo.Context) error {
	status, err := types.StatusTextCode(ctx.Param("status")).ToStatus()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}

	var response []managers.ProductView

	if err = middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {
		products, err := r.GetProductsByStatusId(status.ID)
		if err != nil {
			return errors.Wrap(err, "failed to get products by status id")
		}

		for _, p := range products {
			if p.StatusID != status.ID {
				return errors.New("incorrect status")
			}

			var productView managers.ProductView

			productView.Product = p

			productView.Status, err = productView.Product.StatusID.ToStatus()
			if err != nil {
				return errors.Wrap(err, "failed to parse status")
			}

			productView.Material, err = r.GetMaterialById(p.MaterialId)
			if err != nil {
				return errors.Wrap(err, "failed to get material")
			}

			response = append(response, productView)
		}

		return nil
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, services.NewSuccessResponseWithData(response))
}

// GetAllProducts
// @Summary get all products
// @Description returns all products
// @Tags Users
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} services.Response{data=[]managers.ProductView}
// @Failure 400 {object} services.Response
// @Failure 404 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /users [get]
func (s *service) GetAllProducts(ctx echo.Context) error {
	products, err := middlewares.Repository(ctx).GetAllProduct()
	var response []managers.ProductView
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}

	if err = middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {
		for _, p := range products {
			var productView managers.ProductView
			productView.Product = p
			productView.Status, err = productView.Product.StatusID.ToStatus()
			if err != nil {
				return errors.Wrap(err, "failed to parse status")
			}
			productView.Material, err = r.GetMaterialById(p.MaterialId)
			if err != nil {
				return errors.Wrap(err, "failed to get material")
			}
			response = append(response, productView)
		}
		return nil
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}
	return ctx.JSON(http.StatusOK, services.NewSuccessResponseWithData(response))
}

// CreateOrder
// @Summary create order
// @Description creates order
// @Tags Users
// @Accept  application/json
// @Produce application/json
// @Param   order body types.CreateOrder true "create order request"
// @Success 201 {object} services.Response{data=managers.OrderView}
// @Failure 404 {object} services.Response
// @Failure 500 {obj0ect} services.Response
// @Router /users [post]
func (s *service) CreateOrder(ctx echo.Context) error {
	var request types.CreateOrder

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}

	responseOrder := types.Order{
		CreateOrder: request,
		UUID:        uuid.New().String(),
		ManagerId:   nil,
	}
	var responseView managers.OrderView
	responseView.Order = responseOrder
	if err := middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {

		var err error

		responseView.ProductView.Product, err = r.GetProductById(request.ProductId)
		if err != nil {
			return errors.Wrap(err, "failed to get product by id")
		}

		if responseView.ProductView.Product.StatusID == types.StatusID(1) {
			return errors.New("failed to order product is locked")
		}

		responseView.ProductView.Product.StatusID = types.StatusID(1)

		responseView.ProductView.Status, err = responseView.ProductView.Product.StatusID.ToStatus()
		if err != nil {
			return errors.Wrap(err, "failed to parse status")
		}

		responseView.ProductView.Material, err = r.GetMaterialById(responseView.ProductView.Product.MaterialId)
		if err != nil {
			return errors.Wrap(err, "failed to get material by id")
		}

		_, err = r.CreateOrder(responseOrder)
		if err != nil {
			return errors.Wrap(err, "failed to create order")
		}

		_, err = r.UpdateStatus(responseView.ProductView.Product)
		if err != nil {
			return errors.Wrap(err, "failed to update order")
		}
		return nil
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, services.NewSuccessResponseWithData(responseView))
}
