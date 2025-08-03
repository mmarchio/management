package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoLayerMergeOutput struct {
	EmbedModel
	ID VideoLayerMergeOutputID `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

func (c *VideoLayerMergeOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoLayerMergeOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoLayerMergeOutput) New() VideoLayerMergeOutput {
	c.ID = VideoLayerMergeOutputID(uuid.NewString())
	c.Stats = c.Stats.New(c.ID.String())
	return c
}
