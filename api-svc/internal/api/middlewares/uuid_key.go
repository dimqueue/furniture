package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const uuidKeyHeader = "X-UUID-Key"

var (
	errFailedToCheckExistence = errors.New("failed to check existence")
	errManagerDoesNotExist    = errors.New("manager does not exist")
)

func WithUuidKey() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: fmt.Sprintf("header:%s", uuidKeyHeader),
		Validator: func(key string, ctx echo.Context) (bool, error) {
			exists, err := Repository(ctx).ManagerExists(key)
			if err != nil {
				return false, errFailedToCheckExistence
			}

			if !exists {
				return false, errManagerDoesNotExist
			}

			return true, nil
		},
		ErrorHandler: func(err error, ctx echo.Context) error {
			return ctx.JSON(http.StatusUnauthorized, services.NewErrorResponse(err))
		},
	})
}
