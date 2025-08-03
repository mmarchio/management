package types

import (
	"context"
	"encoding/json"

)

type AudioOutput struct {
	EmbedModel
	ID AudioOutputID `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

func (c *AudioOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c AudioOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
