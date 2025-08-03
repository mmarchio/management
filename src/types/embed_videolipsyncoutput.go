package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoLipsyncOutput struct {
	EmbedModel
	ID VideoLipsyncOutputID `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

func (c *VideoLipsyncOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoLipsyncOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoLipsyncOutput) New() VideoLipsyncOutput {
	c.ID = VideoLipsyncOutputID(uuid.NewString())
	c.Stats = c.Stats.New(c.ID.String())
	return c
}

