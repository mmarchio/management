package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ShallowVideoLipsyncOutput struct {
	ShallowModel
	ID 				VideoLipsyncOutputID `json:"id"`
	StatsModel 		string `json:"stats_model"`
	FilesArrayModel []string `json:"files_array_model"`
}

func (c *ShallowVideoLipsyncOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowVideoLipsyncOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowVideoLipsyncOutput) New() ShallowVideoLipsyncOutput {
	c.ID = VideoLipsyncOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}

