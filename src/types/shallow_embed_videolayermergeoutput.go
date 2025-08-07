package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ShallowVideoLayerMergeOutput struct {
	ShallowModel
	ID 				VideoLayerMergeOutputID `json:"id"`
	StatsModel 		string `json:"stats_model"`
	FilesArrayModel []string `json:"files_array_model"`
}

func (c *ShallowVideoLayerMergeOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowVideoLayerMergeOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowVideoLayerMergeOutput) New() ShallowVideoLayerMergeOutput {
	c.ID = VideoLayerMergeOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}
