package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dmytroserhiienko02/furniture/api-svc/internal/api/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const apiKeyHeader = "X-Api-Key"

var (
	errWrongApiKeyProvided = errors.New("wrong api key provided")
)

func WithApiKey(apiKey string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: fmt.Sprintf("header:%s", apiKeyHeader),
		Validator: func(key string, ctx echo.Context) (bool, error) {
			if key != apiKey {
				return false, errWrongApiKeyProvided
			}
			return true, nil
		},
		ErrorHandler: func(err error, ctx echo.Context) error {
			return ctx.JSON(http.StatusUnauthorized, services.NewErrorResponse(err))
		},
	})
}
