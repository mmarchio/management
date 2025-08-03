package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
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