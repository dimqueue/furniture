package middlewares

import (
	"github.com/dmytroserhiienko02/furniture/api-svc/internal/data"
	"github.com/labstack/echo/v4"
	"gitlab.com/distributed_lab/logan/v3"
)

const (
	loggerCtxKey = "logger"
	repository   = "repository"
)

func CtxLog(log *logan.Entry) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(loggerCtxKey, log.WithFields(logan.F{
				"method": ctx.Request().Method,
				"path":   ctx.Request().URL.Path,
			}))
			return next(ctx)
		}
	}
}

func Log(ctx echo.Context) *logan.Entry {
	return ctx.Get(loggerCtxKey).(*logan.Entry)
}

func CtxRepository(r data.IRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(repository, r)
			return next(ctx)
		}
	}
}

func Repository(ctx echo.Context) data.IRepository {
	return ctx.Get(repository).(data.IRepository)
}

func AdminUUID(ctx echo.Context) string {
	return ctx.Request().Header.Get(uuidKeyHeader)
}
