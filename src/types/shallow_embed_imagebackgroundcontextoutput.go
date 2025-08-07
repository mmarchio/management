package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ShallowImageBackgroundContextOutput struct {
	ShallowModel
	ID 					ImageBackgroundContextOutputID `json:"id"`
	StatsModel 			string `json:"stats_model"`
	FilesArrayModel 	[]string `json:"files_array_model"`
}

func (c *ShallowImageBackgroundContextOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowImageBackgroundContextOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowImageBackgroundContextOutput) New() ShallowImageBackgroundContextOutput {
	c.ID = ImageBackgroundContextOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}