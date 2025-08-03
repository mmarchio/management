package types

import (
	"context"
	"encoding/json"
)

type Character struct{}

func (c *Character) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Character) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
