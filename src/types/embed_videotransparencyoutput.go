package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoTransparancyOutput struct {
	EmbedModel
	ID VideoTransparancyOutputID `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

func (c *VideoTransparancyOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoTransparancyOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoTransparancyOutput) New() VideoTransparancyOutput {
	c.ID = VideoTransparancyOutputID(uuid.NewString())
	c.Stats = c.Stats.New(c.ID.String())
	return c
}

