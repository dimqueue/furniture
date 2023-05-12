package admins

import (
	"net/http"
	"strconv"

	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/middlewares"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"
	"github.com/labstack/echo/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewService() services.AdminService {
	return &service{}
}

type service struct {
}

// UpdateProduct
// @Summary update product
// @Description updates product
// @Tags Admins
// @Accept  application/json
// @Produce application/json
// @Param X-Api-Key header string true "admin`s X-Api-Key"
// @Param id   path string true "product id"
// @Param body   body types.UpdateProduct true "update product body"
// @Success 200 {object} services.Response{data=types.Product}
// @Failure 400 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /admins/products/{id} [patch]
func (s service) UpdateProduct(ctx echo.Context) error {
	var request types.UpdateProduct

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, services.NewErrorResponse(err))
	}
	var err error

	request.Id, err = strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse int")
	}
	var response types.Product
	err = middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {
		var err error
		response, err = r.UpdateProduct(request)
		if err != nil {
			return errors.Wrap(err, "failed to  update product")
		}

		return nil
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, services.NewSuccessResponseWithData(response))
}

// DeleteOrderById
// @Summary delete order by id
// @Description deletes product by id
// @Tags Admins
// @Accept  application/json
// @Produce application/json
// @Param X-Api-Key header string true "admin`s X-Api-Key"
// @Param uuid   path string true "order uuid"
// @Success 204
// @Failure 400 {object} services.Response
// @Failure 500 {object} services.Response
// @Router /admins/orders/{uuid} [delete]
func (s *service) DeleteOrderById(ctx echo.Context) error {
	uuid := ctx.Param("uuid")
	var exists bool
	err := middlewares.Repository(ctx).Transaction(func(r data.IRepository) error {
		var err error
		exists, err = r.OrderExists(uuid)
		if err != nil {
			return errors.Wrap(err, "failed to check order existence")
		}
		if !exists {
			return nil
		}

		if err = r.DeleteOrderById(uuid); err != nil {
			return errors.Wrap(err, "failed to delete order")
		}

		return nil
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, services.NewErrorResponse(err))
	}
	if !exists {
		return ctx.JSON(http.StatusNotFound, services.NewDefaultResponse(http.StatusNotFound))
	}

	return ctx.NoContent(http.StatusNoContent)
}
