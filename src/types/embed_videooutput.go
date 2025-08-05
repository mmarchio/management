package types

import (
	"context"
	"encoding/json"

)

type VideoOutput struct {
	EmbedModel
	ID 				VideoOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_model"`
}

func (c *VideoOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	c.StatsModel = c.StatsModel.New(c.ID.String())
	return string(b), err
}


