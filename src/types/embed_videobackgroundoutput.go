package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoBackgroundOutput struct {
	EmbedModel
	ID 				VideoBackgroundOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
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
	c.StatsModel = c.StatsModel.New(nil)
	return c
}
