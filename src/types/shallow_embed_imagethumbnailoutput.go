package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ShallowImageThumbnailOutput struct {
	ShallowModel
	ID 					ImageThumbnailOutputID `json:"id"`
	StatsModel 			string `json:"stats_model"`
	FilesArrayModel 	[]string `json:"files_array_model"`
}

func (c *ShallowImageThumbnailOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowImageThumbnailOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowImageThumbnailOutput) New() ShallowImageThumbnailOutput {
	c.ID = ImageThumbnailOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}
