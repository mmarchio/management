package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/logger"
)

func GetEchoCtx(c echo.Context) context.Context {
	ctxInterface := c.Get("context")
	if ctx, ok := ctxInterface.(context.Context); ok {
		return ctx
	}
	return nil
}

func SetEchoCtx(c echo.Context, ctx context.Context) echo.Context {
	c.Set("context", ctx)
	return c
}

func GetLogger(ctx context.Context) logger.LoggerFuncT {
	if fn, ok := ctx.Value(logger.LoggerKey).(logger.LoggerFuncT); ok {
		return fn
	}
	return nil
}