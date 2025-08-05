package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoLayerMergeOutput struct {
	EmbedModel
	ID 				VideoLayerMergeOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
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
	c.StatsModel = c.StatsModel.New(c.ID.String())
	return c
}
