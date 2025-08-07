package types

import (
	"context"
	"encoding/json"
)

type ShallowVideoOutput struct {
	EmbedModel
	ID 				VideoOutputID `json:"id"`
	StatsModel 		string `json:"stats_model"`
	FilesArrayModel []string `json:"files_model"`
}

func (c *ShallowVideoOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowVideoOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return string(b), err
}


