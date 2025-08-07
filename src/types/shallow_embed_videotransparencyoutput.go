package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ShallowVideoTransparencyOutput struct {
	ShallowModel
	ID 				VideoTransparancyOutputID `json:"id"`
	StatsModel 		string `json:"stats_model"`
	FilesArrayModel []string `json:"files_array_model"`
}

func (c *ShallowVideoTransparencyOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowVideoTransparencyOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowVideoTransparencyOutput) New() ShallowVideoTransparencyOutput {
	c.ID = VideoTransparancyOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}

