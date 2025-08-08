package models

import "context"

type ITable interface {
	Scan(ctx context.Context, rows Scannable) (ITable, error) 
	Values(ctx context.Context) ([]any, error)
	GetID() string
	GetContentType() string
}

type Scannable interface {
	Scan(...any) error
	Next() bool
	Err() error
	Close() error
}
