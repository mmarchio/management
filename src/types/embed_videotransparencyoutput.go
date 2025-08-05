package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoTransparancyOutput struct {
	EmbedModel
	ID 				VideoTransparancyOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
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
	c.StatsModel = c.StatsModel.New(c.ID.String())
	return c
}

