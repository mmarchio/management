package models

import (
	"github.com/labstack/echo/v4"
)

type ITable interface {
	Scan(ctx echo.Context, rows Scannable) (ITable, error) 
	Values(ctx echo.Context) ([]any, error)
	GetID() string
	GetContentType() string
}

type Scannable interface {
	Scan(...any) error
	Next() bool
	Err() error
	Close() error
}
