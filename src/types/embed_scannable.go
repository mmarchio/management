package types

type Scannable interface {
	Scan(...any) error
	Next() bool
	Err() error
	Close() error
}
