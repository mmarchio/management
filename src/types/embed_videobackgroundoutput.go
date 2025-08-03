package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoBackgroundOutput struct {
	EmbedModel
	ID VideoBackgroundOutputID `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

func (c *VideoBackgroundOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoBackgroundOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoBackgroundOutput) New() VideoBackgroundOutput {
	c.ID = VideoBackgroundOutputID(uuid.NewString())
	c.Stats = c.Stats.New(c.ID.String())
	return c
}
