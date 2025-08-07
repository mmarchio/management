package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ShallowVideoBackgroundOutput struct {
	ShallowModel
	ID 				VideoBackgroundOutputID `json:"id"`
	StatsModel 		string `json:"stats_model"`
	FilesArrayModel []string `json:"files_array_model"`
}

func (c *ShallowVideoBackgroundOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowVideoBackgroundOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowVideoBackgroundOutput) New() ShallowVideoBackgroundOutput {
	c.ID = VideoBackgroundOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}
