package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ImageBackgroundContextOutput struct {
	EmbedModel
	ID ImageBackgroundContextOutputID `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

func (c *ImageBackgroundContextOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ImageBackgroundContextOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ImageBackgroundContextOutput) New() ImageBackgroundContextOutput {
	c.ID = ImageBackgroundContextOutputID(uuid.NewString())
	c.Stats = c.Stats.New(c.ID.String())
	return c
}