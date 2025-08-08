package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoTransparencyOutput struct {
	EmbedModel
	ID 				VideoTransparencyOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
}

func (c *VideoTransparencyOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoTransparencyOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoTransparencyOutput) New() VideoTransparencyOutput {
	c.ID = VideoTransparencyOutputID(uuid.NewString())
	c.StatsModel = c.StatsModel.New(nil)
	return c
}

