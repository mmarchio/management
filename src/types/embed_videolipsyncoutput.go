package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoLipsyncOutput struct {
	EmbedModel
	ID 				VideoLipsyncOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
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
	c.StatsModel = c.StatsModel.New(nil)
	return c
}

