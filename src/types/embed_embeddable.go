package types

import "context"

type Embeddable interface{
	Marshal(context.Context) (string, error)
	GetContentType() string
	GetID() string
}